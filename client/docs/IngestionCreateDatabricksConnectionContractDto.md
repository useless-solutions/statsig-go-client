# IngestionCreateDatabricksConnectionContractDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Host** | **string** |  | 
**Path** | **string** |  | 
**DeltaSharingCredentials** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestionCreateDatabricksConnectionContractDto

`func NewIngestionCreateDatabricksConnectionContractDto(token string, host string, path string, ) *IngestionCreateDatabricksConnectionContractDto`

NewIngestionCreateDatabricksConnectionContractDto instantiates a new IngestionCreateDatabricksConnectionContractDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionCreateDatabricksConnectionContractDtoWithDefaults

`func NewIngestionCreateDatabricksConnectionContractDtoWithDefaults() *IngestionCreateDatabricksConnectionContractDto`

NewIngestionCreateDatabricksConnectionContractDtoWithDefaults instantiates a new IngestionCreateDatabricksConnectionContractDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *IngestionCreateDatabricksConnectionContractDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IngestionCreateDatabricksConnectionContractDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IngestionCreateDatabricksConnectionContractDto) SetToken(v string)`

SetToken sets Token field to given value.


### GetHost

`func (o *IngestionCreateDatabricksConnectionContractDto) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IngestionCreateDatabricksConnectionContractDto) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IngestionCreateDatabricksConnectionContractDto) SetHost(v string)`

SetHost sets Host field to given value.


### GetPath

`func (o *IngestionCreateDatabricksConnectionContractDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IngestionCreateDatabricksConnectionContractDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IngestionCreateDatabricksConnectionContractDto) SetPath(v string)`

SetPath sets Path field to given value.


### GetDeltaSharingCredentials

`func (o *IngestionCreateDatabricksConnectionContractDto) GetDeltaSharingCredentials() string`

GetDeltaSharingCredentials returns the DeltaSharingCredentials field if non-nil, zero value otherwise.

### GetDeltaSharingCredentialsOk

`func (o *IngestionCreateDatabricksConnectionContractDto) GetDeltaSharingCredentialsOk() (*string, bool)`

GetDeltaSharingCredentialsOk returns a tuple with the DeltaSharingCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaSharingCredentials

`func (o *IngestionCreateDatabricksConnectionContractDto) SetDeltaSharingCredentials(v string)`

SetDeltaSharingCredentials sets DeltaSharingCredentials field to given value.

### HasDeltaSharingCredentials

`func (o *IngestionCreateDatabricksConnectionContractDto) HasDeltaSharingCredentials() bool`

HasDeltaSharingCredentials returns a boolean if a field has been set.

### GetVerified

`func (o *IngestionCreateDatabricksConnectionContractDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *IngestionCreateDatabricksConnectionContractDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *IngestionCreateDatabricksConnectionContractDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *IngestionCreateDatabricksConnectionContractDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


