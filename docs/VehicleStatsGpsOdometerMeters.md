# VehicleStatsGpsOdometerMeters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | UTC timestamp in RFC 3339 format. Example: &#x60;2020-01-27T07:06:25Z&#x60;. | 
**Value** | **int64** | Number of meters the vehicle has traveled according to the GPS calculations and the manually-specified odometer reading. | 

## Methods

### NewVehicleStatsGpsOdometerMeters

`func NewVehicleStatsGpsOdometerMeters(time time.Time, value int64, ) *VehicleStatsGpsOdometerMeters`

NewVehicleStatsGpsOdometerMeters instantiates a new VehicleStatsGpsOdometerMeters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStatsGpsOdometerMetersWithDefaults

`func NewVehicleStatsGpsOdometerMetersWithDefaults() *VehicleStatsGpsOdometerMeters`

NewVehicleStatsGpsOdometerMetersWithDefaults instantiates a new VehicleStatsGpsOdometerMeters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *VehicleStatsGpsOdometerMeters) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VehicleStatsGpsOdometerMeters) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VehicleStatsGpsOdometerMeters) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetValue

`func (o *VehicleStatsGpsOdometerMeters) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VehicleStatsGpsOdometerMeters) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VehicleStatsGpsOdometerMeters) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


