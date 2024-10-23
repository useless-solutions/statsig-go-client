# MultipleEventCountDtoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Events** | [**[]MultipleEventCountDtoInnerEventsInner**](MultipleEventCountDtoInnerEventsInner.md) |  | 

## Methods

### NewMultipleEventCountDtoInner

`func NewMultipleEventCountDtoInner(date string, events []MultipleEventCountDtoInnerEventsInner, ) *MultipleEventCountDtoInner`

NewMultipleEventCountDtoInner instantiates a new MultipleEventCountDtoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleEventCountDtoInnerWithDefaults

`func NewMultipleEventCountDtoInnerWithDefaults() *MultipleEventCountDtoInner`

NewMultipleEventCountDtoInnerWithDefaults instantiates a new MultipleEventCountDtoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *MultipleEventCountDtoInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MultipleEventCountDtoInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MultipleEventCountDtoInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetEvents

`func (o *MultipleEventCountDtoInner) GetEvents() []MultipleEventCountDtoInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MultipleEventCountDtoInner) GetEventsOk() (*[]MultipleEventCountDtoInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MultipleEventCountDtoInner) SetEvents(v []MultipleEventCountDtoInnerEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


