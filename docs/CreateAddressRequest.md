# CreateAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressTypes** | Pointer to **[]string** | Reporting location type associated with the address (used for ELD reporting purposes). | [optional] 
**ContactIds** | Pointer to **[]string** | An array of Contact IDs associated with this Address. | [optional] 
**ExternalIds** | Pointer to **map[string]string** | The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object. | [optional] 
**FormattedAddress** | **string** | The full street address for this address/geofence, as it might be recognized by Google Maps. | 
**Geofence** | [**AddressGeofence**](AddressGeofence.md) |  | 
**Latitude** | Pointer to **float64** | Latitude of the address. Will be geocoded from &#x60;formattedAddress&#x60; if not provided. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the address. Will be geocoded from &#x60;formattedAddress&#x60; if not provided. | [optional] 
**Name** | **string** | Name of the address. | 
**Notes** | Pointer to **string** | Notes about the address. | [optional] 
**TagIds** | Pointer to **[]string** | An array of IDs of tags to associate with this address. | [optional] 

## Methods

### NewCreateAddressRequest

`func NewCreateAddressRequest(formattedAddress string, geofence AddressGeofence, name string, ) *CreateAddressRequest`

NewCreateAddressRequest instantiates a new CreateAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAddressRequestWithDefaults

`func NewCreateAddressRequestWithDefaults() *CreateAddressRequest`

NewCreateAddressRequestWithDefaults instantiates a new CreateAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressTypes

`func (o *CreateAddressRequest) GetAddressTypes() []string`

GetAddressTypes returns the AddressTypes field if non-nil, zero value otherwise.

### GetAddressTypesOk

`func (o *CreateAddressRequest) GetAddressTypesOk() (*[]string, bool)`

GetAddressTypesOk returns a tuple with the AddressTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTypes

`func (o *CreateAddressRequest) SetAddressTypes(v []string)`

SetAddressTypes sets AddressTypes field to given value.

### HasAddressTypes

`func (o *CreateAddressRequest) HasAddressTypes() bool`

HasAddressTypes returns a boolean if a field has been set.

### GetContactIds

`func (o *CreateAddressRequest) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *CreateAddressRequest) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *CreateAddressRequest) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *CreateAddressRequest) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetExternalIds

`func (o *CreateAddressRequest) GetExternalIds() map[string]string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *CreateAddressRequest) GetExternalIdsOk() (*map[string]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *CreateAddressRequest) SetExternalIds(v map[string]string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *CreateAddressRequest) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetFormattedAddress

`func (o *CreateAddressRequest) GetFormattedAddress() string`

GetFormattedAddress returns the FormattedAddress field if non-nil, zero value otherwise.

### GetFormattedAddressOk

`func (o *CreateAddressRequest) GetFormattedAddressOk() (*string, bool)`

GetFormattedAddressOk returns a tuple with the FormattedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAddress

`func (o *CreateAddressRequest) SetFormattedAddress(v string)`

SetFormattedAddress sets FormattedAddress field to given value.


### GetGeofence

`func (o *CreateAddressRequest) GetGeofence() AddressGeofence`

GetGeofence returns the Geofence field if non-nil, zero value otherwise.

### GetGeofenceOk

`func (o *CreateAddressRequest) GetGeofenceOk() (*AddressGeofence, bool)`

GetGeofenceOk returns a tuple with the Geofence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofence

`func (o *CreateAddressRequest) SetGeofence(v AddressGeofence)`

SetGeofence sets Geofence field to given value.


### GetLatitude

`func (o *CreateAddressRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CreateAddressRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CreateAddressRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CreateAddressRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *CreateAddressRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CreateAddressRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CreateAddressRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CreateAddressRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *CreateAddressRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAddressRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAddressRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNotes

`func (o *CreateAddressRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateAddressRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateAddressRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateAddressRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTagIds

`func (o *CreateAddressRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *CreateAddressRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *CreateAddressRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *CreateAddressRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


