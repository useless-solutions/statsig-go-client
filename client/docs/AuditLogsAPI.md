# \AuditLogsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1AuditLogsControllerGenList**](AuditLogsAPI.md#ConsoleV1AuditLogsControllerGenList) | **Get** /console/v1/audit_logs | List Audit Logs



## ConsoleV1AuditLogsControllerGenList

> ConsoleV1AuditLogsControllerGenList200Response ConsoleV1AuditLogsControllerGenList(ctx).Id(id).SortKey(sortKey).SortOrder(sortOrder).LatestID(latestID).Tags(tags).ActionType(actionType).ActionTypes(actionTypes).StartDate(startDate).EndDate(endDate).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Audit Logs

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
	id := "id_example" // string |  (optional)
	sortKey := "sortKey_example" // string |  (optional)
	sortOrder := "sortOrder_example" // string |  (optional)
	latestID := "latestID_example" // string |  (optional)
	tags := []string{"Inner_example"} // []string |  (optional)
	actionType := "actionType_example" // string |  (optional)
	actionTypes := []string{"Inner_example"} // []string |  (optional)
	startDate := "startDate_example" // string |  (optional)
	endDate := "endDate_example" // string |  (optional)
	limit := int32(10) // int32 | Results per page (optional)
	page := int32(1) // int32 | Page number (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.ConsoleV1AuditLogsControllerGenList(context.Background()).Id(id).SortKey(sortKey).SortOrder(sortOrder).LatestID(latestID).Tags(tags).ActionType(actionType).ActionTypes(actionTypes).StartDate(startDate).EndDate(endDate).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.ConsoleV1AuditLogsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AuditLogsControllerGenList`: ConsoleV1AuditLogsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.ConsoleV1AuditLogsControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AuditLogsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **sortKey** | **string** |  | 
 **sortOrder** | **string** |  | 
 **latestID** | **string** |  | 
 **tags** | **[]string** |  | 
 **actionType** | **string** |  | 
 **actionTypes** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AuditLogsControllerGenList200Response**](ConsoleV1AuditLogsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

