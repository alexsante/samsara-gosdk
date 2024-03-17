# DriverVehicleGroupTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the tag. | [optional] 
**Name** | Pointer to **string** | Name of the tag. | [optional] 
**ParentTagId** | Pointer to **string** | If this tag is part a hierarchical tag tree, this is the ID of the parent tag, otherwise this will be omitted. | [optional] 

## Methods

### NewDriverVehicleGroupTag

`func NewDriverVehicleGroupTag() *DriverVehicleGroupTag`

NewDriverVehicleGroupTag instantiates a new DriverVehicleGroupTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverVehicleGroupTagWithDefaults

`func NewDriverVehicleGroupTagWithDefaults() *DriverVehicleGroupTag`

NewDriverVehicleGroupTagWithDefaults instantiates a new DriverVehicleGroupTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverVehicleGroupTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverVehicleGroupTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverVehicleGroupTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DriverVehicleGroupTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DriverVehicleGroupTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DriverVehicleGroupTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DriverVehicleGroupTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DriverVehicleGroupTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentTagId

`func (o *DriverVehicleGroupTag) GetParentTagId() string`

GetParentTagId returns the ParentTagId field if non-nil, zero value otherwise.

### GetParentTagIdOk

`func (o *DriverVehicleGroupTag) GetParentTagIdOk() (*string, bool)`

GetParentTagIdOk returns a tuple with the ParentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTagId

`func (o *DriverVehicleGroupTag) SetParentTagId(v string)`

SetParentTagId sets ParentTagId field to given value.

### HasParentTagId

`func (o *DriverVehicleGroupTag) HasParentTagId() bool`

HasParentTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


