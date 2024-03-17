/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the VehicleStatsGpsOdometerMeters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleStatsGpsOdometerMeters{}

// VehicleStatsGpsOdometerMeters Vehicle GPS odometer event.
type VehicleStatsGpsOdometerMeters struct {
	// UTC timestamp in RFC 3339 format. Example: `2020-01-27T07:06:25Z`.
	Time time.Time `json:"time"`
	// Number of meters the vehicle has traveled according to the GPS calculations and the manually-specified odometer reading.
	Value int64 `json:"value"`
}

type _VehicleStatsGpsOdometerMeters VehicleStatsGpsOdometerMeters

// NewVehicleStatsGpsOdometerMeters instantiates a new VehicleStatsGpsOdometerMeters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleStatsGpsOdometerMeters(time time.Time, value int64) *VehicleStatsGpsOdometerMeters {
	this := VehicleStatsGpsOdometerMeters{}
	this.Time = time
	this.Value = value
	return &this
}

// NewVehicleStatsGpsOdometerMetersWithDefaults instantiates a new VehicleStatsGpsOdometerMeters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleStatsGpsOdometerMetersWithDefaults() *VehicleStatsGpsOdometerMeters {
	this := VehicleStatsGpsOdometerMeters{}
	return &this
}

// GetTime returns the Time field value
func (o *VehicleStatsGpsOdometerMeters) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *VehicleStatsGpsOdometerMeters) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *VehicleStatsGpsOdometerMeters) SetTime(v time.Time) {
	o.Time = v
}

// GetValue returns the Value field value
func (o *VehicleStatsGpsOdometerMeters) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VehicleStatsGpsOdometerMeters) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VehicleStatsGpsOdometerMeters) SetValue(v int64) {
	o.Value = v
}

func (o VehicleStatsGpsOdometerMeters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleStatsGpsOdometerMeters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time"] = o.Time
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *VehicleStatsGpsOdometerMeters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"time",
		"value",
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

	varVehicleStatsGpsOdometerMeters := _VehicleStatsGpsOdometerMeters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVehicleStatsGpsOdometerMeters)

	if err != nil {
		return err
	}

	*o = VehicleStatsGpsOdometerMeters(varVehicleStatsGpsOdometerMeters)

	return err
}

type NullableVehicleStatsGpsOdometerMeters struct {
	value *VehicleStatsGpsOdometerMeters
	isSet bool
}

func (v NullableVehicleStatsGpsOdometerMeters) Get() *VehicleStatsGpsOdometerMeters {
	return v.value
}

func (v *NullableVehicleStatsGpsOdometerMeters) Set(val *VehicleStatsGpsOdometerMeters) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleStatsGpsOdometerMeters) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleStatsGpsOdometerMeters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleStatsGpsOdometerMeters(val *VehicleStatsGpsOdometerMeters) *NullableVehicleStatsGpsOdometerMeters {
	return &NullableVehicleStatsGpsOdometerMeters{value: val, isSet: true}
}

func (v NullableVehicleStatsGpsOdometerMeters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleStatsGpsOdometerMeters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


