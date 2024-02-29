// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type NetworkpoliciesundodeploymentsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestNetworkpoliciesundodeploymentsStore(t *testing.T) {
	suite.Run(t, new(NetworkpoliciesundodeploymentsStoreSuite))
}

func (s *NetworkpoliciesundodeploymentsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *NetworkpoliciesundodeploymentsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE networkpoliciesundodeployments CASCADE")
	s.T().Log("networkpoliciesundodeployments", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *NetworkpoliciesundodeploymentsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *NetworkpoliciesundodeploymentsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	networkPolicyApplicationUndoDeploymentRecord := &storage.NetworkPolicyApplicationUndoDeploymentRecord{}
	s.NoError(testutils.FullInit(networkPolicyApplicationUndoDeploymentRecord, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err := store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkPolicyApplicationUndoDeploymentRecord)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, networkPolicyApplicationUndoDeploymentRecord))
	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkPolicyApplicationUndoDeploymentRecord, foundNetworkPolicyApplicationUndoDeploymentRecord)

	networkPolicyApplicationUndoDeploymentRecordCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, networkPolicyApplicationUndoDeploymentRecordCount)
	networkPolicyApplicationUndoDeploymentRecordCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(networkPolicyApplicationUndoDeploymentRecordCount)

	networkPolicyApplicationUndoDeploymentRecordExists, err := store.Exists(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.True(networkPolicyApplicationUndoDeploymentRecordExists)
	s.NoError(store.Upsert(ctx, networkPolicyApplicationUndoDeploymentRecord))
	s.ErrorIs(store.Upsert(withNoAccessCtx, networkPolicyApplicationUndoDeploymentRecord), sac.ErrResourceAccessDenied)

	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkPolicyApplicationUndoDeploymentRecord, foundNetworkPolicyApplicationUndoDeploymentRecord)

	s.NoError(store.Delete(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId()))
	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkPolicyApplicationUndoDeploymentRecord)
	s.NoError(store.Delete(withNoAccessCtx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId()))

	var networkPolicyApplicationUndoDeploymentRecords []*storage.NetworkPolicyApplicationUndoDeploymentRecord
	var networkPolicyApplicationUndoDeploymentRecordIDs []string
	for i := 0; i < 200; i++ {
		networkPolicyApplicationUndoDeploymentRecord := &storage.NetworkPolicyApplicationUndoDeploymentRecord{}
		s.NoError(testutils.FullInit(networkPolicyApplicationUndoDeploymentRecord, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		networkPolicyApplicationUndoDeploymentRecords = append(networkPolicyApplicationUndoDeploymentRecords, networkPolicyApplicationUndoDeploymentRecord)
		networkPolicyApplicationUndoDeploymentRecordIDs = append(networkPolicyApplicationUndoDeploymentRecordIDs, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	}

	s.NoError(store.UpsertMany(ctx, networkPolicyApplicationUndoDeploymentRecords))

	networkPolicyApplicationUndoDeploymentRecordCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, networkPolicyApplicationUndoDeploymentRecordCount)

	s.NoError(store.DeleteMany(ctx, networkPolicyApplicationUndoDeploymentRecordIDs))

	networkPolicyApplicationUndoDeploymentRecordCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, networkPolicyApplicationUndoDeploymentRecordCount)
}
