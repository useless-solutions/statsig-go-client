# ExternalMetricDefinitionContractDtoMetricEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the metric event. | 
**Type** | Pointer to [**nil**](nil.md) | The type of metric event. Allowed values include: count, count_distinct, value, and metadata. | [optional] 
**MetadataKey** | Pointer to **string** | The key for associated metadata, if applicable. | [optional] 
**Criteria** | Pointer to [**[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner**](ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner.md) | Filtering criteria for the metric event, including conditions and values to refine the event data. | [optional] 

## Methods

### NewExternalMetricDefinitionContractDtoMetricEventsInner

`func NewExternalMetricDefinitionContractDtoMetricEventsInner(name string, ) *ExternalMetricDefinitionContractDtoMetricEventsInner`

NewExternalMetricDefinitionContractDtoMetricEventsInner instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMetricDefinitionContractDtoMetricEventsInnerWithDefaults

`func NewExternalMetricDefinitionContractDtoMetricEventsInnerWithDefaults() *ExternalMetricDefinitionContractDtoMetricEventsInner`

NewExternalMetricDefinitionContractDtoMetricEventsInnerWithDefaults instantiates a new ExternalMetricDefinitionContractDtoMetricEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetType() nil`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetTypeOk() (*nil, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetType(v nil)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadataKey

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetMetadataKey() string`

GetMetadataKey returns the MetadataKey field if non-nil, zero value otherwise.

### GetMetadataKeyOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetMetadataKeyOk() (*string, bool)`

GetMetadataKeyOk returns a tuple with the MetadataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataKey

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetMetadataKey(v string)`

SetMetadataKey sets MetadataKey field to given value.

### HasMetadataKey

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) HasMetadataKey() bool`

HasMetadataKey returns a boolean if a field has been set.

### GetCriteria

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetCriteria() []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) GetCriteriaOk() (*[]ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) SetCriteria(v []ExternalMetricDefinitionContractDtoMetricEventsInnerCriteriaInner)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ExternalMetricDefinitionContractDtoMetricEventsInner) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


