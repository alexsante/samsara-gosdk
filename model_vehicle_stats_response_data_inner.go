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

// checks if the VehicleStatsResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleStatsResponseDataInner{}

// VehicleStatsResponseDataInner A vehicle and its most recent stat.
type VehicleStatsResponseDataInner struct {
	EngineState *VehicleStatsEngineState `json:"engineState,omitempty"`
	FuelPercent *VehicleStatsFuelPercent `json:"fuelPercent,omitempty"`
	GpsDistanceMeters *VehicleStatsGpsDistanceMeters `json:"gpsDistanceMeters,omitempty"`
	GpsOdometerMeters *VehicleStatsGpsOdometerMeters `json:"gpsOdometerMeters,omitempty"`
	// The unique Samsara ID of the Vehicle. This is automatically generated when the Vehicle object is created. It cannot be changed.
	Id string `json:"id"`
	// The human-readable name of the Vehicle. This is set by a fleet administrator and will appear in both Samsara’s cloud dashboard as well as the Samsara Driver mobile app. **By default**, this name is the serial number of the Samsara Vehicle Gateway. It can be set or updated through the Samsara Dashboard or through the API at any time.
	Name string `json:"name"`
	ObdEngineSeconds *VehicleStatsObdEngineSeconds `json:"obdEngineSeconds,omitempty"`
	ObdOdometerMeters *VehicleStatsObdOdometerMeters `json:"obdOdometerMeters,omitempty"`
}

type _VehicleStatsResponseDataInner VehicleStatsResponseDataInner

// NewVehicleStatsResponseDataInner instantiates a new VehicleStatsResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleStatsResponseDataInner(id string, name string) *VehicleStatsResponseDataInner {
	this := VehicleStatsResponseDataInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewVehicleStatsResponseDataInnerWithDefaults instantiates a new VehicleStatsResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleStatsResponseDataInnerWithDefaults() *VehicleStatsResponseDataInner {
	this := VehicleStatsResponseDataInner{}
	return &this
}

// GetEngineState returns the EngineState field value if set, zero value otherwise.
func (o *VehicleStatsResponseDataInner) GetEngineState() VehicleStatsEngineState {
	if o == nil || IsNil(o.EngineState) {
		var ret VehicleStatsEngineState
		return ret
	}
	return *o.EngineState
}

// GetEngineStateOk returns a tuple with the EngineState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetEngineStateOk() (*VehicleStatsEngineState, bool) {
	if o == nil || IsNil(o.EngineState) {
		return nil, false
	}
	return o.EngineState, true
}

// HasEngineState returns a boolean if a field has been set.
func (o *VehicleStatsResponseDataInner) HasEngineState() bool {
	if o != nil && !IsNil(o.EngineState) {
		return true
	}

	return false
}

// SetEngineState gets a reference to the given VehicleStatsEngineState and assigns it to the EngineState field.
func (o *VehicleStatsResponseDataInner) SetEngineState(v VehicleStatsEngineState) {
	o.EngineState = &v
}

// GetFuelPercent returns the FuelPercent field value if set, zero value otherwise.
func (o *VehicleStatsResponseDataInner) GetFuelPercent() VehicleStatsFuelPercent {
	if o == nil || IsNil(o.FuelPercent) {
		var ret VehicleStatsFuelPercent
		return ret
	}
	return *o.FuelPercent
}

// GetFuelPercentOk returns a tuple with the FuelPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetFuelPercentOk() (*VehicleStatsFuelPercent, bool) {
	if o == nil || IsNil(o.FuelPercent) {
		return nil, false
	}
	return o.FuelPercent, true
}

// HasFuelPercent returns a boolean if a field has been set.
func (o *VehicleStatsResponseDataInner) HasFuelPercent() bool {
	if o != nil && !IsNil(o.FuelPercent) {
		return true
	}

	return false
}

// SetFuelPercent gets a reference to the given VehicleStatsFuelPercent and assigns it to the FuelPercent field.
func (o *VehicleStatsResponseDataInner) SetFuelPercent(v VehicleStatsFuelPercent) {
	o.FuelPercent = &v
}

// GetGpsDistanceMeters returns the GpsDistanceMeters field value if set, zero value otherwise.
func (o *VehicleStatsResponseDataInner) GetGpsDistanceMeters() VehicleStatsGpsDistanceMeters {
	if o == nil || IsNil(o.GpsDistanceMeters) {
		var ret VehicleStatsGpsDistanceMeters
		return ret
	}
	return *o.GpsDistanceMeters
}

// GetGpsDistanceMetersOk returns a tuple with the GpsDistanceMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetGpsDistanceMetersOk() (*VehicleStatsGpsDistanceMeters, bool) {
	if o == nil || IsNil(o.GpsDistanceMeters) {
		return nil, false
	}
	return o.GpsDistanceMeters, true
}

// HasGpsDistanceMeters returns a boolean if a field has been set.
func (o *VehicleStatsResponseDataInner) HasGpsDistanceMeters() bool {
	if o != nil && !IsNil(o.GpsDistanceMeters) {
		return true
	}

	return false
}

// SetGpsDistanceMeters gets a reference to the given VehicleStatsGpsDistanceMeters and assigns it to the GpsDistanceMeters field.
func (o *VehicleStatsResponseDataInner) SetGpsDistanceMeters(v VehicleStatsGpsDistanceMeters) {
	o.GpsDistanceMeters = &v
}

// GetGpsOdometerMeters returns the GpsOdometerMeters field value if set, zero value otherwise.
func (o *VehicleStatsResponseDataInner) GetGpsOdometerMeters() VehicleStatsGpsOdometerMeters {
	if o == nil || IsNil(o.GpsOdometerMeters) {
		var ret VehicleStatsGpsOdometerMeters
		return ret
	}
	return *o.GpsOdometerMeters
}

// GetGpsOdometerMetersOk returns a tuple with the GpsOdometerMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetGpsOdometerMetersOk() (*VehicleStatsGpsOdometerMeters, bool) {
	if o == nil || IsNil(o.GpsOdometerMeters) {
		return nil, false
	}
	return o.GpsOdometerMeters, true
}

// HasGpsOdometerMeters returns a boolean if a field has been set.
func (o *VehicleStatsResponseDataInner) HasGpsOdometerMeters() bool {
	if o != nil && !IsNil(o.GpsOdometerMeters) {
		return true
	}

	return false
}

// SetGpsOdometerMeters gets a reference to the given VehicleStatsGpsOdometerMeters and assigns it to the GpsOdometerMeters field.
func (o *VehicleStatsResponseDataInner) SetGpsOdometerMeters(v VehicleStatsGpsOdometerMeters) {
	o.GpsOdometerMeters = &v
}

// GetId returns the Id field value
func (o *VehicleStatsResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VehicleStatsResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VehicleStatsResponseDataInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VehicleStatsResponseDataInner) SetName(v string) {
	o.Name = v
}

// GetObdEngineSeconds returns the ObdEngineSeconds field value if set, zero value otherwise.
func (o *VehicleStatsResponseDataInner) GetObdEngineSeconds() VehicleStatsObdEngineSeconds {
	if o == nil || IsNil(o.ObdEngineSeconds) {
		var ret VehicleStatsObdEngineSeconds
		return ret
	}
	return *o.ObdEngineSeconds
}

// GetObdEngineSecondsOk returns a tuple with the ObdEngineSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetObdEngineSecondsOk() (*VehicleStatsObdEngineSeconds, bool) {
	if o == nil || IsNil(o.ObdEngineSeconds) {
		return nil, false
	}
	return o.ObdEngineSeconds, true
}

// HasObdEngineSeconds returns a boolean if a field has been set.
func (o *VehicleStatsResponseDataInner) HasObdEngineSeconds() bool {
	if o != nil && !IsNil(o.ObdEngineSeconds) {
		return true
	}

	return false
}

// SetObdEngineSeconds gets a reference to the given VehicleStatsObdEngineSeconds and assigns it to the ObdEngineSeconds field.
func (o *VehicleStatsResponseDataInner) SetObdEngineSeconds(v VehicleStatsObdEngineSeconds) {
	o.ObdEngineSeconds = &v
}

// GetObdOdometerMeters returns the ObdOdometerMeters field value if set, zero value otherwise.
func (o *VehicleStatsResponseDataInner) GetObdOdometerMeters() VehicleStatsObdOdometerMeters {
	if o == nil || IsNil(o.ObdOdometerMeters) {
		var ret VehicleStatsObdOdometerMeters
		return ret
	}
	return *o.ObdOdometerMeters
}

// GetObdOdometerMetersOk returns a tuple with the ObdOdometerMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleStatsResponseDataInner) GetObdOdometerMetersOk() (*VehicleStatsObdOdometerMeters, bool) {
	if o == nil || IsNil(o.ObdOdometerMeters) {
		return nil, false
	}
	return o.ObdOdometerMeters, true
}

// HasObdOdometerMeters returns a boolean if a field has been set.
func (o *VehicleStatsResponseDataInner) HasObdOdometerMeters() bool {
	if o != nil && !IsNil(o.ObdOdometerMeters) {
		return true
	}

	return false
}

// SetObdOdometerMeters gets a reference to the given VehicleStatsObdOdometerMeters and assigns it to the ObdOdometerMeters field.
func (o *VehicleStatsResponseDataInner) SetObdOdometerMeters(v VehicleStatsObdOdometerMeters) {
	o.ObdOdometerMeters = &v
}

func (o VehicleStatsResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleStatsResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineState) {
		toSerialize["engineState"] = o.EngineState
	}
	if !IsNil(o.FuelPercent) {
		toSerialize["fuelPercent"] = o.FuelPercent
	}
	if !IsNil(o.GpsDistanceMeters) {
		toSerialize["gpsDistanceMeters"] = o.GpsDistanceMeters
	}
	if !IsNil(o.GpsOdometerMeters) {
		toSerialize["gpsOdometerMeters"] = o.GpsOdometerMeters
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.ObdEngineSeconds) {
		toSerialize["obdEngineSeconds"] = o.ObdEngineSeconds
	}
	if !IsNil(o.ObdOdometerMeters) {
		toSerialize["obdOdometerMeters"] = o.ObdOdometerMeters
	}
	return toSerialize, nil
}

func (o *VehicleStatsResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varVehicleStatsResponseDataInner := _VehicleStatsResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVehicleStatsResponseDataInner)

	if err != nil {
		return err
	}

	*o = VehicleStatsResponseDataInner(varVehicleStatsResponseDataInner)

	return err
}

type NullableVehicleStatsResponseDataInner struct {
	value *VehicleStatsResponseDataInner
	isSet bool
}

func (v NullableVehicleStatsResponseDataInner) Get() *VehicleStatsResponseDataInner {
	return v.value
}

func (v *NullableVehicleStatsResponseDataInner) Set(val *VehicleStatsResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleStatsResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleStatsResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleStatsResponseDataInner(val *VehicleStatsResponseDataInner) *NullableVehicleStatsResponseDataInner {
	return &NullableVehicleStatsResponseDataInner{value: val, isSet: true}
}

func (v NullableVehicleStatsResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleStatsResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


