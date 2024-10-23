# TagUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsCore** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewTagUpdateDto

`func NewTagUpdateDto() *TagUpdateDto`

NewTagUpdateDto instantiates a new TagUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateDtoWithDefaults

`func NewTagUpdateDtoWithDefaults() *TagUpdateDto`

NewTagUpdateDtoWithDefaults instantiates a new TagUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TagUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsCore

`func (o *TagUpdateDto) GetIsCore() bool`

GetIsCore returns the IsCore field if non-nil, zero value otherwise.

### GetIsCoreOk

`func (o *TagUpdateDto) GetIsCoreOk() (*bool, bool)`

GetIsCoreOk returns a tuple with the IsCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCore

`func (o *TagUpdateDto) SetIsCore(v bool)`

SetIsCore sets IsCore field to given value.

### HasIsCore

`func (o *TagUpdateDto) HasIsCore() bool`

HasIsCore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


