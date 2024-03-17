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

// checks if the DriverResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DriverResponse{}

// DriverResponse A single driver.
type DriverResponse struct {
	Data *Driver `json:"data,omitempty"`
}

// NewDriverResponse instantiates a new DriverResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriverResponse() *DriverResponse {
	this := DriverResponse{}
	return &this
}

// NewDriverResponseWithDefaults instantiates a new DriverResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriverResponseWithDefaults() *DriverResponse {
	this := DriverResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DriverResponse) GetData() Driver {
	if o == nil || IsNil(o.Data) {
		var ret Driver
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverResponse) GetDataOk() (*Driver, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DriverResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Driver and assigns it to the Data field.
func (o *DriverResponse) SetData(v Driver) {
	o.Data = &v
}

func (o DriverResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DriverResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDriverResponse struct {
	value *DriverResponse
	isSet bool
}

func (v NullableDriverResponse) Get() *DriverResponse {
	return v.value
}

func (v *NullableDriverResponse) Set(val *DriverResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDriverResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDriverResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriverResponse(val *DriverResponse) *NullableDriverResponse {
	return &NullableDriverResponse{value: val, isSet: true}
}

func (v NullableDriverResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriverResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


