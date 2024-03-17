# VehicleLocationsListResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique Samsara ID of the Vehicle. This is automatically generated when the Vehicle object is created. It cannot be changed. | 
**Locations** | [**[]VehicleLocation**](VehicleLocation.md) | A list of location events for the given vehicle. | 
**Name** | **string** | The human-readable name of the Vehicle. This is set by a fleet administrator and will appear in both Samsara’s cloud dashboard as well as the Samsara Driver mobile app. **By default**, this name is the serial number of the Samsara Vehicle Gateway. It can be set or updated through the Samsara Dashboard or through the API at any time. | 

## Methods

### NewVehicleLocationsListResponseDataInner

`func NewVehicleLocationsListResponseDataInner(id string, locations []VehicleLocation, name string, ) *VehicleLocationsListResponseDataInner`

NewVehicleLocationsListResponseDataInner instantiates a new VehicleLocationsListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLocationsListResponseDataInnerWithDefaults

`func NewVehicleLocationsListResponseDataInnerWithDefaults() *VehicleLocationsListResponseDataInner`

NewVehicleLocationsListResponseDataInnerWithDefaults instantiates a new VehicleLocationsListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VehicleLocationsListResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleLocationsListResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleLocationsListResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetLocations

`func (o *VehicleLocationsListResponseDataInner) GetLocations() []VehicleLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *VehicleLocationsListResponseDataInner) GetLocationsOk() (*[]VehicleLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *VehicleLocationsListResponseDataInner) SetLocations(v []VehicleLocation)`

SetLocations sets Locations field to given value.


### GetName

`func (o *VehicleLocationsListResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleLocationsListResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleLocationsListResponseDataInner) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


