# UserInvitesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | Role assigned to the invited users. Can be &#39;Admin&#39;, &#39;Read Only&#39;, &#39;Member&#39;, or any custom role name. | 
**Emails** | **[]string** | List of email addresses to send invitations to. Invitee Emails must have the same domain to your company email domain. | 
**Teams** | Pointer to **[]string** | Optional list of teams that the invited users will be associated with. | [optional] 

## Methods

### NewUserInvitesDto

`func NewUserInvitesDto(role string, emails []string, ) *UserInvitesDto`

NewUserInvitesDto instantiates a new UserInvitesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitesDtoWithDefaults

`func NewUserInvitesDtoWithDefaults() *UserInvitesDto`

NewUserInvitesDtoWithDefaults instantiates a new UserInvitesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *UserInvitesDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserInvitesDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserInvitesDto) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmails

`func (o *UserInvitesDto) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UserInvitesDto) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UserInvitesDto) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetTeams

`func (o *UserInvitesDto) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *UserInvitesDto) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *UserInvitesDto) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *UserInvitesDto) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


