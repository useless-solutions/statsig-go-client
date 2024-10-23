# MetricsUpdateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A detailed description of the metric, providing insights into its purpose and application. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags associated with the metric, used for categorization and easier retrieval. | [optional] 
**IsHidden** | Pointer to **bool** | Indicates whether the metric is hidden from general view, useful for internal metrics. | [optional] 
**IsVerified** | Pointer to **bool** | Flag to mark the metric as verified, ensuring it is deemed trustworthy within the organization. | [optional] 
**IsReadOnly** | Pointer to **bool** | Specifies if the metric definition can only be edited via the Console API, enhancing control over modifications. | [optional] 
**IsPermanent** | Pointer to **bool** | Determines if the metric is permanent, preventing it from being deleted or modified inadvertently. | [optional] 
**WarehouseNative** | Pointer to [**MetricsUpdateContractDtoWarehouseNative**](MetricsUpdateContractDtoWarehouseNative.md) |  | [optional] 
**UnitTypes** | Pointer to **[]string** | Array of unit types that the metric can utilize, such as stableID, userID, or other custom identifiers. | [optional] 
**Team** | Pointer to [**nil**](nil.md) | Optional field indicating the team responsible for the metric, aiding in accountability and management. | [optional] 

## Methods

### NewMetricsUpdateContractDto

`func NewMetricsUpdateContractDto() *MetricsUpdateContractDto`

NewMetricsUpdateContractDto instantiates a new MetricsUpdateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsUpdateContractDtoWithDefaults

`func NewMetricsUpdateContractDtoWithDefaults() *MetricsUpdateContractDto`

NewMetricsUpdateContractDtoWithDefaults instantiates a new MetricsUpdateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MetricsUpdateContractDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsUpdateContractDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsUpdateContractDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsUpdateContractDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *MetricsUpdateContractDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricsUpdateContractDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricsUpdateContractDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricsUpdateContractDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsHidden

`func (o *MetricsUpdateContractDto) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *MetricsUpdateContractDto) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *MetricsUpdateContractDto) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *MetricsUpdateContractDto) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsVerified

`func (o *MetricsUpdateContractDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *MetricsUpdateContractDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *MetricsUpdateContractDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *MetricsUpdateContractDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *MetricsUpdateContractDto) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MetricsUpdateContractDto) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *MetricsUpdateContractDto) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *MetricsUpdateContractDto) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetIsPermanent

`func (o *MetricsUpdateContractDto) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *MetricsUpdateContractDto) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *MetricsUpdateContractDto) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *MetricsUpdateContractDto) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetWarehouseNative

`func (o *MetricsUpdateContractDto) GetWarehouseNative() MetricsUpdateContractDtoWarehouseNative`

GetWarehouseNative returns the WarehouseNative field if non-nil, zero value otherwise.

### GetWarehouseNativeOk

`func (o *MetricsUpdateContractDto) GetWarehouseNativeOk() (*MetricsUpdateContractDtoWarehouseNative, bool)`

GetWarehouseNativeOk returns a tuple with the WarehouseNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseNative

`func (o *MetricsUpdateContractDto) SetWarehouseNative(v MetricsUpdateContractDtoWarehouseNative)`

SetWarehouseNative sets WarehouseNative field to given value.

### HasWarehouseNative

`func (o *MetricsUpdateContractDto) HasWarehouseNative() bool`

HasWarehouseNative returns a boolean if a field has been set.

### GetUnitTypes

`func (o *MetricsUpdateContractDto) GetUnitTypes() []string`

GetUnitTypes returns the UnitTypes field if non-nil, zero value otherwise.

### GetUnitTypesOk

`func (o *MetricsUpdateContractDto) GetUnitTypesOk() (*[]string, bool)`

GetUnitTypesOk returns a tuple with the UnitTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTypes

`func (o *MetricsUpdateContractDto) SetUnitTypes(v []string)`

SetUnitTypes sets UnitTypes field to given value.

### HasUnitTypes

`func (o *MetricsUpdateContractDto) HasUnitTypes() bool`

HasUnitTypes returns a boolean if a field has been set.

### GetTeam

`func (o *MetricsUpdateContractDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MetricsUpdateContractDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MetricsUpdateContractDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MetricsUpdateContractDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


