# ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to [**[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner**](ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner.md) | Optional array of criteria to filter the funnel events, defined by various types and conditions. | [optional] 
**MetricSourceName** | Pointer to **string** | Optional name of the metric source associated with the funnel event. | [optional] 
**Name** | Pointer to [**nil**](nil.md) | Optional step name for the funnel event, can be null if not specified. | [optional] 
**SessionIdentifierField** | Pointer to [**nil**](nil.md) | Name of column which being used as session identifier. Funnel event with the same metric source | [optional] 

## Methods

### NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner

`func NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner() *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner`

NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner instantiates a new ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInnerWithDefaults

`func NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInnerWithDefaults() *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner`

NewExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInnerWithDefaults instantiates a new ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetCriteria() []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetCriteriaOk() (*[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetCriteria(v []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetMetricSourceName

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetMetricSourceName() string`

GetMetricSourceName returns the MetricSourceName field if non-nil, zero value otherwise.

### GetMetricSourceNameOk

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetMetricSourceNameOk() (*string, bool)`

GetMetricSourceNameOk returns a tuple with the MetricSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSourceName

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetMetricSourceName(v string)`

SetMetricSourceName sets MetricSourceName field to given value.

### HasMetricSourceName

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasMetricSourceName() bool`

HasMetricSourceName returns a boolean if a field has been set.

### GetName

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetName() nil`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetNameOk() (*nil, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetName(v nil)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSessionIdentifierField

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetSessionIdentifierField() nil`

GetSessionIdentifierField returns the SessionIdentifierField field if non-nil, zero value otherwise.

### GetSessionIdentifierFieldOk

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) GetSessionIdentifierFieldOk() (*nil, bool)`

GetSessionIdentifierFieldOk returns a tuple with the SessionIdentifierField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdentifierField

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) SetSessionIdentifierField(v nil)`

SetSessionIdentifierField sets SessionIdentifierField field to given value.

### HasSessionIdentifierField

`func (o *ExternalMetricDefinitionContractDtoWarehouseNativeFunnelEventsInner) HasSessionIdentifierField() bool`

HasSessionIdentifierField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


