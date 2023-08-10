// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type ComplianceIntegrationsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestComplianceIntegrationsStore(t *testing.T) {
	suite.Run(t, new(ComplianceIntegrationsStoreSuite))
}

func (s *ComplianceIntegrationsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ComplianceIntegrationsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_integrations CASCADE")
	s.T().Log("compliance_integrations", tag)
	s.NoError(err)
}

func (s *ComplianceIntegrationsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ComplianceIntegrationsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceIntegration := &storage.ComplianceIntegration{}
	s.NoError(testutils.FullInit(complianceIntegration, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceIntegration, exists, err := store.Get(ctx, complianceIntegration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceIntegration)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceIntegration))
	foundComplianceIntegration, exists, err = store.Get(ctx, complianceIntegration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceIntegration, foundComplianceIntegration)

	complianceIntegrationCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, complianceIntegrationCount)
	complianceIntegrationCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(complianceIntegrationCount)

	complianceIntegrationExists, err := store.Exists(ctx, complianceIntegration.GetId())
	s.NoError(err)
	s.True(complianceIntegrationExists)
	s.NoError(store.Upsert(ctx, complianceIntegration))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceIntegration), sac.ErrResourceAccessDenied)

	foundComplianceIntegration, exists, err = store.Get(ctx, complianceIntegration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceIntegration, foundComplianceIntegration)

	s.NoError(store.Delete(ctx, complianceIntegration.GetId()))
	foundComplianceIntegration, exists, err = store.Get(ctx, complianceIntegration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceIntegration)
	s.ErrorIs(store.Delete(withNoAccessCtx, complianceIntegration.GetId()), sac.ErrResourceAccessDenied)

	var complianceIntegrations []*storage.ComplianceIntegration
	var complianceIntegrationIDs []string
	for i := 0; i < 200; i++ {
		complianceIntegration := &storage.ComplianceIntegration{}
		s.NoError(testutils.FullInit(complianceIntegration, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceIntegrations = append(complianceIntegrations, complianceIntegration)
		complianceIntegrationIDs = append(complianceIntegrationIDs, complianceIntegration.GetId())
	}

	s.NoError(store.UpsertMany(ctx, complianceIntegrations))
	allComplianceIntegration, err := store.GetAll(ctx)
	s.NoError(err)
	s.ElementsMatch(complianceIntegrations, allComplianceIntegration)

	complianceIntegrationCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, complianceIntegrationCount)

	s.NoError(store.DeleteMany(ctx, complianceIntegrationIDs))

	complianceIntegrationCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(0, complianceIntegrationCount)
}