# AlertSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the alert | 
**Name** | **string** | Name of the alert | 
**AlertType** | **string** | Type of alert | 
**Metrics** | **map[string]interface{}** | List of metrics associated with this alert | 
**MetricGroupBys** | **map[string]interface{}** | Metric groupbys | 
**Formula** | Pointer to **string** | Formula for the alert | [optional] 
**Message** | **string** | Alert message | 
**Instruction** | **string** | Instructions for resolving the alert | 
**CreatorID** | Pointer to **string** |  | [optional] 
**CompanyID** | **string** |  | 
**Priority** | **string** | Priority of this alert | 
**AlertThreshold** | **float32** |  | 
**WarningThreshold** | Pointer to **float32** |  | [optional] 
**WindowMs** | **float32** | How far back and how frequently a metric should be checked, in milliseconds | 
**Condition** | **string** | Condition under which a metric change triggers an alert in milliseconds | 
**RenotificationCondition** | Pointer to **string** | Condition under which a re-notification is sent | [optional] 
**RenotificationWindowMs** | Pointer to **float32** | How long to wait before re-notifying in milliseconds | [optional] 
**RenotificationMessage** | Pointer to **string** | Re-notification message | [optional] 
**Team** | Pointer to [**nil**](nil.md) | Team associated with this alert | [optional] 
**Tags** | **[]string** | Tags associated with this alert | 

## Methods

### NewAlertSchemaDto

`func NewAlertSchemaDto(id string, name string, alertType string, metrics map[string]interface{}, metricGroupBys map[string]interface{}, message string, instruction string, companyID string, priority string, alertThreshold float32, windowMs float32, condition string, tags []string, ) *AlertSchemaDto`

NewAlertSchemaDto instantiates a new AlertSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertSchemaDtoWithDefaults

`func NewAlertSchemaDtoWithDefaults() *AlertSchemaDto`

NewAlertSchemaDtoWithDefaults instantiates a new AlertSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertSchemaDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertSchemaDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertSchemaDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AlertSchemaDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertSchemaDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertSchemaDto) SetName(v string)`

SetName sets Name field to given value.


### GetAlertType

`func (o *AlertSchemaDto) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertSchemaDto) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertSchemaDto) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.


### GetMetrics

`func (o *AlertSchemaDto) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AlertSchemaDto) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AlertSchemaDto) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.


### GetMetricGroupBys

`func (o *AlertSchemaDto) GetMetricGroupBys() map[string]interface{}`

GetMetricGroupBys returns the MetricGroupBys field if non-nil, zero value otherwise.

### GetMetricGroupBysOk

`func (o *AlertSchemaDto) GetMetricGroupBysOk() (*map[string]interface{}, bool)`

GetMetricGroupBysOk returns a tuple with the MetricGroupBys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroupBys

`func (o *AlertSchemaDto) SetMetricGroupBys(v map[string]interface{})`

SetMetricGroupBys sets MetricGroupBys field to given value.


### GetFormula

`func (o *AlertSchemaDto) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *AlertSchemaDto) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *AlertSchemaDto) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *AlertSchemaDto) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetMessage

`func (o *AlertSchemaDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertSchemaDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertSchemaDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInstruction

`func (o *AlertSchemaDto) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *AlertSchemaDto) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *AlertSchemaDto) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.


### GetCreatorID

`func (o *AlertSchemaDto) GetCreatorID() string`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *AlertSchemaDto) GetCreatorIDOk() (*string, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *AlertSchemaDto) SetCreatorID(v string)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *AlertSchemaDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCompanyID

`func (o *AlertSchemaDto) GetCompanyID() string`

GetCompanyID returns the CompanyID field if non-nil, zero value otherwise.

### GetCompanyIDOk

`func (o *AlertSchemaDto) GetCompanyIDOk() (*string, bool)`

GetCompanyIDOk returns a tuple with the CompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyID

`func (o *AlertSchemaDto) SetCompanyID(v string)`

SetCompanyID sets CompanyID field to given value.


### GetPriority

`func (o *AlertSchemaDto) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AlertSchemaDto) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AlertSchemaDto) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetAlertThreshold

`func (o *AlertSchemaDto) GetAlertThreshold() float32`

GetAlertThreshold returns the AlertThreshold field if non-nil, zero value otherwise.

### GetAlertThresholdOk

`func (o *AlertSchemaDto) GetAlertThresholdOk() (*float32, bool)`

GetAlertThresholdOk returns a tuple with the AlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThreshold

`func (o *AlertSchemaDto) SetAlertThreshold(v float32)`

SetAlertThreshold sets AlertThreshold field to given value.


### GetWarningThreshold

`func (o *AlertSchemaDto) GetWarningThreshold() float32`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *AlertSchemaDto) GetWarningThresholdOk() (*float32, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *AlertSchemaDto) SetWarningThreshold(v float32)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *AlertSchemaDto) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetWindowMs

`func (o *AlertSchemaDto) GetWindowMs() float32`

GetWindowMs returns the WindowMs field if non-nil, zero value otherwise.

### GetWindowMsOk

`func (o *AlertSchemaDto) GetWindowMsOk() (*float32, bool)`

GetWindowMsOk returns a tuple with the WindowMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMs

`func (o *AlertSchemaDto) SetWindowMs(v float32)`

SetWindowMs sets WindowMs field to given value.


### GetCondition

`func (o *AlertSchemaDto) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AlertSchemaDto) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AlertSchemaDto) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetRenotificationCondition

`func (o *AlertSchemaDto) GetRenotificationCondition() string`

GetRenotificationCondition returns the RenotificationCondition field if non-nil, zero value otherwise.

### GetRenotificationConditionOk

`func (o *AlertSchemaDto) GetRenotificationConditionOk() (*string, bool)`

GetRenotificationConditionOk returns a tuple with the RenotificationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotificationCondition

`func (o *AlertSchemaDto) SetRenotificationCondition(v string)`

SetRenotificationCondition sets RenotificationCondition field to given value.

### HasRenotificationCondition

`func (o *AlertSchemaDto) HasRenotificationCondition() bool`

HasRenotificationCondition returns a boolean if a field has been set.

### GetRenotificationWindowMs

`func (o *AlertSchemaDto) GetRenotificationWindowMs() float32`

GetRenotificationWindowMs returns the RenotificationWindowMs field if non-nil, zero value otherwise.

### GetRenotificationWindowMsOk

`func (o *AlertSchemaDto) GetRenotificationWindowMsOk() (*float32, bool)`

GetRenotificationWindowMsOk returns a tuple with the RenotificationWindowMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotificationWindowMs

`func (o *AlertSchemaDto) SetRenotificationWindowMs(v float32)`

SetRenotificationWindowMs sets RenotificationWindowMs field to given value.

### HasRenotificationWindowMs

`func (o *AlertSchemaDto) HasRenotificationWindowMs() bool`

HasRenotificationWindowMs returns a boolean if a field has been set.

### GetRenotificationMessage

`func (o *AlertSchemaDto) GetRenotificationMessage() string`

GetRenotificationMessage returns the RenotificationMessage field if non-nil, zero value otherwise.

### GetRenotificationMessageOk

`func (o *AlertSchemaDto) GetRenotificationMessageOk() (*string, bool)`

GetRenotificationMessageOk returns a tuple with the RenotificationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotificationMessage

`func (o *AlertSchemaDto) SetRenotificationMessage(v string)`

SetRenotificationMessage sets RenotificationMessage field to given value.

### HasRenotificationMessage

`func (o *AlertSchemaDto) HasRenotificationMessage() bool`

HasRenotificationMessage returns a boolean if a field has been set.

### GetTeam

`func (o *AlertSchemaDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *AlertSchemaDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *AlertSchemaDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *AlertSchemaDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTags

`func (o *AlertSchemaDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AlertSchemaDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AlertSchemaDto) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


