# \IngestionsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks**](IngestionsAPI.md#ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks) | **Post** /console/v1/ingestion/connection/databricks | Create Ingestion Databricks
[**ConsoleV1IngestionControllerGenBackfillIngestion**](IngestionsAPI.md#ConsoleV1IngestionControllerGenBackfillIngestion) | **Post** /console/v1/ingestion/backfill | Backfill Ingestion
[**ConsoleV1IngestionControllerGenCreateIngestionSource**](IngestionsAPI.md#ConsoleV1IngestionControllerGenCreateIngestionSource) | **Post** /console/v1/ingestion | Create Ingestion Source
[**ConsoleV1IngestionControllerGenDeleteIngestionSource**](IngestionsAPI.md#ConsoleV1IngestionControllerGenDeleteIngestionSource) | **Delete** /console/v1/ingestion | Delete Ingestion Source
[**ConsoleV1IngestionControllerGenIngestionRun**](IngestionsAPI.md#ConsoleV1IngestionControllerGenIngestionRun) | **Get** /console/v1/ingestion/runs/{id} | Read Ingestion Run
[**ConsoleV1IngestionControllerGenIngestionRuns**](IngestionsAPI.md#ConsoleV1IngestionControllerGenIngestionRuns) | **Get** /console/v1/ingestion/runs | List Ingestion Runs
[**ConsoleV1IngestionControllerGenIngestionStatusList**](IngestionsAPI.md#ConsoleV1IngestionControllerGenIngestionStatusList) | **Get** /console/v1/ingestion/status | List Ingestions Status
[**ConsoleV1IngestionControllerGenReadIngestion**](IngestionsAPI.md#ConsoleV1IngestionControllerGenReadIngestion) | **Get** /console/v1/ingestion | Read Ingestion
[**ConsoleV1IngestionControllerGenReadIngestionSchedule**](IngestionsAPI.md#ConsoleV1IngestionControllerGenReadIngestionSchedule) | **Get** /console/v1/ingestion/schedule | Read Ingestion Schedule
[**ConsoleV1IngestionControllerGenUpdateIngestion**](IngestionsAPI.md#ConsoleV1IngestionControllerGenUpdateIngestion) | **Patch** /console/v1/ingestion | Update Ingestion Source
[**ConsoleV1IngestionControllerGenUpdateIngestionSchedule**](IngestionsAPI.md#ConsoleV1IngestionControllerGenUpdateIngestionSchedule) | **Post** /console/v1/ingestion/schedule | Update Ingestion Schedule
[**ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger**](IngestionsAPI.md#ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger) | **Get** /console/v1/ingestion/events/delta | Get Ingestion Event Delta Ledger
[**ConsoleV1IngestionEventsControllerGenIngestionEventCount**](IngestionsAPI.md#ConsoleV1IngestionEventsControllerGenIngestionEventCount) | **Get** /console/v1/ingestion/events/count | Get Ingestion Event Count



## ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks

> ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks200Response ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks(ctx).IngestionCreateDatabricksConnectionContractDto(ingestionCreateDatabricksConnectionContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Ingestion Databricks

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
	ingestionCreateDatabricksConnectionContractDto := *openapiclient.NewIngestionCreateDatabricksConnectionContractDto("Token_example", "Host_example", "Path_example") // IngestionCreateDatabricksConnectionContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks(context.Background()).IngestionCreateDatabricksConnectionContractDto(ingestionCreateDatabricksConnectionContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks`: ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestionCreateDatabricksConnectionContractDto** | [**IngestionCreateDatabricksConnectionContractDto**](IngestionCreateDatabricksConnectionContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks200Response**](ConsoleV1IngestionConnectionControllerGenCreateIngestionDatabricks200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenBackfillIngestion

> ConsoleV1IngestionControllerGenBackfillIngestion200Response ConsoleV1IngestionControllerGenBackfillIngestion(ctx).IngestionBackfillContractDto(ingestionBackfillContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Backfill Ingestion

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
	ingestionBackfillContractDto := *openapiclient.NewIngestionBackfillContractDto("DatestampStart_example", "DatestampEnd_example", "Type_example", "Dataset_example") // IngestionBackfillContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenBackfillIngestion(context.Background()).IngestionBackfillContractDto(ingestionBackfillContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenBackfillIngestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenBackfillIngestion`: ConsoleV1IngestionControllerGenBackfillIngestion200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenBackfillIngestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenBackfillIngestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestionBackfillContractDto** | [**IngestionBackfillContractDto**](IngestionBackfillContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenBackfillIngestion200Response**](ConsoleV1IngestionControllerGenBackfillIngestion200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenCreateIngestionSource

> ConsoleV1IngestionControllerGenCreateIngestionSource200Response ConsoleV1IngestionControllerGenCreateIngestionSource(ctx).IngestionSourceCreateContractDto(ingestionSourceCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Ingestion Source

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
	ingestionSourceCreateContractDto := openapiclient.IngestionSourceCreateContractDto{IngestionSourceCreateContractDtoOneOf: openapiclient.NewIngestionSourceCreateContractDtoOneOf("Dataset_example", "Type_example", "SourceName_example")} // IngestionSourceCreateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenCreateIngestionSource(context.Background()).IngestionSourceCreateContractDto(ingestionSourceCreateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenCreateIngestionSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenCreateIngestionSource`: ConsoleV1IngestionControllerGenCreateIngestionSource200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenCreateIngestionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenCreateIngestionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestionSourceCreateContractDto** | [**IngestionSourceCreateContractDto**](IngestionSourceCreateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenCreateIngestionSource200Response**](ConsoleV1IngestionControllerGenCreateIngestionSource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenDeleteIngestionSource

> ConsoleV1IngestionControllerGenDeleteIngestionSource200Response ConsoleV1IngestionControllerGenDeleteIngestionSource(ctx).Type_(type_).Dataset(dataset).SourceName(sourceName).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Ingestion Source

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
	type_ := "type__example" // string | 
	dataset := "dataset_example" // string | 
	sourceName := "sourceName_example" // string |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenDeleteIngestionSource(context.Background()).Type_(type_).Dataset(dataset).SourceName(sourceName).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenDeleteIngestionSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenDeleteIngestionSource`: ConsoleV1IngestionControllerGenDeleteIngestionSource200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenDeleteIngestionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenDeleteIngestionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **dataset** | **string** |  | 
 **sourceName** | **string** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenDeleteIngestionSource200Response**](ConsoleV1IngestionControllerGenDeleteIngestionSource200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenIngestionRun

> ConsoleV1IngestionControllerGenIngestionRun200Response ConsoleV1IngestionControllerGenIngestionRun(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Ingestion Run

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
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenIngestionRun(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenIngestionRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenIngestionRun`: ConsoleV1IngestionControllerGenIngestionRun200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenIngestionRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenIngestionRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenIngestionRun200Response**](ConsoleV1IngestionControllerGenIngestionRun200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenIngestionRuns

> ConsoleV1IngestionControllerGenIngestionRuns200Response ConsoleV1IngestionControllerGenIngestionRuns(ctx).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Ingestion Runs

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
	page := TODO // interface{} |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenIngestionRuns(context.Background()).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenIngestionRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenIngestionRuns`: ConsoleV1IngestionControllerGenIngestionRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenIngestionRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenIngestionRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**interface{}**](interface{}.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenIngestionRuns200Response**](ConsoleV1IngestionControllerGenIngestionRuns200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenIngestionStatusList

> ConsoleV1IngestionControllerGenIngestionStatusList200Response ConsoleV1IngestionControllerGenIngestionStatusList(ctx).StartDate(startDate).EndDate(endDate).Source(source).Dataset(dataset).Status(status).Statuses(statuses).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Ingestions Status

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
	startDate := "startDate_example" // string | 
	endDate := "endDate_example" // string | 
	source := "source_example" // string |  (optional)
	dataset := "dataset_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	statuses := []string{"Inner_example"} // []string |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenIngestionStatusList(context.Background()).StartDate(startDate).EndDate(endDate).Source(source).Dataset(dataset).Status(status).Statuses(statuses).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenIngestionStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenIngestionStatusList`: ConsoleV1IngestionControllerGenIngestionStatusList200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenIngestionStatusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenIngestionStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **source** | **string** |  | 
 **dataset** | **string** |  | 
 **status** | **string** |  | 
 **statuses** | **[]string** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenIngestionStatusList200Response**](ConsoleV1IngestionControllerGenIngestionStatusList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenReadIngestion

> ConsoleV1IngestionControllerGenReadIngestion200Response ConsoleV1IngestionControllerGenReadIngestion(ctx).Type_(type_).Dataset(dataset).SourceName(sourceName).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Ingestion

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
	type_ := "type__example" // string | 
	dataset := "dataset_example" // string | 
	sourceName := "sourceName_example" // string |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenReadIngestion(context.Background()).Type_(type_).Dataset(dataset).SourceName(sourceName).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenReadIngestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenReadIngestion`: ConsoleV1IngestionControllerGenReadIngestion200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenReadIngestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenReadIngestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **dataset** | **string** |  | 
 **sourceName** | **string** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenReadIngestion200Response**](ConsoleV1IngestionControllerGenReadIngestion200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenReadIngestionSchedule

> ConsoleV1IngestionControllerGenReadIngestionSchedule200Response ConsoleV1IngestionControllerGenReadIngestionSchedule(ctx).Dataset(dataset).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Ingestion Schedule

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
	dataset := "dataset_example" // string | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenReadIngestionSchedule(context.Background()).Dataset(dataset).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenReadIngestionSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenReadIngestionSchedule`: ConsoleV1IngestionControllerGenReadIngestionSchedule200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenReadIngestionSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenReadIngestionScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataset** | **string** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenReadIngestionSchedule200Response**](ConsoleV1IngestionControllerGenReadIngestionSchedule200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenUpdateIngestion

> ConsoleV1IngestionControllerGenUpdateIngestion200Response ConsoleV1IngestionControllerGenUpdateIngestion(ctx).IngestionUpdateContractDto(ingestionUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Ingestion Source

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
	ingestionUpdateContractDto := openapiclient.IngestionUpdateContractDto{IngestionUpdateContractDtoOneOf: openapiclient.NewIngestionUpdateContractDtoOneOf("Dataset_example", "Type_example")} // IngestionUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenUpdateIngestion(context.Background()).IngestionUpdateContractDto(ingestionUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenUpdateIngestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenUpdateIngestion`: ConsoleV1IngestionControllerGenUpdateIngestion200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenUpdateIngestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenUpdateIngestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestionUpdateContractDto** | [**IngestionUpdateContractDto**](IngestionUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenUpdateIngestion200Response**](ConsoleV1IngestionControllerGenUpdateIngestion200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionControllerGenUpdateIngestionSchedule

> ConsoleV1IngestionControllerGenUpdateIngestionSchedule200Response ConsoleV1IngestionControllerGenUpdateIngestionSchedule(ctx).IngestionScheduleUpdateContractDto(ingestionScheduleUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Ingestion Schedule

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
	ingestionScheduleUpdateContractDto := *openapiclient.NewIngestionScheduleUpdateContractDto("Dataset_example") // IngestionScheduleUpdateContractDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionControllerGenUpdateIngestionSchedule(context.Background()).IngestionScheduleUpdateContractDto(ingestionScheduleUpdateContractDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionControllerGenUpdateIngestionSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionControllerGenUpdateIngestionSchedule`: ConsoleV1IngestionControllerGenUpdateIngestionSchedule200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionControllerGenUpdateIngestionSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionControllerGenUpdateIngestionScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestionScheduleUpdateContractDto** | [**IngestionScheduleUpdateContractDto**](IngestionScheduleUpdateContractDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionControllerGenUpdateIngestionSchedule200Response**](ConsoleV1IngestionControllerGenUpdateIngestionSchedule200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger

> ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger200Response ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger(ctx).StartDate(startDate).EndDate(endDate).SourceName(sourceName).EventName(eventName).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Ingestion Event Delta Ledger

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
	startDate := "startDate_example" // string | 
	endDate := "endDate_example" // string | 
	sourceName := "sourceName_example" // string |  (optional)
	eventName := "eventName_example" // string |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger(context.Background()).StartDate(startDate).EndDate(endDate).SourceName(sourceName).EventName(eventName).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger`: ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionEventsControllerGenIngestionDeltaLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **sourceName** | **string** |  | 
 **eventName** | **string** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger200Response**](ConsoleV1IngestionEventsControllerGenIngestionDeltaLedger200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1IngestionEventsControllerGenIngestionEventCount

> ConsoleV1IngestionEventsControllerGenIngestionEventCount200Response ConsoleV1IngestionEventsControllerGenIngestionEventCount(ctx).StartDate(startDate).EndDate(endDate).SourceName(sourceName).EventName(eventName).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Ingestion Event Count

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
	startDate := "startDate_example" // string | 
	endDate := "endDate_example" // string | 
	sourceName := "sourceName_example" // string |  (optional)
	eventName := "eventName_example" // string |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionsAPI.ConsoleV1IngestionEventsControllerGenIngestionEventCount(context.Background()).StartDate(startDate).EndDate(endDate).SourceName(sourceName).EventName(eventName).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionsAPI.ConsoleV1IngestionEventsControllerGenIngestionEventCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1IngestionEventsControllerGenIngestionEventCount`: ConsoleV1IngestionEventsControllerGenIngestionEventCount200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionsAPI.ConsoleV1IngestionEventsControllerGenIngestionEventCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1IngestionEventsControllerGenIngestionEventCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **sourceName** | **string** |  | 
 **eventName** | **string** |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1IngestionEventsControllerGenIngestionEventCount200Response**](ConsoleV1IngestionEventsControllerGenIngestionEventCount200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

