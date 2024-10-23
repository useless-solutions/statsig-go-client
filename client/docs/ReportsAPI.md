# \ReportsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1ReportsControllerGenReport**](ReportsAPI.md#ConsoleV1ReportsControllerGenReport) | **Get** /console/v1/reports | Get Reports



## ConsoleV1ReportsControllerGenReport

> ConsoleV1ReportsControllerGenReport200Response ConsoleV1ReportsControllerGenReport(ctx).Type_(type_).Date(date).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Reports

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
	type_ := "first_exposures" // string | report type
	date := "2024-09-01" // string | date for the report
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ConsoleV1ReportsControllerGenReport(context.Background()).Type_(type_).Date(date).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ConsoleV1ReportsControllerGenReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ReportsControllerGenReport`: ConsoleV1ReportsControllerGenReport200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ConsoleV1ReportsControllerGenReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ReportsControllerGenReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | report type | 
 **date** | **string** | date for the report | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ReportsControllerGenReport200Response**](ConsoleV1ReportsControllerGenReport200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

