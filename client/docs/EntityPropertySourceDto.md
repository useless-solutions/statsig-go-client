# EntityPropertySourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique identifier for the entity property source. | 
**Description** | **string** | Detailed context and purpose of the entity property source. | 
**Tags** | **[]string** | Tags for categorization and search. | 
**Sql** | **string** | SQL query defining the data source. | 
**TimestampColumn** | Pointer to **string** | Optional column name for timestamp. | [optional] 
**TimestampAsDay** | Pointer to **bool** | Indicates if the timestamp is treated as a day. | [optional] 
**IdTypeMapping** | [**[]EntityPropertySourceDtoIdTypeMappingInner**](EntityPropertySourceDtoIdTypeMappingInner.md) | Mappings of Statsig units to their columns. | 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewEntityPropertySourceDto

`func NewEntityPropertySourceDto(name string, description string, tags []string, sql string, idTypeMapping []EntityPropertySourceDtoIdTypeMappingInner, ) *EntityPropertySourceDto`

NewEntityPropertySourceDto instantiates a new EntityPropertySourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertySourceDtoWithDefaults

`func NewEntityPropertySourceDtoWithDefaults() *EntityPropertySourceDto`

NewEntityPropertySourceDtoWithDefaults instantiates a new EntityPropertySourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityPropertySourceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityPropertySourceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityPropertySourceDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntityPropertySourceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityPropertySourceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityPropertySourceDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *EntityPropertySourceDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityPropertySourceDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityPropertySourceDto) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetSql

`func (o *EntityPropertySourceDto) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *EntityPropertySourceDto) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *EntityPropertySourceDto) SetSql(v string)`

SetSql sets Sql field to given value.


### GetTimestampColumn

`func (o *EntityPropertySourceDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *EntityPropertySourceDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *EntityPropertySourceDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.

### HasTimestampColumn

`func (o *EntityPropertySourceDto) HasTimestampColumn() bool`

HasTimestampColumn returns a boolean if a field has been set.

### GetTimestampAsDay

`func (o *EntityPropertySourceDto) GetTimestampAsDay() bool`

GetTimestampAsDay returns the TimestampAsDay field if non-nil, zero value otherwise.

### GetTimestampAsDayOk

`func (o *EntityPropertySourceDto) GetTimestampAsDayOk() (*bool, bool)`

GetTimestampAsDayOk returns a tuple with the TimestampAsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAsDay

`func (o *EntityPropertySourceDto) SetTimestampAsDay(v bool)`

SetTimestampAsDay sets TimestampAsDay field to given value.

### HasTimestampAsDay

`func (o *EntityPropertySourceDto) HasTimestampAsDay() bool`

HasTimestampAsDay returns a boolean if a field has been set.

### GetIdTypeMapping

`func (o *EntityPropertySourceDto) GetIdTypeMapping() []EntityPropertySourceDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *EntityPropertySourceDto) GetIdTypeMappingOk() (*[]EntityPropertySourceDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *EntityPropertySourceDto) SetIdTypeMapping(v []EntityPropertySourceDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.


### GetIsReadOnly

`func (o *EntityPropertySourceDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *EntityPropertySourceDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *EntityPropertySourceDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *EntityPropertySourceDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


