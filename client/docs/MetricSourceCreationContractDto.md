# MetricSourceCreationContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the metric source, serving as its primary identifier. | 
**Description** | Pointer to **string** | An optional description for the metric source, providing context and details about its purpose and usage. | [optional] 
**Tags** | Pointer to **[]string** | Optional array of tags to categorize the metric source, facilitating easier organization and retrieval. | [optional] 
**Sql** | **string** | The SQL query or statement used to extract data from the metric source. | 
**TimestampColumn** | **string** | The name of the column containing timestamp data for the metric source. | 
**TimestampAsDay** | Pointer to **bool** | Indicates whether the timestamp should be treated as a day-level granularity. | [optional] 
**IdTypeMapping** | [**[]MetricSourceContractDtoIdTypeMappingInner**](MetricSourceContractDtoIdTypeMappingInner.md) | Array defining the mapping between Statsig unit IDs and their respective source columns. | 
**SourceType** | Pointer to **string** | The type of source, indicating whether it is a database table or a custom query. | [optional] 
**TableName** | Pointer to **string** | The name of the database table if the source type is \&quot;table\&quot;. | [optional] 
**CustomFieldMapping** | Pointer to [**[]MetricSourceContractDtoCustomFieldMappingInner**](MetricSourceContractDtoCustomFieldMappingInner.md) | Optional array defining mappings for custom fields using specific formulas. | [optional] 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewMetricSourceCreationContractDto

`func NewMetricSourceCreationContractDto(name string, sql string, timestampColumn string, idTypeMapping []MetricSourceContractDtoIdTypeMappingInner, ) *MetricSourceCreationContractDto`

NewMetricSourceCreationContractDto instantiates a new MetricSourceCreationContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSourceCreationContractDtoWithDefaults

`func NewMetricSourceCreationContractDtoWithDefaults() *MetricSourceCreationContractDto`

NewMetricSourceCreationContractDtoWithDefaults instantiates a new MetricSourceCreationContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricSourceCreationContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricSourceCreationContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricSourceCreationContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MetricSourceCreationContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricSourceCreationContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricSourceCreationContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricSourceCreationContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *MetricSourceCreationContractDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricSourceCreationContractDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricSourceCreationContractDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricSourceCreationContractDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSql

`func (o *MetricSourceCreationContractDto) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *MetricSourceCreationContractDto) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *MetricSourceCreationContractDto) SetSql(v string)`

SetSql sets Sql field to given value.


### GetTimestampColumn

`func (o *MetricSourceCreationContractDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *MetricSourceCreationContractDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *MetricSourceCreationContractDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.


### GetTimestampAsDay

`func (o *MetricSourceCreationContractDto) GetTimestampAsDay() bool`

GetTimestampAsDay returns the TimestampAsDay field if non-nil, zero value otherwise.

### GetTimestampAsDayOk

`func (o *MetricSourceCreationContractDto) GetTimestampAsDayOk() (*bool, bool)`

GetTimestampAsDayOk returns a tuple with the TimestampAsDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAsDay

`func (o *MetricSourceCreationContractDto) SetTimestampAsDay(v bool)`

SetTimestampAsDay sets TimestampAsDay field to given value.

### HasTimestampAsDay

`func (o *MetricSourceCreationContractDto) HasTimestampAsDay() bool`

HasTimestampAsDay returns a boolean if a field has been set.

### GetIdTypeMapping

`func (o *MetricSourceCreationContractDto) GetIdTypeMapping() []MetricSourceContractDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *MetricSourceCreationContractDto) GetIdTypeMappingOk() (*[]MetricSourceContractDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *MetricSourceCreationContractDto) SetIdTypeMapping(v []MetricSourceContractDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.


### GetSourceType

`func (o *MetricSourceCreationContractDto) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MetricSourceCreationContractDto) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MetricSourceCreationContractDto) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *MetricSourceCreationContractDto) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTableName

`func (o *MetricSourceCreationContractDto) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *MetricSourceCreationContractDto) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *MetricSourceCreationContractDto) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *MetricSourceCreationContractDto) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetCustomFieldMapping

`func (o *MetricSourceCreationContractDto) GetCustomFieldMapping() []MetricSourceContractDtoCustomFieldMappingInner`

GetCustomFieldMapping returns the CustomFieldMapping field if non-nil, zero value otherwise.

### GetCustomFieldMappingOk

`func (o *MetricSourceCreationContractDto) GetCustomFieldMappingOk() (*[]MetricSourceContractDtoCustomFieldMappingInner, bool)`

GetCustomFieldMappingOk returns a tuple with the CustomFieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldMapping

`func (o *MetricSourceCreationContractDto) SetCustomFieldMapping(v []MetricSourceContractDtoCustomFieldMappingInner)`

SetCustomFieldMapping sets CustomFieldMapping field to given value.

### HasCustomFieldMapping

`func (o *MetricSourceCreationContractDto) HasCustomFieldMapping() bool`

HasCustomFieldMapping returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *MetricSourceCreationContractDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MetricSourceCreationContractDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *MetricSourceCreationContractDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *MetricSourceCreationContractDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


