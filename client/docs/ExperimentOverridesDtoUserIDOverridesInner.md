# ExperimentOverridesDtoUserIDOverridesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupID** | **string** | The experiment group the user will be forced into. | 
**Ids** | **[]string** | Array of user IDs to be assigned to the specified experiment group. | 
**Environment** | Pointer to [**nil**](nil.md) | Optional environment designation (e.g., production, staging) for the experiment. | [optional] 
**UnitType** | Pointer to [**nil**](nil.md) | Optional type of unit for the experiment, defining the scope of the override. | [optional] 

## Methods

### NewExperimentOverridesDtoUserIDOverridesInner

`func NewExperimentOverridesDtoUserIDOverridesInner(groupID string, ids []string, ) *ExperimentOverridesDtoUserIDOverridesInner`

NewExperimentOverridesDtoUserIDOverridesInner instantiates a new ExperimentOverridesDtoUserIDOverridesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentOverridesDtoUserIDOverridesInnerWithDefaults

`func NewExperimentOverridesDtoUserIDOverridesInnerWithDefaults() *ExperimentOverridesDtoUserIDOverridesInner`

NewExperimentOverridesDtoUserIDOverridesInnerWithDefaults instantiates a new ExperimentOverridesDtoUserIDOverridesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupID

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetGroupID() string`

GetGroupID returns the GroupID field if non-nil, zero value otherwise.

### GetGroupIDOk

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetGroupIDOk() (*string, bool)`

GetGroupIDOk returns a tuple with the GroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupID

`func (o *ExperimentOverridesDtoUserIDOverridesInner) SetGroupID(v string)`

SetGroupID sets GroupID field to given value.


### GetIds

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ExperimentOverridesDtoUserIDOverridesInner) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetEnvironment

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetEnvironment() nil`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetEnvironmentOk() (*nil, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ExperimentOverridesDtoUserIDOverridesInner) SetEnvironment(v nil)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ExperimentOverridesDtoUserIDOverridesInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetUnitType

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetUnitType() nil`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *ExperimentOverridesDtoUserIDOverridesInner) GetUnitTypeOk() (*nil, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *ExperimentOverridesDtoUserIDOverridesInner) SetUnitType(v nil)`

SetUnitType sets UnitType field to given value.

### HasUnitType

`func (o *ExperimentOverridesDtoUserIDOverridesInner) HasUnitType() bool`

HasUnitType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


