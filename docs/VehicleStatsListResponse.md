# VehicleStatsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VehicleStatsListResponseDataInner**](VehicleStatsListResponseDataInner.md) | A list of vehicles and an array of stat events for each vehicle. | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewVehicleStatsListResponse

`func NewVehicleStatsListResponse(data []VehicleStatsListResponseDataInner, pagination PaginationResponse, ) *VehicleStatsListResponse`

NewVehicleStatsListResponse instantiates a new VehicleStatsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStatsListResponseWithDefaults

`func NewVehicleStatsListResponseWithDefaults() *VehicleStatsListResponse`

NewVehicleStatsListResponseWithDefaults instantiates a new VehicleStatsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VehicleStatsListResponse) GetData() []VehicleStatsListResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VehicleStatsListResponse) GetDataOk() (*[]VehicleStatsListResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VehicleStatsListResponse) SetData(v []VehicleStatsListResponseDataInner)`

SetData sets Data field to given value.


### GetPagination

`func (o *VehicleStatsListResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VehicleStatsListResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VehicleStatsListResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


