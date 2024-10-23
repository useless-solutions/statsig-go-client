# EchidnaGateLoadPulseQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refresh** | Pointer to **string** |  | [optional] [default to "full"]
**RuleId** | **string** |  | 
**TurboMode** | Pointer to **bool** |  | [optional] 

## Methods

### NewEchidnaGateLoadPulseQueryDto

`func NewEchidnaGateLoadPulseQueryDto(ruleId string, ) *EchidnaGateLoadPulseQueryDto`

NewEchidnaGateLoadPulseQueryDto instantiates a new EchidnaGateLoadPulseQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEchidnaGateLoadPulseQueryDtoWithDefaults

`func NewEchidnaGateLoadPulseQueryDtoWithDefaults() *EchidnaGateLoadPulseQueryDto`

NewEchidnaGateLoadPulseQueryDtoWithDefaults instantiates a new EchidnaGateLoadPulseQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefresh

`func (o *EchidnaGateLoadPulseQueryDto) GetRefresh() string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *EchidnaGateLoadPulseQueryDto) GetRefreshOk() (*string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *EchidnaGateLoadPulseQueryDto) SetRefresh(v string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *EchidnaGateLoadPulseQueryDto) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRuleId

`func (o *EchidnaGateLoadPulseQueryDto) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EchidnaGateLoadPulseQueryDto) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EchidnaGateLoadPulseQueryDto) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetTurboMode

`func (o *EchidnaGateLoadPulseQueryDto) GetTurboMode() bool`

GetTurboMode returns the TurboMode field if non-nil, zero value otherwise.

### GetTurboModeOk

`func (o *EchidnaGateLoadPulseQueryDto) GetTurboModeOk() (*bool, bool)`

GetTurboModeOk returns a tuple with the TurboMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurboMode

`func (o *EchidnaGateLoadPulseQueryDto) SetTurboMode(v bool)`

SetTurboMode sets TurboMode field to given value.

### HasTurboMode

`func (o *EchidnaGateLoadPulseQueryDto) HasTurboMode() bool`

HasTurboMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


