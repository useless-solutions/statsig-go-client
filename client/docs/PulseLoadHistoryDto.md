# PulseLoadHistoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorName** | Pointer to **string** |  | [optional] 
**CreatedTime** | **float32** |  | 
**FinishedTime** | Pointer to **float32** |  | [optional] 
**FinishedState** | Pointer to **string** |  | [optional] 
**StartDs** | **string** |  | 
**EndDs** | **string** |  | 
**ReloadType** | **string** |  | 
**TurboMode** | **bool** |  | 

## Methods

### NewPulseLoadHistoryDto

`func NewPulseLoadHistoryDto(createdTime float32, startDs string, endDs string, reloadType string, turboMode bool, ) *PulseLoadHistoryDto`

NewPulseLoadHistoryDto instantiates a new PulseLoadHistoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulseLoadHistoryDtoWithDefaults

`func NewPulseLoadHistoryDtoWithDefaults() *PulseLoadHistoryDto`

NewPulseLoadHistoryDtoWithDefaults instantiates a new PulseLoadHistoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatorName

`func (o *PulseLoadHistoryDto) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *PulseLoadHistoryDto) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *PulseLoadHistoryDto) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *PulseLoadHistoryDto) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PulseLoadHistoryDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PulseLoadHistoryDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PulseLoadHistoryDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetFinishedTime

`func (o *PulseLoadHistoryDto) GetFinishedTime() float32`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *PulseLoadHistoryDto) GetFinishedTimeOk() (*float32, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *PulseLoadHistoryDto) SetFinishedTime(v float32)`

SetFinishedTime sets FinishedTime field to given value.

### HasFinishedTime

`func (o *PulseLoadHistoryDto) HasFinishedTime() bool`

HasFinishedTime returns a boolean if a field has been set.

### GetFinishedState

`func (o *PulseLoadHistoryDto) GetFinishedState() string`

GetFinishedState returns the FinishedState field if non-nil, zero value otherwise.

### GetFinishedStateOk

`func (o *PulseLoadHistoryDto) GetFinishedStateOk() (*string, bool)`

GetFinishedStateOk returns a tuple with the FinishedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedState

`func (o *PulseLoadHistoryDto) SetFinishedState(v string)`

SetFinishedState sets FinishedState field to given value.

### HasFinishedState

`func (o *PulseLoadHistoryDto) HasFinishedState() bool`

HasFinishedState returns a boolean if a field has been set.

### GetStartDs

`func (o *PulseLoadHistoryDto) GetStartDs() string`

GetStartDs returns the StartDs field if non-nil, zero value otherwise.

### GetStartDsOk

`func (o *PulseLoadHistoryDto) GetStartDsOk() (*string, bool)`

GetStartDsOk returns a tuple with the StartDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDs

`func (o *PulseLoadHistoryDto) SetStartDs(v string)`

SetStartDs sets StartDs field to given value.


### GetEndDs

`func (o *PulseLoadHistoryDto) GetEndDs() string`

GetEndDs returns the EndDs field if non-nil, zero value otherwise.

### GetEndDsOk

`func (o *PulseLoadHistoryDto) GetEndDsOk() (*string, bool)`

GetEndDsOk returns a tuple with the EndDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDs

`func (o *PulseLoadHistoryDto) SetEndDs(v string)`

SetEndDs sets EndDs field to given value.


### GetReloadType

`func (o *PulseLoadHistoryDto) GetReloadType() string`

GetReloadType returns the ReloadType field if non-nil, zero value otherwise.

### GetReloadTypeOk

`func (o *PulseLoadHistoryDto) GetReloadTypeOk() (*string, bool)`

GetReloadTypeOk returns a tuple with the ReloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloadType

`func (o *PulseLoadHistoryDto) SetReloadType(v string)`

SetReloadType sets ReloadType field to given value.


### GetTurboMode

`func (o *PulseLoadHistoryDto) GetTurboMode() bool`

GetTurboMode returns the TurboMode field if non-nil, zero value otherwise.

### GetTurboModeOk

`func (o *PulseLoadHistoryDto) GetTurboModeOk() (*bool, bool)`

GetTurboModeOk returns a tuple with the TurboMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurboMode

`func (o *PulseLoadHistoryDto) SetTurboMode(v bool)`

SetTurboMode sets TurboMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


