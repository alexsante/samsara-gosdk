# UpdateDriverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierSettings** | Pointer to [**DriverCarrierSettings**](DriverCarrierSettings.md) |  | [optional] 
**EldAdverseWeatherExemptionEnabled** | Pointer to **bool** | Flag indicating this driver may use Adverse Weather exemptions in ELD logs. | [optional] [default to false]
**EldBigDayExemptionEnabled** | Pointer to **bool** | Flag indicating this driver may use Big Day exemption in ELD logs. | [optional] [default to false]
**EldDayStartHour** | Pointer to **int32** | &#x60;0&#x60; indicating midnight-to-midnight ELD driving hours, &#x60;12&#x60; to indicate noon-to-noon driving hours. | [optional] [default to 0]
**EldExempt** | Pointer to **bool** | Flag indicating this driver is exempt from the Electronic Logging Mandate. | [optional] [default to false]
**EldExemptReason** | Pointer to **string** | Reason that this driver is exempt from the Electronic Logging Mandate (see eldExempt). | [optional] 
**EldPcEnabled** | Pointer to **bool** | Flag indicating this driver may select the Personal Conveyance duty status in ELD logs. | [optional] [default to false]
**EldYmEnabled** | Pointer to **bool** | Flag indicating this driver may select the Yard Move duty status in ELD logs. | [optional] [default to false]
**ExternalIds** | Pointer to **map[string]string** | The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object. | [optional] 
**LicenseNumber** | Pointer to **string** | Driver&#39;s state issued license number. The combination of this number and &#x60;licenseState&#x60; must be unique. | [optional] 
**LicenseState** | Pointer to **string** | Abbreviation of state that issued driver&#39;s license. | [optional] 
**Locale** | Pointer to [**DriverLocale**](DriverLocale.md) |  | [optional] 
**Name** | Pointer to **string** | Driver&#39;s name. | [optional] 
**Notes** | Pointer to **string** | Notes about the driver. | [optional] 
**Password** | Pointer to **string** | Password that the driver can use to login to the Samsara driver app. | [optional] 
**Phone** | Pointer to **string** | Phone number of the driver. | [optional] 
**StaticAssignedVehicleId** | Pointer to **string** | ID of vehicle that the driver is permanently assigned to. (uncommon). | [optional] 
**TachographCardNumber** | Pointer to **string** | Driver&#39;s assigned tachograph card number (Europe specific) | [optional] 
**TagIds** | Pointer to **[]string** | IDs of tags the driver is associated with. | [optional] 
**Timezone** | Pointer to **string** | Home terminal timezone, in order to indicate what time zone should be used to calculate the ELD logs. Driver timezones use [IANA timezone database](https://www.iana.org/time-zones) keys (e.g. &#x60;America/Los_Angeles&#x60;, &#x60;America/New_York&#x60;, &#x60;Europe/London&#x60;, etc.). You can find a mapping of common timezone formats to IANA timezone keys [here](https://unicode.org/cldr/charts/latest/supplemental/zone_tzid.html). | [optional] [default to "America/Los_Angeles"]
**Username** | Pointer to **string** | Driver&#39;s login username into the driver app. The username may not contain spaces or the &#39;@&#39; symbol. The username must be unique. | [optional] 
**VehicleGroupTagId** | Pointer to **string** | Tag ID which determines which vehicles a driver will see when selecting vehicles. | [optional] 

## Methods

### NewUpdateDriverRequest

`func NewUpdateDriverRequest() *UpdateDriverRequest`

NewUpdateDriverRequest instantiates a new UpdateDriverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDriverRequestWithDefaults

`func NewUpdateDriverRequestWithDefaults() *UpdateDriverRequest`

NewUpdateDriverRequestWithDefaults instantiates a new UpdateDriverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierSettings

`func (o *UpdateDriverRequest) GetCarrierSettings() DriverCarrierSettings`

GetCarrierSettings returns the CarrierSettings field if non-nil, zero value otherwise.

### GetCarrierSettingsOk

`func (o *UpdateDriverRequest) GetCarrierSettingsOk() (*DriverCarrierSettings, bool)`

GetCarrierSettingsOk returns a tuple with the CarrierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSettings

`func (o *UpdateDriverRequest) SetCarrierSettings(v DriverCarrierSettings)`

SetCarrierSettings sets CarrierSettings field to given value.

### HasCarrierSettings

`func (o *UpdateDriverRequest) HasCarrierSettings() bool`

HasCarrierSettings returns a boolean if a field has been set.

### GetEldAdverseWeatherExemptionEnabled

`func (o *UpdateDriverRequest) GetEldAdverseWeatherExemptionEnabled() bool`

GetEldAdverseWeatherExemptionEnabled returns the EldAdverseWeatherExemptionEnabled field if non-nil, zero value otherwise.

### GetEldAdverseWeatherExemptionEnabledOk

`func (o *UpdateDriverRequest) GetEldAdverseWeatherExemptionEnabledOk() (*bool, bool)`

GetEldAdverseWeatherExemptionEnabledOk returns a tuple with the EldAdverseWeatherExemptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldAdverseWeatherExemptionEnabled

`func (o *UpdateDriverRequest) SetEldAdverseWeatherExemptionEnabled(v bool)`

SetEldAdverseWeatherExemptionEnabled sets EldAdverseWeatherExemptionEnabled field to given value.

### HasEldAdverseWeatherExemptionEnabled

`func (o *UpdateDriverRequest) HasEldAdverseWeatherExemptionEnabled() bool`

HasEldAdverseWeatherExemptionEnabled returns a boolean if a field has been set.

### GetEldBigDayExemptionEnabled

`func (o *UpdateDriverRequest) GetEldBigDayExemptionEnabled() bool`

GetEldBigDayExemptionEnabled returns the EldBigDayExemptionEnabled field if non-nil, zero value otherwise.

### GetEldBigDayExemptionEnabledOk

`func (o *UpdateDriverRequest) GetEldBigDayExemptionEnabledOk() (*bool, bool)`

GetEldBigDayExemptionEnabledOk returns a tuple with the EldBigDayExemptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldBigDayExemptionEnabled

`func (o *UpdateDriverRequest) SetEldBigDayExemptionEnabled(v bool)`

SetEldBigDayExemptionEnabled sets EldBigDayExemptionEnabled field to given value.

### HasEldBigDayExemptionEnabled

`func (o *UpdateDriverRequest) HasEldBigDayExemptionEnabled() bool`

HasEldBigDayExemptionEnabled returns a boolean if a field has been set.

### GetEldDayStartHour

`func (o *UpdateDriverRequest) GetEldDayStartHour() int32`

GetEldDayStartHour returns the EldDayStartHour field if non-nil, zero value otherwise.

### GetEldDayStartHourOk

`func (o *UpdateDriverRequest) GetEldDayStartHourOk() (*int32, bool)`

GetEldDayStartHourOk returns a tuple with the EldDayStartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldDayStartHour

`func (o *UpdateDriverRequest) SetEldDayStartHour(v int32)`

SetEldDayStartHour sets EldDayStartHour field to given value.

### HasEldDayStartHour

`func (o *UpdateDriverRequest) HasEldDayStartHour() bool`

HasEldDayStartHour returns a boolean if a field has been set.

### GetEldExempt

`func (o *UpdateDriverRequest) GetEldExempt() bool`

GetEldExempt returns the EldExempt field if non-nil, zero value otherwise.

### GetEldExemptOk

`func (o *UpdateDriverRequest) GetEldExemptOk() (*bool, bool)`

GetEldExemptOk returns a tuple with the EldExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldExempt

`func (o *UpdateDriverRequest) SetEldExempt(v bool)`

SetEldExempt sets EldExempt field to given value.

### HasEldExempt

`func (o *UpdateDriverRequest) HasEldExempt() bool`

HasEldExempt returns a boolean if a field has been set.

### GetEldExemptReason

`func (o *UpdateDriverRequest) GetEldExemptReason() string`

GetEldExemptReason returns the EldExemptReason field if non-nil, zero value otherwise.

### GetEldExemptReasonOk

`func (o *UpdateDriverRequest) GetEldExemptReasonOk() (*string, bool)`

GetEldExemptReasonOk returns a tuple with the EldExemptReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldExemptReason

`func (o *UpdateDriverRequest) SetEldExemptReason(v string)`

SetEldExemptReason sets EldExemptReason field to given value.

### HasEldExemptReason

`func (o *UpdateDriverRequest) HasEldExemptReason() bool`

HasEldExemptReason returns a boolean if a field has been set.

### GetEldPcEnabled

`func (o *UpdateDriverRequest) GetEldPcEnabled() bool`

GetEldPcEnabled returns the EldPcEnabled field if non-nil, zero value otherwise.

### GetEldPcEnabledOk

`func (o *UpdateDriverRequest) GetEldPcEnabledOk() (*bool, bool)`

GetEldPcEnabledOk returns a tuple with the EldPcEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldPcEnabled

`func (o *UpdateDriverRequest) SetEldPcEnabled(v bool)`

SetEldPcEnabled sets EldPcEnabled field to given value.

### HasEldPcEnabled

`func (o *UpdateDriverRequest) HasEldPcEnabled() bool`

HasEldPcEnabled returns a boolean if a field has been set.

### GetEldYmEnabled

`func (o *UpdateDriverRequest) GetEldYmEnabled() bool`

GetEldYmEnabled returns the EldYmEnabled field if non-nil, zero value otherwise.

### GetEldYmEnabledOk

`func (o *UpdateDriverRequest) GetEldYmEnabledOk() (*bool, bool)`

GetEldYmEnabledOk returns a tuple with the EldYmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldYmEnabled

`func (o *UpdateDriverRequest) SetEldYmEnabled(v bool)`

SetEldYmEnabled sets EldYmEnabled field to given value.

### HasEldYmEnabled

`func (o *UpdateDriverRequest) HasEldYmEnabled() bool`

HasEldYmEnabled returns a boolean if a field has been set.

### GetExternalIds

`func (o *UpdateDriverRequest) GetExternalIds() map[string]string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *UpdateDriverRequest) GetExternalIdsOk() (*map[string]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *UpdateDriverRequest) SetExternalIds(v map[string]string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *UpdateDriverRequest) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetLicenseNumber

`func (o *UpdateDriverRequest) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *UpdateDriverRequest) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *UpdateDriverRequest) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *UpdateDriverRequest) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### GetLicenseState

`func (o *UpdateDriverRequest) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *UpdateDriverRequest) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *UpdateDriverRequest) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *UpdateDriverRequest) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLocale

`func (o *UpdateDriverRequest) GetLocale() DriverLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateDriverRequest) GetLocaleOk() (*DriverLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateDriverRequest) SetLocale(v DriverLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateDriverRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *UpdateDriverRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDriverRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDriverRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDriverRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateDriverRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateDriverRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateDriverRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateDriverRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateDriverRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateDriverRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateDriverRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateDriverRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateDriverRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateDriverRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateDriverRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateDriverRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStaticAssignedVehicleId

`func (o *UpdateDriverRequest) GetStaticAssignedVehicleId() string`

GetStaticAssignedVehicleId returns the StaticAssignedVehicleId field if non-nil, zero value otherwise.

### GetStaticAssignedVehicleIdOk

`func (o *UpdateDriverRequest) GetStaticAssignedVehicleIdOk() (*string, bool)`

GetStaticAssignedVehicleIdOk returns a tuple with the StaticAssignedVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAssignedVehicleId

`func (o *UpdateDriverRequest) SetStaticAssignedVehicleId(v string)`

SetStaticAssignedVehicleId sets StaticAssignedVehicleId field to given value.

### HasStaticAssignedVehicleId

`func (o *UpdateDriverRequest) HasStaticAssignedVehicleId() bool`

HasStaticAssignedVehicleId returns a boolean if a field has been set.

### GetTachographCardNumber

`func (o *UpdateDriverRequest) GetTachographCardNumber() string`

GetTachographCardNumber returns the TachographCardNumber field if non-nil, zero value otherwise.

### GetTachographCardNumberOk

`func (o *UpdateDriverRequest) GetTachographCardNumberOk() (*string, bool)`

GetTachographCardNumberOk returns a tuple with the TachographCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTachographCardNumber

`func (o *UpdateDriverRequest) SetTachographCardNumber(v string)`

SetTachographCardNumber sets TachographCardNumber field to given value.

### HasTachographCardNumber

`func (o *UpdateDriverRequest) HasTachographCardNumber() bool`

HasTachographCardNumber returns a boolean if a field has been set.

### GetTagIds

`func (o *UpdateDriverRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *UpdateDriverRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *UpdateDriverRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *UpdateDriverRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateDriverRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateDriverRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateDriverRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateDriverRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateDriverRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateDriverRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateDriverRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateDriverRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVehicleGroupTagId

`func (o *UpdateDriverRequest) GetVehicleGroupTagId() string`

GetVehicleGroupTagId returns the VehicleGroupTagId field if non-nil, zero value otherwise.

### GetVehicleGroupTagIdOk

`func (o *UpdateDriverRequest) GetVehicleGroupTagIdOk() (*string, bool)`

GetVehicleGroupTagIdOk returns a tuple with the VehicleGroupTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleGroupTagId

`func (o *UpdateDriverRequest) SetVehicleGroupTagId(v string)`

SetVehicleGroupTagId sets VehicleGroupTagId field to given value.

### HasVehicleGroupTagId

`func (o *UpdateDriverRequest) HasVehicleGroupTagId() bool`

HasVehicleGroupTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


