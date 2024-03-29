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

// checks if the UpdateAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAddressRequest{}

// UpdateAddressRequest A request body to update an Address.
type UpdateAddressRequest struct {
	// Reporting location type associated with the address (used for ELD reporting purposes).
	AddressTypes []string `json:"addressTypes,omitempty"`
	// An array of Contact IDs associated with this Address.
	ContactIds []string `json:"contactIds,omitempty"`
	// The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object.
	ExternalIds *map[string]string `json:"externalIds,omitempty"`
	// The full street address for this address/geofence, as it might be recognized by Google Maps.
	FormattedAddress *string `json:"formattedAddress,omitempty"`
	Geofence *AddressGeofence `json:"geofence,omitempty"`
	// Latitude of the address. Will be geocoded from `formattedAddress` if not provided.
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude of the address. Will be geocoded from `formattedAddress` if not provided.
	Longitude *float64 `json:"longitude,omitempty"`
	// Name of the address.
	Name *string `json:"name,omitempty"`
	// Notes about the address.
	Notes *string `json:"notes,omitempty"`
	// An array of IDs of tags to associate with this address.
	TagIds []string `json:"tagIds,omitempty"`
}

// NewUpdateAddressRequest instantiates a new UpdateAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAddressRequest() *UpdateAddressRequest {
	this := UpdateAddressRequest{}
	return &this
}

// NewUpdateAddressRequestWithDefaults instantiates a new UpdateAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAddressRequestWithDefaults() *UpdateAddressRequest {
	this := UpdateAddressRequest{}
	return &this
}

// GetAddressTypes returns the AddressTypes field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetAddressTypes() []string {
	if o == nil || IsNil(o.AddressTypes) {
		var ret []string
		return ret
	}
	return o.AddressTypes
}

// GetAddressTypesOk returns a tuple with the AddressTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetAddressTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AddressTypes) {
		return nil, false
	}
	return o.AddressTypes, true
}

// HasAddressTypes returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasAddressTypes() bool {
	if o != nil && !IsNil(o.AddressTypes) {
		return true
	}

	return false
}

// SetAddressTypes gets a reference to the given []string and assigns it to the AddressTypes field.
func (o *UpdateAddressRequest) SetAddressTypes(v []string) {
	o.AddressTypes = v
}

// GetContactIds returns the ContactIds field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetContactIds() []string {
	if o == nil || IsNil(o.ContactIds) {
		var ret []string
		return ret
	}
	return o.ContactIds
}

// GetContactIdsOk returns a tuple with the ContactIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetContactIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ContactIds) {
		return nil, false
	}
	return o.ContactIds, true
}

// HasContactIds returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasContactIds() bool {
	if o != nil && !IsNil(o.ContactIds) {
		return true
	}

	return false
}

// SetContactIds gets a reference to the given []string and assigns it to the ContactIds field.
func (o *UpdateAddressRequest) SetContactIds(v []string) {
	o.ContactIds = v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetExternalIds() map[string]string {
	if o == nil || IsNil(o.ExternalIds) {
		var ret map[string]string
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetExternalIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given map[string]string and assigns it to the ExternalIds field.
func (o *UpdateAddressRequest) SetExternalIds(v map[string]string) {
	o.ExternalIds = &v
}

// GetFormattedAddress returns the FormattedAddress field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetFormattedAddress() string {
	if o == nil || IsNil(o.FormattedAddress) {
		var ret string
		return ret
	}
	return *o.FormattedAddress
}

// GetFormattedAddressOk returns a tuple with the FormattedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetFormattedAddressOk() (*string, bool) {
	if o == nil || IsNil(o.FormattedAddress) {
		return nil, false
	}
	return o.FormattedAddress, true
}

// HasFormattedAddress returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasFormattedAddress() bool {
	if o != nil && !IsNil(o.FormattedAddress) {
		return true
	}

	return false
}

// SetFormattedAddress gets a reference to the given string and assigns it to the FormattedAddress field.
func (o *UpdateAddressRequest) SetFormattedAddress(v string) {
	o.FormattedAddress = &v
}

// GetGeofence returns the Geofence field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetGeofence() AddressGeofence {
	if o == nil || IsNil(o.Geofence) {
		var ret AddressGeofence
		return ret
	}
	return *o.Geofence
}

// GetGeofenceOk returns a tuple with the Geofence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetGeofenceOk() (*AddressGeofence, bool) {
	if o == nil || IsNil(o.Geofence) {
		return nil, false
	}
	return o.Geofence, true
}

// HasGeofence returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasGeofence() bool {
	if o != nil && !IsNil(o.Geofence) {
		return true
	}

	return false
}

// SetGeofence gets a reference to the given AddressGeofence and assigns it to the Geofence field.
func (o *UpdateAddressRequest) SetGeofence(v AddressGeofence) {
	o.Geofence = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *UpdateAddressRequest) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *UpdateAddressRequest) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAddressRequest) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateAddressRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *UpdateAddressRequest) GetTagIds() []string {
	if o == nil || IsNil(o.TagIds) {
		var ret []string
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAddressRequest) GetTagIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TagIds) {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *UpdateAddressRequest) HasTagIds() bool {
	if o != nil && !IsNil(o.TagIds) {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []string and assigns it to the TagIds field.
func (o *UpdateAddressRequest) SetTagIds(v []string) {
	o.TagIds = v
}

func (o UpdateAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressTypes) {
		toSerialize["addressTypes"] = o.AddressTypes
	}
	if !IsNil(o.ContactIds) {
		toSerialize["contactIds"] = o.ContactIds
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !IsNil(o.FormattedAddress) {
		toSerialize["formattedAddress"] = o.FormattedAddress
	}
	if !IsNil(o.Geofence) {
		toSerialize["geofence"] = o.Geofence
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.TagIds) {
		toSerialize["tagIds"] = o.TagIds
	}
	return toSerialize, nil
}

type NullableUpdateAddressRequest struct {
	value *UpdateAddressRequest
	isSet bool
}

func (v NullableUpdateAddressRequest) Get() *UpdateAddressRequest {
	return v.value
}

func (v *NullableUpdateAddressRequest) Set(val *UpdateAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAddressRequest(val *UpdateAddressRequest) *NullableUpdateAddressRequest {
	return &NullableUpdateAddressRequest{value: val, isSet: true}
}

func (v NullableUpdateAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


