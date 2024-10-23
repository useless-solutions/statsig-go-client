# LoggedEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | The timestamp when the event occurred, represented as a string. | 
**Name** | **string** | The name of the event (e.g., \&quot;add_to_cart\&quot;). | 
**Source** | **string** | The source of the event, indicating where it was triggered from. | 
**Value** | **string** | The value associated with the event, providing additional context. | 
**UserID** | **string** | The ID of the user who triggered the event. | 

## Methods

### NewLoggedEventDto

`func NewLoggedEventDto(timestamp string, name string, source string, value string, userID string, ) *LoggedEventDto`

NewLoggedEventDto instantiates a new LoggedEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggedEventDtoWithDefaults

`func NewLoggedEventDtoWithDefaults() *LoggedEventDto`

NewLoggedEventDtoWithDefaults instantiates a new LoggedEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *LoggedEventDto) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LoggedEventDto) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LoggedEventDto) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetName

`func (o *LoggedEventDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggedEventDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggedEventDto) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *LoggedEventDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LoggedEventDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LoggedEventDto) SetSource(v string)`

SetSource sets Source field to given value.


### GetValue

`func (o *LoggedEventDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoggedEventDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoggedEventDto) SetValue(v string)`

SetValue sets Value field to given value.


### GetUserID

`func (o *LoggedEventDto) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *LoggedEventDto) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *LoggedEventDto) SetUserID(v string)`

SetUserID sets UserID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


