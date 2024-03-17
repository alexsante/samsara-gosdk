# PaginationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndCursor** | **string** | Cursor identifier representing the last element in the response. This value should be used in conjunction with a subsequent request&#39;s &#39;after&#39; query parameter. This may be an empty string if there are no more pages left to view. | 
**HasNextPage** | **bool** | True if there are more pages of results immediately available after this endCursor. | 

## Methods

### NewPaginationResponse

`func NewPaginationResponse(endCursor string, hasNextPage bool, ) *PaginationResponse`

NewPaginationResponse instantiates a new PaginationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResponseWithDefaults

`func NewPaginationResponseWithDefaults() *PaginationResponse`

NewPaginationResponseWithDefaults instantiates a new PaginationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndCursor

`func (o *PaginationResponse) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *PaginationResponse) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *PaginationResponse) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.


### GetHasNextPage

`func (o *PaginationResponse) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *PaginationResponse) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *PaginationResponse) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


