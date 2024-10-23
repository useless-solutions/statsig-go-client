# IngestionUpdateContractDtoOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**ColumnMapping** | Pointer to [**IngestionSourceCreateContractDtoOneOf2ColumnMapping**](IngestionSourceCreateContractDtoOneOf2ColumnMapping.md) |  | [optional] 
**Type** | **string** |  | 
**SourceName** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Share** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestionUpdateContractDtoOneOf2

`func NewIngestionUpdateContractDtoOneOf2(dataset string, type_ string, ) *IngestionUpdateContractDtoOneOf2`

NewIngestionUpdateContractDtoOneOf2 instantiates a new IngestionUpdateContractDtoOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionUpdateContractDtoOneOf2WithDefaults

`func NewIngestionUpdateContractDtoOneOf2WithDefaults() *IngestionUpdateContractDtoOneOf2`

NewIngestionUpdateContractDtoOneOf2WithDefaults instantiates a new IngestionUpdateContractDtoOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *IngestionUpdateContractDtoOneOf2) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionUpdateContractDtoOneOf2) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionUpdateContractDtoOneOf2) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetColumnMapping

`func (o *IngestionUpdateContractDtoOneOf2) GetColumnMapping() IngestionSourceCreateContractDtoOneOf2ColumnMapping`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *IngestionUpdateContractDtoOneOf2) GetColumnMappingOk() (*IngestionSourceCreateContractDtoOneOf2ColumnMapping, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *IngestionUpdateContractDtoOneOf2) SetColumnMapping(v IngestionSourceCreateContractDtoOneOf2ColumnMapping)`

SetColumnMapping sets ColumnMapping field to given value.

### HasColumnMapping

`func (o *IngestionUpdateContractDtoOneOf2) HasColumnMapping() bool`

HasColumnMapping returns a boolean if a field has been set.

### GetType

`func (o *IngestionUpdateContractDtoOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionUpdateContractDtoOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionUpdateContractDtoOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetSourceName

`func (o *IngestionUpdateContractDtoOneOf2) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *IngestionUpdateContractDtoOneOf2) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *IngestionUpdateContractDtoOneOf2) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *IngestionUpdateContractDtoOneOf2) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetQuery

`func (o *IngestionUpdateContractDtoOneOf2) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IngestionUpdateContractDtoOneOf2) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IngestionUpdateContractDtoOneOf2) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IngestionUpdateContractDtoOneOf2) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetShare

`func (o *IngestionUpdateContractDtoOneOf2) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *IngestionUpdateContractDtoOneOf2) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *IngestionUpdateContractDtoOneOf2) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *IngestionUpdateContractDtoOneOf2) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSchema

`func (o *IngestionUpdateContractDtoOneOf2) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IngestionUpdateContractDtoOneOf2) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IngestionUpdateContractDtoOneOf2) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IngestionUpdateContractDtoOneOf2) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *IngestionUpdateContractDtoOneOf2) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *IngestionUpdateContractDtoOneOf2) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *IngestionUpdateContractDtoOneOf2) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *IngestionUpdateContractDtoOneOf2) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetEnabled

`func (o *IngestionUpdateContractDtoOneOf2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IngestionUpdateContractDtoOneOf2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IngestionUpdateContractDtoOneOf2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IngestionUpdateContractDtoOneOf2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


