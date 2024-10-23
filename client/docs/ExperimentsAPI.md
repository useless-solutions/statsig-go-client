# \ExperimentsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1ExperimentOverridesControllerAdd**](ExperimentsAPI.md#ConsoleV1ExperimentOverridesControllerAdd) | **Patch** /console/v1/experiments/{id}/overrides | Partially Update Experiment Overrides
[**ConsoleV1ExperimentOverridesControllerGenRead**](ExperimentsAPI.md#ConsoleV1ExperimentOverridesControllerGenRead) | **Get** /console/v1/experiments/{id}/overrides | Get Experiment Overrides
[**ConsoleV1ExperimentOverridesControllerGenRemove**](ExperimentsAPI.md#ConsoleV1ExperimentOverridesControllerGenRemove) | **Delete** /console/v1/experiments/{id}/overrides | Delete Experiment Overrides
[**ConsoleV1ExperimentOverridesControllerUpdate**](ExperimentsAPI.md#ConsoleV1ExperimentOverridesControllerUpdate) | **Post** /console/v1/experiments/{id}/overrides | Update Experiment Overrides
[**ConsoleV1ExperimentsControllerGenAbandon**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenAbandon) | **Put** /console/v1/experiments/{id}/abandon | Abandon Experiment
[**ConsoleV1ExperimentsControllerGenArchive**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenArchive) | **Put** /console/v1/experiments/{id}/archive | Archive Experiment
[**ConsoleV1ExperimentsControllerGenCreate**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenCreate) | **Post** /console/v1/experiments | Create Experiment
[**ConsoleV1ExperimentsControllerGenFullUpdate**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenFullUpdate) | **Post** /console/v1/experiments/{id} | Fully Update Experiment
[**ConsoleV1ExperimentsControllerGenList**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenList) | **Get** /console/v1/experiments | List Experiments
[**ConsoleV1ExperimentsControllerGenMakeDecision**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenMakeDecision) | **Put** /console/v1/experiments/{id}/make_decision | Finish Experiment Early
[**ConsoleV1ExperimentsControllerGenPartialUpdate**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenPartialUpdate) | **Patch** /console/v1/experiments/{id} | Partially Update Experiment
[**ConsoleV1ExperimentsControllerGenPulseResults**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenPulseResults) | **Get** /console/v1/experiments/{id}/pulse_results | Retrieve Pulse Results (Beta)
[**ConsoleV1ExperimentsControllerGenRead**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenRead) | **Get** /console/v1/experiments/{id} | Get Experiment
[**ConsoleV1ExperimentsControllerGenRemove**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenRemove) | **Delete** /console/v1/experiments/{id} | Deleted Experiment
[**ConsoleV1ExperimentsControllerGenReset**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenReset) | **Put** /console/v1/experiments/{id}/reset | Reset Experiment
[**ConsoleV1ExperimentsControllerGenStart**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenStart) | **Put** /console/v1/experiments/{id}/start | Start Experiment
[**ConsoleV1ExperimentsControllerGenUnarchive**](ExperimentsAPI.md#ConsoleV1ExperimentsControllerGenUnarchive) | **Put** /console/v1/experiments/{id}/unarchive | Unarchive Experiment
[**ConsoleV1ExposuresControllerExposureCount**](ExperimentsAPI.md#ConsoleV1ExposuresControllerExposureCount) | **Get** /console/v1/experiments/{experimentName}/exposures | Get Exposure Count for Experiment



## ConsoleV1ExperimentOverridesControllerAdd

> ConsoleV1ExperimentOverridesControllerAdd200Response ConsoleV1ExperimentOverridesControllerAdd(ctx, id).ExperimentOverridesDto(experimentOverridesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Partially Update Experiment Overrides

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
	experimentOverridesDto := *openapiclient.NewExperimentOverridesDto([]openapiclient.ExperimentOverridesDtoOverridesInner{*openapiclient.NewExperimentOverridesDtoOverridesInner("Type_example", "Id_example", "GroupID_example")}, []openapiclient.ExperimentOverridesDtoUserIDOverridesInner{*openapiclient.NewExperimentOverridesDtoUserIDOverridesInner("GroupID_example", []string{"Ids_example"})}) // ExperimentOverridesDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentOverridesControllerAdd(context.Background(), id).ExperimentOverridesDto(experimentOverridesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentOverridesControllerAdd`: ConsoleV1ExperimentOverridesControllerAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentOverridesControllerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentOverridesDto** | [**ExperimentOverridesDto**](ExperimentOverridesDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentOverridesControllerAdd200Response**](ConsoleV1ExperimentOverridesControllerAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentOverridesControllerGenRead

> ConsoleV1ExperimentOverridesControllerGenRead200Response ConsoleV1ExperimentOverridesControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Experiment Overrides

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
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentOverridesControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentOverridesControllerGenRead`: ConsoleV1ExperimentOverridesControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentOverridesControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentOverridesControllerGenRead200Response**](ConsoleV1ExperimentOverridesControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentOverridesControllerGenRemove

> ConsoleV1ExperimentOverridesControllerGenRemove200Response ConsoleV1ExperimentOverridesControllerGenRemove(ctx, id).ExperimentOverridesDto(experimentOverridesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Experiment Overrides

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
	experimentOverridesDto := *openapiclient.NewExperimentOverridesDto([]openapiclient.ExperimentOverridesDtoOverridesInner{*openapiclient.NewExperimentOverridesDtoOverridesInner("Type_example", "Id_example", "GroupID_example")}, []openapiclient.ExperimentOverridesDtoUserIDOverridesInner{*openapiclient.NewExperimentOverridesDtoUserIDOverridesInner("GroupID_example", []string{"Ids_example"})}) // ExperimentOverridesDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentOverridesControllerGenRemove(context.Background(), id).ExperimentOverridesDto(experimentOverridesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentOverridesControllerGenRemove`: ConsoleV1ExperimentOverridesControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentOverridesControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentOverridesDto** | [**ExperimentOverridesDto**](ExperimentOverridesDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentOverridesControllerGenRemove200Response**](ConsoleV1ExperimentOverridesControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentOverridesControllerUpdate

> ConsoleV1ExperimentOverridesControllerUpdate200Response ConsoleV1ExperimentOverridesControllerUpdate(ctx, id).ExperimentOverridesDto(experimentOverridesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Experiment Overrides

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
	experimentOverridesDto := *openapiclient.NewExperimentOverridesDto([]openapiclient.ExperimentOverridesDtoOverridesInner{*openapiclient.NewExperimentOverridesDtoOverridesInner("Type_example", "Id_example", "GroupID_example")}, []openapiclient.ExperimentOverridesDtoUserIDOverridesInner{*openapiclient.NewExperimentOverridesDtoUserIDOverridesInner("GroupID_example", []string{"Ids_example"})}) // ExperimentOverridesDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentOverridesControllerUpdate(context.Background(), id).ExperimentOverridesDto(experimentOverridesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentOverridesControllerUpdate`: ConsoleV1ExperimentOverridesControllerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentOverridesControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentOverridesControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentOverridesDto** | [**ExperimentOverridesDto**](ExperimentOverridesDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentOverridesControllerUpdate200Response**](ConsoleV1ExperimentOverridesControllerUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenAbandon

> ConsoleV1ExperimentsControllerGenAbandon200Response ConsoleV1ExperimentsControllerGenAbandon(ctx, id).ExperimentAbandonDto(experimentAbandonDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Abandon Experiment

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
	experimentAbandonDto := *openapiclient.NewExperimentAbandonDto("Your reason for stopping early") // ExperimentAbandonDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenAbandon(context.Background(), id).ExperimentAbandonDto(experimentAbandonDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenAbandon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenAbandon`: ConsoleV1ExperimentsControllerGenAbandon200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenAbandon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenAbandonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentAbandonDto** | [**ExperimentAbandonDto**](ExperimentAbandonDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenAbandon200Response**](ConsoleV1ExperimentsControllerGenAbandon200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenArchive

> ConsoleV1ExperimentsControllerGenArchive200Response ConsoleV1ExperimentsControllerGenArchive(ctx, id).ExperimentArchiveDto(experimentArchiveDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Archive Experiment

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
	experimentArchiveDto := *openapiclient.NewExperimentArchiveDto() // ExperimentArchiveDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenArchive(context.Background(), id).ExperimentArchiveDto(experimentArchiveDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenArchive`: ConsoleV1ExperimentsControllerGenArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentArchiveDto** | [**ExperimentArchiveDto**](ExperimentArchiveDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenArchive200Response**](ConsoleV1ExperimentsControllerGenArchive200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenCreate

> ConsoleV1ExperimentsControllerGenCreate201Response ConsoleV1ExperimentsControllerGenCreate(ctx).ExperimentCreateDto(experimentCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Experiment

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
	experimentCreateDto := *openapiclient.NewExperimentCreateDto("Name_example") // ExperimentCreateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenCreate(context.Background()).ExperimentCreateDto(experimentCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenCreate`: ConsoleV1ExperimentsControllerGenCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **experimentCreateDto** | [**ExperimentCreateDto**](ExperimentCreateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenCreate201Response**](ConsoleV1ExperimentsControllerGenCreate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenFullUpdate

> ConsoleV1ExperimentsControllerGenFullUpdate200Response ConsoleV1ExperimentsControllerGenFullUpdate(ctx, id).ExperimentFullUpdateDto(experimentFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Fully Update Experiment

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
	experimentFullUpdateDto := *openapiclient.NewExperimentFullUpdateDto("Description_example", "IdType_example", "Hypothesis_example", []openapiclient.ExternalExperimentDtoGroupsInner{*openapiclient.NewExternalExperimentDtoGroupsInner("Name_example", float32(123), map[string]interface{}{"key": interface{}(123)})}, float32(123), "TODO", false, "DefaultConfidenceInterval_example", "Status_example") // ExperimentFullUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenFullUpdate(context.Background(), id).ExperimentFullUpdateDto(experimentFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenFullUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenFullUpdate`: ConsoleV1ExperimentsControllerGenFullUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenFullUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenFullUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentFullUpdateDto** | [**ExperimentFullUpdateDto**](ExperimentFullUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenFullUpdate200Response**](ConsoleV1ExperimentsControllerGenFullUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenList

> ConsoleV1ExperimentsControllerGenList200Response ConsoleV1ExperimentsControllerGenList(ctx).LayerID(layerID).IdType(idType).Status(status).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Experiments

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
	layerID := "layerID_example" // string | Which layer to place the experiment into. (optional)
	idType := "idType_example" // string | The idType the experiment will be performed on (optional)
	status := TODO // interface{} | The current status of the experiment (optional)
	creatorName := "creatorName_example" // string | Name of the creator. (optional)
	creatorID := "creatorID_example" // string | ID of the user who created the entity. (optional)
	tags := []string{"Inner_example"} // []string | Filter by tags (optional)
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenList(context.Background()).LayerID(layerID).IdType(idType).Status(status).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenList`: ConsoleV1ExperimentsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerID** | **string** | Which layer to place the experiment into. | 
 **idType** | **string** | The idType the experiment will be performed on | 
 **status** | [**interface{}**](interface{}.md) | The current status of the experiment | 
 **creatorName** | **string** | Name of the creator. | 
 **creatorID** | **string** | ID of the user who created the entity. | 
 **tags** | **[]string** | Filter by tags | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenList200Response**](ConsoleV1ExperimentsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenMakeDecision

> ConsoleV1ExperimentsControllerGenMakeDecision200Response ConsoleV1ExperimentsControllerGenMakeDecision(ctx, id).ExperimentStatusUpdateDto(experimentStatusUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Finish Experiment Early

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
	experimentStatusUpdateDto := *openapiclient.NewExperimentStatusUpdateDto("groupid123", "Your reason for stopping early") // ExperimentStatusUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenMakeDecision(context.Background(), id).ExperimentStatusUpdateDto(experimentStatusUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenMakeDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenMakeDecision`: ConsoleV1ExperimentsControllerGenMakeDecision200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenMakeDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenMakeDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentStatusUpdateDto** | [**ExperimentStatusUpdateDto**](ExperimentStatusUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenMakeDecision200Response**](ConsoleV1ExperimentsControllerGenMakeDecision200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenPartialUpdate

> ConsoleV1ExperimentsControllerGenFullUpdate200Response ConsoleV1ExperimentsControllerGenPartialUpdate(ctx, id).ExperimentPartialUpdateDto(experimentPartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Partially Update Experiment

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
	experimentPartialUpdateDto := *openapiclient.NewExperimentPartialUpdateDto() // ExperimentPartialUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenPartialUpdate(context.Background(), id).ExperimentPartialUpdateDto(experimentPartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenPartialUpdate`: ConsoleV1ExperimentsControllerGenFullUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentPartialUpdateDto** | [**ExperimentPartialUpdateDto**](ExperimentPartialUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenFullUpdate200Response**](ConsoleV1ExperimentsControllerGenFullUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenPulseResults

> ConsoleV1ExperimentsControllerGenPulseResults200Response ConsoleV1ExperimentsControllerGenPulseResults(ctx, id).Control(control).Test(test).Cuped(cuped).Confidence(confidence).ApplyBonferroniPerVariant(applyBonferroniPerVariant).ApplyBonferroniPerMetric(applyBonferroniPerMetric).BonferroniPrimaryMetricWeight(bonferroniPrimaryMetricWeight).ApplyBenjaminiHochbergPerMetric(applyBenjaminiHochbergPerMetric).ApplyBenjaminiHochbergPerVariant(applyBenjaminiHochbergPerVariant).XRespectReviewSettings(xRespectReviewSettings).Execute()

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
	id := "id_example" // string | id
	control := "control_example" // string | Control Group ID
	test := "test_example" // string | Test Group ID
	cuped := "cuped_example" // string | Whether to apply CUPED. Allowed values are \"true\" or \"false\".
	confidence := "confidence_example" // string | Confidence interval (0-100)
	applyBonferroniPerVariant := "applyBonferroniPerVariant_example" // string | Whether to apply Bonferroni Per Variant. Allowed values are \"true\" or \"false\".
	applyBonferroniPerMetric := "applyBonferroniPerMetric_example" // string | Whether to apply Bonferroni Per Metric. Allowed values are \"true\" or \"false\".
	bonferroniPrimaryMetricWeight := "bonferroniPrimaryMetricWeight_example" // string | α allocated to primary metrics
	applyBenjaminiHochbergPerMetric := "applyBenjaminiHochbergPerMetric_example" // string | Whether to apply Benjamini-Hochberg Correction Per Metric. Allowed values are \"true\" or \"false\".
	applyBenjaminiHochbergPerVariant := "applyBenjaminiHochbergPerVariant_example" // string | Whether to apply Benjamini-Hochberg Correction Per Variant. Allowed values are \"true\" or \"false\".
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenPulseResults(context.Background(), id).Control(control).Test(test).Cuped(cuped).Confidence(confidence).ApplyBonferroniPerVariant(applyBonferroniPerVariant).ApplyBonferroniPerMetric(applyBonferroniPerMetric).BonferroniPrimaryMetricWeight(bonferroniPrimaryMetricWeight).ApplyBenjaminiHochbergPerMetric(applyBenjaminiHochbergPerMetric).ApplyBenjaminiHochbergPerVariant(applyBenjaminiHochbergPerVariant).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenPulseResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenPulseResults`: ConsoleV1ExperimentsControllerGenPulseResults200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenPulseResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenPulseResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **control** | **string** | Control Group ID | 
 **test** | **string** | Test Group ID | 
 **cuped** | **string** | Whether to apply CUPED. Allowed values are \&quot;true\&quot; or \&quot;false\&quot;. | 
 **confidence** | **string** | Confidence interval (0-100) | 
 **applyBonferroniPerVariant** | **string** | Whether to apply Bonferroni Per Variant. Allowed values are \&quot;true\&quot; or \&quot;false\&quot;. | 
 **applyBonferroniPerMetric** | **string** | Whether to apply Bonferroni Per Metric. Allowed values are \&quot;true\&quot; or \&quot;false\&quot;. | 
 **bonferroniPrimaryMetricWeight** | **string** | α allocated to primary metrics | 
 **applyBenjaminiHochbergPerMetric** | **string** | Whether to apply Benjamini-Hochberg Correction Per Metric. Allowed values are \&quot;true\&quot; or \&quot;false\&quot;. | 
 **applyBenjaminiHochbergPerVariant** | **string** | Whether to apply Benjamini-Hochberg Correction Per Variant. Allowed values are \&quot;true\&quot; or \&quot;false\&quot;. | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenPulseResults200Response**](ConsoleV1ExperimentsControllerGenPulseResults200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenRead

> ConsoleV1ExperimentsControllerGenRead200Response ConsoleV1ExperimentsControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Experiment

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
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenRead`: ConsoleV1ExperimentsControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenRead200Response**](ConsoleV1ExperimentsControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenRemove

> ConsoleV1ExperimentsControllerGenRemove200Response ConsoleV1ExperimentsControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Deleted Experiment

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
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenRemove`: ConsoleV1ExperimentsControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenRemove200Response**](ConsoleV1ExperimentsControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenReset

> ConsoleV1ExperimentsControllerGenReset200Response ConsoleV1ExperimentsControllerGenReset(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Reset Experiment

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
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenReset(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenReset`: ConsoleV1ExperimentsControllerGenReset200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenReset200Response**](ConsoleV1ExperimentsControllerGenReset200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenStart

> ConsoleV1ExperimentsControllerGenStart200Response ConsoleV1ExperimentsControllerGenStart(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Start Experiment

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
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenStart(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenStart`: ConsoleV1ExperimentsControllerGenStart200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenStart200Response**](ConsoleV1ExperimentsControllerGenStart200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenUnarchive

> ConsoleV1ExperimentsControllerGenUnarchive200Response ConsoleV1ExperimentsControllerGenUnarchive(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Unarchive Experiment

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
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExperimentsControllerGenUnarchive(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExperimentsControllerGenUnarchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenUnarchive`: ConsoleV1ExperimentsControllerGenUnarchive200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExperimentsControllerGenUnarchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenUnarchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenUnarchive200Response**](ConsoleV1ExperimentsControllerGenUnarchive200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExposuresControllerExposureCount

> ExposureCountDto ConsoleV1ExposuresControllerExposureCount(ctx, experimentName).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Exposure Count for Experiment

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
	experimentName := "experimentName_example" // string | Exposure name
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ConsoleV1ExposuresControllerExposureCount(context.Background(), experimentName).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ConsoleV1ExposuresControllerExposureCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExposuresControllerExposureCount`: ExposureCountDto
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ConsoleV1ExposuresControllerExposureCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experimentName** | **string** | Exposure name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExposuresControllerExposureCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ExposureCountDto**](ExposureCountDto.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

