# VehicleStatsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineState** | Pointer to [**VehicleStatsEngineState**](VehicleStatsEngineState.md) |  | [optional] 
**FuelPercent** | Pointer to [**VehicleStatsFuelPercent**](VehicleStatsFuelPercent.md) |  | [optional] 
**GpsDistanceMeters** | Pointer to [**VehicleStatsGpsDistanceMeters**](VehicleStatsGpsDistanceMeters.md) |  | [optional] 
**GpsOdometerMeters** | Pointer to [**VehicleStatsGpsOdometerMeters**](VehicleStatsGpsOdometerMeters.md) |  | [optional] 
**Id** | **string** | The unique Samsara ID of the Vehicle. This is automatically generated when the Vehicle object is created. It cannot be changed. | 
**Name** | **string** | The human-readable name of the Vehicle. This is set by a fleet administrator and will appear in both Samsara’s cloud dashboard as well as the Samsara Driver mobile app. **By default**, this name is the serial number of the Samsara Vehicle Gateway. It can be set or updated through the Samsara Dashboard or through the API at any time. | 
**ObdEngineSeconds** | Pointer to [**VehicleStatsObdEngineSeconds**](VehicleStatsObdEngineSeconds.md) |  | [optional] 
**ObdOdometerMeters** | Pointer to [**VehicleStatsObdOdometerMeters**](VehicleStatsObdOdometerMeters.md) |  | [optional] 

## Methods

### NewVehicleStatsResponseDataInner

`func NewVehicleStatsResponseDataInner(id string, name string, ) *VehicleStatsResponseDataInner`

NewVehicleStatsResponseDataInner instantiates a new VehicleStatsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStatsResponseDataInnerWithDefaults

`func NewVehicleStatsResponseDataInnerWithDefaults() *VehicleStatsResponseDataInner`

NewVehicleStatsResponseDataInnerWithDefaults instantiates a new VehicleStatsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineState

`func (o *VehicleStatsResponseDataInner) GetEngineState() VehicleStatsEngineState`

GetEngineState returns the EngineState field if non-nil, zero value otherwise.

### GetEngineStateOk

`func (o *VehicleStatsResponseDataInner) GetEngineStateOk() (*VehicleStatsEngineState, bool)`

GetEngineStateOk returns a tuple with the EngineState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineState

`func (o *VehicleStatsResponseDataInner) SetEngineState(v VehicleStatsEngineState)`

SetEngineState sets EngineState field to given value.

### HasEngineState

`func (o *VehicleStatsResponseDataInner) HasEngineState() bool`

HasEngineState returns a boolean if a field has been set.

### GetFuelPercent

`func (o *VehicleStatsResponseDataInner) GetFuelPercent() VehicleStatsFuelPercent`

GetFuelPercent returns the FuelPercent field if non-nil, zero value otherwise.

### GetFuelPercentOk

`func (o *VehicleStatsResponseDataInner) GetFuelPercentOk() (*VehicleStatsFuelPercent, bool)`

GetFuelPercentOk returns a tuple with the FuelPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelPercent

`func (o *VehicleStatsResponseDataInner) SetFuelPercent(v VehicleStatsFuelPercent)`

SetFuelPercent sets FuelPercent field to given value.

### HasFuelPercent

`func (o *VehicleStatsResponseDataInner) HasFuelPercent() bool`

HasFuelPercent returns a boolean if a field has been set.

### GetGpsDistanceMeters

`func (o *VehicleStatsResponseDataInner) GetGpsDistanceMeters() VehicleStatsGpsDistanceMeters`

GetGpsDistanceMeters returns the GpsDistanceMeters field if non-nil, zero value otherwise.

### GetGpsDistanceMetersOk

`func (o *VehicleStatsResponseDataInner) GetGpsDistanceMetersOk() (*VehicleStatsGpsDistanceMeters, bool)`

GetGpsDistanceMetersOk returns a tuple with the GpsDistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsDistanceMeters

`func (o *VehicleStatsResponseDataInner) SetGpsDistanceMeters(v VehicleStatsGpsDistanceMeters)`

SetGpsDistanceMeters sets GpsDistanceMeters field to given value.

### HasGpsDistanceMeters

`func (o *VehicleStatsResponseDataInner) HasGpsDistanceMeters() bool`

HasGpsDistanceMeters returns a boolean if a field has been set.

### GetGpsOdometerMeters

`func (o *VehicleStatsResponseDataInner) GetGpsOdometerMeters() VehicleStatsGpsOdometerMeters`

GetGpsOdometerMeters returns the GpsOdometerMeters field if non-nil, zero value otherwise.

### GetGpsOdometerMetersOk

`func (o *VehicleStatsResponseDataInner) GetGpsOdometerMetersOk() (*VehicleStatsGpsOdometerMeters, bool)`

GetGpsOdometerMetersOk returns a tuple with the GpsOdometerMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsOdometerMeters

`func (o *VehicleStatsResponseDataInner) SetGpsOdometerMeters(v VehicleStatsGpsOdometerMeters)`

SetGpsOdometerMeters sets GpsOdometerMeters field to given value.

### HasGpsOdometerMeters

`func (o *VehicleStatsResponseDataInner) HasGpsOdometerMeters() bool`

HasGpsOdometerMeters returns a boolean if a field has been set.

### GetId

`func (o *VehicleStatsResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleStatsResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleStatsResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VehicleStatsResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleStatsResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleStatsResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetObdEngineSeconds

`func (o *VehicleStatsResponseDataInner) GetObdEngineSeconds() VehicleStatsObdEngineSeconds`

GetObdEngineSeconds returns the ObdEngineSeconds field if non-nil, zero value otherwise.

### GetObdEngineSecondsOk

`func (o *VehicleStatsResponseDataInner) GetObdEngineSecondsOk() (*VehicleStatsObdEngineSeconds, bool)`

GetObdEngineSecondsOk returns a tuple with the ObdEngineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObdEngineSeconds

`func (o *VehicleStatsResponseDataInner) SetObdEngineSeconds(v VehicleStatsObdEngineSeconds)`

SetObdEngineSeconds sets ObdEngineSeconds field to given value.

### HasObdEngineSeconds

`func (o *VehicleStatsResponseDataInner) HasObdEngineSeconds() bool`

HasObdEngineSeconds returns a boolean if a field has been set.

### GetObdOdometerMeters

`func (o *VehicleStatsResponseDataInner) GetObdOdometerMeters() VehicleStatsObdOdometerMeters`

GetObdOdometerMeters returns the ObdOdometerMeters field if non-nil, zero value otherwise.

### GetObdOdometerMetersOk

`func (o *VehicleStatsResponseDataInner) GetObdOdometerMetersOk() (*VehicleStatsObdOdometerMeters, bool)`

GetObdOdometerMetersOk returns a tuple with the ObdOdometerMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObdOdometerMeters

`func (o *VehicleStatsResponseDataInner) SetObdOdometerMeters(v VehicleStatsObdOdometerMeters)`

SetObdOdometerMeters sets ObdOdometerMeters field to given value.

### HasObdOdometerMeters

`func (o *VehicleStatsResponseDataInner) HasObdOdometerMeters() bool`

HasObdOdometerMeters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


