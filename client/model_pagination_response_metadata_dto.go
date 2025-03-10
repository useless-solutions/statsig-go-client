/*
Console API

       The \"Console API\" is the CRUD API for performing the actions offered on console.statsig.com without needing to go through the web UI.       If you have any feature requests, drop on in to our [slack channel](https://www.statsig.com/slack) and let us know.       <br /><br />       <b>Authorization</b>       <br />       All requests must include the **STATSIG-API-KEY** field in the header. The value should be a **Console API Key** which can be created in the Project Settings on [console.statsig.com/api_keys](https://console.statsig.com/api_keys)       <br /><br />       <b>Rate Limiting</b>       <br />       Requests to the Console API are limited to <code>~ 100reqs / 10secs and ~ 900reqs / 15 mins</code>.       <br /><br />       <b>Keyboard Search</b>       <br />       Use <code>Ctrl/Cmd + K</code> to search for specific endpoints.       

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PaginationResponseMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationResponseMetadataDto{}

// PaginationResponseMetadataDto struct for PaginationResponseMetadataDto
type PaginationResponseMetadataDto struct {
	ItemsPerPage float32 `json:"itemsPerPage"`
	PageNumber float32 `json:"pageNumber"`
	NextPage nil `json:"nextPage"`
	PreviousPage nil `json:"previousPage"`
	TotalItems *float32 `json:"totalItems,omitempty"`
	All *string `json:"all,omitempty"`
}

type _PaginationResponseMetadataDto PaginationResponseMetadataDto

// NewPaginationResponseMetadataDto instantiates a new PaginationResponseMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationResponseMetadataDto(itemsPerPage float32, pageNumber float32, nextPage nil, previousPage nil) *PaginationResponseMetadataDto {
	this := PaginationResponseMetadataDto{}
	this.ItemsPerPage = itemsPerPage
	this.PageNumber = pageNumber
	this.NextPage = nextPage
	this.PreviousPage = previousPage
	return &this
}

// NewPaginationResponseMetadataDtoWithDefaults instantiates a new PaginationResponseMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationResponseMetadataDtoWithDefaults() *PaginationResponseMetadataDto {
	this := PaginationResponseMetadataDto{}
	return &this
}

// GetItemsPerPage returns the ItemsPerPage field value
func (o *PaginationResponseMetadataDto) GetItemsPerPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value
// and a boolean to check if the value has been set.
func (o *PaginationResponseMetadataDto) GetItemsPerPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsPerPage, true
}

// SetItemsPerPage sets field value
func (o *PaginationResponseMetadataDto) SetItemsPerPage(v float32) {
	o.ItemsPerPage = v
}

// GetPageNumber returns the PageNumber field value
func (o *PaginationResponseMetadataDto) GetPageNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *PaginationResponseMetadataDto) GetPageNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *PaginationResponseMetadataDto) SetPageNumber(v float32) {
	o.PageNumber = v
}

// GetNextPage returns the NextPage field value
func (o *PaginationResponseMetadataDto) GetNextPage() nil {
	if o == nil {
		var ret nil
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *PaginationResponseMetadataDto) GetNextPageOk() (*nil, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *PaginationResponseMetadataDto) SetNextPage(v nil) {
	o.NextPage = v
}

// GetPreviousPage returns the PreviousPage field value
func (o *PaginationResponseMetadataDto) GetPreviousPage() nil {
	if o == nil {
		var ret nil
		return ret
	}

	return o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value
// and a boolean to check if the value has been set.
func (o *PaginationResponseMetadataDto) GetPreviousPageOk() (*nil, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousPage, true
}

// SetPreviousPage sets field value
func (o *PaginationResponseMetadataDto) SetPreviousPage(v nil) {
	o.PreviousPage = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PaginationResponseMetadataDto) GetTotalItems() float32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret float32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationResponseMetadataDto) GetTotalItemsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PaginationResponseMetadataDto) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given float32 and assigns it to the TotalItems field.
func (o *PaginationResponseMetadataDto) SetTotalItems(v float32) {
	o.TotalItems = &v
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *PaginationResponseMetadataDto) GetAll() string {
	if o == nil || IsNil(o.All) {
		var ret string
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationResponseMetadataDto) GetAllOk() (*string, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *PaginationResponseMetadataDto) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given string and assigns it to the All field.
func (o *PaginationResponseMetadataDto) SetAll(v string) {
	o.All = &v
}

func (o PaginationResponseMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationResponseMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["itemsPerPage"] = o.ItemsPerPage
	toSerialize["pageNumber"] = o.PageNumber
	toSerialize["nextPage"] = o.NextPage
	toSerialize["previousPage"] = o.PreviousPage
	if !IsNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	return toSerialize, nil
}

func (o *PaginationResponseMetadataDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"itemsPerPage",
		"pageNumber",
		"nextPage",
		"previousPage",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaginationResponseMetadataDto := _PaginationResponseMetadataDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginationResponseMetadataDto)

	if err != nil {
		return err
	}

	*o = PaginationResponseMetadataDto(varPaginationResponseMetadataDto)

	return err
}

type NullablePaginationResponseMetadataDto struct {
	value *PaginationResponseMetadataDto
	isSet bool
}

func (v NullablePaginationResponseMetadataDto) Get() *PaginationResponseMetadataDto {
	return v.value
}

func (v *NullablePaginationResponseMetadataDto) Set(val *PaginationResponseMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationResponseMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationResponseMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationResponseMetadataDto(val *PaginationResponseMetadataDto) *NullablePaginationResponseMetadataDto {
	return &NullablePaginationResponseMetadataDto{value: val, isSet: true}
}

func (v NullablePaginationResponseMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationResponseMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


