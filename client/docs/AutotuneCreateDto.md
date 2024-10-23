# AutotuneCreateDto

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
**Name** | **string** | The name that was originally given to the autotune on creation but formatted as an ID (\&quot;A Autotune\&quot; -&gt; \&quot;a_autotune\&quot;). | 
**IdType** | Pointer to **string** | idType of the autotune (userID, stableID, or a customID). Defaults to userID if not provided | [optional] 
**IsContextual** | Pointer to **bool** | Makes this autotune contextual | [optional] 

## Methods

### NewAutotuneCreateDto

`func NewAutotuneCreateDto(variants []AutotuneCreateDtoVariantsInner, successEvent string, explorationWindow string, attributionWindow string, winnerThreshold string, name string, ) *AutotuneCreateDto`

NewAutotuneCreateDto instantiates a new AutotuneCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotuneCreateDtoWithDefaults

`func NewAutotuneCreateDtoWithDefaults() *AutotuneCreateDto`

NewAutotuneCreateDtoWithDefaults instantiates a new AutotuneCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AutotuneCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutotuneCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutotuneCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutotuneCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariants

`func (o *AutotuneCreateDto) GetVariants() []AutotuneCreateDtoVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *AutotuneCreateDto) GetVariantsOk() (*[]AutotuneCreateDtoVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *AutotuneCreateDto) SetVariants(v []AutotuneCreateDtoVariantsInner)`

SetVariants sets Variants field to given value.


### GetSuccessEvent

`func (o *AutotuneCreateDto) GetSuccessEvent() string`

GetSuccessEvent returns the SuccessEvent field if non-nil, zero value otherwise.

### GetSuccessEventOk

`func (o *AutotuneCreateDto) GetSuccessEventOk() (*string, bool)`

GetSuccessEventOk returns a tuple with the SuccessEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEvent

`func (o *AutotuneCreateDto) SetSuccessEvent(v string)`

SetSuccessEvent sets SuccessEvent field to given value.


### GetSuccessEventValue

`func (o *AutotuneCreateDto) GetSuccessEventValue() string`

GetSuccessEventValue returns the SuccessEventValue field if non-nil, zero value otherwise.

### GetSuccessEventValueOk

`func (o *AutotuneCreateDto) GetSuccessEventValueOk() (*string, bool)`

GetSuccessEventValueOk returns a tuple with the SuccessEventValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEventValue

`func (o *AutotuneCreateDto) SetSuccessEventValue(v string)`

SetSuccessEventValue sets SuccessEventValue field to given value.

### HasSuccessEventValue

`func (o *AutotuneCreateDto) HasSuccessEventValue() bool`

HasSuccessEventValue returns a boolean if a field has been set.

### GetExplorationWindow

`func (o *AutotuneCreateDto) GetExplorationWindow() string`

GetExplorationWindow returns the ExplorationWindow field if non-nil, zero value otherwise.

### GetExplorationWindowOk

`func (o *AutotuneCreateDto) GetExplorationWindowOk() (*string, bool)`

GetExplorationWindowOk returns a tuple with the ExplorationWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorationWindow

`func (o *AutotuneCreateDto) SetExplorationWindow(v string)`

SetExplorationWindow sets ExplorationWindow field to given value.


### GetAttributionWindow

`func (o *AutotuneCreateDto) GetAttributionWindow() string`

GetAttributionWindow returns the AttributionWindow field if non-nil, zero value otherwise.

### GetAttributionWindowOk

`func (o *AutotuneCreateDto) GetAttributionWindowOk() (*string, bool)`

GetAttributionWindowOk returns a tuple with the AttributionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionWindow

`func (o *AutotuneCreateDto) SetAttributionWindow(v string)`

SetAttributionWindow sets AttributionWindow field to given value.


### GetWinnerThreshold

`func (o *AutotuneCreateDto) GetWinnerThreshold() string`

GetWinnerThreshold returns the WinnerThreshold field if non-nil, zero value otherwise.

### GetWinnerThresholdOk

`func (o *AutotuneCreateDto) GetWinnerThresholdOk() (*string, bool)`

GetWinnerThresholdOk returns a tuple with the WinnerThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerThreshold

`func (o *AutotuneCreateDto) SetWinnerThreshold(v string)`

SetWinnerThreshold sets WinnerThreshold field to given value.


### GetMetadataField

`func (o *AutotuneCreateDto) GetMetadataField() string`

GetMetadataField returns the MetadataField field if non-nil, zero value otherwise.

### GetMetadataFieldOk

`func (o *AutotuneCreateDto) GetMetadataFieldOk() (*string, bool)`

GetMetadataFieldOk returns a tuple with the MetadataField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataField

`func (o *AutotuneCreateDto) SetMetadataField(v string)`

SetMetadataField sets MetadataField field to given value.

### HasMetadataField

`func (o *AutotuneCreateDto) HasMetadataField() bool`

HasMetadataField returns a boolean if a field has been set.

### GetHigherIsBetter

`func (o *AutotuneCreateDto) GetHigherIsBetter() bool`

GetHigherIsBetter returns the HigherIsBetter field if non-nil, zero value otherwise.

### GetHigherIsBetterOk

`func (o *AutotuneCreateDto) GetHigherIsBetterOk() (*bool, bool)`

GetHigherIsBetterOk returns a tuple with the HigherIsBetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherIsBetter

`func (o *AutotuneCreateDto) SetHigherIsBetter(v bool)`

SetHigherIsBetter sets HigherIsBetter field to given value.

### HasHigherIsBetter

`func (o *AutotuneCreateDto) HasHigherIsBetter() bool`

HasHigherIsBetter returns a boolean if a field has been set.

### GetName

`func (o *AutotuneCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotuneCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotuneCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetIdType

`func (o *AutotuneCreateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *AutotuneCreateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *AutotuneCreateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *AutotuneCreateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetIsContextual

`func (o *AutotuneCreateDto) GetIsContextual() bool`

GetIsContextual returns the IsContextual field if non-nil, zero value otherwise.

### GetIsContextualOk

`func (o *AutotuneCreateDto) GetIsContextualOk() (*bool, bool)`

GetIsContextualOk returns a tuple with the IsContextual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContextual

`func (o *AutotuneCreateDto) SetIsContextual(v bool)`

SetIsContextual sets IsContextual field to given value.

### HasIsContextual

`func (o *AutotuneCreateDto) HasIsContextual() bool`

HasIsContextual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


