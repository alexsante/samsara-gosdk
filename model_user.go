/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package samsarago

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User A user object.
type User struct {
	AuthType UserAuthType `json:"authType"`
	// The email address of this user.
	Email string `json:"email"`
	// ID of the user.
	Id string `json:"id"`
	// The first and last name of the user.
	Name string `json:"name"`
	// The list of roles that applies to this user. A user may have \"organizational\" roles, which apply to the user at the organizational level, and \"tag-specific\" roles, which apply to the user for a given tag.
	Roles []UserRoleAssignment `json:"roles"`
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(authType UserAuthType, email string, id string, name string, roles []UserRoleAssignment) *User {
	this := User{}
	this.AuthType = authType
	this.Email = email
	this.Id = id
	this.Name = name
	this.Roles = roles
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetAuthType returns the AuthType field value
func (o *User) GetAuthType() UserAuthType {
	if o == nil {
		var ret UserAuthType
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *User) GetAuthTypeOk() (*UserAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *User) SetAuthType(v UserAuthType) {
	o.AuthType = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *User) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *User) SetName(v string) {
	o.Name = v
}

// GetRoles returns the Roles field value
func (o *User) GetRoles() []UserRoleAssignment {
	if o == nil {
		var ret []UserRoleAssignment
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *User) GetRolesOk() ([]UserRoleAssignment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *User) SetRoles(v []UserRoleAssignment) {
	o.Roles = v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authType"] = o.AuthType
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *User) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authType",
		"email",
		"id",
		"name",
		"roles",
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

	varUser := _User{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


