# AddressGeofencePolygon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vertices** | [**[]AddressGeofencePolygonVerticesInner**](AddressGeofencePolygonVerticesInner.md) | The vertices of the polygon geofence. These geofence vertices describe the perimeter of the polygon, and must consist of at least 3 vertices and less than 40. | 

## Methods

### NewAddressGeofencePolygon

`func NewAddressGeofencePolygon(vertices []AddressGeofencePolygonVerticesInner, ) *AddressGeofencePolygon`

NewAddressGeofencePolygon instantiates a new AddressGeofencePolygon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressGeofencePolygonWithDefaults

`func NewAddressGeofencePolygonWithDefaults() *AddressGeofencePolygon`

NewAddressGeofencePolygonWithDefaults instantiates a new AddressGeofencePolygon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVertices

`func (o *AddressGeofencePolygon) GetVertices() []AddressGeofencePolygonVerticesInner`

GetVertices returns the Vertices field if non-nil, zero value otherwise.

### GetVerticesOk

`func (o *AddressGeofencePolygon) GetVerticesOk() (*[]AddressGeofencePolygonVerticesInner, bool)`

GetVerticesOk returns a tuple with the Vertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertices

`func (o *AddressGeofencePolygon) SetVertices(v []AddressGeofencePolygonVerticesInner)`

SetVertices sets Vertices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


