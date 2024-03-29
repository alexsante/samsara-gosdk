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

// checks if the TinyTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TinyTag{}

// TinyTag struct for TinyTag
type TinyTag struct {
	// Unique Samsara ID of this tag.
	Id *string `json:"id,omitempty"`
	// Name of this tag.
	Name *string `json:"name,omitempty"`
	// If this tag is part a hierarchical tag tree, this is the ID of the parent tag, otherwise this will be omitted.
	ParentTagId *string `json:"parentTagId,omitempty"`
}

// NewTinyTag instantiates a new TinyTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTinyTag() *TinyTag {
	this := TinyTag{}
	return &this
}

// NewTinyTagWithDefaults instantiates a new TinyTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTinyTagWithDefaults() *TinyTag {
	this := TinyTag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TinyTag) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TinyTag) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TinyTag) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TinyTag) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TinyTag) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TinyTag) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TinyTag) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TinyTag) SetName(v string) {
	o.Name = &v
}

// GetParentTagId returns the ParentTagId field value if set, zero value otherwise.
func (o *TinyTag) GetParentTagId() string {
	if o == nil || IsNil(o.ParentTagId) {
		var ret string
		return ret
	}
	return *o.ParentTagId
}

// GetParentTagIdOk returns a tuple with the ParentTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TinyTag) GetParentTagIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentTagId) {
		return nil, false
	}
	return o.ParentTagId, true
}

// HasParentTagId returns a boolean if a field has been set.
func (o *TinyTag) HasParentTagId() bool {
	if o != nil && !IsNil(o.ParentTagId) {
		return true
	}

	return false
}

// SetParentTagId gets a reference to the given string and assigns it to the ParentTagId field.
func (o *TinyTag) SetParentTagId(v string) {
	o.ParentTagId = &v
}

func (o TinyTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TinyTag) ToMap() (map[string]interface{}, error) {
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

type NullableTinyTag struct {
	value *TinyTag
	isSet bool
}

func (v NullableTinyTag) Get() *TinyTag {
	return v.value
}

func (v *NullableTinyTag) Set(val *TinyTag) {
	v.value = val
	v.isSet = true
}

func (v NullableTinyTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTinyTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTinyTag(val *TinyTag) *NullableTinyTag {
	return &NullableTinyTag{value: val, isSet: true}
}

func (v NullableTinyTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTinyTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


