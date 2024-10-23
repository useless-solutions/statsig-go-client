# AssignmentSourceQueryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional new name for the assignment source. | [optional] 
**Description** | Pointer to **string** | Optional updated context for the assignment source. | [optional] 
**IsVerified** | Pointer to **bool** | Marks the assignment source as verified for internal trustworthiness. | [optional] 
**Tags** | Pointer to **[]string** | Optional updated tags for categorization. | [optional] 
**Sql** | **string** | SQL query defining the data source for assignments. | 
**TimestampColumn** | **string** | Column name representing the timestamp of assignments. | 
**ExperimentIDColumn** | **string** | Column name for the experiment ID associated with the assignments. | 
**GroupIDColumn** | **string** | Column name for the group ID linked to the assignments. | 
**IdTypeMapping** | [**[]AssignmentSourceCreationDtoIdTypeMappingInner**](AssignmentSourceCreationDtoIdTypeMappingInner.md) | Mappings of Statsig units to their respective columns. | 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewAssignmentSourceQueryUpdateDto

`func NewAssignmentSourceQueryUpdateDto(sql string, timestampColumn string, experimentIDColumn string, groupIDColumn string, idTypeMapping []AssignmentSourceCreationDtoIdTypeMappingInner, ) *AssignmentSourceQueryUpdateDto`

NewAssignmentSourceQueryUpdateDto instantiates a new AssignmentSourceQueryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentSourceQueryUpdateDtoWithDefaults

`func NewAssignmentSourceQueryUpdateDtoWithDefaults() *AssignmentSourceQueryUpdateDto`

NewAssignmentSourceQueryUpdateDtoWithDefaults instantiates a new AssignmentSourceQueryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssignmentSourceQueryUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssignmentSourceQueryUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssignmentSourceQueryUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssignmentSourceQueryUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AssignmentSourceQueryUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssignmentSourceQueryUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssignmentSourceQueryUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssignmentSourceQueryUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsVerified

`func (o *AssignmentSourceQueryUpdateDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *AssignmentSourceQueryUpdateDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *AssignmentSourceQueryUpdateDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *AssignmentSourceQueryUpdateDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetTags

`func (o *AssignmentSourceQueryUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssignmentSourceQueryUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssignmentSourceQueryUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssignmentSourceQueryUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSql

`func (o *AssignmentSourceQueryUpdateDto) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *AssignmentSourceQueryUpdateDto) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *AssignmentSourceQueryUpdateDto) SetSql(v string)`

SetSql sets Sql field to given value.


### GetTimestampColumn

`func (o *AssignmentSourceQueryUpdateDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *AssignmentSourceQueryUpdateDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *AssignmentSourceQueryUpdateDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.


### GetExperimentIDColumn

`func (o *AssignmentSourceQueryUpdateDto) GetExperimentIDColumn() string`

GetExperimentIDColumn returns the ExperimentIDColumn field if non-nil, zero value otherwise.

### GetExperimentIDColumnOk

`func (o *AssignmentSourceQueryUpdateDto) GetExperimentIDColumnOk() (*string, bool)`

GetExperimentIDColumnOk returns a tuple with the ExperimentIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIDColumn

`func (o *AssignmentSourceQueryUpdateDto) SetExperimentIDColumn(v string)`

SetExperimentIDColumn sets ExperimentIDColumn field to given value.


### GetGroupIDColumn

`func (o *AssignmentSourceQueryUpdateDto) GetGroupIDColumn() string`

GetGroupIDColumn returns the GroupIDColumn field if non-nil, zero value otherwise.

### GetGroupIDColumnOk

`func (o *AssignmentSourceQueryUpdateDto) GetGroupIDColumnOk() (*string, bool)`

GetGroupIDColumnOk returns a tuple with the GroupIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIDColumn

`func (o *AssignmentSourceQueryUpdateDto) SetGroupIDColumn(v string)`

SetGroupIDColumn sets GroupIDColumn field to given value.


### GetIdTypeMapping

`func (o *AssignmentSourceQueryUpdateDto) GetIdTypeMapping() []AssignmentSourceCreationDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *AssignmentSourceQueryUpdateDto) GetIdTypeMappingOk() (*[]AssignmentSourceCreationDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *AssignmentSourceQueryUpdateDto) SetIdTypeMapping(v []AssignmentSourceCreationDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.


### GetIsReadOnly

`func (o *AssignmentSourceQueryUpdateDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *AssignmentSourceQueryUpdateDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *AssignmentSourceQueryUpdateDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *AssignmentSourceQueryUpdateDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


