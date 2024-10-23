# ExperimentFullUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A helpful summary of what this experiment does | 
**IdType** | **string** | The type of ID which the experiment is based on | 
**SecondaryIDType** | Pointer to [**nil**](nil.md) | The secondary ID type for the experiment used in WHN for ID resolution | [optional] 
**Hypothesis** | **string** | A statement that will be tested by this experiment | 
**Links** | Pointer to [**[]ExternalExperimentDtoLinksInner**](ExternalExperimentDtoLinksInner.md) | Links to relevant documentation or resources | [optional] 
**Groups** | [**[]ExternalExperimentDtoGroupsInner**](ExternalExperimentDtoGroupsInner.md) | The test groups for your experiment | 
**ControlGroupID** | Pointer to **string** | Optional control group ID | [optional] 
**Allocation** | **float32** | Percent of layer allocated to this experiment | 
**PrimaryMetricTags** | Pointer to **[]string** | Primary metric tags for the experiment | [optional] 
**SecondaryMetricTags** | Pointer to **[]string** | Secondary metric tags for the experiment | [optional] 
**PrimaryMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) | Main metrics needed to evaluate your hypothesis | [optional] 
**SecondaryMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) | Additional metrics to monitor that might impact the analysis or final decision of the experiment | [optional] 
**TargetApps** | Pointer to [**ExternalExperimentDtoTargetApps**](ExternalExperimentDtoTargetApps.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the experiment | [optional] 
**Duration** | Pointer to **int32** | How long the experiment is expected to last in days | [optional] 
**TargetExposures** | Pointer to **int32** | Target exposures for the experiment | [optional] 
**TargetingGateID** | [**nil**](nil.md) | Restrict your experiment to users passing the selected feature gate | 
**BonferroniCorrection** | **bool** | Is Bonferroni correction applied? | 
**DefaultConfidenceInterval** | **string** | Default error margin used for results | 
**Status** | **string** | The current status of the experiment | 
**LaunchedGroupID** | Pointer to [**nil**](nil.md) | ID of the launched group, null otherwise | [optional] 
**AssignmentSourceName** | Pointer to **string** | Source name of the assignment | [optional] 
**AssignmentSourceExperimentName** | Pointer to **string** | Name of the source experiment for assignment | [optional] 
**CreatorID** | Pointer to [**nil**](nil.md) | The Statsig ID of the creator of this experiment | [optional] 
**CreatorEmail** | Pointer to [**nil**](nil.md) | The email of the creator of this experiment | [optional] 
**IsAnalysisOnly** | Pointer to [**nil**](nil.md) | For Warehouse Native | [optional] 
**Team** | Pointer to [**nil**](nil.md) | Enterprise only | [optional] 
**AllocationDuration** | Pointer to **int32** | Allocation duration in days | [optional] 
**CohortedAnalysisDuration** | Pointer to **int32** | Cohorted analysis duration in days | [optional] 
**FixedAnalysisDuration** | Pointer to **int32** | Fixed analysis duration in days | [optional] 
**ScheduledReloadHour** | Pointer to **int32** | Warehouse Native only - UTC hour at which to run scheduled pulse loads | [optional] 
**ScheduledReloadType** | Pointer to **string** | Warehouse Native only - reload type for scheduled reloads | [optional] 

## Methods

### NewExperimentFullUpdateDto

`func NewExperimentFullUpdateDto(description string, idType string, hypothesis string, groups []ExternalExperimentDtoGroupsInner, allocation float32, targetingGateID nil, bonferroniCorrection bool, defaultConfidenceInterval string, status string, ) *ExperimentFullUpdateDto`

NewExperimentFullUpdateDto instantiates a new ExperimentFullUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentFullUpdateDtoWithDefaults

`func NewExperimentFullUpdateDtoWithDefaults() *ExperimentFullUpdateDto`

NewExperimentFullUpdateDtoWithDefaults instantiates a new ExperimentFullUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ExperimentFullUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperimentFullUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperimentFullUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIdType

`func (o *ExperimentFullUpdateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *ExperimentFullUpdateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *ExperimentFullUpdateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetSecondaryIDType

`func (o *ExperimentFullUpdateDto) GetSecondaryIDType() nil`

GetSecondaryIDType returns the SecondaryIDType field if non-nil, zero value otherwise.

### GetSecondaryIDTypeOk

`func (o *ExperimentFullUpdateDto) GetSecondaryIDTypeOk() (*nil, bool)`

GetSecondaryIDTypeOk returns a tuple with the SecondaryIDType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIDType

`func (o *ExperimentFullUpdateDto) SetSecondaryIDType(v nil)`

SetSecondaryIDType sets SecondaryIDType field to given value.

### HasSecondaryIDType

`func (o *ExperimentFullUpdateDto) HasSecondaryIDType() bool`

HasSecondaryIDType returns a boolean if a field has been set.

### GetHypothesis

`func (o *ExperimentFullUpdateDto) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *ExperimentFullUpdateDto) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *ExperimentFullUpdateDto) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.


### GetLinks

`func (o *ExperimentFullUpdateDto) GetLinks() []ExternalExperimentDtoLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentFullUpdateDto) GetLinksOk() (*[]ExternalExperimentDtoLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentFullUpdateDto) SetLinks(v []ExternalExperimentDtoLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentFullUpdateDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetGroups

`func (o *ExperimentFullUpdateDto) GetGroups() []ExternalExperimentDtoGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ExperimentFullUpdateDto) GetGroupsOk() (*[]ExternalExperimentDtoGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ExperimentFullUpdateDto) SetGroups(v []ExternalExperimentDtoGroupsInner)`

SetGroups sets Groups field to given value.


### GetControlGroupID

`func (o *ExperimentFullUpdateDto) GetControlGroupID() string`

GetControlGroupID returns the ControlGroupID field if non-nil, zero value otherwise.

### GetControlGroupIDOk

`func (o *ExperimentFullUpdateDto) GetControlGroupIDOk() (*string, bool)`

GetControlGroupIDOk returns a tuple with the ControlGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlGroupID

`func (o *ExperimentFullUpdateDto) SetControlGroupID(v string)`

SetControlGroupID sets ControlGroupID field to given value.

### HasControlGroupID

`func (o *ExperimentFullUpdateDto) HasControlGroupID() bool`

HasControlGroupID returns a boolean if a field has been set.

### GetAllocation

`func (o *ExperimentFullUpdateDto) GetAllocation() float32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *ExperimentFullUpdateDto) GetAllocationOk() (*float32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *ExperimentFullUpdateDto) SetAllocation(v float32)`

SetAllocation sets Allocation field to given value.


### GetPrimaryMetricTags

`func (o *ExperimentFullUpdateDto) GetPrimaryMetricTags() []string`

GetPrimaryMetricTags returns the PrimaryMetricTags field if non-nil, zero value otherwise.

### GetPrimaryMetricTagsOk

`func (o *ExperimentFullUpdateDto) GetPrimaryMetricTagsOk() (*[]string, bool)`

GetPrimaryMetricTagsOk returns a tuple with the PrimaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetricTags

`func (o *ExperimentFullUpdateDto) SetPrimaryMetricTags(v []string)`

SetPrimaryMetricTags sets PrimaryMetricTags field to given value.

### HasPrimaryMetricTags

`func (o *ExperimentFullUpdateDto) HasPrimaryMetricTags() bool`

HasPrimaryMetricTags returns a boolean if a field has been set.

### GetSecondaryMetricTags

`func (o *ExperimentFullUpdateDto) GetSecondaryMetricTags() []string`

GetSecondaryMetricTags returns the SecondaryMetricTags field if non-nil, zero value otherwise.

### GetSecondaryMetricTagsOk

`func (o *ExperimentFullUpdateDto) GetSecondaryMetricTagsOk() (*[]string, bool)`

GetSecondaryMetricTagsOk returns a tuple with the SecondaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetricTags

`func (o *ExperimentFullUpdateDto) SetSecondaryMetricTags(v []string)`

SetSecondaryMetricTags sets SecondaryMetricTags field to given value.

### HasSecondaryMetricTags

`func (o *ExperimentFullUpdateDto) HasSecondaryMetricTags() bool`

HasSecondaryMetricTags returns a boolean if a field has been set.

### GetPrimaryMetrics

`func (o *ExperimentFullUpdateDto) GetPrimaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetPrimaryMetrics returns the PrimaryMetrics field if non-nil, zero value otherwise.

### GetPrimaryMetricsOk

`func (o *ExperimentFullUpdateDto) GetPrimaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetPrimaryMetricsOk returns a tuple with the PrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetrics

`func (o *ExperimentFullUpdateDto) SetPrimaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetPrimaryMetrics sets PrimaryMetrics field to given value.

### HasPrimaryMetrics

`func (o *ExperimentFullUpdateDto) HasPrimaryMetrics() bool`

HasPrimaryMetrics returns a boolean if a field has been set.

### GetSecondaryMetrics

`func (o *ExperimentFullUpdateDto) GetSecondaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *ExperimentFullUpdateDto) GetSecondaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *ExperimentFullUpdateDto) SetSecondaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *ExperimentFullUpdateDto) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.

### GetTargetApps

`func (o *ExperimentFullUpdateDto) GetTargetApps() ExternalExperimentDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *ExperimentFullUpdateDto) GetTargetAppsOk() (*ExternalExperimentDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *ExperimentFullUpdateDto) SetTargetApps(v ExternalExperimentDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *ExperimentFullUpdateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetTags

`func (o *ExperimentFullUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExperimentFullUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExperimentFullUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExperimentFullUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDuration

`func (o *ExperimentFullUpdateDto) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ExperimentFullUpdateDto) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ExperimentFullUpdateDto) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ExperimentFullUpdateDto) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTargetExposures

`func (o *ExperimentFullUpdateDto) GetTargetExposures() int32`

GetTargetExposures returns the TargetExposures field if non-nil, zero value otherwise.

### GetTargetExposuresOk

`func (o *ExperimentFullUpdateDto) GetTargetExposuresOk() (*int32, bool)`

GetTargetExposuresOk returns a tuple with the TargetExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetExposures

`func (o *ExperimentFullUpdateDto) SetTargetExposures(v int32)`

SetTargetExposures sets TargetExposures field to given value.

### HasTargetExposures

`func (o *ExperimentFullUpdateDto) HasTargetExposures() bool`

HasTargetExposures returns a boolean if a field has been set.

### GetTargetingGateID

`func (o *ExperimentFullUpdateDto) GetTargetingGateID() nil`

GetTargetingGateID returns the TargetingGateID field if non-nil, zero value otherwise.

### GetTargetingGateIDOk

`func (o *ExperimentFullUpdateDto) GetTargetingGateIDOk() (*nil, bool)`

GetTargetingGateIDOk returns a tuple with the TargetingGateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingGateID

`func (o *ExperimentFullUpdateDto) SetTargetingGateID(v nil)`

SetTargetingGateID sets TargetingGateID field to given value.


### GetBonferroniCorrection

`func (o *ExperimentFullUpdateDto) GetBonferroniCorrection() bool`

GetBonferroniCorrection returns the BonferroniCorrection field if non-nil, zero value otherwise.

### GetBonferroniCorrectionOk

`func (o *ExperimentFullUpdateDto) GetBonferroniCorrectionOk() (*bool, bool)`

GetBonferroniCorrectionOk returns a tuple with the BonferroniCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonferroniCorrection

`func (o *ExperimentFullUpdateDto) SetBonferroniCorrection(v bool)`

SetBonferroniCorrection sets BonferroniCorrection field to given value.


### GetDefaultConfidenceInterval

`func (o *ExperimentFullUpdateDto) GetDefaultConfidenceInterval() string`

GetDefaultConfidenceInterval returns the DefaultConfidenceInterval field if non-nil, zero value otherwise.

### GetDefaultConfidenceIntervalOk

`func (o *ExperimentFullUpdateDto) GetDefaultConfidenceIntervalOk() (*string, bool)`

GetDefaultConfidenceIntervalOk returns a tuple with the DefaultConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfidenceInterval

`func (o *ExperimentFullUpdateDto) SetDefaultConfidenceInterval(v string)`

SetDefaultConfidenceInterval sets DefaultConfidenceInterval field to given value.


### GetStatus

`func (o *ExperimentFullUpdateDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExperimentFullUpdateDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExperimentFullUpdateDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLaunchedGroupID

`func (o *ExperimentFullUpdateDto) GetLaunchedGroupID() nil`

GetLaunchedGroupID returns the LaunchedGroupID field if non-nil, zero value otherwise.

### GetLaunchedGroupIDOk

`func (o *ExperimentFullUpdateDto) GetLaunchedGroupIDOk() (*nil, bool)`

GetLaunchedGroupIDOk returns a tuple with the LaunchedGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedGroupID

`func (o *ExperimentFullUpdateDto) SetLaunchedGroupID(v nil)`

SetLaunchedGroupID sets LaunchedGroupID field to given value.

### HasLaunchedGroupID

`func (o *ExperimentFullUpdateDto) HasLaunchedGroupID() bool`

HasLaunchedGroupID returns a boolean if a field has been set.

### GetAssignmentSourceName

`func (o *ExperimentFullUpdateDto) GetAssignmentSourceName() string`

GetAssignmentSourceName returns the AssignmentSourceName field if non-nil, zero value otherwise.

### GetAssignmentSourceNameOk

`func (o *ExperimentFullUpdateDto) GetAssignmentSourceNameOk() (*string, bool)`

GetAssignmentSourceNameOk returns a tuple with the AssignmentSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceName

`func (o *ExperimentFullUpdateDto) SetAssignmentSourceName(v string)`

SetAssignmentSourceName sets AssignmentSourceName field to given value.

### HasAssignmentSourceName

`func (o *ExperimentFullUpdateDto) HasAssignmentSourceName() bool`

HasAssignmentSourceName returns a boolean if a field has been set.

### GetAssignmentSourceExperimentName

`func (o *ExperimentFullUpdateDto) GetAssignmentSourceExperimentName() string`

GetAssignmentSourceExperimentName returns the AssignmentSourceExperimentName field if non-nil, zero value otherwise.

### GetAssignmentSourceExperimentNameOk

`func (o *ExperimentFullUpdateDto) GetAssignmentSourceExperimentNameOk() (*string, bool)`

GetAssignmentSourceExperimentNameOk returns a tuple with the AssignmentSourceExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceExperimentName

`func (o *ExperimentFullUpdateDto) SetAssignmentSourceExperimentName(v string)`

SetAssignmentSourceExperimentName sets AssignmentSourceExperimentName field to given value.

### HasAssignmentSourceExperimentName

`func (o *ExperimentFullUpdateDto) HasAssignmentSourceExperimentName() bool`

HasAssignmentSourceExperimentName returns a boolean if a field has been set.

### GetCreatorID

`func (o *ExperimentFullUpdateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *ExperimentFullUpdateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *ExperimentFullUpdateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *ExperimentFullUpdateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *ExperimentFullUpdateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *ExperimentFullUpdateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *ExperimentFullUpdateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *ExperimentFullUpdateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetIsAnalysisOnly

`func (o *ExperimentFullUpdateDto) GetIsAnalysisOnly() nil`

GetIsAnalysisOnly returns the IsAnalysisOnly field if non-nil, zero value otherwise.

### GetIsAnalysisOnlyOk

`func (o *ExperimentFullUpdateDto) GetIsAnalysisOnlyOk() (*nil, bool)`

GetIsAnalysisOnlyOk returns a tuple with the IsAnalysisOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalysisOnly

`func (o *ExperimentFullUpdateDto) SetIsAnalysisOnly(v nil)`

SetIsAnalysisOnly sets IsAnalysisOnly field to given value.

### HasIsAnalysisOnly

`func (o *ExperimentFullUpdateDto) HasIsAnalysisOnly() bool`

HasIsAnalysisOnly returns a boolean if a field has been set.

### GetTeam

`func (o *ExperimentFullUpdateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExperimentFullUpdateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExperimentFullUpdateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExperimentFullUpdateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetAllocationDuration

`func (o *ExperimentFullUpdateDto) GetAllocationDuration() int32`

GetAllocationDuration returns the AllocationDuration field if non-nil, zero value otherwise.

### GetAllocationDurationOk

`func (o *ExperimentFullUpdateDto) GetAllocationDurationOk() (*int32, bool)`

GetAllocationDurationOk returns a tuple with the AllocationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDuration

`func (o *ExperimentFullUpdateDto) SetAllocationDuration(v int32)`

SetAllocationDuration sets AllocationDuration field to given value.

### HasAllocationDuration

`func (o *ExperimentFullUpdateDto) HasAllocationDuration() bool`

HasAllocationDuration returns a boolean if a field has been set.

### GetCohortedAnalysisDuration

`func (o *ExperimentFullUpdateDto) GetCohortedAnalysisDuration() int32`

GetCohortedAnalysisDuration returns the CohortedAnalysisDuration field if non-nil, zero value otherwise.

### GetCohortedAnalysisDurationOk

`func (o *ExperimentFullUpdateDto) GetCohortedAnalysisDurationOk() (*int32, bool)`

GetCohortedAnalysisDurationOk returns a tuple with the CohortedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortedAnalysisDuration

`func (o *ExperimentFullUpdateDto) SetCohortedAnalysisDuration(v int32)`

SetCohortedAnalysisDuration sets CohortedAnalysisDuration field to given value.

### HasCohortedAnalysisDuration

`func (o *ExperimentFullUpdateDto) HasCohortedAnalysisDuration() bool`

HasCohortedAnalysisDuration returns a boolean if a field has been set.

### GetFixedAnalysisDuration

`func (o *ExperimentFullUpdateDto) GetFixedAnalysisDuration() int32`

GetFixedAnalysisDuration returns the FixedAnalysisDuration field if non-nil, zero value otherwise.

### GetFixedAnalysisDurationOk

`func (o *ExperimentFullUpdateDto) GetFixedAnalysisDurationOk() (*int32, bool)`

GetFixedAnalysisDurationOk returns a tuple with the FixedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAnalysisDuration

`func (o *ExperimentFullUpdateDto) SetFixedAnalysisDuration(v int32)`

SetFixedAnalysisDuration sets FixedAnalysisDuration field to given value.

### HasFixedAnalysisDuration

`func (o *ExperimentFullUpdateDto) HasFixedAnalysisDuration() bool`

HasFixedAnalysisDuration returns a boolean if a field has been set.

### GetScheduledReloadHour

`func (o *ExperimentFullUpdateDto) GetScheduledReloadHour() int32`

GetScheduledReloadHour returns the ScheduledReloadHour field if non-nil, zero value otherwise.

### GetScheduledReloadHourOk

`func (o *ExperimentFullUpdateDto) GetScheduledReloadHourOk() (*int32, bool)`

GetScheduledReloadHourOk returns a tuple with the ScheduledReloadHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadHour

`func (o *ExperimentFullUpdateDto) SetScheduledReloadHour(v int32)`

SetScheduledReloadHour sets ScheduledReloadHour field to given value.

### HasScheduledReloadHour

`func (o *ExperimentFullUpdateDto) HasScheduledReloadHour() bool`

HasScheduledReloadHour returns a boolean if a field has been set.

### GetScheduledReloadType

`func (o *ExperimentFullUpdateDto) GetScheduledReloadType() string`

GetScheduledReloadType returns the ScheduledReloadType field if non-nil, zero value otherwise.

### GetScheduledReloadTypeOk

`func (o *ExperimentFullUpdateDto) GetScheduledReloadTypeOk() (*string, bool)`

GetScheduledReloadTypeOk returns a tuple with the ScheduledReloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadType

`func (o *ExperimentFullUpdateDto) SetScheduledReloadType(v string)`

SetScheduledReloadType sets ScheduledReloadType field to given value.

### HasScheduledReloadType

`func (o *ExperimentFullUpdateDto) HasScheduledReloadType() bool`

HasScheduledReloadType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


