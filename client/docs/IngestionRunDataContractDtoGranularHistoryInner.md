# IngestionRunDataContractDtoGranularHistoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**LatestSourceStatus** | **string** |  | 
**StatusByDate** | [**[]IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner**](IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner.md) |  | 

## Methods

### NewIngestionRunDataContractDtoGranularHistoryInner

`func NewIngestionRunDataContractDtoGranularHistoryInner(source string, latestSourceStatus string, statusByDate []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner, ) *IngestionRunDataContractDtoGranularHistoryInner`

NewIngestionRunDataContractDtoGranularHistoryInner instantiates a new IngestionRunDataContractDtoGranularHistoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionRunDataContractDtoGranularHistoryInnerWithDefaults

`func NewIngestionRunDataContractDtoGranularHistoryInnerWithDefaults() *IngestionRunDataContractDtoGranularHistoryInner`

NewIngestionRunDataContractDtoGranularHistoryInnerWithDefaults instantiates a new IngestionRunDataContractDtoGranularHistoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *IngestionRunDataContractDtoGranularHistoryInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IngestionRunDataContractDtoGranularHistoryInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IngestionRunDataContractDtoGranularHistoryInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetLatestSourceStatus

`func (o *IngestionRunDataContractDtoGranularHistoryInner) GetLatestSourceStatus() string`

GetLatestSourceStatus returns the LatestSourceStatus field if non-nil, zero value otherwise.

### GetLatestSourceStatusOk

`func (o *IngestionRunDataContractDtoGranularHistoryInner) GetLatestSourceStatusOk() (*string, bool)`

GetLatestSourceStatusOk returns a tuple with the LatestSourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSourceStatus

`func (o *IngestionRunDataContractDtoGranularHistoryInner) SetLatestSourceStatus(v string)`

SetLatestSourceStatus sets LatestSourceStatus field to given value.


### GetStatusByDate

`func (o *IngestionRunDataContractDtoGranularHistoryInner) GetStatusByDate() []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner`

GetStatusByDate returns the StatusByDate field if non-nil, zero value otherwise.

### GetStatusByDateOk

`func (o *IngestionRunDataContractDtoGranularHistoryInner) GetStatusByDateOk() (*[]IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner, bool)`

GetStatusByDateOk returns a tuple with the StatusByDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusByDate

`func (o *IngestionRunDataContractDtoGranularHistoryInner) SetStatusByDate(v []IngestionRunDataContractDtoGranularHistoryInnerStatusByDateInner)`

SetStatusByDate sets StatusByDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


