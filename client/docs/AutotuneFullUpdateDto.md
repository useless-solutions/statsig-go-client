# AutotuneFullUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief summary of what the autotune is being used for. | [optional] 
**Variants** | [**[]AutotuneCreateDtoVariantsInner**](AutotuneCreateDtoVariantsInner.md) | An array of Variant objects. | 
**SuccessEvent** | **string** | The event you are trying to optimize for. | 
**SuccessEventValue** | Pointer to **string** | The value that should come with the event for it to be considered successful. | [optional] 
**ExplorationWindow** | **string** | The initial time period where Autotune will equally split the traffic. | 
**AttributionWindow** | **string** | The maximum duration between the exposure and success event that counts as a success. | 
**WinnerThreshold** | **string** | The \&quot;probability of best\&quot; threshold a variant needs to achieve for Autotune to declare it the winner, stop collecting data, and direct all traffic. | 
**MetadataField** | Pointer to **string** | Metadata field containing the numeric value to optimize for. If this field is null, autotune optimizes for the existence of a follow-up event. This is only used for contextual autotunes. | [optional] 
**HigherIsBetter** | Pointer to **bool** | Whether to optimize for an increase or decrease in the metadata field value. Default is true. This is only used for contextual autotunes. | [optional] 

## Methods

### NewAutotuneFullUpdateDto

`func NewAutotuneFullUpdateDto(variants []AutotuneCreateDtoVariantsInner, successEvent string, explorationWindow string, attributionWindow string, winnerThreshold string, ) *AutotuneFullUpdateDto`

NewAutotuneFullUpdateDto instantiates a new AutotuneFullUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotuneFullUpdateDtoWithDefaults

`func NewAutotuneFullUpdateDtoWithDefaults() *AutotuneFullUpdateDto`

NewAutotuneFullUpdateDtoWithDefaults instantiates a new AutotuneFullUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AutotuneFullUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutotuneFullUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutotuneFullUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutotuneFullUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariants

`func (o *AutotuneFullUpdateDto) GetVariants() []AutotuneCreateDtoVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *AutotuneFullUpdateDto) GetVariantsOk() (*[]AutotuneCreateDtoVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *AutotuneFullUpdateDto) SetVariants(v []AutotuneCreateDtoVariantsInner)`

SetVariants sets Variants field to given value.


### GetSuccessEvent

`func (o *AutotuneFullUpdateDto) GetSuccessEvent() string`

GetSuccessEvent returns the SuccessEvent field if non-nil, zero value otherwise.

### GetSuccessEventOk

`func (o *AutotuneFullUpdateDto) GetSuccessEventOk() (*string, bool)`

GetSuccessEventOk returns a tuple with the SuccessEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEvent

`func (o *AutotuneFullUpdateDto) SetSuccessEvent(v string)`

SetSuccessEvent sets SuccessEvent field to given value.


### GetSuccessEventValue

`func (o *AutotuneFullUpdateDto) GetSuccessEventValue() string`

GetSuccessEventValue returns the SuccessEventValue field if non-nil, zero value otherwise.

### GetSuccessEventValueOk

`func (o *AutotuneFullUpdateDto) GetSuccessEventValueOk() (*string, bool)`

GetSuccessEventValueOk returns a tuple with the SuccessEventValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEventValue

`func (o *AutotuneFullUpdateDto) SetSuccessEventValue(v string)`

SetSuccessEventValue sets SuccessEventValue field to given value.

### HasSuccessEventValue

`func (o *AutotuneFullUpdateDto) HasSuccessEventValue() bool`

HasSuccessEventValue returns a boolean if a field has been set.

### GetExplorationWindow

`func (o *AutotuneFullUpdateDto) GetExplorationWindow() string`

GetExplorationWindow returns the ExplorationWindow field if non-nil, zero value otherwise.

### GetExplorationWindowOk

`func (o *AutotuneFullUpdateDto) GetExplorationWindowOk() (*string, bool)`

GetExplorationWindowOk returns a tuple with the ExplorationWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorationWindow

`func (o *AutotuneFullUpdateDto) SetExplorationWindow(v string)`

SetExplorationWindow sets ExplorationWindow field to given value.


### GetAttributionWindow

`func (o *AutotuneFullUpdateDto) GetAttributionWindow() string`

GetAttributionWindow returns the AttributionWindow field if non-nil, zero value otherwise.

### GetAttributionWindowOk

`func (o *AutotuneFullUpdateDto) GetAttributionWindowOk() (*string, bool)`

GetAttributionWindowOk returns a tuple with the AttributionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionWindow

`func (o *AutotuneFullUpdateDto) SetAttributionWindow(v string)`

SetAttributionWindow sets AttributionWindow field to given value.


### GetWinnerThreshold

`func (o *AutotuneFullUpdateDto) GetWinnerThreshold() string`

GetWinnerThreshold returns the WinnerThreshold field if non-nil, zero value otherwise.

### GetWinnerThresholdOk

`func (o *AutotuneFullUpdateDto) GetWinnerThresholdOk() (*string, bool)`

GetWinnerThresholdOk returns a tuple with the WinnerThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerThreshold

`func (o *AutotuneFullUpdateDto) SetWinnerThreshold(v string)`

SetWinnerThreshold sets WinnerThreshold field to given value.


### GetMetadataField

`func (o *AutotuneFullUpdateDto) GetMetadataField() string`

GetMetadataField returns the MetadataField field if non-nil, zero value otherwise.

### GetMetadataFieldOk

`func (o *AutotuneFullUpdateDto) GetMetadataFieldOk() (*string, bool)`

GetMetadataFieldOk returns a tuple with the MetadataField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataField

`func (o *AutotuneFullUpdateDto) SetMetadataField(v string)`

SetMetadataField sets MetadataField field to given value.

### HasMetadataField

`func (o *AutotuneFullUpdateDto) HasMetadataField() bool`

HasMetadataField returns a boolean if a field has been set.

### GetHigherIsBetter

`func (o *AutotuneFullUpdateDto) GetHigherIsBetter() bool`

GetHigherIsBetter returns the HigherIsBetter field if non-nil, zero value otherwise.

### GetHigherIsBetterOk

`func (o *AutotuneFullUpdateDto) GetHigherIsBetterOk() (*bool, bool)`

GetHigherIsBetterOk returns a tuple with the HigherIsBetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherIsBetter

`func (o *AutotuneFullUpdateDto) SetHigherIsBetter(v bool)`

SetHigherIsBetter sets HigherIsBetter field to given value.

### HasHigherIsBetter

`func (o *AutotuneFullUpdateDto) HasHigherIsBetter() bool`

HasHigherIsBetter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


