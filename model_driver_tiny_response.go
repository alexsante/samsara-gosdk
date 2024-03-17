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

// checks if the DriverTinyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DriverTinyResponse{}

// DriverTinyResponse A minified driver object.
type DriverTinyResponse struct {
	// ID of the driver.
	Id *string `json:"id,omitempty"`
	// Name of the driver.
	Name *string `json:"name,omitempty"`
}

// NewDriverTinyResponse instantiates a new DriverTinyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriverTinyResponse() *DriverTinyResponse {
	this := DriverTinyResponse{}
	return &this
}

// NewDriverTinyResponseWithDefaults instantiates a new DriverTinyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriverTinyResponseWithDefaults() *DriverTinyResponse {
	this := DriverTinyResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DriverTinyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverTinyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DriverTinyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DriverTinyResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DriverTinyResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverTinyResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DriverTinyResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DriverTinyResponse) SetName(v string) {
	o.Name = &v
}

func (o DriverTinyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DriverTinyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDriverTinyResponse struct {
	value *DriverTinyResponse
	isSet bool
}

func (v NullableDriverTinyResponse) Get() *DriverTinyResponse {
	return v.value
}

func (v *NullableDriverTinyResponse) Set(val *DriverTinyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDriverTinyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDriverTinyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriverTinyResponse(val *DriverTinyResponse) *NullableDriverTinyResponse {
	return &NullableDriverTinyResponse{value: val, isSet: true}
}

func (v NullableDriverTinyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriverTinyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


