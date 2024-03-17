# VehicleLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Heading** | Pointer to **float64** | Heading of the vehicle in degrees. | [optional] 
**Latitude** | **float64** | GPS latitude represented in degrees | 
**Longitude** | **float64** | GPS longitude represented in degrees | 
**ReverseGeo** | Pointer to [**VehicleLocationReverseGeo**](VehicleLocationReverseGeo.md) |  | [optional] 
**Speed** | Pointer to **float64** | GPS speed of the vehicle in miles per hour. | [optional] 
**Time** | **time.Time** | UTC timestamp in RFC 3339 format. Example: &#x60;2020-01-27T07:06:25Z&#x60;. | 

## Methods

### NewVehicleLocation

`func NewVehicleLocation(latitude float64, longitude float64, time time.Time, ) *VehicleLocation`

NewVehicleLocation instantiates a new VehicleLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLocationWithDefaults

`func NewVehicleLocationWithDefaults() *VehicleLocation`

NewVehicleLocationWithDefaults instantiates a new VehicleLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeading

`func (o *VehicleLocation) GetHeading() float64`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *VehicleLocation) GetHeadingOk() (*float64, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *VehicleLocation) SetHeading(v float64)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *VehicleLocation) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetLatitude

`func (o *VehicleLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VehicleLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VehicleLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *VehicleLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VehicleLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VehicleLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetReverseGeo

`func (o *VehicleLocation) GetReverseGeo() VehicleLocationReverseGeo`

GetReverseGeo returns the ReverseGeo field if non-nil, zero value otherwise.

### GetReverseGeoOk

`func (o *VehicleLocation) GetReverseGeoOk() (*VehicleLocationReverseGeo, bool)`

GetReverseGeoOk returns a tuple with the ReverseGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseGeo

`func (o *VehicleLocation) SetReverseGeo(v VehicleLocationReverseGeo)`

SetReverseGeo sets ReverseGeo field to given value.

### HasReverseGeo

`func (o *VehicleLocation) HasReverseGeo() bool`

HasReverseGeo returns a boolean if a field has been set.

### GetSpeed

`func (o *VehicleLocation) GetSpeed() float64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VehicleLocation) GetSpeedOk() (*float64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VehicleLocation) SetSpeed(v float64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VehicleLocation) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTime

`func (o *VehicleLocation) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VehicleLocation) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VehicleLocation) SetTime(v time.Time)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


