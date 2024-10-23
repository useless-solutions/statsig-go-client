# TargetAppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of target app | [optional] 
**Name** | **string** | name of the target app | 

## Methods

### NewTargetAppDto

`func NewTargetAppDto(name string, ) *TargetAppDto`

NewTargetAppDto instantiates a new TargetAppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetAppDtoWithDefaults

`func NewTargetAppDtoWithDefaults() *TargetAppDto`

NewTargetAppDtoWithDefaults instantiates a new TargetAppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TargetAppDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetAppDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetAppDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TargetAppDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TargetAppDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetAppDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetAppDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


