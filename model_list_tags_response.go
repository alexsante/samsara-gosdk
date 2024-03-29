/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package samsarago

import (
	"encoding/json"
)

// checks if the ListTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTagsResponse{}

// ListTagsResponse A list of tags.
type ListTagsResponse struct {
	Data []Tag `json:"data,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewListTagsResponse instantiates a new ListTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTagsResponse() *ListTagsResponse {
	this := ListTagsResponse{}
	return &this
}

// NewListTagsResponseWithDefaults instantiates a new ListTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTagsResponseWithDefaults() *ListTagsResponse {
	this := ListTagsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListTagsResponse) GetData() []Tag {
	if o == nil || IsNil(o.Data) {
		var ret []Tag
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTagsResponse) GetDataOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListTagsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Tag and assigns it to the Data field.
func (o *ListTagsResponse) SetData(v []Tag) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListTagsResponse) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTagsResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListTagsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *ListTagsResponse) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o ListTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListTagsResponse struct {
	value *ListTagsResponse
	isSet bool
}

func (v NullableListTagsResponse) Get() *ListTagsResponse {
	return v.value
}

func (v *NullableListTagsResponse) Set(val *ListTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTagsResponse(val *ListTagsResponse) *NullableListTagsResponse {
	return &NullableListTagsResponse{value: val, isSet: true}
}

func (v NullableListTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


