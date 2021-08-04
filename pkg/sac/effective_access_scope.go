package sac

import (
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	"k8s.io/apimachinery/pkg/labels"
)

// ScopeState represents possible states of a scope.
type ScopeState int

const (
	// Excluded indicates that neither the scope nor its subtree are included.
	Excluded ScopeState = iota
	// Partial indicates that the scope is not included in its entirety but some
	// of its children are included.
	Partial
	// Included indicates that the scope is included with its subtree.
	Included

	// FQSN stands for "fully qualified scope name".
	clusterNameLabel   = "stackrox.io/authz.metadata.cluster.fqsn"
	namespaceNameLabel = "stackrox.io/authz.metadata.namespace.fqsn"
	scopeSeparator     = "::"
)

var (
	log = logging.LoggerForModule()
)

// EffectiveAccessScopeTree is a tree of scopes with their states.
type EffectiveAccessScopeTree struct {
	State           ScopeState
	Clusters        map[string]*ClustersScopeSubTree
	clusterIDToName map[string]string
}

// ClustersScopeSubTree is a subtree of cluster scopes with their states. Extras
// field can be used by clients to augment the tree with additional info like
// cluster id, labels, etc.
type ClustersScopeSubTree struct {
	State      ScopeState
	Namespaces map[string]*NamespacesScopeSubTree
	Extras     interface{}
}

// NamespacesScopeSubTree is a subtree of namespace scopes with their states.
// Extras field can be used by clients to augment the tree with additional info
// like namespace id, labels, etc.
type NamespacesScopeSubTree struct {
	State  ScopeState
	Extras interface{}
}

// UnrestrictedEffectiveAccessScope returns EffectiveAccessScopeTree
// allowing everything implicitly via marking the root Included.
func UnrestrictedEffectiveAccessScope() *EffectiveAccessScopeTree {
	return newEffectiveAccessScopeTree(Included)
}

// ComputeEffectiveAccessScope applies a simple access scope to provided
// clusters and namespaces and yields EffectiveAccessScopeTree. Empty access
// scope rules mean nothing is included.
func ComputeEffectiveAccessScope(scopeRules *storage.SimpleAccessScope_Rules, clusters []*storage.Cluster, namespaces []*storage.NamespaceMetadata, detail v1.ComputeEffectiveAccessScopeRequest_Detail) (*EffectiveAccessScopeTree, error) {
	root := newEffectiveAccessScopeTree(Excluded)

	// Compile scope into cluster and namespace selectors.
	clusterSelectors, namespaceSelectors, err := convertRulesToLabelSelectors(scopeRules)
	if err != nil {
		return nil, err
	}

	// Check every cluster against corresponding access scope rules represented
	// by clusterSelectors (note cluster name to label conversion). Partial
	// state is not possible here yet.
	for _, cluster := range clusters {
		root.populateStateForCluster(cluster, clusterSelectors, detail)
	}

	// Check every namespace not indirectly included by its parent cluster
	// against corresponding access scope rules represented by
	// namespaceSelectors (note namespace name to label conversion).
	for _, namespace := range namespaces {
		clusterName := namespace.GetClusterName()
		namespaceFQSN := getNamespaceFQSN(clusterName, namespace.GetName())

		// If parent cluster is unknown, log and add cluster as Excluded.
		parentCluster := root.Clusters[clusterName]
		if parentCluster == nil {
			log.Warnf("namespace %q belongs to unknown cluster %q", namespaceFQSN, clusterName)
			parentCluster = newClusterScopeSubTreeWithExtras(Excluded, EffectiveAccessScopeTreeExtras{Name: clusterName})
			root.Clusters[clusterName] = parentCluster
		}

		parentCluster.populateStateForNamespace(namespace, namespaceSelectors, detail)
	}

	root.bubbleUpStatesAndCompactify(detail)

	return root, nil
}

// GetClusterByID returns ClusterScopeSubTree for given cluster ID.
// Returns nil when clusterID is not known.
func (root *EffectiveAccessScopeTree) GetClusterByID(clusterID string) *ClustersScopeSubTree {
	return root.Clusters[root.clusterIDToName[clusterID]]
}

// populateStateForCluster adds given cluster as Included or Excluded to root.
// Only the last observed cluster is considered if multiple ones with the same
// name exist.
func (root *EffectiveAccessScopeTree) populateStateForCluster(cluster *storage.Cluster, clusterSelectors []labels.Selector, detail v1.ComputeEffectiveAccessScopeRequest_Detail) {
	clusterName := cluster.GetName()

	extras := extrasForCluster(cluster, detail)

	// There is no need to check if root is Included as we start with Excluded root.
	// If it will be Included then we can include the cluster and short-circuit:
	// no need to match if parent is included.

	// Augment cluster labels with cluster's name.
	clusterLabels := augmentLabels(cluster.GetLabels(), clusterNameLabel, clusterName)

	// Match and update the tree.
	matched := matchLabels(clusterSelectors, clusterLabels)
	root.Clusters[clusterName] = newClusterScopeSubTreeWithExtras(matched, extras)
	root.clusterIDToName[cluster.GetId()] = clusterName
}

// bubbleUpStatesAndCompactify updates the state of parent nodes based on the
// state of their children and compactifies the tree iff the requested level of
// detail is MINIMAL.
//
// If any child is Included or Partial, its parent becomes at least Partial. If
// all children are Included, the parent is still Partial unless it has been
// included directly.
//
// For MINIMAL level of detail, delete from the tree:
//   * subtrees *with roots* in the Excluded state,
//   * subtrees *of nodes* in the Included state.
func (root *EffectiveAccessScopeTree) bubbleUpStatesAndCompactify(detail v1.ComputeEffectiveAccessScopeRequest_Detail) {
	deleteUnnecessaryNodes := detail == v1.ComputeEffectiveAccessScopeRequest_MINIMAL
	for clusterName, cluster := range root.Clusters {
		for namespaceName, namespace := range cluster.Namespaces {
			// Update the cluster's state from Excluded to Partial
			// if any of its namespaces is included.
			if cluster.State == Excluded &&
				(namespace.State == Included || namespace.State == Partial) {
				cluster.State = Partial

				// If we don't need to delete nodes, we can short-circuit.
				if !deleteUnnecessaryNodes {
					break
				}
			}
			// Delete Excluded namespaces if desired.
			if deleteUnnecessaryNodes && namespace.State == Excluded {
				delete(cluster.Namespaces, namespaceName)
			}
		}

		// Delete all namespaces for Included clusters and Excluded clusters
		// if desired.
		if deleteUnnecessaryNodes {
			if cluster.State == Included {
				cluster.Namespaces = nil
			} else if cluster.State == Excluded {
				delete(root.Clusters, clusterName)
			}
		}

		// Update the root's state from to Partial if any cluster is included.
		if root.State == Excluded && (cluster.State == Included || cluster.State == Partial) {
			root.State = Partial
		}
	}
}

// populateStateForNamespace adds given namespace as Included or Excluded to
// parent cluster. Only the last observed namespace is considered if multiple
// ones with the same <cluster name, namespace name> exist.
func (cluster *ClustersScopeSubTree) populateStateForNamespace(namespace *storage.NamespaceMetadata, namespaceSelectors []labels.Selector, detail v1.ComputeEffectiveAccessScopeRequest_Detail) {
	clusterName := namespace.GetClusterName()
	namespaceName := namespace.GetName()
	namespaceFQSN := getNamespaceFQSN(clusterName, namespaceName)

	// If parent is Included, include the namespace and short-circuit:
	// no need to match if parent is included.
	if cluster.State == Included {
		cluster.Namespaces[namespaceName] = newNamespacesScopeSubTree(Included, extrasForNamespace(namespace, detail))
		return
	}

	// Augment namespace labels with namespace's FQSN.
	namespaceLabels := augmentLabels(namespace.GetLabels(), namespaceNameLabel, namespaceFQSN)

	// Match and update the tree.
	matched := matchLabels(namespaceSelectors, namespaceLabels)
	cluster.Namespaces[namespaceName] = newNamespacesScopeSubTree(matched, extrasForNamespace(namespace, detail))
}

func newEffectiveAccessScopeTree(state ScopeState) *EffectiveAccessScopeTree {
	return &EffectiveAccessScopeTree{
		State:           state,
		Clusters:        make(map[string]*ClustersScopeSubTree),
		clusterIDToName: make(map[string]string),
	}
}

func newClusterScopeSubTreeWithExtras(state ScopeState, extras EffectiveAccessScopeTreeExtras) *ClustersScopeSubTree {
	return &ClustersScopeSubTree{
		State:      state,
		Namespaces: make(map[string]*NamespacesScopeSubTree),
		Extras:     &extras,
	}
}

func newNamespacesScopeSubTree(state ScopeState, extras EffectiveAccessScopeTreeExtras) *NamespacesScopeSubTree {
	return &NamespacesScopeSubTree{
		State:  state,
		Extras: &extras,
	}
}

func getNamespaceFQSN(cluster string, namespace string) string {
	return cluster + scopeSeparator + namespace
}

func augmentLabels(labels map[string]string, key string, value string) map[string]string {
	result := make(map[string]string)
	for k, v := range labels {
		result[k] = v
	}
	result[key] = value

	return result
}

// matchLabels checks if any of the given selectors matches the given label map.
func matchLabels(selectors []labels.Selector, lbls map[string]string) ScopeState {
	for _, s := range selectors {
		if s.Matches(labels.Set(lbls)) {
			return Included
		}
	}
	return Excluded
}
