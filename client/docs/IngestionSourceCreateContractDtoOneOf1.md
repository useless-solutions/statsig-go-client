# IngestionSourceCreateContractDtoOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** |  | 
**ColumnMapping** | Pointer to [**IngestionSourceCreateContractDtoOneOf1ColumnMapping**](IngestionSourceCreateContractDtoOneOf1ColumnMapping.md) |  | [optional] 
**Type** | **string** |  | 
**SourceName** | **string** |  | 
**Query** | Pointer to **string** |  | [optional] 
**UseDeltaSharing** | Pointer to **bool** |  | [optional] 
**Share** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestionSourceCreateContractDtoOneOf1

`func NewIngestionSourceCreateContractDtoOneOf1(dataset string, type_ string, sourceName string, ) *IngestionSourceCreateContractDtoOneOf1`

NewIngestionSourceCreateContractDtoOneOf1 instantiates a new IngestionSourceCreateContractDtoOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSourceCreateContractDtoOneOf1WithDefaults

`func NewIngestionSourceCreateContractDtoOneOf1WithDefaults() *IngestionSourceCreateContractDtoOneOf1`

NewIngestionSourceCreateContractDtoOneOf1WithDefaults instantiates a new IngestionSourceCreateContractDtoOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *IngestionSourceCreateContractDtoOneOf1) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *IngestionSourceCreateContractDtoOneOf1) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetColumnMapping

`func (o *IngestionSourceCreateContractDtoOneOf1) GetColumnMapping() IngestionSourceCreateContractDtoOneOf1ColumnMapping`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetColumnMappingOk() (*IngestionSourceCreateContractDtoOneOf1ColumnMapping, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *IngestionSourceCreateContractDtoOneOf1) SetColumnMapping(v IngestionSourceCreateContractDtoOneOf1ColumnMapping)`

SetColumnMapping sets ColumnMapping field to given value.

### HasColumnMapping

`func (o *IngestionSourceCreateContractDtoOneOf1) HasColumnMapping() bool`

HasColumnMapping returns a boolean if a field has been set.

### GetType

`func (o *IngestionSourceCreateContractDtoOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionSourceCreateContractDtoOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetSourceName

`func (o *IngestionSourceCreateContractDtoOneOf1) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *IngestionSourceCreateContractDtoOneOf1) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetQuery

`func (o *IngestionSourceCreateContractDtoOneOf1) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IngestionSourceCreateContractDtoOneOf1) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IngestionSourceCreateContractDtoOneOf1) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetUseDeltaSharing

`func (o *IngestionSourceCreateContractDtoOneOf1) GetUseDeltaSharing() bool`

GetUseDeltaSharing returns the UseDeltaSharing field if non-nil, zero value otherwise.

### GetUseDeltaSharingOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetUseDeltaSharingOk() (*bool, bool)`

GetUseDeltaSharingOk returns a tuple with the UseDeltaSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDeltaSharing

`func (o *IngestionSourceCreateContractDtoOneOf1) SetUseDeltaSharing(v bool)`

SetUseDeltaSharing sets UseDeltaSharing field to given value.

### HasUseDeltaSharing

`func (o *IngestionSourceCreateContractDtoOneOf1) HasUseDeltaSharing() bool`

HasUseDeltaSharing returns a boolean if a field has been set.

### GetShare

`func (o *IngestionSourceCreateContractDtoOneOf1) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *IngestionSourceCreateContractDtoOneOf1) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *IngestionSourceCreateContractDtoOneOf1) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSchema

`func (o *IngestionSourceCreateContractDtoOneOf1) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IngestionSourceCreateContractDtoOneOf1) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *IngestionSourceCreateContractDtoOneOf1) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *IngestionSourceCreateContractDtoOneOf1) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *IngestionSourceCreateContractDtoOneOf1) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *IngestionSourceCreateContractDtoOneOf1) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetEnabled

`func (o *IngestionSourceCreateContractDtoOneOf1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IngestionSourceCreateContractDtoOneOf1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IngestionSourceCreateContractDtoOneOf1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IngestionSourceCreateContractDtoOneOf1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


