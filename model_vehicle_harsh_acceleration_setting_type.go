/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package samsarago

import (
	"encoding/json"
	"fmt"
)

// VehicleHarshAccelerationSettingType The harsh acceleration setting type. This setting influences the acceleration sensitivity from which a [harsh event](https://kb.samsara.com/hc/en-us/articles/360006938891-Harsh-Events) is triggered. **By default**, this setting is inferred by the Samsara Vehicle Gateway from the engine computer, but it may be set or updated through the Samsara Dashboard or the API at any time. If set to `off`, then no acceleration based harsh events are triggered for the vehicle.
type VehicleHarshAccelerationSettingType string

// List of VehicleHarshAccelerationSettingType
const (
	PASSENGER_CAR VehicleHarshAccelerationSettingType = "passengerCar"
	LIGHT_TRUCK VehicleHarshAccelerationSettingType = "lightTruck"
	HEAVY_DUTY VehicleHarshAccelerationSettingType = "heavyDuty"
	OFF VehicleHarshAccelerationSettingType = "off"
	AUTOMATIC VehicleHarshAccelerationSettingType = "automatic"
)

// All allowed values of VehicleHarshAccelerationSettingType enum
var AllowedVehicleHarshAccelerationSettingTypeEnumValues = []VehicleHarshAccelerationSettingType{
	"passengerCar",
	"lightTruck",
	"heavyDuty",
	"off",
	"automatic",
}

func (v *VehicleHarshAccelerationSettingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VehicleHarshAccelerationSettingType(value)
	for _, existing := range AllowedVehicleHarshAccelerationSettingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VehicleHarshAccelerationSettingType", value)
}

// NewVehicleHarshAccelerationSettingTypeFromValue returns a pointer to a valid VehicleHarshAccelerationSettingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVehicleHarshAccelerationSettingTypeFromValue(v string) (*VehicleHarshAccelerationSettingType, error) {
	ev := VehicleHarshAccelerationSettingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VehicleHarshAccelerationSettingType: valid values are %v", v, AllowedVehicleHarshAccelerationSettingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleHarshAccelerationSettingType) IsValid() bool {
	for _, existing := range AllowedVehicleHarshAccelerationSettingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VehicleHarshAccelerationSettingType value
func (v VehicleHarshAccelerationSettingType) Ptr() *VehicleHarshAccelerationSettingType {
	return &v
}

type NullableVehicleHarshAccelerationSettingType struct {
	value *VehicleHarshAccelerationSettingType
	isSet bool
}

func (v NullableVehicleHarshAccelerationSettingType) Get() *VehicleHarshAccelerationSettingType {
	return v.value
}

func (v *NullableVehicleHarshAccelerationSettingType) Set(val *VehicleHarshAccelerationSettingType) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleHarshAccelerationSettingType) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleHarshAccelerationSettingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleHarshAccelerationSettingType(val *VehicleHarshAccelerationSettingType) *NullableVehicleHarshAccelerationSettingType {
	return &NullableVehicleHarshAccelerationSettingType{value: val, isSet: true}
}

func (v NullableVehicleHarshAccelerationSettingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleHarshAccelerationSettingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

