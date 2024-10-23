# IngestionSourceCreateContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**ColumnMapping** | Pointer to [**IngestionSourceCreateContractDtoOneOf2ColumnMapping**](IngestionSourceCreateContractDtoOneOf2ColumnMapping.md) |  | [optional] 
**Type** | **string** |  | 
**SourceName** | **string** |  | 
**Query** | Pointer to **string** |  | [optional] 
**UseDeltaSharing** | Pointer to **bool** |  | [optional] 
**Share** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestionSourceCreateContractDto

`func NewIngestionSourceCreateContractDto(dataset string, type_ string, sourceName string, ) *IngestionSourceCreateContractDto`

NewIngestionSourceCreateContractDto instantiates a new IngestionSourceCreateContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSourceCreateContractDtoWithDefaults

`func NewIngestionSourceCreateContractDtoWithDefaults() *IngestionSourceCreateContractDto`

NewIngestionSourceCreateContractDtoWithDefaults instantiates a new IngestionSourceCreateContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *IngestionSourceCreateContractDto) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionSourceCreateContractDto) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionSourceCreateContractDto) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetColumnMapping

`func (o *IngestionSourceCreateContractDto) GetColumnMapping() IngestionSourceCreateContractDtoOneOf2ColumnMapping`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *IngestionSourceCreateContractDto) GetColumnMappingOk() (*IngestionSourceCreateContractDtoOneOf2ColumnMapping, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *IngestionSourceCreateContractDto) SetColumnMapping(v IngestionSourceCreateContractDtoOneOf2ColumnMapping)`

SetColumnMapping sets ColumnMapping field to given value.

### HasColumnMapping

`func (o *IngestionSourceCreateContractDto) HasColumnMapping() bool`

HasColumnMapping returns a boolean if a field has been set.

### GetType

`func (o *IngestionSourceCreateContractDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionSourceCreateContractDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionSourceCreateContractDto) SetType(v string)`

SetType sets Type field to given value.


### GetSourceName

`func (o *IngestionSourceCreateContractDto) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *IngestionSourceCreateContractDto) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *IngestionSourceCreateContractDto) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetQuery

`func (o *IngestionSourceCreateContractDto) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IngestionSourceCreateContractDto) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IngestionSourceCreateContractDto) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IngestionSourceCreateContractDto) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetUseDeltaSharing

`func (o *IngestionSourceCreateContractDto) GetUseDeltaSharing() bool`

GetUseDeltaSharing returns the UseDeltaSharing field if non-nil, zero value otherwise.

### GetUseDeltaSharingOk

`func (o *IngestionSourceCreateContractDto) GetUseDeltaSharingOk() (*bool, bool)`

GetUseDeltaSharingOk returns a tuple with the UseDeltaSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDeltaSharing

`func (o *IngestionSourceCreateContractDto) SetUseDeltaSharing(v bool)`

SetUseDeltaSharing sets UseDeltaSharing field to given value.

### HasUseDeltaSharing

`func (o *IngestionSourceCreateContractDto) HasUseDeltaSharing() bool`

HasUseDeltaSharing returns a boolean if a field has been set.

### GetShare

`func (o *IngestionSourceCreateContractDto) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *IngestionSourceCreateContractDto) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *IngestionSourceCreateContractDto) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *IngestionSourceCreateContractDto) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSchema

`func (o *IngestionSourceCreateContractDto) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IngestionSourceCreateContractDto) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IngestionSourceCreateContractDto) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IngestionSourceCreateContractDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *IngestionSourceCreateContractDto) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *IngestionSourceCreateContractDto) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *IngestionSourceCreateContractDto) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *IngestionSourceCreateContractDto) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetEnabled

`func (o *IngestionSourceCreateContractDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IngestionSourceCreateContractDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IngestionSourceCreateContractDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IngestionSourceCreateContractDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


