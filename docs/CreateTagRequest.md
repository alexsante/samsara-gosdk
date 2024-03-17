# CreateTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | The addresses that belong to this tag. | [optional] 
**Assets** | Pointer to **[]string** | The trailers, unpowered, and powered assets that belong to this tag. | [optional] 
**Drivers** | Pointer to **[]string** | The drivers that belong to this tag. | [optional] 
**Machines** | Pointer to **[]string** | The machines that belong to this tag. | [optional] 
**Name** | **string** | Name of this tag. | 
**ParentTagId** | Pointer to **string** | If this tag is part a hierarchical tag tree, this is the ID of the parent tag, otherwise this will be omitted. | [optional] 
**Sensors** | Pointer to **[]string** | The sensors that belong to this tag. | [optional] 
**Vehicles** | Pointer to **[]string** | The vehicles that belong to this tag. | [optional] 

## Methods

### NewCreateTagRequest

`func NewCreateTagRequest(name string, ) *CreateTagRequest`

NewCreateTagRequest instantiates a new CreateTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagRequestWithDefaults

`func NewCreateTagRequestWithDefaults() *CreateTagRequest`

NewCreateTagRequestWithDefaults instantiates a new CreateTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *CreateTagRequest) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CreateTagRequest) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CreateTagRequest) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CreateTagRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAssets

`func (o *CreateTagRequest) GetAssets() []string`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CreateTagRequest) GetAssetsOk() (*[]string, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CreateTagRequest) SetAssets(v []string)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CreateTagRequest) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetDrivers

`func (o *CreateTagRequest) GetDrivers() []string`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *CreateTagRequest) GetDriversOk() (*[]string, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *CreateTagRequest) SetDrivers(v []string)`

SetDrivers sets Drivers field to given value.

### HasDrivers

`func (o *CreateTagRequest) HasDrivers() bool`

HasDrivers returns a boolean if a field has been set.

### GetMachines

`func (o *CreateTagRequest) GetMachines() []string`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *CreateTagRequest) GetMachinesOk() (*[]string, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *CreateTagRequest) SetMachines(v []string)`

SetMachines sets Machines field to given value.

### HasMachines

`func (o *CreateTagRequest) HasMachines() bool`

HasMachines returns a boolean if a field has been set.

### GetName

`func (o *CreateTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTagRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentTagId

`func (o *CreateTagRequest) GetParentTagId() string`

GetParentTagId returns the ParentTagId field if non-nil, zero value otherwise.

### GetParentTagIdOk

`func (o *CreateTagRequest) GetParentTagIdOk() (*string, bool)`

GetParentTagIdOk returns a tuple with the ParentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTagId

`func (o *CreateTagRequest) SetParentTagId(v string)`

SetParentTagId sets ParentTagId field to given value.

### HasParentTagId

`func (o *CreateTagRequest) HasParentTagId() bool`

HasParentTagId returns a boolean if a field has been set.

### GetSensors

`func (o *CreateTagRequest) GetSensors() []string`

GetSensors returns the Sensors field if non-nil, zero value otherwise.

### GetSensorsOk

`func (o *CreateTagRequest) GetSensorsOk() (*[]string, bool)`

GetSensorsOk returns a tuple with the Sensors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensors

`func (o *CreateTagRequest) SetSensors(v []string)`

SetSensors sets Sensors field to given value.

### HasSensors

`func (o *CreateTagRequest) HasSensors() bool`

HasSensors returns a boolean if a field has been set.

### GetVehicles

`func (o *CreateTagRequest) GetVehicles() []string`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *CreateTagRequest) GetVehiclesOk() (*[]string, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *CreateTagRequest) SetVehicles(v []string)`

SetVehicles sets Vehicles field to given value.

### HasVehicles

`func (o *CreateTagRequest) HasVehicles() bool`

HasVehicles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


