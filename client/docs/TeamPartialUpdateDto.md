# TeamPartialUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the team. | [optional] 
**Description** | Pointer to **string** | Description of the team. | [optional] 
**Id** | Pointer to **string** | The ID of the team. | [optional] 
**Members** | Pointer to **[]string** | Array of member email addresses in the team. | [optional] 
**Admins** | Pointer to **[]string** | Array of admin email addresses in the team. | [optional] 
**DefaultGateMetrics** | Pointer to [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default gate metrics for the team. | [optional] 
**DefaultExperimentPrimaryMetrics** | Pointer to [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default primary metrics for experiments in the team. | [optional] 
**DefaultExperimentSecondaryMetrics** | Pointer to [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default secondary metrics for experiments in the team. | [optional] 
**DefaultHoldoutMetrics** | Pointer to [**[]TeamDtoDefaultGateMetricsInner**](TeamDtoDefaultGateMetricsInner.md) | Default holdout metrics for the team. | [optional] 
**ChangeTeamConfigs** | Pointer to **string** | Who can change team configurations: \&quot;anyone\&quot; or \&quot;team_only\&quot;. | [optional] 
**ReviewApproval** | Pointer to **string** | Who can review and approve changes: \&quot;anyone\&quot;, \&quot;team_only\&quot;, or \&quot;admin_only\&quot;. | [optional] 
**DefaultTargetApplications** | Pointer to **[]string** | Default target applications for the team. | [optional] 
**DefaultHoldoutID** | Pointer to [**nil**](nil.md) | Default holdout ID for the team, if applicable. | [optional] 
**RequireReviews** | Pointer to [**nil**](nil.md) | Whether reviews are required for changes, if applicable. | [optional] 
**RequireGateTemplates** | Pointer to [**nil**](nil.md) | Whether gate templates are required for the team, if applicable. | [optional] 
**RequireExperimentTemplates** | Pointer to [**nil**](nil.md) | Whether experiment templates are required for the team, if applicable. | [optional] 

## Methods

### NewTeamPartialUpdateDto

`func NewTeamPartialUpdateDto() *TeamPartialUpdateDto`

NewTeamPartialUpdateDto instantiates a new TeamPartialUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPartialUpdateDtoWithDefaults

`func NewTeamPartialUpdateDtoWithDefaults() *TeamPartialUpdateDto`

NewTeamPartialUpdateDtoWithDefaults instantiates a new TeamPartialUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamPartialUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamPartialUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamPartialUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamPartialUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TeamPartialUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamPartialUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamPartialUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamPartialUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TeamPartialUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamPartialUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamPartialUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamPartialUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMembers

`func (o *TeamPartialUpdateDto) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamPartialUpdateDto) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamPartialUpdateDto) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamPartialUpdateDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetAdmins

`func (o *TeamPartialUpdateDto) GetAdmins() []string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *TeamPartialUpdateDto) GetAdminsOk() (*[]string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *TeamPartialUpdateDto) SetAdmins(v []string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *TeamPartialUpdateDto) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetDefaultGateMetrics

`func (o *TeamPartialUpdateDto) GetDefaultGateMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultGateMetrics returns the DefaultGateMetrics field if non-nil, zero value otherwise.

### GetDefaultGateMetricsOk

`func (o *TeamPartialUpdateDto) GetDefaultGateMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultGateMetricsOk returns a tuple with the DefaultGateMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateMetrics

`func (o *TeamPartialUpdateDto) SetDefaultGateMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultGateMetrics sets DefaultGateMetrics field to given value.

### HasDefaultGateMetrics

`func (o *TeamPartialUpdateDto) HasDefaultGateMetrics() bool`

HasDefaultGateMetrics returns a boolean if a field has been set.

### GetDefaultExperimentPrimaryMetrics

`func (o *TeamPartialUpdateDto) GetDefaultExperimentPrimaryMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultExperimentPrimaryMetrics returns the DefaultExperimentPrimaryMetrics field if non-nil, zero value otherwise.

### GetDefaultExperimentPrimaryMetricsOk

`func (o *TeamPartialUpdateDto) GetDefaultExperimentPrimaryMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultExperimentPrimaryMetricsOk returns a tuple with the DefaultExperimentPrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentPrimaryMetrics

`func (o *TeamPartialUpdateDto) SetDefaultExperimentPrimaryMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultExperimentPrimaryMetrics sets DefaultExperimentPrimaryMetrics field to given value.

### HasDefaultExperimentPrimaryMetrics

`func (o *TeamPartialUpdateDto) HasDefaultExperimentPrimaryMetrics() bool`

HasDefaultExperimentPrimaryMetrics returns a boolean if a field has been set.

### GetDefaultExperimentSecondaryMetrics

`func (o *TeamPartialUpdateDto) GetDefaultExperimentSecondaryMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultExperimentSecondaryMetrics returns the DefaultExperimentSecondaryMetrics field if non-nil, zero value otherwise.

### GetDefaultExperimentSecondaryMetricsOk

`func (o *TeamPartialUpdateDto) GetDefaultExperimentSecondaryMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultExperimentSecondaryMetricsOk returns a tuple with the DefaultExperimentSecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentSecondaryMetrics

`func (o *TeamPartialUpdateDto) SetDefaultExperimentSecondaryMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultExperimentSecondaryMetrics sets DefaultExperimentSecondaryMetrics field to given value.

### HasDefaultExperimentSecondaryMetrics

`func (o *TeamPartialUpdateDto) HasDefaultExperimentSecondaryMetrics() bool`

HasDefaultExperimentSecondaryMetrics returns a boolean if a field has been set.

### GetDefaultHoldoutMetrics

`func (o *TeamPartialUpdateDto) GetDefaultHoldoutMetrics() []TeamDtoDefaultGateMetricsInner`

GetDefaultHoldoutMetrics returns the DefaultHoldoutMetrics field if non-nil, zero value otherwise.

### GetDefaultHoldoutMetricsOk

`func (o *TeamPartialUpdateDto) GetDefaultHoldoutMetricsOk() (*[]TeamDtoDefaultGateMetricsInner, bool)`

GetDefaultHoldoutMetricsOk returns a tuple with the DefaultHoldoutMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHoldoutMetrics

`func (o *TeamPartialUpdateDto) SetDefaultHoldoutMetrics(v []TeamDtoDefaultGateMetricsInner)`

SetDefaultHoldoutMetrics sets DefaultHoldoutMetrics field to given value.

### HasDefaultHoldoutMetrics

`func (o *TeamPartialUpdateDto) HasDefaultHoldoutMetrics() bool`

HasDefaultHoldoutMetrics returns a boolean if a field has been set.

### GetChangeTeamConfigs

`func (o *TeamPartialUpdateDto) GetChangeTeamConfigs() string`

GetChangeTeamConfigs returns the ChangeTeamConfigs field if non-nil, zero value otherwise.

### GetChangeTeamConfigsOk

`func (o *TeamPartialUpdateDto) GetChangeTeamConfigsOk() (*string, bool)`

GetChangeTeamConfigsOk returns a tuple with the ChangeTeamConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTeamConfigs

`func (o *TeamPartialUpdateDto) SetChangeTeamConfigs(v string)`

SetChangeTeamConfigs sets ChangeTeamConfigs field to given value.

### HasChangeTeamConfigs

`func (o *TeamPartialUpdateDto) HasChangeTeamConfigs() bool`

HasChangeTeamConfigs returns a boolean if a field has been set.

### GetReviewApproval

`func (o *TeamPartialUpdateDto) GetReviewApproval() string`

GetReviewApproval returns the ReviewApproval field if non-nil, zero value otherwise.

### GetReviewApprovalOk

`func (o *TeamPartialUpdateDto) GetReviewApprovalOk() (*string, bool)`

GetReviewApprovalOk returns a tuple with the ReviewApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewApproval

`func (o *TeamPartialUpdateDto) SetReviewApproval(v string)`

SetReviewApproval sets ReviewApproval field to given value.

### HasReviewApproval

`func (o *TeamPartialUpdateDto) HasReviewApproval() bool`

HasReviewApproval returns a boolean if a field has been set.

### GetDefaultTargetApplications

`func (o *TeamPartialUpdateDto) GetDefaultTargetApplications() []string`

GetDefaultTargetApplications returns the DefaultTargetApplications field if non-nil, zero value otherwise.

### GetDefaultTargetApplicationsOk

`func (o *TeamPartialUpdateDto) GetDefaultTargetApplicationsOk() (*[]string, bool)`

GetDefaultTargetApplicationsOk returns a tuple with the DefaultTargetApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetApplications

`func (o *TeamPartialUpdateDto) SetDefaultTargetApplications(v []string)`

SetDefaultTargetApplications sets DefaultTargetApplications field to given value.

### HasDefaultTargetApplications

`func (o *TeamPartialUpdateDto) HasDefaultTargetApplications() bool`

HasDefaultTargetApplications returns a boolean if a field has been set.

### GetDefaultHoldoutID

`func (o *TeamPartialUpdateDto) GetDefaultHoldoutID() nil`

GetDefaultHoldoutID returns the DefaultHoldoutID field if non-nil, zero value otherwise.

### GetDefaultHoldoutIDOk

`func (o *TeamPartialUpdateDto) GetDefaultHoldoutIDOk() (*nil, bool)`

GetDefaultHoldoutIDOk returns a tuple with the DefaultHoldoutID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHoldoutID

`func (o *TeamPartialUpdateDto) SetDefaultHoldoutID(v nil)`

SetDefaultHoldoutID sets DefaultHoldoutID field to given value.

### HasDefaultHoldoutID

`func (o *TeamPartialUpdateDto) HasDefaultHoldoutID() bool`

HasDefaultHoldoutID returns a boolean if a field has been set.

### GetRequireReviews

`func (o *TeamPartialUpdateDto) GetRequireReviews() nil`

GetRequireReviews returns the RequireReviews field if non-nil, zero value otherwise.

### GetRequireReviewsOk

`func (o *TeamPartialUpdateDto) GetRequireReviewsOk() (*nil, bool)`

GetRequireReviewsOk returns a tuple with the RequireReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReviews

`func (o *TeamPartialUpdateDto) SetRequireReviews(v nil)`

SetRequireReviews sets RequireReviews field to given value.

### HasRequireReviews

`func (o *TeamPartialUpdateDto) HasRequireReviews() bool`

HasRequireReviews returns a boolean if a field has been set.

### GetRequireGateTemplates

`func (o *TeamPartialUpdateDto) GetRequireGateTemplates() nil`

GetRequireGateTemplates returns the RequireGateTemplates field if non-nil, zero value otherwise.

### GetRequireGateTemplatesOk

`func (o *TeamPartialUpdateDto) GetRequireGateTemplatesOk() (*nil, bool)`

GetRequireGateTemplatesOk returns a tuple with the RequireGateTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireGateTemplates

`func (o *TeamPartialUpdateDto) SetRequireGateTemplates(v nil)`

SetRequireGateTemplates sets RequireGateTemplates field to given value.

### HasRequireGateTemplates

`func (o *TeamPartialUpdateDto) HasRequireGateTemplates() bool`

HasRequireGateTemplates returns a boolean if a field has been set.

### GetRequireExperimentTemplates

`func (o *TeamPartialUpdateDto) GetRequireExperimentTemplates() nil`

GetRequireExperimentTemplates returns the RequireExperimentTemplates field if non-nil, zero value otherwise.

### GetRequireExperimentTemplatesOk

`func (o *TeamPartialUpdateDto) GetRequireExperimentTemplatesOk() (*nil, bool)`

GetRequireExperimentTemplatesOk returns a tuple with the RequireExperimentTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExperimentTemplates

`func (o *TeamPartialUpdateDto) SetRequireExperimentTemplates(v nil)`

SetRequireExperimentTemplates sets RequireExperimentTemplates field to given value.

### HasRequireExperimentTemplates

`func (o *TeamPartialUpdateDto) HasRequireExperimentTemplates() bool`

HasRequireExperimentTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


