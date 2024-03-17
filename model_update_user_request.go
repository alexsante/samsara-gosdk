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

// checks if the UpdateUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserRequest{}

// UpdateUserRequest The user update arguments
type UpdateUserRequest struct {
	AuthType *UserAuthType `json:"authType,omitempty"`
	// The first and last name of the user.
	Name *string `json:"name,omitempty"`
	// The list of roles that applies to this user. A user may have \"organizational\" roles, which apply to the user at the organizational level, and \"tag-specific\" roles, which apply to the user for a given tag.
	Roles []UserRoleAssignmentRequest `json:"roles,omitempty"`
}

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetAuthType() UserAuthType {
	if o == nil || IsNil(o.AuthType) {
		var ret UserAuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetAuthTypeOk() (*UserAuthType, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given UserAuthType and assigns it to the AuthType field.
func (o *UpdateUserRequest) SetAuthType(v UserAuthType) {
	o.AuthType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateUserRequest) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetRoles() []UserRoleAssignmentRequest {
	if o == nil || IsNil(o.Roles) {
		var ret []UserRoleAssignmentRequest
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetRolesOk() ([]UserRoleAssignmentRequest, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRoleAssignmentRequest and assigns it to the Roles field.
func (o *UpdateUserRequest) SetRoles(v []UserRoleAssignmentRequest) {
	o.Roles = v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

