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

// checks if the VehicleLocationReverseGeo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleLocationReverseGeo{}

// VehicleLocationReverseGeo Reverse geocoded information.
type VehicleLocationReverseGeo struct {
	// Formatted address of the reverse geocoding data.
	FormattedLocation *string `json:"formattedLocation,omitempty"`
}

// NewVehicleLocationReverseGeo instantiates a new VehicleLocationReverseGeo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleLocationReverseGeo() *VehicleLocationReverseGeo {
	this := VehicleLocationReverseGeo{}
	return &this
}

// NewVehicleLocationReverseGeoWithDefaults instantiates a new VehicleLocationReverseGeo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleLocationReverseGeoWithDefaults() *VehicleLocationReverseGeo {
	this := VehicleLocationReverseGeo{}
	return &this
}

// GetFormattedLocation returns the FormattedLocation field value if set, zero value otherwise.
func (o *VehicleLocationReverseGeo) GetFormattedLocation() string {
	if o == nil || IsNil(o.FormattedLocation) {
		var ret string
		return ret
	}
	return *o.FormattedLocation
}

// GetFormattedLocationOk returns a tuple with the FormattedLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleLocationReverseGeo) GetFormattedLocationOk() (*string, bool) {
	if o == nil || IsNil(o.FormattedLocation) {
		return nil, false
	}
	return o.FormattedLocation, true
}

// HasFormattedLocation returns a boolean if a field has been set.
func (o *VehicleLocationReverseGeo) HasFormattedLocation() bool {
	if o != nil && !IsNil(o.FormattedLocation) {
		return true
	}

	return false
}

// SetFormattedLocation gets a reference to the given string and assigns it to the FormattedLocation field.
func (o *VehicleLocationReverseGeo) SetFormattedLocation(v string) {
	o.FormattedLocation = &v
}

func (o VehicleLocationReverseGeo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleLocationReverseGeo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FormattedLocation) {
		toSerialize["formattedLocation"] = o.FormattedLocation
	}
	return toSerialize, nil
}

type NullableVehicleLocationReverseGeo struct {
	value *VehicleLocationReverseGeo
	isSet bool
}

func (v NullableVehicleLocationReverseGeo) Get() *VehicleLocationReverseGeo {
	return v.value
}

func (v *NullableVehicleLocationReverseGeo) Set(val *VehicleLocationReverseGeo) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleLocationReverseGeo) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleLocationReverseGeo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleLocationReverseGeo(val *VehicleLocationReverseGeo) *NullableVehicleLocationReverseGeo {
	return &NullableVehicleLocationReverseGeo{value: val, isSet: true}
}

func (v NullableVehicleLocationReverseGeo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleLocationReverseGeo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


