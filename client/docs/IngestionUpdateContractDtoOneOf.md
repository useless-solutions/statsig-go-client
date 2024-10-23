# IngestionUpdateContractDtoOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**ColumnMapping** | Pointer to [**IngestionSourceCreateContractDtoOneOfColumnMapping**](IngestionSourceCreateContractDtoOneOfColumnMapping.md) |  | [optional] 
**Type** | **string** |  | 
**SourceName** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Share** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestionUpdateContractDtoOneOf

`func NewIngestionUpdateContractDtoOneOf(dataset string, type_ string, ) *IngestionUpdateContractDtoOneOf`

NewIngestionUpdateContractDtoOneOf instantiates a new IngestionUpdateContractDtoOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionUpdateContractDtoOneOfWithDefaults

`func NewIngestionUpdateContractDtoOneOfWithDefaults() *IngestionUpdateContractDtoOneOf`

NewIngestionUpdateContractDtoOneOfWithDefaults instantiates a new IngestionUpdateContractDtoOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *IngestionUpdateContractDtoOneOf) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionUpdateContractDtoOneOf) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionUpdateContractDtoOneOf) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetColumnMapping

`func (o *IngestionUpdateContractDtoOneOf) GetColumnMapping() IngestionSourceCreateContractDtoOneOfColumnMapping`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *IngestionUpdateContractDtoOneOf) GetColumnMappingOk() (*IngestionSourceCreateContractDtoOneOfColumnMapping, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *IngestionUpdateContractDtoOneOf) SetColumnMapping(v IngestionSourceCreateContractDtoOneOfColumnMapping)`

SetColumnMapping sets ColumnMapping field to given value.

### HasColumnMapping

`func (o *IngestionUpdateContractDtoOneOf) HasColumnMapping() bool`

HasColumnMapping returns a boolean if a field has been set.

### GetType

`func (o *IngestionUpdateContractDtoOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionUpdateContractDtoOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionUpdateContractDtoOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetSourceName

`func (o *IngestionUpdateContractDtoOneOf) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *IngestionUpdateContractDtoOneOf) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *IngestionUpdateContractDtoOneOf) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *IngestionUpdateContractDtoOneOf) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetQuery

`func (o *IngestionUpdateContractDtoOneOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IngestionUpdateContractDtoOneOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IngestionUpdateContractDtoOneOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IngestionUpdateContractDtoOneOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetShare

`func (o *IngestionUpdateContractDtoOneOf) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *IngestionUpdateContractDtoOneOf) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *IngestionUpdateContractDtoOneOf) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *IngestionUpdateContractDtoOneOf) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSchema

`func (o *IngestionUpdateContractDtoOneOf) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IngestionUpdateContractDtoOneOf) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IngestionUpdateContractDtoOneOf) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IngestionUpdateContractDtoOneOf) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *IngestionUpdateContractDtoOneOf) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *IngestionUpdateContractDtoOneOf) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *IngestionUpdateContractDtoOneOf) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *IngestionUpdateContractDtoOneOf) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetEnabled

`func (o *IngestionUpdateContractDtoOneOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IngestionUpdateContractDtoOneOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IngestionUpdateContractDtoOneOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IngestionUpdateContractDtoOneOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

