# EntityPropertySourceCreationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique identifier for the entity property source. | 
**Description** | Pointer to **string** | Optional detailed context for the entity property source. | [optional] 
**Tags** | Pointer to **[]string** | Optional tags for categorization. | [optional] 
**Sql** | **string** | SQL query defining the data source. | 
**TimestampColumn** | Pointer to **string** | Optional column name for timestamp. | [optional] 
**TimestampAsDay** | Pointer to **bool** | Indicates if the timestamp is treated as a day. | [optional] 
**IdTypeMapping** | [**[]EntityPropertySourceDtoIdTypeMappingInner**](EntityPropertySourceDtoIdTypeMappingInner.md) | Mappings of Statsig units to their columns. | 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewEntityPropertySourceCreationDto

`func NewEntityPropertySourceCreationDto(name string, sql string, idTypeMapping []EntityPropertySourceDtoIdTypeMappingInner, ) *EntityPropertySourceCreationDto`

NewEntityPropertySourceCreationDto instantiates a new EntityPropertySourceCreationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertySourceCreationDtoWithDefaults

`func NewEntityPropertySourceCreationDtoWithDefaults() *EntityPropertySourceCreationDto`

NewEntityPropertySourceCreationDtoWithDefaults instantiates a new EntityPropertySourceCreationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityPropertySourceCreationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityPropertySourceCreationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityPropertySourceCreationDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntityPropertySourceCreationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityPropertySourceCreationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityPropertySourceCreationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityPropertySourceCreationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *EntityPropertySourceCreationDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityPropertySourceCreationDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityPropertySourceCreationDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntityPropertySourceCreationDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSql

`func (o *EntityPropertySourceCreationDto) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *EntityPropertySourceCreationDto) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *EntityPropertySourceCreationDto) SetSql(v string)`

SetSql sets Sql field to given value.


### GetTimestampColumn

`func (o *EntityPropertySourceCreationDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *EntityPropertySourceCreationDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *EntityPropertySourceCreationDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.

### HasTimestampColumn

`func (o *EntityPropertySourceCreationDto) HasTimestampColumn() bool`

HasTimestampColumn returns a boolean if a field has been set.

### GetTimestampAsDay

`func (o *EntityPropertySourceCreationDto) GetTimestampAsDay() bool`

GetTimestampAsDay returns the TimestampAsDay field if non-nil, zero value otherwise.

### GetTimestampAsDayOk

`func (o *EntityPropertySourceCreationDto) GetTimestampAsDayOk() (*bool, bool)`

GetTimestampAsDayOk returns a tuple with the TimestampAsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAsDay

`func (o *EntityPropertySourceCreationDto) SetTimestampAsDay(v bool)`

SetTimestampAsDay sets TimestampAsDay field to given value.

### HasTimestampAsDay

`func (o *EntityPropertySourceCreationDto) HasTimestampAsDay() bool`

HasTimestampAsDay returns a boolean if a field has been set.

### GetIdTypeMapping

`func (o *EntityPropertySourceCreationDto) GetIdTypeMapping() []EntityPropertySourceDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *EntityPropertySourceCreationDto) GetIdTypeMappingOk() (*[]EntityPropertySourceDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *EntityPropertySourceCreationDto) SetIdTypeMapping(v []EntityPropertySourceDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.


### GetIsReadOnly

`func (o *EntityPropertySourceCreationDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *EntityPropertySourceCreationDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *EntityPropertySourceCreationDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *EntityPropertySourceCreationDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


