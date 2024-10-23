# \HoldoutsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1HoldoutOverridesControllerGenAdd**](HoldoutsAPI.md#ConsoleV1HoldoutOverridesControllerGenAdd) | **Patch** /console/v1/holdouts/{id}/overrides | Add Holdout Overrides
[**ConsoleV1HoldoutOverridesControllerGenRead**](HoldoutsAPI.md#ConsoleV1HoldoutOverridesControllerGenRead) | **Get** /console/v1/holdouts/{id}/overrides | Read Holdout Overrides
[**ConsoleV1HoldoutOverridesControllerGenRemove**](HoldoutsAPI.md#ConsoleV1HoldoutOverridesControllerGenRemove) | **Delete** /console/v1/holdouts/{id}/overrides | Remove Holdout Overrides
[**ConsoleV1HoldoutOverridesControllerGenUpdate**](HoldoutsAPI.md#ConsoleV1HoldoutOverridesControllerGenUpdate) | **Post** /console/v1/holdouts/{id}/overrides | Update Holdout Overrides
[**ConsoleV1HoldoutsControllerGenCreate**](HoldoutsAPI.md#ConsoleV1HoldoutsControllerGenCreate) | **Post** /console/v1/holdouts | Create holdout
[**ConsoleV1HoldoutsControllerGenFullUpdate**](HoldoutsAPI.md#ConsoleV1HoldoutsControllerGenFullUpdate) | **Post** /console/v1/holdouts/{id} | Update holdout by id
[**ConsoleV1HoldoutsControllerGenList**](HoldoutsAPI.md#ConsoleV1HoldoutsControllerGenList) | **Get** /console/v1/holdouts | List Holdouts
[**ConsoleV1HoldoutsControllerGenPartialUpdate**](HoldoutsAPI.md#ConsoleV1HoldoutsControllerGenPartialUpdate) | **Patch** /console/v1/holdouts/{id} | Patch holdout by id
[**ConsoleV1HoldoutsControllerGenRead**](HoldoutsAPI.md#ConsoleV1HoldoutsControllerGenRead) | **Get** /console/v1/holdouts/{id} | Get holdout by id
[**ConsoleV1HoldoutsControllerGenRemove**](HoldoutsAPI.md#ConsoleV1HoldoutsControllerGenRemove) | **Delete** /console/v1/holdouts/{id} | Delete holdout by id



## ConsoleV1HoldoutOverridesControllerGenAdd

> ConsoleV1HoldoutOverridesControllerGenUpdate200Response ConsoleV1HoldoutOverridesControllerGenAdd(ctx, id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Add Holdout Overrides

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
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenAdd(context.Background(), id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutOverridesControllerGenAdd`: ConsoleV1HoldoutOverridesControllerGenUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutOverridesControllerGenAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOverridesContractDto** | [**UpdateOverridesContractDto**](UpdateOverridesContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutOverridesControllerGenUpdate200Response**](ConsoleV1HoldoutOverridesControllerGenUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutOverridesControllerGenRead

> ConsoleV1HoldoutOverridesControllerGenRead200Response ConsoleV1HoldoutOverridesControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Holdout Overrides

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
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutOverridesControllerGenRead`: ConsoleV1HoldoutOverridesControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutOverridesControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutOverridesControllerGenRead200Response**](ConsoleV1HoldoutOverridesControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutOverridesControllerGenRemove

> ConsoleV1HoldoutOverridesControllerGenRemove200Response ConsoleV1HoldoutOverridesControllerGenRemove(ctx, id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Remove Holdout Overrides



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
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenRemove(context.Background(), id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutOverridesControllerGenRemove`: ConsoleV1HoldoutOverridesControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutOverridesControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOverridesContractDto** | [**UpdateOverridesContractDto**](UpdateOverridesContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutOverridesControllerGenRemove200Response**](ConsoleV1HoldoutOverridesControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutOverridesControllerGenUpdate

> ConsoleV1HoldoutOverridesControllerGenUpdate200Response ConsoleV1HoldoutOverridesControllerGenUpdate(ctx, id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Holdout Overrides

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
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenUpdate(context.Background(), id).UpdateOverridesContractDto(updateOverridesContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutOverridesControllerGenUpdate`: ConsoleV1HoldoutOverridesControllerGenUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutOverridesControllerGenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutOverridesControllerGenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOverridesContractDto** | [**UpdateOverridesContractDto**](UpdateOverridesContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutOverridesControllerGenUpdate200Response**](ConsoleV1HoldoutOverridesControllerGenUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutsControllerGenCreate

> ConsoleV1HoldoutsControllerGenCreate200Response ConsoleV1HoldoutsControllerGenCreate(ctx).HoldoutCreateContractDto(holdoutCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create holdout

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
	holdoutCreateContractDto := *openapiclient.NewHoldoutCreateContractDto("team holdout") // HoldoutCreateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutsControllerGenCreate(context.Background()).HoldoutCreateContractDto(holdoutCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutsControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutsControllerGenCreate`: ConsoleV1HoldoutsControllerGenCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutsControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutsControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **holdoutCreateContractDto** | [**HoldoutCreateContractDto**](HoldoutCreateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutsControllerGenCreate200Response**](ConsoleV1HoldoutsControllerGenCreate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutsControllerGenFullUpdate

> ConsoleV1HoldoutsControllerGenFullUpdate200Response ConsoleV1HoldoutsControllerGenFullUpdate(ctx, id).HoldoutFullUpdateContractDto(holdoutFullUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update holdout by id

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
	holdoutFullUpdateContractDto := *openapiclient.NewHoldoutFullUpdateContractDto(true, "example holdout description", float32(5), []string{"GateIDs_example"}, []string{"ExperimentIDs_example"}, []string{"LayerIDs_example"}, false, "TODO") // HoldoutFullUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutsControllerGenFullUpdate(context.Background(), id).HoldoutFullUpdateContractDto(holdoutFullUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutsControllerGenFullUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutsControllerGenFullUpdate`: ConsoleV1HoldoutsControllerGenFullUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutsControllerGenFullUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutsControllerGenFullUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **holdoutFullUpdateContractDto** | [**HoldoutFullUpdateContractDto**](HoldoutFullUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutsControllerGenFullUpdate200Response**](ConsoleV1HoldoutsControllerGenFullUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutsControllerGenList

> ConsoleV1HoldoutsControllerGenList200Response ConsoleV1HoldoutsControllerGenList(ctx).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Holdouts

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
	creatorName := "creatorName_example" // string | Name of the creator. (optional)
	creatorID := "creatorID_example" // string | ID of the user who created the entity. (optional)
	tags := []string{"Inner_example"} // []string | Filter by tags (optional)
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutsControllerGenList(context.Background()).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutsControllerGenList`: ConsoleV1HoldoutsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutsControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creatorName** | **string** | Name of the creator. | 
 **creatorID** | **string** | ID of the user who created the entity. | 
 **tags** | **[]string** | Filter by tags | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutsControllerGenList200Response**](ConsoleV1HoldoutsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutsControllerGenPartialUpdate

> ConsoleV1HoldoutsControllerGenPartialUpdate200Response ConsoleV1HoldoutsControllerGenPartialUpdate(ctx, id).HoldoutPartialUpdateContractDto(holdoutPartialUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Patch holdout by id

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
	holdoutPartialUpdateContractDto := *openapiclient.NewHoldoutPartialUpdateContractDto() // HoldoutPartialUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutsControllerGenPartialUpdate(context.Background(), id).HoldoutPartialUpdateContractDto(holdoutPartialUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutsControllerGenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutsControllerGenPartialUpdate`: ConsoleV1HoldoutsControllerGenPartialUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutsControllerGenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutsControllerGenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **holdoutPartialUpdateContractDto** | [**HoldoutPartialUpdateContractDto**](HoldoutPartialUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutsControllerGenPartialUpdate200Response**](ConsoleV1HoldoutsControllerGenPartialUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutsControllerGenRead

> ConsoleV1HoldoutsControllerGenRead200Response ConsoleV1HoldoutsControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get holdout by id

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
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutsControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutsControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutsControllerGenRead`: ConsoleV1HoldoutsControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutsControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutsControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutsControllerGenRead200Response**](ConsoleV1HoldoutsControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1HoldoutsControllerGenRemove

> ConsoleV1HoldoutsControllerGenRemove200Response ConsoleV1HoldoutsControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete holdout by id

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
	resp, r, err := apiClient.HoldoutsAPI.ConsoleV1HoldoutsControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsAPI.ConsoleV1HoldoutsControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1HoldoutsControllerGenRemove`: ConsoleV1HoldoutsControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsAPI.ConsoleV1HoldoutsControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1HoldoutsControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1HoldoutsControllerGenRemove200Response**](ConsoleV1HoldoutsControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

