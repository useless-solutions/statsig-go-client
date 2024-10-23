# ExperimentArchiveDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveReason** | Pointer to **string** | The reason for archiving the experiment | [optional] 

## Methods

### NewExperimentArchiveDto

`func NewExperimentArchiveDto() *ExperimentArchiveDto`

NewExperimentArchiveDto instantiates a new ExperimentArchiveDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentArchiveDtoWithDefaults

`func NewExperimentArchiveDtoWithDefaults() *ExperimentArchiveDto`

NewExperimentArchiveDtoWithDefaults instantiates a new ExperimentArchiveDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveReason

`func (o *ExperimentArchiveDto) GetArchiveReason() string`

GetArchiveReason returns the ArchiveReason field if non-nil, zero value otherwise.

### GetArchiveReasonOk

`func (o *ExperimentArchiveDto) GetArchiveReasonOk() (*string, bool)`

GetArchiveReasonOk returns a tuple with the ArchiveReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveReason

`func (o *ExperimentArchiveDto) SetArchiveReason(v string)`

SetArchiveReason sets ArchiveReason field to given value.

### HasArchiveReason

`func (o *ExperimentArchiveDto) HasArchiveReason() bool`

HasArchiveReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


