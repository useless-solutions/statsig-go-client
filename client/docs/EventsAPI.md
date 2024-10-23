# \EventsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1EventsControllerGenList**](EventsAPI.md#ConsoleV1EventsControllerGenList) | **Get** /console/v1/events | List Events
[**ConsoleV1EventsControllerGenListMetricsByEvent**](EventsAPI.md#ConsoleV1EventsControllerGenListMetricsByEvent) | **Get** /console/v1/events/{eventName}/metrics | Get metrics using event name
[**ConsoleV1EventsControllerGenListSpecificEvent**](EventsAPI.md#ConsoleV1EventsControllerGenListSpecificEvent) | **Get** /console/v1/events/{eventName} | Get specific events



## ConsoleV1EventsControllerGenList

> ConsoleV1EventsControllerGenList200Response ConsoleV1EventsControllerGenList(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Events

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
	resp, r, err := apiClient.EventsAPI.ConsoleV1EventsControllerGenList(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ConsoleV1EventsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1EventsControllerGenList`: ConsoleV1EventsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ConsoleV1EventsControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1EventsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1EventsControllerGenList200Response**](ConsoleV1EventsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1EventsControllerGenListMetricsByEvent

> ConsoleV1EventsControllerGenListMetricsByEvent200Response ConsoleV1EventsControllerGenListMetricsByEvent(ctx, eventName).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get metrics using event name

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
	eventName := "eventName_example" // string | 
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ConsoleV1EventsControllerGenListMetricsByEvent(context.Background(), eventName).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ConsoleV1EventsControllerGenListMetricsByEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1EventsControllerGenListMetricsByEvent`: ConsoleV1EventsControllerGenListMetricsByEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ConsoleV1EventsControllerGenListMetricsByEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1EventsControllerGenListMetricsByEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1EventsControllerGenListMetricsByEvent200Response**](ConsoleV1EventsControllerGenListMetricsByEvent200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1EventsControllerGenListSpecificEvent

> ConsoleV1EventsControllerGenListSpecificEvent200Response ConsoleV1EventsControllerGenListSpecificEvent(ctx, eventName).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get specific events

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
	eventName := "eventName_example" // string | 
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ConsoleV1EventsControllerGenListSpecificEvent(context.Background(), eventName).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ConsoleV1EventsControllerGenListSpecificEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1EventsControllerGenListSpecificEvent`: ConsoleV1EventsControllerGenListSpecificEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ConsoleV1EventsControllerGenListSpecificEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1EventsControllerGenListSpecificEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1EventsControllerGenListSpecificEvent200Response**](ConsoleV1EventsControllerGenListSpecificEvent200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

