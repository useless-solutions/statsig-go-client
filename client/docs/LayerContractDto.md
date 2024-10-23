# LayerContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID | 
**Name** | Pointer to **string** | Optional name for the configuration. | [optional] 
**IdType** | **string** | The ID type used for this layer, important for validation. | 
**Description** | **string** | A detailed description of the layer, explaining its purpose and functionality. | 
**LastModifierID** | [**nil**](nil.md) | ID of the last modifier. | 
**LastModifiedTime** | [**nil**](nil.md) | Time of the last modification. | 
**LastModifierEmail** | [**nil**](nil.md) | Email of the last modifier. | 
**LastModifierName** | [**nil**](nil.md) | Name of the last modifier. | 
**CreatorID** | [**nil**](nil.md) | ID of the user who created the entity. | 
**CreatedTime** | **float32** | Timestamp when the entity was created. | 
**CreatorName** | [**nil**](nil.md) | Name of the creator. | 
**CreatorEmail** | [**nil**](nil.md) | Email of the creator. | 
**Tags** | Pointer to **[]string** | Optional tags for categorization. | [optional] 
**TargetApps** | Pointer to [**LayerContractDtoTargetApps**](LayerContractDtoTargetApps.md) |  | [optional] 
**HoldoutIDs** | Pointer to **[]string** | Holdouts applied to this configuration. | [optional] 
**Team** | Pointer to [**nil**](nil.md) | Optional identifier for the responsible team. | [optional] 
**Version** | Pointer to **float32** | Version number | [optional] 
**IsImplicitLayer** | **bool** | Indicates if the layer was automatically created by Statsig during experiment creation. | 
**Parameters** | [**[]LayerContractDtoParametersInner**](LayerContractDtoParametersInner.md) | An array of parameters associated with the layer, each defining specific attributes. | 

## Methods

### NewLayerContractDto

`func NewLayerContractDto(id string, idType string, description string, lastModifierID nil, lastModifiedTime nil, lastModifierEmail nil, lastModifierName nil, creatorID nil, createdTime float32, creatorName nil, creatorEmail nil, isImplicitLayer bool, parameters []LayerContractDtoParametersInner, ) *LayerContractDto`

NewLayerContractDto instantiates a new LayerContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerContractDtoWithDefaults

`func NewLayerContractDtoWithDefaults() *LayerContractDto`

NewLayerContractDtoWithDefaults instantiates a new LayerContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LayerContractDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LayerContractDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LayerContractDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LayerContractDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayerContractDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayerContractDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LayerContractDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdType

`func (o *LayerContractDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *LayerContractDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *LayerContractDto) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetDescription

`func (o *LayerContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LayerContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LayerContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLastModifierID

`func (o *LayerContractDto) GetLastModifierID() nil`

GetLastModifierID returns the LastModifierID field if non-nil, zero value otherwise.

### GetLastModifierIDOk

`func (o *LayerContractDto) GetLastModifierIDOk() (*nil, bool)`

GetLastModifierIDOk returns a tuple with the LastModifierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierID

`func (o *LayerContractDto) SetLastModifierID(v nil)`

SetLastModifierID sets LastModifierID field to given value.


### GetLastModifiedTime

`func (o *LayerContractDto) GetLastModifiedTime() nil`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *LayerContractDto) GetLastModifiedTimeOk() (*nil, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *LayerContractDto) SetLastModifiedTime(v nil)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastModifierEmail

`func (o *LayerContractDto) GetLastModifierEmail() nil`

GetLastModifierEmail returns the LastModifierEmail field if non-nil, zero value otherwise.

### GetLastModifierEmailOk

`func (o *LayerContractDto) GetLastModifierEmailOk() (*nil, bool)`

GetLastModifierEmailOk returns a tuple with the LastModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierEmail

`func (o *LayerContractDto) SetLastModifierEmail(v nil)`

SetLastModifierEmail sets LastModifierEmail field to given value.


### GetLastModifierName

`func (o *LayerContractDto) GetLastModifierName() nil`

GetLastModifierName returns the LastModifierName field if non-nil, zero value otherwise.

### GetLastModifierNameOk

`func (o *LayerContractDto) GetLastModifierNameOk() (*nil, bool)`

GetLastModifierNameOk returns a tuple with the LastModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierName

`func (o *LayerContractDto) SetLastModifierName(v nil)`

SetLastModifierName sets LastModifierName field to given value.


### GetCreatorID

`func (o *LayerContractDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *LayerContractDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *LayerContractDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.


### GetCreatedTime

`func (o *LayerContractDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *LayerContractDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *LayerContractDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreatorName

`func (o *LayerContractDto) GetCreatorName() nil`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *LayerContractDto) GetCreatorNameOk() (*nil, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *LayerContractDto) SetCreatorName(v nil)`

SetCreatorName sets CreatorName field to given value.


### GetCreatorEmail

`func (o *LayerContractDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *LayerContractDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *LayerContractDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetTags

`func (o *LayerContractDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LayerContractDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LayerContractDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LayerContractDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetApps

`func (o *LayerContractDto) GetTargetApps() LayerContractDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *LayerContractDto) GetTargetAppsOk() (*LayerContractDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *LayerContractDto) SetTargetApps(v LayerContractDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *LayerContractDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetHoldoutIDs

`func (o *LayerContractDto) GetHoldoutIDs() []string`

GetHoldoutIDs returns the HoldoutIDs field if non-nil, zero value otherwise.

### GetHoldoutIDsOk

`func (o *LayerContractDto) GetHoldoutIDsOk() (*[]string, bool)`

GetHoldoutIDsOk returns a tuple with the HoldoutIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutIDs

`func (o *LayerContractDto) SetHoldoutIDs(v []string)`

SetHoldoutIDs sets HoldoutIDs field to given value.

### HasHoldoutIDs

`func (o *LayerContractDto) HasHoldoutIDs() bool`

HasHoldoutIDs returns a boolean if a field has been set.

### GetTeam

`func (o *LayerContractDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *LayerContractDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *LayerContractDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *LayerContractDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetVersion

`func (o *LayerContractDto) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LayerContractDto) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LayerContractDto) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LayerContractDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsImplicitLayer

`func (o *LayerContractDto) GetIsImplicitLayer() bool`

GetIsImplicitLayer returns the IsImplicitLayer field if non-nil, zero value otherwise.

### GetIsImplicitLayerOk

`func (o *LayerContractDto) GetIsImplicitLayerOk() (*bool, bool)`

GetIsImplicitLayerOk returns a tuple with the IsImplicitLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImplicitLayer

`func (o *LayerContractDto) SetIsImplicitLayer(v bool)`

SetIsImplicitLayer sets IsImplicitLayer field to given value.


### GetParameters

`func (o *LayerContractDto) GetParameters() []LayerContractDtoParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LayerContractDto) GetParametersOk() (*[]LayerContractDtoParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LayerContractDto) SetParameters(v []LayerContractDtoParametersInner)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


