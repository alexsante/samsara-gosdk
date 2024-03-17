# VehicleLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VehicleLocationsResponseDataInner**](VehicleLocationsResponseDataInner.md) | List of the most recent locations for the specified vehicles. | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewVehicleLocationsResponse

`func NewVehicleLocationsResponse(data []VehicleLocationsResponseDataInner, pagination PaginationResponse, ) *VehicleLocationsResponse`

NewVehicleLocationsResponse instantiates a new VehicleLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLocationsResponseWithDefaults

`func NewVehicleLocationsResponseWithDefaults() *VehicleLocationsResponse`

NewVehicleLocationsResponseWithDefaults instantiates a new VehicleLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VehicleLocationsResponse) GetData() []VehicleLocationsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VehicleLocationsResponse) GetDataOk() (*[]VehicleLocationsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VehicleLocationsResponse) SetData(v []VehicleLocationsResponseDataInner)`

SetData sets Data field to given value.


### GetPagination

`func (o *VehicleLocationsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VehicleLocationsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VehicleLocationsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


