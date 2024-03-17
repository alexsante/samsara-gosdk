# VehicleStatsListResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineStates** | Pointer to [**[]VehicleStatsEngineState**](VehicleStatsEngineState.md) | A list of engine state events for the given vehicle. | [optional] 
**FuelPercents** | Pointer to [**[]VehicleStatsFuelPercent**](VehicleStatsFuelPercent.md) | A list of fuel percentage readings for the given vehicle. | [optional] 
**GpsDistanceMeters** | Pointer to [**[]VehicleStatsGpsDistanceMeters**](VehicleStatsGpsDistanceMeters.md) | A list of GPS distance events for the given vehicle. | [optional] 
**GpsOdometerMeters** | Pointer to [**[]VehicleStatsGpsOdometerMeters**](VehicleStatsGpsOdometerMeters.md) | A list of GPS odometer events for the given vehicle. | [optional] 
**Id** | **string** | The unique Samsara ID of the Vehicle. This is automatically generated when the Vehicle object is created. It cannot be changed. | 
**Name** | **string** | The human-readable name of the Vehicle. This is set by a fleet administrator and will appear in both Samsara’s cloud dashboard as well as the Samsara Driver mobile app. **By default**, this name is the serial number of the Samsara Vehicle Gateway. It can be set or updated through the Samsara Dashboard or through the API at any time. | 
**ObdEngineSeconds** | Pointer to [**[]VehicleStatsObdEngineSeconds**](VehicleStatsObdEngineSeconds.md) | A list of OBD engine seconds readings for the given vehicle. | [optional] 
**ObdOdometerMeters** | Pointer to [**[]VehicleStatsObdOdometerMeters**](VehicleStatsObdOdometerMeters.md) | A list of OBD odometer readings for the given vehicle. | [optional] 

## Methods

### NewVehicleStatsListResponseDataInner

`func NewVehicleStatsListResponseDataInner(id string, name string, ) *VehicleStatsListResponseDataInner`

NewVehicleStatsListResponseDataInner instantiates a new VehicleStatsListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStatsListResponseDataInnerWithDefaults

`func NewVehicleStatsListResponseDataInnerWithDefaults() *VehicleStatsListResponseDataInner`

NewVehicleStatsListResponseDataInnerWithDefaults instantiates a new VehicleStatsListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineStates

`func (o *VehicleStatsListResponseDataInner) GetEngineStates() []VehicleStatsEngineState`

GetEngineStates returns the EngineStates field if non-nil, zero value otherwise.

### GetEngineStatesOk

`func (o *VehicleStatsListResponseDataInner) GetEngineStatesOk() (*[]VehicleStatsEngineState, bool)`

GetEngineStatesOk returns a tuple with the EngineStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineStates

`func (o *VehicleStatsListResponseDataInner) SetEngineStates(v []VehicleStatsEngineState)`

SetEngineStates sets EngineStates field to given value.

### HasEngineStates

`func (o *VehicleStatsListResponseDataInner) HasEngineStates() bool`

HasEngineStates returns a boolean if a field has been set.

### GetFuelPercents

`func (o *VehicleStatsListResponseDataInner) GetFuelPercents() []VehicleStatsFuelPercent`

GetFuelPercents returns the FuelPercents field if non-nil, zero value otherwise.

### GetFuelPercentsOk

`func (o *VehicleStatsListResponseDataInner) GetFuelPercentsOk() (*[]VehicleStatsFuelPercent, bool)`

GetFuelPercentsOk returns a tuple with the FuelPercents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelPercents

`func (o *VehicleStatsListResponseDataInner) SetFuelPercents(v []VehicleStatsFuelPercent)`

SetFuelPercents sets FuelPercents field to given value.

### HasFuelPercents

`func (o *VehicleStatsListResponseDataInner) HasFuelPercents() bool`

HasFuelPercents returns a boolean if a field has been set.

### GetGpsDistanceMeters

`func (o *VehicleStatsListResponseDataInner) GetGpsDistanceMeters() []VehicleStatsGpsDistanceMeters`

GetGpsDistanceMeters returns the GpsDistanceMeters field if non-nil, zero value otherwise.

### GetGpsDistanceMetersOk

`func (o *VehicleStatsListResponseDataInner) GetGpsDistanceMetersOk() (*[]VehicleStatsGpsDistanceMeters, bool)`

GetGpsDistanceMetersOk returns a tuple with the GpsDistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsDistanceMeters

`func (o *VehicleStatsListResponseDataInner) SetGpsDistanceMeters(v []VehicleStatsGpsDistanceMeters)`

SetGpsDistanceMeters sets GpsDistanceMeters field to given value.

### HasGpsDistanceMeters

`func (o *VehicleStatsListResponseDataInner) HasGpsDistanceMeters() bool`

HasGpsDistanceMeters returns a boolean if a field has been set.

### GetGpsOdometerMeters

`func (o *VehicleStatsListResponseDataInner) GetGpsOdometerMeters() []VehicleStatsGpsOdometerMeters`

GetGpsOdometerMeters returns the GpsOdometerMeters field if non-nil, zero value otherwise.

### GetGpsOdometerMetersOk

`func (o *VehicleStatsListResponseDataInner) GetGpsOdometerMetersOk() (*[]VehicleStatsGpsOdometerMeters, bool)`

GetGpsOdometerMetersOk returns a tuple with the GpsOdometerMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsOdometerMeters

`func (o *VehicleStatsListResponseDataInner) SetGpsOdometerMeters(v []VehicleStatsGpsOdometerMeters)`

SetGpsOdometerMeters sets GpsOdometerMeters field to given value.

### HasGpsOdometerMeters

`func (o *VehicleStatsListResponseDataInner) HasGpsOdometerMeters() bool`

HasGpsOdometerMeters returns a boolean if a field has been set.

### GetId

`func (o *VehicleStatsListResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleStatsListResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleStatsListResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VehicleStatsListResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleStatsListResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleStatsListResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetObdEngineSeconds

`func (o *VehicleStatsListResponseDataInner) GetObdEngineSeconds() []VehicleStatsObdEngineSeconds`

GetObdEngineSeconds returns the ObdEngineSeconds field if non-nil, zero value otherwise.

### GetObdEngineSecondsOk

`func (o *VehicleStatsListResponseDataInner) GetObdEngineSecondsOk() (*[]VehicleStatsObdEngineSeconds, bool)`

GetObdEngineSecondsOk returns a tuple with the ObdEngineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObdEngineSeconds

`func (o *VehicleStatsListResponseDataInner) SetObdEngineSeconds(v []VehicleStatsObdEngineSeconds)`

SetObdEngineSeconds sets ObdEngineSeconds field to given value.

### HasObdEngineSeconds

`func (o *VehicleStatsListResponseDataInner) HasObdEngineSeconds() bool`

HasObdEngineSeconds returns a boolean if a field has been set.

### GetObdOdometerMeters

`func (o *VehicleStatsListResponseDataInner) GetObdOdometerMeters() []VehicleStatsObdOdometerMeters`

GetObdOdometerMeters returns the ObdOdometerMeters field if non-nil, zero value otherwise.

### GetObdOdometerMetersOk

`func (o *VehicleStatsListResponseDataInner) GetObdOdometerMetersOk() (*[]VehicleStatsObdOdometerMeters, bool)`

GetObdOdometerMetersOk returns a tuple with the ObdOdometerMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObdOdometerMeters

`func (o *VehicleStatsListResponseDataInner) SetObdOdometerMeters(v []VehicleStatsObdOdometerMeters)`

SetObdOdometerMeters sets ObdOdometerMeters field to given value.

### HasObdOdometerMeters

`func (o *VehicleStatsListResponseDataInner) HasObdOdometerMeters() bool`

HasObdOdometerMeters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


