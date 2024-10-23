# EntityPropertySourcePartialUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique identifier for the entity property source. | [optional] 
**Description** | Pointer to **string** | Detailed context and purpose of the entity property source. | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorization and search. | [optional] 
**Sql** | Pointer to **interface{}** | SQL query defining the data source | [optional] [readonly] 
**TimestampColumn** | Pointer to **string** | Optional column name for timestamp. | [optional] 
**TimestampAsDay** | Pointer to **bool** | Indicates if the timestamp is treated as a day. | [optional] 
**IdTypeMapping** | Pointer to [**[]EntityPropertySourceDtoIdTypeMappingInner**](EntityPropertySourceDtoIdTypeMappingInner.md) | Mappings of Statsig units to their columns. | [optional] 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewEntityPropertySourcePartialUpdateDto

`func NewEntityPropertySourcePartialUpdateDto() *EntityPropertySourcePartialUpdateDto`

NewEntityPropertySourcePartialUpdateDto instantiates a new EntityPropertySourcePartialUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertySourcePartialUpdateDtoWithDefaults

`func NewEntityPropertySourcePartialUpdateDtoWithDefaults() *EntityPropertySourcePartialUpdateDto`

NewEntityPropertySourcePartialUpdateDtoWithDefaults instantiates a new EntityPropertySourcePartialUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityPropertySourcePartialUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityPropertySourcePartialUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityPropertySourcePartialUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityPropertySourcePartialUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EntityPropertySourcePartialUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityPropertySourcePartialUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityPropertySourcePartialUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityPropertySourcePartialUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *EntityPropertySourcePartialUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityPropertySourcePartialUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityPropertySourcePartialUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntityPropertySourcePartialUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSql

`func (o *EntityPropertySourcePartialUpdateDto) GetSql() interface{}`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *EntityPropertySourcePartialUpdateDto) GetSqlOk() (*interface{}, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *EntityPropertySourcePartialUpdateDto) SetSql(v interface{})`

SetSql sets Sql field to given value.

### HasSql

`func (o *EntityPropertySourcePartialUpdateDto) HasSql() bool`

HasSql returns a boolean if a field has been set.

### SetSqlNil

`func (o *EntityPropertySourcePartialUpdateDto) SetSqlNil(b bool)`

 SetSqlNil sets the value for Sql to be an explicit nil

### UnsetSql
`func (o *EntityPropertySourcePartialUpdateDto) UnsetSql()`

UnsetSql ensures that no value is present for Sql, not even an explicit nil
### GetTimestampColumn

`func (o *EntityPropertySourcePartialUpdateDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *EntityPropertySourcePartialUpdateDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *EntityPropertySourcePartialUpdateDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.

### HasTimestampColumn

`func (o *EntityPropertySourcePartialUpdateDto) HasTimestampColumn() bool`

HasTimestampColumn returns a boolean if a field has been set.

### GetTimestampAsDay

`func (o *EntityPropertySourcePartialUpdateDto) GetTimestampAsDay() bool`

GetTimestampAsDay returns the TimestampAsDay field if non-nil, zero value otherwise.

### GetTimestampAsDayOk

`func (o *EntityPropertySourcePartialUpdateDto) GetTimestampAsDayOk() (*bool, bool)`

GetTimestampAsDayOk returns a tuple with the TimestampAsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAsDay

`func (o *EntityPropertySourcePartialUpdateDto) SetTimestampAsDay(v bool)`

SetTimestampAsDay sets TimestampAsDay field to given value.

### HasTimestampAsDay

`func (o *EntityPropertySourcePartialUpdateDto) HasTimestampAsDay() bool`

HasTimestampAsDay returns a boolean if a field has been set.

### GetIdTypeMapping

`func (o *EntityPropertySourcePartialUpdateDto) GetIdTypeMapping() []EntityPropertySourceDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *EntityPropertySourcePartialUpdateDto) GetIdTypeMappingOk() (*[]EntityPropertySourceDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *EntityPropertySourcePartialUpdateDto) SetIdTypeMapping(v []EntityPropertySourceDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.

### HasIdTypeMapping

`func (o *EntityPropertySourcePartialUpdateDto) HasIdTypeMapping() bool`

HasIdTypeMapping returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *EntityPropertySourcePartialUpdateDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *EntityPropertySourcePartialUpdateDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *EntityPropertySourcePartialUpdateDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *EntityPropertySourcePartialUpdateDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


