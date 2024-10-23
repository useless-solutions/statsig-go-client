# TeamCreationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the team. | 
**Description** | Pointer to **string** | Description of the team. | [optional] 
**Members** | **[]string** | Array of member email addresses in the team. | 
**Admins** | **[]string** | Array of admin email addresses in the team. | 
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

## Methods

### NewTeamCreationDto

`func NewTeamCreationDto(name string, members []string, admins []string, defaultGateMetrics []TeamDtoDefaultGateMetricsInner, defaultExperimentPrimaryMetrics []TeamDtoDefaultGateMetricsInner, defaultExperimentSecondaryMetrics []TeamDtoDefaultGateMetricsInner, defaultHoldoutMetrics []TeamDtoDefaultGateMetricsInner, changeTeamConfigs string, reviewApproval string, defaultTargetApplications []string, ) *TeamCreationDto`

NewTeamCreationDto instantiates a new TeamCreationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCreationDtoWithDefaults

`func NewTeamCreationDtoWithDefaults() *TeamCreationDto`

NewTeamCreationDtoWithDefaults instantiates a new TeamCreationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamCreationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamCreationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamCreationDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamCreationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamCreationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamCreationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamCreationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMembers

`func (o *TeamCreationDto) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamCreationDto) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamCreationDto) SetMembers(v []string)`

SetMembers sets Members field to given value.


### GetAdmins

`func (o *TeamCreationDto) GetAdmins() []string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *TeamCreationDto) GetAdminsOk() (*[]string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *TeamCreationDto) SetAdmins(v []string)`

SetAdmins sets Admins field to given value.


### GetDefaultGateMetrics

`func (o *TeamCreationDto) GetDefaultGateMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultGateMetrics returns the DefaultGateMetrics field if non-nil, zero value otherwise.

### GetDefaultGateMetricsOk

`func (o *TeamCreationDto) GetDefaultGateMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultGateMetricsOk returns a tuple with the DefaultGateMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateMetrics

`func (o *TeamCreationDto) SetDefaultGateMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultGateMetrics sets DefaultGateMetrics field to given value.


### GetDefaultExperimentPrimaryMetrics

`func (o *TeamCreationDto) GetDefaultExperimentPrimaryMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultExperimentPrimaryMetrics returns the DefaultExperimentPrimaryMetrics field if non-nil, zero value otherwise.

### GetDefaultExperimentPrimaryMetricsOk

`func (o *TeamCreationDto) GetDefaultExperimentPrimaryMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultExperimentPrimaryMetricsOk returns a tuple with the DefaultExperimentPrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentPrimaryMetrics

`func (o *TeamCreationDto) SetDefaultExperimentPrimaryMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultExperimentPrimaryMetrics sets DefaultExperimentPrimaryMetrics field to given value.


### GetDefaultExperimentSecondaryMetrics

`func (o *TeamCreationDto) GetDefaultExperimentSecondaryMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultExperimentSecondaryMetrics returns the DefaultExperimentSecondaryMetrics field if non-nil, zero value otherwise.

### GetDefaultExperimentSecondaryMetricsOk

`func (o *TeamCreationDto) GetDefaultExperimentSecondaryMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultExperimentSecondaryMetricsOk returns a tuple with the DefaultExperimentSecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentSecondaryMetrics

`func (o *TeamCreationDto) SetDefaultExperimentSecondaryMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultExperimentSecondaryMetrics sets DefaultExperimentSecondaryMetrics field to given value.


### GetDefaultHoldoutMetrics

`func (o *TeamCreationDto) GetDefaultHoldoutMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultHoldoutMetrics returns the DefaultHoldoutMetrics field if non-nil, zero value otherwise.

### GetDefaultHoldoutMetricsOk

`func (o *TeamCreationDto) GetDefaultHoldoutMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultHoldoutMetricsOk returns a tuple with the DefaultHoldoutMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHoldoutMetrics

`func (o *TeamCreationDto) SetDefaultHoldoutMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultHoldoutMetrics sets DefaultHoldoutMetrics field to given value.


### GetChangeTeamConfigs

`func (o *TeamCreationDto) GetChangeTeamConfigs() string`

GetChangeTeamConfigs returns the ChangeTeamConfigs field if non-nil, zero value otherwise.

### GetChangeTeamConfigsOk

`func (o *TeamCreationDto) GetChangeTeamConfigsOk() (*string, bool)`

GetChangeTeamConfigsOk returns a tuple with the ChangeTeamConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTeamConfigs

`func (o *TeamCreationDto) SetChangeTeamConfigs(v string)`

SetChangeTeamConfigs sets ChangeTeamConfigs field to given value.


### GetReviewApproval

`func (o *TeamCreationDto) GetReviewApproval() string`

GetReviewApproval returns the ReviewApproval field if non-nil, zero value otherwise.

### GetReviewApprovalOk

`func (o *TeamCreationDto) GetReviewApprovalOk() (*string, bool)`

GetReviewApprovalOk returns a tuple with the ReviewApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewApproval

`func (o *TeamCreationDto) SetReviewApproval(v string)`

SetReviewApproval sets ReviewApproval field to given value.


### GetDefaultTargetApplications

`func (o *TeamCreationDto) GetDefaultTargetApplications() []string`

GetDefaultTargetApplications returns the DefaultTargetApplications field if non-nil, zero value otherwise.

### GetDefaultTargetApplicationsOk

`func (o *TeamCreationDto) GetDefaultTargetApplicationsOk() (*[]string, bool)`

GetDefaultTargetApplicationsOk returns a tuple with the DefaultTargetApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetApplications

`func (o *TeamCreationDto) SetDefaultTargetApplications(v []string)`

SetDefaultTargetApplications sets DefaultTargetApplications field to given value.


### GetDefaultHoldoutID

`func (o *TeamCreationDto) GetDefaultHoldoutID() nil`

GetDefaultHoldoutID returns the DefaultHoldoutID field if non-nil, zero value otherwise.

### GetDefaultHoldoutIDOk

`func (o *TeamCreationDto) GetDefaultHoldoutIDOk() (*nil, bool)`

GetDefaultHoldoutIDOk returns a tuple with the DefaultHoldoutID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHoldoutID

`func (o *TeamCreationDto) SetDefaultHoldoutID(v nil)`

SetDefaultHoldoutID sets DefaultHoldoutID field to given value.

### HasDefaultHoldoutID

`func (o *TeamCreationDto) HasDefaultHoldoutID() bool`

HasDefaultHoldoutID returns a boolean if a field has been set.

### GetRequireReviews

`func (o *TeamCreationDto) GetRequireReviews() nil`

GetRequireReviews returns the RequireReviews field if non-nil, zero value otherwise.

### GetRequireReviewsOk

`func (o *TeamCreationDto) GetRequireReviewsOk() (*nil, bool)`

GetRequireReviewsOk returns a tuple with the RequireReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReviews

`func (o *TeamCreationDto) SetRequireReviews(v nil)`

SetRequireReviews sets RequireReviews field to given value.

### HasRequireReviews

`func (o *TeamCreationDto) HasRequireReviews() bool`

HasRequireReviews returns a boolean if a field has been set.

### GetRequireGateTemplates

`func (o *TeamCreationDto) GetRequireGateTemplates() nil`

GetRequireGateTemplates returns the RequireGateTemplates field if non-nil, zero value otherwise.

### GetRequireGateTemplatesOk

`func (o *TeamCreationDto) GetRequireGateTemplatesOk() (*nil, bool)`

GetRequireGateTemplatesOk returns a tuple with the RequireGateTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGateTemplates

`func (o *TeamCreationDto) SetRequireGateTemplates(v nil)`

SetRequireGateTemplates sets RequireGateTemplates field to given value.

### HasRequireGateTemplates

`func (o *TeamCreationDto) HasRequireGateTemplates() bool`

HasRequireGateTemplates returns a boolean if a field has been set.

### GetRequireExperimentTemplates

`func (o *TeamCreationDto) GetRequireExperimentTemplates() nil`

GetRequireExperimentTemplates returns the RequireExperimentTemplates field if non-nil, zero value otherwise.

### GetRequireExperimentTemplatesOk

`func (o *TeamCreationDto) GetRequireExperimentTemplatesOk() (*nil, bool)`

GetRequireExperimentTemplatesOk returns a tuple with the RequireExperimentTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExperimentTemplates

`func (o *TeamCreationDto) SetRequireExperimentTemplates(v nil)`

SetRequireExperimentTemplates sets RequireExperimentTemplates field to given value.

### HasRequireExperimentTemplates

`func (o *TeamCreationDto) HasRequireExperimentTemplates() bool`

HasRequireExperimentTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


