# \AutotunesAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1AutotunesControllerGenCreate**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenCreate) | **Post** /console/v1/autotunes | Create Autotune
[**ConsoleV1AutotunesControllerGenFullUpdate**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenFullUpdate) | **Post** /console/v1/autotunes/{id} | Fully Update Autotune
[**ConsoleV1AutotunesControllerGenList**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenList) | **Get** /console/v1/autotunes | List Autotune
[**ConsoleV1AutotunesControllerGenMakeDecision**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenMakeDecision) | **Put** /console/v1/autotunes/{id}/make_decision | Finish Experiment Early
[**ConsoleV1AutotunesControllerGenPartialUpdate**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenPartialUpdate) | **Patch** /console/v1/autotunes/{id} | Partially Update Autotune
[**ConsoleV1AutotunesControllerGenRead**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenRead) | **Get** /console/v1/autotunes/{id} | Read Autotune
[**ConsoleV1AutotunesControllerGenRemove**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenRemove) | **Delete** /console/v1/autotunes/{id} | Delete Autotune
[**ConsoleV1AutotunesControllerGenReset**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenReset) | **Put** /console/v1/autotunes/{id}/reset | Reset Experiment
[**ConsoleV1AutotunesControllerGenStart**](AutotunesAPI.md#ConsoleV1AutotunesControllerGenStart) | **Put** /console/v1/autotunes/{id}/start | Start Autotune Experiment



## ConsoleV1AutotunesControllerGenCreate

> ConsoleV1AutotunesControllerGenCreate201Response ConsoleV1AutotunesControllerGenCreate(ctx).AutotuneCreateDto(autotuneCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Autotune

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
	autotuneCreateDto := *openapiclient.NewAutotuneCreateDto([]openapiclient.AutotuneCreateDtoVariantsInner{*openapiclient.NewAutotuneCreateDtoVariantsInner("Name_example", interface{}(123))}, "SuccessEvent_example", "ExplorationWindow_example", "AttributionWindow_example", "WinnerThreshold_example", "Name_example") // AutotuneCreateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenCreate(context.Background()).AutotuneCreateDto(autotuneCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenCreate`: ConsoleV1AutotunesControllerGenCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autotuneCreateDto** | [**AutotuneCreateDto**](AutotuneCreateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenCreate201Response**](ConsoleV1AutotunesControllerGenCreate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenFullUpdate

> ConsoleV1AutotunesControllerGenFullUpdate200Response ConsoleV1AutotunesControllerGenFullUpdate(ctx, id).AutotuneFullUpdateDto(autotuneFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Fully Update Autotune



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
	autotuneFullUpdateDto := *openapiclient.NewAutotuneFullUpdateDto([]openapiclient.AutotuneCreateDtoVariantsInner{*openapiclient.NewAutotuneCreateDtoVariantsInner("Name_example", interface{}(123))}, "SuccessEvent_example", "ExplorationWindow_example", "AttributionWindow_example", "WinnerThreshold_example") // AutotuneFullUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenFullUpdate(context.Background(), id).AutotuneFullUpdateDto(autotuneFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenFullUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenFullUpdate`: ConsoleV1AutotunesControllerGenFullUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenFullUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenFullUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autotuneFullUpdateDto** | [**AutotuneFullUpdateDto**](AutotuneFullUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenFullUpdate200Response**](ConsoleV1AutotunesControllerGenFullUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenList

> ConsoleV1AutotunesControllerGenList200Response ConsoleV1AutotunesControllerGenList(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Autotune

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
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenList(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenList`: ConsoleV1AutotunesControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenList200Response**](ConsoleV1AutotunesControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenMakeDecision

> ConsoleV1AutotunesControllerGenMakeDecision200Response ConsoleV1AutotunesControllerGenMakeDecision(ctx, id).ExperimentStatusUpdateDto(experimentStatusUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

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
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenMakeDecision(context.Background(), id).ExperimentStatusUpdateDto(experimentStatusUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenMakeDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenMakeDecision`: ConsoleV1AutotunesControllerGenMakeDecision200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenMakeDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenMakeDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentStatusUpdateDto** | [**ExperimentStatusUpdateDto**](ExperimentStatusUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenMakeDecision200Response**](ConsoleV1AutotunesControllerGenMakeDecision200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenPartialUpdate

> ConsoleV1AutotunesControllerGenPartialUpdate200Response ConsoleV1AutotunesControllerGenPartialUpdate(ctx, id).AutotunePartialUpdateDto(autotunePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Partially Update Autotune



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
	autotunePartialUpdateDto := *openapiclient.NewAutotunePartialUpdateDto() // AutotunePartialUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenPartialUpdate(context.Background(), id).AutotunePartialUpdateDto(autotunePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenPartialUpdate`: ConsoleV1AutotunesControllerGenPartialUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autotunePartialUpdateDto** | [**AutotunePartialUpdateDto**](AutotunePartialUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenPartialUpdate200Response**](ConsoleV1AutotunesControllerGenPartialUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenRead

> ConsoleV1AutotunesControllerGenRead200Response ConsoleV1AutotunesControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Autotune

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
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenRead`: ConsoleV1AutotunesControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenRead200Response**](ConsoleV1AutotunesControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenRemove

> ConsoleV1AutotunesControllerGenRemove200Response ConsoleV1AutotunesControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Autotune

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
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenRemove`: ConsoleV1AutotunesControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenRemove200Response**](ConsoleV1AutotunesControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenReset

> ConsoleV1AutotunesControllerGenReset200Response ConsoleV1AutotunesControllerGenReset(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

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
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenReset(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenReset`: ConsoleV1AutotunesControllerGenReset200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenReset200Response**](ConsoleV1AutotunesControllerGenReset200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1AutotunesControllerGenStart

> ConsoleV1AutotunesControllerGenStart200Response ConsoleV1AutotunesControllerGenStart(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Start Autotune Experiment

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
	resp, r, err := apiClient.AutotunesAPI.ConsoleV1AutotunesControllerGenStart(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutotunesAPI.ConsoleV1AutotunesControllerGenStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AutotunesControllerGenStart`: ConsoleV1AutotunesControllerGenStart200Response
	fmt.Fprintf(os.Stdout, "Response from `AutotunesAPI.ConsoleV1AutotunesControllerGenStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AutotunesControllerGenStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AutotunesControllerGenStart200Response**](ConsoleV1AutotunesControllerGenStart200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

