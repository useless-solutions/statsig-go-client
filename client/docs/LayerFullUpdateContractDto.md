# LayerFullUpdateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A detailed description of the layer, explaining its purpose and functionality. | 
**Parameters** | [**[]LayerContractDtoParametersInner**](LayerContractDtoParametersInner.md) | An array of parameters associated with the layer, each defining specific attributes. | 
**TargetApps** | Pointer to [**LayerContractDtoTargetApps**](LayerContractDtoTargetApps.md) |  | [optional] 

## Methods

### NewLayerFullUpdateContractDto

`func NewLayerFullUpdateContractDto(description string, parameters []LayerContractDtoParametersInner, ) *LayerFullUpdateContractDto`

NewLayerFullUpdateContractDto instantiates a new LayerFullUpdateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerFullUpdateContractDtoWithDefaults

`func NewLayerFullUpdateContractDtoWithDefaults() *LayerFullUpdateContractDto`

NewLayerFullUpdateContractDtoWithDefaults instantiates a new LayerFullUpdateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LayerFullUpdateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LayerFullUpdateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LayerFullUpdateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetParameters

`func (o *LayerFullUpdateContractDto) GetParameters() []LayerContractDtoParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LayerFullUpdateContractDto) GetParametersOk() (*[]LayerContractDtoParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LayerFullUpdateContractDto) SetParameters(v []LayerContractDtoParametersInner)`

SetParameters sets Parameters field to given value.


### GetTargetApps

`func (o *LayerFullUpdateContractDto) GetTargetApps() LayerContractDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *LayerFullUpdateContractDto) GetTargetAppsOk() (*LayerContractDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *LayerFullUpdateContractDto) SetTargetApps(v LayerContractDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *LayerFullUpdateContractDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


