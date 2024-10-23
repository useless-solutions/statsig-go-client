# AssignmentSourcePartialUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique identifier for the assignment source. | [optional] 
**Description** | Pointer to **string** | Detailed context and purpose of the assignment source. | [optional] 
**IsVerified** | Pointer to **bool** | Marks the assignment source as verified for internal trustworthiness. | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorization and search. | [optional] 
**Sql** | Pointer to **interface{}** | SQL query defining the data source for assignments. | [optional] [readonly] 
**TimestampColumn** | Pointer to **string** | Column name representing the timestamp of assignments. | [optional] 
**ExperimentIDColumn** | Pointer to **string** | Column name for the experiment ID associated with the assignments. | [optional] 
**GroupIDColumn** | Pointer to **string** | Column name for the group ID linked to the assignments. | [optional] 
**IdTypeMapping** | Pointer to [**[]AssignmentSourceCreationDtoIdTypeMappingInner**](AssignmentSourceCreationDtoIdTypeMappingInner.md) | Mappings of Statsig units to their respective columns. | [optional] 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewAssignmentSourcePartialUpdateDto

`func NewAssignmentSourcePartialUpdateDto() *AssignmentSourcePartialUpdateDto`

NewAssignmentSourcePartialUpdateDto instantiates a new AssignmentSourcePartialUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentSourcePartialUpdateDtoWithDefaults

`func NewAssignmentSourcePartialUpdateDtoWithDefaults() *AssignmentSourcePartialUpdateDto`

NewAssignmentSourcePartialUpdateDtoWithDefaults instantiates a new AssignmentSourcePartialUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssignmentSourcePartialUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssignmentSourcePartialUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssignmentSourcePartialUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssignmentSourcePartialUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AssignmentSourcePartialUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssignmentSourcePartialUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssignmentSourcePartialUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssignmentSourcePartialUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsVerified

`func (o *AssignmentSourcePartialUpdateDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *AssignmentSourcePartialUpdateDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *AssignmentSourcePartialUpdateDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *AssignmentSourcePartialUpdateDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetTags

`func (o *AssignmentSourcePartialUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssignmentSourcePartialUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssignmentSourcePartialUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssignmentSourcePartialUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSql

`func (o *AssignmentSourcePartialUpdateDto) GetSql() interface{}`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *AssignmentSourcePartialUpdateDto) GetSqlOk() (*interface{}, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *AssignmentSourcePartialUpdateDto) SetSql(v interface{})`

SetSql sets Sql field to given value.

### HasSql

`func (o *AssignmentSourcePartialUpdateDto) HasSql() bool`

HasSql returns a boolean if a field has been set.

### SetSqlNil

`func (o *AssignmentSourcePartialUpdateDto) SetSqlNil(b bool)`

 SetSqlNil sets the value for Sql to be an explicit nil

### UnsetSql
`func (o *AssignmentSourcePartialUpdateDto) UnsetSql()`

UnsetSql ensures that no value is present for Sql, not even an explicit nil
### GetTimestampColumn

`func (o *AssignmentSourcePartialUpdateDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *AssignmentSourcePartialUpdateDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *AssignmentSourcePartialUpdateDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.

### HasTimestampColumn

`func (o *AssignmentSourcePartialUpdateDto) HasTimestampColumn() bool`

HasTimestampColumn returns a boolean if a field has been set.

### GetExperimentIDColumn

`func (o *AssignmentSourcePartialUpdateDto) GetExperimentIDColumn() string`

GetExperimentIDColumn returns the ExperimentIDColumn field if non-nil, zero value otherwise.

### GetExperimentIDColumnOk

`func (o *AssignmentSourcePartialUpdateDto) GetExperimentIDColumnOk() (*string, bool)`

GetExperimentIDColumnOk returns a tuple with the ExperimentIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIDColumn

`func (o *AssignmentSourcePartialUpdateDto) SetExperimentIDColumn(v string)`

SetExperimentIDColumn sets ExperimentIDColumn field to given value.

### HasExperimentIDColumn

`func (o *AssignmentSourcePartialUpdateDto) HasExperimentIDColumn() bool`

HasExperimentIDColumn returns a boolean if a field has been set.

### GetGroupIDColumn

`func (o *AssignmentSourcePartialUpdateDto) GetGroupIDColumn() string`

GetGroupIDColumn returns the GroupIDColumn field if non-nil, zero value otherwise.

### GetGroupIDColumnOk

`func (o *AssignmentSourcePartialUpdateDto) GetGroupIDColumnOk() (*string, bool)`

GetGroupIDColumnOk returns a tuple with the GroupIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIDColumn

`func (o *AssignmentSourcePartialUpdateDto) SetGroupIDColumn(v string)`

SetGroupIDColumn sets GroupIDColumn field to given value.

### HasGroupIDColumn

`func (o *AssignmentSourcePartialUpdateDto) HasGroupIDColumn() bool`

HasGroupIDColumn returns a boolean if a field has been set.

### GetIdTypeMapping

`func (o *AssignmentSourcePartialUpdateDto) GetIdTypeMapping() []AssignmentSourceCreationDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *AssignmentSourcePartialUpdateDto) GetIdTypeMappingOk() (*[]AssignmentSourceCreationDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *AssignmentSourcePartialUpdateDto) SetIdTypeMapping(v []AssignmentSourceCreationDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.

### HasIdTypeMapping

`func (o *AssignmentSourcePartialUpdateDto) HasIdTypeMapping() bool`

HasIdTypeMapping returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *AssignmentSourcePartialUpdateDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *AssignmentSourcePartialUpdateDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *AssignmentSourcePartialUpdateDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *AssignmentSourcePartialUpdateDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


