# UpdateVehicleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuxInputType1** | Pointer to [**VehicleAuxInputType**](VehicleAuxInputType.md) |  | [optional] 
**AuxInputType2** | Pointer to [**VehicleAuxInputType**](VehicleAuxInputType.md) |  | [optional] 
**EngineHours** | Pointer to **int64** | A manual override for the vehicle&#39;s engine hours. You may only override a vehicle&#39;s engine hours if it cannot be read from on-board diagnostics. When you provide a manual engine hours override, Samsara will begin updating a vehicle&#39;s engine hours based on when the Samsara Vehicle Gateway is recieving power or not. | [optional] 
**ExternalIds** | Pointer to **map[string]string** | The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object. | [optional] 
**HarshAccelerationSettingType** | Pointer to [**VehicleHarshAccelerationSettingType**](VehicleHarshAccelerationSettingType.md) |  | [optional] 
**LicensePlate** | Pointer to **string** | The license plate of the Vehicle. **By default**: empty. Can be set or updated through the Samsara Dashboard or the API at any time. | [optional] 
**Name** | Pointer to **string** | The human-readable name of the Vehicle. This is set by a fleet administrator and will appear in both Samsara’s cloud dashboard as well as the Samsara Driver mobile app. **By default**, this name is the serial number of the Samsara Vehicle Gateway. It can be set or updated through the Samsara Dashboard or through the API at any time. | [optional] 
**Notes** | Pointer to **string** | These are generic notes about the Vehicle. Empty by default. Can be set or updated through the Samsara Dashboard or the API at any time. | [optional] [default to ""]
**OdometerMeters** | Pointer to **int64** | A manual override for the vehicle&#39;s odometer. You may only override a vehicle&#39;s odometer if it cannot be read from on-board diagnostics. When you provide a manual odometer override, Samsara will begin updating a vehicle&#39;s odometer using GPS distance traveled since this override was set. See [here](https://kb.samsara.com/hc/en-us/articles/115005273667) for more details. | [optional] 
**StaticAssignedDriverId** | Pointer to **string** | ID for the static assigned driver of the vehicle. | [optional] 
**TagIds** | Pointer to **[]string** | An array of IDs of tags to associate with this vehicle. | [optional] 
**Vin** | Pointer to **string** | The VIN of the Vehicle. Most of the time, this will be automatically read from the engine computer by the Samsara Vehicle Gateway. It will be empty if it cannot be read. It can be set or updated through the Samsara Dashboard or the API at any time. | [optional] 

## Methods

### NewUpdateVehicleRequest

`func NewUpdateVehicleRequest() *UpdateVehicleRequest`

NewUpdateVehicleRequest instantiates a new UpdateVehicleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVehicleRequestWithDefaults

`func NewUpdateVehicleRequestWithDefaults() *UpdateVehicleRequest`

NewUpdateVehicleRequestWithDefaults instantiates a new UpdateVehicleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxInputType1

`func (o *UpdateVehicleRequest) GetAuxInputType1() VehicleAuxInputType`

GetAuxInputType1 returns the AuxInputType1 field if non-nil, zero value otherwise.

### GetAuxInputType1Ok

`func (o *UpdateVehicleRequest) GetAuxInputType1Ok() (*VehicleAuxInputType, bool)`

GetAuxInputType1Ok returns a tuple with the AuxInputType1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxInputType1

`func (o *UpdateVehicleRequest) SetAuxInputType1(v VehicleAuxInputType)`

SetAuxInputType1 sets AuxInputType1 field to given value.

### HasAuxInputType1

`func (o *UpdateVehicleRequest) HasAuxInputType1() bool`

HasAuxInputType1 returns a boolean if a field has been set.

### GetAuxInputType2

`func (o *UpdateVehicleRequest) GetAuxInputType2() VehicleAuxInputType`

GetAuxInputType2 returns the AuxInputType2 field if non-nil, zero value otherwise.

### GetAuxInputType2Ok

`func (o *UpdateVehicleRequest) GetAuxInputType2Ok() (*VehicleAuxInputType, bool)`

GetAuxInputType2Ok returns a tuple with the AuxInputType2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxInputType2

`func (o *UpdateVehicleRequest) SetAuxInputType2(v VehicleAuxInputType)`

SetAuxInputType2 sets AuxInputType2 field to given value.

### HasAuxInputType2

`func (o *UpdateVehicleRequest) HasAuxInputType2() bool`

HasAuxInputType2 returns a boolean if a field has been set.

### GetEngineHours

`func (o *UpdateVehicleRequest) GetEngineHours() int64`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *UpdateVehicleRequest) GetEngineHoursOk() (*int64, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *UpdateVehicleRequest) SetEngineHours(v int64)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *UpdateVehicleRequest) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### GetExternalIds

`func (o *UpdateVehicleRequest) GetExternalIds() map[string]string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *UpdateVehicleRequest) GetExternalIdsOk() (*map[string]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *UpdateVehicleRequest) SetExternalIds(v map[string]string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *UpdateVehicleRequest) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetHarshAccelerationSettingType

`func (o *UpdateVehicleRequest) GetHarshAccelerationSettingType() VehicleHarshAccelerationSettingType`

GetHarshAccelerationSettingType returns the HarshAccelerationSettingType field if non-nil, zero value otherwise.

### GetHarshAccelerationSettingTypeOk

`func (o *UpdateVehicleRequest) GetHarshAccelerationSettingTypeOk() (*VehicleHarshAccelerationSettingType, bool)`

GetHarshAccelerationSettingTypeOk returns a tuple with the HarshAccelerationSettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarshAccelerationSettingType

`func (o *UpdateVehicleRequest) SetHarshAccelerationSettingType(v VehicleHarshAccelerationSettingType)`

SetHarshAccelerationSettingType sets HarshAccelerationSettingType field to given value.

### HasHarshAccelerationSettingType

`func (o *UpdateVehicleRequest) HasHarshAccelerationSettingType() bool`

HasHarshAccelerationSettingType returns a boolean if a field has been set.

### GetLicensePlate

`func (o *UpdateVehicleRequest) GetLicensePlate() string`

GetLicensePlate returns the LicensePlate field if non-nil, zero value otherwise.

### GetLicensePlateOk

`func (o *UpdateVehicleRequest) GetLicensePlateOk() (*string, bool)`

GetLicensePlateOk returns a tuple with the LicensePlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlate

`func (o *UpdateVehicleRequest) SetLicensePlate(v string)`

SetLicensePlate sets LicensePlate field to given value.

### HasLicensePlate

`func (o *UpdateVehicleRequest) HasLicensePlate() bool`

HasLicensePlate returns a boolean if a field has been set.

### GetName

`func (o *UpdateVehicleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVehicleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVehicleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVehicleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateVehicleRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateVehicleRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateVehicleRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateVehicleRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOdometerMeters

`func (o *UpdateVehicleRequest) GetOdometerMeters() int64`

GetOdometerMeters returns the OdometerMeters field if non-nil, zero value otherwise.

### GetOdometerMetersOk

`func (o *UpdateVehicleRequest) GetOdometerMetersOk() (*int64, bool)`

GetOdometerMetersOk returns a tuple with the OdometerMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerMeters

`func (o *UpdateVehicleRequest) SetOdometerMeters(v int64)`

SetOdometerMeters sets OdometerMeters field to given value.

### HasOdometerMeters

`func (o *UpdateVehicleRequest) HasOdometerMeters() bool`

HasOdometerMeters returns a boolean if a field has been set.

### GetStaticAssignedDriverId

`func (o *UpdateVehicleRequest) GetStaticAssignedDriverId() string`

GetStaticAssignedDriverId returns the StaticAssignedDriverId field if non-nil, zero value otherwise.

### GetStaticAssignedDriverIdOk

`func (o *UpdateVehicleRequest) GetStaticAssignedDriverIdOk() (*string, bool)`

GetStaticAssignedDriverIdOk returns a tuple with the StaticAssignedDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAssignedDriverId

`func (o *UpdateVehicleRequest) SetStaticAssignedDriverId(v string)`

SetStaticAssignedDriverId sets StaticAssignedDriverId field to given value.

### HasStaticAssignedDriverId

`func (o *UpdateVehicleRequest) HasStaticAssignedDriverId() bool`

HasStaticAssignedDriverId returns a boolean if a field has been set.

### GetTagIds

`func (o *UpdateVehicleRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *UpdateVehicleRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *UpdateVehicleRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *UpdateVehicleRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetVin

`func (o *UpdateVehicleRequest) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *UpdateVehicleRequest) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *UpdateVehicleRequest) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *UpdateVehicleRequest) HasVin() bool`

HasVin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


