# VehicleLocationsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VehicleLocationsListResponseDataInner**](VehicleLocationsListResponseDataInner.md) | A list of vehicles and an array of location events for each vehicle. | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewVehicleLocationsListResponse

`func NewVehicleLocationsListResponse(data []VehicleLocationsListResponseDataInner, pagination PaginationResponse, ) *VehicleLocationsListResponse`

NewVehicleLocationsListResponse instantiates a new VehicleLocationsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLocationsListResponseWithDefaults

`func NewVehicleLocationsListResponseWithDefaults() *VehicleLocationsListResponse`

NewVehicleLocationsListResponseWithDefaults instantiates a new VehicleLocationsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VehicleLocationsListResponse) GetData() []VehicleLocationsListResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VehicleLocationsListResponse) GetDataOk() (*[]VehicleLocationsListResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VehicleLocationsListResponse) SetData(v []VehicleLocationsListResponseDataInner)`

SetData sets Data field to given value.


### GetPagination

`func (o *VehicleLocationsListResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VehicleLocationsListResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VehicleLocationsListResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


