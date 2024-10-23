# IngestionBackfillContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatestampStart** | **string** |  | 
**DatestampEnd** | **string** |  | 
**Type** | **string** |  | 
**Source** | Pointer to [**IngestionBackfillContractDtoSource**](IngestionBackfillContractDtoSource.md) |  | [optional] 
**Dataset** | **string** |  | 

## Methods

### NewIngestionBackfillContractDto

`func NewIngestionBackfillContractDto(datestampStart string, datestampEnd string, type_ string, dataset string, ) *IngestionBackfillContractDto`

NewIngestionBackfillContractDto instantiates a new IngestionBackfillContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionBackfillContractDtoWithDefaults

`func NewIngestionBackfillContractDtoWithDefaults() *IngestionBackfillContractDto`

NewIngestionBackfillContractDtoWithDefaults instantiates a new IngestionBackfillContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatestampStart

`func (o *IngestionBackfillContractDto) GetDatestampStart() string`

GetDatestampStart returns the DatestampStart field if non-nil, zero value otherwise.

### GetDatestampStartOk

`func (o *IngestionBackfillContractDto) GetDatestampStartOk() (*string, bool)`

GetDatestampStartOk returns a tuple with the DatestampStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatestampStart

`func (o *IngestionBackfillContractDto) SetDatestampStart(v string)`

SetDatestampStart sets DatestampStart field to given value.


### GetDatestampEnd

`func (o *IngestionBackfillContractDto) GetDatestampEnd() string`

GetDatestampEnd returns the DatestampEnd field if non-nil, zero value otherwise.

### GetDatestampEndOk

`func (o *IngestionBackfillContractDto) GetDatestampEndOk() (*string, bool)`

GetDatestampEndOk returns a tuple with the DatestampEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatestampEnd

`func (o *IngestionBackfillContractDto) SetDatestampEnd(v string)`

SetDatestampEnd sets DatestampEnd field to given value.


### GetType

`func (o *IngestionBackfillContractDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionBackfillContractDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionBackfillContractDto) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *IngestionBackfillContractDto) GetSource() IngestionBackfillContractDtoSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IngestionBackfillContractDto) GetSourceOk() (*IngestionBackfillContractDtoSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IngestionBackfillContractDto) SetSource(v IngestionBackfillContractDtoSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *IngestionBackfillContractDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDataset

`func (o *IngestionBackfillContractDto) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionBackfillContractDto) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionBackfillContractDto) SetDataset(v string)`

SetDataset sets Dataset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


