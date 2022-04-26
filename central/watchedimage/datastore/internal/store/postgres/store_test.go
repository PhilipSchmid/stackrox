// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type WatchedimagesStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestWatchedimagesStore(t *testing.T) {
	suite.Run(t, new(WatchedimagesStoreSuite))
}

func (s *WatchedimagesStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *WatchedimagesStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *WatchedimagesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	watchedImage := &storage.WatchedImage{}
	s.NoError(testutils.FullInit(watchedImage, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundWatchedImage, exists, err := store.Get(ctx, watchedImage.GetName())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundWatchedImage)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, watchedImage))
	foundWatchedImage, exists, err = store.Get(ctx, watchedImage.GetName())
	s.NoError(err)
	s.True(exists)
	s.Equal(watchedImage, foundWatchedImage)

	watchedImageCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(watchedImageCount, 1)
	watchedImageCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(watchedImageCount)

	watchedImageExists, err := store.Exists(ctx, watchedImage.GetName())
	s.NoError(err)
	s.True(watchedImageExists)
	s.NoError(store.Upsert(ctx, watchedImage))
	s.ErrorIs(store.Upsert(withNoAccessCtx, watchedImage), sac.ErrResourceAccessDenied)

	foundWatchedImage, exists, err = store.Get(ctx, watchedImage.GetName())
	s.NoError(err)
	s.True(exists)
	s.Equal(watchedImage, foundWatchedImage)

	s.NoError(store.Delete(ctx, watchedImage.GetName()))
	foundWatchedImage, exists, err = store.Get(ctx, watchedImage.GetName())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundWatchedImage)
	s.ErrorIs(store.Delete(withNoAccessCtx, watchedImage.GetName()), sac.ErrResourceAccessDenied)

	var watchedImages []*storage.WatchedImage
	for i := 0; i < 200; i++ {
		watchedImage := &storage.WatchedImage{}
		s.NoError(testutils.FullInit(watchedImage, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		watchedImages = append(watchedImages, watchedImage)
	}
	s.NoError(store.UpsertMany(ctx, watchedImages))

	watchedImageCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(watchedImageCount, 200)
}
