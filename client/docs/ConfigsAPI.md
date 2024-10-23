# \ConfigsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1ChangeValidationControllerChangeValidation**](ConfigsAPI.md#ConsoleV1ChangeValidationControllerChangeValidation) | **Post** /console/v1/change_validation | Change Validation
[**ConsoleV1ExposureCountsControllerGenExposureCount**](ConfigsAPI.md#ConsoleV1ExposureCountsControllerGenExposureCount) | **Get** /console/v1/exposure_count | Read Exposure Event Count



## ConsoleV1ChangeValidationControllerChangeValidation

> ConsoleV1ChangeValidationControllerChangeValidation200Response ConsoleV1ChangeValidationControllerChangeValidation(ctx).ChangeValidationDto(changeValidationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Change Validation

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
	changeValidationDto := *openapiclient.NewChangeValidationDto("Name_example", false) // ChangeValidationDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.ConsoleV1ChangeValidationControllerChangeValidation(context.Background()).ChangeValidationDto(changeValidationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.ConsoleV1ChangeValidationControllerChangeValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ChangeValidationControllerChangeValidation`: ConsoleV1ChangeValidationControllerChangeValidation200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.ConsoleV1ChangeValidationControllerChangeValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ChangeValidationControllerChangeValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeValidationDto** | [**ChangeValidationDto**](ChangeValidationDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ChangeValidationControllerChangeValidation200Response**](ConsoleV1ChangeValidationControllerChangeValidation200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1ExposureCountsControllerGenExposureCount

> ConsoleV1ExposureCountsControllerGenExposureCount200Response ConsoleV1ExposureCountsControllerGenExposureCount(ctx).Experiments(experiments).Gates(gates).XRespectReviewSettings(xRespectReviewSettings).Execute()

Read Exposure Event Count



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
	experiments := TODO // interface{} |  (optional)
	gates := TODO // interface{} |  (optional)
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.ConsoleV1ExposureCountsControllerGenExposureCount(context.Background()).Experiments(experiments).Gates(gates).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.ConsoleV1ExposureCountsControllerGenExposureCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1ExposureCountsControllerGenExposureCount`: ConsoleV1ExposureCountsControllerGenExposureCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.ConsoleV1ExposureCountsControllerGenExposureCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1ExposureCountsControllerGenExposureCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **experiments** | [**interface{}**](interface{}.md) |  | 
 **gates** | [**interface{}**](interface{}.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1ExposureCountsControllerGenExposureCount200Response**](ConsoleV1ExposureCountsControllerGenExposureCount200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

