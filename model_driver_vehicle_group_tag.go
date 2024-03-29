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

// checks if the DriverVehicleGroupTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DriverVehicleGroupTag{}

// DriverVehicleGroupTag struct for DriverVehicleGroupTag
type DriverVehicleGroupTag struct {
	// ID of the tag.
	Id *string `json:"id,omitempty"`
	// Name of the tag.
	Name *string `json:"name,omitempty"`
	// If this tag is part a hierarchical tag tree, this is the ID of the parent tag, otherwise this will be omitted.
	ParentTagId *string `json:"parentTagId,omitempty"`
}

// NewDriverVehicleGroupTag instantiates a new DriverVehicleGroupTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriverVehicleGroupTag() *DriverVehicleGroupTag {
	this := DriverVehicleGroupTag{}
	return &this
}

// NewDriverVehicleGroupTagWithDefaults instantiates a new DriverVehicleGroupTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriverVehicleGroupTagWithDefaults() *DriverVehicleGroupTag {
	this := DriverVehicleGroupTag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DriverVehicleGroupTag) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverVehicleGroupTag) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DriverVehicleGroupTag) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DriverVehicleGroupTag) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DriverVehicleGroupTag) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverVehicleGroupTag) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DriverVehicleGroupTag) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DriverVehicleGroupTag) SetName(v string) {
	o.Name = &v
}

// GetParentTagId returns the ParentTagId field value if set, zero value otherwise.
func (o *DriverVehicleGroupTag) GetParentTagId() string {
	if o == nil || IsNil(o.ParentTagId) {
		var ret string
		return ret
	}
	return *o.ParentTagId
}

// GetParentTagIdOk returns a tuple with the ParentTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverVehicleGroupTag) GetParentTagIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentTagId) {
		return nil, false
	}
	return o.ParentTagId, true
}

// HasParentTagId returns a boolean if a field has been set.
func (o *DriverVehicleGroupTag) HasParentTagId() bool {
	if o != nil && !IsNil(o.ParentTagId) {
		return true
	}

	return false
}

// SetParentTagId gets a reference to the given string and assigns it to the ParentTagId field.
func (o *DriverVehicleGroupTag) SetParentTagId(v string) {
	o.ParentTagId = &v
}

func (o DriverVehicleGroupTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DriverVehicleGroupTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParentTagId) {
		toSerialize["parentTagId"] = o.ParentTagId
	}
	return toSerialize, nil
}

type NullableDriverVehicleGroupTag struct {
	value *DriverVehicleGroupTag
	isSet bool
}

func (v NullableDriverVehicleGroupTag) Get() *DriverVehicleGroupTag {
	return v.value
}

func (v *NullableDriverVehicleGroupTag) Set(val *DriverVehicleGroupTag) {
	v.value = val
	v.isSet = true
}

func (v NullableDriverVehicleGroupTag) IsSet() bool {
	return v.isSet
}

func (v *NullableDriverVehicleGroupTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriverVehicleGroupTag(val *DriverVehicleGroupTag) *NullableDriverVehicleGroupTag {
	return &NullableDriverVehicleGroupTag{value: val, isSet: true}
}

func (v NullableDriverVehicleGroupTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriverVehicleGroupTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


