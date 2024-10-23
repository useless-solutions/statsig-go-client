# MetricSourceContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the metric source, serving as its primary identifier. | 
**Description** | **string** | A detailed description of the metric source, providing context and usage information. | 
**Tags** | Pointer to **[]string** | Optional tags for categorizing the metric source and improving searchability. | [optional] 
**Sql** | **string** | The SQL query or statement used to extract data from the metric source. | 
**TimestampColumn** | **string** | The name of the column containing timestamp data for the metric source. | 
**TimestampAsDay** | Pointer to **bool** | Indicates whether the timestamp should be treated as a day-level granularity. | [optional] 
**IdTypeMapping** | [**[]MetricSourceContractDtoIdTypeMappingInner**](MetricSourceContractDtoIdTypeMappingInner.md) | Array defining the mapping between Statsig unit IDs and their respective source columns. | 
**SourceType** | Pointer to **string** | The type of source, indicating whether it is a database table or a custom query. | [optional] 
**TableName** | Pointer to **string** | The name of the database table if the source type is \&quot;table\&quot;. | [optional] 
**CustomFieldMapping** | Pointer to [**[]MetricSourceContractDtoCustomFieldMappingInner**](MetricSourceContractDtoCustomFieldMappingInner.md) | Optional array defining mappings for custom fields using specific formulas. | [optional] 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewMetricSourceContractDto

`func NewMetricSourceContractDto(name string, description string, sql string, timestampColumn string, idTypeMapping []MetricSourceContractDtoIdTypeMappingInner, ) *MetricSourceContractDto`

NewMetricSourceContractDto instantiates a new MetricSourceContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSourceContractDtoWithDefaults

`func NewMetricSourceContractDtoWithDefaults() *MetricSourceContractDto`

NewMetricSourceContractDtoWithDefaults instantiates a new MetricSourceContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricSourceContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricSourceContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricSourceContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MetricSourceContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricSourceContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricSourceContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *MetricSourceContractDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricSourceContractDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricSourceContractDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricSourceContractDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSql

`func (o *MetricSourceContractDto) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *MetricSourceContractDto) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *MetricSourceContractDto) SetSql(v string)`

SetSql sets Sql field to given value.


### GetTimestampColumn

`func (o *MetricSourceContractDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *MetricSourceContractDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *MetricSourceContractDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.


### GetTimestampAsDay

`func (o *MetricSourceContractDto) GetTimestampAsDay() bool`

GetTimestampAsDay returns the TimestampAsDay field if non-nil, zero value otherwise.

### GetTimestampAsDayOk

`func (o *MetricSourceContractDto) GetTimestampAsDayOk() (*bool, bool)`

GetTimestampAsDayOk returns a tuple with the TimestampAsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAsDay

`func (o *MetricSourceContractDto) SetTimestampAsDay(v bool)`

SetTimestampAsDay sets TimestampAsDay field to given value.

### HasTimestampAsDay

`func (o *MetricSourceContractDto) HasTimestampAsDay() bool`

HasTimestampAsDay returns a boolean if a field has been set.

### GetIdTypeMapping

`func (o *MetricSourceContractDto) GetIdTypeMapping() []MetricSourceContractDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *MetricSourceContractDto) GetIdTypeMappingOk() (*[]MetricSourceContractDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *MetricSourceContractDto) SetIdTypeMapping(v []MetricSourceContractDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.


### GetSourceType

`func (o *MetricSourceContractDto) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MetricSourceContractDto) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MetricSourceContractDto) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *MetricSourceContractDto) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTableName

`func (o *MetricSourceContractDto) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *MetricSourceContractDto) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *MetricSourceContractDto) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *MetricSourceContractDto) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetCustomFieldMapping

`func (o *MetricSourceContractDto) GetCustomFieldMapping() []MetricSourceContractDtoCustomFieldMappingInner`

GetCustomFieldMapping returns the CustomFieldMapping field if non-nil, zero value otherwise.

### GetCustomFieldMappingOk

`func (o *MetricSourceContractDto) GetCustomFieldMappingOk() (*[]MetricSourceContractDtoCustomFieldMappingInner, bool)`

GetCustomFieldMappingOk returns a tuple with the CustomFieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldMapping

`func (o *MetricSourceContractDto) SetCustomFieldMapping(v []MetricSourceContractDtoCustomFieldMappingInner)`

SetCustomFieldMapping sets CustomFieldMapping field to given value.

### HasCustomFieldMapping

`func (o *MetricSourceContractDto) HasCustomFieldMapping() bool`

HasCustomFieldMapping returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *MetricSourceContractDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MetricSourceContractDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *MetricSourceContractDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *MetricSourceContractDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


