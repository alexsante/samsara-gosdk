/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address An Address object.
type Address struct {
	// Reporting location type associated with the address (used for ELD reporting purposes).
	AddressTypes []string `json:"addressTypes,omitempty"`
	// An array Contact mini-objects that are associated the Address.
	Contacts []ContactTinyResponse `json:"contacts,omitempty"`
	// The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object.
	ExternalIds *map[string]string `json:"externalIds,omitempty"`
	// The full street address for this address/geofence, as it might be recognized by Google Maps.
	FormattedAddress string `json:"formattedAddress"`
	Geofence AddressGeofence `json:"geofence"`
	// ID of the Address.
	Id string `json:"id"`
	// Latitude of the address. Will be geocoded from `formattedAddress` if not provided.
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude of the address. Will be geocoded from `formattedAddress` if not provided.
	Longitude *float64 `json:"longitude,omitempty"`
	// Name of the address.
	Name string `json:"name"`
	// Notes about the address.
	Notes *string `json:"notes,omitempty"`
	// An array of all tag mini-objects that are associated with the given address entry.
	Tags []TagTinyResponse `json:"tags,omitempty"`
}

type _Address Address

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(formattedAddress string, geofence AddressGeofence, id string, name string) *Address {
	this := Address{}
	this.FormattedAddress = formattedAddress
	this.Geofence = geofence
	this.Id = id
	this.Name = name
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetAddressTypes returns the AddressTypes field value if set, zero value otherwise.
func (o *Address) GetAddressTypes() []string {
	if o == nil || IsNil(o.AddressTypes) {
		var ret []string
		return ret
	}
	return o.AddressTypes
}

// GetAddressTypesOk returns a tuple with the AddressTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetAddressTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AddressTypes) {
		return nil, false
	}
	return o.AddressTypes, true
}

// HasAddressTypes returns a boolean if a field has been set.
func (o *Address) HasAddressTypes() bool {
	if o != nil && !IsNil(o.AddressTypes) {
		return true
	}

	return false
}

// SetAddressTypes gets a reference to the given []string and assigns it to the AddressTypes field.
func (o *Address) SetAddressTypes(v []string) {
	o.AddressTypes = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *Address) GetContacts() []ContactTinyResponse {
	if o == nil || IsNil(o.Contacts) {
		var ret []ContactTinyResponse
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetContactsOk() ([]ContactTinyResponse, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *Address) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []ContactTinyResponse and assigns it to the Contacts field.
func (o *Address) SetContacts(v []ContactTinyResponse) {
	o.Contacts = v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *Address) GetExternalIds() map[string]string {
	if o == nil || IsNil(o.ExternalIds) {
		var ret map[string]string
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetExternalIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *Address) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given map[string]string and assigns it to the ExternalIds field.
func (o *Address) SetExternalIds(v map[string]string) {
	o.ExternalIds = &v
}

// GetFormattedAddress returns the FormattedAddress field value
func (o *Address) GetFormattedAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormattedAddress
}

// GetFormattedAddressOk returns a tuple with the FormattedAddress field value
// and a boolean to check if the value has been set.
func (o *Address) GetFormattedAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormattedAddress, true
}

// SetFormattedAddress sets field value
func (o *Address) SetFormattedAddress(v string) {
	o.FormattedAddress = v
}

// GetGeofence returns the Geofence field value
func (o *Address) GetGeofence() AddressGeofence {
	if o == nil {
		var ret AddressGeofence
		return ret
	}

	return o.Geofence
}

// GetGeofenceOk returns a tuple with the Geofence field value
// and a boolean to check if the value has been set.
func (o *Address) GetGeofenceOk() (*AddressGeofence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Geofence, true
}

// SetGeofence sets field value
func (o *Address) SetGeofence(v AddressGeofence) {
	o.Geofence = v
}

// GetId returns the Id field value
func (o *Address) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Address) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Address) SetId(v string) {
	o.Id = v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *Address) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *Address) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *Address) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *Address) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *Address) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *Address) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetName returns the Name field value
func (o *Address) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Address) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Address) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Address) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Address) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Address) SetNotes(v string) {
	o.Notes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Address) GetTags() []TagTinyResponse {
	if o == nil || IsNil(o.Tags) {
		var ret []TagTinyResponse
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetTagsOk() ([]TagTinyResponse, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Address) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagTinyResponse and assigns it to the Tags field.
func (o *Address) SetTags(v []TagTinyResponse) {
	o.Tags = v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressTypes) {
		toSerialize["addressTypes"] = o.AddressTypes
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	toSerialize["formattedAddress"] = o.FormattedAddress
	toSerialize["geofence"] = o.Geofence
	toSerialize["id"] = o.Id
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *Address) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"formattedAddress",
		"geofence",
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

	varAddress := _Address{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddress)

	if err != nil {
		return err
	}

	*o = Address(varAddress)

	return err
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

