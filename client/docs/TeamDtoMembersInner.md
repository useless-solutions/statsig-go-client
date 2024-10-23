# TeamDtoMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the user. | 
**FirstName** | **string** | The user&#39;s first name. | 
**LastName** | **string** | The user&#39;s last name. | 
**Role** | **string** | The user&#39;s role, which can be &#39;Admin&#39;, &#39;Read Only&#39;, &#39;Member&#39;, or any custom role name. | 

## Methods

### NewTeamDtoMembersInner

`func NewTeamDtoMembersInner(email string, firstName string, lastName string, role string, ) *TeamDtoMembersInner`

NewTeamDtoMembersInner instantiates a new TeamDtoMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamDtoMembersInnerWithDefaults

`func NewTeamDtoMembersInnerWithDefaults() *TeamDtoMembersInner`

NewTeamDtoMembersInnerWithDefaults instantiates a new TeamDtoMembersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TeamDtoMembersInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TeamDtoMembersInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TeamDtoMembersInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *TeamDtoMembersInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TeamDtoMembersInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TeamDtoMembersInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *TeamDtoMembersInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TeamDtoMembersInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TeamDtoMembersInner) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetRole

`func (o *TeamDtoMembersInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TeamDtoMembersInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TeamDtoMembersInner) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


