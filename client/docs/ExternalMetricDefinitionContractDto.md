# ExternalMetricDefinitionContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the metric, serving as its primary identifier. | 
**Type** | **string** | The type of the metric, defining its aggregation method and characteristics. | 
**IsHidden** | Pointer to **bool** | Indicates if the metric is hidden from the user interface. | [optional] 
**IsVerified** | Pointer to **bool** | Marks the metric as verified, indicating trustworthiness within the organization. | [optional] 
**IsReadOnly** | Pointer to **bool** | Set to true to make the metric definition editable only through the Console API. | [optional] 
**UnitTypes** | Pointer to **[]string** | Array of unit types associated with the metric, such as stableID or userID. | [optional] 
**MetricEvents** | Pointer to [**[]ExternalMetricDefinitionContractDtoMetricEventsInner**](ExternalMetricDefinitionContractDtoMetricEventsInner.md) | An array of event definitions used to compute the metric. | [optional] 
**MetricComponentMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) | List of input metrics used to calculate the new metric for composite types. | [optional] 
**Description** | Pointer to **string** | An optional description providing additional context about the metric. | [optional] 
**Directionality** | **string** | Specifies the desired directionality for the metric, indicating whether an increase or decrease is favorable. | 
**Tags** | Pointer to **[]string** | Optional tags for categorizing the metric and improving searchability. | [optional] 
**IsPermanent** | Pointer to **bool** | Indicates whether the metric is permanent and should not be deleted. | [optional] 
**RollupTimeWindow** | Pointer to **string** | Time window for the metric rollup. Specify \&quot;custom\&quot; for a customized time window. | [optional] 
**CustomRollUpStart** | Pointer to **float32** | Custom time window start date in days since exposure. | [optional] 
**CustomRollUpEnd** | Pointer to **float32** | Custom time window end date in days since exposure. | [optional] 
**FunnelEventList** | Pointer to [**[]ExternalMetricDefinitionContractDtoFunnelEventListInner**](ExternalMetricDefinitionContractDtoFunnelEventListInner.md) | List of events used to create funnel metrics. | [optional] 
**FunnelCountDistinct** | Pointer to **string** | Specifies whether to count events or distinct users for the funnel metric. | [optional] 
**WarehouseNative** | Pointer to [**ExternalMetricDefinitionContractDtoWarehouseNative**](ExternalMetricDefinitionContractDtoWarehouseNative.md) |  | [optional] 
**Team** | Pointer to [**nil**](nil.md) | The team associated with the metric, applicable for enterprise environments. | [optional] 
**Id** | **string** | Unique identifier for the metric, used for referencing within the system. | 
**Lineage** | [**ExternalMetricDefinitionContractDtoLineage**](ExternalMetricDefinitionContractDtoLineage.md) |  | 
**CreatorName** | Pointer to [**nil**](nil.md) | Name of the person who created the metric, if available. | [optional] 
**CreatorEmail** | Pointer to [**nil**](nil.md) | Email address of the metric creator for contact purposes. | [optional] 
**CreatedTime** | Pointer to **float32** | Timestamp indicating when the metric was created. | [optional] 
**Owner** | Pointer to [**ExternalMetricDefinitionContractDtoOwner**](ExternalMetricDefinitionContractDtoOwner.md) |  | [optional] 

## Methods

### NewExternalMetricDefinitionContractDto

`func NewExternalMetricDefinitionContractDto(name string, type_ string, directionality string, id string, lineage ExternalMetricDefinitionContractDtoLineage, ) *ExternalMetricDefinitionContractDto`

NewExternalMetricDefinitionContractDto instantiates a new ExternalMetricDefinitionContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMetricDefinitionContractDtoWithDefaults

`func NewExternalMetricDefinitionContractDtoWithDefaults() *ExternalMetricDefinitionContractDto`

NewExternalMetricDefinitionContractDtoWithDefaults instantiates a new ExternalMetricDefinitionContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalMetricDefinitionContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalMetricDefinitionContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalMetricDefinitionContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ExternalMetricDefinitionContractDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalMetricDefinitionContractDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalMetricDefinitionContractDto) SetType(v string)`

SetType sets Type field to given value.


### GetIsHidden

`func (o *ExternalMetricDefinitionContractDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ExternalMetricDefinitionContractDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ExternalMetricDefinitionContractDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *ExternalMetricDefinitionContractDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsVerified

`func (o *ExternalMetricDefinitionContractDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *ExternalMetricDefinitionContractDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *ExternalMetricDefinitionContractDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *ExternalMetricDefinitionContractDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *ExternalMetricDefinitionContractDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *ExternalMetricDefinitionContractDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *ExternalMetricDefinitionContractDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *ExternalMetricDefinitionContractDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetUnitTypes

`func (o *ExternalMetricDefinitionContractDto) GetUnitTypes() []string`

GetUnitTypes returns the UnitTypes field if non-nil, zero value otherwise.

### GetUnitTypesOk

`func (o *ExternalMetricDefinitionContractDto) GetUnitTypesOk() (*[]string, bool)`

GetUnitTypesOk returns a tuple with the UnitTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTypes

`func (o *ExternalMetricDefinitionContractDto) SetUnitTypes(v []string)`

SetUnitTypes sets UnitTypes field to given value.

### HasUnitTypes

`func (o *ExternalMetricDefinitionContractDto) HasUnitTypes() bool`

HasUnitTypes returns a boolean if a field has been set.

### GetMetricEvents

`func (o *ExternalMetricDefinitionContractDto) GetMetricEvents() []ExternalMetricDefinitionContractDtoMetricEventsInner`

GetMetricEvents returns the MetricEvents field if non-nil, zero value otherwise.

### GetMetricEventsOk

`func (o *ExternalMetricDefinitionContractDto) GetMetricEventsOk() (*[]ExternalMetricDefinitionContractDtoMetricEventsInner, bool)`

GetMetricEventsOk returns a tuple with the MetricEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEvents

`func (o *ExternalMetricDefinitionContractDto) SetMetricEvents(v []ExternalMetricDefinitionContractDtoMetricEventsInner)`

SetMetricEvents sets MetricEvents field to given value.

### HasMetricEvents

`func (o *ExternalMetricDefinitionContractDto) HasMetricEvents() bool`

HasMetricEvents returns a boolean if a field has been set.

### GetMetricComponentMetrics

`func (o *ExternalMetricDefinitionContractDto) GetMetricComponentMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetMetricComponentMetrics returns the MetricComponentMetrics field if non-nil, zero value otherwise.

### GetMetricComponentMetricsOk

`func (o *ExternalMetricDefinitionContractDto) GetMetricComponentMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetMetricComponentMetricsOk returns a tuple with the MetricComponentMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricComponentMetrics

`func (o *ExternalMetricDefinitionContractDto) SetMetricComponentMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetMetricComponentMetrics sets MetricComponentMetrics field to given value.

### HasMetricComponentMetrics

`func (o *ExternalMetricDefinitionContractDto) HasMetricComponentMetrics() bool`

HasMetricComponentMetrics returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalMetricDefinitionContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalMetricDefinitionContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalMetricDefinitionContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalMetricDefinitionContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectionality

`func (o *ExternalMetricDefinitionContractDto) GetDirectionality() string`

GetDirectionality returns the Directionality field if non-nil, zero value otherwise.

### GetDirectionalityOk

`func (o *ExternalMetricDefinitionContractDto) GetDirectionalityOk() (*string, bool)`

GetDirectionalityOk returns a tuple with the Directionality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionality

`func (o *ExternalMetricDefinitionContractDto) SetDirectionality(v string)`

SetDirectionality sets Directionality field to given value.


### GetTags

`func (o *ExternalMetricDefinitionContractDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternalMetricDefinitionContractDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternalMetricDefinitionContractDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExternalMetricDefinitionContractDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsPermanent

`func (o *ExternalMetricDefinitionContractDto) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *ExternalMetricDefinitionContractDto) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *ExternalMetricDefinitionContractDto) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *ExternalMetricDefinitionContractDto) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetRollupTimeWindow

`func (o *ExternalMetricDefinitionContractDto) GetRollupTimeWindow() string`

GetRollupTimeWindow returns the RollupTimeWindow field if non-nil, zero value otherwise.

### GetRollupTimeWindowOk

`func (o *ExternalMetricDefinitionContractDto) GetRollupTimeWindowOk() (*string, bool)`

GetRollupTimeWindowOk returns a tuple with the RollupTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupTimeWindow

`func (o *ExternalMetricDefinitionContractDto) SetRollupTimeWindow(v string)`

SetRollupTimeWindow sets RollupTimeWindow field to given value.

### HasRollupTimeWindow

`func (o *ExternalMetricDefinitionContractDto) HasRollupTimeWindow() bool`

HasRollupTimeWindow returns a boolean if a field has been set.

### GetCustomRollUpStart

`func (o *ExternalMetricDefinitionContractDto) GetCustomRollUpStart() float32`

GetCustomRollUpStart returns the CustomRollUpStart field if non-nil, zero value otherwise.

### GetCustomRollUpStartOk

`func (o *ExternalMetricDefinitionContractDto) GetCustomRollUpStartOk() (*float32, bool)`

GetCustomRollUpStartOk returns a tuple with the CustomRollUpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRollUpStart

`func (o *ExternalMetricDefinitionContractDto) SetCustomRollUpStart(v float32)`

SetCustomRollUpStart sets CustomRollUpStart field to given value.

### HasCustomRollUpStart

`func (o *ExternalMetricDefinitionContractDto) HasCustomRollUpStart() bool`

HasCustomRollUpStart returns a boolean if a field has been set.

### GetCustomRollUpEnd

`func (o *ExternalMetricDefinitionContractDto) GetCustomRollUpEnd() float32`

GetCustomRollUpEnd returns the CustomRollUpEnd field if non-nil, zero value otherwise.

### GetCustomRollUpEndOk

`func (o *ExternalMetricDefinitionContractDto) GetCustomRollUpEndOk() (*float32, bool)`

GetCustomRollUpEndOk returns a tuple with the CustomRollUpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRollUpEnd

`func (o *ExternalMetricDefinitionContractDto) SetCustomRollUpEnd(v float32)`

SetCustomRollUpEnd sets CustomRollUpEnd field to given value.

### HasCustomRollUpEnd

`func (o *ExternalMetricDefinitionContractDto) HasCustomRollUpEnd() bool`

HasCustomRollUpEnd returns a boolean if a field has been set.

### GetFunnelEventList

`func (o *ExternalMetricDefinitionContractDto) GetFunnelEventList() []ExternalMetricDefinitionContractDtoFunnelEventListInner`

GetFunnelEventList returns the FunnelEventList field if non-nil, zero value otherwise.

### GetFunnelEventListOk

`func (o *ExternalMetricDefinitionContractDto) GetFunnelEventListOk() (*[]ExternalMetricDefinitionContractDtoFunnelEventListInner, bool)`

GetFunnelEventListOk returns a tuple with the FunnelEventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelEventList

`func (o *ExternalMetricDefinitionContractDto) SetFunnelEventList(v []ExternalMetricDefinitionContractDtoFunnelEventListInner)`

SetFunnelEventList sets FunnelEventList field to given value.

### HasFunnelEventList

`func (o *ExternalMetricDefinitionContractDto) HasFunnelEventList() bool`

HasFunnelEventList returns a boolean if a field has been set.

### GetFunnelCountDistinct

`func (o *ExternalMetricDefinitionContractDto) GetFunnelCountDistinct() string`

GetFunnelCountDistinct returns the FunnelCountDistinct field if non-nil, zero value otherwise.

### GetFunnelCountDistinctOk

`func (o *ExternalMetricDefinitionContractDto) GetFunnelCountDistinctOk() (*string, bool)`

GetFunnelCountDistinctOk returns a tuple with the FunnelCountDistinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCountDistinct

`func (o *ExternalMetricDefinitionContractDto) SetFunnelCountDistinct(v string)`

SetFunnelCountDistinct sets FunnelCountDistinct field to given value.

### HasFunnelCountDistinct

`func (o *ExternalMetricDefinitionContractDto) HasFunnelCountDistinct() bool`

HasFunnelCountDistinct returns a boolean if a field has been set.

### GetWarehouseNative

`func (o *ExternalMetricDefinitionContractDto) GetWarehouseNative() ExternalMetricDefinitionContractDtoWarehouseNative`

GetWarehouseNative returns the WarehouseNative field if non-nil, zero value otherwise.

### GetWarehouseNativeOk

`func (o *ExternalMetricDefinitionContractDto) GetWarehouseNativeOk() (*ExternalMetricDefinitionContractDtoWarehouseNative, bool)`

GetWarehouseNativeOk returns a tuple with the WarehouseNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseNative

`func (o *ExternalMetricDefinitionContractDto) SetWarehouseNative(v ExternalMetricDefinitionContractDtoWarehouseNative)`

SetWarehouseNative sets WarehouseNative field to given value.

### HasWarehouseNative

`func (o *ExternalMetricDefinitionContractDto) HasWarehouseNative() bool`

HasWarehouseNative returns a boolean if a field has been set.

### GetTeam

`func (o *ExternalMetricDefinitionContractDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExternalMetricDefinitionContractDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExternalMetricDefinitionContractDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExternalMetricDefinitionContractDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetId

`func (o *ExternalMetricDefinitionContractDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalMetricDefinitionContractDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalMetricDefinitionContractDto) SetId(v string)`

SetId sets Id field to given value.


### GetLineage

`func (o *ExternalMetricDefinitionContractDto) GetLineage() ExternalMetricDefinitionContractDtoLineage`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *ExternalMetricDefinitionContractDto) GetLineageOk() (*ExternalMetricDefinitionContractDtoLineage, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *ExternalMetricDefinitionContractDto) SetLineage(v ExternalMetricDefinitionContractDtoLineage)`

SetLineage sets Lineage field to given value.


### GetCreatorName

`func (o *ExternalMetricDefinitionContractDto) GetCreatorName() nil`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ExternalMetricDefinitionContractDto) GetCreatorNameOk() (*nil, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ExternalMetricDefinitionContractDto) SetCreatorName(v nil)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *ExternalMetricDefinitionContractDto) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *ExternalMetricDefinitionContractDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *ExternalMetricDefinitionContractDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *ExternalMetricDefinitionContractDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *ExternalMetricDefinitionContractDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ExternalMetricDefinitionContractDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ExternalMetricDefinitionContractDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ExternalMetricDefinitionContractDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ExternalMetricDefinitionContractDto) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetOwner

`func (o *ExternalMetricDefinitionContractDto) GetOwner() ExternalMetricDefinitionContractDtoOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExternalMetricDefinitionContractDto) GetOwnerOk() (*ExternalMetricDefinitionContractDtoOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExternalMetricDefinitionContractDto) SetOwner(v ExternalMetricDefinitionContractDtoOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExternalMetricDefinitionContractDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


