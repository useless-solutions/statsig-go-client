# \ExperimentsWarehouseNativeAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1ExperimentsControllerGenCreateAssignmentSource**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenCreateAssignmentSource) | **Post** /console/v1/experiments/assignment_sources | Create Assignment Source
[**ConsoleV1ExperimentsControllerGenCreateEntityPropertySource**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenCreateEntityPropertySource) | **Post** /console/v1/experiments/entity_properties | Create Entity Property Source
[**ConsoleV1ExperimentsControllerGenEntityPropertySource**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenEntityPropertySource) | **Get** /console/v1/experiments/entity_property/{name} | Get Entity Property Source
[**ConsoleV1ExperimentsControllerGenListAssignmentSources**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenListAssignmentSources) | **Get** /console/v1/experiments/assignment_sources | List Assignment Sources
[**ConsoleV1ExperimentsControllerGenListEntityPropertySources**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenListEntityPropertySources) | **Get** /console/v1/experiments/entity_properties | List Entity Property Sources
[**ConsoleV1ExperimentsControllerGenLoadPulse**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenLoadPulse) | **Post** /console/v1/experiments/{id}/load_pulse | Load Pulse (Warehouse Native)
[**ConsoleV1ExperimentsControllerGenPulseLoadHistory**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenPulseLoadHistory) | **Get** /console/v1/experiments/{id}/pulse_load_history | Pulse Load History (Warehouse Native)
[**ConsoleV1ExperimentsControllerGenRemoveAssignmentSource**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenRemoveAssignmentSource) | **Delete** /console/v1/experiments/assignment_source/{name} | Delete Assignment Source
[**ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource) | **Delete** /console/v1/experiments/entity_property/{name} | Delete Entity Property Source
[**ConsoleV1ExperimentsControllerGenUpdateAssignmentSource**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenUpdateAssignmentSource) | **Patch** /console/v1/experiments/assignment_source/{name} | Patch Assignment Source
[**ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery) | **Post** /console/v1/experiments/assignment_source/{name} | Post Assignment Source
[**ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource) | **Patch** /console/v1/experiments/entity_property/{name} | Patch Entity Property Source
[**ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery**](ExperimentsWarehouseNativeAPI.md#ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery) | **Post** /console/v1/experiments/entity_property/{name} | Post Entity Property Source



## ConsoleV1ExperimentsControllerGenCreateAssignmentSource

> ConsoleV1ExperimentsControllerGenCreateAssignmentSource201Response ConsoleV1ExperimentsControllerGenCreateAssignmentSource(ctx).AssignmentSourceCreationDto(assignmentSourceCreationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Assignment Source

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
	assignmentSourceCreationDto := *openapiclient.NewAssignmentSourceCreationDto("Name_example", "Sql_example", "TimestampColumn_example", "ExperimentIDColumn_example", "GroupIDColumn_example", []openapiclient.AssignmentSourceCreationDtoIdTypeMappingInner{*openapiclient.NewAssignmentSourceCreationDtoIdTypeMappingInner("StatsigUnitID_example", "Column_example")}) // AssignmentSourceCreationDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenCreateAssignmentSource(context.Background()).AssignmentSourceCreationDto(assignmentSourceCreationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenCreateAssignmentSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenCreateAssignmentSource`: ConsoleV1ExperimentsControllerGenCreateAssignmentSource201Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenCreateAssignmentSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenCreateAssignmentSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignmentSourceCreationDto** | [**AssignmentSourceCreationDto**](AssignmentSourceCreationDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenCreateAssignmentSource201Response**](ConsoleV1ExperimentsControllerGenCreateAssignmentSource201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenCreateEntityPropertySource

> ConsoleV1ExperimentsControllerGenCreateEntityPropertySource201Response ConsoleV1ExperimentsControllerGenCreateEntityPropertySource(ctx).EntityPropertySourceCreationDto(entityPropertySourceCreationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Entity Property Source

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
	entityPropertySourceCreationDto := *openapiclient.NewEntityPropertySourceCreationDto("Name_example", "Sql_example", []openapiclient.EntityPropertySourceDtoIdTypeMappingInner{*openapiclient.NewEntityPropertySourceDtoIdTypeMappingInner("StatsigUnitID_example", "Column_example")}) // EntityPropertySourceCreationDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenCreateEntityPropertySource(context.Background()).EntityPropertySourceCreationDto(entityPropertySourceCreationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenCreateEntityPropertySource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenCreateEntityPropertySource`: ConsoleV1ExperimentsControllerGenCreateEntityPropertySource201Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenCreateEntityPropertySource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenCreateEntityPropertySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityPropertySourceCreationDto** | [**EntityPropertySourceCreationDto**](EntityPropertySourceCreationDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenCreateEntityPropertySource201Response**](ConsoleV1ExperimentsControllerGenCreateEntityPropertySource201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenEntityPropertySource

> ConsoleV1ExperimentsControllerGenEntityPropertySource200Response ConsoleV1ExperimentsControllerGenEntityPropertySource(ctx, name).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Entity Property Source

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
	name := "name_example" // string | Name of entity property source
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenEntityPropertySource(context.Background(), name).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenEntityPropertySource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenEntityPropertySource`: ConsoleV1ExperimentsControllerGenEntityPropertySource200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenEntityPropertySource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of entity property source | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenEntityPropertySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenEntityPropertySource200Response**](ConsoleV1ExperimentsControllerGenEntityPropertySource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenListAssignmentSources

> ConsoleV1ExperimentsControllerGenListAssignmentSources200Response ConsoleV1ExperimentsControllerGenListAssignmentSources(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Assignment Sources

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
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenListAssignmentSources(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenListAssignmentSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenListAssignmentSources`: ConsoleV1ExperimentsControllerGenListAssignmentSources200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenListAssignmentSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenListAssignmentSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenListAssignmentSources200Response**](ConsoleV1ExperimentsControllerGenListAssignmentSources200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenListEntityPropertySources

> ConsoleV1ExperimentsControllerGenListEntityPropertySources200Response ConsoleV1ExperimentsControllerGenListEntityPropertySources(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Entity Property Sources

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
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenListEntityPropertySources(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenListEntityPropertySources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenListEntityPropertySources`: ConsoleV1ExperimentsControllerGenListEntityPropertySources200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenListEntityPropertySources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenListEntityPropertySourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenListEntityPropertySources200Response**](ConsoleV1ExperimentsControllerGenListEntityPropertySources200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenLoadPulse

> ConsoleV1ExperimentsControllerGenLoadPulse200Response ConsoleV1ExperimentsControllerGenLoadPulse(ctx, id).EchidnaLoadPulseQueryDto(echidnaLoadPulseQueryDto).Refresh(refresh).RuleId(ruleId).TurboMode(turboMode).XRespectReviewSettings(xRespectReviewSettings).Execute()

Load Pulse (Warehouse Native)

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
	echidnaLoadPulseQueryDto := *openapiclient.NewEchidnaLoadPulseQueryDto() // EchidnaLoadPulseQueryDto | 
	refresh := "refresh_example" // string |  (optional) (default to "full")
	ruleId := "ruleId_example" // string |  (optional)
	turboMode := true // bool |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenLoadPulse(context.Background(), id).EchidnaLoadPulseQueryDto(echidnaLoadPulseQueryDto).Refresh(refresh).RuleId(ruleId).TurboMode(turboMode).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenLoadPulse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenLoadPulse`: ConsoleV1ExperimentsControllerGenLoadPulse200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenLoadPulse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenLoadPulseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **echidnaLoadPulseQueryDto** | [**EchidnaLoadPulseQueryDto**](EchidnaLoadPulseQueryDto.md) |  | 
 **refresh** | **string** |  | [default to &quot;full&quot;]
 **ruleId** | **string** |  | 
 **turboMode** | **bool** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenLoadPulse200Response**](ConsoleV1ExperimentsControllerGenLoadPulse200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenPulseLoadHistory

> ConsoleV1GatesControllerGenPulseLoadHistory200Response ConsoleV1ExperimentsControllerGenPulseLoadHistory(ctx, id).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

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
	id := "id_example" // string | id
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenPulseLoadHistory(context.Background(), id).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenPulseLoadHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenPulseLoadHistory`: ConsoleV1GatesControllerGenPulseLoadHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenPulseLoadHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenPulseLoadHistoryRequest struct via the builder pattern


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


## ConsoleV1ExperimentsControllerGenRemoveAssignmentSource

> ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response ConsoleV1ExperimentsControllerGenRemoveAssignmentSource(ctx, name).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Assignment Source

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
	name := "name_example" // string | Name of the assignment source
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenRemoveAssignmentSource(context.Background(), name).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenRemoveAssignmentSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenRemoveAssignmentSource`: ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenRemoveAssignmentSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the assignment source | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenRemoveAssignmentSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response**](ConsoleV1ExperimentsControllerGenRemoveAssignmentSource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource

> ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource200Response ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource(ctx, name).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Entity Property Source

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
	name := "name_example" // string | name
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource(context.Background(), name).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource`: ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenRemoveEntityPropertySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource200Response**](ConsoleV1ExperimentsControllerGenRemoveEntityPropertySource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenUpdateAssignmentSource

> ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response ConsoleV1ExperimentsControllerGenUpdateAssignmentSource(ctx, name).AssignmentSourcePartialUpdateDto(assignmentSourcePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Patch Assignment Source

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
	name := "name_example" // string | Name of the assignment source
	assignmentSourcePartialUpdateDto := *openapiclient.NewAssignmentSourcePartialUpdateDto() // AssignmentSourcePartialUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateAssignmentSource(context.Background(), name).AssignmentSourcePartialUpdateDto(assignmentSourcePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateAssignmentSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenUpdateAssignmentSource`: ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateAssignmentSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the assignment source | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenUpdateAssignmentSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignmentSourcePartialUpdateDto** | [**AssignmentSourcePartialUpdateDto**](AssignmentSourcePartialUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response**](ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery

> ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery(ctx, name).AssignmentSourceQueryUpdateDto(assignmentSourceQueryUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Post Assignment Source

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
	name := "name_example" // string | Name of the assignment source
	assignmentSourceQueryUpdateDto := *openapiclient.NewAssignmentSourceQueryUpdateDto("Sql_example", "TimestampColumn_example", "ExperimentIDColumn_example", "GroupIDColumn_example", []openapiclient.AssignmentSourceCreationDtoIdTypeMappingInner{*openapiclient.NewAssignmentSourceCreationDtoIdTypeMappingInner("StatsigUnitID_example", "Column_example")}) // AssignmentSourceQueryUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery(context.Background(), name).AssignmentSourceQueryUpdateDto(assignmentSourceQueryUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery`: ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the assignment source | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignmentSourceQueryUpdateDto** | [**AssignmentSourceQueryUpdateDto**](AssignmentSourceQueryUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response**](ConsoleV1ExperimentsControllerGenUpdateAssignmentSourceQuery200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource

> ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource200Response ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource(ctx, name).EntityPropertySourcePartialUpdateDto(entityPropertySourcePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Patch Entity Property Source

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
	name := "name_example" // string | Name of entity property source
	entityPropertySourcePartialUpdateDto := *openapiclient.NewEntityPropertySourcePartialUpdateDto() // EntityPropertySourcePartialUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource(context.Background(), name).EntityPropertySourcePartialUpdateDto(entityPropertySourcePartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource`: ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of entity property source | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityPropertySourcePartialUpdateDto** | [**EntityPropertySourcePartialUpdateDto**](EntityPropertySourcePartialUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource200Response**](ConsoleV1ExperimentsControllerGenUpdateEntityPropertySource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery

> ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery200Response ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery(ctx, name).EntityPropertySourceQueryUpdateDto(entityPropertySourceQueryUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Post Entity Property Source

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
	name := "name_example" // string | name
	entityPropertySourceQueryUpdateDto := *openapiclient.NewEntityPropertySourceQueryUpdateDto("Sql_example", []openapiclient.EntityPropertySourceDtoIdTypeMappingInner{*openapiclient.NewEntityPropertySourceDtoIdTypeMappingInner("StatsigUnitID_example", "Column_example")}) // EntityPropertySourceQueryUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery(context.Background(), name).EntityPropertySourceQueryUpdateDto(entityPropertySourceQueryUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery`: ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsWarehouseNativeAPI.ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityPropertySourceQueryUpdateDto** | [**EntityPropertySourceQueryUpdateDto**](EntityPropertySourceQueryUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery200Response**](ConsoleV1ExperimentsControllerGenUpdateEntityPropertySourceQuery200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

