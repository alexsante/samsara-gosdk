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

// checks if the VehicleStatsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleStatsResponse{}

// VehicleStatsResponse Most recent vehicle stats and pagination info.
type VehicleStatsResponse struct {
	// List of the most recent stats for the specified vehicles and stat types.
	Data []VehicleStatsResponseDataInner `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type _VehicleStatsResponse VehicleStatsResponse

// NewVehicleStatsResponse instantiates a new VehicleStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleStatsResponse(data []VehicleStatsResponseDataInner, pagination PaginationResponse) *VehicleStatsResponse {
	this := VehicleStatsResponse{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewVehicleStatsResponseWithDefaults instantiates a new VehicleStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleStatsResponseWithDefaults() *VehicleStatsResponse {
	this := VehicleStatsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *VehicleStatsResponse) GetData() []VehicleStatsResponseDataInner {
	if o == nil {
		var ret []VehicleStatsResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponse) GetDataOk() ([]VehicleStatsResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *VehicleStatsResponse) SetData(v []VehicleStatsResponseDataInner) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *VehicleStatsResponse) GetPagination() PaginationResponse {
	if o == nil {
		var ret PaginationResponse
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *VehicleStatsResponse) SetPagination(v PaginationResponse) {
	o.Pagination = v
}

func (o VehicleStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleStatsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

func (o *VehicleStatsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varVehicleStatsResponse := _VehicleStatsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVehicleStatsResponse)

	if err != nil {
		return err
	}

	*o = VehicleStatsResponse(varVehicleStatsResponse)

	return err
}

type NullableVehicleStatsResponse struct {
	value *VehicleStatsResponse
	isSet bool
}

func (v NullableVehicleStatsResponse) Get() *VehicleStatsResponse {
	return v.value
}

func (v *NullableVehicleStatsResponse) Set(val *VehicleStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleStatsResponse(val *VehicleStatsResponse) *NullableVehicleStatsResponse {
	return &NullableVehicleStatsResponse{value: val, isSet: true}
}

func (v NullableVehicleStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


