# SegmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Is the segment enabled. | 
**Type** | **string** |  | 
**Count** | Pointer to **float32** | For id_list segments: the length of the ID list | [optional] 
**Rules** | Pointer to [**[]ExternalGateDtoRulesInner**](ExternalGateDtoRulesInner.md) | Rule Object | [optional] 
**Tags** | Pointer to **[]string** | Optional tags for categorization. | [optional] 
**Id** | **string** | ID | 
**Name** | Pointer to **string** | Optional name for the configuration. | [optional] 
**IdType** | **string** | Type of ID | 
**Description** | **string** | Detailed description of the configuration’s purpose. | 
**LastModifierID** | [**nil**](nil.md) | ID of the last modifier. | 
**LastModifiedTime** | [**nil**](nil.md) | Time of the last modification. | 
**LastModifierEmail** | [**nil**](nil.md) | Email of the last modifier. | 
**LastModifierName** | [**nil**](nil.md) | Name of the last modifier. | 
**CreatorID** | [**nil**](nil.md) | ID of the user who created the entity. | 
**CreatedTime** | **float32** | Timestamp when the entity was created. | 
**CreatorName** | [**nil**](nil.md) | Name of the creator. | 
**CreatorEmail** | [**nil**](nil.md) | Email of the creator. | 
**TargetApps** | Pointer to **[]string** | List of target applications associated with this configuration. | [optional] 
**HoldoutIDs** | Pointer to **[]string** | Holdouts applied to this configuration. | [optional] 
**Team** | Pointer to [**nil**](nil.md) | Optional identifier for the responsible team. | [optional] 
**Version** | Pointer to **float32** | Version number | [optional] 

## Methods

### NewSegmentDto

`func NewSegmentDto(isEnabled bool, type_ string, id string, idType string, description string, lastModifierID nil, lastModifiedTime nil, lastModifierEmail nil, lastModifierName nil, creatorID nil, createdTime float32, creatorName nil, creatorEmail nil, ) *SegmentDto`

NewSegmentDto instantiates a new SegmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentDtoWithDefaults

`func NewSegmentDtoWithDefaults() *SegmentDto`

NewSegmentDtoWithDefaults instantiates a new SegmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SegmentDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SegmentDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SegmentDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetType

`func (o *SegmentDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SegmentDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SegmentDto) SetType(v string)`

SetType sets Type field to given value.


### GetCount

`func (o *SegmentDto) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SegmentDto) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SegmentDto) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SegmentDto) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRules

`func (o *SegmentDto) GetRules() []ExternalGateDtoRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SegmentDto) GetRulesOk() (*[]ExternalGateDtoRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SegmentDto) SetRules(v []ExternalGateDtoRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SegmentDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *SegmentDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SegmentDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SegmentDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SegmentDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetId

`func (o *SegmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SegmentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdType

`func (o *SegmentDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *SegmentDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *SegmentDto) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetDescription

`func (o *SegmentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLastModifierID

`func (o *SegmentDto) GetLastModifierID() nil`

GetLastModifierID returns the LastModifierID field if non-nil, zero value otherwise.

### GetLastModifierIDOk

`func (o *SegmentDto) GetLastModifierIDOk() (*nil, bool)`

GetLastModifierIDOk returns a tuple with the LastModifierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierID

`func (o *SegmentDto) SetLastModifierID(v nil)`

SetLastModifierID sets LastModifierID field to given value.


### GetLastModifiedTime

`func (o *SegmentDto) GetLastModifiedTime() nil`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *SegmentDto) GetLastModifiedTimeOk() (*nil, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *SegmentDto) SetLastModifiedTime(v nil)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastModifierEmail

`func (o *SegmentDto) GetLastModifierEmail() nil`

GetLastModifierEmail returns the LastModifierEmail field if non-nil, zero value otherwise.

### GetLastModifierEmailOk

`func (o *SegmentDto) GetLastModifierEmailOk() (*nil, bool)`

GetLastModifierEmailOk returns a tuple with the LastModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierEmail

`func (o *SegmentDto) SetLastModifierEmail(v nil)`

SetLastModifierEmail sets LastModifierEmail field to given value.


### GetLastModifierName

`func (o *SegmentDto) GetLastModifierName() nil`

GetLastModifierName returns the LastModifierName field if non-nil, zero value otherwise.

### GetLastModifierNameOk

`func (o *SegmentDto) GetLastModifierNameOk() (*nil, bool)`

GetLastModifierNameOk returns a tuple with the LastModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierName

`func (o *SegmentDto) SetLastModifierName(v nil)`

SetLastModifierName sets LastModifierName field to given value.


### GetCreatorID

`func (o *SegmentDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *SegmentDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *SegmentDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.


### GetCreatedTime

`func (o *SegmentDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SegmentDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SegmentDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreatorName

`func (o *SegmentDto) GetCreatorName() nil`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *SegmentDto) GetCreatorNameOk() (*nil, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *SegmentDto) SetCreatorName(v nil)`

SetCreatorName sets CreatorName field to given value.


### GetCreatorEmail

`func (o *SegmentDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SegmentDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SegmentDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetTargetApps

`func (o *SegmentDto) GetTargetApps() []string`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *SegmentDto) GetTargetAppsOk() (*[]string, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *SegmentDto) SetTargetApps(v []string)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *SegmentDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetHoldoutIDs

`func (o *SegmentDto) GetHoldoutIDs() []string`

GetHoldoutIDs returns the HoldoutIDs field if non-nil, zero value otherwise.

### GetHoldoutIDsOk

`func (o *SegmentDto) GetHoldoutIDsOk() (*[]string, bool)`

GetHoldoutIDsOk returns a tuple with the HoldoutIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutIDs

`func (o *SegmentDto) SetHoldoutIDs(v []string)`

SetHoldoutIDs sets HoldoutIDs field to given value.

### HasHoldoutIDs

`func (o *SegmentDto) HasHoldoutIDs() bool`

HasHoldoutIDs returns a boolean if a field has been set.

### GetTeam

`func (o *SegmentDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *SegmentDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *SegmentDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *SegmentDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetVersion

`func (o *SegmentDto) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SegmentDto) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SegmentDto) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SegmentDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


