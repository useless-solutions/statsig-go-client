# LayerContractDtoParametersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this parameter, used for identification within the layer. | 
**Type** | **string** | The data type that this parameter returns. Allowed types include: string, boolean, number, object, and array. | 
**DefaultValue** | [**LayerContractDtoParametersInnerDefaultValue**](LayerContractDtoParametersInnerDefaultValue.md) |  | 

## Methods

### NewLayerContractDtoParametersInner

`func NewLayerContractDtoParametersInner(name string, type_ string, defaultValue LayerContractDtoParametersInnerDefaultValue, ) *LayerContractDtoParametersInner`

NewLayerContractDtoParametersInner instantiates a new LayerContractDtoParametersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerContractDtoParametersInnerWithDefaults

`func NewLayerContractDtoParametersInnerWithDefaults() *LayerContractDtoParametersInner`

NewLayerContractDtoParametersInnerWithDefaults instantiates a new LayerContractDtoParametersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LayerContractDtoParametersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayerContractDtoParametersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayerContractDtoParametersInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LayerContractDtoParametersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LayerContractDtoParametersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LayerContractDtoParametersInner) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultValue

`func (o *LayerContractDtoParametersInner) GetDefaultValue() LayerContractDtoParametersInnerDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *LayerContractDtoParametersInner) GetDefaultValueOk() (*LayerContractDtoParametersInnerDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *LayerContractDtoParametersInner) SetDefaultValue(v LayerContractDtoParametersInnerDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


