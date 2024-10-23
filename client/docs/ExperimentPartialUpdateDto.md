# ExperimentPartialUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A helpful summary of what this experiment does | [optional] 
**IdType** | Pointer to **string** | The type of ID which the experiment is based on | [optional] 
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

## Methods

### NewExperimentPartialUpdateDto

`func NewExperimentPartialUpdateDto() *ExperimentPartialUpdateDto`

NewExperimentPartialUpdateDto instantiates a new ExperimentPartialUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentPartialUpdateDtoWithDefaults

`func NewExperimentPartialUpdateDtoWithDefaults() *ExperimentPartialUpdateDto`

NewExperimentPartialUpdateDtoWithDefaults instantiates a new ExperimentPartialUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ExperimentPartialUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperimentPartialUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperimentPartialUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExperimentPartialUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdType

`func (o *ExperimentPartialUpdateDto) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *ExperimentPartialUpdateDto) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *ExperimentPartialUpdateDto) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *ExperimentPartialUpdateDto) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetSecondaryIDType

`func (o *ExperimentPartialUpdateDto) GetSecondaryIDType() nil`

GetSecondaryIDType returns the SecondaryIDType field if non-nil, zero value otherwise.

### GetSecondaryIDTypeOk

`func (o *ExperimentPartialUpdateDto) GetSecondaryIDTypeOk() (*nil, bool)`

GetSecondaryIDTypeOk returns a tuple with the SecondaryIDType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIDType

`func (o *ExperimentPartialUpdateDto) SetSecondaryIDType(v nil)`

SetSecondaryIDType sets SecondaryIDType field to given value.

### HasSecondaryIDType

`func (o *ExperimentPartialUpdateDto) HasSecondaryIDType() bool`

HasSecondaryIDType returns a boolean if a field has been set.

### GetHypothesis

`func (o *ExperimentPartialUpdateDto) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *ExperimentPartialUpdateDto) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *ExperimentPartialUpdateDto) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.

### HasHypothesis

`func (o *ExperimentPartialUpdateDto) HasHypothesis() bool`

HasHypothesis returns a boolean if a field has been set.

### GetLinks

`func (o *ExperimentPartialUpdateDto) GetLinks() []ExternalExperimentDtoLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentPartialUpdateDto) GetLinksOk() (*[]ExternalExperimentDtoLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentPartialUpdateDto) SetLinks(v []ExternalExperimentDtoLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentPartialUpdateDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetGroups

`func (o *ExperimentPartialUpdateDto) GetGroups() []ExternalExperimentDtoGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ExperimentPartialUpdateDto) GetGroupsOk() (*[]ExternalExperimentDtoGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ExperimentPartialUpdateDto) SetGroups(v []ExternalExperimentDtoGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ExperimentPartialUpdateDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetControlGroupID

`func (o *ExperimentPartialUpdateDto) GetControlGroupID() string`

GetControlGroupID returns the ControlGroupID field if non-nil, zero value otherwise.

### GetControlGroupIDOk

`func (o *ExperimentPartialUpdateDto) GetControlGroupIDOk() (*string, bool)`

GetControlGroupIDOk returns a tuple with the ControlGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlGroupID

`func (o *ExperimentPartialUpdateDto) SetControlGroupID(v string)`

SetControlGroupID sets ControlGroupID field to given value.

### HasControlGroupID

`func (o *ExperimentPartialUpdateDto) HasControlGroupID() bool`

HasControlGroupID returns a boolean if a field has been set.

### GetAllocation

`func (o *ExperimentPartialUpdateDto) GetAllocation() float32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *ExperimentPartialUpdateDto) GetAllocationOk() (*float32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *ExperimentPartialUpdateDto) SetAllocation(v float32)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *ExperimentPartialUpdateDto) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.

### GetPrimaryMetricTags

`func (o *ExperimentPartialUpdateDto) GetPrimaryMetricTags() []string`

GetPrimaryMetricTags returns the PrimaryMetricTags field if non-nil, zero value otherwise.

### GetPrimaryMetricTagsOk

`func (o *ExperimentPartialUpdateDto) GetPrimaryMetricTagsOk() (*[]string, bool)`

GetPrimaryMetricTagsOk returns a tuple with the PrimaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetricTags

`func (o *ExperimentPartialUpdateDto) SetPrimaryMetricTags(v []string)`

SetPrimaryMetricTags sets PrimaryMetricTags field to given value.

### HasPrimaryMetricTags

`func (o *ExperimentPartialUpdateDto) HasPrimaryMetricTags() bool`

HasPrimaryMetricTags returns a boolean if a field has been set.

### GetSecondaryMetricTags

`func (o *ExperimentPartialUpdateDto) GetSecondaryMetricTags() []string`

GetSecondaryMetricTags returns the SecondaryMetricTags field if non-nil, zero value otherwise.

### GetSecondaryMetricTagsOk

`func (o *ExperimentPartialUpdateDto) GetSecondaryMetricTagsOk() (*[]string, bool)`

GetSecondaryMetricTagsOk returns a tuple with the SecondaryMetricTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetricTags

`func (o *ExperimentPartialUpdateDto) SetSecondaryMetricTags(v []string)`

SetSecondaryMetricTags sets SecondaryMetricTags field to given value.

### HasSecondaryMetricTags

`func (o *ExperimentPartialUpdateDto) HasSecondaryMetricTags() bool`

HasSecondaryMetricTags returns a boolean if a field has been set.

### GetPrimaryMetrics

`func (o *ExperimentPartialUpdateDto) GetPrimaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetPrimaryMetrics returns the PrimaryMetrics field if non-nil, zero value otherwise.

### GetPrimaryMetricsOk

`func (o *ExperimentPartialUpdateDto) GetPrimaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetPrimaryMetricsOk returns a tuple with the PrimaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetrics

`func (o *ExperimentPartialUpdateDto) SetPrimaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetPrimaryMetrics sets PrimaryMetrics field to given value.

### HasPrimaryMetrics

`func (o *ExperimentPartialUpdateDto) HasPrimaryMetrics() bool`

HasPrimaryMetrics returns a boolean if a field has been set.

### GetSecondaryMetrics

`func (o *ExperimentPartialUpdateDto) GetSecondaryMetrics() []ExternalGateDtoMonitoringMetricsInner`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *ExperimentPartialUpdateDto) GetSecondaryMetricsOk() (*[]ExternalGateDtoMonitoringMetricsInner, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *ExperimentPartialUpdateDto) SetSecondaryMetrics(v []ExternalGateDtoMonitoringMetricsInner)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *ExperimentPartialUpdateDto) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.

### GetTargetApps

`func (o *ExperimentPartialUpdateDto) GetTargetApps() ExternalExperimentDtoTargetApps`

GetTargetApps returns the TargetApps field if non-nil, zero value otherwise.

### GetTargetAppsOk

`func (o *ExperimentPartialUpdateDto) GetTargetAppsOk() (*ExternalExperimentDtoTargetApps, bool)`

GetTargetAppsOk returns a tuple with the TargetApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApps

`func (o *ExperimentPartialUpdateDto) SetTargetApps(v ExternalExperimentDtoTargetApps)`

SetTargetApps sets TargetApps field to given value.

### HasTargetApps

`func (o *ExperimentPartialUpdateDto) HasTargetApps() bool`

HasTargetApps returns a boolean if a field has been set.

### GetTags

`func (o *ExperimentPartialUpdateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExperimentPartialUpdateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExperimentPartialUpdateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExperimentPartialUpdateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDuration

`func (o *ExperimentPartialUpdateDto) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ExperimentPartialUpdateDto) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ExperimentPartialUpdateDto) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ExperimentPartialUpdateDto) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTargetExposures

`func (o *ExperimentPartialUpdateDto) GetTargetExposures() int32`

GetTargetExposures returns the TargetExposures field if non-nil, zero value otherwise.

### GetTargetExposuresOk

`func (o *ExperimentPartialUpdateDto) GetTargetExposuresOk() (*int32, bool)`

GetTargetExposuresOk returns a tuple with the TargetExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetExposures

`func (o *ExperimentPartialUpdateDto) SetTargetExposures(v int32)`

SetTargetExposures sets TargetExposures field to given value.

### HasTargetExposures

`func (o *ExperimentPartialUpdateDto) HasTargetExposures() bool`

HasTargetExposures returns a boolean if a field has been set.

### GetTargetingGateID

`func (o *ExperimentPartialUpdateDto) GetTargetingGateID() nil`

GetTargetingGateID returns the TargetingGateID field if non-nil, zero value otherwise.

### GetTargetingGateIDOk

`func (o *ExperimentPartialUpdateDto) GetTargetingGateIDOk() (*nil, bool)`

GetTargetingGateIDOk returns a tuple with the TargetingGateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingGateID

`func (o *ExperimentPartialUpdateDto) SetTargetingGateID(v nil)`

SetTargetingGateID sets TargetingGateID field to given value.

### HasTargetingGateID

`func (o *ExperimentPartialUpdateDto) HasTargetingGateID() bool`

HasTargetingGateID returns a boolean if a field has been set.

### GetBonferroniCorrection

`func (o *ExperimentPartialUpdateDto) GetBonferroniCorrection() bool`

GetBonferroniCorrection returns the BonferroniCorrection field if non-nil, zero value otherwise.

### GetBonferroniCorrectionOk

`func (o *ExperimentPartialUpdateDto) GetBonferroniCorrectionOk() (*bool, bool)`

GetBonferroniCorrectionOk returns a tuple with the BonferroniCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonferroniCorrection

`func (o *ExperimentPartialUpdateDto) SetBonferroniCorrection(v bool)`

SetBonferroniCorrection sets BonferroniCorrection field to given value.

### HasBonferroniCorrection

`func (o *ExperimentPartialUpdateDto) HasBonferroniCorrection() bool`

HasBonferroniCorrection returns a boolean if a field has been set.

### GetDefaultConfidenceInterval

`func (o *ExperimentPartialUpdateDto) GetDefaultConfidenceInterval() string`

GetDefaultConfidenceInterval returns the DefaultConfidenceInterval field if non-nil, zero value otherwise.

### GetDefaultConfidenceIntervalOk

`func (o *ExperimentPartialUpdateDto) GetDefaultConfidenceIntervalOk() (*string, bool)`

GetDefaultConfidenceIntervalOk returns a tuple with the DefaultConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfidenceInterval

`func (o *ExperimentPartialUpdateDto) SetDefaultConfidenceInterval(v string)`

SetDefaultConfidenceInterval sets DefaultConfidenceInterval field to given value.

### HasDefaultConfidenceInterval

`func (o *ExperimentPartialUpdateDto) HasDefaultConfidenceInterval() bool`

HasDefaultConfidenceInterval returns a boolean if a field has been set.

### GetStatus

`func (o *ExperimentPartialUpdateDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExperimentPartialUpdateDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExperimentPartialUpdateDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExperimentPartialUpdateDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLaunchedGroupID

`func (o *ExperimentPartialUpdateDto) GetLaunchedGroupID() nil`

GetLaunchedGroupID returns the LaunchedGroupID field if non-nil, zero value otherwise.

### GetLaunchedGroupIDOk

`func (o *ExperimentPartialUpdateDto) GetLaunchedGroupIDOk() (*nil, bool)`

GetLaunchedGroupIDOk returns a tuple with the LaunchedGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedGroupID

`func (o *ExperimentPartialUpdateDto) SetLaunchedGroupID(v nil)`

SetLaunchedGroupID sets LaunchedGroupID field to given value.

### HasLaunchedGroupID

`func (o *ExperimentPartialUpdateDto) HasLaunchedGroupID() bool`

HasLaunchedGroupID returns a boolean if a field has been set.

### GetAssignmentSourceName

`func (o *ExperimentPartialUpdateDto) GetAssignmentSourceName() string`

GetAssignmentSourceName returns the AssignmentSourceName field if non-nil, zero value otherwise.

### GetAssignmentSourceNameOk

`func (o *ExperimentPartialUpdateDto) GetAssignmentSourceNameOk() (*string, bool)`

GetAssignmentSourceNameOk returns a tuple with the AssignmentSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceName

`func (o *ExperimentPartialUpdateDto) SetAssignmentSourceName(v string)`

SetAssignmentSourceName sets AssignmentSourceName field to given value.

### HasAssignmentSourceName

`func (o *ExperimentPartialUpdateDto) HasAssignmentSourceName() bool`

HasAssignmentSourceName returns a boolean if a field has been set.

### GetAssignmentSourceExperimentName

`func (o *ExperimentPartialUpdateDto) GetAssignmentSourceExperimentName() string`

GetAssignmentSourceExperimentName returns the AssignmentSourceExperimentName field if non-nil, zero value otherwise.

### GetAssignmentSourceExperimentNameOk

`func (o *ExperimentPartialUpdateDto) GetAssignmentSourceExperimentNameOk() (*string, bool)`

GetAssignmentSourceExperimentNameOk returns a tuple with the AssignmentSourceExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceExperimentName

`func (o *ExperimentPartialUpdateDto) SetAssignmentSourceExperimentName(v string)`

SetAssignmentSourceExperimentName sets AssignmentSourceExperimentName field to given value.

### HasAssignmentSourceExperimentName

`func (o *ExperimentPartialUpdateDto) HasAssignmentSourceExperimentName() bool`

HasAssignmentSourceExperimentName returns a boolean if a field has been set.

### GetCreatorID

`func (o *ExperimentPartialUpdateDto) GetCreatorID() nil`

GetCreatorID returns the CreatorID field if non-nil, zero value otherwise.

### GetCreatorIDOk

`func (o *ExperimentPartialUpdateDto) GetCreatorIDOk() (*nil, bool)`

GetCreatorIDOk returns a tuple with the CreatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorID

`func (o *ExperimentPartialUpdateDto) SetCreatorID(v nil)`

SetCreatorID sets CreatorID field to given value.

### HasCreatorID

`func (o *ExperimentPartialUpdateDto) HasCreatorID() bool`

HasCreatorID returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *ExperimentPartialUpdateDto) GetCreatorEmail() nil`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *ExperimentPartialUpdateDto) GetCreatorEmailOk() (*nil, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *ExperimentPartialUpdateDto) SetCreatorEmail(v nil)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *ExperimentPartialUpdateDto) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetIsAnalysisOnly

`func (o *ExperimentPartialUpdateDto) GetIsAnalysisOnly() nil`

GetIsAnalysisOnly returns the IsAnalysisOnly field if non-nil, zero value otherwise.

### GetIsAnalysisOnlyOk

`func (o *ExperimentPartialUpdateDto) GetIsAnalysisOnlyOk() (*nil, bool)`

GetIsAnalysisOnlyOk returns a tuple with the IsAnalysisOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalysisOnly

`func (o *ExperimentPartialUpdateDto) SetIsAnalysisOnly(v nil)`

SetIsAnalysisOnly sets IsAnalysisOnly field to given value.

### HasIsAnalysisOnly

`func (o *ExperimentPartialUpdateDto) HasIsAnalysisOnly() bool`

HasIsAnalysisOnly returns a boolean if a field has been set.

### GetTeam

`func (o *ExperimentPartialUpdateDto) GetTeam() nil`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExperimentPartialUpdateDto) GetTeamOk() (*nil, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExperimentPartialUpdateDto) SetTeam(v nil)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExperimentPartialUpdateDto) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetAllocationDuration

`func (o *ExperimentPartialUpdateDto) GetAllocationDuration() int32`

GetAllocationDuration returns the AllocationDuration field if non-nil, zero value otherwise.

### GetAllocationDurationOk

`func (o *ExperimentPartialUpdateDto) GetAllocationDurationOk() (*int32, bool)`

GetAllocationDurationOk returns a tuple with the AllocationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDuration

`func (o *ExperimentPartialUpdateDto) SetAllocationDuration(v int32)`

SetAllocationDuration sets AllocationDuration field to given value.

### HasAllocationDuration

`func (o *ExperimentPartialUpdateDto) HasAllocationDuration() bool`

HasAllocationDuration returns a boolean if a field has been set.

### GetCohortedAnalysisDuration

`func (o *ExperimentPartialUpdateDto) GetCohortedAnalysisDuration() int32`

GetCohortedAnalysisDuration returns the CohortedAnalysisDuration field if non-nil, zero value otherwise.

### GetCohortedAnalysisDurationOk

`func (o *ExperimentPartialUpdateDto) GetCohortedAnalysisDurationOk() (*int32, bool)`

GetCohortedAnalysisDurationOk returns a tuple with the CohortedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortedAnalysisDuration

`func (o *ExperimentPartialUpdateDto) SetCohortedAnalysisDuration(v int32)`

SetCohortedAnalysisDuration sets CohortedAnalysisDuration field to given value.

### HasCohortedAnalysisDuration

`func (o *ExperimentPartialUpdateDto) HasCohortedAnalysisDuration() bool`

HasCohortedAnalysisDuration returns a boolean if a field has been set.

### GetFixedAnalysisDuration

`func (o *ExperimentPartialUpdateDto) GetFixedAnalysisDuration() int32`

GetFixedAnalysisDuration returns the FixedAnalysisDuration field if non-nil, zero value otherwise.

### GetFixedAnalysisDurationOk

`func (o *ExperimentPartialUpdateDto) GetFixedAnalysisDurationOk() (*int32, bool)`

GetFixedAnalysisDurationOk returns a tuple with the FixedAnalysisDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAnalysisDuration

`func (o *ExperimentPartialUpdateDto) SetFixedAnalysisDuration(v int32)`

SetFixedAnalysisDuration sets FixedAnalysisDuration field to given value.

### HasFixedAnalysisDuration

`func (o *ExperimentPartialUpdateDto) HasFixedAnalysisDuration() bool`

HasFixedAnalysisDuration returns a boolean if a field has been set.

### GetScheduledReloadHour

`func (o *ExperimentPartialUpdateDto) GetScheduledReloadHour() int32`

GetScheduledReloadHour returns the ScheduledReloadHour field if non-nil, zero value otherwise.

### GetScheduledReloadHourOk

`func (o *ExperimentPartialUpdateDto) GetScheduledReloadHourOk() (*int32, bool)`

GetScheduledReloadHourOk returns a tuple with the ScheduledReloadHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadHour

`func (o *ExperimentPartialUpdateDto) SetScheduledReloadHour(v int32)`

SetScheduledReloadHour sets ScheduledReloadHour field to given value.

### HasScheduledReloadHour

`func (o *ExperimentPartialUpdateDto) HasScheduledReloadHour() bool`

HasScheduledReloadHour returns a boolean if a field has been set.

### GetScheduledReloadType

`func (o *ExperimentPartialUpdateDto) GetScheduledReloadType() string`

GetScheduledReloadType returns the ScheduledReloadType field if non-nil, zero value otherwise.

### GetScheduledReloadTypeOk

`func (o *ExperimentPartialUpdateDto) GetScheduledReloadTypeOk() (*string, bool)`

GetScheduledReloadTypeOk returns a tuple with the ScheduledReloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReloadType

`func (o *ExperimentPartialUpdateDto) SetScheduledReloadType(v string)`

SetScheduledReloadType sets ScheduledReloadType field to given value.

### HasScheduledReloadType

`func (o *ExperimentPartialUpdateDto) HasScheduledReloadType() bool`

HasScheduledReloadType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


