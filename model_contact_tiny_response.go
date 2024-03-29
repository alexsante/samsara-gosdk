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

// checks if the ContactTinyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactTinyResponse{}

// ContactTinyResponse A minified contact object
type ContactTinyResponse struct {
	// First name of the contact.
	FirstName *string `json:"firstName,omitempty"`
	// ID of the contact.
	Id *string `json:"id,omitempty"`
	// Last name of the contact.
	LastName *string `json:"lastName,omitempty"`
}

// NewContactTinyResponse instantiates a new ContactTinyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactTinyResponse() *ContactTinyResponse {
	this := ContactTinyResponse{}
	return &this
}

// NewContactTinyResponseWithDefaults instantiates a new ContactTinyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactTinyResponseWithDefaults() *ContactTinyResponse {
	this := ContactTinyResponse{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ContactTinyResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactTinyResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ContactTinyResponse) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ContactTinyResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContactTinyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactTinyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContactTinyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContactTinyResponse) SetId(v string) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ContactTinyResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactTinyResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ContactTinyResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ContactTinyResponse) SetLastName(v string) {
	o.LastName = &v
}

func (o ContactTinyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactTinyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	return toSerialize, nil
}

type NullableContactTinyResponse struct {
	value *ContactTinyResponse
	isSet bool
}

func (v NullableContactTinyResponse) Get() *ContactTinyResponse {
	return v.value
}

func (v *NullableContactTinyResponse) Set(val *ContactTinyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactTinyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactTinyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactTinyResponse(val *ContactTinyResponse) *NullableContactTinyResponse {
	return &NullableContactTinyResponse{value: val, isSet: true}
}

func (v NullableContactTinyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactTinyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


