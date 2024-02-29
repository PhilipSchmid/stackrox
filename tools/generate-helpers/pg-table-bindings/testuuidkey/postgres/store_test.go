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

type TestSingleUUIDKeyStructsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestTestSingleUUIDKeyStructsStore(t *testing.T) {
	suite.Run(t, new(TestSingleUUIDKeyStructsStoreSuite))
}

func (s *TestSingleUUIDKeyStructsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *TestSingleUUIDKeyStructsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE test_single_uuid_key_structs CASCADE")
	s.T().Log("test_single_uuid_key_structs", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *TestSingleUUIDKeyStructsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *TestSingleUUIDKeyStructsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	testSingleUUIDKeyStruct := &storage.TestSingleUUIDKeyStruct{}
	s.NoError(testutils.FullInit(testSingleUUIDKeyStruct, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundTestSingleUUIDKeyStruct, exists, err := store.Get(ctx, testSingleUUIDKeyStruct.GetKey())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestSingleUUIDKeyStruct)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, testSingleUUIDKeyStruct))
	foundTestSingleUUIDKeyStruct, exists, err = store.Get(ctx, testSingleUUIDKeyStruct.GetKey())
	s.NoError(err)
	s.True(exists)
	s.Equal(testSingleUUIDKeyStruct, foundTestSingleUUIDKeyStruct)

	testSingleUUIDKeyStructCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, testSingleUUIDKeyStructCount)
	testSingleUUIDKeyStructCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(testSingleUUIDKeyStructCount)

	testSingleUUIDKeyStructExists, err := store.Exists(ctx, testSingleUUIDKeyStruct.GetKey())
	s.NoError(err)
	s.True(testSingleUUIDKeyStructExists)
	s.NoError(store.Upsert(ctx, testSingleUUIDKeyStruct))
	s.ErrorIs(store.Upsert(withNoAccessCtx, testSingleUUIDKeyStruct), sac.ErrResourceAccessDenied)

	foundTestSingleUUIDKeyStruct, exists, err = store.Get(ctx, testSingleUUIDKeyStruct.GetKey())
	s.NoError(err)
	s.True(exists)
	s.Equal(testSingleUUIDKeyStruct, foundTestSingleUUIDKeyStruct)

	s.NoError(store.Delete(ctx, testSingleUUIDKeyStruct.GetKey()))
	foundTestSingleUUIDKeyStruct, exists, err = store.Get(ctx, testSingleUUIDKeyStruct.GetKey())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestSingleUUIDKeyStruct)
	s.NoError(store.Delete(withNoAccessCtx, testSingleUUIDKeyStruct.GetKey()))

	var testSingleUUIDKeyStructs []*storage.TestSingleUUIDKeyStruct
	var testSingleUUIDKeyStructIDs []string
	for i := 0; i < 200; i++ {
		testSingleUUIDKeyStruct := &storage.TestSingleUUIDKeyStruct{}
		s.NoError(testutils.FullInit(testSingleUUIDKeyStruct, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		testSingleUUIDKeyStructs = append(testSingleUUIDKeyStructs, testSingleUUIDKeyStruct)
		testSingleUUIDKeyStructIDs = append(testSingleUUIDKeyStructIDs, testSingleUUIDKeyStruct.GetKey())
	}

	s.NoError(store.UpsertMany(ctx, testSingleUUIDKeyStructs))
	allTestSingleUUIDKeyStruct, err := store.GetAll(ctx)
	s.NoError(err)
	s.ElementsMatch(testSingleUUIDKeyStructs, allTestSingleUUIDKeyStruct)

	testSingleUUIDKeyStructCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, testSingleUUIDKeyStructCount)

	s.NoError(store.DeleteMany(ctx, testSingleUUIDKeyStructIDs))

	testSingleUUIDKeyStructCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, testSingleUUIDKeyStructCount)
}
