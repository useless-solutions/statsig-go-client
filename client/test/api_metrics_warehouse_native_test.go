/*
Console API

Testing MetricsWarehouseNativeAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_MetricsWarehouseNativeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetricsWarehouseNativeAPIService ConsoleV1MetricsControllerGenCreateMetricSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenCreateMetricSource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsWarehouseNativeAPIService ConsoleV1MetricsControllerGenDeleteMetricSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenDeleteMetricSource(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsWarehouseNativeAPIService ConsoleV1MetricsControllerGenListMetricSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenListMetricSources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsWarehouseNativeAPIService ConsoleV1MetricsControllerGenReadMetricSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenReadMetricSource(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsWarehouseNativeAPIService ConsoleV1MetricsControllerReloadMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerReloadMetric(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsWarehouseNativeAPIService ConsoleV1MetricsControllerUpdateMetricSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerUpdateMetricSource(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
