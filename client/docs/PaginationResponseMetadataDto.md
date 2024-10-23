# PaginationResponseMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsPerPage** | **float32** |  | 
**PageNumber** | **float32** |  | 
**NextPage** | [**nil**](nil.md) |  | 
**PreviousPage** | [**nil**](nil.md) |  | 
**TotalItems** | Pointer to **float32** |  | [optional] 
**All** | Pointer to **string** |  | [optional] 

## Methods

### NewPaginationResponseMetadataDto

`func NewPaginationResponseMetadataDto(itemsPerPage float32, pageNumber float32, nextPage nil, previousPage nil, ) *PaginationResponseMetadataDto`

NewPaginationResponseMetadataDto instantiates a new PaginationResponseMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResponseMetadataDtoWithDefaults

`func NewPaginationResponseMetadataDtoWithDefaults() *PaginationResponseMetadataDto`

NewPaginationResponseMetadataDtoWithDefaults instantiates a new PaginationResponseMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsPerPage

`func (o *PaginationResponseMetadataDto) GetItemsPerPage() float32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *PaginationResponseMetadataDto) GetItemsPerPageOk() (*float32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *PaginationResponseMetadataDto) SetItemsPerPage(v float32)`

SetItemsPerPage sets ItemsPerPage field to given value.


### GetPageNumber

`func (o *PaginationResponseMetadataDto) GetPageNumber() float32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *PaginationResponseMetadataDto) GetPageNumberOk() (*float32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *PaginationResponseMetadataDto) SetPageNumber(v float32)`

SetPageNumber sets PageNumber field to given value.


### GetNextPage

`func (o *PaginationResponseMetadataDto) GetNextPage() nil`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PaginationResponseMetadataDto) GetNextPageOk() (*nil, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PaginationResponseMetadataDto) SetNextPage(v nil)`

SetNextPage sets NextPage field to given value.


### GetPreviousPage

`func (o *PaginationResponseMetadataDto) GetPreviousPage() nil`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *PaginationResponseMetadataDto) GetPreviousPageOk() (*nil, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *PaginationResponseMetadataDto) SetPreviousPage(v nil)`

SetPreviousPage sets PreviousPage field to given value.


### GetTotalItems

`func (o *PaginationResponseMetadataDto) GetTotalItems() float32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PaginationResponseMetadataDto) GetTotalItemsOk() (*float32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PaginationResponseMetadataDto) SetTotalItems(v float32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *PaginationResponseMetadataDto) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetAll

`func (o *PaginationResponseMetadataDto) GetAll() string`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *PaginationResponseMetadataDto) GetAllOk() (*string, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *PaginationResponseMetadataDto) SetAll(v string)`

SetAll sets All field to given value.

### HasAll

`func (o *PaginationResponseMetadataDto) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


