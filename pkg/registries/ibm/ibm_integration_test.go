package ibm

import (
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/images/types"
	"github.com/stackrox/rox/pkg/images/utils"
	"github.com/stretchr/testify/require"
)

const (
	testImage = "us.gcr.io/sr-testing/nginx:1.10"
	apiToken  = "Z_t3ZI1AcDB_513s91kHw_RXpGVcY-GFUQLLx-UwZqzB"
)

func TestIBM(t *testing.T) {
	reg, err := newRegistry(&storage.ImageIntegration{
		IntegrationConfig: &storage.ImageIntegration_Ibm{
			Ibm: &storage.IBMRegistryConfig{
				Endpoint: "us.icr.io",
				ApiKey:   apiToken,
			},
		},
	})
	require.NoError(t, err)

	image, err := utils.GenerateImageFromString(testImage)
	require.NoError(t, err)

	metadata, err := reg.Metadata(types.ToImage(image))
	require.NoError(t, err)
	require.NotNil(t, metadata)
}
