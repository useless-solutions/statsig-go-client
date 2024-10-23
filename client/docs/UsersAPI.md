# \UsersAPI

All URIs are relative to *https://statsigapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleV1UsersControllerGenCreateTeam**](UsersAPI.md#ConsoleV1UsersControllerGenCreateTeam) | **Post** /console/v1/users/teams | Create Team
[**ConsoleV1UsersControllerGenDeleteTeam**](UsersAPI.md#ConsoleV1UsersControllerGenDeleteTeam) | **Delete** /console/v1/users/teams/{id} | Delete Team
[**ConsoleV1UsersControllerGenList**](UsersAPI.md#ConsoleV1UsersControllerGenList) | **Get** /console/v1/users | List Users
[**ConsoleV1UsersControllerGenListTeams**](UsersAPI.md#ConsoleV1UsersControllerGenListTeams) | **Get** /console/v1/users/teams | List Teams
[**ConsoleV1UsersControllerGenRead**](UsersAPI.md#ConsoleV1UsersControllerGenRead) | **Get** /console/v1/users/{email} | Get user by email
[**ConsoleV1UsersControllerGenReadTeam**](UsersAPI.md#ConsoleV1UsersControllerGenReadTeam) | **Get** /console/v1/users/teams/{id} | Get Team
[**ConsoleV1UsersControllerGenUpdateTeam**](UsersAPI.md#ConsoleV1UsersControllerGenUpdateTeam) | **Patch** /console/v1/users/teams/{id} | Update Team. Ops: Replace. Use GET for current data if you intent to Add.
[**ConsoleV1UsersControllerInviteUsers**](UsersAPI.md#ConsoleV1UsersControllerInviteUsers) | **Post** /console/v1/users/invite | Invite user. To avoid spamming, invitation emails are not sent. Invitee will see invitation notification in-app after logging in.
[**ConsoleV1UsersControllerUpdate**](UsersAPI.md#ConsoleV1UsersControllerUpdate) | **Post** /console/v1/users/{id} | Update user



## ConsoleV1UsersControllerGenCreateTeam

> ConsoleV1UsersControllerGenCreateTeam200Response ConsoleV1UsersControllerGenCreateTeam(ctx).TeamCreationDto(teamCreationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Create Team

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
	teamCreationDto := *openapiclient.NewTeamCreationDto("Name_example", []string{"Members_example"}, []string{"Admins_example"}, []openapiclient.TeamDtoDefaultGateMetricsInner{*openapiclient.NewTeamDtoDefaultGateMetricsInner("Name_example", "Type_example")}, []openapiclient.TeamDtoDefaultGateMetricsInner{*openapiclient.NewTeamDtoDefaultGateMetricsInner("Name_example", "Type_example")}, []openapiclient.TeamDtoDefaultGateMetricsInner{*openapiclient.NewTeamDtoDefaultGateMetricsInner("Name_example", "Type_example")}, []openapiclient.TeamDtoDefaultGateMetricsInner{*openapiclient.NewTeamDtoDefaultGateMetricsInner("Name_example", "Type_example")}, "ChangeTeamConfigs_example", "ReviewApproval_example", []string{"DefaultTargetApplications_example"}) // TeamCreationDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerGenCreateTeam(context.Background()).TeamCreationDto(teamCreationDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerGenCreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerGenCreateTeam`: ConsoleV1UsersControllerGenCreateTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerGenCreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerGenCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamCreationDto** | [**TeamCreationDto**](TeamCreationDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerGenCreateTeam200Response**](ConsoleV1UsersControllerGenCreateTeam200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerGenDeleteTeam

> ConsoleV1UsersControllerGenDeleteTeam200Response ConsoleV1UsersControllerGenDeleteTeam(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Delete Team

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
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerGenDeleteTeam(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerGenDeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerGenDeleteTeam`: ConsoleV1UsersControllerGenDeleteTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerGenDeleteTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerGenDeleteTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerGenDeleteTeam200Response**](ConsoleV1UsersControllerGenDeleteTeam200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerGenList

> ConsoleV1UsersControllerGenList200Response ConsoleV1UsersControllerGenList(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Users

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
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerGenList(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerGenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerGenList`: ConsoleV1UsersControllerGenList200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerGenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerGenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerGenList200Response**](ConsoleV1UsersControllerGenList200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerGenListTeams

> ConsoleV1UsersControllerGenListTeams200Response ConsoleV1UsersControllerGenListTeams(ctx).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()

List Teams

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
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerGenListTeams(context.Background()).Limit(limit).Page(page).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerGenListTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerGenListTeams`: ConsoleV1UsersControllerGenListTeams200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerGenListTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerGenListTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Results per page | 
 **page** | **int32** | Page number | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerGenListTeams200Response**](ConsoleV1UsersControllerGenListTeams200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerGenRead

> ConsoleV1UsersControllerGenRead200Response ConsoleV1UsersControllerGenRead(ctx, email).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get user by email

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
	email := "email_example" // string | email
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerGenRead(context.Background(), email).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerGenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerGenRead`: ConsoleV1UsersControllerGenRead200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerGenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | email | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerGenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerGenRead200Response**](ConsoleV1UsersControllerGenRead200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerGenReadTeam

> ConsoleV1UsersControllerGenReadTeam200Response ConsoleV1UsersControllerGenReadTeam(ctx, id).XRespectReviewSettings(xRespectReviewSettings).Execute()

Get Team

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
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerGenReadTeam(context.Background(), id).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerGenReadTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerGenReadTeam`: ConsoleV1UsersControllerGenReadTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerGenReadTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerGenReadTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerGenReadTeam200Response**](ConsoleV1UsersControllerGenReadTeam200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerGenUpdateTeam

> ConsoleV1UsersControllerGenUpdateTeam200Response ConsoleV1UsersControllerGenUpdateTeam(ctx, id).TeamPartialUpdateDto(teamPartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update Team. Ops: Replace. Use GET for current data if you intent to Add.

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
	teamPartialUpdateDto := *openapiclient.NewTeamPartialUpdateDto() // TeamPartialUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerGenUpdateTeam(context.Background(), id).TeamPartialUpdateDto(teamPartialUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerGenUpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerGenUpdateTeam`: ConsoleV1UsersControllerGenUpdateTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerGenUpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerGenUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamPartialUpdateDto** | [**TeamPartialUpdateDto**](TeamPartialUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerGenUpdateTeam200Response**](ConsoleV1UsersControllerGenUpdateTeam200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerInviteUsers

> ConsoleV1UsersControllerInviteUsers200Response ConsoleV1UsersControllerInviteUsers(ctx).UserInvitesDto(userInvitesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Invite user. To avoid spamming, invitation emails are not sent. Invitee will see invitation notification in-app after logging in.

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
	userInvitesDto := *openapiclient.NewUserInvitesDto("Role_example", []string{"Emails_example"}) // UserInvitesDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerInviteUsers(context.Background()).UserInvitesDto(userInvitesDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerInviteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerInviteUsers`: ConsoleV1UsersControllerInviteUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerInviteUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerInviteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userInvitesDto** | [**UserInvitesDto**](UserInvitesDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerInviteUsers200Response**](ConsoleV1UsersControllerInviteUsers200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleV1UsersControllerUpdate

> ConsoleV1UsersControllerUpdate200Response ConsoleV1UsersControllerUpdate(ctx, id).UserUpdateDto(userUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()

Update user

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
	userUpdateDto := *openapiclient.NewUserUpdateDto() // UserUpdateDto | 
	xRespectReviewSettings := "xRespectReviewSettings_example" // string | Optional header to respect review settings for mutation endpoints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ConsoleV1UsersControllerUpdate(context.Background(), id).UserUpdateDto(userUpdateDto).XRespectReviewSettings(xRespectReviewSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConsoleV1UsersControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleV1UsersControllerUpdate`: ConsoleV1UsersControllerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ConsoleV1UsersControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleV1UsersControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateDto** | [**UserUpdateDto**](UserUpdateDto.md) |  | 
 **xRespectReviewSettings** | **string** | Optional header to respect review settings for mutation endpoints. | 

### Return type

[**ConsoleV1UsersControllerUpdate200Response**](ConsoleV1UsersControllerUpdate200Response.md)

### Authorization

[STATSIG-API-KEY](../README.md#STATSIG-API-KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

