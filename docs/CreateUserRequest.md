# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | [**UserAuthType**](UserAuthType.md) |  | 
**Email** | **string** | The email address of this user. | 
**Name** | **string** | The first and last name of the user. | 
**Roles** | [**[]UserRoleAssignmentRequest**](UserRoleAssignmentRequest.md) | The list of roles that applies to this user. A user may have \&quot;organizational\&quot; roles, which apply to the user at the organizational level, and \&quot;tag-specific\&quot; roles, which apply to the user for a given tag. | 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(authType UserAuthType, email string, name string, roles []UserRoleAssignmentRequest, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *CreateUserRequest) GetAuthType() UserAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateUserRequest) GetAuthTypeOk() (*UserAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateUserRequest) SetAuthType(v UserAuthType)`

SetAuthType sets AuthType field to given value.


### GetEmail

`func (o *CreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CreateUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRoles

`func (o *CreateUserRequest) GetRoles() []UserRoleAssignmentRequest`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserRequest) GetRolesOk() (*[]UserRoleAssignmentRequest, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserRequest) SetRoles(v []UserRoleAssignmentRequest)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


