# VehicleLocationsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique Samsara ID of the Vehicle. This is automatically generated when the Vehicle object is created. It cannot be changed. | 
**Location** | [**VehicleLocation**](VehicleLocation.md) |  | 
**Name** | **string** | The human-readable name of the Vehicle. This is set by a fleet administrator and will appear in both Samsara’s cloud dashboard as well as the Samsara Driver mobile app. **By default**, this name is the serial number of the Samsara Vehicle Gateway. It can be set or updated through the Samsara Dashboard or through the API at any time. | 

## Methods

### NewVehicleLocationsResponseDataInner

`func NewVehicleLocationsResponseDataInner(id string, location VehicleLocation, name string, ) *VehicleLocationsResponseDataInner`

NewVehicleLocationsResponseDataInner instantiates a new VehicleLocationsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLocationsResponseDataInnerWithDefaults

`func NewVehicleLocationsResponseDataInnerWithDefaults() *VehicleLocationsResponseDataInner`

NewVehicleLocationsResponseDataInnerWithDefaults instantiates a new VehicleLocationsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VehicleLocationsResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleLocationsResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleLocationsResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *VehicleLocationsResponseDataInner) GetLocation() VehicleLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VehicleLocationsResponseDataInner) GetLocationOk() (*VehicleLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VehicleLocationsResponseDataInner) SetLocation(v VehicleLocation)`

SetLocation sets Location field to given value.


### GetName

`func (o *VehicleLocationsResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleLocationsResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleLocationsResponseDataInner) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


