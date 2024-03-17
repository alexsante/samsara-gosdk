# ListDriversResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Driver**](Driver.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListDriversResponse

`func NewListDriversResponse() *ListDriversResponse`

NewListDriversResponse instantiates a new ListDriversResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDriversResponseWithDefaults

`func NewListDriversResponseWithDefaults() *ListDriversResponse`

NewListDriversResponseWithDefaults instantiates a new ListDriversResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDriversResponse) GetData() []Driver`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDriversResponse) GetDataOk() (*[]Driver, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDriversResponse) SetData(v []Driver)`

SetData sets Data field to given value.

### HasData

`func (o *ListDriversResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListDriversResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDriversResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDriversResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListDriversResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


