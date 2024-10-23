# GatePartialUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]ExternalGateDtoRulesInner**](ExternalGateDtoRulesInner.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IdType** | Pointer to **string** |  | [optional] 
**TargetApps** | Pointer to [**DynamicConfigDtoTargetApps**](DynamicConfigDtoTargetApps.md) |  | [optional] 
**CreatorID** | Pointer to [**nil**](nil.md) |  | [optional] 
**CreatorEmail** | Pointer to [**nil**](nil.md) |  | [optional] 
**Team** | Pointer to [**nil**](nil.md) |  | [optional] 
**MeasureMetricLifts** | Pointer to **bool** |  | [optional] 
**MonitoringMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) |  | [optional] 

## Methods

### NewGatePartialUpdateDto

`func NewGatePartialUpdateDto() *GatePartialUpdateDto`

NewGatePartialUpdateDto instantiates a new GatePartialUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatePartialUpdateDtoWithDefaults

`func NewGatePartialUpdateDtoWithDefaults() *GatePartialUpdateDto`

NewGatePartialUpdateDtoWithDefaults instantiates a new GatePartialUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GatePartialUpdateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GatePartialUpdateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GatePartialUpdateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *GatePartialUpdateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *GatePartialUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatePartialUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatePartialUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatePartialUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *GatePartialUpdateDto) GetRules() []ExternalGateDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GatePartialUpdateDto) GetRulesOk() (*[]ExternalGateDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GatePartialUpdateDto) SetRules(v []ExternalGateDtoRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GatePartialUpdateDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *GatePartialUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatePartialUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatePartialUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatePartialUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *GatePartialUpdateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatePartialUpdateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatePartialUpdateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatePartialUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdType

`func (o *GatePartialUpdateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *GatePartialUpdateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *GatePartialUpdateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *GatePartialUpdateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTargetApps

`func (o *GatePartialUpdateDto) GetTargetApps() DynamicConfigDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *GatePartialUpdateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *GatePartialUpdateDto) SetTargetApps(v DynamicConfigDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *GatePartialUpdateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetCreatorID

`func (o *GatePartialUpdateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *GatePartialUpdateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *GatePartialUpdateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *GatePartialUpdateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *GatePartialUpdateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *GatePartialUpdateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *GatePartialUpdateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *GatePartialUpdateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetTeam

`func (o *GatePartialUpdateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GatePartialUpdateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GatePartialUpdateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GatePartialUpdateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetMeasureMetricLifts

`func (o *GatePartialUpdateDto) GetMeasureMetricLifts() bool`

GetMeasureMetricLifts returns the MeasureMetricLifts field if non-nil, zero value otherwise.

### GetMeasureMetricLiftsOk

`func (o *GatePartialUpdateDto) GetMeasureMetricLiftsOk() (*bool, bool)`

GetMeasureMetricLiftsOk returns a tuple with the MeasureMetricLifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureMetricLifts

`func (o *GatePartialUpdateDto) SetMeasureMetricLifts(v bool)`

SetMeasureMetricLifts sets MeasureMetricLifts field to given value.

### HasMeasureMetricLifts

`func (o *GatePartialUpdateDto) HasMeasureMetricLifts() bool`

HasMeasureMetricLifts returns a boolean if a field has been set.

### GetMonitoringMetrics

`func (o *GatePartialUpdateDto) GetMonitoringMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetMonitoringMetrics returns the MonitoringMetrics field if non-nil, zero value otherwise.

### GetMonitoringMetricsOk

`func (o *GatePartialUpdateDto) GetMonitoringMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetMonitoringMetricsOk returns a tuple with the MonitoringMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMetrics

`func (o *GatePartialUpdateDto) SetMonitoringMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetMonitoringMetrics sets MonitoringMetrics field to given value.

### HasMonitoringMetrics

`func (o *GatePartialUpdateDto) HasMonitoringMetrics() bool`

HasMonitoringMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


