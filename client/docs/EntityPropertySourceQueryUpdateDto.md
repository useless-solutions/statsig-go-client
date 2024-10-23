# EntityPropertySourceQueryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional new name for the entity property source. | [optional] 
**Description** | Pointer to **string** | Optional updated context for the entity property source. | [optional] 
**Tags** | Pointer to **[]string** | Optional updated tags for categorization. | [optional] 
**Sql** | **string** | SQL query defining the data source. | 
**TimestampColumn** | Pointer to **string** | Optional column name for timestamp. | [optional] 
**TimestampAsDay** | Pointer to **bool** | Indicates if the timestamp is treated as a day. | [optional] 
**IdTypeMapping** | [**[]EntityPropertySourceDtoIdTypeMappingInner**](EntityPropertySourceDtoIdTypeMappingInner.md) | Mappings of Statsig units to their columns. | 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewEntityPropertySourceQueryUpdateDto

`func NewEntityPropertySourceQueryUpdateDto(sql string, idTypeMapping []EntityPropertySourceDtoIdTypeMappingInner, ) *EntityPropertySourceQueryUpdateDto`

NewEntityPropertySourceQueryUpdateDto instantiates a new EntityPropertySourceQueryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertySourceQueryUpdateDtoWithDefaults

`func NewEntityPropertySourceQueryUpdateDtoWithDefaults() *EntityPropertySourceQueryUpdateDto`

NewEntityPropertySourceQueryUpdateDtoWithDefaults instantiates a new EntityPropertySourceQueryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityPropertySourceQueryUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityPropertySourceQueryUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityPropertySourceQueryUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityPropertySourceQueryUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EntityPropertySourceQueryUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityPropertySourceQueryUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityPropertySourceQueryUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityPropertySourceQueryUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *EntityPropertySourceQueryUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityPropertySourceQueryUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityPropertySourceQueryUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntityPropertySourceQueryUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSql

`func (o *EntityPropertySourceQueryUpdateDto) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *EntityPropertySourceQueryUpdateDto) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *EntityPropertySourceQueryUpdateDto) SetSql(v string)`

SetSql sets Sql field to given value.


### GetTimestampColumn

`func (o *EntityPropertySourceQueryUpdateDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *EntityPropertySourceQueryUpdateDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *EntityPropertySourceQueryUpdateDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.

### HasTimestampColumn

`func (o *EntityPropertySourceQueryUpdateDto) HasTimestampColumn() bool`

HasTimestampColumn returns a boolean if a field has been set.

### GetTimestampAsDay

`func (o *EntityPropertySourceQueryUpdateDto) GetTimestampAsDay() bool`

GetTimestampAsDay returns the TimestampAsDay field if non-nil, zero value otherwise.

### GetTimestampAsDayOk

`func (o *EntityPropertySourceQueryUpdateDto) GetTimestampAsDayOk() (*bool, bool)`

GetTimestampAsDayOk returns a tuple with the TimestampAsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAsDay

`func (o *EntityPropertySourceQueryUpdateDto) SetTimestampAsDay(v bool)`

SetTimestampAsDay sets TimestampAsDay field to given value.

### HasTimestampAsDay

`func (o *EntityPropertySourceQueryUpdateDto) HasTimestampAsDay() bool`

HasTimestampAsDay returns a boolean if a field has been set.

### GetIdTypeMapping

`func (o *EntityPropertySourceQueryUpdateDto) GetIdTypeMapping() []EntityPropertySourceDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *EntityPropertySourceQueryUpdateDto) GetIdTypeMappingOk() (*[]EntityPropertySourceDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *EntityPropertySourceQueryUpdateDto) SetIdTypeMapping(v []EntityPropertySourceDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.


### GetIsReadOnly

`func (o *EntityPropertySourceQueryUpdateDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *EntityPropertySourceQueryUpdateDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *EntityPropertySourceQueryUpdateDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *EntityPropertySourceQueryUpdateDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


