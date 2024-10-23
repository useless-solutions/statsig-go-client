# IngestionSourceCreateContractDtoOneOf1ColumnMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitId** | Pointer to **string** | The unique user identifier this event is for. This might not necessarily be a single column for userID - it could be spread across multiple columns for deviceID etc. | [optional] 
**EventName** | **string** | Name of the event. String under 128 characters, using ‘_’ for spaces. | 
**Timestamp** | **string** | Unix timestamp in seconds of the event (ex. 1729692720) | 
**Ids** | **map[string]string** |  | [default to {}]
**Metadata** | Pointer to **map[string]string** |  | [optional] [default to {}]
**MetadataObject** | Pointer to **string** |  | [optional] [default to "null"]
**EventValue** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewIngestionSourceCreateContractDtoOneOf1ColumnMapping

`func NewIngestionSourceCreateContractDtoOneOf1ColumnMapping(eventName string, timestamp string, ids map[string]string, ) *IngestionSourceCreateContractDtoOneOf1ColumnMapping`

NewIngestionSourceCreateContractDtoOneOf1ColumnMapping instantiates a new IngestionSourceCreateContractDtoOneOf1ColumnMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSourceCreateContractDtoOneOf1ColumnMappingWithDefaults

`func NewIngestionSourceCreateContractDtoOneOf1ColumnMappingWithDefaults() *IngestionSourceCreateContractDtoOneOf1ColumnMapping`

NewIngestionSourceCreateContractDtoOneOf1ColumnMappingWithDefaults instantiates a new IngestionSourceCreateContractDtoOneOf1ColumnMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitId

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### GetEventName

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetTimestamp

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetIds

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetIds() map[string]string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetIdsOk() (*map[string]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) SetIds(v map[string]string)`

SetIds sets Ids field to given value.


### GetMetadata

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetadataObject

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetMetadataObject() string`

GetMetadataObject returns the MetadataObject field if non-nil, zero value otherwise.

### GetMetadataObjectOk

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetMetadataObjectOk() (*string, bool)`

GetMetadataObjectOk returns a tuple with the MetadataObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataObject

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) SetMetadataObject(v string)`

SetMetadataObject sets MetadataObject field to given value.

### HasMetadataObject

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) HasMetadataObject() bool`

HasMetadataObject returns a boolean if a field has been set.

### GetEventValue

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetEventValue() string`

GetEventValue returns the EventValue field if non-nil, zero value otherwise.

### GetEventValueOk

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) GetEventValueOk() (*string, bool)`

GetEventValueOk returns a tuple with the EventValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventValue

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) SetEventValue(v string)`

SetEventValue sets EventValue field to given value.

### HasEventValue

`func (o *IngestionSourceCreateContractDtoOneOf1ColumnMapping) HasEventValue() bool`

HasEventValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


