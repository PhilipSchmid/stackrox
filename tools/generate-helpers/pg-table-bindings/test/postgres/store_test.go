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
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type SinglekeyStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestSinglekeyStore(t *testing.T) {
	suite.Run(t, new(SinglekeyStoreSuite))
}

func (s *SinglekeyStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *SinglekeyStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *SinglekeyStoreSuite) TestStore() {
	ctx := context.Background()

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	testSingleKeyStruct := &storage.TestSingleKeyStruct{}
	s.NoError(testutils.FullInit(testSingleKeyStruct, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundTestSingleKeyStruct, exists, err := store.Get(ctx, testSingleKeyStruct.GetKey())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestSingleKeyStruct)

	s.NoError(store.Upsert(ctx, testSingleKeyStruct))
	foundTestSingleKeyStruct, exists, err = store.Get(ctx, testSingleKeyStruct.GetKey())
	s.NoError(err)
	s.True(exists)
	s.Equal(testSingleKeyStruct, foundTestSingleKeyStruct)

	testSingleKeyStructCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(testSingleKeyStructCount, 1)

	testSingleKeyStructExists, err := store.Exists(ctx, testSingleKeyStruct.GetKey())
	s.NoError(err)
	s.True(testSingleKeyStructExists)
	s.NoError(store.Upsert(ctx, testSingleKeyStruct))

	foundTestSingleKeyStruct, exists, err = store.Get(ctx, testSingleKeyStruct.GetKey())
	s.NoError(err)
	s.True(exists)
	s.Equal(testSingleKeyStruct, foundTestSingleKeyStruct)

	s.NoError(store.Delete(ctx, testSingleKeyStruct.GetKey()))
	foundTestSingleKeyStruct, exists, err = store.Get(ctx, testSingleKeyStruct.GetKey())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestSingleKeyStruct)

	var testSingleKeyStructs []*storage.TestSingleKeyStruct
	for i := 0; i < 200; i++ {
		testSingleKeyStruct := &storage.TestSingleKeyStruct{}
		s.NoError(testutils.FullInit(testSingleKeyStruct, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		testSingleKeyStructs = append(testSingleKeyStructs, testSingleKeyStruct)
	}
	s.NoError(store.UpsertMany(ctx, testSingleKeyStructs))

	testSingleKeyStructCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(testSingleKeyStructCount, 200)
}
