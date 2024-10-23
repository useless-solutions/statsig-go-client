# AutotuneExperimentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Detailed description of the configuration’s purpose. | 
**Variants** | [**[]AutotuneExperimentDtoVariantsInner**](AutotuneExperimentDtoVariantsInner.md) |  | 
**SuccessEvent** | **string** | The event you are trying to optimize for. | 
**SuccessEventValue** | **string** | The value that should come with the event for it to be considered successful. | 
**ExplorationWindow** | **string** | The initial time period where Autotune will equally split the traffic. | 
**AttributionWindow** | **string** | The maximum duration between the exposure and success event that counts as a success. | 
**WinnerThreshold** | **string** | The \&quot;probability of best\&quot; threshold a variant needs to achieve for Autotune to declare it the winner, stop collecting data, and direct all traffic. | 
**MetadataField** | Pointer to **string** | Metadata field containing the numeric value to optimize for. If this field is null, autotune optimizes for the existence of a follow-up event. This is only used for contextual autotunes. | [optional] 
**HigherIsBetter** | Pointer to **bool** | Whether to optimize for an increase or decrease in the metadata field value. Default is true. This is only used for contextual autotunes. | [optional] 
**IsContextual** | Pointer to **bool** | Whether this is a contextual autotune | [optional] 
**Id** | **string** | ID | 
**Name** | Pointer to **string** | Optional name for the configuration. | [optional] 
**IdType** | **string** | Type of ID | 
**LastModifierID** | [**nil**](nil.md) | ID of the last modifier. | 
**LastModifiedTime** | [**nil**](nil.md) | Time of the last modification. | 
**LastModifierEmail** | [**nil**](nil.md) | Email of the last modifier. | 
**LastModifierName** | [**nil**](nil.md) | Name of the last modifier. | 
**CreatorID** | [**nil**](nil.md) | ID of the user who created the entity. | 
**CreatedTime** | **float32** | Timestamp when the entity was created. | 
**CreatorName** | [**nil**](nil.md) | Name of the creator. | 
**CreatorEmail** | [**nil**](nil.md) | Email of the creator. | 
**Tags** | Pointer to **[]string** | Optional tags for categorization. | [optional] 
**TargetApps** | Pointer to **[]string** | List of target applications associated with this configuration. | [optional] 
**HoldoutIDs** | Pointer to **[]string** | Holdouts applied to this configuration. | [optional] 
**Team** | Pointer to [**nil**](nil.md) | Optional identifier for the responsible team. | [optional] 
**Version** | Pointer to **float32** | Version number | [optional] 
**IsStarted** | **bool** | Is the autotune experiment currently running. | 
**Winner** | [**nil**](nil.md) |  | 

## Methods

### NewAutotuneExperimentDto

`func NewAutotuneExperimentDto(description string, variants []AutotuneExperimentDtoVariantsInner, successEvent string, successEventValue string, explorationWindow string, attributionWindow string, winnerThreshold string, id string, idType string, lastModifierID nil, lastModifiedTime nil, lastModifierEmail nil, lastModifierName nil, creatorID nil, createdTime float32, creatorName nil, creatorEmail nil, isStarted bool, winner nil, ) *AutotuneExperimentDto`

NewAutotuneExperimentDto instantiates a new AutotuneExperimentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotuneExperimentDtoWithDefaults

`func NewAutotuneExperimentDtoWithDefaults() *AutotuneExperimentDto`

NewAutotuneExperimentDtoWithDefaults instantiates a new AutotuneExperimentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AutotuneExperimentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutotuneExperimentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutotuneExperimentDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVariants

`func (o *AutotuneExperimentDto) GetVariants() []AutotuneExperimentDtoVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *AutotuneExperimentDto) GetVariantsOk() (*[]AutotuneExperimentDtoVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *AutotuneExperimentDto) SetVariants(v []AutotuneExperimentDtoVariantsInner)`

SetVariants sets Variants field to given value.


### GetSuccessEvent

`func (o *AutotuneExperimentDto) GetSuccessEvent() string`

GetSuccessEvent returns the SuccessEvent field if non-nil, zero value otherwise.

### GetSuccessEventOk

`func (o *AutotuneExperimentDto) GetSuccessEventOk() (*string, bool)`

GetSuccessEventOk returns a tuple with the SuccessEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEvent

`func (o *AutotuneExperimentDto) SetSuccessEvent(v string)`

SetSuccessEvent sets SuccessEvent field to given value.


### GetSuccessEventValue

`func (o *AutotuneExperimentDto) GetSuccessEventValue() string`

GetSuccessEventValue returns the SuccessEventValue field if non-nil, zero value otherwise.

### GetSuccessEventValueOk

`func (o *AutotuneExperimentDto) GetSuccessEventValueOk() (*string, bool)`

GetSuccessEventValueOk returns a tuple with the SuccessEventValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEventValue

`func (o *AutotuneExperimentDto) SetSuccessEventValue(v string)`

SetSuccessEventValue sets SuccessEventValue field to given value.


### GetExplorationWindow

`func (o *AutotuneExperimentDto) GetExplorationWindow() string`

GetExplorationWindow returns the ExplorationWindow field if non-nil, zero value otherwise.

### GetExplorationWindowOk

`func (o *AutotuneExperimentDto) GetExplorationWindowOk() (*string, bool)`

GetExplorationWindowOk returns a tuple with the ExplorationWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorationWindow

`func (o *AutotuneExperimentDto) SetExplorationWindow(v string)`

SetExplorationWindow sets ExplorationWindow field to given value.


### GetAttributionWindow

`func (o *AutotuneExperimentDto) GetAttributionWindow() string`

GetAttributionWindow returns the AttributionWindow field if non-nil, zero value otherwise.

### GetAttributionWindowOk

`func (o *AutotuneExperimentDto) GetAttributionWindowOk() (*string, bool)`

GetAttributionWindowOk returns a tuple with the AttributionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionWindow

`func (o *AutotuneExperimentDto) SetAttributionWindow(v string)`

SetAttributionWindow sets AttributionWindow field to given value.


### GetWinnerThreshold

`func (o *AutotuneExperimentDto) GetWinnerThreshold() string`

GetWinnerThreshold returns the WinnerThreshold field if non-nil, zero value otherwise.

### GetWinnerThresholdOk

`func (o *AutotuneExperimentDto) GetWinnerThresholdOk() (*string, bool)`

GetWinnerThresholdOk returns a tuple with the WinnerThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerThreshold

`func (o *AutotuneExperimentDto) SetWinnerThreshold(v string)`

SetWinnerThreshold sets WinnerThreshold field to given value.


### GetMetadataField

`func (o *AutotuneExperimentDto) GetMetadataField() string`

GetMetadataField returns the MetadataField field if non-nil, zero value otherwise.

### GetMetadataFieldOk

`func (o *AutotuneExperimentDto) GetMetadataFieldOk() (*string, bool)`

GetMetadataFieldOk returns a tuple with the MetadataField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataField

`func (o *AutotuneExperimentDto) SetMetadataField(v string)`

SetMetadataField sets MetadataField field to given value.

### HasMetadataField

`func (o *AutotuneExperimentDto) HasMetadataField() bool`

HasMetadataField returns a boolean if a field has been set.

### GetHigherIsBetter

`func (o *AutotuneExperimentDto) GetHigherIsBetter() bool`

GetHigherIsBetter returns the HigherIsBetter field if non-nil, zero value otherwise.

### GetHigherIsBetterOk

`func (o *AutotuneExperimentDto) GetHigherIsBetterOk() (*bool, bool)`

GetHigherIsBetterOk returns a tuple with the HigherIsBetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherIsBetter

`func (o *AutotuneExperimentDto) SetHigherIsBetter(v bool)`

SetHigherIsBetter sets HigherIsBetter field to given value.

### HasHigherIsBetter

`func (o *AutotuneExperimentDto) HasHigherIsBetter() bool`

HasHigherIsBetter returns a boolean if a field has been set.

### GetIsContextual

`func (o *AutotuneExperimentDto) GetIsContextual() bool`

GetIsContextual returns the IsContextual field if non-nil, zero value otherwise.

### GetIsContextualOk

`func (o *AutotuneExperimentDto) GetIsContextualOk() (*bool, bool)`

GetIsContextualOk returns a tuple with the IsContextual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContextual

`func (o *AutotuneExperimentDto) SetIsContextual(v bool)`

SetIsContextual sets IsContextual field to given value.

### HasIsContextual

`func (o *AutotuneExperimentDto) HasIsContextual() bool`

HasIsContextual returns a boolean if a field has been set.

### GetId

`func (o *AutotuneExperimentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutotuneExperimentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutotuneExperimentDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutotuneExperimentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotuneExperimentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotuneExperimentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutotuneExperimentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdType

`func (o *AutotuneExperimentDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *AutotuneExperimentDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *AutotuneExperimentDto) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetLastModifierID

`func (o *AutotuneExperimentDto) GetLastModifierID() nil`

GetLastModifierID returns the LastModifierID field if non-nil, zero value otherwise.

### GetLastModifierIDOk

`func (o *AutotuneExperimentDto) GetLastModifierIDOk() (*nil, bool)`

GetLastModifierIDOk returns a tuple with the LastModifierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierID

`func (o *AutotuneExperimentDto) SetLastModifierID(v nil)`

SetLastModifierID sets LastModifierID field to given value.


### GetLastModifiedTime

`func (o *AutotuneExperimentDto) GetLastModifiedTime() nil`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *AutotuneExperimentDto) GetLastModifiedTimeOk() (*nil, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *AutotuneExperimentDto) SetLastModifiedTime(v nil)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastModifierEmail

`func (o *AutotuneExperimentDto) GetLastModifierEmail() nil`

GetLastModifierEmail returns the LastModifierEmail field if non-nil, zero value otherwise.

### GetLastModifierEmailOk

`func (o *AutotuneExperimentDto) GetLastModifierEmailOk() (*nil, bool)`

GetLastModifierEmailOk returns a tuple with the LastModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierEmail

`func (o *AutotuneExperimentDto) SetLastModifierEmail(v nil)`

SetLastModifierEmail sets LastModifierEmail field to given value.


### GetLastModifierName

`func (o *AutotuneExperimentDto) GetLastModifierName() nil`

GetLastModifierName returns the LastModifierName field if non-nil, zero value otherwise.

### GetLastModifierNameOk

`func (o *AutotuneExperimentDto) GetLastModifierNameOk() (*nil, bool)`

GetLastModifierNameOk returns a tuple with the LastModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierName

`func (o *AutotuneExperimentDto) SetLastModifierName(v nil)`

SetLastModifierName sets LastModifierName field to given value.


### GetCreatorID

`func (o *AutotuneExperimentDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *AutotuneExperimentDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *AutotuneExperimentDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.


### GetCreatedTime

`func (o *AutotuneExperimentDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AutotuneExperimentDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AutotuneExperimentDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreatorName

`func (o *AutotuneExperimentDto) GetCreatorName() nil`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *AutotuneExperimentDto) GetCreatorNameOk() (*nil, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *AutotuneExperimentDto) SetCreatorName(v nil)`

SetCreatorName sets CreatorName field to given value.


### GetCreatorEmail

`func (o *AutotuneExperimentDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *AutotuneExperimentDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *AutotuneExperimentDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetTags

`func (o *AutotuneExperimentDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AutotuneExperimentDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AutotuneExperimentDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AutotuneExperimentDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetApps

`func (o *AutotuneExperimentDto) GetTargetApps() []string`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *AutotuneExperimentDto) GetTargetAppsOk() (*[]string, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *AutotuneExperimentDto) SetTargetApps(v []string)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *AutotuneExperimentDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetHoldoutIDs

`func (o *AutotuneExperimentDto) GetHoldoutIDs() []string`

GetHoldoutIDs returns the HoldoutIDs field if non-nil, zero value otherwise.

### GetHoldoutIDsOk

`func (o *AutotuneExperimentDto) GetHoldoutIDsOk() (*[]string, bool)`

GetHoldoutIDsOk returns a tuple with the HoldoutIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutIDs

`func (o *AutotuneExperimentDto) SetHoldoutIDs(v []string)`

SetHoldoutIDs sets HoldoutIDs field to given value.

### HasHoldoutIDs

`func (o *AutotuneExperimentDto) HasHoldoutIDs() bool`

HasHoldoutIDs returns a boolean if a field has been set.

### GetTeam

`func (o *AutotuneExperimentDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *AutotuneExperimentDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *AutotuneExperimentDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *AutotuneExperimentDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetVersion

`func (o *AutotuneExperimentDto) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AutotuneExperimentDto) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AutotuneExperimentDto) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AutotuneExperimentDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsStarted

`func (o *AutotuneExperimentDto) GetIsStarted() bool`

GetIsStarted returns the IsStarted field if non-nil, zero value otherwise.

### GetIsStartedOk

`func (o *AutotuneExperimentDto) GetIsStartedOk() (*bool, bool)`

GetIsStartedOk returns a tuple with the IsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarted

`func (o *AutotuneExperimentDto) SetIsStarted(v bool)`

SetIsStarted sets IsStarted field to given value.


### GetWinner

`func (o *AutotuneExperimentDto) GetWinner() nil`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *AutotuneExperimentDto) GetWinnerOk() (*nil, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *AutotuneExperimentDto) SetWinner(v nil)`

SetWinner sets Winner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


