# VehicleStatsFuelPercent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | UTC timestamp in RFC 3339 format. Example: &#x60;2020-01-27T07:06:25Z&#x60;. | 
**Value** | **int64** | The engine fuel level in percentage points (e.g. &#x60;99&#x60;, &#x60;50&#x60;, etc). | 

## Methods

### NewVehicleStatsFuelPercent

`func NewVehicleStatsFuelPercent(time time.Time, value int64, ) *VehicleStatsFuelPercent`

NewVehicleStatsFuelPercent instantiates a new VehicleStatsFuelPercent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStatsFuelPercentWithDefaults

`func NewVehicleStatsFuelPercentWithDefaults() *VehicleStatsFuelPercent`

NewVehicleStatsFuelPercentWithDefaults instantiates a new VehicleStatsFuelPercent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *VehicleStatsFuelPercent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VehicleStatsFuelPercent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VehicleStatsFuelPercent) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetValue

`func (o *VehicleStatsFuelPercent) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VehicleStatsFuelPercent) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VehicleStatsFuelPercent) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


