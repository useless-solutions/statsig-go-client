# \AlertsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1AlertsControllerGenList**](AlertsAPI.md#ConsoleV1AlertsControllerGenList) | **Get** /console/v1/alerts | List Alerts



## ConsoleV1AlertsControllerGenList

> ConsoleV1AlertsControllerGenList200Response ConsoleV1AlertsControllerGenList(ctx).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Alerts

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
	resp, r, err := apiClient.AlertsAPI.ConsoleV1AlertsControllerGenList(context.Background()).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ConsoleV1AlertsControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1AlertsControllerGenList`: ConsoleV1AlertsControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ConsoleV1AlertsControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1AlertsControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creatorName** | **string** | Name of the creator. | 
 **creatorID** | **string** | ID of the user who created the entity. | 
 **tags** | **[]string** | Filter by tags | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1AlertsControllerGenList200Response**](ConsoleV1AlertsControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

