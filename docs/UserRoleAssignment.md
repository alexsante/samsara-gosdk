# UserRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**UserRole**](UserRole.md) |  | [optional] 
**Tag** | Pointer to [**UserRoleAssignmentTag**](UserRoleAssignmentTag.md) |  | [optional] 

## Methods

### NewUserRoleAssignment

`func NewUserRoleAssignment() *UserRoleAssignment`

NewUserRoleAssignment instantiates a new UserRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleAssignmentWithDefaults

`func NewUserRoleAssignmentWithDefaults() *UserRoleAssignment`

NewUserRoleAssignmentWithDefaults instantiates a new UserRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *UserRoleAssignment) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserRoleAssignment) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserRoleAssignment) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserRoleAssignment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTag

`func (o *UserRoleAssignment) GetTag() UserRoleAssignmentTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UserRoleAssignment) GetTagOk() (*UserRoleAssignmentTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UserRoleAssignment) SetTag(v UserRoleAssignmentTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *UserRoleAssignment) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


