# AssignmentSourceContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique identifier for the assignment source. | 
**Description** | **string** | Detailed context and purpose of the assignment source. | 
**IsVerified** | Pointer to **bool** | Marks the assignment source as verified for internal trustworthiness. | [optional] 
**Tags** | **[]string** | Tags for categorization and search. | 
**Sql** | **string** | SQL query defining the data source for assignments. | 
**TimestampColumn** | **string** | Column name representing the timestamp of assignments. | 
**ExperimentIDColumn** | **string** | Column name for the experiment ID associated with the assignments. | 
**GroupIDColumn** | **string** | Column name for the group ID linked to the assignments. | 
**IdTypeMapping** | [**[]AssignmentSourceCreationDtoIdTypeMappingInner**](AssignmentSourceCreationDtoIdTypeMappingInner.md) | Mappings of Statsig units to their respective columns. | 
**IsReadOnly** | Pointer to **bool** | Specifies if the source can only be edited via the Console API. | [optional] 

## Methods

### NewAssignmentSourceContractDto

`func NewAssignmentSourceContractDto(name string, description string, tags []string, sql string, timestampColumn string, experimentIDColumn string, groupIDColumn string, idTypeMapping []AssignmentSourceCreationDtoIdTypeMappingInner, ) *AssignmentSourceContractDto`

NewAssignmentSourceContractDto instantiates a new AssignmentSourceContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentSourceContractDtoWithDefaults

`func NewAssignmentSourceContractDtoWithDefaults() *AssignmentSourceContractDto`

NewAssignmentSourceContractDtoWithDefaults instantiates a new AssignmentSourceContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssignmentSourceContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssignmentSourceContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssignmentSourceContractDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AssignmentSourceContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssignmentSourceContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssignmentSourceContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsVerified

`func (o *AssignmentSourceContractDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *AssignmentSourceContractDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *AssignmentSourceContractDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *AssignmentSourceContractDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetTags

`func (o *AssignmentSourceContractDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssignmentSourceContractDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssignmentSourceContractDto) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetSql

`func (o *AssignmentSourceContractDto) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *AssignmentSourceContractDto) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *AssignmentSourceContractDto) SetSql(v string)`

SetSql sets Sql field to given value.


### GetTimestampColumn

`func (o *AssignmentSourceContractDto) GetTimestampColumn() string`

GetTimestampColumn returns the TimestampColumn field if non-nil, zero value otherwise.

### GetTimestampColumnOk

`func (o *AssignmentSourceContractDto) GetTimestampColumnOk() (*string, bool)`

GetTimestampColumnOk returns a tuple with the TimestampColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampColumn

`func (o *AssignmentSourceContractDto) SetTimestampColumn(v string)`

SetTimestampColumn sets TimestampColumn field to given value.


### GetExperimentIDColumn

`func (o *AssignmentSourceContractDto) GetExperimentIDColumn() string`

GetExperimentIDColumn returns the ExperimentIDColumn field if non-nil, zero value otherwise.

### GetExperimentIDColumnOk

`func (o *AssignmentSourceContractDto) GetExperimentIDColumnOk() (*string, bool)`

GetExperimentIDColumnOk returns a tuple with the ExperimentIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIDColumn

`func (o *AssignmentSourceContractDto) SetExperimentIDColumn(v string)`

SetExperimentIDColumn sets ExperimentIDColumn field to given value.


### GetGroupIDColumn

`func (o *AssignmentSourceContractDto) GetGroupIDColumn() string`

GetGroupIDColumn returns the GroupIDColumn field if non-nil, zero value otherwise.

### GetGroupIDColumnOk

`func (o *AssignmentSourceContractDto) GetGroupIDColumnOk() (*string, bool)`

GetGroupIDColumnOk returns a tuple with the GroupIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIDColumn

`func (o *AssignmentSourceContractDto) SetGroupIDColumn(v string)`

SetGroupIDColumn sets GroupIDColumn field to given value.


### GetIdTypeMapping

`func (o *AssignmentSourceContractDto) GetIdTypeMapping() []AssignmentSourceCreationDtoIdTypeMappingInner`

GetIdTypeMapping returns the IdTypeMapping field if non-nil, zero value otherwise.

### GetIdTypeMappingOk

`func (o *AssignmentSourceContractDto) GetIdTypeMappingOk() (*[]AssignmentSourceCreationDtoIdTypeMappingInner, bool)`

GetIdTypeMappingOk returns a tuple with the IdTypeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeMapping

`func (o *AssignmentSourceContractDto) SetIdTypeMapping(v []AssignmentSourceCreationDtoIdTypeMappingInner)`

SetIdTypeMapping sets IdTypeMapping field to given value.


### GetIsReadOnly

`func (o *AssignmentSourceContractDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *AssignmentSourceContractDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *AssignmentSourceContractDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *AssignmentSourceContractDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


