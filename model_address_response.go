/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressResponse{}

// AddressResponse An Address response body.
type AddressResponse struct {
	Data Address `json:"data"`
}

type _AddressResponse AddressResponse

// NewAddressResponse instantiates a new AddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressResponse(data Address) *AddressResponse {
	this := AddressResponse{}
	this.Data = data
	return &this
}

// NewAddressResponseWithDefaults instantiates a new AddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressResponseWithDefaults() *AddressResponse {
	this := AddressResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AddressResponse) GetData() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetDataOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressResponse) SetData(v Address) {
	o.Data = v
}

func (o AddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AddressResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varAddressResponse := _AddressResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressResponse)

	if err != nil {
		return err
	}

	*o = AddressResponse(varAddressResponse)

	return err
}

type NullableAddressResponse struct {
	value *AddressResponse
	isSet bool
}

func (v NullableAddressResponse) Get() *AddressResponse {
	return v.value
}

func (v *NullableAddressResponse) Set(val *AddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressResponse(val *AddressResponse) *NullableAddressResponse {
	return &NullableAddressResponse{value: val, isSet: true}
}

func (v NullableAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

