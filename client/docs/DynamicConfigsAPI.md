# \DynamicConfigsAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1DynamicConfigControllerGenCreate**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenCreate) | **Post** /console/v1/dynamic_configs | Create Dynamic Config
[**ConsoleV1DynamicConfigControllerGenDisable**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenDisable) | **Put** /console/v1/dynamic_configs/{id}/disable | Disable Dynamic Config
[**ConsoleV1DynamicConfigControllerGenEnable**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenEnable) | **Put** /console/v1/dynamic_configs/{id}/enable | Enable Dynamic Config
[**ConsoleV1DynamicConfigControllerGenFullUpdate**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenFullUpdate) | **Post** /console/v1/dynamic_configs/{id} | Fully Update Dynamic Config
[**ConsoleV1DynamicConfigControllerGenList**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenList) | **Get** /console/v1/dynamic_configs | List Dynamic Configs
[**ConsoleV1DynamicConfigControllerGenListVersions**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenListVersions) | **Get** /console/v1/dynamic_configs/{id}/versions | List Dynamic Config Versions
[**ConsoleV1DynamicConfigControllerGenMultiRuleAdd**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenMultiRuleAdd) | **Post** /console/v1/dynamic_configs/{id}/rules | Add Dynamic Config Rules
[**ConsoleV1DynamicConfigControllerGenMultiRuleUpdate**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenMultiRuleUpdate) | **Patch** /console/v1/dynamic_configs/{id}/rules | Update List of Dynamic Config Rules
[**ConsoleV1DynamicConfigControllerGenPartialUpdate**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenPartialUpdate) | **Patch** /console/v1/dynamic_configs/{id} | Partially Update Dynamic Config
[**ConsoleV1DynamicConfigControllerGenRead**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenRead) | **Get** /console/v1/dynamic_configs/{id} | Get Dynamic Config
[**ConsoleV1DynamicConfigControllerGenReadRule**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenReadRule) | **Get** /console/v1/dynamic_configs/{id}/rule/{ruleId} | Get Specific Dynamic Config Rule
[**ConsoleV1DynamicConfigControllerGenReadRules**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenReadRules) | **Get** /console/v1/dynamic_configs/{id}/rules | Get Dynamic Config Rules
[**ConsoleV1DynamicConfigControllerGenRemove**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenRemove) | **Delete** /console/v1/dynamic_configs/{id} | Delete Dynamic Config
[**ConsoleV1DynamicConfigControllerGenRuleAdd**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenRuleAdd) | **Post** /console/v1/dynamic_configs/{id}/rule | Add Dynamic Config Rule
[**ConsoleV1DynamicConfigControllerGenRuleDelete**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenRuleDelete) | **Delete** /console/v1/dynamic_configs/{id}/rule/{ruleId} | Delete Dynamic Config Rule
[**ConsoleV1DynamicConfigControllerGenRuleNameDelete**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenRuleNameDelete) | **Delete** /console/v1/dynamic_configs/{id}/rule/name/{ruleName} | Delete Dynamic Config Rule By Name
[**ConsoleV1DynamicConfigControllerGenRuleNameUpdate**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenRuleNameUpdate) | **Patch** /console/v1/dynamic_configs/{id}/rule/name/{ruleName} | Update Dynamic Config Rule By Name
[**ConsoleV1DynamicConfigControllerGenRuleUpdate**](DynamicConfigsAPI.md#ConsoleV1DynamicConfigControllerGenRuleUpdate) | **Patch** /console/v1/dynamic_configs/{id}/rule/{ruleId} | Update Dynamic Config Rule By Id



## ConsoleV1DynamicConfigControllerGenCreate

> ConsoleV1DynamicConfigControllerGenCreate200Response ConsoleV1DynamicConfigControllerGenCreate(ctx).DynamicConfigCreateDto(dynamicConfigCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Dynamic Config

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
	dynamicConfigCreateDto := *openapiclient.NewDynamicConfigCreateDto("Name_example") // DynamicConfigCreateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenCreate(context.Background()).DynamicConfigCreateDto(dynamicConfigCreateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenCreate`: ConsoleV1DynamicConfigControllerGenCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicConfigCreateDto** | [**DynamicConfigCreateDto**](DynamicConfigCreateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenCreate200Response**](ConsoleV1DynamicConfigControllerGenCreate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenDisable

> ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response ConsoleV1DynamicConfigControllerGenDisable(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Disable Dynamic Config

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
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenDisable(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenDisable`: ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response**](ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenEnable

> ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response ConsoleV1DynamicConfigControllerGenEnable(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Enable Dynamic Config

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
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenEnable(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenEnable`: ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response**](ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenFullUpdate

> ConsoleV1DynamicConfigControllerGenFullUpdate200Response ConsoleV1DynamicConfigControllerGenFullUpdate(ctx, id).DynamicConfigFullUpdateDto(dynamicConfigFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Fully Update Dynamic Config

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
	dynamicConfigFullUpdateDto := *openapiclient.NewDynamicConfigFullUpdateDto(false, "helpful summary of what this dynamic config does", []openapiclient.DynamicConfigDtoRulesInner{*openapiclient.NewDynamicConfigDtoRulesInner("Name_example", float32(123), []openapiclient.DynamicConfigDtoRulesInnerConditionsInner{*openapiclient.NewDynamicConfigDtoRulesInnerConditionsInner("Type_example")})}) // DynamicConfigFullUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenFullUpdate(context.Background(), id).DynamicConfigFullUpdateDto(dynamicConfigFullUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenFullUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenFullUpdate`: ConsoleV1DynamicConfigControllerGenFullUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenFullUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenFullUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicConfigFullUpdateDto** | [**DynamicConfigFullUpdateDto**](DynamicConfigFullUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenFullUpdate200Response**](ConsoleV1DynamicConfigControllerGenFullUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenList

> ConsoleV1DynamicConfigControllerGenList200Response ConsoleV1DynamicConfigControllerGenList(ctx).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Dynamic Configs

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
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenList(context.Background()).CreatorName(creatorName).CreatorID(creatorID).Tags(tags).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenList`: ConsoleV1DynamicConfigControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creatorName** | **string** | Name of the creator. | 
 **creatorID** | **string** | ID of the user who created the entity. | 
 **tags** | **[]string** | Filter by tags | 
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenList200Response**](ConsoleV1DynamicConfigControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenListVersions

> ConsoleV1DynamicConfigControllerGenListVersions200Response ConsoleV1DynamicConfigControllerGenListVersions(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Dynamic Config Versions

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
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenListVersions(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenListVersions`: ConsoleV1DynamicConfigControllerGenListVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenListVersions200Response**](ConsoleV1DynamicConfigControllerGenListVersions200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenMultiRuleAdd

> ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response ConsoleV1DynamicConfigControllerGenMultiRuleAdd(ctx, id).MultiDynamicConfigRuleAddDto(multiDynamicConfigRuleAddDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Add Dynamic Config Rules

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
	multiDynamicConfigRuleAddDto := *openapiclient.NewMultiDynamicConfigRuleAddDto([]openapiclient.DynamicConfigDtoRulesInner{*openapiclient.NewDynamicConfigDtoRulesInner("Name_example", float32(123), []openapiclient.DynamicConfigDtoRulesInnerConditionsInner{*openapiclient.NewDynamicConfigDtoRulesInnerConditionsInner("Type_example")})}) // MultiDynamicConfigRuleAddDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenMultiRuleAdd(context.Background(), id).MultiDynamicConfigRuleAddDto(multiDynamicConfigRuleAddDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenMultiRuleAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenMultiRuleAdd`: ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenMultiRuleAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenMultiRuleAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multiDynamicConfigRuleAddDto** | [**MultiDynamicConfigRuleAddDto**](MultiDynamicConfigRuleAddDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response**](ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenMultiRuleUpdate

> ConsoleV1DynamicConfigControllerGenMultiRuleUpdate200Response ConsoleV1DynamicConfigControllerGenMultiRuleUpdate(ctx, id).MultiRuleUpdateDto(multiRuleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update List of Dynamic Config Rules

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
	multiRuleUpdateDto := *openapiclient.NewMultiRuleUpdateDto([]openapiclient.MultiRuleUpdateDtoRulesInner{*openapiclient.NewMultiRuleUpdateDtoRulesInner()}) // MultiRuleUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenMultiRuleUpdate(context.Background(), id).MultiRuleUpdateDto(multiRuleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenMultiRuleUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenMultiRuleUpdate`: ConsoleV1DynamicConfigControllerGenMultiRuleUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenMultiRuleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenMultiRuleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multiRuleUpdateDto** | [**MultiRuleUpdateDto**](MultiRuleUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenMultiRuleUpdate200Response**](ConsoleV1DynamicConfigControllerGenMultiRuleUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenPartialUpdate

> ConsoleV1DynamicConfigControllerGenPartialUpdate200Response ConsoleV1DynamicConfigControllerGenPartialUpdate(ctx, id).DynamicConfigPartialUpdateDto(dynamicConfigPartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Partially Update Dynamic Config

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
	dynamicConfigPartialUpdateDto := *openapiclient.NewDynamicConfigPartialUpdateDto() // DynamicConfigPartialUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenPartialUpdate(context.Background(), id).DynamicConfigPartialUpdateDto(dynamicConfigPartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenPartialUpdate`: ConsoleV1DynamicConfigControllerGenPartialUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicConfigPartialUpdateDto** | [**DynamicConfigPartialUpdateDto**](DynamicConfigPartialUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenPartialUpdate200Response**](ConsoleV1DynamicConfigControllerGenPartialUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenRead

> ConsoleV1DynamicConfigControllerGenRead200Response ConsoleV1DynamicConfigControllerGenRead(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Dynamic Config

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
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRead(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenRead`: ConsoleV1DynamicConfigControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenRead200Response**](ConsoleV1DynamicConfigControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenReadRule

> ConsoleV1DynamicConfigControllerGenReadRule200Response ConsoleV1DynamicConfigControllerGenReadRule(ctx, id, ruleId).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Specific Dynamic Config Rule

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
	id := "id_example" // string | Dynamic Config ID
	ruleId := "ruleId_example" // string | Rule ID
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenReadRule(context.Background(), id, ruleId).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenReadRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenReadRule`: ConsoleV1DynamicConfigControllerGenReadRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenReadRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Config ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenReadRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenReadRule200Response**](ConsoleV1DynamicConfigControllerGenReadRule200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenReadRules

> ConsoleV1DynamicConfigControllerGenReadRules200Response ConsoleV1DynamicConfigControllerGenReadRules(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Dynamic Config Rules

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
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenReadRules(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenReadRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenReadRules`: ConsoleV1DynamicConfigControllerGenReadRules200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenReadRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenReadRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenReadRules200Response**](ConsoleV1DynamicConfigControllerGenReadRules200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenRemove

> ConsoleV1DynamicConfigControllerGenRemove200Response ConsoleV1DynamicConfigControllerGenRemove(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Dynamic Config

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
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRemove(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenRemove`: ConsoleV1DynamicConfigControllerGenRemove200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenRemove200Response**](ConsoleV1DynamicConfigControllerGenRemove200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenRuleAdd

> ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response ConsoleV1DynamicConfigControllerGenRuleAdd(ctx, id).DynamicConfigRuleAddDto(dynamicConfigRuleAddDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Add Dynamic Config Rule

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
	dynamicConfigRuleAddDto := *openapiclient.NewDynamicConfigRuleAddDto("Name_example", float32(123), []openapiclient.DynamicConfigDtoRulesInnerConditionsInner{*openapiclient.NewDynamicConfigDtoRulesInnerConditionsInner("Type_example")}) // DynamicConfigRuleAddDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleAdd(context.Background(), id).DynamicConfigRuleAddDto(dynamicConfigRuleAddDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenRuleAdd`: ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenRuleAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicConfigRuleAddDto** | [**DynamicConfigRuleAddDto**](DynamicConfigRuleAddDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response**](ConsoleV1DynamicConfigControllerGenMultiRuleAdd200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenRuleDelete

> ConsoleV1DynamicConfigControllerGenRuleDelete200Response ConsoleV1DynamicConfigControllerGenRuleDelete(ctx, id, ruleId).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Dynamic Config Rule

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
	id := "id_example" // string | Dynamic Config ID
	ruleId := "ruleId_example" // string | Rule ID
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleDelete(context.Background(), id, ruleId).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenRuleDelete`: ConsoleV1DynamicConfigControllerGenRuleDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Config ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenRuleDelete200Response**](ConsoleV1DynamicConfigControllerGenRuleDelete200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenRuleNameDelete

> ConsoleV1DynamicConfigControllerGenRuleDelete200Response ConsoleV1DynamicConfigControllerGenRuleNameDelete(ctx, id, ruleName).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Dynamic Config Rule By Name

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
	id := "id_example" // string | Dynamic Config ID
	ruleName := "ruleName_example" // string | Rule Name
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleNameDelete(context.Background(), id, ruleName).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenRuleNameDelete`: ConsoleV1DynamicConfigControllerGenRuleDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Config ID | 
**ruleName** | **string** | Rule Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenRuleNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenRuleDelete200Response**](ConsoleV1DynamicConfigControllerGenRuleDelete200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenRuleNameUpdate

> ConsoleV1DynamicConfigControllerGenRuleUpdate200Response ConsoleV1DynamicConfigControllerGenRuleNameUpdate(ctx, id, ruleName).RuleUpdateDto(ruleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Dynamic Config Rule By Name

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
	id := "id_example" // string | Dynamic Config ID
	ruleName := "ruleName_example" // string | Rule Name
	ruleUpdateDto := *openapiclient.NewRuleUpdateDto() // RuleUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleNameUpdate(context.Background(), id, ruleName).RuleUpdateDto(ruleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleNameUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenRuleNameUpdate`: ConsoleV1DynamicConfigControllerGenRuleUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleNameUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Config ID | 
**ruleName** | **string** | Rule Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenRuleNameUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleUpdateDto** | [**RuleUpdateDto**](RuleUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenRuleUpdate200Response**](ConsoleV1DynamicConfigControllerGenRuleUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1DynamicConfigControllerGenRuleUpdate

> ConsoleV1DynamicConfigControllerGenRuleUpdate200Response ConsoleV1DynamicConfigControllerGenRuleUpdate(ctx, id, ruleId).RuleUpdateDto(ruleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Dynamic Config Rule By Id

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
	id := "id_example" // string | Dynamic Config ID
	ruleId := "ruleId_example" // string | Rule ID
	ruleUpdateDto := *openapiclient.NewRuleUpdateDto() // RuleUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleUpdate(context.Background(), id, ruleId).RuleUpdateDto(ruleUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1DynamicConfigControllerGenRuleUpdate`: ConsoleV1DynamicConfigControllerGenRuleUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `DynamicConfigsAPI.ConsoleV1DynamicConfigControllerGenRuleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dynamic Config ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1DynamicConfigControllerGenRuleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleUpdateDto** | [**RuleUpdateDto**](RuleUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1DynamicConfigControllerGenRuleUpdate200Response**](ConsoleV1DynamicConfigControllerGenRuleUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

