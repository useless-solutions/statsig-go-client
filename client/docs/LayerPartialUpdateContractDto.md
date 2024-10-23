# LayerPartialUpdateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A detailed description of the layer, explaining its purpose and functionality. | [optional] 
**Parameters** | Pointer to [**[]LayerContractDtoParametersInner**](LayerContractDtoParametersInner.md) | An array of parameters associated with the layer, each defining specific attributes. | [optional] 
**TargetApps** | Pointer to [**LayerContractDtoTargetApps**](LayerContractDtoTargetApps.md) |  | [optional] 

## Methods

### NewLayerPartialUpdateContractDto

`func NewLayerPartialUpdateContractDto() *LayerPartialUpdateContractDto`

NewLayerPartialUpdateContractDto instantiates a new LayerPartialUpdateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerPartialUpdateContractDtoWithDefaults

`func NewLayerPartialUpdateContractDtoWithDefaults() *LayerPartialUpdateContractDto`

NewLayerPartialUpdateContractDtoWithDefaults instantiates a new LayerPartialUpdateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LayerPartialUpdateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LayerPartialUpdateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LayerPartialUpdateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LayerPartialUpdateContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *LayerPartialUpdateContractDto) GetParameters() []LayerContractDtoParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LayerPartialUpdateContractDto) GetParametersOk() (*[]LayerContractDtoParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LayerPartialUpdateContractDto) SetParameters(v []LayerContractDtoParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *LayerPartialUpdateContractDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTargetApps

`func (o *LayerPartialUpdateContractDto) GetTargetApps() LayerContractDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *LayerPartialUpdateContractDto) GetTargetAppsOk() (*LayerContractDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *LayerPartialUpdateContractDto) SetTargetApps(v LayerContractDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *LayerPartialUpdateContractDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


