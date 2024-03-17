# TagTinyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the tag. | [optional] 
**Name** | Pointer to **string** | Name of the tag. | [optional] 
**ParentTagId** | Pointer to **string** | If this tag is part a hierarchical tag tree, this is the ID of the parent tag, otherwise this will be omitted. | [optional] 

## Methods

### NewTagTinyResponse

`func NewTagTinyResponse() *TagTinyResponse`

NewTagTinyResponse instantiates a new TagTinyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTinyResponseWithDefaults

`func NewTagTinyResponseWithDefaults() *TagTinyResponse`

NewTagTinyResponseWithDefaults instantiates a new TagTinyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagTinyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagTinyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagTinyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TagTinyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TagTinyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagTinyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagTinyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagTinyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentTagId

`func (o *TagTinyResponse) GetParentTagId() string`

GetParentTagId returns the ParentTagId field if non-nil, zero value otherwise.

### GetParentTagIdOk

`func (o *TagTinyResponse) GetParentTagIdOk() (*string, bool)`

GetParentTagIdOk returns a tuple with the ParentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTagId

`func (o *TagTinyResponse) SetParentTagId(v string)`

SetParentTagId sets ParentTagId field to given value.

### HasParentTagId

`func (o *TagTinyResponse) HasParentTagId() bool`

HasParentTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


