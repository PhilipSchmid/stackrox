// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	storage "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type ProcessbaselinesStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestProcessbaselinesStore(t *testing.T) {
	suite.Run(t, new(ProcessbaselinesStoreSuite))
}

func (s *ProcessbaselinesStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *ProcessbaselinesStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *ProcessbaselinesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	processBaseline := &storage.ProcessBaseline{}
	s.NoError(testutils.FullInit(processBaseline, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundProcessBaseline, exists, err := store.Get(ctx, processBaseline.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundProcessBaseline)

	withNoAccessCtx := sac.WithNoAccess(ctx)
	withAccessToDifferentNsCtx := sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(targetResource),
			sac.ClusterScopeKeys(processBaseline.GetKey().GetClusterId()),
			sac.NamespaceScopeKeys("unknown ns"),
		))
	withAccessCtx := sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(targetResource),
			sac.ClusterScopeKeys(processBaseline.GetKey().GetClusterId()),
			sac.NamespaceScopeKeys(processBaseline.GetKey().GetNamespace()),
		))
	withAccessToClusterCtx := sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(targetResource),
			sac.ClusterScopeKeys(processBaseline.GetKey().GetClusterId()),
		))
	withNoAccessToClusterCtx := sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(targetResource),
			sac.ClusterScopeKeys("unknown cluster"),
		))

	s.NoError(store.Upsert(ctx, processBaseline))
	foundProcessBaseline, exists, err = store.Get(ctx, processBaseline.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(processBaseline, foundProcessBaseline)

	processBaselineCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(processBaselineCount, 1)

	processBaselineExists, err := store.Exists(ctx, processBaseline.GetId())
	s.NoError(err)
	s.True(processBaselineExists)
	s.NoError(store.Upsert(ctx, processBaseline))
	s.ErrorIs(store.Upsert(withNoAccessCtx, processBaseline), sac.ErrResourceAccessDenied)
	s.ErrorIs(store.Upsert(withNoAccessToClusterCtx, processBaseline), sac.ErrResourceAccessDenied)
	s.ErrorIs(store.Upsert(withAccessToDifferentNsCtx, processBaseline), sac.ErrResourceAccessDenied)
	s.NoError(store.Upsert(withAccessCtx, processBaseline))
	s.NoError(store.Upsert(withAccessToClusterCtx, processBaseline))
	s.ErrorIs(store.UpsertMany(withAccessToDifferentNsCtx, []*storage.ProcessBaseline{processBaseline}), sac.ErrResourceAccessDenied)
	s.ErrorIs(store.UpsertMany(withNoAccessToClusterCtx, []*storage.ProcessBaseline{processBaseline}), sac.ErrResourceAccessDenied)
	s.NoError(store.UpsertMany(withAccessCtx, []*storage.ProcessBaseline{processBaseline}))
	s.NoError(store.UpsertMany(withAccessToClusterCtx, []*storage.ProcessBaseline{processBaseline}))

	foundProcessBaseline, exists, err = store.Get(ctx, processBaseline.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(processBaseline, foundProcessBaseline)

	s.NoError(store.Delete(ctx, processBaseline.GetId()))
	foundProcessBaseline, exists, err = store.Get(ctx, processBaseline.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundProcessBaseline)

	var processBaselines []*storage.ProcessBaseline
	for i := 0; i < 200; i++ {
		processBaseline := &storage.ProcessBaseline{}
		s.NoError(testutils.FullInit(processBaseline, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		processBaselines = append(processBaselines, processBaseline)
	}
	s.ErrorIs(store.UpsertMany(withAccessToDifferentNsCtx, processBaselines), sac.ErrResourceAccessDenied)
	s.NoError(store.UpsertMany(ctx, processBaselines))

	processBaselineCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(processBaselineCount, 200)
}
