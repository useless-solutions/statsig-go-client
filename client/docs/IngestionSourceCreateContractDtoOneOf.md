# IngestionSourceCreateContractDtoOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**ColumnMapping** | Pointer to [**IngestionSourceCreateContractDtoOneOfColumnMapping**](IngestionSourceCreateContractDtoOneOfColumnMapping.md) |  | [optional] 
**Type** | **string** |  | 
**SourceName** | **string** |  | 
**Query** | Pointer to **string** |  | [optional] 
**UseDeltaSharing** | Pointer to **bool** |  | [optional] 
**Share** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestionSourceCreateContractDtoOneOf

`func NewIngestionSourceCreateContractDtoOneOf(dataset string, type_ string, sourceName string, ) *IngestionSourceCreateContractDtoOneOf`

NewIngestionSourceCreateContractDtoOneOf instantiates a new IngestionSourceCreateContractDtoOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSourceCreateContractDtoOneOfWithDefaults

`func NewIngestionSourceCreateContractDtoOneOfWithDefaults() *IngestionSourceCreateContractDtoOneOf`

NewIngestionSourceCreateContractDtoOneOfWithDefaults instantiates a new IngestionSourceCreateContractDtoOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *IngestionSourceCreateContractDtoOneOf) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionSourceCreateContractDtoOneOf) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetColumnMapping

`func (o *IngestionSourceCreateContractDtoOneOf) GetColumnMapping() IngestionSourceCreateContractDtoOneOfColumnMapping`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetColumnMappingOk() (*IngestionSourceCreateContractDtoOneOfColumnMapping, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *IngestionSourceCreateContractDtoOneOf) SetColumnMapping(v IngestionSourceCreateContractDtoOneOfColumnMapping)`

SetColumnMapping sets ColumnMapping field to given value.

### HasColumnMapping

`func (o *IngestionSourceCreateContractDtoOneOf) HasColumnMapping() bool`

HasColumnMapping returns a boolean if a field has been set.

### GetType

`func (o *IngestionSourceCreateContractDtoOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionSourceCreateContractDtoOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetSourceName

`func (o *IngestionSourceCreateContractDtoOneOf) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *IngestionSourceCreateContractDtoOneOf) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetQuery

`func (o *IngestionSourceCreateContractDtoOneOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IngestionSourceCreateContractDtoOneOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IngestionSourceCreateContractDtoOneOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetUseDeltaSharing

`func (o *IngestionSourceCreateContractDtoOneOf) GetUseDeltaSharing() bool`

GetUseDeltaSharing returns the UseDeltaSharing field if non-nil, zero value otherwise.

### GetUseDeltaSharingOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetUseDeltaSharingOk() (*bool, bool)`

GetUseDeltaSharingOk returns a tuple with the UseDeltaSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDeltaSharing

`func (o *IngestionSourceCreateContractDtoOneOf) SetUseDeltaSharing(v bool)`

SetUseDeltaSharing sets UseDeltaSharing field to given value.

### HasUseDeltaSharing

`func (o *IngestionSourceCreateContractDtoOneOf) HasUseDeltaSharing() bool`

HasUseDeltaSharing returns a boolean if a field has been set.

### GetShare

`func (o *IngestionSourceCreateContractDtoOneOf) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *IngestionSourceCreateContractDtoOneOf) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *IngestionSourceCreateContractDtoOneOf) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSchema

`func (o *IngestionSourceCreateContractDtoOneOf) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IngestionSourceCreateContractDtoOneOf) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IngestionSourceCreateContractDtoOneOf) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *IngestionSourceCreateContractDtoOneOf) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *IngestionSourceCreateContractDtoOneOf) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *IngestionSourceCreateContractDtoOneOf) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetEnabled

`func (o *IngestionSourceCreateContractDtoOneOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IngestionSourceCreateContractDtoOneOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IngestionSourceCreateContractDtoOneOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IngestionSourceCreateContractDtoOneOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


