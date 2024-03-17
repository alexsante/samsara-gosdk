/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VehicleTinyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleTinyResponse{}

// VehicleTinyResponse A minified vehicle object.
type VehicleTinyResponse struct {
	// ID of the vehicle.
	Id *string `json:"id,omitempty"`
	// Name of the vehicle.
	Name *string `json:"name,omitempty"`
}

// NewVehicleTinyResponse instantiates a new VehicleTinyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleTinyResponse() *VehicleTinyResponse {
	this := VehicleTinyResponse{}
	return &this
}

// NewVehicleTinyResponseWithDefaults instantiates a new VehicleTinyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleTinyResponseWithDefaults() *VehicleTinyResponse {
	this := VehicleTinyResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VehicleTinyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTinyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VehicleTinyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VehicleTinyResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VehicleTinyResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTinyResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VehicleTinyResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VehicleTinyResponse) SetName(v string) {
	o.Name = &v
}

func (o VehicleTinyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleTinyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableVehicleTinyResponse struct {
	value *VehicleTinyResponse
	isSet bool
}

func (v NullableVehicleTinyResponse) Get() *VehicleTinyResponse {
	return v.value
}

func (v *NullableVehicleTinyResponse) Set(val *VehicleTinyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleTinyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleTinyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleTinyResponse(val *VehicleTinyResponse) *NullableVehicleTinyResponse {
	return &NullableVehicleTinyResponse{value: val, isSet: true}
}

func (v NullableVehicleTinyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleTinyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


