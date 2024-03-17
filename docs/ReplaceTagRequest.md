# ReplaceTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | The addresses that belong to this tag. | [optional] 
**Assets** | Pointer to **[]string** | The trailers, unpowered, and powered assets that belong to this tag. | [optional] 
**Drivers** | Pointer to **[]string** | The drivers that belong to this tag. | [optional] 
**Machines** | Pointer to **[]string** | The machines that belong to this tag. | [optional] 
**Name** | Pointer to **string** | Name of this tag. | [optional] 
**ParentTagId** | Pointer to **string** | If this tag is part a hierarchical tag tree, this is the ID of the parent tag, otherwise this will be omitted. | [optional] 
**Sensors** | Pointer to **[]string** | The sensors that belong to this tag. | [optional] 
**Vehicles** | Pointer to **[]string** | The vehicles that belong to this tag. | [optional] 

## Methods

### NewReplaceTagRequest

`func NewReplaceTagRequest() *ReplaceTagRequest`

NewReplaceTagRequest instantiates a new ReplaceTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceTagRequestWithDefaults

`func NewReplaceTagRequestWithDefaults() *ReplaceTagRequest`

NewReplaceTagRequestWithDefaults instantiates a new ReplaceTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ReplaceTagRequest) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ReplaceTagRequest) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ReplaceTagRequest) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ReplaceTagRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAssets

`func (o *ReplaceTagRequest) GetAssets() []string`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ReplaceTagRequest) GetAssetsOk() (*[]string, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ReplaceTagRequest) SetAssets(v []string)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ReplaceTagRequest) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetDrivers

`func (o *ReplaceTagRequest) GetDrivers() []string`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *ReplaceTagRequest) GetDriversOk() (*[]string, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *ReplaceTagRequest) SetDrivers(v []string)`

SetDrivers sets Drivers field to given value.

### HasDrivers

`func (o *ReplaceTagRequest) HasDrivers() bool`

HasDrivers returns a boolean if a field has been set.

### GetMachines

`func (o *ReplaceTagRequest) GetMachines() []string`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *ReplaceTagRequest) GetMachinesOk() (*[]string, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *ReplaceTagRequest) SetMachines(v []string)`

SetMachines sets Machines field to given value.

### HasMachines

`func (o *ReplaceTagRequest) HasMachines() bool`

HasMachines returns a boolean if a field has been set.

### GetName

`func (o *ReplaceTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplaceTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplaceTagRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReplaceTagRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentTagId

`func (o *ReplaceTagRequest) GetParentTagId() string`

GetParentTagId returns the ParentTagId field if non-nil, zero value otherwise.

### GetParentTagIdOk

`func (o *ReplaceTagRequest) GetParentTagIdOk() (*string, bool)`

GetParentTagIdOk returns a tuple with the ParentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTagId

`func (o *ReplaceTagRequest) SetParentTagId(v string)`

SetParentTagId sets ParentTagId field to given value.

### HasParentTagId

`func (o *ReplaceTagRequest) HasParentTagId() bool`

HasParentTagId returns a boolean if a field has been set.

### GetSensors

`func (o *ReplaceTagRequest) GetSensors() []string`

GetSensors returns the Sensors field if non-nil, zero value otherwise.

### GetSensorsOk

`func (o *ReplaceTagRequest) GetSensorsOk() (*[]string, bool)`

GetSensorsOk returns a tuple with the Sensors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensors

`func (o *ReplaceTagRequest) SetSensors(v []string)`

SetSensors sets Sensors field to given value.

### HasSensors

`func (o *ReplaceTagRequest) HasSensors() bool`

HasSensors returns a boolean if a field has been set.

### GetVehicles

`func (o *ReplaceTagRequest) GetVehicles() []string`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *ReplaceTagRequest) GetVehiclesOk() (*[]string, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *ReplaceTagRequest) SetVehicles(v []string)`

SetVehicles sets Vehicles field to given value.

### HasVehicles

`func (o *ReplaceTagRequest) HasVehicles() bool`

HasVehicles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


