# AuditLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the audit log | 
**Name** | **string** | name of the audit log | 
**ChangeLog** | **string** | change log of the audit log | 
**ActionType** | **map[string]interface{}** |  | 
**Date** | **string** |  | 
**Time** | **string** |  | 
**UpdatedBy** | **string** |  | 
**UpdatedByUserID** | **string** |  | 
**ModifierEmail** | **map[string]interface{}** |  | 
**Changes** | **map[string]interface{}** |  | 
**Tags** | **[]string** |  | 
**TargetAppIDs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAuditLogDto

`func NewAuditLogDto(id string, name string, changeLog string, actionType map[string]interface{}, date string, time string, updatedBy string, updatedByUserID string, modifierEmail map[string]interface{}, changes map[string]interface{}, tags []string, ) *AuditLogDto`

NewAuditLogDto instantiates a new AuditLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogDtoWithDefaults

`func NewAuditLogDtoWithDefaults() *AuditLogDto`

NewAuditLogDtoWithDefaults instantiates a new AuditLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditLogDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AuditLogDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogDto) SetName(v string)`

SetName sets Name field to given value.


### GetChangeLog

`func (o *AuditLogDto) GetChangeLog() string`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AuditLogDto) GetChangeLogOk() (*string, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AuditLogDto) SetChangeLog(v string)`

SetChangeLog sets ChangeLog field to given value.


### GetActionType

`func (o *AuditLogDto) GetActionType() map[string]interface{}`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AuditLogDto) GetActionTypeOk() (*map[string]interface{}, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AuditLogDto) SetActionType(v map[string]interface{})`

SetActionType sets ActionType field to given value.


### GetDate

`func (o *AuditLogDto) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuditLogDto) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuditLogDto) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *AuditLogDto) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AuditLogDto) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AuditLogDto) SetTime(v string)`

SetTime sets Time field to given value.


### GetUpdatedBy

`func (o *AuditLogDto) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AuditLogDto) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AuditLogDto) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetUpdatedByUserID

`func (o *AuditLogDto) GetUpdatedByUserID() string`

GetUpdatedByUserID returns the UpdatedByUserID field if non-nil, zero value otherwise.

### GetUpdatedByUserIDOk

`func (o *AuditLogDto) GetUpdatedByUserIDOk() (*string, bool)`

GetUpdatedByUserIDOk returns a tuple with the UpdatedByUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserID

`func (o *AuditLogDto) SetUpdatedByUserID(v string)`

SetUpdatedByUserID sets UpdatedByUserID field to given value.


### GetModifierEmail

`func (o *AuditLogDto) GetModifierEmail() map[string]interface{}`

GetModifierEmail returns the ModifierEmail field if non-nil, zero value otherwise.

### GetModifierEmailOk

`func (o *AuditLogDto) GetModifierEmailOk() (*map[string]interface{}, bool)`

GetModifierEmailOk returns a tuple with the ModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierEmail

`func (o *AuditLogDto) SetModifierEmail(v map[string]interface{})`

SetModifierEmail sets ModifierEmail field to given value.


### GetChanges

`func (o *AuditLogDto) GetChanges() map[string]interface{}`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *AuditLogDto) GetChangesOk() (*map[string]interface{}, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *AuditLogDto) SetChanges(v map[string]interface{})`

SetChanges sets Changes field to given value.


### GetTags

`func (o *AuditLogDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AuditLogDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AuditLogDto) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTargetAppIDs

`func (o *AuditLogDto) GetTargetAppIDs() []string`

GetTargetAppIDs returns the TargetAppIDs field if non-nil, zero value otherwise.

### GetTargetAppIDsOk

`func (o *AuditLogDto) GetTargetAppIDsOk() (*[]string, bool)`

GetTargetAppIDsOk returns a tuple with the TargetAppIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAppIDs

`func (o *AuditLogDto) SetTargetAppIDs(v []string)`

SetTargetAppIDs sets TargetAppIDs field to given value.

### HasTargetAppIDs

`func (o *AuditLogDto) HasTargetAppIDs() bool`

HasTargetAppIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


