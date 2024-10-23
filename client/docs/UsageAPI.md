# \UsageAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1UsageControllerGenReport**](UsageAPI.md#ConsoleV1UsageControllerGenReport) | **Get** /console/v1/project/usage_billing/report | Get Report URL



## ConsoleV1UsageControllerGenReport

> ConsoleV1UsageControllerGenReport200Response ConsoleV1UsageControllerGenReport(ctx).End(end).Start(start).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Report URL

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
	end := int32(56) // int32 | date to query (YYYY-MM-DD)
	start := int32(56) // int32 | date to query (YYYY-MM-DD) (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.ConsoleV1UsageControllerGenReport(context.Background()).End(end).Start(start).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.ConsoleV1UsageControllerGenReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsageControllerGenReport`: ConsoleV1UsageControllerGenReport200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.ConsoleV1UsageControllerGenReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsageControllerGenReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **end** | **int32** | date to query (YYYY-MM-DD) | 
 **start** | **int32** | date to query (YYYY-MM-DD) | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsageControllerGenReport200Response**](ConsoleV1UsageControllerGenReport200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

