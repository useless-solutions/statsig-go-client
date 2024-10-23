# \LayersAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1LayerExperimentsControllerGenList**](LayersAPI.md#ConsoleV1LayerExperimentsControllerGenList) | **Get** /console/v1/layers/{id}/experiments | Lineage: List Experiment related to Layer
[**ConsoleV1LayerOverridesControllerGenRead**](LayersAPI.md#ConsoleV1LayerOverridesControllerGenRead) | **Get** /console/v1/layers/{id}/overrides | Get Layer Overrides
[**ConsoleV1LayersControllerGenCreate**](LayersAPI.md#ConsoleV1LayersControllerGenCreate) | **Post** /console/v1/layers | Create a Layer
[**ConsoleV1LayersControllerGenFullUpdate**](LayersAPI.md#ConsoleV1LayersControllerGenFullUpdate) | **Post** /console/v1/layers/{id} | Update a layer
[**ConsoleV1LayersControllerGenList**](LayersAPI.md#ConsoleV1LayersControllerGenList) | **Get** /console/v1/layers | Get Layers
[**ConsoleV1LayersControllerGenPartialUpdate**](LayersAPI.md#ConsoleV1LayersControllerGenPartialUpdate) | **Patch** /console/v1/layers/{id} | Partially update a layer
[**ConsoleV1LayersControllerGenRead**](LayersAPI.md#ConsoleV1LayersControllerGenRead) | **Get** /console/v1/layers/{id} | Get one layer
[**ConsoleV1LayersControllerGenRemove**](LayersAPI.md#ConsoleV1LayersControllerGenRemove) | **Delete** /console/v1/layers/{id} | Delete a layer



## ConsoleV1LayerExperimentsControllerGenList

> ConsoleV1LayerExperimentsControllerGenList200Response ConsoleV1LayerExperimentsControllerGenList(ctx, id).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

Lineage: List Experiment related to Layer

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
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayerExperimentsControllerGenList(context.Background(), id).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayerExperimentsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayerExperimentsControllerGenList`: ConsoleV1LayerExperimentsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayerExperimentsControllerGenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayerExperimentsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayerExperimentsControllerGenList200Response**](ConsoleV1LayerExperimentsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1LayerOverridesControllerGenRead

> ConsoleV1LayerOverridesControllerGenRead200Response ConsoleV1LayerOverridesControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Layer Overrides

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
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayerOverridesControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayerOverridesControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayerOverridesControllerGenRead`: ConsoleV1LayerOverridesControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayerOverridesControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayerOverridesControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayerOverridesControllerGenRead200Response**](ConsoleV1LayerOverridesControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1LayersControllerGenCreate

> ConsoleV1LayersControllerGenCreate201Response ConsoleV1LayersControllerGenCreate(ctx).LayerCreateContractDto(layerCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create a Layer

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
	layerCreateContractDto := *openapiclient.NewLayerCreateContractDto("Name_example", "IdType_example") // LayerCreateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayersControllerGenCreate(context.Background()).LayerCreateContractDto(layerCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayersControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayersControllerGenCreate`: ConsoleV1LayersControllerGenCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayersControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayersControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerCreateContractDto** | [**LayerCreateContractDto**](LayerCreateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayersControllerGenCreate201Response**](ConsoleV1LayersControllerGenCreate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1LayersControllerGenFullUpdate

> ConsoleV1LayersControllerGenFullUpdate200Response ConsoleV1LayersControllerGenFullUpdate(ctx, id).LayerFullUpdateContractDto(layerFullUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update a layer

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
	layerFullUpdateContractDto := *openapiclient.NewLayerFullUpdateContractDto("Description_example", []openapiclient.LayerContractDtoParametersInner{*openapiclient.NewLayerContractDtoParametersInner("Name_example", "Type_example", openapiclient.LayerContractDto_parameters_inner_defaultValue{ArrayOfLayerContractDtoParametersInnerDefaultValueOneOfInner: new([]LayerContractDtoParametersInnerDefaultValueOneOfInner)})}) // LayerFullUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayersControllerGenFullUpdate(context.Background(), id).LayerFullUpdateContractDto(layerFullUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayersControllerGenFullUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayersControllerGenFullUpdate`: ConsoleV1LayersControllerGenFullUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayersControllerGenFullUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayersControllerGenFullUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **layerFullUpdateContractDto** | [**LayerFullUpdateContractDto**](LayerFullUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayersControllerGenFullUpdate200Response**](ConsoleV1LayersControllerGenFullUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1LayersControllerGenList

> ConsoleV1LayersControllerGenList200Response ConsoleV1LayersControllerGenList(ctx).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Layers

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
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayersControllerGenList(context.Background()).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayersControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayersControllerGenList`: ConsoleV1LayersControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayersControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayersControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creatorName** | **string** | Name of the creator. | 
 **creatorID** | **string** | ID of the user who created the entity. | 
 **tags** | **[]string** | Filter by tags | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayersControllerGenList200Response**](ConsoleV1LayersControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1LayersControllerGenPartialUpdate

> ConsoleV1LayersControllerGenPartialUpdate200Response ConsoleV1LayersControllerGenPartialUpdate(ctx, id).LayerPartialUpdateContractDto(layerPartialUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Partially update a layer

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
	layerPartialUpdateContractDto := *openapiclient.NewLayerPartialUpdateContractDto() // LayerPartialUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayersControllerGenPartialUpdate(context.Background(), id).LayerPartialUpdateContractDto(layerPartialUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayersControllerGenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayersControllerGenPartialUpdate`: ConsoleV1LayersControllerGenPartialUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayersControllerGenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayersControllerGenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **layerPartialUpdateContractDto** | [**LayerPartialUpdateContractDto**](LayerPartialUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayersControllerGenPartialUpdate200Response**](ConsoleV1LayersControllerGenPartialUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1LayersControllerGenRead

> ConsoleV1LayersControllerGenRead200Response ConsoleV1LayersControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get one layer

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
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayersControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayersControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayersControllerGenRead`: ConsoleV1LayersControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayersControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayersControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayersControllerGenRead200Response**](ConsoleV1LayersControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1LayersControllerGenRemove

> ConsoleV1LayersControllerGenRemove200Response ConsoleV1LayersControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete a layer

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
	resp, r, err := apiClient.LayersAPI.ConsoleV1LayersControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersAPI.ConsoleV1LayersControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1LayersControllerGenRemove`: ConsoleV1LayersControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `LayersAPI.ConsoleV1LayersControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1LayersControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1LayersControllerGenRemove200Response**](ConsoleV1LayersControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

