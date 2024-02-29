package search

import (
	"context"

	"github.com/stackrox/rox/central/networkpolicies/datastore/internal/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
)

// searcherImpl provides a search implementation for network policies.
type searcherImpl struct {
	store store.Store
}

// Count returns the number of search results from the query.
func (s *searcherImpl) Count(ctx context.Context, q *v1.Query) (int, error) {
	return s.store.Count(ctx, q)
}
