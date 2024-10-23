# ExternalExperimentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewSettings** | Pointer to [**ExternalGateDtoReviewSettings**](ExternalGateDtoReviewSettings.md) |  | [optional] 
**ActiveReview** | Pointer to [**ExternalGateDtoActiveReview**](ExternalGateDtoActiveReview.md) |  | [optional] 
**Id** | **string** | ID | 
**Name** | Pointer to **string** | Optional name for the configuration. | [optional] 
**IdType** | **string** |  | 
**Description** | **string** | A helpful summary of what this experiment does | 
**LastModifierID** | [**nil**](nil.md) | ID of the last modifier. | 
**LastModifiedTime** | [**nil**](nil.md) | Time of the last modification. | 
**LastModifierEmail** | [**nil**](nil.md) | Email of the last modifier. | 
**LastModifierName** | [**nil**](nil.md) | Name of the last modifier. | 
**CreatorID** | Pointer to [**nil**](nil.md) | The Statsig ID of the creator of this experiment | [optional] 
**CreatedTime** | **float32** | Timestamp when the entity was created. | 
**CreatorName** | [**nil**](nil.md) | Name of the creator. | 
**CreatorEmail** | Pointer to [**nil**](nil.md) | The email of the creator of this experiment | [optional] 
**Tags** | **[]string** |  | 
**TargetApps** | Pointer to [**ExternalExperimentDtoTargetApps**](ExternalExperimentDtoTargetApps.md) |  | [optional] 
**HoldoutIDs** | Pointer to **[]string** | Holdouts applied to this configuration. | [optional] 
**Team** | Pointer to [**nil**](nil.md) | Enterprise only | [optional] 
**Version** | Pointer to **float32** | Version number | [optional] 
**SecondaryIDType** | Pointer to [**nil**](nil.md) | The secondary ID type for the experiment used in WHN for ID resolution | [optional] 
**Hypothesis** | **string** | A statement that will be tested by this experiment | 
**Links** | Pointer to [**[]ExternalExperimentDtoLinksInner**](ExternalExperimentDtoLinksInner.md) | Links to relevant documentation or resources | [optional] 
**Groups** | [**[]ExternalExperimentDtoGroupsInner**](ExternalExperimentDtoGroupsInner.md) | The test groups for your experiment | 
**ControlGroupID** | Pointer to **string** | Optional control group ID | [optional] 
**Allocation** | **float32** | Percent of layer allocated to this experiment | 
**PrimaryMetricTags** | Pointer to **[]string** | Primary metric tags for the experiment | [optional] 
**SecondaryMetricTags** | Pointer to **[]string** | Secondary metric tags for the experiment | [optional] 
**PrimaryMetrics** | [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) |  | 
**SecondaryMetrics** | [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) |  | 
**Duration** | Pointer to **int32** | How long the experiment is expected to last in days | [optional] 
**TargetExposures** | Pointer to **int32** | Target exposures for the experiment | [optional] 
**TargetingGateID** | [**nil**](nil.md) | Restrict your experiment to users passing the selected feature gate | 
**BonferroniCorrection** | **bool** | Is Bonferroni correction applied? | 
**DefaultConfidenceInterval** | **string** | Default error margin used for results | 
**Status** | **string** | The current status of the experiment | 
**LaunchedGroupID** | Pointer to [**nil**](nil.md) | ID of the launched group, null otherwise | [optional] 
**AssignmentSourceName** | Pointer to **string** |  | [optional] 
**AssignmentSourceExperimentName** | Pointer to **string** | Name of the source experiment for assignment | [optional] 
**IsAnalysisOnly** | Pointer to **bool** |  | [optional] 
**AllocationDuration** | Pointer to **int32** | Allocation duration in days | [optional] 
**CohortedAnalysisDuration** | Pointer to **int32** | Cohorted analysis duration in days | [optional] 
**FixedAnalysisDuration** | Pointer to **int32** | Fixed analysis duration in days | [optional] 
**ScheduledReloadHour** | Pointer to **int32** | Warehouse Native only - UTC hour at which to run scheduled pulse loads | [optional] 
**ScheduledReloadType** | Pointer to **string** | Warehouse Native only - reload type for scheduled reloads | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**ExternalExperimentName** | Pointer to **string** |  | [optional] 
**LayerID** | [**nil**](nil.md) |  | 
**StartTime** | [**nil**](nil.md) |  | 
**EndTime** | [**nil**](nil.md) |  | 
**DecisionReason** | [**nil**](nil.md) |  | 
**DecisionTime** | [**nil**](nil.md) |  | 
**HealthChecks** | Pointer to [**[]ExternalExperimentDtoHealthChecksInner**](ExternalExperimentDtoHealthChecksInner.md) |  | [optional] 
**Owner** | Pointer to [**nil**](nil.md) | Schema for owner data including ID, type, name. Note that if Entity is created by CONSOLE API, owner will be undefined. | [optional] 

## Methods

### NewExternalExperimentDto

`func NewExternalExperimentDto(id string, idType string, description string, lastModifierID nil, lastModifiedTime nil, lastModifierEmail nil, lastModifierName nil, createdTime float32, creatorName nil, tags []string, hypothesis string, groups []ExternalExperimentDtoGroupsInner, allocation float32, primaryMetrics []ExternalGateDtoMonitoringMetricsInner, secondaryMetrics []ExternalGateDtoMonitoringMetricsInner, targetingGateID nil, bonferroniCorrection bool, defaultConfidenceInterval string, status string, layerID nil, startTime nil, endTime nil, decisionReason nil, decisionTime nil, ) *ExternalExperimentDto`

NewExternalExperimentDto instantiates a new ExternalExperimentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalExperimentDtoWithDefaults

`func NewExternalExperimentDtoWithDefaults() *ExternalExperimentDto`

NewExternalExperimentDtoWithDefaults instantiates a new ExternalExperimentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewSettings

`func (o *ExternalExperimentDto) GetReviewSettings() ExternalGateDtoReviewSettings`

GetReviewSettings returns the ReviewSettings field if non-nil, zero value otherwise.

### GetReviewSettingsOk

`func (o *ExternalExperimentDto) GetReviewSettingsOk() (*ExternalGateDtoReviewSettings, bool)`

GetReviewSettingsOk returns a tuple with the ReviewSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewSettings

`func (o *ExternalExperimentDto) SetReviewSettings(v ExternalGateDtoReviewSettings)`

SetReviewSettings sets ReviewSettings field to given value.

### HasReviewSettings

`func (o *ExternalExperimentDto) HasReviewSettings() bool`

HasReviewSettings returns a boolean if a field has been set.

### GetActiveReview

`func (o *ExternalExperimentDto) GetActiveReview() ExternalGateDtoActiveReview`

GetActiveReview returns the ActiveReview field if non-nil, zero value otherwise.

### GetActiveReviewOk

`func (o *ExternalExperimentDto) GetActiveReviewOk() (*ExternalGateDtoActiveReview, bool)`

GetActiveReviewOk returns a tuple with the ActiveReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveReview

`func (o *ExternalExperimentDto) SetActiveReview(v ExternalGateDtoActiveReview)`

SetActiveReview sets ActiveReview field to given value.

### HasActiveReview

`func (o *ExternalExperimentDto) HasActiveReview() bool`

HasActiveReview returns a boolean if a field has been set.

### GetId

`func (o *ExternalExperimentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalExperimentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalExperimentDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ExternalExperimentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalExperimentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalExperimentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalExperimentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdType

`func (o *ExternalExperimentDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *ExternalExperimentDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *ExternalExperimentDto) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetDescription

`func (o *ExternalExperimentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalExperimentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalExperimentDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLastModifierID

`func (o *ExternalExperimentDto) GetLastModifierID() nil`

GetLastModifierID returns the LastModifierID field if non-nil, zero value otherwise.

### GetLastModifierIDOk

`func (o *ExternalExperimentDto) GetLastModifierIDOk() (*nil, bool)`

GetLastModifierIDOk returns a tuple with the LastModifierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierID

`func (o *ExternalExperimentDto) SetLastModifierID(v nil)`

SetLastModifierID sets LastModifierID field to given value.


### GetLastModifiedTime

`func (o *ExternalExperimentDto) GetLastModifiedTime() nil`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *ExternalExperimentDto) GetLastModifiedTimeOk() (*nil, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *ExternalExperimentDto) SetLastModifiedTime(v nil)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastModifierEmail

`func (o *ExternalExperimentDto) GetLastModifierEmail() nil`

GetLastModifierEmail returns the LastModifierEmail field if non-nil, zero value otherwise.

### GetLastModifierEmailOk

`func (o *ExternalExperimentDto) GetLastModifierEmailOk() (*nil, bool)`

GetLastModifierEmailOk returns a tuple with the LastModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierEmail

`func (o *ExternalExperimentDto) SetLastModifierEmail(v nil)`

SetLastModifierEmail sets LastModifierEmail field to given value.


### GetLastModifierName

`func (o *ExternalExperimentDto) GetLastModifierName() nil`

GetLastModifierName returns the LastModifierName field if non-nil, zero value otherwise.

### GetLastModifierNameOk

`func (o *ExternalExperimentDto) GetLastModifierNameOk() (*nil, bool)`

GetLastModifierNameOk returns a tuple with the LastModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierName

`func (o *ExternalExperimentDto) SetLastModifierName(v nil)`

SetLastModifierName sets LastModifierName field to given value.


### GetCreatorID

`func (o *ExternalExperimentDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *ExternalExperimentDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *ExternalExperimentDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *ExternalExperimentDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ExternalExperimentDto) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ExternalExperimentDto) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ExternalExperimentDto) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreatorName

`func (o *ExternalExperimentDto) GetCreatorName() nil`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ExternalExperimentDto) GetCreatorNameOk() (*nil, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ExternalExperimentDto) SetCreatorName(v nil)`

SetCreatorName sets CreatorName field to given value.


### GetCreatorEmail

`func (o *ExternalExperimentDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *ExternalExperimentDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *ExternalExperimentDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *ExternalExperimentDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetTags

`func (o *ExternalExperimentDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternalExperimentDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternalExperimentDto) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTargetApps

`func (o *ExternalExperimentDto) GetTargetApps() ExternalExperimentDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *ExternalExperimentDto) GetTargetAppsOk() (*ExternalExperimentDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *ExternalExperimentDto) SetTargetApps(v ExternalExperimentDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *ExternalExperimentDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetHoldoutIDs

`func (o *ExternalExperimentDto) GetHoldoutIDs() []string`

GetHoldoutIDs returns the HoldoutIDs field if non-nil, zero value otherwise.

### GetHoldoutIDsOk

`func (o *ExternalExperimentDto) GetHoldoutIDsOk() (*[]string, bool)`

GetHoldoutIDsOk returns a tuple with the HoldoutIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutIDs

`func (o *ExternalExperimentDto) SetHoldoutIDs(v []string)`

SetHoldoutIDs sets HoldoutIDs field to given value.

### HasHoldoutIDs

`func (o *ExternalExperimentDto) HasHoldoutIDs() bool`

HasHoldoutIDs returns a boolean if a field has been set.

### GetTeam

`func (o *ExternalExperimentDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExternalExperimentDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExternalExperimentDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExternalExperimentDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetVersion

`func (o *ExternalExperimentDto) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExternalExperimentDto) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExternalExperimentDto) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExternalExperimentDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSecondaryIDType

`func (o *ExternalExperimentDto) GetSecondaryIDType() nil`

GetSecondaryIDType returns the SecondaryIDType field if non-nil, zero value otherwise.

### GetSecondaryIDTypeOk

`func (o *ExternalExperimentDto) GetSecondaryIDTypeOk() (*nil, bool)`

GetSecondaryIDTypeOk returns a tuple with the SecondaryIDType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIDType

`func (o *ExternalExperimentDto) SetSecondaryIDType(v nil)`

SetSecondaryIDType sets SecondaryIDType field to given value.

### HasSecondaryIDType

`func (o *ExternalExperimentDto) HasSecondaryIDType() bool`

HasSecondaryIDType returns a boolean if a field has been set.

### GetHypothesis

`func (o *ExternalExperimentDto) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *ExternalExperimentDto) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *ExternalExperimentDto) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.


### GetLinks

`func (o *ExternalExperimentDto) GetLinks() []ExternalExperimentDtoLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalExperimentDto) GetLinksOk() (*[]ExternalExperimentDtoLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalExperimentDto) SetLinks(v []ExternalExperimentDtoLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalExperimentDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetGroups

`func (o *ExternalExperimentDto) GetGroups() []ExternalExperimentDtoGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ExternalExperimentDto) GetGroupsOk() (*[]ExternalExperimentDtoGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ExternalExperimentDto) SetGroups(v []ExternalExperimentDtoGroupsInner)`

SetGroups sets Groups field to given value.


### GetControlGroupID

`func (o *ExternalExperimentDto) GetControlGroupID() string`

GetControlGroupID returns the ControlGroupID field if non-nil, zero value otherwise.

### GetControlGroupIDOk

`func (o *ExternalExperimentDto) GetControlGroupIDOk() (*string, bool)`

GetControlGroupIDOk returns a tuple with the ControlGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlGroupID

`func (o *ExternalExperimentDto) SetControlGroupID(v string)`

SetControlGroupID sets ControlGroupID field to given value.

### HasControlGroupID

`func (o *ExternalExperimentDto) HasControlGroupID() bool`

HasControlGroupID returns a boolean if a field has been set.

### GetAllocation

`func (o *ExternalExperimentDto) GetAllocation() float32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *ExternalExperimentDto) GetAllocationOk() (*float32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *ExternalExperimentDto) SetAllocation(v float32)`

SetAllocation sets Allocation field to given value.


### GetPrimaryMetricTags

`func (o *ExternalExperimentDto) GetPrimaryMetricTags() []string`

GetPrimaryMetricTags returns the PrimaryMetricTags field if non-nil, zero value otherwise.

### GetPrimaryMetricTagsOk

`func (o *ExternalExperimentDto) GetPrimaryMetricTagsOk() (*[]string, bool)`

GetPrimaryMetricTagsOk returns a tuple with the PrimaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetricTags

`func (o *ExternalExperimentDto) SetPrimaryMetricTags(v []string)`

SetPrimaryMetricTags sets PrimaryMetricTags field to given value.

### HasPrimaryMetricTags

`func (o *ExternalExperimentDto) HasPrimaryMetricTags() bool`

HasPrimaryMetricTags returns a boolean if a field has been set.

### GetSecondaryMetricTags

`func (o *ExternalExperimentDto) GetSecondaryMetricTags() []string`

GetSecondaryMetricTags returns the SecondaryMetricTags field if non-nil, zero value otherwise.

### GetSecondaryMetricTagsOk

`func (o *ExternalExperimentDto) GetSecondaryMetricTagsOk() (*[]string, bool)`

GetSecondaryMetricTagsOk returns a tuple with the SecondaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetricTags

`func (o *ExternalExperimentDto) SetSecondaryMetricTags(v []string)`

SetSecondaryMetricTags sets SecondaryMetricTags field to given value.

### HasSecondaryMetricTags

`func (o *ExternalExperimentDto) HasSecondaryMetricTags() bool`

HasSecondaryMetricTags returns a boolean if a field has been set.

### GetPrimaryMetrics

`func (o *ExternalExperimentDto) GetPrimaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetPrimaryMetrics returns the PrimaryMetrics field if non-nil, zero value otherwise.

### GetPrimaryMetricsOk

`func (o *ExternalExperimentDto) GetPrimaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetPrimaryMetricsOk returns a tuple with the PrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetrics

`func (o *ExternalExperimentDto) SetPrimaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetPrimaryMetrics sets PrimaryMetrics field to given value.


### GetSecondaryMetrics

`func (o *ExternalExperimentDto) GetSecondaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *ExternalExperimentDto) GetSecondaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *ExternalExperimentDto) SetSecondaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.


### GetDuration

`func (o *ExternalExperimentDto) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ExternalExperimentDto) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ExternalExperimentDto) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ExternalExperimentDto) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTargetExposures

`func (o *ExternalExperimentDto) GetTargetExposures() int32`

GetTargetExposures returns the TargetExposures field if non-nil, zero value otherwise.

### GetTargetExposuresOk

`func (o *ExternalExperimentDto) GetTargetExposuresOk() (*int32, bool)`

GetTargetExposuresOk returns a tuple with the TargetExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetExposures

`func (o *ExternalExperimentDto) SetTargetExposures(v int32)`

SetTargetExposures sets TargetExposures field to given value.

### HasTargetExposures

`func (o *ExternalExperimentDto) HasTargetExposures() bool`

HasTargetExposures returns a boolean if a field has been set.

### GetTargetingGateID

`func (o *ExternalExperimentDto) GetTargetingGateID() nil`

GetTargetingGateID returns the TargetingGateID field if non-nil, zero value otherwise.

### GetTargetingGateIDOk

`func (o *ExternalExperimentDto) GetTargetingGateIDOk() (*nil, bool)`

GetTargetingGateIDOk returns a tuple with the TargetingGateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingGateID

`func (o *ExternalExperimentDto) SetTargetingGateID(v nil)`

SetTargetingGateID sets TargetingGateID field to given value.


### GetBonferroniCorrection

`func (o *ExternalExperimentDto) GetBonferroniCorrection() bool`

GetBonferroniCorrection returns the BonferroniCorrection field if non-nil, zero value otherwise.

### GetBonferroniCorrectionOk

`func (o *ExternalExperimentDto) GetBonferroniCorrectionOk() (*bool, bool)`

GetBonferroniCorrectionOk returns a tuple with the BonferroniCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonferroniCorrection

`func (o *ExternalExperimentDto) SetBonferroniCorrection(v bool)`

SetBonferroniCorrection sets BonferroniCorrection field to given value.


### GetDefaultConfidenceInterval

`func (o *ExternalExperimentDto) GetDefaultConfidenceInterval() string`

GetDefaultConfidenceInterval returns the DefaultConfidenceInterval field if non-nil, zero value otherwise.

### GetDefaultConfidenceIntervalOk

`func (o *ExternalExperimentDto) GetDefaultConfidenceIntervalOk() (*string, bool)`

GetDefaultConfidenceIntervalOk returns a tuple with the DefaultConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfidenceInterval

`func (o *ExternalExperimentDto) SetDefaultConfidenceInterval(v string)`

SetDefaultConfidenceInterval sets DefaultConfidenceInterval field to given value.


### GetStatus

`func (o *ExternalExperimentDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalExperimentDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalExperimentDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLaunchedGroupID

`func (o *ExternalExperimentDto) GetLaunchedGroupID() nil`

GetLaunchedGroupID returns the LaunchedGroupID field if non-nil, zero value otherwise.

### GetLaunchedGroupIDOk

`func (o *ExternalExperimentDto) GetLaunchedGroupIDOk() (*nil, bool)`

GetLaunchedGroupIDOk returns a tuple with the LaunchedGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedGroupID

`func (o *ExternalExperimentDto) SetLaunchedGroupID(v nil)`

SetLaunchedGroupID sets LaunchedGroupID field to given value.

### HasLaunchedGroupID

`func (o *ExternalExperimentDto) HasLaunchedGroupID() bool`

HasLaunchedGroupID returns a boolean if a field has been set.

### GetAssignmentSourceName

`func (o *ExternalExperimentDto) GetAssignmentSourceName() string`

GetAssignmentSourceName returns the AssignmentSourceName field if non-nil, zero value otherwise.

### GetAssignmentSourceNameOk

`func (o *ExternalExperimentDto) GetAssignmentSourceNameOk() (*string, bool)`

GetAssignmentSourceNameOk returns a tuple with the AssignmentSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceName

`func (o *ExternalExperimentDto) SetAssignmentSourceName(v string)`

SetAssignmentSourceName sets AssignmentSourceName field to given value.

### HasAssignmentSourceName

`func (o *ExternalExperimentDto) HasAssignmentSourceName() bool`

HasAssignmentSourceName returns a boolean if a field has been set.

### GetAssignmentSourceExperimentName

`func (o *ExternalExperimentDto) GetAssignmentSourceExperimentName() string`

GetAssignmentSourceExperimentName returns the AssignmentSourceExperimentName field if non-nil, zero value otherwise.

### GetAssignmentSourceExperimentNameOk

`func (o *ExternalExperimentDto) GetAssignmentSourceExperimentNameOk() (*string, bool)`

GetAssignmentSourceExperimentNameOk returns a tuple with the AssignmentSourceExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceExperimentName

`func (o *ExternalExperimentDto) SetAssignmentSourceExperimentName(v string)`

SetAssignmentSourceExperimentName sets AssignmentSourceExperimentName field to given value.

### HasAssignmentSourceExperimentName

`func (o *ExternalExperimentDto) HasAssignmentSourceExperimentName() bool`

HasAssignmentSourceExperimentName returns a boolean if a field has been set.

### GetIsAnalysisOnly

`func (o *ExternalExperimentDto) GetIsAnalysisOnly() bool`

GetIsAnalysisOnly returns the IsAnalysisOnly field if non-nil, zero value otherwise.

### GetIsAnalysisOnlyOk

`func (o *ExternalExperimentDto) GetIsAnalysisOnlyOk() (*bool, bool)`

GetIsAnalysisOnlyOk returns a tuple with the IsAnalysisOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalysisOnly

`func (o *ExternalExperimentDto) SetIsAnalysisOnly(v bool)`

SetIsAnalysisOnly sets IsAnalysisOnly field to given value.

### HasIsAnalysisOnly

`func (o *ExternalExperimentDto) HasIsAnalysisOnly() bool`

HasIsAnalysisOnly returns a boolean if a field has been set.

### GetAllocationDuration

`func (o *ExternalExperimentDto) GetAllocationDuration() int32`

GetAllocationDuration returns the AllocationDuration field if non-nil, zero value otherwise.

### GetAllocationDurationOk

`func (o *ExternalExperimentDto) GetAllocationDurationOk() (*int32, bool)`

GetAllocationDurationOk returns a tuple with the AllocationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDuration

`func (o *ExternalExperimentDto) SetAllocationDuration(v int32)`

SetAllocationDuration sets AllocationDuration field to given value.

### HasAllocationDuration

`func (o *ExternalExperimentDto) HasAllocationDuration() bool`

HasAllocationDuration returns a boolean if a field has been set.

### GetCohortedAnalysisDuration

`func (o *ExternalExperimentDto) GetCohortedAnalysisDuration() int32`

GetCohortedAnalysisDuration returns the CohortedAnalysisDuration field if non-nil, zero value otherwise.

### GetCohortedAnalysisDurationOk

`func (o *ExternalExperimentDto) GetCohortedAnalysisDurationOk() (*int32, bool)`

GetCohortedAnalysisDurationOk returns a tuple with the CohortedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortedAnalysisDuration

`func (o *ExternalExperimentDto) SetCohortedAnalysisDuration(v int32)`

SetCohortedAnalysisDuration sets CohortedAnalysisDuration field to given value.

### HasCohortedAnalysisDuration

`func (o *ExternalExperimentDto) HasCohortedAnalysisDuration() bool`

HasCohortedAnalysisDuration returns a boolean if a field has been set.

### GetFixedAnalysisDuration

`func (o *ExternalExperimentDto) GetFixedAnalysisDuration() int32`

GetFixedAnalysisDuration returns the FixedAnalysisDuration field if non-nil, zero value otherwise.

### GetFixedAnalysisDurationOk

`func (o *ExternalExperimentDto) GetFixedAnalysisDurationOk() (*int32, bool)`

GetFixedAnalysisDurationOk returns a tuple with the FixedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAnalysisDuration

`func (o *ExternalExperimentDto) SetFixedAnalysisDuration(v int32)`

SetFixedAnalysisDuration sets FixedAnalysisDuration field to given value.

### HasFixedAnalysisDuration

`func (o *ExternalExperimentDto) HasFixedAnalysisDuration() bool`

HasFixedAnalysisDuration returns a boolean if a field has been set.

### GetScheduledReloadHour

`func (o *ExternalExperimentDto) GetScheduledReloadHour() int32`

GetScheduledReloadHour returns the ScheduledReloadHour field if non-nil, zero value otherwise.

### GetScheduledReloadHourOk

`func (o *ExternalExperimentDto) GetScheduledReloadHourOk() (*int32, bool)`

GetScheduledReloadHourOk returns a tuple with the ScheduledReloadHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadHour

`func (o *ExternalExperimentDto) SetScheduledReloadHour(v int32)`

SetScheduledReloadHour sets ScheduledReloadHour field to given value.

### HasScheduledReloadHour

`func (o *ExternalExperimentDto) HasScheduledReloadHour() bool`

HasScheduledReloadHour returns a boolean if a field has been set.

### GetScheduledReloadType

`func (o *ExternalExperimentDto) GetScheduledReloadType() string`

GetScheduledReloadType returns the ScheduledReloadType field if non-nil, zero value otherwise.

### GetScheduledReloadTypeOk

`func (o *ExternalExperimentDto) GetScheduledReloadTypeOk() (*string, bool)`

GetScheduledReloadTypeOk returns a tuple with the ScheduledReloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadType

`func (o *ExternalExperimentDto) SetScheduledReloadType(v string)`

SetScheduledReloadType sets ScheduledReloadType field to given value.

### HasScheduledReloadType

`func (o *ExternalExperimentDto) HasScheduledReloadType() bool`

HasScheduledReloadType returns a boolean if a field has been set.

### GetSubtype

`func (o *ExternalExperimentDto) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ExternalExperimentDto) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ExternalExperimentDto) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ExternalExperimentDto) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetExternalExperimentName

`func (o *ExternalExperimentDto) GetExternalExperimentName() string`

GetExternalExperimentName returns the ExternalExperimentName field if non-nil, zero value otherwise.

### GetExternalExperimentNameOk

`func (o *ExternalExperimentDto) GetExternalExperimentNameOk() (*string, bool)`

GetExternalExperimentNameOk returns a tuple with the ExternalExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalExperimentName

`func (o *ExternalExperimentDto) SetExternalExperimentName(v string)`

SetExternalExperimentName sets ExternalExperimentName field to given value.

### HasExternalExperimentName

`func (o *ExternalExperimentDto) HasExternalExperimentName() bool`

HasExternalExperimentName returns a boolean if a field has been set.

### GetLayerID

`func (o *ExternalExperimentDto) GetLayerID() nil`

GetLayerID returns the LayerID field if non-nil, zero value otherwise.

### GetLayerIDOk

`func (o *ExternalExperimentDto) GetLayerIDOk() (*nil, bool)`

GetLayerIDOk returns a tuple with the LayerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerID

`func (o *ExternalExperimentDto) SetLayerID(v nil)`

SetLayerID sets LayerID field to given value.


### GetStartTime

`func (o *ExternalExperimentDto) GetStartTime() nil`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExternalExperimentDto) GetStartTimeOk() (*nil, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExternalExperimentDto) SetStartTime(v nil)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ExternalExperimentDto) GetEndTime() nil`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ExternalExperimentDto) GetEndTimeOk() (*nil, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ExternalExperimentDto) SetEndTime(v nil)`

SetEndTime sets EndTime field to given value.


### GetDecisionReason

`func (o *ExternalExperimentDto) GetDecisionReason() nil`

GetDecisionReason returns the DecisionReason field if non-nil, zero value otherwise.

### GetDecisionReasonOk

`func (o *ExternalExperimentDto) GetDecisionReasonOk() (*nil, bool)`

GetDecisionReasonOk returns a tuple with the DecisionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionReason

`func (o *ExternalExperimentDto) SetDecisionReason(v nil)`

SetDecisionReason sets DecisionReason field to given value.


### GetDecisionTime

`func (o *ExternalExperimentDto) GetDecisionTime() nil`

GetDecisionTime returns the DecisionTime field if non-nil, zero value otherwise.

### GetDecisionTimeOk

`func (o *ExternalExperimentDto) GetDecisionTimeOk() (*nil, bool)`

GetDecisionTimeOk returns a tuple with the DecisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionTime

`func (o *ExternalExperimentDto) SetDecisionTime(v nil)`

SetDecisionTime sets DecisionTime field to given value.


### GetHealthChecks

`func (o *ExternalExperimentDto) GetHealthChecks() []ExternalExperimentDtoHealthChecksInner`

GetHealthChecks returns the HealthChecks field if non-nil, zero value otherwise.

### GetHealthChecksOk

`func (o *ExternalExperimentDto) GetHealthChecksOk() (*[]ExternalExperimentDtoHealthChecksInner, bool)`

GetHealthChecksOk returns a tuple with the HealthChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthChecks

`func (o *ExternalExperimentDto) SetHealthChecks(v []ExternalExperimentDtoHealthChecksInner)`

SetHealthChecks sets HealthChecks field to given value.

### HasHealthChecks

`func (o *ExternalExperimentDto) HasHealthChecks() bool`

HasHealthChecks returns a boolean if a field has been set.

### GetOwner

`func (o *ExternalExperimentDto) GetOwner() nil`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExternalExperimentDto) GetOwnerOk() (*nil, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExternalExperimentDto) SetOwner(v nil)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExternalExperimentDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


