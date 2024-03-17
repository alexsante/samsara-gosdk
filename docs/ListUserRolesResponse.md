# ListUserRolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]UserRole**](UserRole.md) | A list of user roles | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListUserRolesResponse

`func NewListUserRolesResponse() *ListUserRolesResponse`

NewListUserRolesResponse instantiates a new ListUserRolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserRolesResponseWithDefaults

`func NewListUserRolesResponseWithDefaults() *ListUserRolesResponse`

NewListUserRolesResponseWithDefaults instantiates a new ListUserRolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListUserRolesResponse) GetData() []UserRole`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListUserRolesResponse) GetDataOk() (*[]UserRole, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListUserRolesResponse) SetData(v []UserRole)`

SetData sets Data field to given value.

### HasData

`func (o *ListUserRolesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListUserRolesResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUserRolesResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUserRolesResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListUserRolesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


