# UserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID for the role. | [optional] 
**Name** | Pointer to **string** | The name of the role. | [optional] 

## Methods

### NewUserRole

`func NewUserRole() *UserRole`

NewUserRole instantiates a new UserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleWithDefaults

`func NewUserRoleWithDefaults() *UserRole`

NewUserRoleWithDefaults instantiates a new UserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRole) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


