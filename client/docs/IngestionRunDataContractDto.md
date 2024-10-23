# IngestionRunDataContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunID** | **string** |  | 
**LatestStatus** | **string** |  | 
**LastUpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Trigger** | **string** |  | 
**Sources** | **[]string** |  | 
**DateStamps** | **[]string** |  | 
**RunHistory** | [**[]IngestionRunDataContractDtoRunHistoryInner**](IngestionRunDataContractDtoRunHistoryInner.md) |  | 
**GranularHistory** | [**[]IngestionRunDataContractDtoGranularHistoryInner**](IngestionRunDataContractDtoGranularHistoryInner.md) |  | 

## Methods

### NewIngestionRunDataContractDto

`func NewIngestionRunDataContractDto(runID string, latestStatus string, lastUpdatedAt time.Time, createdAt time.Time, trigger string, sources []string, dateStamps []string, runHistory []IngestionRunDataContractDtoRunHistoryInner, granularHistory []IngestionRunDataContractDtoGranularHistoryInner, ) *IngestionRunDataContractDto`

NewIngestionRunDataContractDto instantiates a new IngestionRunDataContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionRunDataContractDtoWithDefaults

`func NewIngestionRunDataContractDtoWithDefaults() *IngestionRunDataContractDto`

NewIngestionRunDataContractDtoWithDefaults instantiates a new IngestionRunDataContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunID

`func (o *IngestionRunDataContractDto) GetRunID() string`

GetRunID returns the RunID field if non-nil, zero value otherwise.

### GetRunIDOk

`func (o *IngestionRunDataContractDto) GetRunIDOk() (*string, bool)`

GetRunIDOk returns a tuple with the RunID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunID

`func (o *IngestionRunDataContractDto) SetRunID(v string)`

SetRunID sets RunID field to given value.


### GetLatestStatus

`func (o *IngestionRunDataContractDto) GetLatestStatus() string`

GetLatestStatus returns the LatestStatus field if non-nil, zero value otherwise.

### GetLatestStatusOk

`func (o *IngestionRunDataContractDto) GetLatestStatusOk() (*string, bool)`

GetLatestStatusOk returns a tuple with the LatestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestStatus

`func (o *IngestionRunDataContractDto) SetLatestStatus(v string)`

SetLatestStatus sets LatestStatus field to given value.


### GetLastUpdatedAt

`func (o *IngestionRunDataContractDto) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *IngestionRunDataContractDto) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *IngestionRunDataContractDto) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.


### GetCreatedAt

`func (o *IngestionRunDataContractDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngestionRunDataContractDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngestionRunDataContractDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTrigger

`func (o *IngestionRunDataContractDto) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *IngestionRunDataContractDto) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *IngestionRunDataContractDto) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetSources

`func (o *IngestionRunDataContractDto) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *IngestionRunDataContractDto) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *IngestionRunDataContractDto) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetDateStamps

`func (o *IngestionRunDataContractDto) GetDateStamps() []string`

GetDateStamps returns the DateStamps field if non-nil, zero value otherwise.

### GetDateStampsOk

`func (o *IngestionRunDataContractDto) GetDateStampsOk() (*[]string, bool)`

GetDateStampsOk returns a tuple with the DateStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStamps

`func (o *IngestionRunDataContractDto) SetDateStamps(v []string)`

SetDateStamps sets DateStamps field to given value.


### GetRunHistory

`func (o *IngestionRunDataContractDto) GetRunHistory() []IngestionRunDataContractDtoRunHistoryInner`

GetRunHistory returns the RunHistory field if non-nil, zero value otherwise.

### GetRunHistoryOk

`func (o *IngestionRunDataContractDto) GetRunHistoryOk() (*[]IngestionRunDataContractDtoRunHistoryInner, bool)`

GetRunHistoryOk returns a tuple with the RunHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunHistory

`func (o *IngestionRunDataContractDto) SetRunHistory(v []IngestionRunDataContractDtoRunHistoryInner)`

SetRunHistory sets RunHistory field to given value.


### GetGranularHistory

`func (o *IngestionRunDataContractDto) GetGranularHistory() []IngestionRunDataContractDtoGranularHistoryInner`

GetGranularHistory returns the GranularHistory field if non-nil, zero value otherwise.

### GetGranularHistoryOk

`func (o *IngestionRunDataContractDto) GetGranularHistoryOk() (*[]IngestionRunDataContractDtoGranularHistoryInner, bool)`

GetGranularHistoryOk returns a tuple with the GranularHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularHistory

`func (o *IngestionRunDataContractDto) SetGranularHistory(v []IngestionRunDataContractDtoGranularHistoryInner)`

SetGranularHistory sets GranularHistory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


