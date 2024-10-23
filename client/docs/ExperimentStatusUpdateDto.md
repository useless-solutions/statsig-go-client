# ExperimentStatusUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the group to launch | 
**DecisionReason** | **string** | The reason for making the decision to update the experiment status | 
**RemoveTargeting** | Pointer to **bool** | Indicates whether to remove targeting from the experiment | [optional] [default to false]

## Methods

### NewExperimentStatusUpdateDto

`func NewExperimentStatusUpdateDto(id string, decisionReason string, ) *ExperimentStatusUpdateDto`

NewExperimentStatusUpdateDto instantiates a new ExperimentStatusUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentStatusUpdateDtoWithDefaults

`func NewExperimentStatusUpdateDtoWithDefaults() *ExperimentStatusUpdateDto`

NewExperimentStatusUpdateDtoWithDefaults instantiates a new ExperimentStatusUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExperimentStatusUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentStatusUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentStatusUpdateDto) SetId(v string)`

SetId sets Id field to given value.


### GetDecisionReason

`func (o *ExperimentStatusUpdateDto) GetDecisionReason() string`

GetDecisionReason returns the DecisionReason field if non-nil, zero value otherwise.

### GetDecisionReasonOk

`func (o *ExperimentStatusUpdateDto) GetDecisionReasonOk() (*string, bool)`

GetDecisionReasonOk returns a tuple with the DecisionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionReason

`func (o *ExperimentStatusUpdateDto) SetDecisionReason(v string)`

SetDecisionReason sets DecisionReason field to given value.


### GetRemoveTargeting

`func (o *ExperimentStatusUpdateDto) GetRemoveTargeting() bool`

GetRemoveTargeting returns the RemoveTargeting field if non-nil, zero value otherwise.

### GetRemoveTargetingOk

`func (o *ExperimentStatusUpdateDto) GetRemoveTargetingOk() (*bool, bool)`

GetRemoveTargetingOk returns a tuple with the RemoveTargeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTargeting

`func (o *ExperimentStatusUpdateDto) SetRemoveTargeting(v bool)`

SetRemoveTargeting sets RemoveTargeting field to given value.

### HasRemoveTargeting

`func (o *ExperimentStatusUpdateDto) HasRemoveTargeting() bool`

HasRemoveTargeting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


