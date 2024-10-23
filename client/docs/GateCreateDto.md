# GateCreateDto

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
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewGateCreateDto

`func NewGateCreateDto() *GateCreateDto`

NewGateCreateDto instantiates a new GateCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGateCreateDtoWithDefaults

`func NewGateCreateDtoWithDefaults() *GateCreateDto`

NewGateCreateDtoWithDefaults instantiates a new GateCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GateCreateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GateCreateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GateCreateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *GateCreateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *GateCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GateCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GateCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GateCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *GateCreateDto) GetRules() []ExternalGateDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GateCreateDto) GetRulesOk() (*[]ExternalGateDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GateCreateDto) SetRules(v []ExternalGateDtoRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GateCreateDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *GateCreateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GateCreateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GateCreateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GateCreateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *GateCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GateCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GateCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GateCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdType

`func (o *GateCreateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *GateCreateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *GateCreateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *GateCreateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetTargetApps

`func (o *GateCreateDto) GetTargetApps() DynamicConfigDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *GateCreateDto) GetTargetAppsOk() (*DynamicConfigDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *GateCreateDto) SetTargetApps(v DynamicConfigDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *GateCreateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetCreatorID

`func (o *GateCreateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *GateCreateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *GateCreateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *GateCreateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *GateCreateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *GateCreateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *GateCreateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *GateCreateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetTeam

`func (o *GateCreateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GateCreateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GateCreateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GateCreateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetMeasureMetricLifts

`func (o *GateCreateDto) GetMeasureMetricLifts() bool`

GetMeasureMetricLifts returns the MeasureMetricLifts field if non-nil, zero value otherwise.

### GetMeasureMetricLiftsOk

`func (o *GateCreateDto) GetMeasureMetricLiftsOk() (*bool, bool)`

GetMeasureMetricLiftsOk returns a tuple with the MeasureMetricLifts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureMetricLifts

`func (o *GateCreateDto) SetMeasureMetricLifts(v bool)`

SetMeasureMetricLifts sets MeasureMetricLifts field to given value.

### HasMeasureMetricLifts

`func (o *GateCreateDto) HasMeasureMetricLifts() bool`

HasMeasureMetricLifts returns a boolean if a field has been set.

### GetMonitoringMetrics

`func (o *GateCreateDto) GetMonitoringMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetMonitoringMetrics returns the MonitoringMetrics field if non-nil, zero value otherwise.

### GetMonitoringMetricsOk

`func (o *GateCreateDto) GetMonitoringMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetMonitoringMetricsOk returns a tuple with the MonitoringMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMetrics

`func (o *GateCreateDto) SetMonitoringMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetMonitoringMetrics sets MonitoringMetrics field to given value.

### HasMonitoringMetrics

`func (o *GateCreateDto) HasMonitoringMetrics() bool`

HasMonitoringMetrics returns a boolean if a field has been set.

### GetName

`func (o *GateCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GateCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GateCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GateCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *GateCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GateCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GateCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GateCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


