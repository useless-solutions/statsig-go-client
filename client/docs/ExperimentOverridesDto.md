# ExperimentOverridesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | [**[]ExperimentOverridesDtoOverridesInner**](ExperimentOverridesDtoOverridesInner.md) | Array of experiment overrides, each specifying type, ID, and group ID. | 
**UserIDOverrides** | [**[]ExperimentOverridesDtoUserIDOverridesInner**](ExperimentOverridesDtoUserIDOverridesInner.md) | Array of user ID overrides, specifying which users to force into experiment groups. | 

## Methods

### NewExperimentOverridesDto

`func NewExperimentOverridesDto(overrides []ExperimentOverridesDtoOverridesInner, userIDOverrides []ExperimentOverridesDtoUserIDOverridesInner, ) *ExperimentOverridesDto`

NewExperimentOverridesDto instantiates a new ExperimentOverridesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentOverridesDtoWithDefaults

`func NewExperimentOverridesDtoWithDefaults() *ExperimentOverridesDto`

NewExperimentOverridesDtoWithDefaults instantiates a new ExperimentOverridesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *ExperimentOverridesDto) GetOverrides() []ExperimentOverridesDtoOverridesInner`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ExperimentOverridesDto) GetOverridesOk() (*[]ExperimentOverridesDtoOverridesInner, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ExperimentOverridesDto) SetOverrides(v []ExperimentOverridesDtoOverridesInner)`

SetOverrides sets Overrides field to given value.


### GetUserIDOverrides

`func (o *ExperimentOverridesDto) GetUserIDOverrides() []ExperimentOverridesDtoUserIDOverridesInner`

GetUserIDOverrides returns the UserIDOverrides field if non-nil, zero value otherwise.

### GetUserIDOverridesOk

`func (o *ExperimentOverridesDto) GetUserIDOverridesOk() (*[]ExperimentOverridesDtoUserIDOverridesInner, bool)`

GetUserIDOverridesOk returns a tuple with the UserIDOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIDOverrides

`func (o *ExperimentOverridesDto) SetUserIDOverrides(v []ExperimentOverridesDtoUserIDOverridesInner)`

SetUserIDOverrides sets UserIDOverrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


