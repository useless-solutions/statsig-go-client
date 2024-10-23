# PaginationResponseWithMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A simple string explaining the result of the operation. | 
**Data** | **[]map[string]interface{}** | Array of results returned by pagination limit. | 
**Pagination** | [**PaginationResponseMetadataDto**](PaginationResponseMetadataDto.md) | Pagination metadata for checking if there is next page for example. | 

## Methods

### NewPaginationResponseWithMessage

`func NewPaginationResponseWithMessage(message string, data []map[string]interface{}, pagination PaginationResponseMetadataDto, ) *PaginationResponseWithMessage`

NewPaginationResponseWithMessage instantiates a new PaginationResponseWithMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResponseWithMessageWithDefaults

`func NewPaginationResponseWithMessageWithDefaults() *PaginationResponseWithMessage`

NewPaginationResponseWithMessageWithDefaults instantiates a new PaginationResponseWithMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PaginationResponseWithMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaginationResponseWithMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaginationResponseWithMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *PaginationResponseWithMessage) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginationResponseWithMessage) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginationResponseWithMessage) SetData(v []map[string]interface{})`

SetData sets Data field to given value.


### GetPagination

`func (o *PaginationResponseWithMessage) GetPagination() PaginationResponseMetadataDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginationResponseWithMessage) GetPaginationOk() (*PaginationResponseMetadataDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginationResponseWithMessage) SetPagination(v PaginationResponseMetadataDto)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


