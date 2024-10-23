/*
Console API

Testing GatesAPIService

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

func Test_openapi_GatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GatesAPIService ConsoleV1GateOverridesControllerGenAdd", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenAdd(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GateOverridesControllerGenRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenRead(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GateOverridesControllerGenRemove", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenRemove(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GateOverridesControllerGenUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenArchive(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenDisable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenDisable(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenEnable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenEnable(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenFullUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenFullUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenLaunch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenLaunch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenLoadPulse", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenLoadPulse(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenMultiRuleAdd", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenMultiRuleAdd(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenMultiRuleUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenMultiRuleUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenPartialUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenPulseLoadHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var ruleID string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenPulseLoadHistory(context.Background(), id, ruleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenPulseResults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var ruleID string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenPulseResults(context.Background(), id, ruleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRead(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenReadRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenReadRules(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenRemove", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRemove(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenRuleAdd", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRuleAdd(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenRuleDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var ruleID string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRuleDelete(context.Background(), id, ruleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatesAPIService ConsoleV1GatesControllerGenRuleUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var ruleID string

		resp, httpRes, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRuleUpdate(context.Background(), id, ruleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
