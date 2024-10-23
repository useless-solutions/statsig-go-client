# ExternalGateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID | 
**Name** | Pointer to **string** | Optional name for the configuration. | [optional] 
**IdType** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**LastModifierID** | [**nil**](nil.md) | ID of the last modifier. | 
**LastModifiedTime** | [**nil**](nil.md) | Time of the last modification. | 
**LastModifierEmail** | [**nil**](nil.md) | Email of the last modifier. | 
**LastModifierName** | [**nil**](nil.md) | Name of the last modifier. | 
**CreatorID** | Pointer to [**nil**](nil.md) |  | [optional] 
**CreatedTime** | **float32** | Timestamp when the entity was created. | 
**CreatorName** | [**nil**](nil.md) | Name of the creator. | 
**CreatorEmail** | Pointer to [**nil**](nil.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TargetApps** | Pointer to [**DynamicConfigDtoTargetApps**](DynamicConfigDtoTargetApps.md) |  | [optional] 
**HoldoutIDs** | Pointer to **[]string** | Holdouts applied to this configuration. | [optional] 
**Team** | Pointer to [**nil**](nil.md) |  | [optional] 
**Version** | Pointer to **float32** | Version number | [optional] 
**ChecksPerHour** | [**nil**](nil.md) |  | 
**Status** | **string** |  | 
**Type** | **string** |  | 
**TypeReason** | **string** |  | 
**Owner** | Pointer to [**nil**](nil.md) | Schema for owner data including ID, type, name. Note that if Entity is created by CONSOLE API, owner will be undefined. | [optional] 
**IsEnabled** | **bool** |  | 
**Rules** | [**[]ExternalGateDtoRulesInner**](ExternalGateDtoRulesInner.md) |  | 
**MeasureMetricLifts** | Pointer to **bool** |  | [optional] 
**MonitoringMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) |  | [optional] 
**ReviewSettings** | Pointer to [**ExternalGateDtoReviewSettings**](ExternalGateDtoReviewSettings.md) |  | [optional] 
**ActiveReview** | Pointer to [**ExternalGateDtoActiveReview**](ExternalGateDtoActiveReview.md) |  | [optional] 

## Methods

### NewExternalGateDto

`func NewExternalGateDto(id string, description string, lastModifierID nil, lastModifiedTime nil, lastModifierEmail nil, lastModifierName nil, createdTime float32, creatorName nil, checksPerHour nil, status string, type_ string, typeReason string, isEnabled bool, rules []ExternalGateDtoRulesInner, ) *ExternalGateDto`

NewExternalGateDto instantiates a new ExternalGateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGateDtoWithDefaults

`func NewExternalGateDtoWithDefaults() *ExternalGateDto`

NewExternalGateDtoWithDefaults instantiates a new ExternalGateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalGateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalGateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalGateDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ExternalGateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalGateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalGateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalGateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdType

`func (o *ExternalGateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *ExternalGateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *ExternalGateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *ExternalGateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalGateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalGateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalGateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLastModifierID

`func (o *ExternalGateDto) GetLastModifierID() nil`

GetLastModifierID returns the LastModifierID field if non-nil, zero value otherwise.

### GetLastModifierIDOk

`func (o *ExternalGateDto) GetLastModifierIDOk() (*nil, bool)`

GetLastModifierIDOk returns a tuple with the LastModifierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierID

`func (o *ExternalGateDto) SetLastModifierID(v nil)`

SetLastModifierID sets LastModifierID field to given value.


### GetLastModifiedTime

`func (o *ExternalGateDto) GetLastModifiedTime() nil`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *ExternalGateDto) GetLastModifiedTimeOk() (*nil, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *ExternalGateDto) SetLastModifiedTime(v nil)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastModifierEmail

`func (o *ExternalGateDto) GetLastModifierEmail() nil`

GetLastModifierEmail returns the LastModifierEmail field if non-nil, zero value otherwise.

### GetLastModifierEmailOk

`func (o *ExternalGateDto) GetLastModifierEmailOk() (*nil, bool)`

GetLastModifierEmailOk returns a tuple with the LastModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierEmail

`func (o *ExternalGateDto) SetLastModifierEmail(v nil)`

SetLastModifierEmail sets LastModifierEmail field to given value.


### GetLastModifierName

`func (o *ExternalGateDto) GetLastModifierName() nil`

GetLastModifierName returns the LastModifierName field if non-nil, zero value otherwise.

### GetLastModifierNameOk

`func (o *ExternalGateDto) GetLastModifierNameOk() (*nil, bool)`

GetLastModifierNameOk returns a tuple with the LastModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierName

`func (o *ExternalGateDto) SetLastModifierName(v nil)`

SetLastModifierName sets LastModifierName field to given value.


### GetCreatorID

`func (o *ExternalGateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *ExternalGateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *ExternalGateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *ExternalGateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ExternalGateDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ExternalGateDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ExternalGateDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreatorName

`func (o *ExternalGateDto) GetCreatorName() nil`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ExternalGateDto) GetCreatorNameOk() (*nil, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ExternalGateDto) SetCreatorName(v nil)`

SetCreatorName sets CreatorName field to given value.


### GetCreatorEmail

`func (o *ExternalGateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *ExternalGateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *ExternalGateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *ExternalGateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetTags

`func (o *ExternalGateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternalGateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternalGateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExternalGateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetApps

`func (o *ExternalGateDto) GetTargetApps() DynamicConfigDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *ExternalGateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *ExternalGateDto) SetTargetApps(v DynamicConfigDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *ExternalGateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetHoldoutIDs

`func (o *ExternalGateDto) GetHoldoutIDs() []string`

GetHoldoutIDs returns the HoldoutIDs field if non-nil, zero value otherwise.

### GetHoldoutIDsOk

`func (o *ExternalGateDto) GetHoldoutIDsOk() (*[]string, bool)`

GetHoldoutIDsOk returns a tuple with the HoldoutIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutIDs

`func (o *ExternalGateDto) SetHoldoutIDs(v []string)`

SetHoldoutIDs sets HoldoutIDs field to given value.

### HasHoldoutIDs

`func (o *ExternalGateDto) HasHoldoutIDs() bool`

HasHoldoutIDs returns a boolean if a field has been set.

### GetTeam

`func (o *ExternalGateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExternalGateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExternalGateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExternalGateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetVersion

`func (o *ExternalGateDto) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExternalGateDto) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExternalGateDto) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExternalGateDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetChecksPerHour

`func (o *ExternalGateDto) GetChecksPerHour() nil`

GetChecksPerHour returns the ChecksPerHour field if non-nil, zero value otherwise.

### GetChecksPerHourOk

`func (o *ExternalGateDto) GetChecksPerHourOk() (*nil, bool)`

GetChecksPerHourOk returns a tuple with the ChecksPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksPerHour

`func (o *ExternalGateDto) SetChecksPerHour(v nil)`

SetChecksPerHour sets ChecksPerHour field to given value.


### GetStatus

`func (o *ExternalGateDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalGateDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalGateDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ExternalGateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalGateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalGateDto) SetType(v string)`

SetType sets Type field to given value.


### GetTypeReason

`func (o *ExternalGateDto) GetTypeReason() string`

GetTypeReason returns the TypeReason field if non-nil, zero value otherwise.

### GetTypeReasonOk

`func (o *ExternalGateDto) GetTypeReasonOk() (*string, bool)`

GetTypeReasonOk returns a tuple with the TypeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeReason

`func (o *ExternalGateDto) SetTypeReason(v string)`

SetTypeReason sets TypeReason field to given value.


### GetOwner

`func (o *ExternalGateDto) GetOwner() nil`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExternalGateDto) GetOwnerOk() (*nil, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExternalGateDto) SetOwner(v nil)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExternalGateDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ExternalGateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ExternalGateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ExternalGateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetRules

`func (o *ExternalGateDto) GetRules() []ExternalGateDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ExternalGateDto) GetRulesOk() (*[]ExternalGateDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ExternalGateDto) SetRules(v []ExternalGateDtoRulesInner)`

SetRules sets Rules field to given value.


### GetMeasureMetricLifts

`func (o *ExternalGateDto) GetMeasureMetricLifts() bool`

GetMeasureMetricLifts returns the MeasureMetricLifts field if non-nil, zero value otherwise.

### GetMeasureMetricLiftsOk

`func (o *ExternalGateDto) GetMeasureMetricLiftsOk() (*bool, bool)`

GetMeasureMetricLiftsOk returns a tuple with the MeasureMetricLifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureMetricLifts

`func (o *ExternalGateDto) SetMeasureMetricLifts(v bool)`

SetMeasureMetricLifts sets MeasureMetricLifts field to given value.

### HasMeasureMetricLifts

`func (o *ExternalGateDto) HasMeasureMetricLifts() bool`

HasMeasureMetricLifts returns a boolean if a field has been set.

### GetMonitoringMetrics

`func (o *ExternalGateDto) GetMonitoringMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetMonitoringMetrics returns the MonitoringMetrics field if non-nil, zero value otherwise.

### GetMonitoringMetricsOk

`func (o *ExternalGateDto) GetMonitoringMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetMonitoringMetricsOk returns a tuple with the MonitoringMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMetrics

`func (o *ExternalGateDto) SetMonitoringMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetMonitoringMetrics sets MonitoringMetrics field to given value.

### HasMonitoringMetrics

`func (o *ExternalGateDto) HasMonitoringMetrics() bool`

HasMonitoringMetrics returns a boolean if a field has been set.

### GetReviewSettings

`func (o *ExternalGateDto) GetReviewSettings() ExternalGateDtoReviewSettings`

GetReviewSettings returns the ReviewSettings field if non-nil, zero value otherwise.

### GetReviewSettingsOk

`func (o *ExternalGateDto) GetReviewSettingsOk() (*ExternalGateDtoReviewSettings, bool)`

GetReviewSettingsOk returns a tuple with the ReviewSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewSettings

`func (o *ExternalGateDto) SetReviewSettings(v ExternalGateDtoReviewSettings)`

SetReviewSettings sets ReviewSettings field to given value.

### HasReviewSettings

`func (o *ExternalGateDto) HasReviewSettings() bool`

HasReviewSettings returns a boolean if a field has been set.

### GetActiveReview

`func (o *ExternalGateDto) GetActiveReview() ExternalGateDtoActiveReview`

GetActiveReview returns the ActiveReview field if non-nil, zero value otherwise.

### GetActiveReviewOk

`func (o *ExternalGateDto) GetActiveReviewOk() (*ExternalGateDtoActiveReview, bool)`

GetActiveReviewOk returns a tuple with the ActiveReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveReview

`func (o *ExternalGateDto) SetActiveReview(v ExternalGateDtoActiveReview)`

SetActiveReview sets ActiveReview field to given value.

### HasActiveReview

`func (o *ExternalGateDto) HasActiveReview() bool`

HasActiveReview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


