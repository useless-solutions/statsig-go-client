# UpdateOverridesContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentOverrides** | [**[]OverrideDtoEnvironmentOverridesInner**](OverrideDtoEnvironmentOverridesInner.md) |  | 
**PassingUserIDs** | **[]string** | List of user IDs | 
**FailingUserIDs** | **[]string** | List of user IDs | 
**PassingCustomIDs** | Pointer to **[]string** | Optional list of custom IDs | [optional] 
**FailingCustomIDs** | Pointer to **[]string** | Optional list of custom IDs | [optional] 

## Methods

### NewUpdateOverridesContractDto

`func NewUpdateOverridesContractDto(environmentOverrides []OverrideDtoEnvironmentOverridesInner, passingUserIDs []string, failingUserIDs []string, ) *UpdateOverridesContractDto`

NewUpdateOverridesContractDto instantiates a new UpdateOverridesContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOverridesContractDtoWithDefaults

`func NewUpdateOverridesContractDtoWithDefaults() *UpdateOverridesContractDto`

NewUpdateOverridesContractDtoWithDefaults instantiates a new UpdateOverridesContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentOverrides

`func (o *UpdateOverridesContractDto) GetEnvironmentOverrides() []OverrideDtoEnvironmentOverridesInner`

GetEnvironmentOverrides returns the EnvironmentOverrides field if non-nil, zero value otherwise.

### GetEnvironmentOverridesOk

`func (o *UpdateOverridesContractDto) GetEnvironmentOverridesOk() (*[]OverrideDtoEnvironmentOverridesInner, bool)`

GetEnvironmentOverridesOk returns a tuple with the EnvironmentOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentOverrides

`func (o *UpdateOverridesContractDto) SetEnvironmentOverrides(v []OverrideDtoEnvironmentOverridesInner)`

SetEnvironmentOverrides sets EnvironmentOverrides field to given value.


### GetPassingUserIDs

`func (o *UpdateOverridesContractDto) GetPassingUserIDs() []string`

GetPassingUserIDs returns the PassingUserIDs field if non-nil, zero value otherwise.

### GetPassingUserIDsOk

`func (o *UpdateOverridesContractDto) GetPassingUserIDsOk() (*[]string, bool)`

GetPassingUserIDsOk returns a tuple with the PassingUserIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassingUserIDs

`func (o *UpdateOverridesContractDto) SetPassingUserIDs(v []string)`

SetPassingUserIDs sets PassingUserIDs field to given value.


### GetFailingUserIDs

`func (o *UpdateOverridesContractDto) GetFailingUserIDs() []string`

GetFailingUserIDs returns the FailingUserIDs field if non-nil, zero value otherwise.

### GetFailingUserIDsOk

`func (o *UpdateOverridesContractDto) GetFailingUserIDsOk() (*[]string, bool)`

GetFailingUserIDsOk returns a tuple with the FailingUserIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingUserIDs

`func (o *UpdateOverridesContractDto) SetFailingUserIDs(v []string)`

SetFailingUserIDs sets FailingUserIDs field to given value.


### GetPassingCustomIDs

`func (o *UpdateOverridesContractDto) GetPassingCustomIDs() []string`

GetPassingCustomIDs returns the PassingCustomIDs field if non-nil, zero value otherwise.

### GetPassingCustomIDsOk

`func (o *UpdateOverridesContractDto) GetPassingCustomIDsOk() (*[]string, bool)`

GetPassingCustomIDsOk returns a tuple with the PassingCustomIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassingCustomIDs

`func (o *UpdateOverridesContractDto) SetPassingCustomIDs(v []string)`

SetPassingCustomIDs sets PassingCustomIDs field to given value.

### HasPassingCustomIDs

`func (o *UpdateOverridesContractDto) HasPassingCustomIDs() bool`

HasPassingCustomIDs returns a boolean if a field has been set.

### GetFailingCustomIDs

`func (o *UpdateOverridesContractDto) GetFailingCustomIDs() []string`

GetFailingCustomIDs returns the FailingCustomIDs field if non-nil, zero value otherwise.

### GetFailingCustomIDsOk

`func (o *UpdateOverridesContractDto) GetFailingCustomIDsOk() (*[]string, bool)`

GetFailingCustomIDsOk returns a tuple with the FailingCustomIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingCustomIDs

`func (o *UpdateOverridesContractDto) SetFailingCustomIDs(v []string)`

SetFailingCustomIDs sets FailingCustomIDs field to given value.

### HasFailingCustomIDs

`func (o *UpdateOverridesContractDto) HasFailingCustomIDs() bool`

HasFailingCustomIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


