# \GatesAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1GateOverridesControllerGenAdd**](GatesAPI.md#ConsoleV1GateOverridesControllerGenAdd) | **Patch** /console/v1/gates/{id}/overrides | Add Gate Overrides
[**ConsoleV1GateOverridesControllerGenRead**](GatesAPI.md#ConsoleV1GateOverridesControllerGenRead) | **Get** /console/v1/gates/{id}/overrides | Get Gate Override
[**ConsoleV1GateOverridesControllerGenRemove**](GatesAPI.md#ConsoleV1GateOverridesControllerGenRemove) | **Delete** /console/v1/gates/{id}/overrides | Delete Gate Overrides
[**ConsoleV1GateOverridesControllerGenUpdate**](GatesAPI.md#ConsoleV1GateOverridesControllerGenUpdate) | **Post** /console/v1/gates/{id}/overrides | Update Gate Overrides
[**ConsoleV1GatesControllerGenArchive**](GatesAPI.md#ConsoleV1GatesControllerGenArchive) | **Put** /console/v1/gates/{id}/archive | Archive Gate
[**ConsoleV1GatesControllerGenCreate**](GatesAPI.md#ConsoleV1GatesControllerGenCreate) | **Post** /console/v1/gates | Create Gate
[**ConsoleV1GatesControllerGenDisable**](GatesAPI.md#ConsoleV1GatesControllerGenDisable) | **Put** /console/v1/gates/{id}/disable | Disable Gate
[**ConsoleV1GatesControllerGenEnable**](GatesAPI.md#ConsoleV1GatesControllerGenEnable) | **Put** /console/v1/gates/{id}/enable | Enable Gate
[**ConsoleV1GatesControllerGenFullUpdate**](GatesAPI.md#ConsoleV1GatesControllerGenFullUpdate) | **Post** /console/v1/gates/{id} | Fully Update Gates
[**ConsoleV1GatesControllerGenLaunch**](GatesAPI.md#ConsoleV1GatesControllerGenLaunch) | **Put** /console/v1/gates/{id}/launch | Launch Gate
[**ConsoleV1GatesControllerGenList**](GatesAPI.md#ConsoleV1GatesControllerGenList) | **Get** /console/v1/gates | List Gates
[**ConsoleV1GatesControllerGenLoadPulse**](GatesAPI.md#ConsoleV1GatesControllerGenLoadPulse) | **Post** /console/v1/gates/{id}/load_pulse | Load Pulse Gate
[**ConsoleV1GatesControllerGenMultiRuleAdd**](GatesAPI.md#ConsoleV1GatesControllerGenMultiRuleAdd) | **Post** /console/v1/gates/{id}/rules | Add Multiple Gate Rule
[**ConsoleV1GatesControllerGenMultiRuleUpdate**](GatesAPI.md#ConsoleV1GatesControllerGenMultiRuleUpdate) | **Patch** /console/v1/gates/{id}/rules | Update list of Gate Rules
[**ConsoleV1GatesControllerGenPartialUpdate**](GatesAPI.md#ConsoleV1GatesControllerGenPartialUpdate) | **Patch** /console/v1/gates/{id} | Partially Update Gates
[**ConsoleV1GatesControllerGenPulseLoadHistory**](GatesAPI.md#ConsoleV1GatesControllerGenPulseLoadHistory) | **Get** /console/v1/gates/{id}/rules/{ruleID}/pulse_load_history | Pulse Load History (Warehouse Native)
[**ConsoleV1GatesControllerGenPulseResults**](GatesAPI.md#ConsoleV1GatesControllerGenPulseResults) | **Get** /console/v1/gates/{id}/rules/{ruleID}/pulse_results | Retrieve Pulse Results (Beta)
[**ConsoleV1GatesControllerGenRead**](GatesAPI.md#ConsoleV1GatesControllerGenRead) | **Get** /console/v1/gates/{id} | Read Gate
[**ConsoleV1GatesControllerGenReadRules**](GatesAPI.md#ConsoleV1GatesControllerGenReadRules) | **Get** /console/v1/gates/{id}/rules | Read Gate Rules
[**ConsoleV1GatesControllerGenRemove**](GatesAPI.md#ConsoleV1GatesControllerGenRemove) | **Delete** /console/v1/gates/{id} | Delete Gates
[**ConsoleV1GatesControllerGenRuleAdd**](GatesAPI.md#ConsoleV1GatesControllerGenRuleAdd) | **Post** /console/v1/gates/{id}/rule | Add Gate Rule
[**ConsoleV1GatesControllerGenRuleDelete**](GatesAPI.md#ConsoleV1GatesControllerGenRuleDelete) | **Delete** /console/v1/gates/{id}/rules/{ruleID} | Delete Gate Rules
[**ConsoleV1GatesControllerGenRuleUpdate**](GatesAPI.md#ConsoleV1GatesControllerGenRuleUpdate) | **Patch** /console/v1/gates/{id}/rules/{ruleID} | Update Gate Rules



## ConsoleV1GateOverridesControllerGenAdd

> ConsoleV1GateOverridesControllerGenUpdate201Response ConsoleV1GateOverridesControllerGenAdd(ctx, id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Add Gate Overrides

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	updateOverridesContractDto := openapiclient.UpdateOverridesContractDto{UpdateOverridesContractDtoOneOf: openapiclient.NewUpdateOverridesContractDtoOneOf([]openapiclient.OverrideDtoEnvironmentOverridesInner{*openapiclient.NewOverrideDtoEnvironmentOverridesInner("TODO", []string{"PassingIDs_example"}, []string{"FailingIDs_example"})})} // UpdateOverridesContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenAdd(context.Background(), id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GateOverridesControllerGenAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GateOverridesControllerGenAdd`: ConsoleV1GateOverridesControllerGenUpdate201Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GateOverridesControllerGenAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GateOverridesControllerGenAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOverridesContractDto** | [**UpdateOverridesContractDto**](UpdateOverridesContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GateOverridesControllerGenUpdate201Response**](ConsoleV1GateOverridesControllerGenUpdate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GateOverridesControllerGenRead

> ConsoleV1GateOverridesControllerGenRead201Response ConsoleV1GateOverridesControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Gate Override

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GateOverridesControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GateOverridesControllerGenRead`: ConsoleV1GateOverridesControllerGenRead201Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GateOverridesControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GateOverridesControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GateOverridesControllerGenRead201Response**](ConsoleV1GateOverridesControllerGenRead201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GateOverridesControllerGenRemove

> ConsoleV1GateOverridesControllerGenUpdate201Response ConsoleV1GateOverridesControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Gate Overrides

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GateOverridesControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GateOverridesControllerGenRemove`: ConsoleV1GateOverridesControllerGenUpdate201Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GateOverridesControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GateOverridesControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GateOverridesControllerGenUpdate201Response**](ConsoleV1GateOverridesControllerGenUpdate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GateOverridesControllerGenUpdate

> ConsoleV1GateOverridesControllerGenUpdate201Response ConsoleV1GateOverridesControllerGenUpdate(ctx, id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Gate Overrides

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	updateOverridesContractDto := openapiclient.UpdateOverridesContractDto{UpdateOverridesContractDtoOneOf: openapiclient.NewUpdateOverridesContractDtoOneOf([]openapiclient.OverrideDtoEnvironmentOverridesInner{*openapiclient.NewOverrideDtoEnvironmentOverridesInner("TODO", []string{"PassingIDs_example"}, []string{"FailingIDs_example"})})} // UpdateOverridesContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GateOverridesControllerGenUpdate(context.Background(), id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GateOverridesControllerGenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GateOverridesControllerGenUpdate`: ConsoleV1GateOverridesControllerGenUpdate201Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GateOverridesControllerGenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GateOverridesControllerGenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOverridesContractDto** | [**UpdateOverridesContractDto**](UpdateOverridesContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GateOverridesControllerGenUpdate201Response**](ConsoleV1GateOverridesControllerGenUpdate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenArchive

> ConsoleV1GatesControllerGenArchive200Response ConsoleV1GatesControllerGenArchive(ctx, id).ArchiveSchemaDto(archiveSchemaDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Archive Gate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	archiveSchemaDto := *openapiclient.NewArchiveSchemaDto() // ArchiveSchemaDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenArchive(context.Background(), id).ArchiveSchemaDto(archiveSchemaDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenArchive`: ConsoleV1GatesControllerGenArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveSchemaDto** | [**ArchiveSchemaDto**](ArchiveSchemaDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenArchive200Response**](ConsoleV1GatesControllerGenArchive200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenCreate

> ConsoleV1GatesControllerGenCreate200Response ConsoleV1GatesControllerGenCreate(ctx).GateCreateDto(gateCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Gate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gateCreateDto := *openapiclient.NewGateCreateDto() // GateCreateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenCreate(context.Background()).GateCreateDto(gateCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenCreate`: ConsoleV1GatesControllerGenCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gateCreateDto** | [**GateCreateDto**](GateCreateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenCreate200Response**](ConsoleV1GatesControllerGenCreate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenDisable

> ConsoleV1GatesControllerGenLaunch200Response ConsoleV1GatesControllerGenDisable(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Disable Gate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenDisable(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenDisable`: ConsoleV1GatesControllerGenLaunch200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenLaunch200Response**](ConsoleV1GatesControllerGenLaunch200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenEnable

> ConsoleV1GatesControllerGenLaunch200Response ConsoleV1GatesControllerGenEnable(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Enable Gate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenEnable(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenEnable`: ConsoleV1GatesControllerGenLaunch200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenLaunch200Response**](ConsoleV1GatesControllerGenLaunch200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenFullUpdate

> ConsoleV1GatesControllerGenFullUpdate200Response ConsoleV1GatesControllerGenFullUpdate(ctx, id).GateFullUpdateDto(gateFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Fully Update Gates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	gateFullUpdateDto := *openapiclient.NewGateFullUpdateDto(false, "Description_example", []openapiclient.ExternalGateDtoRulesInner{*openapiclient.NewExternalGateDtoRulesInner("Name_example", float32(123), []openapiclient.DynamicConfigDtoRulesInnerConditionsInner{*openapiclient.NewDynamicConfigDtoRulesInnerConditionsInner("Type_example")})}) // GateFullUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenFullUpdate(context.Background(), id).GateFullUpdateDto(gateFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenFullUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenFullUpdate`: ConsoleV1GatesControllerGenFullUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenFullUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenFullUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gateFullUpdateDto** | [**GateFullUpdateDto**](GateFullUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenFullUpdate200Response**](ConsoleV1GatesControllerGenFullUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenLaunch

> ConsoleV1GatesControllerGenLaunch200Response ConsoleV1GatesControllerGenLaunch(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Launch Gate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenLaunch(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenLaunch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenLaunch`: ConsoleV1GatesControllerGenLaunch200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenLaunch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenLaunchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenLaunch200Response**](ConsoleV1GatesControllerGenLaunch200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenList

> ConsoleV1GatesControllerGenList200Response ConsoleV1GatesControllerGenList(ctx).IdType(idType).Type_(type_).TypeReason(typeReason).PassRate(passRate).RolloutRate(rolloutRate).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Gates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	idType := []string{"Inner_example"} // []string | Filter by idType (optional)
	type_ := "type__example" // string | Filter by type (optional)
	typeReason := "typeReason_example" // string | Filter by typeReason (optional)
	passRate := TODO // interface{} | Filter by pass rate of the gates, as determined by a sampling of overall true/false values returned: 0, 100, or INBETWEEN (pass rate greater than zero but less than 100) (optional)
	rolloutRate := TODO // interface{} | Filter by rollout rate of the gates: 0 (all rules are set to pass 0%), 100 (all rules pass 100% including an \"everyone\" catch all rule), or INBETWEEN (at least one rule has a pass rate greater than 0 but less than 100) (optional)
	creatorName := "creatorName_example" // string | Name of the creator. (optional)
	creatorID := "creatorID_example" // string | ID of the user who created the entity. (optional)
	tags := []string{"Inner_example"} // []string | Filter by tags (optional)
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenList(context.Background()).IdType(idType).Type_(type_).TypeReason(typeReason).PassRate(passRate).RolloutRate(rolloutRate).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenList`: ConsoleV1GatesControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idType** | **[]string** | Filter by idType | 
 **type_** | **string** | Filter by type | 
 **typeReason** | **string** | Filter by typeReason | 
 **passRate** | [**interface{}**](interface{}.md) | Filter by pass rate of the gates, as determined by a sampling of overall true/false values returned: 0, 100, or INBETWEEN (pass rate greater than zero but less than 100) | 
 **rolloutRate** | [**interface{}**](interface{}.md) | Filter by rollout rate of the gates: 0 (all rules are set to pass 0%), 100 (all rules pass 100% including an \&quot;everyone\&quot; catch all rule), or INBETWEEN (at least one rule has a pass rate greater than 0 but less than 100) | 
 **creatorName** | **string** | Name of the creator. | 
 **creatorID** | **string** | ID of the user who created the entity. | 
 **tags** | **[]string** | Filter by tags | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenList200Response**](ConsoleV1GatesControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenLoadPulse

> ConsoleV1GatesControllerGenLoadPulse200Response ConsoleV1GatesControllerGenLoadPulse(ctx, id).EchidnaGateLoadPulseQueryDto(echidnaGateLoadPulseQueryDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Load Pulse Gate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	echidnaGateLoadPulseQueryDto := *openapiclient.NewEchidnaGateLoadPulseQueryDto("RuleId_example") // EchidnaGateLoadPulseQueryDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenLoadPulse(context.Background(), id).EchidnaGateLoadPulseQueryDto(echidnaGateLoadPulseQueryDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenLoadPulse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenLoadPulse`: ConsoleV1GatesControllerGenLoadPulse200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenLoadPulse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenLoadPulseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **echidnaGateLoadPulseQueryDto** | [**EchidnaGateLoadPulseQueryDto**](EchidnaGateLoadPulseQueryDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenLoadPulse200Response**](ConsoleV1GatesControllerGenLoadPulse200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenMultiRuleAdd

> ConsoleV1GatesControllerGenMultiRuleAdd200Response ConsoleV1GatesControllerGenMultiRuleAdd(ctx, id).MultiRuleDto(multiRuleDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Add Multiple Gate Rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	multiRuleDto := *openapiclient.NewMultiRuleDto([]openapiclient.ExternalGateDtoRulesInner{*openapiclient.NewExternalGateDtoRulesInner("Name_example", float32(123), []openapiclient.DynamicConfigDtoRulesInnerConditionsInner{*openapiclient.NewDynamicConfigDtoRulesInnerConditionsInner("Type_example")})}) // MultiRuleDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenMultiRuleAdd(context.Background(), id).MultiRuleDto(multiRuleDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenMultiRuleAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenMultiRuleAdd`: ConsoleV1GatesControllerGenMultiRuleAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenMultiRuleAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenMultiRuleAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multiRuleDto** | [**MultiRuleDto**](MultiRuleDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenMultiRuleAdd200Response**](ConsoleV1GatesControllerGenMultiRuleAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenMultiRuleUpdate

> ConsoleV1GatesControllerGenMultiRuleUpdate200Response ConsoleV1GatesControllerGenMultiRuleUpdate(ctx, id).MultiRuleUpdateDto(multiRuleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update list of Gate Rules

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	multiRuleUpdateDto := *openapiclient.NewMultiRuleUpdateDto([]openapiclient.MultiRuleUpdateDtoRulesInner{*openapiclient.NewMultiRuleUpdateDtoRulesInner()}) // MultiRuleUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenMultiRuleUpdate(context.Background(), id).MultiRuleUpdateDto(multiRuleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenMultiRuleUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenMultiRuleUpdate`: ConsoleV1GatesControllerGenMultiRuleUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenMultiRuleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenMultiRuleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multiRuleUpdateDto** | [**MultiRuleUpdateDto**](MultiRuleUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenMultiRuleUpdate200Response**](ConsoleV1GatesControllerGenMultiRuleUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenPartialUpdate

> ConsoleV1GatesControllerGenPartialUpdate200Response ConsoleV1GatesControllerGenPartialUpdate(ctx, id).GatePartialUpdateDto(gatePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Partially Update Gates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	gatePartialUpdateDto := *openapiclient.NewGatePartialUpdateDto() // GatePartialUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenPartialUpdate(context.Background(), id).GatePartialUpdateDto(gatePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenPartialUpdate`: ConsoleV1GatesControllerGenPartialUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatePartialUpdateDto** | [**GatePartialUpdateDto**](GatePartialUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenPartialUpdate200Response**](ConsoleV1GatesControllerGenPartialUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenPulseLoadHistory

> ConsoleV1GatesControllerGenPulseLoadHistory200Response ConsoleV1GatesControllerGenPulseLoadHistory(ctx, id, ruleID).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

Pulse Load History (Warehouse Native)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Gate ID
	ruleID := "ruleID_example" // string | Rule ID
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenPulseLoadHistory(context.Background(), id, ruleID).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenPulseLoadHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenPulseLoadHistory`: ConsoleV1GatesControllerGenPulseLoadHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenPulseLoadHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Gate ID | 
**ruleID** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenPulseLoadHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenPulseLoadHistory200Response**](ConsoleV1GatesControllerGenPulseLoadHistory200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenPulseResults

> ConsoleV1GatesControllerGenPulseResults200Response ConsoleV1GatesControllerGenPulseResults(ctx, id, ruleID).Cuped(cuped).XRespectReviewSettings(xRespectReviewSettings).Execute()

Retrieve Pulse Results (Beta)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Gate ID
	ruleID := "ruleID_example" // string | Rule ID
	cuped := "cuped_example" // string | Whether to apply CUPED. Allowed values are \"true\" or \"false\".
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenPulseResults(context.Background(), id, ruleID).Cuped(cuped).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenPulseResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenPulseResults`: ConsoleV1GatesControllerGenPulseResults200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenPulseResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Gate ID | 
**ruleID** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenPulseResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cuped** | **string** | Whether to apply CUPED. Allowed values are \&quot;true\&quot; or \&quot;false\&quot;. | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenPulseResults200Response**](ConsoleV1GatesControllerGenPulseResults200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenRead

> ConsoleV1GatesControllerGenRead200Response ConsoleV1GatesControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Gate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenRead`: ConsoleV1GatesControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenRead200Response**](ConsoleV1GatesControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenReadRules

> ConsoleV1GatesControllerGenReadRules200Response ConsoleV1GatesControllerGenReadRules(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Gate Rules

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenReadRules(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenReadRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenReadRules`: ConsoleV1GatesControllerGenReadRules200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenReadRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenReadRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenReadRules200Response**](ConsoleV1GatesControllerGenReadRules200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenRemove

> ConsoleV1GatesControllerGenRemove200Response ConsoleV1GatesControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Gates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenRemove`: ConsoleV1GatesControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenRemove200Response**](ConsoleV1GatesControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenRuleAdd

> ConsoleV1GatesControllerGenMultiRuleAdd200Response ConsoleV1GatesControllerGenRuleAdd(ctx, id).RuleDto(ruleDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Add Gate Rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | id
	ruleDto := *openapiclient.NewRuleDto("Name_example", float32(123), []openapiclient.DynamicConfigDtoRulesInnerConditionsInner{*openapiclient.NewDynamicConfigDtoRulesInnerConditionsInner("Type_example")}) // RuleDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRuleAdd(context.Background(), id).RuleDto(ruleDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenRuleAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenRuleAdd`: ConsoleV1GatesControllerGenMultiRuleAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenRuleAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenRuleAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleDto** | [**RuleDto**](RuleDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenMultiRuleAdd200Response**](ConsoleV1GatesControllerGenMultiRuleAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenRuleDelete

> ConsoleV1GatesControllerGenPartialUpdate200Response ConsoleV1GatesControllerGenRuleDelete(ctx, id, ruleID).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Gate Rules

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Gate ID
	ruleID := "ruleID_example" // string | Rule ID
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRuleDelete(context.Background(), id, ruleID).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenRuleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenRuleDelete`: ConsoleV1GatesControllerGenPartialUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Gate ID | 
**ruleID** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenPartialUpdate200Response**](ConsoleV1GatesControllerGenPartialUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1GatesControllerGenRuleUpdate

> ConsoleV1GatesControllerGenMultiRuleUpdate200Response ConsoleV1GatesControllerGenRuleUpdate(ctx, id, ruleID).RuleUpdateDto(ruleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Gate Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Gate ID
	ruleID := "ruleID_example" // string | Rule ID
	ruleUpdateDto := *openapiclient.NewRuleUpdateDto() // RuleUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatesAPI.ConsoleV1GatesControllerGenRuleUpdate(context.Background(), id, ruleID).RuleUpdateDto(ruleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatesAPI.ConsoleV1GatesControllerGenRuleUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1GatesControllerGenRuleUpdate`: ConsoleV1GatesControllerGenMultiRuleUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatesAPI.ConsoleV1GatesControllerGenRuleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Gate ID | 
**ruleID** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1GatesControllerGenRuleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleUpdateDto** | [**RuleUpdateDto**](RuleUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1GatesControllerGenMultiRuleUpdate200Response**](ConsoleV1GatesControllerGenMultiRuleUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

