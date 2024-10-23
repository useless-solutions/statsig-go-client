# AutotunePartialUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief summary of what the autotune is being used for. | [optional] 
**Variants** | Pointer to [**[]AutotuneCreateDtoVariantsInner**](AutotuneCreateDtoVariantsInner.md) | An array of Variant objects. | [optional] 
**SuccessEvent** | Pointer to **string** | The event you are trying to optimize for. | [optional] 
**SuccessEventValue** | Pointer to **string** | The value that should come with the event for it to be considered successful. | [optional] 
**ExplorationWindow** | Pointer to **string** | The initial time period where Autotune will equally split the traffic. | [optional] 
**AttributionWindow** | Pointer to **string** | The maximum duration between the exposure and success event that counts as a success. | [optional] 
**WinnerThreshold** | Pointer to **string** | The \&quot;probability of best\&quot; threshold a variant needs to achieve for Autotune to declare it the winner, stop collecting data, and direct all traffic. | [optional] 
**MetadataField** | Pointer to **string** | Metadata field containing the numeric value to optimize for. If this field is null, autotune optimizes for the existence of a follow-up event. This is only used for contextual autotunes. | [optional] 
**HigherIsBetter** | Pointer to **bool** | Whether to optimize for an increase or decrease in the metadata field value. Default is true. This is only used for contextual autotunes. | [optional] 

## Methods

### NewAutotunePartialUpdateDto

`func NewAutotunePartialUpdateDto() *AutotunePartialUpdateDto`

NewAutotunePartialUpdateDto instantiates a new AutotunePartialUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotunePartialUpdateDtoWithDefaults

`func NewAutotunePartialUpdateDtoWithDefaults() *AutotunePartialUpdateDto`

NewAutotunePartialUpdateDtoWithDefaults instantiates a new AutotunePartialUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AutotunePartialUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutotunePartialUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutotunePartialUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutotunePartialUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariants

`func (o *AutotunePartialUpdateDto) GetVariants() []AutotuneCreateDtoVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *AutotunePartialUpdateDto) GetVariantsOk() (*[]AutotuneCreateDtoVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *AutotunePartialUpdateDto) SetVariants(v []AutotuneCreateDtoVariantsInner)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *AutotunePartialUpdateDto) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetSuccessEvent

`func (o *AutotunePartialUpdateDto) GetSuccessEvent() string`

GetSuccessEvent returns the SuccessEvent field if non-nil, zero value otherwise.

### GetSuccessEventOk

`func (o *AutotunePartialUpdateDto) GetSuccessEventOk() (*string, bool)`

GetSuccessEventOk returns a tuple with the SuccessEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEvent

`func (o *AutotunePartialUpdateDto) SetSuccessEvent(v string)`

SetSuccessEvent sets SuccessEvent field to given value.

### HasSuccessEvent

`func (o *AutotunePartialUpdateDto) HasSuccessEvent() bool`

HasSuccessEvent returns a boolean if a field has been set.

### GetSuccessEventValue

`func (o *AutotunePartialUpdateDto) GetSuccessEventValue() string`

GetSuccessEventValue returns the SuccessEventValue field if non-nil, zero value otherwise.

### GetSuccessEventValueOk

`func (o *AutotunePartialUpdateDto) GetSuccessEventValueOk() (*string, bool)`

GetSuccessEventValueOk returns a tuple with the SuccessEventValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEventValue

`func (o *AutotunePartialUpdateDto) SetSuccessEventValue(v string)`

SetSuccessEventValue sets SuccessEventValue field to given value.

### HasSuccessEventValue

`func (o *AutotunePartialUpdateDto) HasSuccessEventValue() bool`

HasSuccessEventValue returns a boolean if a field has been set.

### GetExplorationWindow

`func (o *AutotunePartialUpdateDto) GetExplorationWindow() string`

GetExplorationWindow returns the ExplorationWindow field if non-nil, zero value otherwise.

### GetExplorationWindowOk

`func (o *AutotunePartialUpdateDto) GetExplorationWindowOk() (*string, bool)`

GetExplorationWindowOk returns a tuple with the ExplorationWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorationWindow

`func (o *AutotunePartialUpdateDto) SetExplorationWindow(v string)`

SetExplorationWindow sets ExplorationWindow field to given value.

### HasExplorationWindow

`func (o *AutotunePartialUpdateDto) HasExplorationWindow() bool`

HasExplorationWindow returns a boolean if a field has been set.

### GetAttributionWindow

`func (o *AutotunePartialUpdateDto) GetAttributionWindow() string`

GetAttributionWindow returns the AttributionWindow field if non-nil, zero value otherwise.

### GetAttributionWindowOk

`func (o *AutotunePartialUpdateDto) GetAttributionWindowOk() (*string, bool)`

GetAttributionWindowOk returns a tuple with the AttributionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionWindow

`func (o *AutotunePartialUpdateDto) SetAttributionWindow(v string)`

SetAttributionWindow sets AttributionWindow field to given value.

### HasAttributionWindow

`func (o *AutotunePartialUpdateDto) HasAttributionWindow() bool`

HasAttributionWindow returns a boolean if a field has been set.

### GetWinnerThreshold

`func (o *AutotunePartialUpdateDto) GetWinnerThreshold() string`

GetWinnerThreshold returns the WinnerThreshold field if non-nil, zero value otherwise.

### GetWinnerThresholdOk

`func (o *AutotunePartialUpdateDto) GetWinnerThresholdOk() (*string, bool)`

GetWinnerThresholdOk returns a tuple with the WinnerThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerThreshold

`func (o *AutotunePartialUpdateDto) SetWinnerThreshold(v string)`

SetWinnerThreshold sets WinnerThreshold field to given value.

### HasWinnerThreshold

`func (o *AutotunePartialUpdateDto) HasWinnerThreshold() bool`

HasWinnerThreshold returns a boolean if a field has been set.

### GetMetadataField

`func (o *AutotunePartialUpdateDto) GetMetadataField() string`

GetMetadataField returns the MetadataField field if non-nil, zero value otherwise.

### GetMetadataFieldOk

`func (o *AutotunePartialUpdateDto) GetMetadataFieldOk() (*string, bool)`

GetMetadataFieldOk returns a tuple with the MetadataField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataField

`func (o *AutotunePartialUpdateDto) SetMetadataField(v string)`

SetMetadataField sets MetadataField field to given value.

### HasMetadataField

`func (o *AutotunePartialUpdateDto) HasMetadataField() bool`

HasMetadataField returns a boolean if a field has been set.

### GetHigherIsBetter

`func (o *AutotunePartialUpdateDto) GetHigherIsBetter() bool`

GetHigherIsBetter returns the HigherIsBetter field if non-nil, zero value otherwise.

### GetHigherIsBetterOk

`func (o *AutotunePartialUpdateDto) GetHigherIsBetterOk() (*bool, bool)`

GetHigherIsBetterOk returns a tuple with the HigherIsBetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherIsBetter

`func (o *AutotunePartialUpdateDto) SetHigherIsBetter(v bool)`

SetHigherIsBetter sets HigherIsBetter field to given value.

### HasHigherIsBetter

`func (o *AutotunePartialUpdateDto) HasHigherIsBetter() bool`

HasHigherIsBetter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


