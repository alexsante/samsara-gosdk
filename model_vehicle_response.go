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

// checks if the VehicleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleResponse{}

// VehicleResponse A single vehicle.
type VehicleResponse struct {
	Data Vehicle `json:"data"`
}

type _VehicleResponse VehicleResponse

// NewVehicleResponse instantiates a new VehicleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleResponse(data Vehicle) *VehicleResponse {
	this := VehicleResponse{}
	this.Data = data
	return &this
}

// NewVehicleResponseWithDefaults instantiates a new VehicleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleResponseWithDefaults() *VehicleResponse {
	this := VehicleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *VehicleResponse) GetData() Vehicle {
	if o == nil {
		var ret Vehicle
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VehicleResponse) GetDataOk() (*Vehicle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *VehicleResponse) SetData(v Vehicle) {
	o.Data = v
}

func (o VehicleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *VehicleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varVehicleResponse := _VehicleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVehicleResponse)

	if err != nil {
		return err
	}

	*o = VehicleResponse(varVehicleResponse)

	return err
}

type NullableVehicleResponse struct {
	value *VehicleResponse
	isSet bool
}

func (v NullableVehicleResponse) Get() *VehicleResponse {
	return v.value
}

func (v *NullableVehicleResponse) Set(val *VehicleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleResponse(val *VehicleResponse) *NullableVehicleResponse {
	return &NullableVehicleResponse{value: val, isSet: true}
}

func (v NullableVehicleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


