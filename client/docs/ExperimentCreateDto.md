# ExperimentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the new experiment | [optional] 
**IdType** | Pointer to **string** | The idType the experiment will be performed on | [optional] 
**SecondaryIDType** | Pointer to [**nil**](nil.md) | The secondary ID type for the experiment used in WHN for ID resolution | [optional] 
**Hypothesis** | Pointer to **string** | A statement that will be tested by this experiment | [optional] 
**Links** | Pointer to [**[]ExternalExperimentDtoLinksInner**](ExternalExperimentDtoLinksInner.md) | Links to relevant documentation or resources | [optional] 
**Groups** | Pointer to [**[]ExternalExperimentDtoGroupsInner**](ExternalExperimentDtoGroupsInner.md) | The test groups for your experiment | [optional] 
**ControlGroupID** | Pointer to **string** | Optional control group ID | [optional] 
**Allocation** | Pointer to **float32** | Percent of layer allocated to this experiment | [optional] 
**PrimaryMetricTags** | Pointer to **[]string** | Primary metric tags for the experiment | [optional] 
**SecondaryMetricTags** | Pointer to **[]string** | Secondary metric tags for the experiment | [optional] 
**PrimaryMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) | Main metrics needed to evaluate your hypothesis | [optional] 
**SecondaryMetrics** | Pointer to [**[]ExternalGateDtoMonitoringMetricsInner**](ExternalGateDtoMonitoringMetricsInner.md) | Additional metrics to monitor that might impact the analysis or final decision of the experiment | [optional] 
**TargetApps** | Pointer to [**ExternalExperimentDtoTargetApps**](ExternalExperimentDtoTargetApps.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the experiment | [optional] 
**Duration** | Pointer to **int32** | How long the experiment is expected to last in days | [optional] 
**TargetExposures** | Pointer to **int32** | Target exposures for the experiment | [optional] 
**TargetingGateID** | Pointer to [**nil**](nil.md) | Restrict your experiment to users passing the selected feature gate | [optional] 
**BonferroniCorrection** | Pointer to **bool** | Is Bonferroni correction applied? | [optional] 
**DefaultConfidenceInterval** | Pointer to **string** | Default error margin used for results | [optional] 
**Status** | Pointer to **string** | The current status of the experiment | [optional] 
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
**Name** | **string** | The name of the new experiment | 
**Id** | Pointer to **string** |  | [optional] 
**LayerID** | Pointer to **string** | Which layer to place the experiment into. | [optional] 
**ReviewSettings** | Pointer to [**ExternalGateDtoReviewSettings**](ExternalGateDtoReviewSettings.md) |  | [optional] 
**ActiveReview** | Pointer to [**ExternalGateDtoActiveReview**](ExternalGateDtoActiveReview.md) |  | [optional] 

## Methods

### NewExperimentCreateDto

`func NewExperimentCreateDto(name string, ) *ExperimentCreateDto`

NewExperimentCreateDto instantiates a new ExperimentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentCreateDtoWithDefaults

`func NewExperimentCreateDtoWithDefaults() *ExperimentCreateDto`

NewExperimentCreateDtoWithDefaults instantiates a new ExperimentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ExperimentCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperimentCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperimentCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExperimentCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdType

`func (o *ExperimentCreateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *ExperimentCreateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *ExperimentCreateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *ExperimentCreateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetSecondaryIDType

`func (o *ExperimentCreateDto) GetSecondaryIDType() nil`

GetSecondaryIDType returns the SecondaryIDType field if non-nil, zero value otherwise.

### GetSecondaryIDTypeOk

`func (o *ExperimentCreateDto) GetSecondaryIDTypeOk() (*nil, bool)`

GetSecondaryIDTypeOk returns a tuple with the SecondaryIDType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIDType

`func (o *ExperimentCreateDto) SetSecondaryIDType(v nil)`

SetSecondaryIDType sets SecondaryIDType field to given value.

### HasSecondaryIDType

`func (o *ExperimentCreateDto) HasSecondaryIDType() bool`

HasSecondaryIDType returns a boolean if a field has been set.

### GetHypothesis

`func (o *ExperimentCreateDto) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *ExperimentCreateDto) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *ExperimentCreateDto) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.

### HasHypothesis

`func (o *ExperimentCreateDto) HasHypothesis() bool`

HasHypothesis returns a boolean if a field has been set.

### GetLinks

`func (o *ExperimentCreateDto) GetLinks() []ExternalExperimentDtoLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentCreateDto) GetLinksOk() (*[]ExternalExperimentDtoLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentCreateDto) SetLinks(v []ExternalExperimentDtoLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentCreateDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetGroups

`func (o *ExperimentCreateDto) GetGroups() []ExternalExperimentDtoGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ExperimentCreateDto) GetGroupsOk() (*[]ExternalExperimentDtoGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ExperimentCreateDto) SetGroups(v []ExternalExperimentDtoGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ExperimentCreateDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetControlGroupID

`func (o *ExperimentCreateDto) GetControlGroupID() string`

GetControlGroupID returns the ControlGroupID field if non-nil, zero value otherwise.

### GetControlGroupIDOk

`func (o *ExperimentCreateDto) GetControlGroupIDOk() (*string, bool)`

GetControlGroupIDOk returns a tuple with the ControlGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlGroupID

`func (o *ExperimentCreateDto) SetControlGroupID(v string)`

SetControlGroupID sets ControlGroupID field to given value.

### HasControlGroupID

`func (o *ExperimentCreateDto) HasControlGroupID() bool`

HasControlGroupID returns a boolean if a field has been set.

### GetAllocation

`func (o *ExperimentCreateDto) GetAllocation() float32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *ExperimentCreateDto) GetAllocationOk() (*float32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *ExperimentCreateDto) SetAllocation(v float32)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *ExperimentCreateDto) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.

### GetPrimaryMetricTags

`func (o *ExperimentCreateDto) GetPrimaryMetricTags() []string`

GetPrimaryMetricTags returns the PrimaryMetricTags field if non-nil, zero value otherwise.

### GetPrimaryMetricTagsOk

`func (o *ExperimentCreateDto) GetPrimaryMetricTagsOk() (*[]string, bool)`

GetPrimaryMetricTagsOk returns a tuple with the PrimaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetricTags

`func (o *ExperimentCreateDto) SetPrimaryMetricTags(v []string)`

SetPrimaryMetricTags sets PrimaryMetricTags field to given value.

### HasPrimaryMetricTags

`func (o *ExperimentCreateDto) HasPrimaryMetricTags() bool`

HasPrimaryMetricTags returns a boolean if a field has been set.

### GetSecondaryMetricTags

`func (o *ExperimentCreateDto) GetSecondaryMetricTags() []string`

GetSecondaryMetricTags returns the SecondaryMetricTags field if non-nil, zero value otherwise.

### GetSecondaryMetricTagsOk

`func (o *ExperimentCreateDto) GetSecondaryMetricTagsOk() (*[]string, bool)`

GetSecondaryMetricTagsOk returns a tuple with the SecondaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetricTags

`func (o *ExperimentCreateDto) SetSecondaryMetricTags(v []string)`

SetSecondaryMetricTags sets SecondaryMetricTags field to given value.

### HasSecondaryMetricTags

`func (o *ExperimentCreateDto) HasSecondaryMetricTags() bool`

HasSecondaryMetricTags returns a boolean if a field has been set.

### GetPrimaryMetrics

`func (o *ExperimentCreateDto) GetPrimaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetPrimaryMetrics returns the PrimaryMetrics field if non-nil, zero value otherwise.

### GetPrimaryMetricsOk

`func (o *ExperimentCreateDto) GetPrimaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetPrimaryMetricsOk returns a tuple with the PrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetrics

`func (o *ExperimentCreateDto) SetPrimaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetPrimaryMetrics sets PrimaryMetrics field to given value.

### HasPrimaryMetrics

`func (o *ExperimentCreateDto) HasPrimaryMetrics() bool`

HasPrimaryMetrics returns a boolean if a field has been set.

### GetSecondaryMetrics

`func (o *ExperimentCreateDto) GetSecondaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *ExperimentCreateDto) GetSecondaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *ExperimentCreateDto) SetSecondaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *ExperimentCreateDto) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.

### GetTargetApps

`func (o *ExperimentCreateDto) GetTargetApps() ExternalExperimentDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *ExperimentCreateDto) GetTargetAppsOk() (*ExternalExperimentDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *ExperimentCreateDto) SetTargetApps(v ExternalExperimentDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *ExperimentCreateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetTags

`func (o *ExperimentCreateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExperimentCreateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExperimentCreateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExperimentCreateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDuration

`func (o *ExperimentCreateDto) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ExperimentCreateDto) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ExperimentCreateDto) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ExperimentCreateDto) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTargetExposures

`func (o *ExperimentCreateDto) GetTargetExposures() int32`

GetTargetExposures returns the TargetExposures field if non-nil, zero value otherwise.

### GetTargetExposuresOk

`func (o *ExperimentCreateDto) GetTargetExposuresOk() (*int32, bool)`

GetTargetExposuresOk returns a tuple with the TargetExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetExposures

`func (o *ExperimentCreateDto) SetTargetExposures(v int32)`

SetTargetExposures sets TargetExposures field to given value.

### HasTargetExposures

`func (o *ExperimentCreateDto) HasTargetExposures() bool`

HasTargetExposures returns a boolean if a field has been set.

### GetTargetingGateID

`func (o *ExperimentCreateDto) GetTargetingGateID() nil`

GetTargetingGateID returns the TargetingGateID field if non-nil, zero value otherwise.

### GetTargetingGateIDOk

`func (o *ExperimentCreateDto) GetTargetingGateIDOk() (*nil, bool)`

GetTargetingGateIDOk returns a tuple with the TargetingGateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingGateID

`func (o *ExperimentCreateDto) SetTargetingGateID(v nil)`

SetTargetingGateID sets TargetingGateID field to given value.

### HasTargetingGateID

`func (o *ExperimentCreateDto) HasTargetingGateID() bool`

HasTargetingGateID returns a boolean if a field has been set.

### GetBonferroniCorrection

`func (o *ExperimentCreateDto) GetBonferroniCorrection() bool`

GetBonferroniCorrection returns the BonferroniCorrection field if non-nil, zero value otherwise.

### GetBonferroniCorrectionOk

`func (o *ExperimentCreateDto) GetBonferroniCorrectionOk() (*bool, bool)`

GetBonferroniCorrectionOk returns a tuple with the BonferroniCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonferroniCorrection

`func (o *ExperimentCreateDto) SetBonferroniCorrection(v bool)`

SetBonferroniCorrection sets BonferroniCorrection field to given value.

### HasBonferroniCorrection

`func (o *ExperimentCreateDto) HasBonferroniCorrection() bool`

HasBonferroniCorrection returns a boolean if a field has been set.

### GetDefaultConfidenceInterval

`func (o *ExperimentCreateDto) GetDefaultConfidenceInterval() string`

GetDefaultConfidenceInterval returns the DefaultConfidenceInterval field if non-nil, zero value otherwise.

### GetDefaultConfidenceIntervalOk

`func (o *ExperimentCreateDto) GetDefaultConfidenceIntervalOk() (*string, bool)`

GetDefaultConfidenceIntervalOk returns a tuple with the DefaultConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfidenceInterval

`func (o *ExperimentCreateDto) SetDefaultConfidenceInterval(v string)`

SetDefaultConfidenceInterval sets DefaultConfidenceInterval field to given value.

### HasDefaultConfidenceInterval

`func (o *ExperimentCreateDto) HasDefaultConfidenceInterval() bool`

HasDefaultConfidenceInterval returns a boolean if a field has been set.

### GetStatus

`func (o *ExperimentCreateDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExperimentCreateDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExperimentCreateDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExperimentCreateDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLaunchedGroupID

`func (o *ExperimentCreateDto) GetLaunchedGroupID() nil`

GetLaunchedGroupID returns the LaunchedGroupID field if non-nil, zero value otherwise.

### GetLaunchedGroupIDOk

`func (o *ExperimentCreateDto) GetLaunchedGroupIDOk() (*nil, bool)`

GetLaunchedGroupIDOk returns a tuple with the LaunchedGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedGroupID

`func (o *ExperimentCreateDto) SetLaunchedGroupID(v nil)`

SetLaunchedGroupID sets LaunchedGroupID field to given value.

### HasLaunchedGroupID

`func (o *ExperimentCreateDto) HasLaunchedGroupID() bool`

HasLaunchedGroupID returns a boolean if a field has been set.

### GetAssignmentSourceName

`func (o *ExperimentCreateDto) GetAssignmentSourceName() string`

GetAssignmentSourceName returns the AssignmentSourceName field if non-nil, zero value otherwise.

### GetAssignmentSourceNameOk

`func (o *ExperimentCreateDto) GetAssignmentSourceNameOk() (*string, bool)`

GetAssignmentSourceNameOk returns a tuple with the AssignmentSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceName

`func (o *ExperimentCreateDto) SetAssignmentSourceName(v string)`

SetAssignmentSourceName sets AssignmentSourceName field to given value.

### HasAssignmentSourceName

`func (o *ExperimentCreateDto) HasAssignmentSourceName() bool`

HasAssignmentSourceName returns a boolean if a field has been set.

### GetAssignmentSourceExperimentName

`func (o *ExperimentCreateDto) GetAssignmentSourceExperimentName() string`

GetAssignmentSourceExperimentName returns the AssignmentSourceExperimentName field if non-nil, zero value otherwise.

### GetAssignmentSourceExperimentNameOk

`func (o *ExperimentCreateDto) GetAssignmentSourceExperimentNameOk() (*string, bool)`

GetAssignmentSourceExperimentNameOk returns a tuple with the AssignmentSourceExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceExperimentName

`func (o *ExperimentCreateDto) SetAssignmentSourceExperimentName(v string)`

SetAssignmentSourceExperimentName sets AssignmentSourceExperimentName field to given value.

### HasAssignmentSourceExperimentName

`func (o *ExperimentCreateDto) HasAssignmentSourceExperimentName() bool`

HasAssignmentSourceExperimentName returns a boolean if a field has been set.

### GetCreatorID

`func (o *ExperimentCreateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *ExperimentCreateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *ExperimentCreateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *ExperimentCreateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *ExperimentCreateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *ExperimentCreateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *ExperimentCreateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *ExperimentCreateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetIsAnalysisOnly

`func (o *ExperimentCreateDto) GetIsAnalysisOnly() nil`

GetIsAnalysisOnly returns the IsAnalysisOnly field if non-nil, zero value otherwise.

### GetIsAnalysisOnlyOk

`func (o *ExperimentCreateDto) GetIsAnalysisOnlyOk() (*nil, bool)`

GetIsAnalysisOnlyOk returns a tuple with the IsAnalysisOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalysisOnly

`func (o *ExperimentCreateDto) SetIsAnalysisOnly(v nil)`

SetIsAnalysisOnly sets IsAnalysisOnly field to given value.

### HasIsAnalysisOnly

`func (o *ExperimentCreateDto) HasIsAnalysisOnly() bool`

HasIsAnalysisOnly returns a boolean if a field has been set.

### GetTeam

`func (o *ExperimentCreateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExperimentCreateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExperimentCreateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExperimentCreateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetAllocationDuration

`func (o *ExperimentCreateDto) GetAllocationDuration() int32`

GetAllocationDuration returns the AllocationDuration field if non-nil, zero value otherwise.

### GetAllocationDurationOk

`func (o *ExperimentCreateDto) GetAllocationDurationOk() (*int32, bool)`

GetAllocationDurationOk returns a tuple with the AllocationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDuration

`func (o *ExperimentCreateDto) SetAllocationDuration(v int32)`

SetAllocationDuration sets AllocationDuration field to given value.

### HasAllocationDuration

`func (o *ExperimentCreateDto) HasAllocationDuration() bool`

HasAllocationDuration returns a boolean if a field has been set.

### GetCohortedAnalysisDuration

`func (o *ExperimentCreateDto) GetCohortedAnalysisDuration() int32`

GetCohortedAnalysisDuration returns the CohortedAnalysisDuration field if non-nil, zero value otherwise.

### GetCohortedAnalysisDurationOk

`func (o *ExperimentCreateDto) GetCohortedAnalysisDurationOk() (*int32, bool)`

GetCohortedAnalysisDurationOk returns a tuple with the CohortedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortedAnalysisDuration

`func (o *ExperimentCreateDto) SetCohortedAnalysisDuration(v int32)`

SetCohortedAnalysisDuration sets CohortedAnalysisDuration field to given value.

### HasCohortedAnalysisDuration

`func (o *ExperimentCreateDto) HasCohortedAnalysisDuration() bool`

HasCohortedAnalysisDuration returns a boolean if a field has been set.

### GetFixedAnalysisDuration

`func (o *ExperimentCreateDto) GetFixedAnalysisDuration() int32`

GetFixedAnalysisDuration returns the FixedAnalysisDuration field if non-nil, zero value otherwise.

### GetFixedAnalysisDurationOk

`func (o *ExperimentCreateDto) GetFixedAnalysisDurationOk() (*int32, bool)`

GetFixedAnalysisDurationOk returns a tuple with the FixedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAnalysisDuration

`func (o *ExperimentCreateDto) SetFixedAnalysisDuration(v int32)`

SetFixedAnalysisDuration sets FixedAnalysisDuration field to given value.

### HasFixedAnalysisDuration

`func (o *ExperimentCreateDto) HasFixedAnalysisDuration() bool`

HasFixedAnalysisDuration returns a boolean if a field has been set.

### GetScheduledReloadHour

`func (o *ExperimentCreateDto) GetScheduledReloadHour() int32`

GetScheduledReloadHour returns the ScheduledReloadHour field if non-nil, zero value otherwise.

### GetScheduledReloadHourOk

`func (o *ExperimentCreateDto) GetScheduledReloadHourOk() (*int32, bool)`

GetScheduledReloadHourOk returns a tuple with the ScheduledReloadHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadHour

`func (o *ExperimentCreateDto) SetScheduledReloadHour(v int32)`

SetScheduledReloadHour sets ScheduledReloadHour field to given value.

### HasScheduledReloadHour

`func (o *ExperimentCreateDto) HasScheduledReloadHour() bool`

HasScheduledReloadHour returns a boolean if a field has been set.

### GetScheduledReloadType

`func (o *ExperimentCreateDto) GetScheduledReloadType() string`

GetScheduledReloadType returns the ScheduledReloadType field if non-nil, zero value otherwise.

### GetScheduledReloadTypeOk

`func (o *ExperimentCreateDto) GetScheduledReloadTypeOk() (*string, bool)`

GetScheduledReloadTypeOk returns a tuple with the ScheduledReloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadType

`func (o *ExperimentCreateDto) SetScheduledReloadType(v string)`

SetScheduledReloadType sets ScheduledReloadType field to given value.

### HasScheduledReloadType

`func (o *ExperimentCreateDto) HasScheduledReloadType() bool`

HasScheduledReloadType returns a boolean if a field has been set.

### GetName

`func (o *ExperimentCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ExperimentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExperimentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLayerID

`func (o *ExperimentCreateDto) GetLayerID() string`

GetLayerID returns the LayerID field if non-nil, zero value otherwise.

### GetLayerIDOk

`func (o *ExperimentCreateDto) GetLayerIDOk() (*string, bool)`

GetLayerIDOk returns a tuple with the LayerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerID

`func (o *ExperimentCreateDto) SetLayerID(v string)`

SetLayerID sets LayerID field to given value.

### HasLayerID

`func (o *ExperimentCreateDto) HasLayerID() bool`

HasLayerID returns a boolean if a field has been set.

### GetReviewSettings

`func (o *ExperimentCreateDto) GetReviewSettings() ExternalGateDtoReviewSettings`

GetReviewSettings returns the ReviewSettings field if non-nil, zero value otherwise.

### GetReviewSettingsOk

`func (o *ExperimentCreateDto) GetReviewSettingsOk() (*ExternalGateDtoReviewSettings, bool)`

GetReviewSettingsOk returns a tuple with the ReviewSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewSettings

`func (o *ExperimentCreateDto) SetReviewSettings(v ExternalGateDtoReviewSettings)`

SetReviewSettings sets ReviewSettings field to given value.

### HasReviewSettings

`func (o *ExperimentCreateDto) HasReviewSettings() bool`

HasReviewSettings returns a boolean if a field has been set.

### GetActiveReview

`func (o *ExperimentCreateDto) GetActiveReview() ExternalGateDtoActiveReview`

GetActiveReview returns the ActiveReview field if non-nil, zero value otherwise.

### GetActiveReviewOk

`func (o *ExperimentCreateDto) GetActiveReviewOk() (*ExternalGateDtoActiveReview, bool)`

GetActiveReviewOk returns a tuple with the ActiveReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveReview

`func (o *ExperimentCreateDto) SetActiveReview(v ExternalGateDtoActiveReview)`

SetActiveReview sets ActiveReview field to given value.

### HasActiveReview

`func (o *ExperimentCreateDto) HasActiveReview() bool`

HasActiveReview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


