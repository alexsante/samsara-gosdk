# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressTypes** | Pointer to **[]string** | Reporting location type associated with the address (used for ELD reporting purposes). | [optional] 
**Contacts** | Pointer to [**[]ContactTinyResponse**](ContactTinyResponse.md) | An array Contact mini-objects that are associated the Address. | [optional] 
**ExternalIds** | Pointer to **map[string]string** | The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object. | [optional] 
**FormattedAddress** | **string** | The full street address for this address/geofence, as it might be recognized by Google Maps. | 
**Geofence** | [**AddressGeofence**](AddressGeofence.md) |  | 
**Id** | **string** | ID of the Address. | 
**Latitude** | Pointer to **float64** | Latitude of the address. Will be geocoded from &#x60;formattedAddress&#x60; if not provided. | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the address. Will be geocoded from &#x60;formattedAddress&#x60; if not provided. | [optional] 
**Name** | **string** | Name of the address. | 
**Notes** | Pointer to **string** | Notes about the address. | [optional] 
**Tags** | Pointer to [**[]TagTinyResponse**](TagTinyResponse.md) | An array of all tag mini-objects that are associated with the given address entry. | [optional] 

## Methods

### NewAddress

`func NewAddress(formattedAddress string, geofence AddressGeofence, id string, name string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressTypes

`func (o *Address) GetAddressTypes() []string`

GetAddressTypes returns the AddressTypes field if non-nil, zero value otherwise.

### GetAddressTypesOk

`func (o *Address) GetAddressTypesOk() (*[]string, bool)`

GetAddressTypesOk returns a tuple with the AddressTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTypes

`func (o *Address) SetAddressTypes(v []string)`

SetAddressTypes sets AddressTypes field to given value.

### HasAddressTypes

`func (o *Address) HasAddressTypes() bool`

HasAddressTypes returns a boolean if a field has been set.

### GetContacts

`func (o *Address) GetContacts() []ContactTinyResponse`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Address) GetContactsOk() (*[]ContactTinyResponse, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Address) SetContacts(v []ContactTinyResponse)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Address) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetExternalIds

`func (o *Address) GetExternalIds() map[string]string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *Address) GetExternalIdsOk() (*map[string]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *Address) SetExternalIds(v map[string]string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *Address) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetFormattedAddress

`func (o *Address) GetFormattedAddress() string`

GetFormattedAddress returns the FormattedAddress field if non-nil, zero value otherwise.

### GetFormattedAddressOk

`func (o *Address) GetFormattedAddressOk() (*string, bool)`

GetFormattedAddressOk returns a tuple with the FormattedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAddress

`func (o *Address) SetFormattedAddress(v string)`

SetFormattedAddress sets FormattedAddress field to given value.


### GetGeofence

`func (o *Address) GetGeofence() AddressGeofence`

GetGeofence returns the Geofence field if non-nil, zero value otherwise.

### GetGeofenceOk

`func (o *Address) GetGeofenceOk() (*AddressGeofence, bool)`

GetGeofenceOk returns a tuple with the Geofence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeofence

`func (o *Address) SetGeofence(v AddressGeofence)`

SetGeofence sets Geofence field to given value.


### GetId

`func (o *Address) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v string)`

SetId sets Id field to given value.


### GetLatitude

`func (o *Address) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Address) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Address) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Address) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *Address) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Address) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Address) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Address) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *Address) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Address) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Address) SetName(v string)`

SetName sets Name field to given value.


### GetNotes

`func (o *Address) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Address) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Address) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Address) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *Address) GetTags() []TagTinyResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Address) GetTagsOk() (*[]TagTinyResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Address) SetTags(v []TagTinyResponse)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Address) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


