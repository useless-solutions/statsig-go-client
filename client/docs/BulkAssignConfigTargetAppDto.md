# BulkAssignConfigTargetAppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetApps** | **[]string** | target app ids | 
**Gates** | Pointer to **[]string** | Gate IDs to assign to target app(s) | [optional] 
**DynamicConfigs** | Pointer to **[]string** | Dynamic Config IDs to assign to target app(s) | [optional] 
**Experiments** | Pointer to **[]string** | Experiment IDs to assign to target app(s) | [optional] 

## Methods

### NewBulkAssignConfigTargetAppDto

`func NewBulkAssignConfigTargetAppDto(targetApps []string, ) *BulkAssignConfigTargetAppDto`

NewBulkAssignConfigTargetAppDto instantiates a new BulkAssignConfigTargetAppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAssignConfigTargetAppDtoWithDefaults

`func NewBulkAssignConfigTargetAppDtoWithDefaults() *BulkAssignConfigTargetAppDto`

NewBulkAssignConfigTargetAppDtoWithDefaults instantiates a new BulkAssignConfigTargetAppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetApps

`func (o *BulkAssignConfigTargetAppDto) GetTargetApps() []string`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *BulkAssignConfigTargetAppDto) GetTargetAppsOk() (*[]string, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *BulkAssignConfigTargetAppDto) SetTargetApps(v []string)`

SetTargetApps sets TargetApps field to given value.


### GetGates

`func (o *BulkAssignConfigTargetAppDto) GetGates() []string`

GetGates returns the Gates field if non-nil, zero value otherwise.

### GetGatesOk

`func (o *BulkAssignConfigTargetAppDto) GetGatesOk() (*[]string, bool)`

GetGatesOk returns a tuple with the Gates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGates

`func (o *BulkAssignConfigTargetAppDto) SetGates(v []string)`

SetGates sets Gates field to given value.

### HasGates

`func (o *BulkAssignConfigTargetAppDto) HasGates() bool`

HasGates returns a boolean if a field has been set.

### GetDynamicConfigs

`func (o *BulkAssignConfigTargetAppDto) GetDynamicConfigs() []string`

GetDynamicConfigs returns the DynamicConfigs field if non-nil, zero value otherwise.

### GetDynamicConfigsOk

`func (o *BulkAssignConfigTargetAppDto) GetDynamicConfigsOk() (*[]string, bool)`

GetDynamicConfigsOk returns a tuple with the DynamicConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicConfigs

`func (o *BulkAssignConfigTargetAppDto) SetDynamicConfigs(v []string)`

SetDynamicConfigs sets DynamicConfigs field to given value.

### HasDynamicConfigs

`func (o *BulkAssignConfigTargetAppDto) HasDynamicConfigs() bool`

HasDynamicConfigs returns a boolean if a field has been set.

### GetExperiments

`func (o *BulkAssignConfigTargetAppDto) GetExperiments() []string`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *BulkAssignConfigTargetAppDto) GetExperimentsOk() (*[]string, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *BulkAssignConfigTargetAppDto) SetExperiments(v []string)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *BulkAssignConfigTargetAppDto) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


