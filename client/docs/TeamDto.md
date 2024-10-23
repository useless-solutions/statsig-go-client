# TeamDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the team. | 
**Description** | Pointer to **string** | Description of the team. | [optional] 
**Id** | **string** | The ID of the team. | 
**DefaultGateMetrics** | [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default gate metrics for the team. | 
**DefaultExperimentPrimaryMetrics** | [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default primary metrics for experiments in the team. | 
**DefaultExperimentSecondaryMetrics** | [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default secondary metrics for experiments in the team. | 
**DefaultHoldoutMetrics** | [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default holdout metrics for the team. | 
**ChangeTeamConfigs** | **string** | Who can change team configurations: \&quot;anyone\&quot; or \&quot;team_only\&quot;. | 
**ReviewApproval** | **string** | Who can review and approve changes: \&quot;anyone\&quot;, \&quot;team_only\&quot;, or \&quot;admin_only\&quot;. | 
**DefaultTargetApplications** | **[]string** | Default target applications for the team. | 
**DefaultHoldoutID** | Pointer to [**nil**](nil.md) | Default holdout ID for the team, if applicable. | [optional] 
**RequireReviews** | Pointer to [**nil**](nil.md) | Whether reviews are required for changes, if applicable. | [optional] 
**RequireGateTemplates** | Pointer to [**nil**](nil.md) | Whether gate templates are required for the team, if applicable. | [optional] 
**RequireExperimentTemplates** | Pointer to [**nil**](nil.md) | Whether experiment templates are required for the team, if applicable. | [optional] 
**Members** | [**[]TeamDtoMembersInner**](TeamDtoMembersInner.md) |  | 
**Admins** | [**[]TeamDtoMembersInner**](TeamDtoMembersInner.md) |  | 

## Methods

### NewTeamDto

`func NewTeamDto(name string, id string, defaultGateMetrics []TeamDtoDefaultGateMetricsInner, defaultExperimentPrimaryMetrics []TeamDtoDefaultGateMetricsInner, defaultExperimentSecondaryMetrics []TeamDtoDefaultGateMetricsInner, defaultHoldoutMetrics []TeamDtoDefaultGateMetricsInner, changeTeamConfigs string, reviewApproval string, defaultTargetApplications []string, members []TeamDtoMembersInner, admins []TeamDtoMembersInner, ) *TeamDto`

NewTeamDto instantiates a new TeamDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamDtoWithDefaults

`func NewTeamDtoWithDefaults() *TeamDto`

NewTeamDtoWithDefaults instantiates a new TeamDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TeamDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamDto) SetId(v string)`

SetId sets Id field to given value.


### GetDefaultGateMetrics

`func (o *TeamDto) GetDefaultGateMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultGateMetrics returns the DefaultGateMetrics field if non-nil, zero value otherwise.

### GetDefaultGateMetricsOk

`func (o *TeamDto) GetDefaultGateMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultGateMetricsOk returns a tuple with the DefaultGateMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateMetrics

`func (o *TeamDto) SetDefaultGateMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultGateMetrics sets DefaultGateMetrics field to given value.


### GetDefaultExperimentPrimaryMetrics

`func (o *TeamDto) GetDefaultExperimentPrimaryMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultExperimentPrimaryMetrics returns the DefaultExperimentPrimaryMetrics field if non-nil, zero value otherwise.

### GetDefaultExperimentPrimaryMetricsOk

`func (o *TeamDto) GetDefaultExperimentPrimaryMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultExperimentPrimaryMetricsOk returns a tuple with the DefaultExperimentPrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentPrimaryMetrics

`func (o *TeamDto) SetDefaultExperimentPrimaryMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultExperimentPrimaryMetrics sets DefaultExperimentPrimaryMetrics field to given value.


### GetDefaultExperimentSecondaryMetrics

`func (o *TeamDto) GetDefaultExperimentSecondaryMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultExperimentSecondaryMetrics returns the DefaultExperimentSecondaryMetrics field if non-nil, zero value otherwise.

### GetDefaultExperimentSecondaryMetricsOk

`func (o *TeamDto) GetDefaultExperimentSecondaryMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultExperimentSecondaryMetricsOk returns a tuple with the DefaultExperimentSecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentSecondaryMetrics

`func (o *TeamDto) SetDefaultExperimentSecondaryMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultExperimentSecondaryMetrics sets DefaultExperimentSecondaryMetrics field to given value.


### GetDefaultHoldoutMetrics

`func (o *TeamDto) GetDefaultHoldoutMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultHoldoutMetrics returns the DefaultHoldoutMetrics field if non-nil, zero value otherwise.

### GetDefaultHoldoutMetricsOk

`func (o *TeamDto) GetDefaultHoldoutMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultHoldoutMetricsOk returns a tuple with the DefaultHoldoutMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHoldoutMetrics

`func (o *TeamDto) SetDefaultHoldoutMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultHoldoutMetrics sets DefaultHoldoutMetrics field to given value.


### GetChangeTeamConfigs

`func (o *TeamDto) GetChangeTeamConfigs() string`

GetChangeTeamConfigs returns the ChangeTeamConfigs field if non-nil, zero value otherwise.

### GetChangeTeamConfigsOk

`func (o *TeamDto) GetChangeTeamConfigsOk() (*string, bool)`

GetChangeTeamConfigsOk returns a tuple with the ChangeTeamConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTeamConfigs

`func (o *TeamDto) SetChangeTeamConfigs(v string)`

SetChangeTeamConfigs sets ChangeTeamConfigs field to given value.


### GetReviewApproval

`func (o *TeamDto) GetReviewApproval() string`

GetReviewApproval returns the ReviewApproval field if non-nil, zero value otherwise.

### GetReviewApprovalOk

`func (o *TeamDto) GetReviewApprovalOk() (*string, bool)`

GetReviewApprovalOk returns a tuple with the ReviewApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewApproval

`func (o *TeamDto) SetReviewApproval(v string)`

SetReviewApproval sets ReviewApproval field to given value.


### GetDefaultTargetApplications

`func (o *TeamDto) GetDefaultTargetApplications() []string`

GetDefaultTargetApplications returns the DefaultTargetApplications field if non-nil, zero value otherwise.

### GetDefaultTargetApplicationsOk

`func (o *TeamDto) GetDefaultTargetApplicationsOk() (*[]string, bool)`

GetDefaultTargetApplicationsOk returns a tuple with the DefaultTargetApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetApplications

`func (o *TeamDto) SetDefaultTargetApplications(v []string)`

SetDefaultTargetApplications sets DefaultTargetApplications field to given value.


### GetDefaultHoldoutID

`func (o *TeamDto) GetDefaultHoldoutID() nil`

GetDefaultHoldoutID returns the DefaultHoldoutID field if non-nil, zero value otherwise.

### GetDefaultHoldoutIDOk

`func (o *TeamDto) GetDefaultHoldoutIDOk() (*nil, bool)`

GetDefaultHoldoutIDOk returns a tuple with the DefaultHoldoutID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHoldoutID

`func (o *TeamDto) SetDefaultHoldoutID(v nil)`

SetDefaultHoldoutID sets DefaultHoldoutID field to given value.

### HasDefaultHoldoutID

`func (o *TeamDto) HasDefaultHoldoutID() bool`

HasDefaultHoldoutID returns a boolean if a field has been set.

### GetRequireReviews

`func (o *TeamDto) GetRequireReviews() nil`

GetRequireReviews returns the RequireReviews field if non-nil, zero value otherwise.

### GetRequireReviewsOk

`func (o *TeamDto) GetRequireReviewsOk() (*nil, bool)`

GetRequireReviewsOk returns a tuple with the RequireReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReviews

`func (o *TeamDto) SetRequireReviews(v nil)`

SetRequireReviews sets RequireReviews field to given value.

### HasRequireReviews

`func (o *TeamDto) HasRequireReviews() bool`

HasRequireReviews returns a boolean if a field has been set.

### GetRequireGateTemplates

`func (o *TeamDto) GetRequireGateTemplates() nil`

GetRequireGateTemplates returns the RequireGateTemplates field if non-nil, zero value otherwise.

### GetRequireGateTemplatesOk

`func (o *TeamDto) GetRequireGateTemplatesOk() (*nil, bool)`

GetRequireGateTemplatesOk returns a tuple with the RequireGateTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGateTemplates

`func (o *TeamDto) SetRequireGateTemplates(v nil)`

SetRequireGateTemplates sets RequireGateTemplates field to given value.

### HasRequireGateTemplates

`func (o *TeamDto) HasRequireGateTemplates() bool`

HasRequireGateTemplates returns a boolean if a field has been set.

### GetRequireExperimentTemplates

`func (o *TeamDto) GetRequireExperimentTemplates() nil`

GetRequireExperimentTemplates returns the RequireExperimentTemplates field if non-nil, zero value otherwise.

### GetRequireExperimentTemplatesOk

`func (o *TeamDto) GetRequireExperimentTemplatesOk() (*nil, bool)`

GetRequireExperimentTemplatesOk returns a tuple with the RequireExperimentTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExperimentTemplates

`func (o *TeamDto) SetRequireExperimentTemplates(v nil)`

SetRequireExperimentTemplates sets RequireExperimentTemplates field to given value.

### HasRequireExperimentTemplates

`func (o *TeamDto) HasRequireExperimentTemplates() bool`

HasRequireExperimentTemplates returns a boolean if a field has been set.

### GetMembers

`func (o *TeamDto) GetMembers() []TeamDtoMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamDto) GetMembersOk() (*[]TeamDtoMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamDto) SetMembers(v []TeamDtoMembersInner)`

SetMembers sets Members field to given value.


### GetAdmins

`func (o *TeamDto) GetAdmins() []TeamDtoMembersInner`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *TeamDto) GetAdminsOk() (*[]TeamDtoMembersInner, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *TeamDto) SetAdmins(v []TeamDtoMembersInner)`

SetAdmins sets Admins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


