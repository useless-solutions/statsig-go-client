# ConsoleV1ExperimentsControllerGenList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A simple string explaining the result of the operation. | 
**Data** | [**[]ExternalExperimentDto**](ExternalExperimentDto.md) |  | 
**Pagination** | [**PaginationResponseMetadataDto**](PaginationResponseMetadataDto.md) | Pagination metadata for checking if there is next page for example. | 

## Methods

### NewConsoleV1ExperimentsControllerGenList200Response

`func NewConsoleV1ExperimentsControllerGenList200Response(message string, data []ExternalExperimentDto, pagination PaginationResponseMetadataDto, ) *ConsoleV1ExperimentsControllerGenList200Response`

NewConsoleV1ExperimentsControllerGenList200Response instantiates a new ConsoleV1ExperimentsControllerGenList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleV1ExperimentsControllerGenList200ResponseWithDefaults

`func NewConsoleV1ExperimentsControllerGenList200ResponseWithDefaults() *ConsoleV1ExperimentsControllerGenList200Response`

NewConsoleV1ExperimentsControllerGenList200ResponseWithDefaults instantiates a new ConsoleV1ExperimentsControllerGenList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConsoleV1ExperimentsControllerGenList200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConsoleV1ExperimentsControllerGenList200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConsoleV1ExperimentsControllerGenList200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ConsoleV1ExperimentsControllerGenList200Response) GetData() []ExternalExperimentDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleV1ExperimentsControllerGenList200Response) GetDataOk() (*[]ExternalExperimentDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleV1ExperimentsControllerGenList200Response) SetData(v []ExternalExperimentDto)`

SetData sets Data field to given value.


### GetPagination

`func (o *ConsoleV1ExperimentsControllerGenList200Response) GetPagination() PaginationResponseMetadataDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ConsoleV1ExperimentsControllerGenList200Response) GetPaginationOk() (*PaginationResponseMetadataDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ConsoleV1ExperimentsControllerGenList200Response) SetPagination(v PaginationResponseMetadataDto)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


