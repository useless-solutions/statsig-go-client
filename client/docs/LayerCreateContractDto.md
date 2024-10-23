# LayerCreateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the layer, formatted as an ID (e.g., \&quot;A Layer\&quot; becomes \&quot;a_layer\&quot;). | 
**Description** | Pointer to **string** | A helpful description of what this layer does, providing context for its purpose. | [optional] 
**IdType** | **string** | The type of ID used for this layer, essential for validation in services. | 
**TargetApps** | Pointer to [**LayerCreateContractDtoTargetApps**](LayerCreateContractDtoTargetApps.md) |  | [optional] 
**Team** | Pointer to **string** | Optional identifier for the team responsible for this layer. | [optional] 

## Methods

### NewLayerCreateContractDto

`func NewLayerCreateContractDto(name string, idType string, ) *LayerCreateContractDto`

NewLayerCreateContractDto instantiates a new LayerCreateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerCreateContractDtoWithDefaults

`func NewLayerCreateContractDtoWithDefaults() *LayerCreateContractDto`

NewLayerCreateContractDtoWithDefaults instantiates a new LayerCreateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LayerCreateContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayerCreateContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayerCreateContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LayerCreateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LayerCreateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LayerCreateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LayerCreateContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdType

`func (o *LayerCreateContractDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *LayerCreateContractDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *LayerCreateContractDto) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetTargetApps

`func (o *LayerCreateContractDto) GetTargetApps() LayerCreateContractDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *LayerCreateContractDto) GetTargetAppsOk() (*LayerCreateContractDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *LayerCreateContractDto) SetTargetApps(v LayerCreateContractDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *LayerCreateContractDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetTeam

`func (o *LayerCreateContractDto) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *LayerCreateContractDto) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *LayerCreateContractDto) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *LayerCreateContractDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


