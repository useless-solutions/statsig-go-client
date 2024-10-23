# GateFullUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** |  | 
**Description** | **string** |  | 
**Rules** | [**[]ExternalGateDtoRulesInner**](ExternalGateDtoRulesInner.md) |  | 
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

### NewGateFullUpdateDto

`func NewGateFullUpdateDto(isEnabled bool, description string, rules []ExternalGateDtoRulesInner, ) *GateFullUpdateDto`

NewGateFullUpdateDto instantiates a new GateFullUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGateFullUpdateDtoWithDefaults

`func NewGateFullUpdateDtoWithDefaults() *GateFullUpdateDto`

NewGateFullUpdateDtoWithDefaults instantiates a new GateFullUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GateFullUpdateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GateFullUpdateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GateFullUpdateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDescription

`func (o *GateFullUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GateFullUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GateFullUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *GateFullUpdateDto) GetRules() []ExternalGateDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GateFullUpdateDto) GetRulesOk() (*[]ExternalGateDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GateFullUpdateDto) SetRules(v []ExternalGateDtoRulesInner)`

SetRules sets Rules field to given value.


### GetTags

`func (o *GateFullUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GateFullUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GateFullUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GateFullUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *GateFullUpdateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GateFullUpdateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GateFullUpdateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GateFullUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdType

`func (o *GateFullUpdateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *GateFullUpdateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *GateFullUpdateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *GateFullUpdateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTargetApps

`func (o *GateFullUpdateDto) GetTargetApps() DynamicConfigDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *GateFullUpdateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *GateFullUpdateDto) SetTargetApps(v DynamicConfigDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *GateFullUpdateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetCreatorID

`func (o *GateFullUpdateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *GateFullUpdateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *GateFullUpdateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *GateFullUpdateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *GateFullUpdateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *GateFullUpdateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *GateFullUpdateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *GateFullUpdateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetTeam

`func (o *GateFullUpdateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GateFullUpdateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GateFullUpdateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GateFullUpdateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetMeasureMetricLifts

`func (o *GateFullUpdateDto) GetMeasureMetricLifts() bool`

GetMeasureMetricLifts returns the MeasureMetricLifts field if non-nil, zero value otherwise.

### GetMeasureMetricLiftsOk

`func (o *GateFullUpdateDto) GetMeasureMetricLiftsOk() (*bool, bool)`

GetMeasureMetricLiftsOk returns a tuple with the MeasureMetricLifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureMetricLifts

`func (o *GateFullUpdateDto) SetMeasureMetricLifts(v bool)`

SetMeasureMetricLifts sets MeasureMetricLifts field to given value.

### HasMeasureMetricLifts

`func (o *GateFullUpdateDto) HasMeasureMetricLifts() bool`

HasMeasureMetricLifts returns a boolean if a field has been set.

### GetMonitoringMetrics

`func (o *GateFullUpdateDto) GetMonitoringMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetMonitoringMetrics returns the MonitoringMetrics field if non-nil, zero value otherwise.

### GetMonitoringMetricsOk

`func (o *GateFullUpdateDto) GetMonitoringMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetMonitoringMetricsOk returns a tuple with the MonitoringMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMetrics

`func (o *GateFullUpdateDto) SetMonitoringMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetMonitoringMetrics sets MonitoringMetrics field to given value.

### HasMonitoringMetrics

`func (o *GateFullUpdateDto) HasMonitoringMetrics() bool`

HasMonitoringMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


