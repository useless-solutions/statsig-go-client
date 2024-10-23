# \KeysAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1KeysControllerGenCreate**](KeysAPI.md#ConsoleV1KeysControllerGenCreate) | **Post** /console/v1/keys | Create Key
[**ConsoleV1KeysControllerGenDeactivate**](KeysAPI.md#ConsoleV1KeysControllerGenDeactivate) | **Patch** /console/v1/keys/{id}/deactivate | Deactivate Key
[**ConsoleV1KeysControllerGenDelete**](KeysAPI.md#ConsoleV1KeysControllerGenDelete) | **Delete** /console/v1/keys/{id} | Delete Key
[**ConsoleV1KeysControllerGenList**](KeysAPI.md#ConsoleV1KeysControllerGenList) | **Get** /console/v1/keys | List Keys
[**ConsoleV1KeysControllerGenRead**](KeysAPI.md#ConsoleV1KeysControllerGenRead) | **Get** /console/v1/keys/{id} | Read Key
[**ConsoleV1KeysControllerGenRotate**](KeysAPI.md#ConsoleV1KeysControllerGenRotate) | **Patch** /console/v1/keys/{id}/rotate | Rotate Key
[**ConsoleV1KeysControllerGenUpdate**](KeysAPI.md#ConsoleV1KeysControllerGenUpdate) | **Patch** /console/v1/keys/{id} | Update Key



## ConsoleV1KeysControllerGenCreate

> ConsoleV1KeysControllerGenCreate200Response ConsoleV1KeysControllerGenCreate(ctx).KeyCreateContractDto(keyCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Key

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
	keyCreateContractDto := *openapiclient.NewKeyCreateContractDto("Description_example", "Type_example") // KeyCreateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ConsoleV1KeysControllerGenCreate(context.Background()).KeyCreateContractDto(keyCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ConsoleV1KeysControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1KeysControllerGenCreate`: ConsoleV1KeysControllerGenCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ConsoleV1KeysControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1KeysControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyCreateContractDto** | [**KeyCreateContractDto**](KeyCreateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1KeysControllerGenCreate200Response**](ConsoleV1KeysControllerGenCreate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1KeysControllerGenDeactivate

> ConsoleV1KeysControllerGenDeactivate200Response ConsoleV1KeysControllerGenDeactivate(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Deactivate Key

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
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ConsoleV1KeysControllerGenDeactivate(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ConsoleV1KeysControllerGenDeactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1KeysControllerGenDeactivate`: ConsoleV1KeysControllerGenDeactivate200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ConsoleV1KeysControllerGenDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1KeysControllerGenDeactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1KeysControllerGenDeactivate200Response**](ConsoleV1KeysControllerGenDeactivate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1KeysControllerGenDelete

> ConsoleV1KeysControllerGenDelete200Response ConsoleV1KeysControllerGenDelete(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Key

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
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ConsoleV1KeysControllerGenDelete(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ConsoleV1KeysControllerGenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1KeysControllerGenDelete`: ConsoleV1KeysControllerGenDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ConsoleV1KeysControllerGenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1KeysControllerGenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1KeysControllerGenDelete200Response**](ConsoleV1KeysControllerGenDelete200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1KeysControllerGenList

> ConsoleV1KeysControllerGenList200Response ConsoleV1KeysControllerGenList(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Keys

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
	resp, r, err := apiClient.KeysAPI.ConsoleV1KeysControllerGenList(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ConsoleV1KeysControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1KeysControllerGenList`: ConsoleV1KeysControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ConsoleV1KeysControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1KeysControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1KeysControllerGenList200Response**](ConsoleV1KeysControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1KeysControllerGenRead

> ConsoleV1KeysControllerGenRead200Response ConsoleV1KeysControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Key

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
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ConsoleV1KeysControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ConsoleV1KeysControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1KeysControllerGenRead`: ConsoleV1KeysControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ConsoleV1KeysControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1KeysControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1KeysControllerGenRead200Response**](ConsoleV1KeysControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1KeysControllerGenRotate

> ConsoleV1KeysControllerGenRotate200Response ConsoleV1KeysControllerGenRotate(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Rotate Key

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
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ConsoleV1KeysControllerGenRotate(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ConsoleV1KeysControllerGenRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1KeysControllerGenRotate`: ConsoleV1KeysControllerGenRotate200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ConsoleV1KeysControllerGenRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1KeysControllerGenRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1KeysControllerGenRotate200Response**](ConsoleV1KeysControllerGenRotate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1KeysControllerGenUpdate

> ConsoleV1KeysControllerGenUpdate200Response ConsoleV1KeysControllerGenUpdate(ctx, id).KeyUpdateContractDto(keyUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Key

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
	keyUpdateContractDto := *openapiclient.NewKeyUpdateContractDto() // KeyUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ConsoleV1KeysControllerGenUpdate(context.Background(), id).KeyUpdateContractDto(keyUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ConsoleV1KeysControllerGenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1KeysControllerGenUpdate`: ConsoleV1KeysControllerGenUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ConsoleV1KeysControllerGenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1KeysControllerGenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyUpdateContractDto** | [**KeyUpdateContractDto**](KeyUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1KeysControllerGenUpdate200Response**](ConsoleV1KeysControllerGenUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

