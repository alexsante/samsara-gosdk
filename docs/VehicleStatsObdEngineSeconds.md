# VehicleStatsObdEngineSeconds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | UTC timestamp in RFC 3339 format. Example: &#x60;2020-01-27T07:06:25Z&#x60;. | 
**Value** | **int64** | Number of seconds the vehicle&#39;s engine has been on according to the on-baord diagnostics. | 

## Methods

### NewVehicleStatsObdEngineSeconds

`func NewVehicleStatsObdEngineSeconds(time time.Time, value int64, ) *VehicleStatsObdEngineSeconds`

NewVehicleStatsObdEngineSeconds instantiates a new VehicleStatsObdEngineSeconds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStatsObdEngineSecondsWithDefaults

`func NewVehicleStatsObdEngineSecondsWithDefaults() *VehicleStatsObdEngineSeconds`

NewVehicleStatsObdEngineSecondsWithDefaults instantiates a new VehicleStatsObdEngineSeconds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *VehicleStatsObdEngineSeconds) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VehicleStatsObdEngineSeconds) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VehicleStatsObdEngineSeconds) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetValue

`func (o *VehicleStatsObdEngineSeconds) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VehicleStatsObdEngineSeconds) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VehicleStatsObdEngineSeconds) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


