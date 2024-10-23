# ConsoleV1SegmentsIDListControllerRead200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A simple string explaining the result of the operation. | 
**Data** | [**[]IDListDto**](IDListDto.md) |  | 
**Pagination** | [**PaginationResponseMetadataDto**](PaginationResponseMetadataDto.md) | Pagination metadata for checking if there is next page for example. | 

## Methods

### NewConsoleV1SegmentsIDListControllerRead200Response

`func NewConsoleV1SegmentsIDListControllerRead200Response(message string, data []IDListDto, pagination PaginationResponseMetadataDto, ) *ConsoleV1SegmentsIDListControllerRead200Response`

NewConsoleV1SegmentsIDListControllerRead200Response instantiates a new ConsoleV1SegmentsIDListControllerRead200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleV1SegmentsIDListControllerRead200ResponseWithDefaults

`func NewConsoleV1SegmentsIDListControllerRead200ResponseWithDefaults() *ConsoleV1SegmentsIDListControllerRead200Response`

NewConsoleV1SegmentsIDListControllerRead200ResponseWithDefaults instantiates a new ConsoleV1SegmentsIDListControllerRead200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) GetData() []IDListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) GetDataOk() (*[]IDListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) SetData(v []IDListDto)`

SetData sets Data field to given value.


### GetPagination

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) GetPagination() PaginationResponseMetadataDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) GetPaginationOk() (*PaginationResponseMetadataDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ConsoleV1SegmentsIDListControllerRead200Response) SetPagination(v PaginationResponseMetadataDto)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


