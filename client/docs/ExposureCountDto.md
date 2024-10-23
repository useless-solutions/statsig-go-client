# ExposureCountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gates** | [**[]ExposureCountDtoGatesInner**](ExposureCountDtoGatesInner.md) | ids of gates to query (max 25) | 
**Experiments** | [**[]ExposureCountDtoGatesInner**](ExposureCountDtoGatesInner.md) | ids of experiments to query (max 25) | 

## Methods

### NewExposureCountDto

`func NewExposureCountDto(gates []ExposureCountDtoGatesInner, experiments []ExposureCountDtoGatesInner, ) *ExposureCountDto`

NewExposureCountDto instantiates a new ExposureCountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExposureCountDtoWithDefaults

`func NewExposureCountDtoWithDefaults() *ExposureCountDto`

NewExposureCountDtoWithDefaults instantiates a new ExposureCountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGates

`func (o *ExposureCountDto) GetGates() []ExposureCountDtoGatesInner`

GetGates returns the Gates field if non-nil, zero value otherwise.

### GetGatesOk

`func (o *ExposureCountDto) GetGatesOk() (*[]ExposureCountDtoGatesInner, bool)`

GetGatesOk returns a tuple with the Gates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGates

`func (o *ExposureCountDto) SetGates(v []ExposureCountDtoGatesInner)`

SetGates sets Gates field to given value.


### GetExperiments

`func (o *ExposureCountDto) GetExperiments() []ExposureCountDtoGatesInner`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *ExposureCountDto) GetExperimentsOk() (*[]ExposureCountDtoGatesInner, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *ExposureCountDto) SetExperiments(v []ExposureCountDtoGatesInner)`

SetExperiments sets Experiments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


