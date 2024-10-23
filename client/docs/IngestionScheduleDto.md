# IngestionScheduleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**ScheduledHourPst** | **float32** |  | 

## Methods

### NewIngestionScheduleDto

`func NewIngestionScheduleDto(dataset string, scheduledHourPst float32, ) *IngestionScheduleDto`

NewIngestionScheduleDto instantiates a new IngestionScheduleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionScheduleDtoWithDefaults

`func NewIngestionScheduleDtoWithDefaults() *IngestionScheduleDto`

NewIngestionScheduleDtoWithDefaults instantiates a new IngestionScheduleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *IngestionScheduleDto) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionScheduleDto) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionScheduleDto) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetScheduledHourPst

`func (o *IngestionScheduleDto) GetScheduledHourPst() float32`

GetScheduledHourPst returns the ScheduledHourPst field if non-nil, zero value otherwise.

### GetScheduledHourPstOk

`func (o *IngestionScheduleDto) GetScheduledHourPstOk() (*float32, bool)`

GetScheduledHourPstOk returns a tuple with the ScheduledHourPst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledHourPst

`func (o *IngestionScheduleDto) SetScheduledHourPst(v float32)`

SetScheduledHourPst sets ScheduledHourPst field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


