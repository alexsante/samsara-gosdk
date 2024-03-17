/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package samsarago

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListVehiclesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVehiclesResponse{}

// ListVehiclesResponse Response for retrieving a list of vehicles.
type ListVehiclesResponse struct {
	Data []Vehicle `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type _ListVehiclesResponse ListVehiclesResponse

// NewListVehiclesResponse instantiates a new ListVehiclesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVehiclesResponse(data []Vehicle, pagination PaginationResponse) *ListVehiclesResponse {
	this := ListVehiclesResponse{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewListVehiclesResponseWithDefaults instantiates a new ListVehiclesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVehiclesResponseWithDefaults() *ListVehiclesResponse {
	this := ListVehiclesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListVehiclesResponse) GetData() []Vehicle {
	if o == nil {
		var ret []Vehicle
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListVehiclesResponse) GetDataOk() ([]Vehicle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListVehiclesResponse) SetData(v []Vehicle) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *ListVehiclesResponse) GetPagination() PaginationResponse {
	if o == nil {
		var ret PaginationResponse
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListVehiclesResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListVehiclesResponse) SetPagination(v PaginationResponse) {
	o.Pagination = v
}

func (o ListVehiclesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVehiclesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

func (o *ListVehiclesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"pagination",
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

	varListVehiclesResponse := _ListVehiclesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListVehiclesResponse)

	if err != nil {
		return err
	}

	*o = ListVehiclesResponse(varListVehiclesResponse)

	return err
}

type NullableListVehiclesResponse struct {
	value *ListVehiclesResponse
	isSet bool
}

func (v NullableListVehiclesResponse) Get() *ListVehiclesResponse {
	return v.value
}

func (v *NullableListVehiclesResponse) Set(val *ListVehiclesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVehiclesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVehiclesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVehiclesResponse(val *ListVehiclesResponse) *NullableListVehiclesResponse {
	return &NullableListVehiclesResponse{value: val, isSet: true}
}

func (v NullableListVehiclesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVehiclesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


