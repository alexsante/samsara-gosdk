# UserRoleAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | Pointer to **string** | The unique ID for the role. | [optional] 
**TagId** | Pointer to **string** | ID of the tag this role applies to. | [optional] 

## Methods

### NewUserRoleAssignmentRequest

`func NewUserRoleAssignmentRequest() *UserRoleAssignmentRequest`

NewUserRoleAssignmentRequest instantiates a new UserRoleAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleAssignmentRequestWithDefaults

`func NewUserRoleAssignmentRequestWithDefaults() *UserRoleAssignmentRequest`

NewUserRoleAssignmentRequestWithDefaults instantiates a new UserRoleAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *UserRoleAssignmentRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UserRoleAssignmentRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UserRoleAssignmentRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UserRoleAssignmentRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetTagId

`func (o *UserRoleAssignmentRequest) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *UserRoleAssignmentRequest) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *UserRoleAssignmentRequest) SetTagId(v string)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *UserRoleAssignmentRequest) HasTagId() bool`

HasTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


