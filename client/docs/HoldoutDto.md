# HoldoutDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID | 
**Name** | Pointer to **string** | Optional name for the configuration. | [optional] 
**IdType** | **string** | type of id | 
**Description** | **string** | brief summary of what the holdout is being used for | 
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
**IsEnabled** | **bool** | enable or disable the holdout | 
**PassPercentage** | **float32** | percentage of users between 0-10% to pass through the holdout | 
**GateIDs** | **[]string** | an array of gateIDs which this holdout is applied to | 
**ExperimentIDs** | **[]string** | an array of experimentIDs which this holdout is applied to | 
**LayerIDs** | **[]string** | an array of layerIDs which this holdout is applied to | 
**IsGlobal** | **bool** | whether the holdout is being applied to all new gates | 
**TargetingGateID** | [**nil**](nil.md) | the gateID that the holdout is targeting | 

## Methods

### NewHoldoutDto

`func NewHoldoutDto(id string, idType string, description string, lastModifierID nil, lastModifiedTime nil, lastModifierEmail nil, lastModifierName nil, creatorID nil, createdTime float32, creatorName nil, creatorEmail nil, isEnabled bool, passPercentage float32, gateIDs []string, experimentIDs []string, layerIDs []string, isGlobal bool, targetingGateID nil, ) *HoldoutDto`

NewHoldoutDto instantiates a new HoldoutDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutDtoWithDefaults

`func NewHoldoutDtoWithDefaults() *HoldoutDto`

NewHoldoutDtoWithDefaults instantiates a new HoldoutDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HoldoutDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HoldoutDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HoldoutDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HoldoutDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HoldoutDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HoldoutDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HoldoutDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdType

`func (o *HoldoutDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *HoldoutDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *HoldoutDto) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetDescription

`func (o *HoldoutDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldoutDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldoutDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLastModifierID

`func (o *HoldoutDto) GetLastModifierID() nil`

GetLastModifierID returns the LastModifierID field if non-nil, zero value otherwise.

### GetLastModifierIDOk

`func (o *HoldoutDto) GetLastModifierIDOk() (*nil, bool)`

GetLastModifierIDOk returns a tuple with the LastModifierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierID

`func (o *HoldoutDto) SetLastModifierID(v nil)`

SetLastModifierID sets LastModifierID field to given value.


### GetLastModifiedTime

`func (o *HoldoutDto) GetLastModifiedTime() nil`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *HoldoutDto) GetLastModifiedTimeOk() (*nil, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *HoldoutDto) SetLastModifiedTime(v nil)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastModifierEmail

`func (o *HoldoutDto) GetLastModifierEmail() nil`

GetLastModifierEmail returns the LastModifierEmail field if non-nil, zero value otherwise.

### GetLastModifierEmailOk

`func (o *HoldoutDto) GetLastModifierEmailOk() (*nil, bool)`

GetLastModifierEmailOk returns a tuple with the LastModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierEmail

`func (o *HoldoutDto) SetLastModifierEmail(v nil)`

SetLastModifierEmail sets LastModifierEmail field to given value.


### GetLastModifierName

`func (o *HoldoutDto) GetLastModifierName() nil`

GetLastModifierName returns the LastModifierName field if non-nil, zero value otherwise.

### GetLastModifierNameOk

`func (o *HoldoutDto) GetLastModifierNameOk() (*nil, bool)`

GetLastModifierNameOk returns a tuple with the LastModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierName

`func (o *HoldoutDto) SetLastModifierName(v nil)`

SetLastModifierName sets LastModifierName field to given value.


### GetCreatorID

`func (o *HoldoutDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *HoldoutDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *HoldoutDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.


### GetCreatedTime

`func (o *HoldoutDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *HoldoutDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *HoldoutDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreatorName

`func (o *HoldoutDto) GetCreatorName() nil`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *HoldoutDto) GetCreatorNameOk() (*nil, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *HoldoutDto) SetCreatorName(v nil)`

SetCreatorName sets CreatorName field to given value.


### GetCreatorEmail

`func (o *HoldoutDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *HoldoutDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *HoldoutDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetTags

`func (o *HoldoutDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HoldoutDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HoldoutDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HoldoutDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetApps

`func (o *HoldoutDto) GetTargetApps() []string`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *HoldoutDto) GetTargetAppsOk() (*[]string, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *HoldoutDto) SetTargetApps(v []string)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *HoldoutDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetHoldoutIDs

`func (o *HoldoutDto) GetHoldoutIDs() []string`

GetHoldoutIDs returns the HoldoutIDs field if non-nil, zero value otherwise.

### GetHoldoutIDsOk

`func (o *HoldoutDto) GetHoldoutIDsOk() (*[]string, bool)`

GetHoldoutIDsOk returns a tuple with the HoldoutIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutIDs

`func (o *HoldoutDto) SetHoldoutIDs(v []string)`

SetHoldoutIDs sets HoldoutIDs field to given value.

### HasHoldoutIDs

`func (o *HoldoutDto) HasHoldoutIDs() bool`

HasHoldoutIDs returns a boolean if a field has been set.

### GetTeam

`func (o *HoldoutDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *HoldoutDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *HoldoutDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *HoldoutDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetVersion

`func (o *HoldoutDto) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HoldoutDto) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HoldoutDto) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HoldoutDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsEnabled

`func (o *HoldoutDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *HoldoutDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *HoldoutDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetPassPercentage

`func (o *HoldoutDto) GetPassPercentage() float32`

GetPassPercentage returns the PassPercentage field if non-nil, zero value otherwise.

### GetPassPercentageOk

`func (o *HoldoutDto) GetPassPercentageOk() (*float32, bool)`

GetPassPercentageOk returns a tuple with the PassPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercentage

`func (o *HoldoutDto) SetPassPercentage(v float32)`

SetPassPercentage sets PassPercentage field to given value.


### GetGateIDs

`func (o *HoldoutDto) GetGateIDs() []string`

GetGateIDs returns the GateIDs field if non-nil, zero value otherwise.

### GetGateIDsOk

`func (o *HoldoutDto) GetGateIDsOk() (*[]string, bool)`

GetGateIDsOk returns a tuple with the GateIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateIDs

`func (o *HoldoutDto) SetGateIDs(v []string)`

SetGateIDs sets GateIDs field to given value.


### GetExperimentIDs

`func (o *HoldoutDto) GetExperimentIDs() []string`

GetExperimentIDs returns the ExperimentIDs field if non-nil, zero value otherwise.

### GetExperimentIDsOk

`func (o *HoldoutDto) GetExperimentIDsOk() (*[]string, bool)`

GetExperimentIDsOk returns a tuple with the ExperimentIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIDs

`func (o *HoldoutDto) SetExperimentIDs(v []string)`

SetExperimentIDs sets ExperimentIDs field to given value.


### GetLayerIDs

`func (o *HoldoutDto) GetLayerIDs() []string`

GetLayerIDs returns the LayerIDs field if non-nil, zero value otherwise.

### GetLayerIDsOk

`func (o *HoldoutDto) GetLayerIDsOk() (*[]string, bool)`

GetLayerIDsOk returns a tuple with the LayerIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerIDs

`func (o *HoldoutDto) SetLayerIDs(v []string)`

SetLayerIDs sets LayerIDs field to given value.


### GetIsGlobal

`func (o *HoldoutDto) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *HoldoutDto) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *HoldoutDto) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.


### GetTargetingGateID

`func (o *HoldoutDto) GetTargetingGateID() nil`

GetTargetingGateID returns the TargetingGateID field if non-nil, zero value otherwise.

### GetTargetingGateIDOk

`func (o *HoldoutDto) GetTargetingGateIDOk() (*nil, bool)`

GetTargetingGateIDOk returns a tuple with the TargetingGateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingGateID

`func (o *HoldoutDto) SetTargetingGateID(v nil)`

SetTargetingGateID sets TargetingGateID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


