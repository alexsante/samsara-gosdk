# AddressGeofence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Circle** | Pointer to [**AddressGeofenceCircle**](AddressGeofenceCircle.md) |  | [optional] 
**Polygon** | Pointer to [**AddressGeofencePolygon**](AddressGeofencePolygon.md) |  | [optional] 

## Methods

### NewAddressGeofence

`func NewAddressGeofence() *AddressGeofence`

NewAddressGeofence instantiates a new AddressGeofence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressGeofenceWithDefaults

`func NewAddressGeofenceWithDefaults() *AddressGeofence`

NewAddressGeofenceWithDefaults instantiates a new AddressGeofence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircle

`func (o *AddressGeofence) GetCircle() AddressGeofenceCircle`

GetCircle returns the Circle field if non-nil, zero value otherwise.

### GetCircleOk

`func (o *AddressGeofence) GetCircleOk() (*AddressGeofenceCircle, bool)`

GetCircleOk returns a tuple with the Circle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircle

`func (o *AddressGeofence) SetCircle(v AddressGeofenceCircle)`

SetCircle sets Circle field to given value.

### HasCircle

`func (o *AddressGeofence) HasCircle() bool`

HasCircle returns a boolean if a field has been set.

### GetPolygon

`func (o *AddressGeofence) GetPolygon() AddressGeofencePolygon`

GetPolygon returns the Polygon field if non-nil, zero value otherwise.

### GetPolygonOk

`func (o *AddressGeofence) GetPolygonOk() (*AddressGeofencePolygon, bool)`

GetPolygonOk returns a tuple with the Polygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolygon

`func (o *AddressGeofence) SetPolygon(v AddressGeofencePolygon)`

SetPolygon sets Polygon field to given value.

### HasPolygon

`func (o *AddressGeofence) HasPolygon() bool`

HasPolygon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


