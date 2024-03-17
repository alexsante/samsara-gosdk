# ListContactsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Contact**](Contact.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListContactsResponse

`func NewListContactsResponse() *ListContactsResponse`

NewListContactsResponse instantiates a new ListContactsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContactsResponseWithDefaults

`func NewListContactsResponseWithDefaults() *ListContactsResponse`

NewListContactsResponseWithDefaults instantiates a new ListContactsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListContactsResponse) GetData() []Contact`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListContactsResponse) GetDataOk() (*[]Contact, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListContactsResponse) SetData(v []Contact)`

SetData sets Data field to given value.

### HasData

`func (o *ListContactsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListContactsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListContactsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListContactsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListContactsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


