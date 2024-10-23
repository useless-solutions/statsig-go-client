# \SegmentsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1SegmentsControllerGenArchive**](SegmentsAPI.md#ConsoleV1SegmentsControllerGenArchive) | **Put** /console/v1/segments/{id}/archive | Archive Segment
[**ConsoleV1SegmentsControllerGenCreate**](SegmentsAPI.md#ConsoleV1SegmentsControllerGenCreate) | **Post** /console/v1/segments | Create Segment
[**ConsoleV1SegmentsControllerGenList**](SegmentsAPI.md#ConsoleV1SegmentsControllerGenList) | **Get** /console/v1/segments | List Segments
[**ConsoleV1SegmentsControllerGenRead**](SegmentsAPI.md#ConsoleV1SegmentsControllerGenRead) | **Get** /console/v1/segments/{id} | Get Segment
[**ConsoleV1SegmentsControllerGenRemove**](SegmentsAPI.md#ConsoleV1SegmentsControllerGenRemove) | **Delete** /console/v1/segments/{id} | Delete Segment
[**ConsoleV1SegmentsIDListControllerAdd**](SegmentsAPI.md#ConsoleV1SegmentsIDListControllerAdd) | **Patch** /console/v1/segments/{id}/id_list | Add IDs to Segment
[**ConsoleV1SegmentsIDListControllerGenRemove**](SegmentsAPI.md#ConsoleV1SegmentsIDListControllerGenRemove) | **Delete** /console/v1/segments/{id}/id_list | Remove IDs from Segment
[**ConsoleV1SegmentsIDListControllerRead**](SegmentsAPI.md#ConsoleV1SegmentsIDListControllerRead) | **Get** /console/v1/segments/{id}/id_list | Get IDs in a Segment
[**ConsoleV1SegmentsIDListResetControllerReset**](SegmentsAPI.md#ConsoleV1SegmentsIDListResetControllerReset) | **Post** /console/v1/segments/{id}/id_list/reset | Upsert ID List Segment
[**ConsoleV1SegmentsRulesControllerUpdate**](SegmentsAPI.md#ConsoleV1SegmentsRulesControllerUpdate) | **Post** /console/v1/segments/{id}/conditional | Update Segment Rules



## ConsoleV1SegmentsControllerGenArchive

> ConsoleV1SegmentsControllerGenArchive200Response ConsoleV1SegmentsControllerGenArchive(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Archive Segment

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
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsControllerGenArchive(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsControllerGenArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsControllerGenArchive`: ConsoleV1SegmentsControllerGenArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsControllerGenArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsControllerGenArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsControllerGenArchive200Response**](ConsoleV1SegmentsControllerGenArchive200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsControllerGenCreate

> ConsoleV1SegmentsControllerGenCreate201Response ConsoleV1SegmentsControllerGenCreate(ctx).SegmentCreateContractDto(segmentCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Segment

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
	segmentCreateContractDto := *openapiclient.NewSegmentCreateContractDto("Name_example", "Type_example") // SegmentCreateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsControllerGenCreate(context.Background()).SegmentCreateContractDto(segmentCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsControllerGenCreate`: ConsoleV1SegmentsControllerGenCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **segmentCreateContractDto** | [**SegmentCreateContractDto**](SegmentCreateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsControllerGenCreate201Response**](ConsoleV1SegmentsControllerGenCreate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsControllerGenList

> ConsoleV1SegmentsControllerGenList200Response ConsoleV1SegmentsControllerGenList(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Segments

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
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsControllerGenList(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsControllerGenList`: ConsoleV1SegmentsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsControllerGenList200Response**](ConsoleV1SegmentsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsControllerGenRead

> ConsoleV1SegmentsControllerGenRead200Response ConsoleV1SegmentsControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Segment

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
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsControllerGenRead`: ConsoleV1SegmentsControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsControllerGenRead200Response**](ConsoleV1SegmentsControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsControllerGenRemove

> ConsoleV1SegmentsControllerGenRemove200Response ConsoleV1SegmentsControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Segment

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
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsControllerGenRemove`: ConsoleV1SegmentsControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsControllerGenRemove200Response**](ConsoleV1SegmentsControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsIDListControllerAdd

> ConsoleV1SegmentsIDListControllerAdd200Response ConsoleV1SegmentsIDListControllerAdd(ctx, id).SegmentIDListContractDto(segmentIDListContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Add IDs to Segment

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
	segmentIDListContractDto := *openapiclient.NewSegmentIDListContractDto([]string{"Ids_example"}) // SegmentIDListContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsIDListControllerAdd(context.Background(), id).SegmentIDListContractDto(segmentIDListContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsIDListControllerAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsIDListControllerAdd`: ConsoleV1SegmentsIDListControllerAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsIDListControllerAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsIDListControllerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentIDListContractDto** | [**SegmentIDListContractDto**](SegmentIDListContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsIDListControllerAdd200Response**](ConsoleV1SegmentsIDListControllerAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsIDListControllerGenRemove

> ConsoleV1SegmentsIDListControllerGenRemove200Response ConsoleV1SegmentsIDListControllerGenRemove(ctx, id).SegmentIDListContractDto(segmentIDListContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Remove IDs from Segment

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
	segmentIDListContractDto := *openapiclient.NewSegmentIDListContractDto([]string{"Ids_example"}) // SegmentIDListContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsIDListControllerGenRemove(context.Background(), id).SegmentIDListContractDto(segmentIDListContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsIDListControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsIDListControllerGenRemove`: ConsoleV1SegmentsIDListControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsIDListControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsIDListControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentIDListContractDto** | [**SegmentIDListContractDto**](SegmentIDListContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsIDListControllerGenRemove200Response**](ConsoleV1SegmentsIDListControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsIDListControllerRead

> ConsoleV1SegmentsIDListControllerRead200Response ConsoleV1SegmentsIDListControllerRead(ctx, id).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get IDs in a Segment

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
	id := "id_example" // string | 
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsIDListControllerRead(context.Background(), id).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsIDListControllerRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsIDListControllerRead`: ConsoleV1SegmentsIDListControllerRead200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsIDListControllerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsIDListControllerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsIDListControllerRead200Response**](ConsoleV1SegmentsIDListControllerRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsIDListResetControllerReset

> ConsoleV1SegmentsIDListResetControllerReset200Response ConsoleV1SegmentsIDListResetControllerReset(ctx, id).SegmentIDListResetContractDto(segmentIDListResetContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Upsert ID List Segment

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
	segmentIDListResetContractDto := *openapiclient.NewSegmentIDListResetContractDto([]string{"Ids_example"}) // SegmentIDListResetContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsIDListResetControllerReset(context.Background(), id).SegmentIDListResetContractDto(segmentIDListResetContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsIDListResetControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsIDListResetControllerReset`: ConsoleV1SegmentsIDListResetControllerReset200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsIDListResetControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsIDListResetControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentIDListResetContractDto** | [**SegmentIDListResetContractDto**](SegmentIDListResetContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsIDListResetControllerReset200Response**](ConsoleV1SegmentsIDListResetControllerReset200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1SegmentsRulesControllerUpdate

> ConsoleV1SegmentsRulesControllerUpdate200Response ConsoleV1SegmentsRulesControllerUpdate(ctx, id).SegmentCreateContractDtoRulesInner(segmentCreateContractDtoRulesInner).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Segment Rules

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
	segmentCreateContractDtoRulesInner := []openapiclient.SegmentCreateContractDtoRulesInner{*openapiclient.NewSegmentCreateContractDtoRulesInner("Name_example", float32(123), []openapiclient.SegmentCreateContractDtoRulesInnerConditionsInner{*openapiclient.NewSegmentCreateContractDtoRulesInnerConditionsInner("Type_example")})} // []SegmentCreateContractDtoRulesInner | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.ConsoleV1SegmentsRulesControllerUpdate(context.Background(), id).SegmentCreateContractDtoRulesInner(segmentCreateContractDtoRulesInner).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ConsoleV1SegmentsRulesControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1SegmentsRulesControllerUpdate`: ConsoleV1SegmentsRulesControllerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ConsoleV1SegmentsRulesControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1SegmentsRulesControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentCreateContractDtoRulesInner** | [**[]SegmentCreateContractDtoRulesInner**](SegmentCreateContractDtoRulesInner.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1SegmentsRulesControllerUpdate200Response**](ConsoleV1SegmentsRulesControllerUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

