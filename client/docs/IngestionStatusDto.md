# IngestionStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ds** | Pointer to **time.Time** |  | [optional] 
**IngestionDataset** | [**nil**](nil.md) |  | 
**IngestionSource** | [**nil**](nil.md) |  | 
**SourceName** | [**nil**](nil.md) |  | 
**Message** | [**nil**](nil.md) |  | 
**Status** | [**nil**](nil.md) |  | 
**RowCount** | Pointer to **float32** |  | [optional] 
**MetricCount** | Pointer to **float32** |  | [optional] 
**Timestamp** | [**nil**](nil.md) |  | 

## Methods

### NewIngestionStatusDto

`func NewIngestionStatusDto(ingestionDataset nil, ingestionSource nil, sourceName nil, message nil, status nil, timestamp nil, ) *IngestionStatusDto`

NewIngestionStatusDto instantiates a new IngestionStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionStatusDtoWithDefaults

`func NewIngestionStatusDtoWithDefaults() *IngestionStatusDto`

NewIngestionStatusDtoWithDefaults instantiates a new IngestionStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDs

`func (o *IngestionStatusDto) GetDs() time.Time`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *IngestionStatusDto) GetDsOk() (*time.Time, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *IngestionStatusDto) SetDs(v time.Time)`

SetDs sets Ds field to given value.

### HasDs

`func (o *IngestionStatusDto) HasDs() bool`

HasDs returns a boolean if a field has been set.

### GetIngestionDataset

`func (o *IngestionStatusDto) GetIngestionDataset() nil`

GetIngestionDataset returns the IngestionDataset field if non-nil, zero value otherwise.

### GetIngestionDatasetOk

`func (o *IngestionStatusDto) GetIngestionDatasetOk() (*nil, bool)`

GetIngestionDatasetOk returns a tuple with the IngestionDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionDataset

`func (o *IngestionStatusDto) SetIngestionDataset(v nil)`

SetIngestionDataset sets IngestionDataset field to given value.


### GetIngestionSource

`func (o *IngestionStatusDto) GetIngestionSource() nil`

GetIngestionSource returns the IngestionSource field if non-nil, zero value otherwise.

### GetIngestionSourceOk

`func (o *IngestionStatusDto) GetIngestionSourceOk() (*nil, bool)`

GetIngestionSourceOk returns a tuple with the IngestionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionSource

`func (o *IngestionStatusDto) SetIngestionSource(v nil)`

SetIngestionSource sets IngestionSource field to given value.


### GetSourceName

`func (o *IngestionStatusDto) GetSourceName() nil`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *IngestionStatusDto) GetSourceNameOk() (*nil, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *IngestionStatusDto) SetSourceName(v nil)`

SetSourceName sets SourceName field to given value.


### GetMessage

`func (o *IngestionStatusDto) GetMessage() nil`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IngestionStatusDto) GetMessageOk() (*nil, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IngestionStatusDto) SetMessage(v nil)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *IngestionStatusDto) GetStatus() nil`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IngestionStatusDto) GetStatusOk() (*nil, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IngestionStatusDto) SetStatus(v nil)`

SetStatus sets Status field to given value.


### GetRowCount

`func (o *IngestionStatusDto) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *IngestionStatusDto) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *IngestionStatusDto) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *IngestionStatusDto) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetMetricCount

`func (o *IngestionStatusDto) GetMetricCount() float32`

GetMetricCount returns the MetricCount field if non-nil, zero value otherwise.

### GetMetricCountOk

`func (o *IngestionStatusDto) GetMetricCountOk() (*float32, bool)`

GetMetricCountOk returns a tuple with the MetricCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCount

`func (o *IngestionStatusDto) SetMetricCount(v float32)`

SetMetricCount sets MetricCount field to given value.

### HasMetricCount

`func (o *IngestionStatusDto) HasMetricCount() bool`

HasMetricCount returns a boolean if a field has been set.

### GetTimestamp

`func (o *IngestionStatusDto) GetTimestamp() nil`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IngestionStatusDto) GetTimestampOk() (*nil, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IngestionStatusDto) SetTimestamp(v nil)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


