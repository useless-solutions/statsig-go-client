# HoldoutFullUpdateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | enable or disable the holdout | 
**Description** | **string** | brief summary of what the holdout is being used for | 
**PassPercentage** | **float32** | percentage of users between 0-10% to pass through the holdout | 
**GateIDs** | **[]string** | an array of gateIDs which this holdout is applied to | 
**ExperimentIDs** | **[]string** | an array of experimentIDs which this holdout is applied to | 
**LayerIDs** | **[]string** | an array of layerIDs which this holdout is applied to | 
**IsGlobal** | **bool** | whether the holdout is being applied to all new gates | 
**TargetingGateID** | [**nil**](nil.md) | the gateID that the holdout is targeting | 

## Methods

### NewHoldoutFullUpdateContractDto

`func NewHoldoutFullUpdateContractDto(isEnabled bool, description string, passPercentage float32, gateIDs []string, experimentIDs []string, layerIDs []string, isGlobal bool, targetingGateID nil, ) *HoldoutFullUpdateContractDto`

NewHoldoutFullUpdateContractDto instantiates a new HoldoutFullUpdateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutFullUpdateContractDtoWithDefaults

`func NewHoldoutFullUpdateContractDtoWithDefaults() *HoldoutFullUpdateContractDto`

NewHoldoutFullUpdateContractDtoWithDefaults instantiates a new HoldoutFullUpdateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *HoldoutFullUpdateContractDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *HoldoutFullUpdateContractDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *HoldoutFullUpdateContractDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDescription

`func (o *HoldoutFullUpdateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldoutFullUpdateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldoutFullUpdateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPassPercentage

`func (o *HoldoutFullUpdateContractDto) GetPassPercentage() float32`

GetPassPercentage returns the PassPercentage field if non-nil, zero value otherwise.

### GetPassPercentageOk

`func (o *HoldoutFullUpdateContractDto) GetPassPercentageOk() (*float32, bool)`

GetPassPercentageOk returns a tuple with the PassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercentage

`func (o *HoldoutFullUpdateContractDto) SetPassPercentage(v float32)`

SetPassPercentage sets PassPercentage field to given value.


### GetGateIDs

`func (o *HoldoutFullUpdateContractDto) GetGateIDs() []string`

GetGateIDs returns the GateIDs field if non-nil, zero value otherwise.

### GetGateIDsOk

`func (o *HoldoutFullUpdateContractDto) GetGateIDsOk() (*[]string, bool)`

GetGateIDsOk returns a tuple with the GateIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateIDs

`func (o *HoldoutFullUpdateContractDto) SetGateIDs(v []string)`

SetGateIDs sets GateIDs field to given value.


### GetExperimentIDs

`func (o *HoldoutFullUpdateContractDto) GetExperimentIDs() []string`

GetExperimentIDs returns the ExperimentIDs field if non-nil, zero value otherwise.

### GetExperimentIDsOk

`func (o *HoldoutFullUpdateContractDto) GetExperimentIDsOk() (*[]string, bool)`

GetExperimentIDsOk returns a tuple with the ExperimentIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIDs

`func (o *HoldoutFullUpdateContractDto) SetExperimentIDs(v []string)`

SetExperimentIDs sets ExperimentIDs field to given value.


### GetLayerIDs

`func (o *HoldoutFullUpdateContractDto) GetLayerIDs() []string`

GetLayerIDs returns the LayerIDs field if non-nil, zero value otherwise.

### GetLayerIDsOk

`func (o *HoldoutFullUpdateContractDto) GetLayerIDsOk() (*[]string, bool)`

GetLayerIDsOk returns a tuple with the LayerIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerIDs

`func (o *HoldoutFullUpdateContractDto) SetLayerIDs(v []string)`

SetLayerIDs sets LayerIDs field to given value.


### GetIsGlobal

`func (o *HoldoutFullUpdateContractDto) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *HoldoutFullUpdateContractDto) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *HoldoutFullUpdateContractDto) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.


### GetTargetingGateID

`func (o *HoldoutFullUpdateContractDto) GetTargetingGateID() nil`

GetTargetingGateID returns the TargetingGateID field if non-nil, zero value otherwise.

### GetTargetingGateIDOk

`func (o *HoldoutFullUpdateContractDto) GetTargetingGateIDOk() (*nil, bool)`

GetTargetingGateIDOk returns a tuple with the TargetingGateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingGateID

`func (o *HoldoutFullUpdateContractDto) SetTargetingGateID(v nil)`

SetTargetingGateID sets TargetingGateID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


