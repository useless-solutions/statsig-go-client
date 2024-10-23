# TargetAppCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the target app | 
**Description** | **string** | a description of the target app | 
**Gates** | Pointer to **[]string** | Gate IDs to assign to target app(s) | [optional] 
**DynamicConfigs** | Pointer to **[]string** | Dynamic Config IDs to assign to target app(s) | [optional] 
**Experiments** | Pointer to **[]string** | Experiment IDs to assign to target app(s) | [optional] 

## Methods

### NewTargetAppCreateDto

`func NewTargetAppCreateDto(name string, description string, ) *TargetAppCreateDto`

NewTargetAppCreateDto instantiates a new TargetAppCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetAppCreateDtoWithDefaults

`func NewTargetAppCreateDtoWithDefaults() *TargetAppCreateDto`

NewTargetAppCreateDtoWithDefaults instantiates a new TargetAppCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TargetAppCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetAppCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetAppCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TargetAppCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetAppCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetAppCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGates

`func (o *TargetAppCreateDto) GetGates() []string`

GetGates returns the Gates field if non-nil, zero value otherwise.

### GetGatesOk

`func (o *TargetAppCreateDto) GetGatesOk() (*[]string, bool)`

GetGatesOk returns a tuple with the Gates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGates

`func (o *TargetAppCreateDto) SetGates(v []string)`

SetGates sets Gates field to given value.

### HasGates

`func (o *TargetAppCreateDto) HasGates() bool`

HasGates returns a boolean if a field has been set.

### GetDynamicConfigs

`func (o *TargetAppCreateDto) GetDynamicConfigs() []string`

GetDynamicConfigs returns the DynamicConfigs field if non-nil, zero value otherwise.

### GetDynamicConfigsOk

`func (o *TargetAppCreateDto) GetDynamicConfigsOk() (*[]string, bool)`

GetDynamicConfigsOk returns a tuple with the DynamicConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicConfigs

`func (o *TargetAppCreateDto) SetDynamicConfigs(v []string)`

SetDynamicConfigs sets DynamicConfigs field to given value.

### HasDynamicConfigs

`func (o *TargetAppCreateDto) HasDynamicConfigs() bool`

HasDynamicConfigs returns a boolean if a field has been set.

### GetExperiments

`func (o *TargetAppCreateDto) GetExperiments() []string`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *TargetAppCreateDto) GetExperimentsOk() (*[]string, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *TargetAppCreateDto) SetExperiments(v []string)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *TargetAppCreateDto) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


