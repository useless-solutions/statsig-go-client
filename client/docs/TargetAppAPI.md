# \TargetAppAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1TargetAppControllerBulkUpdate**](TargetAppAPI.md#ConsoleV1TargetAppControllerBulkUpdate) | **Patch** /console/v1/target_app | Bulk Assign Target Apps
[**ConsoleV1TargetAppControllerGenCreate**](TargetAppAPI.md#ConsoleV1TargetAppControllerGenCreate) | **Post** /console/v1/target_app | Create Target App
[**ConsoleV1TargetAppControllerGenDelete**](TargetAppAPI.md#ConsoleV1TargetAppControllerGenDelete) | **Delete** /console/v1/target_app/{id} | Delete Target App
[**ConsoleV1TargetAppControllerGenList**](TargetAppAPI.md#ConsoleV1TargetAppControllerGenList) | **Get** /console/v1/target_app | List Target Apps
[**ConsoleV1TargetAppControllerGenRead**](TargetAppAPI.md#ConsoleV1TargetAppControllerGenRead) | **Get** /console/v1/target_app/{id} | Read Target App
[**ConsoleV1TargetAppControllerGenUpdate**](TargetAppAPI.md#ConsoleV1TargetAppControllerGenUpdate) | **Patch** /console/v1/target_app/{id} | Update Target App



## ConsoleV1TargetAppControllerBulkUpdate

> ConsoleV1TargetAppControllerGenUpdate200Response ConsoleV1TargetAppControllerBulkUpdate(ctx).BulkAssignConfigTargetAppDto(bulkAssignConfigTargetAppDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Bulk Assign Target Apps

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
	bulkAssignConfigTargetAppDto := *openapiclient.NewBulkAssignConfigTargetAppDto([]string{"TargetApps_example"}) // BulkAssignConfigTargetAppDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetAppAPI.ConsoleV1TargetAppControllerBulkUpdate(context.Background()).BulkAssignConfigTargetAppDto(bulkAssignConfigTargetAppDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetAppAPI.ConsoleV1TargetAppControllerBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TargetAppControllerBulkUpdate`: ConsoleV1TargetAppControllerGenUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `TargetAppAPI.ConsoleV1TargetAppControllerBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TargetAppControllerBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkAssignConfigTargetAppDto** | [**BulkAssignConfigTargetAppDto**](BulkAssignConfigTargetAppDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TargetAppControllerGenUpdate200Response**](ConsoleV1TargetAppControllerGenUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TargetAppControllerGenCreate

> ConsoleV1TargetAppControllerGenRead200Response ConsoleV1TargetAppControllerGenCreate(ctx).TargetAppCreateDto(targetAppCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Target App

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
	targetAppCreateDto := *openapiclient.NewTargetAppCreateDto("string", "a description") // TargetAppCreateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetAppAPI.ConsoleV1TargetAppControllerGenCreate(context.Background()).TargetAppCreateDto(targetAppCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetAppAPI.ConsoleV1TargetAppControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TargetAppControllerGenCreate`: ConsoleV1TargetAppControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `TargetAppAPI.ConsoleV1TargetAppControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TargetAppControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetAppCreateDto** | [**TargetAppCreateDto**](TargetAppCreateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TargetAppControllerGenRead200Response**](ConsoleV1TargetAppControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TargetAppControllerGenDelete

> ConsoleV1TargetAppControllerGenDelete200Response ConsoleV1TargetAppControllerGenDelete(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Target App

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
	id := "string" // string | id of target app
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetAppAPI.ConsoleV1TargetAppControllerGenDelete(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetAppAPI.ConsoleV1TargetAppControllerGenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TargetAppControllerGenDelete`: ConsoleV1TargetAppControllerGenDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `TargetAppAPI.ConsoleV1TargetAppControllerGenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target app | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TargetAppControllerGenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TargetAppControllerGenDelete200Response**](ConsoleV1TargetAppControllerGenDelete200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TargetAppControllerGenList

> ConsoleV1TargetAppControllerGenList200Response ConsoleV1TargetAppControllerGenList(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Target Apps

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
	resp, r, err := apiClient.TargetAppAPI.ConsoleV1TargetAppControllerGenList(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetAppAPI.ConsoleV1TargetAppControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TargetAppControllerGenList`: ConsoleV1TargetAppControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `TargetAppAPI.ConsoleV1TargetAppControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TargetAppControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TargetAppControllerGenList200Response**](ConsoleV1TargetAppControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TargetAppControllerGenRead

> ConsoleV1TargetAppControllerGenRead200Response ConsoleV1TargetAppControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Target App

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
	resp, r, err := apiClient.TargetAppAPI.ConsoleV1TargetAppControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetAppAPI.ConsoleV1TargetAppControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TargetAppControllerGenRead`: ConsoleV1TargetAppControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `TargetAppAPI.ConsoleV1TargetAppControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TargetAppControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TargetAppControllerGenRead200Response**](ConsoleV1TargetAppControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TargetAppControllerGenUpdate

> ConsoleV1TargetAppControllerGenUpdate200Response ConsoleV1TargetAppControllerGenUpdate(ctx, id).UpdateTargetAppDto(updateTargetAppDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Target App

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
	updateTargetAppDto := *openapiclient.NewUpdateTargetAppDto() // UpdateTargetAppDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetAppAPI.ConsoleV1TargetAppControllerGenUpdate(context.Background(), id).UpdateTargetAppDto(updateTargetAppDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetAppAPI.ConsoleV1TargetAppControllerGenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TargetAppControllerGenUpdate`: ConsoleV1TargetAppControllerGenUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `TargetAppAPI.ConsoleV1TargetAppControllerGenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TargetAppControllerGenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTargetAppDto** | [**UpdateTargetAppDto**](UpdateTargetAppDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TargetAppControllerGenUpdate200Response**](ConsoleV1TargetAppControllerGenUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

