# \TagsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1TagsControllerGenCreate**](TagsAPI.md#ConsoleV1TagsControllerGenCreate) | **Post** /console/v1/tags | Create Tag
[**ConsoleV1TagsControllerGenList**](TagsAPI.md#ConsoleV1TagsControllerGenList) | **Get** /console/v1/tags | List Tags
[**ConsoleV1TagsControllerGenRead**](TagsAPI.md#ConsoleV1TagsControllerGenRead) | **Get** /console/v1/tags/{id} | Read Tag
[**ConsoleV1TagsControllerGenRemove**](TagsAPI.md#ConsoleV1TagsControllerGenRemove) | **Delete** /console/v1/tags/{id} | Delete Tag
[**ConsoleV1TagsControllerGenUpdate**](TagsAPI.md#ConsoleV1TagsControllerGenUpdate) | **Patch** /console/v1/tags/{id} | Update Tag



## ConsoleV1TagsControllerGenCreate

> ConsoleV1TagsControllerGenCreate200Response ConsoleV1TagsControllerGenCreate(ctx).TagCreateDto(tagCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Tag

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
	tagCreateDto := *openapiclient.NewTagCreateDto("Name_example", "Description_example") // TagCreateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.ConsoleV1TagsControllerGenCreate(context.Background()).TagCreateDto(tagCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ConsoleV1TagsControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TagsControllerGenCreate`: ConsoleV1TagsControllerGenCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ConsoleV1TagsControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TagsControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagCreateDto** | [**TagCreateDto**](TagCreateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TagsControllerGenCreate200Response**](ConsoleV1TagsControllerGenCreate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TagsControllerGenList

> ConsoleV1TagsControllerGenList200Response ConsoleV1TagsControllerGenList(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Tags

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
	resp, r, err := apiClient.TagsAPI.ConsoleV1TagsControllerGenList(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ConsoleV1TagsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TagsControllerGenList`: ConsoleV1TagsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ConsoleV1TagsControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TagsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TagsControllerGenList200Response**](ConsoleV1TagsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TagsControllerGenRead

> ConsoleV1TagsControllerGenRead200Response ConsoleV1TagsControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Tag

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
	resp, r, err := apiClient.TagsAPI.ConsoleV1TagsControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ConsoleV1TagsControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TagsControllerGenRead`: ConsoleV1TagsControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ConsoleV1TagsControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TagsControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TagsControllerGenRead200Response**](ConsoleV1TagsControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TagsControllerGenRemove

> ConsoleV1TagsControllerGenRemove200Response ConsoleV1TagsControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Tag

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
	resp, r, err := apiClient.TagsAPI.ConsoleV1TagsControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ConsoleV1TagsControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TagsControllerGenRemove`: ConsoleV1TagsControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ConsoleV1TagsControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TagsControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TagsControllerGenRemove200Response**](ConsoleV1TagsControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1TagsControllerGenUpdate

> ConsoleV1TagsControllerGenUpdate200Response ConsoleV1TagsControllerGenUpdate(ctx, id).TagUpdateDto(tagUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Tag

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
	tagUpdateDto := *openapiclient.NewTagUpdateDto() // TagUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.ConsoleV1TagsControllerGenUpdate(context.Background(), id).TagUpdateDto(tagUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ConsoleV1TagsControllerGenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1TagsControllerGenUpdate`: ConsoleV1TagsControllerGenUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ConsoleV1TagsControllerGenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1TagsControllerGenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateDto** | [**TagUpdateDto**](TagUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1TagsControllerGenUpdate200Response**](ConsoleV1TagsControllerGenUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

