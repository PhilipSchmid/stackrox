package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	migrationSchema "github.com/stackrox/rox/migrator/migrations/m_194_to_m_195_vuln_request_global_scope/schema"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
)

const (
	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	log    = logging.LoggerForModule()
	schema = migrationSchema.VulnerabilityRequestsSchema
)

type storeType = storage.VulnerabilityRequest

// Store is the interface to interact with the storage for storage.VulnerabilityRequest
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, id string) error
	DeleteMany(ctx context.Context, identifiers []string) error

	Exists(ctx context.Context, id string) (bool, error)

	Get(ctx context.Context, id string) (*storeType, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storeType, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return pgSearch.NewGenericStoreWithPermissionChecker[storeType, *storeType](
		db,
		schema,
		pkGetter,
		insertIntoVulnerabilityRequests,
		copyFromVulnerabilityRequests,
		nil,
		nil,
		sac.NewAnyGlobalResourceAllowedPermissionChecker(resources.VulnerabilityManagementRequests, resources.VulnerabilityManagementApprovals),
	)
}

// region Helper functions

func pkGetter(obj *storeType) string {
	return obj.GetId()
}

func insertIntoVulnerabilityRequests(batch *pgx.Batch, obj *storage.VulnerabilityRequest) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetName(),
		obj.GetTargetState(),
		obj.GetStatus(),
		obj.GetExpired(),
		obj.GetRequestor().GetName(),
		pgutils.NilOrTime(obj.GetCreatedAt()),
		pgutils.NilOrTime(obj.GetLastUpdated()),
		obj.GetScope().GetImageScope().GetRegistry(),
		obj.GetScope().GetImageScope().GetRemote(),
		obj.GetScope().GetImageScope().GetTag(),
		obj.GetDeferralReq().GetExpiry().GetExpiresWhenFixed(),
		pgutils.NilOrTime(obj.GetDeferralReq().GetExpiry().GetExpiresOn()),
		obj.GetDeferralReq().GetExpiry().GetExpiryType(),
		obj.GetCves().GetCves(),
		serialized,
	}

	finalStr := "INSERT INTO vulnerability_requests (Id, Name, TargetState, Status, Expired, Requestor_Name, CreatedAt, LastUpdated, Scope_ImageScope_Registry, Scope_ImageScope_Remote, Scope_ImageScope_Tag, DeferralReq_Expiry_ExpiresWhenFixed, DeferralReq_Expiry_ExpiresOn, DeferralReq_Expiry_ExpiryType, Cves_Cves, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, TargetState = EXCLUDED.TargetState, Status = EXCLUDED.Status, Expired = EXCLUDED.Expired, Requestor_Name = EXCLUDED.Requestor_Name, CreatedAt = EXCLUDED.CreatedAt, LastUpdated = EXCLUDED.LastUpdated, Scope_ImageScope_Registry = EXCLUDED.Scope_ImageScope_Registry, Scope_ImageScope_Remote = EXCLUDED.Scope_ImageScope_Remote, Scope_ImageScope_Tag = EXCLUDED.Scope_ImageScope_Tag, DeferralReq_Expiry_ExpiresWhenFixed = EXCLUDED.DeferralReq_Expiry_ExpiresWhenFixed, DeferralReq_Expiry_ExpiresOn = EXCLUDED.DeferralReq_Expiry_ExpiresOn, DeferralReq_Expiry_ExpiryType = EXCLUDED.DeferralReq_Expiry_ExpiryType, Cves_Cves = EXCLUDED.Cves_Cves, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetApprovers() {
		if err := insertIntoVulnerabilityRequestsApprovers(batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from vulnerability_requests_approvers where vulnerability_requests_Id = $1 AND idx >= $2"
	batch.Queue(query, obj.GetId(), len(obj.GetApprovers()))
	for childIndex, child := range obj.GetComments() {
		if err := insertIntoVulnerabilityRequestsComments(batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from vulnerability_requests_comments where vulnerability_requests_Id = $1 AND idx >= $2"
	batch.Queue(query, obj.GetId(), len(obj.GetComments()))
	return nil
}

func insertIntoVulnerabilityRequestsApprovers(batch *pgx.Batch, obj *storage.SlimUser, vulnerabilityRequestID string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		vulnerabilityRequestID,
		idx,
		obj.GetName(),
	}

	finalStr := "INSERT INTO vulnerability_requests_approvers (vulnerability_requests_Id, idx, Name) VALUES($1, $2, $3) ON CONFLICT(vulnerability_requests_Id, idx) DO UPDATE SET vulnerability_requests_Id = EXCLUDED.vulnerability_requests_Id, idx = EXCLUDED.idx, Name = EXCLUDED.Name"
	batch.Queue(finalStr, values...)

	return nil
}

func insertIntoVulnerabilityRequestsComments(batch *pgx.Batch, obj *storage.RequestComment, vulnerabilityRequestID string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		vulnerabilityRequestID,
		idx,
		obj.GetUser().GetName(),
	}

	finalStr := "INSERT INTO vulnerability_requests_comments (vulnerability_requests_Id, idx, User_Name) VALUES($1, $2, $3) ON CONFLICT(vulnerability_requests_Id, idx) DO UPDATE SET vulnerability_requests_Id = EXCLUDED.vulnerability_requests_Id, idx = EXCLUDED.idx, User_Name = EXCLUDED.User_Name"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromVulnerabilityRequests(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.VulnerabilityRequest) error {
	inputRows := make([][]interface{}, 0, batchSize)

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	deletes := make([]string, 0, batchSize)

	copyCols := []string{
		"id",
		"name",
		"targetstate",
		"status",
		"expired",
		"requestor_name",
		"createdat",
		"lastupdated",
		"scope_imagescope_registry",
		"scope_imagescope_remote",
		"scope_imagescope_tag",
		"deferralreq_expiry_expireswhenfixed",
		"deferralreq_expiry_expireson",
		"deferralreq_expiry_expirytype",
		"cves_cves",
		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{
			obj.GetId(),
			obj.GetName(),
			obj.GetTargetState(),
			obj.GetStatus(),
			obj.GetExpired(),
			obj.GetRequestor().GetName(),
			pgutils.NilOrTime(obj.GetCreatedAt()),
			pgutils.NilOrTime(obj.GetLastUpdated()),
			obj.GetScope().GetImageScope().GetRegistry(),
			obj.GetScope().GetImageScope().GetRemote(),
			obj.GetScope().GetImageScope().GetTag(),
			obj.GetDeferralReq().GetExpiry().GetExpiresWhenFixed(),
			pgutils.NilOrTime(obj.GetDeferralReq().GetExpiry().GetExpiresOn()),
			obj.GetDeferralReq().GetExpiry().GetExpiryType(),
			obj.GetCves().GetCves(),
			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = deletes[:0]

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"vulnerability_requests"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err := copyFromVulnerabilityRequestsApprovers(ctx, s, tx, obj.GetId(), obj.GetApprovers()...); err != nil {
			return err
		}
		if err := copyFromVulnerabilityRequestsComments(ctx, s, tx, obj.GetId(), obj.GetComments()...); err != nil {
			return err
		}
	}

	return nil
}

func copyFromVulnerabilityRequestsApprovers(ctx context.Context, _ pgSearch.Deleter, tx *postgres.Tx, vulnerabilityRequestID string, objs ...*storage.SlimUser) error {
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"vulnerability_requests_id",
		"idx",
		"name",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			vulnerabilityRequestID,
			idx,
			obj.GetName(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"vulnerability_requests_approvers"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return nil
}

func copyFromVulnerabilityRequestsComments(ctx context.Context, _ pgSearch.Deleter, tx *postgres.Tx, vulnerabilityRequestID string, objs ...*storage.RequestComment) error {
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"vulnerability_requests_id",
		"idx",
		"user_name",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			vulnerabilityRequestID,
			idx,
			obj.GetUser().GetName(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"vulnerability_requests_comments"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return nil
}

// endregion Helper functions
