/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// UserAuthType The authentication type the user uses to authenticate. To use SAML this organization must have a configured SAML integration.
type UserAuthType string

// List of UserAuthType
const (
	DEFAULT UserAuthType = "default"
	SAML UserAuthType = "saml"
)

// All allowed values of UserAuthType enum
var AllowedUserAuthTypeEnumValues = []UserAuthType{
	"default",
	"saml",
}

func (v *UserAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserAuthType(value)
	for _, existing := range AllowedUserAuthTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserAuthType", value)
}

// NewUserAuthTypeFromValue returns a pointer to a valid UserAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserAuthTypeFromValue(v string) (*UserAuthType, error) {
	ev := UserAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserAuthType: valid values are %v", v, AllowedUserAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserAuthType) IsValid() bool {
	for _, existing := range AllowedUserAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAuthType value
func (v UserAuthType) Ptr() *UserAuthType {
	return &v
}

type NullableUserAuthType struct {
	value *UserAuthType
	isSet bool
}

func (v NullableUserAuthType) Get() *UserAuthType {
	return v.value
}

func (v *NullableUserAuthType) Set(val *UserAuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuthType(val *UserAuthType) *NullableUserAuthType {
	return &NullableUserAuthType{value: val, isSet: true}
}

func (v NullableUserAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
