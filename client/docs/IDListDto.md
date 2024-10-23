# IDListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name associated with the list of IDs | 
**Count** | **float32** | The count of IDs in the list | 
**Ids** | **[]string** | The array of IDs | 

## Methods

### NewIDListDto

`func NewIDListDto(name string, count float32, ids []string, ) *IDListDto`

NewIDListDto instantiates a new IDListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDListDtoWithDefaults

`func NewIDListDtoWithDefaults() *IDListDto`

NewIDListDtoWithDefaults instantiates a new IDListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IDListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IDListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IDListDto) SetName(v string)`

SetName sets Name field to given value.


### GetCount

`func (o *IDListDto) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IDListDto) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IDListDto) SetCount(v float32)`

SetCount sets Count field to given value.


### GetIds

`func (o *IDListDto) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *IDListDto) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *IDListDto) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


