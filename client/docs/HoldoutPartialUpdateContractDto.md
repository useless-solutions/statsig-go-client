# HoldoutPartialUpdateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | enable or disable the holdout | [optional] 
**Description** | Pointer to **string** | brief summary of what the holdout is being used for | [optional] 
**PassPercentage** | Pointer to **float32** | percentage of users between 0-10% to pass through the holdout | [optional] 
**GateIDs** | Pointer to **[]string** | an array of gateIDs which this holdout is applied to | [optional] 
**ExperimentIDs** | Pointer to **[]string** | an array of experimentIDs which this holdout is applied to | [optional] 
**LayerIDs** | Pointer to **[]string** | an array of layerIDs which this holdout is applied to | [optional] 
**IsGlobal** | Pointer to **bool** | whether the holdout is being applied to all new gates | [optional] 
**TargetingGateID** | Pointer to [**nil**](nil.md) | the gateID that the holdout is targeting | [optional] 

## Methods

### NewHoldoutPartialUpdateContractDto

`func NewHoldoutPartialUpdateContractDto() *HoldoutPartialUpdateContractDto`

NewHoldoutPartialUpdateContractDto instantiates a new HoldoutPartialUpdateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutPartialUpdateContractDtoWithDefaults

`func NewHoldoutPartialUpdateContractDtoWithDefaults() *HoldoutPartialUpdateContractDto`

NewHoldoutPartialUpdateContractDtoWithDefaults instantiates a new HoldoutPartialUpdateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *HoldoutPartialUpdateContractDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *HoldoutPartialUpdateContractDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *HoldoutPartialUpdateContractDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *HoldoutPartialUpdateContractDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *HoldoutPartialUpdateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldoutPartialUpdateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldoutPartialUpdateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldoutPartialUpdateContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPassPercentage

`func (o *HoldoutPartialUpdateContractDto) GetPassPercentage() float32`

GetPassPercentage returns the PassPercentage field if non-nil, zero value otherwise.

### GetPassPercentageOk

`func (o *HoldoutPartialUpdateContractDto) GetPassPercentageOk() (*float32, bool)`

GetPassPercentageOk returns a tuple with the PassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercentage

`func (o *HoldoutPartialUpdateContractDto) SetPassPercentage(v float32)`

SetPassPercentage sets PassPercentage field to given value.

### HasPassPercentage

`func (o *HoldoutPartialUpdateContractDto) HasPassPercentage() bool`

HasPassPercentage returns a boolean if a field has been set.

### GetGateIDs

`func (o *HoldoutPartialUpdateContractDto) GetGateIDs() []string`

GetGateIDs returns the GateIDs field if non-nil, zero value otherwise.

### GetGateIDsOk

`func (o *HoldoutPartialUpdateContractDto) GetGateIDsOk() (*[]string, bool)`

GetGateIDsOk returns a tuple with the GateIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateIDs

`func (o *HoldoutPartialUpdateContractDto) SetGateIDs(v []string)`

SetGateIDs sets GateIDs field to given value.

### HasGateIDs

`func (o *HoldoutPartialUpdateContractDto) HasGateIDs() bool`

HasGateIDs returns a boolean if a field has been set.

### GetExperimentIDs

`func (o *HoldoutPartialUpdateContractDto) GetExperimentIDs() []string`

GetExperimentIDs returns the ExperimentIDs field if non-nil, zero value otherwise.

### GetExperimentIDsOk

`func (o *HoldoutPartialUpdateContractDto) GetExperimentIDsOk() (*[]string, bool)`

GetExperimentIDsOk returns a tuple with the ExperimentIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIDs

`func (o *HoldoutPartialUpdateContractDto) SetExperimentIDs(v []string)`

SetExperimentIDs sets ExperimentIDs field to given value.

### HasExperimentIDs

`func (o *HoldoutPartialUpdateContractDto) HasExperimentIDs() bool`

HasExperimentIDs returns a boolean if a field has been set.

### GetLayerIDs

`func (o *HoldoutPartialUpdateContractDto) GetLayerIDs() []string`

GetLayerIDs returns the LayerIDs field if non-nil, zero value otherwise.

### GetLayerIDsOk

`func (o *HoldoutPartialUpdateContractDto) GetLayerIDsOk() (*[]string, bool)`

GetLayerIDsOk returns a tuple with the LayerIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerIDs

`func (o *HoldoutPartialUpdateContractDto) SetLayerIDs(v []string)`

SetLayerIDs sets LayerIDs field to given value.

### HasLayerIDs

`func (o *HoldoutPartialUpdateContractDto) HasLayerIDs() bool`

HasLayerIDs returns a boolean if a field has been set.

### GetIsGlobal

`func (o *HoldoutPartialUpdateContractDto) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *HoldoutPartialUpdateContractDto) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *HoldoutPartialUpdateContractDto) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *HoldoutPartialUpdateContractDto) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### GetTargetingGateID

`func (o *HoldoutPartialUpdateContractDto) GetTargetingGateID() nil`

GetTargetingGateID returns the TargetingGateID field if non-nil, zero value otherwise.

### GetTargetingGateIDOk

`func (o *HoldoutPartialUpdateContractDto) GetTargetingGateIDOk() (*nil, bool)`

GetTargetingGateIDOk returns a tuple with the TargetingGateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingGateID

`func (o *HoldoutPartialUpdateContractDto) SetTargetingGateID(v nil)`

SetTargetingGateID sets TargetingGateID field to given value.

### HasTargetingGateID

`func (o *HoldoutPartialUpdateContractDto) HasTargetingGateID() bool`

HasTargetingGateID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


