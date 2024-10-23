# \MetricsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1MetricValuesControllerGenListMetricValues**](MetricsAPI.md#ConsoleV1MetricValuesControllerGenListMetricValues) | **Get** /console/v1/metrics/values | List All Metric Values
[**ConsoleV1MetricsControllerDeleteMetric**](MetricsAPI.md#ConsoleV1MetricsControllerDeleteMetric) | **Delete** /console/v1/metrics/{id} | Delete a metric
[**ConsoleV1MetricsControllerGenCreate**](MetricsAPI.md#ConsoleV1MetricsControllerGenCreate) | **Post** /console/v1/metrics | Create Metric
[**ConsoleV1MetricsControllerGenExperimentListFromMetric**](MetricsAPI.md#ConsoleV1MetricsControllerGenExperimentListFromMetric) | **Get** /console/v1/metrics/{id}/experiments | Lineage: List experiments related to Metric
[**ConsoleV1MetricsControllerGenListMetrics**](MetricsAPI.md#ConsoleV1MetricsControllerGenListMetrics) | **Get** /console/v1/metrics/list | List all Metrics
[**ConsoleV1MetricsControllerGenReadMetric**](MetricsAPI.md#ConsoleV1MetricsControllerGenReadMetric) | **Get** /console/v1/metrics | Read Single Metric Value
[**ConsoleV1MetricsControllerGenReadMetricDefinition**](MetricsAPI.md#ConsoleV1MetricsControllerGenReadMetricDefinition) | **Get** /console/v1/metrics/{id} | Read Metric Definition
[**ConsoleV1MetricsControllerUpdateMetric**](MetricsAPI.md#ConsoleV1MetricsControllerUpdateMetric) | **Post** /console/v1/metrics/{id} | Update a metric



## ConsoleV1MetricValuesControllerGenListMetricValues

> ConsoleV1MetricValuesControllerGenListMetricValues200Response ConsoleV1MetricValuesControllerGenListMetricValues(ctx).Date(date).MetricName(metricName).MetricType(metricType).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List All Metric Values



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
	date := "date_example" // string | 
	metricName := "metricName_example" // string |  (optional)
	metricType := "metricType_example" // string |  (optional)
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricValuesControllerGenListMetricValues(context.Background()).Date(date).MetricName(metricName).MetricType(metricType).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricValuesControllerGenListMetricValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricValuesControllerGenListMetricValues`: ConsoleV1MetricValuesControllerGenListMetricValues200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricValuesControllerGenListMetricValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricValuesControllerGenListMetricValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** |  | 
 **metricName** | **string** |  | 
 **metricType** | **string** |  | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricValuesControllerGenListMetricValues200Response**](ConsoleV1MetricValuesControllerGenListMetricValues200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerDeleteMetric

> ConsoleV1MetricsControllerDeleteMetric200Response ConsoleV1MetricsControllerDeleteMetric(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete a metric

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
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricsControllerDeleteMetric(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricsControllerDeleteMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerDeleteMetric`: ConsoleV1MetricsControllerDeleteMetric200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricsControllerDeleteMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerDeleteMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerDeleteMetric200Response**](ConsoleV1MetricsControllerDeleteMetric200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenCreate

> ConsoleV1MetricsControllerGenCreate201Response ConsoleV1MetricsControllerGenCreate(ctx).MetricCreationContractDto(metricCreationContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Metric

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
	metricCreationContractDto := *openapiclient.NewMetricCreationContractDto("Name_example", "Type_example") // MetricCreationContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricsControllerGenCreate(context.Background()).MetricCreationContractDto(metricCreationContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricsControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenCreate`: ConsoleV1MetricsControllerGenCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricsControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricCreationContractDto** | [**MetricCreationContractDto**](MetricCreationContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenCreate201Response**](ConsoleV1MetricsControllerGenCreate201Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenExperimentListFromMetric

> ConsoleV1MetricsControllerGenExperimentListFromMetric200Response ConsoleV1MetricsControllerGenExperimentListFromMetric(ctx, id).LayerID(layerID).IdType(idType).Status(status).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

Lineage: List experiments related to Metric

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
	layerID := "layerID_example" // string | Which layer to place the experiment into. (optional)
	idType := "idType_example" // string | The idType the experiment will be performed on (optional)
	status := TODO // interface{} | The current status of the experiment (optional)
	creatorName := "creatorName_example" // string | Name of the creator. (optional)
	creatorID := "creatorID_example" // string | ID of the user who created the entity. (optional)
	tags := []string{"Inner_example"} // []string | Filter by tags (optional)
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricsControllerGenExperimentListFromMetric(context.Background(), id).LayerID(layerID).IdType(idType).Status(status).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricsControllerGenExperimentListFromMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenExperimentListFromMetric`: ConsoleV1MetricsControllerGenExperimentListFromMetric200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricsControllerGenExperimentListFromMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenExperimentListFromMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **layerID** | **string** | Which layer to place the experiment into. | 
 **idType** | **string** | The idType the experiment will be performed on | 
 **status** | [**interface{}**](interface{}.md) | The current status of the experiment | 
 **creatorName** | **string** | Name of the creator. | 
 **creatorID** | **string** | ID of the user who created the entity. | 
 **tags** | **[]string** | Filter by tags | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenExperimentListFromMetric200Response**](ConsoleV1MetricsControllerGenExperimentListFromMetric200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenListMetrics

> ConsoleV1MetricsControllerGenListMetrics200Response ConsoleV1MetricsControllerGenListMetrics(ctx).ShowHiddenMetrics(showHiddenMetrics).Tags(tags).Filters(filters).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List all Metrics

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
	showHiddenMetrics := "showHiddenMetrics_example" // string | Should hidden metrics be returned: Allowed values are \"true\" or \"false\". (optional)
	tags := TODO // interface{} | Filter metrics based on a given tagID, found on /tags endpoint. Can be a single string or an array of strings. (optional)
	filters := TODO // interface{} | Additional filters for metrics. Can be a string or an object with tags filter. (optional)
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricsControllerGenListMetrics(context.Background()).ShowHiddenMetrics(showHiddenMetrics).Tags(tags).Filters(filters).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricsControllerGenListMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenListMetrics`: ConsoleV1MetricsControllerGenListMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricsControllerGenListMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenListMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showHiddenMetrics** | **string** | Should hidden metrics be returned: Allowed values are \&quot;true\&quot; or \&quot;false\&quot;. | 
 **tags** | [**interface{}**](interface{}.md) | Filter metrics based on a given tagID, found on /tags endpoint. Can be a single string or an array of strings. | 
 **filters** | [**interface{}**](interface{}.md) | Additional filters for metrics. Can be a string or an object with tags filter. | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenListMetrics200Response**](ConsoleV1MetricsControllerGenListMetrics200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenReadMetric

> ConsoleV1MetricsControllerGenReadMetric200Response ConsoleV1MetricsControllerGenReadMetric(ctx).Id(id).Date(date).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Single Metric Value

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
	id := "id_example" // string | The unique identifier of the metric
	date := "date_example" // string | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricsControllerGenReadMetric(context.Background()).Id(id).Date(date).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricsControllerGenReadMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenReadMetric`: ConsoleV1MetricsControllerGenReadMetric200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricsControllerGenReadMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenReadMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The unique identifier of the metric | 
 **date** | **string** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenReadMetric200Response**](ConsoleV1MetricsControllerGenReadMetric200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerGenReadMetricDefinition

> ConsoleV1MetricsControllerGenReadMetricDefinition200Response ConsoleV1MetricsControllerGenReadMetricDefinition(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Metric Definition

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
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricsControllerGenReadMetricDefinition(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricsControllerGenReadMetricDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerGenReadMetricDefinition`: ConsoleV1MetricsControllerGenReadMetricDefinition200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricsControllerGenReadMetricDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerGenReadMetricDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerGenReadMetricDefinition200Response**](ConsoleV1MetricsControllerGenReadMetricDefinition200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1MetricsControllerUpdateMetric

> ConsoleV1MetricsControllerUpdateMetric200Response ConsoleV1MetricsControllerUpdateMetric(ctx, id).MetricsUpdateContractDto(metricsUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update a metric

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
	metricsUpdateContractDto := *openapiclient.NewMetricsUpdateContractDto() // MetricsUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ConsoleV1MetricsControllerUpdateMetric(context.Background(), id).MetricsUpdateContractDto(metricsUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ConsoleV1MetricsControllerUpdateMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1MetricsControllerUpdateMetric`: ConsoleV1MetricsControllerUpdateMetric200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ConsoleV1MetricsControllerUpdateMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1MetricsControllerUpdateMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsUpdateContractDto** | [**MetricsUpdateContractDto**](MetricsUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1MetricsControllerUpdateMetric200Response**](ConsoleV1MetricsControllerUpdateMetric200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

