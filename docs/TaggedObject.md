# TaggedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The object ID. | 
**Name** | Pointer to **string** | The object name. | [optional] 

## Methods

### NewTaggedObject

`func NewTaggedObject(id string, ) *TaggedObject`

NewTaggedObject instantiates a new TaggedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedObjectWithDefaults

`func NewTaggedObjectWithDefaults() *TaggedObject`

NewTaggedObjectWithDefaults instantiates a new TaggedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaggedObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaggedObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaggedObject) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TaggedObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaggedObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaggedObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaggedObject) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


