# DriverCarrierSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierName** | Pointer to **string** | Carrier for a given driver. | [optional] 
**DotNumber** | Pointer to **int64** | Carrier US DOT Number. If this differs from the general organization&#39;s settings, the override value is used. Updating this value only updates the override setting for this driver. | [optional] 
**MainOfficeAddress** | Pointer to **string** | Main office address for a given driver. If this differs from the general organization&#39;s settings, the override value is used.  | [optional] 

## Methods

### NewDriverCarrierSettings

`func NewDriverCarrierSettings() *DriverCarrierSettings`

NewDriverCarrierSettings instantiates a new DriverCarrierSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverCarrierSettingsWithDefaults

`func NewDriverCarrierSettingsWithDefaults() *DriverCarrierSettings`

NewDriverCarrierSettingsWithDefaults instantiates a new DriverCarrierSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierName

`func (o *DriverCarrierSettings) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *DriverCarrierSettings) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *DriverCarrierSettings) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *DriverCarrierSettings) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetDotNumber

`func (o *DriverCarrierSettings) GetDotNumber() int64`

GetDotNumber returns the DotNumber field if non-nil, zero value otherwise.

### GetDotNumberOk

`func (o *DriverCarrierSettings) GetDotNumberOk() (*int64, bool)`

GetDotNumberOk returns a tuple with the DotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotNumber

`func (o *DriverCarrierSettings) SetDotNumber(v int64)`

SetDotNumber sets DotNumber field to given value.

### HasDotNumber

`func (o *DriverCarrierSettings) HasDotNumber() bool`

HasDotNumber returns a boolean if a field has been set.

### GetMainOfficeAddress

`func (o *DriverCarrierSettings) GetMainOfficeAddress() string`

GetMainOfficeAddress returns the MainOfficeAddress field if non-nil, zero value otherwise.

### GetMainOfficeAddressOk

`func (o *DriverCarrierSettings) GetMainOfficeAddressOk() (*string, bool)`

GetMainOfficeAddressOk returns a tuple with the MainOfficeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainOfficeAddress

`func (o *DriverCarrierSettings) SetMainOfficeAddress(v string)`

SetMainOfficeAddress sets MainOfficeAddress field to given value.

### HasMainOfficeAddress

`func (o *DriverCarrierSettings) HasMainOfficeAddress() bool`

HasMainOfficeAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


