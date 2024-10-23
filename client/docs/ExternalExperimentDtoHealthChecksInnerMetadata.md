# ExternalExperimentDtoHealthChecksInnerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**LastPulseLoadTime** | Pointer to **float32** |  | [optional] 
**CrossoverPercent** | Pointer to **float32** |  | [optional] 
**AssignmentSourceID** | Pointer to **string** |  | [optional] 
**AssignmentSourceName** | Pointer to **string** |  | [optional] 
**ForeignExperimentID** | Pointer to **string** |  | [optional] 
**DeduplicationRate** | Pointer to **float32** |  | [optional] 
**DeduplicationRates** | Pointer to [**[]ExternalExperimentDtoHealthChecksInnerMetadataDeduplicationRatesInner**](ExternalExperimentDtoHealthChecksInnerMetadataDeduplicationRatesInner.md) |  | [optional] 
**PrimaryIdType** | **string** |  | 
**SecondaryIdType** | Pointer to **string** |  | [optional] 
**MissingMetrics** | Pointer to **[]string** |  | [optional] 
**Metrics** | Pointer to [**[]ExternalExperimentDtoHealthChecksInnerMetadataMetricsInner**](ExternalExperimentDtoHealthChecksInnerMetadataMetricsInner.md) |  | [optional] 
**LastUpdatedDs** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalExperimentDtoHealthChecksInnerMetadata

`func NewExternalExperimentDtoHealthChecksInnerMetadata(type_ string, primaryIdType string, ) *ExternalExperimentDtoHealthChecksInnerMetadata`

NewExternalExperimentDtoHealthChecksInnerMetadata instantiates a new ExternalExperimentDtoHealthChecksInnerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalExperimentDtoHealthChecksInnerMetadataWithDefaults

`func NewExternalExperimentDtoHealthChecksInnerMetadataWithDefaults() *ExternalExperimentDtoHealthChecksInnerMetadata`

NewExternalExperimentDtoHealthChecksInnerMetadataWithDefaults instantiates a new ExternalExperimentDtoHealthChecksInnerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetLastPulseLoadTime

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetLastPulseLoadTime() float32`

GetLastPulseLoadTime returns the LastPulseLoadTime field if non-nil, zero value otherwise.

### GetLastPulseLoadTimeOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetLastPulseLoadTimeOk() (*float32, bool)`

GetLastPulseLoadTimeOk returns a tuple with the LastPulseLoadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPulseLoadTime

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetLastPulseLoadTime(v float32)`

SetLastPulseLoadTime sets LastPulseLoadTime field to given value.

### HasLastPulseLoadTime

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasLastPulseLoadTime() bool`

HasLastPulseLoadTime returns a boolean if a field has been set.

### GetCrossoverPercent

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetCrossoverPercent() float32`

GetCrossoverPercent returns the CrossoverPercent field if non-nil, zero value otherwise.

### GetCrossoverPercentOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetCrossoverPercentOk() (*float32, bool)`

GetCrossoverPercentOk returns a tuple with the CrossoverPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossoverPercent

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetCrossoverPercent(v float32)`

SetCrossoverPercent sets CrossoverPercent field to given value.

### HasCrossoverPercent

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasCrossoverPercent() bool`

HasCrossoverPercent returns a boolean if a field has been set.

### GetAssignmentSourceID

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetAssignmentSourceID() string`

GetAssignmentSourceID returns the AssignmentSourceID field if non-nil, zero value otherwise.

### GetAssignmentSourceIDOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetAssignmentSourceIDOk() (*string, bool)`

GetAssignmentSourceIDOk returns a tuple with the AssignmentSourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceID

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetAssignmentSourceID(v string)`

SetAssignmentSourceID sets AssignmentSourceID field to given value.

### HasAssignmentSourceID

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasAssignmentSourceID() bool`

HasAssignmentSourceID returns a boolean if a field has been set.

### GetAssignmentSourceName

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetAssignmentSourceName() string`

GetAssignmentSourceName returns the AssignmentSourceName field if non-nil, zero value otherwise.

### GetAssignmentSourceNameOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetAssignmentSourceNameOk() (*string, bool)`

GetAssignmentSourceNameOk returns a tuple with the AssignmentSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSourceName

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetAssignmentSourceName(v string)`

SetAssignmentSourceName sets AssignmentSourceName field to given value.

### HasAssignmentSourceName

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasAssignmentSourceName() bool`

HasAssignmentSourceName returns a boolean if a field has been set.

### GetForeignExperimentID

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetForeignExperimentID() string`

GetForeignExperimentID returns the ForeignExperimentID field if non-nil, zero value otherwise.

### GetForeignExperimentIDOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetForeignExperimentIDOk() (*string, bool)`

GetForeignExperimentIDOk returns a tuple with the ForeignExperimentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignExperimentID

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetForeignExperimentID(v string)`

SetForeignExperimentID sets ForeignExperimentID field to given value.

### HasForeignExperimentID

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasForeignExperimentID() bool`

HasForeignExperimentID returns a boolean if a field has been set.

### GetDeduplicationRate

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetDeduplicationRate() float32`

GetDeduplicationRate returns the DeduplicationRate field if non-nil, zero value otherwise.

### GetDeduplicationRateOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetDeduplicationRateOk() (*float32, bool)`

GetDeduplicationRateOk returns a tuple with the DeduplicationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationRate

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetDeduplicationRate(v float32)`

SetDeduplicationRate sets DeduplicationRate field to given value.

### HasDeduplicationRate

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasDeduplicationRate() bool`

HasDeduplicationRate returns a boolean if a field has been set.

### GetDeduplicationRates

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetDeduplicationRates() []ExternalExperimentDtoHealthChecksInnerMetadataDeduplicationRatesInner`

GetDeduplicationRates returns the DeduplicationRates field if non-nil, zero value otherwise.

### GetDeduplicationRatesOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetDeduplicationRatesOk() (*[]ExternalExperimentDtoHealthChecksInnerMetadataDeduplicationRatesInner, bool)`

GetDeduplicationRatesOk returns a tuple with the DeduplicationRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationRates

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetDeduplicationRates(v []ExternalExperimentDtoHealthChecksInnerMetadataDeduplicationRatesInner)`

SetDeduplicationRates sets DeduplicationRates field to given value.

### HasDeduplicationRates

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasDeduplicationRates() bool`

HasDeduplicationRates returns a boolean if a field has been set.

### GetPrimaryIdType

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetPrimaryIdType() string`

GetPrimaryIdType returns the PrimaryIdType field if non-nil, zero value otherwise.

### GetPrimaryIdTypeOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetPrimaryIdTypeOk() (*string, bool)`

GetPrimaryIdTypeOk returns a tuple with the PrimaryIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIdType

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetPrimaryIdType(v string)`

SetPrimaryIdType sets PrimaryIdType field to given value.


### GetSecondaryIdType

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetSecondaryIdType() string`

GetSecondaryIdType returns the SecondaryIdType field if non-nil, zero value otherwise.

### GetSecondaryIdTypeOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetSecondaryIdTypeOk() (*string, bool)`

GetSecondaryIdTypeOk returns a tuple with the SecondaryIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIdType

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetSecondaryIdType(v string)`

SetSecondaryIdType sets SecondaryIdType field to given value.

### HasSecondaryIdType

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasSecondaryIdType() bool`

HasSecondaryIdType returns a boolean if a field has been set.

### GetMissingMetrics

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetMissingMetrics() []string`

GetMissingMetrics returns the MissingMetrics field if non-nil, zero value otherwise.

### GetMissingMetricsOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetMissingMetricsOk() (*[]string, bool)`

GetMissingMetricsOk returns a tuple with the MissingMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingMetrics

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetMissingMetrics(v []string)`

SetMissingMetrics sets MissingMetrics field to given value.

### HasMissingMetrics

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasMissingMetrics() bool`

HasMissingMetrics returns a boolean if a field has been set.

### GetMetrics

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetMetrics() []ExternalExperimentDtoHealthChecksInnerMetadataMetricsInner`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetMetricsOk() (*[]ExternalExperimentDtoHealthChecksInnerMetadataMetricsInner, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetMetrics(v []ExternalExperimentDtoHealthChecksInnerMetadataMetricsInner)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetLastUpdatedDs

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetLastUpdatedDs() string`

GetLastUpdatedDs returns the LastUpdatedDs field if non-nil, zero value otherwise.

### GetLastUpdatedDsOk

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) GetLastUpdatedDsOk() (*string, bool)`

GetLastUpdatedDsOk returns a tuple with the LastUpdatedDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDs

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) SetLastUpdatedDs(v string)`

SetLastUpdatedDs sets LastUpdatedDs field to given value.

### HasLastUpdatedDs

`func (o *ExternalExperimentDtoHealthChecksInnerMetadata) HasLastUpdatedDs() bool`

HasLastUpdatedDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


