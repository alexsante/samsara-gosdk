# VehicleStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VehicleStatsResponseDataInner**](VehicleStatsResponseDataInner.md) | List of the most recent stats for the specified vehicles and stat types. | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewVehicleStatsResponse

`func NewVehicleStatsResponse(data []VehicleStatsResponseDataInner, pagination PaginationResponse, ) *VehicleStatsResponse`

NewVehicleStatsResponse instantiates a new VehicleStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStatsResponseWithDefaults

`func NewVehicleStatsResponseWithDefaults() *VehicleStatsResponse`

NewVehicleStatsResponseWithDefaults instantiates a new VehicleStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VehicleStatsResponse) GetData() []VehicleStatsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VehicleStatsResponse) GetDataOk() (*[]VehicleStatsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VehicleStatsResponse) SetData(v []VehicleStatsResponseDataInner)`

SetData sets Data field to given value.


### GetPagination

`func (o *VehicleStatsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VehicleStatsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VehicleStatsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


