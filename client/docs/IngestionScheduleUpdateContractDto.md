# IngestionScheduleUpdateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**ScheduledHourPst** | Pointer to **float32** |  | [optional] [default to 10]

## Methods

### NewIngestionScheduleUpdateContractDto

`func NewIngestionScheduleUpdateContractDto(dataset string, ) *IngestionScheduleUpdateContractDto`

NewIngestionScheduleUpdateContractDto instantiates a new IngestionScheduleUpdateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionScheduleUpdateContractDtoWithDefaults

`func NewIngestionScheduleUpdateContractDtoWithDefaults() *IngestionScheduleUpdateContractDto`

NewIngestionScheduleUpdateContractDtoWithDefaults instantiates a new IngestionScheduleUpdateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *IngestionScheduleUpdateContractDto) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionScheduleUpdateContractDto) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionScheduleUpdateContractDto) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetScheduledHourPst

`func (o *IngestionScheduleUpdateContractDto) GetScheduledHourPst() float32`

GetScheduledHourPst returns the ScheduledHourPst field if non-nil, zero value otherwise.

### GetScheduledHourPstOk

`func (o *IngestionScheduleUpdateContractDto) GetScheduledHourPstOk() (*float32, bool)`

GetScheduledHourPstOk returns a tuple with the ScheduledHourPst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledHourPst

`func (o *IngestionScheduleUpdateContractDto) SetScheduledHourPst(v float32)`

SetScheduledHourPst sets ScheduledHourPst field to given value.

### HasScheduledHourPst

`func (o *IngestionScheduleUpdateContractDto) HasScheduledHourPst() bool`

HasScheduledHourPst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


