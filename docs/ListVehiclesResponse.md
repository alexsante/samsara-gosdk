# ListVehiclesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Vehicle**](Vehicle.md) |  | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewListVehiclesResponse

`func NewListVehiclesResponse(data []Vehicle, pagination PaginationResponse, ) *ListVehiclesResponse`

NewListVehiclesResponse instantiates a new ListVehiclesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVehiclesResponseWithDefaults

`func NewListVehiclesResponseWithDefaults() *ListVehiclesResponse`

NewListVehiclesResponseWithDefaults instantiates a new ListVehiclesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListVehiclesResponse) GetData() []Vehicle`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListVehiclesResponse) GetDataOk() (*[]Vehicle, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListVehiclesResponse) SetData(v []Vehicle)`

SetData sets Data field to given value.


### GetPagination

`func (o *ListVehiclesResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListVehiclesResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListVehiclesResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


