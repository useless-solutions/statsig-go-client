# \MetricsWarehouseNativeAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1MetricsControllerGenCreateMetricSource**](MetricsWarehouseNativeAPI.md#ConsoleV1MetricsControllerGenCreateMetricSource) | **Post** /console/v1/metrics/metric_source | Create Metric Source
[**ConsoleV1MetricsControllerGenDeleteMetricSource**](MetricsWarehouseNativeAPI.md#ConsoleV1MetricsControllerGenDeleteMetricSource) | **Delete** /console/v1/metrics/metric_source/{name} | Delete Metric Source
[**ConsoleV1MetricsControllerGenListMetricSources**](MetricsWarehouseNativeAPI.md#ConsoleV1MetricsControllerGenListMetricSources) | **Get** /console/v1/metrics/metric_source/list | List metric source
[**ConsoleV1MetricsControllerGenReadMetricSource**](MetricsWarehouseNativeAPI.md#ConsoleV1MetricsControllerGenReadMetricSource) | **Get** /console/v1/metrics/metric_source/{name} | Read Metric Source
[**ConsoleV1MetricsControllerReloadMetric**](MetricsWarehouseNativeAPI.md#ConsoleV1MetricsControllerReloadMetric) | **Post** /console/v1/metrics/{id}/reload | Reload metric data
[**ConsoleV1MetricsControllerUpdateMetricSource**](MetricsWarehouseNativeAPI.md#ConsoleV1MetricsControllerUpdateMetricSource) | **Post** /console/v1/metrics/metric_source/{name} | Update Metric Source



## ConsoleV1MetricsControllerGenCreateMetricSource

> ConsoleV1MetricsControllerGenCreateMetricSource201Response ConsoleV1MetricsControllerGenCreateMetricSource(ctx).MetricSourceCreationContractDto(metricSourceCreationContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Metric Source

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
	metricSourceCreationContractDto := *openapiclient.NewMetricSourceCreationContractDto("Name_example", "Sql_example", "TimestampColumn_example", []openapiclient.MetricSourceContractDtoIdTypeMappingInner{*openapiclient.NewMetricSourceContractDtoIdTypeMappingInner("StatsigUnitID_example", "Column_example")}) // MetricSourceCreationContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenCreateMetricSource(context.Background()).MetricSourceCreationContractDto(metricSourceCreationContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenCreateMetricSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenCreateMetricSource`: ConsoleV1MetricsControllerGenCreateMetricSource201Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenCreateMetricSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenCreateMetricSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSourceCreationContractDto** | [**MetricSourceCreationContractDto**](MetricSourceCreationContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenCreateMetricSource201Response**](ConsoleV1MetricsControllerGenCreateMetricSource201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenDeleteMetricSource

> ConsoleV1MetricsControllerGenDeleteMetricSource200Response ConsoleV1MetricsControllerGenDeleteMetricSource(ctx, name).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Metric Source

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
	resp, r, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenDeleteMetricSource(context.Background(), name).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenDeleteMetricSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenDeleteMetricSource`: ConsoleV1MetricsControllerGenDeleteMetricSource200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenDeleteMetricSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenDeleteMetricSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenDeleteMetricSource200Response**](ConsoleV1MetricsControllerGenDeleteMetricSource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenListMetricSources

> ConsoleV1MetricsControllerGenListMetricSources200Response ConsoleV1MetricsControllerGenListMetricSources(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List metric source

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
	resp, r, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenListMetricSources(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenListMetricSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenListMetricSources`: ConsoleV1MetricsControllerGenListMetricSources200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenListMetricSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenListMetricSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenListMetricSources200Response**](ConsoleV1MetricsControllerGenListMetricSources200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenReadMetricSource

> ConsoleV1MetricsControllerGenReadMetricSource200Response ConsoleV1MetricsControllerGenReadMetricSource(ctx, name).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Metric Source

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
	resp, r, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenReadMetricSource(context.Background(), name).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenReadMetricSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenReadMetricSource`: ConsoleV1MetricsControllerGenReadMetricSource200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerGenReadMetricSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenReadMetricSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenReadMetricSource200Response**](ConsoleV1MetricsControllerGenReadMetricSource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerReloadMetric

> ConsoleV1MetricsControllerReloadMetric200Response ConsoleV1MetricsControllerReloadMetric(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Reload metric data

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
	resp, r, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerReloadMetric(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerReloadMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerReloadMetric`: ConsoleV1MetricsControllerReloadMetric200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerReloadMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerReloadMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerReloadMetric200Response**](ConsoleV1MetricsControllerReloadMetric200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerUpdateMetricSource

> ConsoleV1MetricsControllerUpdateMetricSource200Response ConsoleV1MetricsControllerUpdateMetricSource(ctx, name).MetricSourceUpdateContractDto(metricSourceUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Metric Source

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
	metricSourceUpdateContractDto := *openapiclient.NewMetricSourceUpdateContractDto("Sql_example", "TimestampColumn_example", []openapiclient.MetricSourceContractDtoIdTypeMappingInner{*openapiclient.NewMetricSourceContractDtoIdTypeMappingInner("StatsigUnitID_example", "Column_example")}) // MetricSourceUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerUpdateMetricSource(context.Background(), name).MetricSourceUpdateContractDto(metricSourceUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerUpdateMetricSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerUpdateMetricSource`: ConsoleV1MetricsControllerUpdateMetricSource200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsWarehouseNativeAPI.ConsoleV1MetricsControllerUpdateMetricSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerUpdateMetricSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricSourceUpdateContractDto** | [**MetricSourceUpdateContractDto**](MetricSourceUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerUpdateMetricSource200Response**](ConsoleV1MetricsControllerUpdateMetricSource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

