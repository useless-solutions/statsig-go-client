# IngestionSourceCreateContractDtoOneOf2ColumnMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experiment** | **string** | Unique identifier for the experiment. | 
**GroupId** | **string** | Unique identifier for the experiment groups. | 
**UnitId** | Pointer to **string** | The unique user identifier this exposure is for. This might not necessarily be a single column for userID - it could be spread across multiple columns for deviceID etc. | [optional] 
**Timestamp** | **string** | Unix timestamp in seconds of the event (ex. 1729692720) | 
**Ids** | **map[string]string** |  | [default to {}]
**Metadata** | Pointer to **map[string]string** |  | [optional] [default to {}]
**MetadataObject** | Pointer to **string** |  | [optional] [default to "null"]
**EventValue** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewIngestionSourceCreateContractDtoOneOf2ColumnMapping

`func NewIngestionSourceCreateContractDtoOneOf2ColumnMapping(experiment string, groupId string, timestamp string, ids map[string]string, ) *IngestionSourceCreateContractDtoOneOf2ColumnMapping`

NewIngestionSourceCreateContractDtoOneOf2ColumnMapping instantiates a new IngestionSourceCreateContractDtoOneOf2ColumnMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSourceCreateContractDtoOneOf2ColumnMappingWithDefaults

`func NewIngestionSourceCreateContractDtoOneOf2ColumnMappingWithDefaults() *IngestionSourceCreateContractDtoOneOf2ColumnMapping`

NewIngestionSourceCreateContractDtoOneOf2ColumnMappingWithDefaults instantiates a new IngestionSourceCreateContractDtoOneOf2ColumnMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperiment

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetExperiment() string`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetExperimentOk() (*string, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetExperiment(v string)`

SetExperiment sets Experiment field to given value.


### GetGroupId

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetUnitId

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### GetTimestamp

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetIds

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetIds() map[string]string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetIdsOk() (*map[string]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetIds(v map[string]string)`

SetIds sets Ids field to given value.


### GetMetadata

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetadataObject

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadataObject() string`

GetMetadataObject returns the MetadataObject field if non-nil, zero value otherwise.

### GetMetadataObjectOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetMetadataObjectOk() (*string, bool)`

GetMetadataObjectOk returns a tuple with the MetadataObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataObject

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetMetadataObject(v string)`

SetMetadataObject sets MetadataObject field to given value.

### HasMetadataObject

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasMetadataObject() bool`

HasMetadataObject returns a boolean if a field has been set.

### GetEventValue

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetEventValue() string`

GetEventValue returns the EventValue field if non-nil, zero value otherwise.

### GetEventValueOk

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) GetEventValueOk() (*string, bool)`

GetEventValueOk returns a tuple with the EventValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventValue

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) SetEventValue(v string)`

SetEventValue sets EventValue field to given value.

### HasEventValue

`func (o *IngestionSourceCreateContractDtoOneOf2ColumnMapping) HasEventValue() bool`

HasEventValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


