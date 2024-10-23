# MetricCreationContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new metric, which identifies it within the system. | 
**Type** | **string** | The type of the metric. Allowed values include sum, ratio, mean, event_count_sum, composite, composite_sum, undefined, funnel, user_warehouse. | 
**IsHidden** | Pointer to **bool** | Indicates whether the metric should be hidden from the user interface. | [optional] 
**IsVerified** | Pointer to **bool** | Marks the metric as verified for internal trustworthiness. | [optional] 
**IsReadOnly** | Pointer to **bool** | Set to true to make the metric definition editable only from the Console API. | [optional] 
**UnitTypes** | Pointer to **[]string** | Array of unit types associated with the metric, such as stableID or userID. | [optional] 
**MetricEvents** | Pointer to [**[]ExternalMetricDefinitionContractDtoMetricEventsInner**](ExternalMetricDefinitionContractDtoMetricEventsInner.md) | An array of event definitions used to compute the metric. | [optional] 
**MetricComponentMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) | List of input metrics used to calculate the new metric for composite types. | [optional] 
**Description** | Pointer to **string** | A description of the new metric, providing context and purpose. | [optional] 
**Directionality** | Pointer to **string** | Indicates the desired change direction for the metric. Use \&quot;increase\&quot; for positive changes and \&quot;decrease\&quot; for negative changes. | [optional] [default to "increase"]
**Tags** | Pointer to [**MetricCreationContractDtoTags**](MetricCreationContractDtoTags.md) |  | [optional] 
**IsPermanent** | Pointer to **bool** | Indicates whether the metric is permanent and should not be deleted. | [optional] 
**RollupTimeWindow** | Pointer to **string** | Time window for the metric rollup. Specify \&quot;custom\&quot; for a customized time window. | [optional] 
**CustomRollUpStart** | Pointer to **float32** | Custom time window start date in days since exposure. | [optional] 
**CustomRollUpEnd** | Pointer to **float32** | Custom time window end date in days since exposure. | [optional] 
**FunnelEventList** | Pointer to [**[]ExternalMetricDefinitionContractDtoFunnelEventListInner**](ExternalMetricDefinitionContractDtoFunnelEventListInner.md) | List of events used to create funnel metrics. | [optional] 
**FunnelCountDistinct** | Pointer to **string** | Specifies whether to count events or distinct users for the funnel metric. | [optional] 
**WarehouseNative** | Pointer to [**ExternalMetricDefinitionContractDtoWarehouseNative**](ExternalMetricDefinitionContractDtoWarehouseNative.md) |  | [optional] 
**Team** | Pointer to [**nil**](nil.md) | The team associated with the metric, applicable for enterprise environments. | [optional] 

## Methods

### NewMetricCreationContractDto

`func NewMetricCreationContractDto(name string, type_ string, ) *MetricCreationContractDto`

NewMetricCreationContractDto instantiates a new MetricCreationContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCreationContractDtoWithDefaults

`func NewMetricCreationContractDtoWithDefaults() *MetricCreationContractDto`

NewMetricCreationContractDtoWithDefaults instantiates a new MetricCreationContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricCreationContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricCreationContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricCreationContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MetricCreationContractDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricCreationContractDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricCreationContractDto) SetType(v string)`

SetType sets Type field to given value.


### GetIsHidden

`func (o *MetricCreationContractDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *MetricCreationContractDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *MetricCreationContractDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *MetricCreationContractDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsVerified

`func (o *MetricCreationContractDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *MetricCreationContractDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *MetricCreationContractDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *MetricCreationContractDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *MetricCreationContractDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MetricCreationContractDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *MetricCreationContractDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *MetricCreationContractDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetUnitTypes

`func (o *MetricCreationContractDto) GetUnitTypes() []string`

GetUnitTypes returns the UnitTypes field if non-nil, zero value otherwise.

### GetUnitTypesOk

`func (o *MetricCreationContractDto) GetUnitTypesOk() (*[]string, bool)`

GetUnitTypesOk returns a tuple with the UnitTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTypes

`func (o *MetricCreationContractDto) SetUnitTypes(v []string)`

SetUnitTypes sets UnitTypes field to given value.

### HasUnitTypes

`func (o *MetricCreationContractDto) HasUnitTypes() bool`

HasUnitTypes returns a boolean if a field has been set.

### GetMetricEvents

`func (o *MetricCreationContractDto) GetMetricEvents() []ExternalMetricDefinitionContractDtoMetricEventsInner`

GetMetricEvents returns the MetricEvents field if non-nil, zero value otherwise.

### GetMetricEventsOk

`func (o *MetricCreationContractDto) GetMetricEventsOk() (*[]ExternalMetricDefinitionContractDtoMetricEventsInner, bool)`

GetMetricEventsOk returns a tuple with the MetricEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEvents

`func (o *MetricCreationContractDto) SetMetricEvents(v []ExternalMetricDefinitionContractDtoMetricEventsInner)`

SetMetricEvents sets MetricEvents field to given value.

### HasMetricEvents

`func (o *MetricCreationContractDto) HasMetricEvents() bool`

HasMetricEvents returns a boolean if a field has been set.

### GetMetricComponentMetrics

`func (o *MetricCreationContractDto) GetMetricComponentMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetMetricComponentMetrics returns the MetricComponentMetrics field if non-nil, zero value otherwise.

### GetMetricComponentMetricsOk

`func (o *MetricCreationContractDto) GetMetricComponentMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetMetricComponentMetricsOk returns a tuple with the MetricComponentMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricComponentMetrics

`func (o *MetricCreationContractDto) SetMetricComponentMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetMetricComponentMetrics sets MetricComponentMetrics field to given value.

### HasMetricComponentMetrics

`func (o *MetricCreationContractDto) HasMetricComponentMetrics() bool`

HasMetricComponentMetrics returns a boolean if a field has been set.

### GetDescription

`func (o *MetricCreationContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricCreationContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricCreationContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricCreationContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectionality

`func (o *MetricCreationContractDto) GetDirectionality() string`

GetDirectionality returns the Directionality field if non-nil, zero value otherwise.

### GetDirectionalityOk

`func (o *MetricCreationContractDto) GetDirectionalityOk() (*string, bool)`

GetDirectionalityOk returns a tuple with the Directionality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionality

`func (o *MetricCreationContractDto) SetDirectionality(v string)`

SetDirectionality sets Directionality field to given value.

### HasDirectionality

`func (o *MetricCreationContractDto) HasDirectionality() bool`

HasDirectionality returns a boolean if a field has been set.

### GetTags

`func (o *MetricCreationContractDto) GetTags() MetricCreationContractDtoTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricCreationContractDto) GetTagsOk() (*MetricCreationContractDtoTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricCreationContractDto) SetTags(v MetricCreationContractDtoTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricCreationContractDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsPermanent

`func (o *MetricCreationContractDto) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *MetricCreationContractDto) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *MetricCreationContractDto) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *MetricCreationContractDto) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetRollupTimeWindow

`func (o *MetricCreationContractDto) GetRollupTimeWindow() string`

GetRollupTimeWindow returns the RollupTimeWindow field if non-nil, zero value otherwise.

### GetRollupTimeWindowOk

`func (o *MetricCreationContractDto) GetRollupTimeWindowOk() (*string, bool)`

GetRollupTimeWindowOk returns a tuple with the RollupTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupTimeWindow

`func (o *MetricCreationContractDto) SetRollupTimeWindow(v string)`

SetRollupTimeWindow sets RollupTimeWindow field to given value.

### HasRollupTimeWindow

`func (o *MetricCreationContractDto) HasRollupTimeWindow() bool`

HasRollupTimeWindow returns a boolean if a field has been set.

### GetCustomRollUpStart

`func (o *MetricCreationContractDto) GetCustomRollUpStart() float32`

GetCustomRollUpStart returns the CustomRollUpStart field if non-nil, zero value otherwise.

### GetCustomRollUpStartOk

`func (o *MetricCreationContractDto) GetCustomRollUpStartOk() (*float32, bool)`

GetCustomRollUpStartOk returns a tuple with the CustomRollUpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRollUpStart

`func (o *MetricCreationContractDto) SetCustomRollUpStart(v float32)`

SetCustomRollUpStart sets CustomRollUpStart field to given value.

### HasCustomRollUpStart

`func (o *MetricCreationContractDto) HasCustomRollUpStart() bool`

HasCustomRollUpStart returns a boolean if a field has been set.

### GetCustomRollUpEnd

`func (o *MetricCreationContractDto) GetCustomRollUpEnd() float32`

GetCustomRollUpEnd returns the CustomRollUpEnd field if non-nil, zero value otherwise.

### GetCustomRollUpEndOk

`func (o *MetricCreationContractDto) GetCustomRollUpEndOk() (*float32, bool)`

GetCustomRollUpEndOk returns a tuple with the CustomRollUpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRollUpEnd

`func (o *MetricCreationContractDto) SetCustomRollUpEnd(v float32)`

SetCustomRollUpEnd sets CustomRollUpEnd field to given value.

### HasCustomRollUpEnd

`func (o *MetricCreationContractDto) HasCustomRollUpEnd() bool`

HasCustomRollUpEnd returns a boolean if a field has been set.

### GetFunnelEventList

`func (o *MetricCreationContractDto) GetFunnelEventList() []ExternalMetricDefinitionContractDtoFunnelEventListInner`

GetFunnelEventList returns the FunnelEventList field if non-nil, zero value otherwise.

### GetFunnelEventListOk

`func (o *MetricCreationContractDto) GetFunnelEventListOk() (*[]ExternalMetricDefinitionContractDtoFunnelEventListInner, bool)`

GetFunnelEventListOk returns a tuple with the FunnelEventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelEventList

`func (o *MetricCreationContractDto) SetFunnelEventList(v []ExternalMetricDefinitionContractDtoFunnelEventListInner)`

SetFunnelEventList sets FunnelEventList field to given value.

### HasFunnelEventList

`func (o *MetricCreationContractDto) HasFunnelEventList() bool`

HasFunnelEventList returns a boolean if a field has been set.

### GetFunnelCountDistinct

`func (o *MetricCreationContractDto) GetFunnelCountDistinct() string`

GetFunnelCountDistinct returns the FunnelCountDistinct field if non-nil, zero value otherwise.

### GetFunnelCountDistinctOk

`func (o *MetricCreationContractDto) GetFunnelCountDistinctOk() (*string, bool)`

GetFunnelCountDistinctOk returns a tuple with the FunnelCountDistinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCountDistinct

`func (o *MetricCreationContractDto) SetFunnelCountDistinct(v string)`

SetFunnelCountDistinct sets FunnelCountDistinct field to given value.

### HasFunnelCountDistinct

`func (o *MetricCreationContractDto) HasFunnelCountDistinct() bool`

HasFunnelCountDistinct returns a boolean if a field has been set.

### GetWarehouseNative

`func (o *MetricCreationContractDto) GetWarehouseNative() ExternalMetricDefinitionContractDtoWarehouseNative`

GetWarehouseNative returns the WarehouseNative field if non-nil, zero value otherwise.

### GetWarehouseNativeOk

`func (o *MetricCreationContractDto) GetWarehouseNativeOk() (*ExternalMetricDefinitionContractDtoWarehouseNative, bool)`

GetWarehouseNativeOk returns a tuple with the WarehouseNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseNative

`func (o *MetricCreationContractDto) SetWarehouseNative(v ExternalMetricDefinitionContractDtoWarehouseNative)`

SetWarehouseNative sets WarehouseNative field to given value.

### HasWarehouseNative

`func (o *MetricCreationContractDto) HasWarehouseNative() bool`

HasWarehouseNative returns a boolean if a field has been set.

### GetTeam

`func (o *MetricCreationContractDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MetricCreationContractDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MetricCreationContractDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MetricCreationContractDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


