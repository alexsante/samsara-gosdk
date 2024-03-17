# Driver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierSettings** | Pointer to [**DriverCarrierSettings**](DriverCarrierSettings.md) |  | [optional] 
**CreatedAtTime** | Pointer to **time.Time** | The date and time this driver was created in RFC 3339 format. | [optional] 
**EldAdverseWeatherExemptionEnabled** | Pointer to **bool** | Flag indicating this driver may use Adverse Weather exemptions in ELD logs. | [optional] [default to false]
**EldBigDayExemptionEnabled** | Pointer to **bool** | Flag indicating this driver may use Big Day exemption in ELD logs. | [optional] [default to false]
**EldDayStartHour** | Pointer to **int32** | &#x60;0&#x60; indicating midnight-to-midnight ELD driving hours, &#x60;12&#x60; to indicate noon-to-noon driving hours. | [optional] [default to 0]
**EldExempt** | Pointer to **bool** | Flag indicating this driver is exempt from the Electronic Logging Mandate. | [optional] [default to false]
**EldExemptReason** | Pointer to **string** | Reason that this driver is exempt from the Electronic Logging Mandate (see eldExempt). | [optional] 
**EldPcEnabled** | Pointer to **bool** | Flag indicating this driver may select the Personal Conveyance duty status in ELD logs. | [optional] [default to false]
**EldYmEnabled** | Pointer to **bool** | Flag indicating this driver may select the Yard Move duty status in ELD logs. | [optional] [default to false]
**ExternalIds** | Pointer to **map[string]string** | The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object. | [optional] 
**Id** | Pointer to **string** | Samsara ID for the driver. | [optional] 
**IsDeactivated** | Pointer to **bool** | A boolean indicating whether or not the driver is deactivated. | [optional] 
**LicenseNumber** | Pointer to **string** | Driver&#39;s state issued license number. The combination of this number and &#x60;licenseState&#x60; must be unique. | [optional] 
**LicenseState** | Pointer to **string** | Abbreviation of state that issued driver&#39;s license. | [optional] 
**Locale** | Pointer to [**DriverLocale**](DriverLocale.md) |  | [optional] 
**Name** | Pointer to **string** | Driver&#39;s name. | [optional] 
**Notes** | Pointer to **string** | Notes about the driver. | [optional] 
**Phone** | Pointer to **string** | Phone number of the driver. | [optional] 
**StaticAssignedVehicle** | Pointer to [**DriverStaticAssignedVehicle**](DriverStaticAssignedVehicle.md) |  | [optional] 
**TachographCardNumber** | Pointer to **string** | Driver&#39;s assigned tachograph card number (Europe specific) | [optional] 
**Tags** | Pointer to [**[]TagTinyResponse**](TagTinyResponse.md) | The tags this driver belongs to. | [optional] 
**Timezone** | Pointer to **string** | Home terminal timezone, in order to indicate what time zone should be used to calculate the ELD logs. Driver timezones use [IANA timezone database](https://www.iana.org/time-zones) keys (e.g. &#x60;America/Los_Angeles&#x60;, &#x60;America/New_York&#x60;, &#x60;Europe/London&#x60;, etc.). You can find a mapping of common timezone formats to IANA timezone keys [here](https://unicode.org/cldr/charts/latest/supplemental/zone_tzid.html). | [optional] [default to "America/Los_Angeles"]
**UpdatedAtTime** | Pointer to **time.Time** | The date and time this driver was last updated in RFC 3339 format. | [optional] 
**Username** | Pointer to **string** | Driver&#39;s login username into the driver app. The username may not contain spaces or the &#39;@&#39; symbol. The username must be unique. | [optional] 
**VehicleGroupTag** | Pointer to [**DriverVehicleGroupTag**](DriverVehicleGroupTag.md) |  | [optional] 

## Methods

### NewDriver

`func NewDriver() *Driver`

NewDriver instantiates a new Driver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverWithDefaults

`func NewDriverWithDefaults() *Driver`

NewDriverWithDefaults instantiates a new Driver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierSettings

`func (o *Driver) GetCarrierSettings() DriverCarrierSettings`

GetCarrierSettings returns the CarrierSettings field if non-nil, zero value otherwise.

### GetCarrierSettingsOk

`func (o *Driver) GetCarrierSettingsOk() (*DriverCarrierSettings, bool)`

GetCarrierSettingsOk returns a tuple with the CarrierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSettings

`func (o *Driver) SetCarrierSettings(v DriverCarrierSettings)`

SetCarrierSettings sets CarrierSettings field to given value.

### HasCarrierSettings

`func (o *Driver) HasCarrierSettings() bool`

HasCarrierSettings returns a boolean if a field has been set.

### GetCreatedAtTime

`func (o *Driver) GetCreatedAtTime() time.Time`

GetCreatedAtTime returns the CreatedAtTime field if non-nil, zero value otherwise.

### GetCreatedAtTimeOk

`func (o *Driver) GetCreatedAtTimeOk() (*time.Time, bool)`

GetCreatedAtTimeOk returns a tuple with the CreatedAtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTime

`func (o *Driver) SetCreatedAtTime(v time.Time)`

SetCreatedAtTime sets CreatedAtTime field to given value.

### HasCreatedAtTime

`func (o *Driver) HasCreatedAtTime() bool`

HasCreatedAtTime returns a boolean if a field has been set.

### GetEldAdverseWeatherExemptionEnabled

`func (o *Driver) GetEldAdverseWeatherExemptionEnabled() bool`

GetEldAdverseWeatherExemptionEnabled returns the EldAdverseWeatherExemptionEnabled field if non-nil, zero value otherwise.

### GetEldAdverseWeatherExemptionEnabledOk

`func (o *Driver) GetEldAdverseWeatherExemptionEnabledOk() (*bool, bool)`

GetEldAdverseWeatherExemptionEnabledOk returns a tuple with the EldAdverseWeatherExemptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldAdverseWeatherExemptionEnabled

`func (o *Driver) SetEldAdverseWeatherExemptionEnabled(v bool)`

SetEldAdverseWeatherExemptionEnabled sets EldAdverseWeatherExemptionEnabled field to given value.

### HasEldAdverseWeatherExemptionEnabled

`func (o *Driver) HasEldAdverseWeatherExemptionEnabled() bool`

HasEldAdverseWeatherExemptionEnabled returns a boolean if a field has been set.

### GetEldBigDayExemptionEnabled

`func (o *Driver) GetEldBigDayExemptionEnabled() bool`

GetEldBigDayExemptionEnabled returns the EldBigDayExemptionEnabled field if non-nil, zero value otherwise.

### GetEldBigDayExemptionEnabledOk

`func (o *Driver) GetEldBigDayExemptionEnabledOk() (*bool, bool)`

GetEldBigDayExemptionEnabledOk returns a tuple with the EldBigDayExemptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldBigDayExemptionEnabled

`func (o *Driver) SetEldBigDayExemptionEnabled(v bool)`

SetEldBigDayExemptionEnabled sets EldBigDayExemptionEnabled field to given value.

### HasEldBigDayExemptionEnabled

`func (o *Driver) HasEldBigDayExemptionEnabled() bool`

HasEldBigDayExemptionEnabled returns a boolean if a field has been set.

### GetEldDayStartHour

`func (o *Driver) GetEldDayStartHour() int32`

GetEldDayStartHour returns the EldDayStartHour field if non-nil, zero value otherwise.

### GetEldDayStartHourOk

`func (o *Driver) GetEldDayStartHourOk() (*int32, bool)`

GetEldDayStartHourOk returns a tuple with the EldDayStartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldDayStartHour

`func (o *Driver) SetEldDayStartHour(v int32)`

SetEldDayStartHour sets EldDayStartHour field to given value.

### HasEldDayStartHour

`func (o *Driver) HasEldDayStartHour() bool`

HasEldDayStartHour returns a boolean if a field has been set.

### GetEldExempt

`func (o *Driver) GetEldExempt() bool`

GetEldExempt returns the EldExempt field if non-nil, zero value otherwise.

### GetEldExemptOk

`func (o *Driver) GetEldExemptOk() (*bool, bool)`

GetEldExemptOk returns a tuple with the EldExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldExempt

`func (o *Driver) SetEldExempt(v bool)`

SetEldExempt sets EldExempt field to given value.

### HasEldExempt

`func (o *Driver) HasEldExempt() bool`

HasEldExempt returns a boolean if a field has been set.

### GetEldExemptReason

`func (o *Driver) GetEldExemptReason() string`

GetEldExemptReason returns the EldExemptReason field if non-nil, zero value otherwise.

### GetEldExemptReasonOk

`func (o *Driver) GetEldExemptReasonOk() (*string, bool)`

GetEldExemptReasonOk returns a tuple with the EldExemptReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldExemptReason

`func (o *Driver) SetEldExemptReason(v string)`

SetEldExemptReason sets EldExemptReason field to given value.

### HasEldExemptReason

`func (o *Driver) HasEldExemptReason() bool`

HasEldExemptReason returns a boolean if a field has been set.

### GetEldPcEnabled

`func (o *Driver) GetEldPcEnabled() bool`

GetEldPcEnabled returns the EldPcEnabled field if non-nil, zero value otherwise.

### GetEldPcEnabledOk

`func (o *Driver) GetEldPcEnabledOk() (*bool, bool)`

GetEldPcEnabledOk returns a tuple with the EldPcEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldPcEnabled

`func (o *Driver) SetEldPcEnabled(v bool)`

SetEldPcEnabled sets EldPcEnabled field to given value.

### HasEldPcEnabled

`func (o *Driver) HasEldPcEnabled() bool`

HasEldPcEnabled returns a boolean if a field has been set.

### GetEldYmEnabled

`func (o *Driver) GetEldYmEnabled() bool`

GetEldYmEnabled returns the EldYmEnabled field if non-nil, zero value otherwise.

### GetEldYmEnabledOk

`func (o *Driver) GetEldYmEnabledOk() (*bool, bool)`

GetEldYmEnabledOk returns a tuple with the EldYmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldYmEnabled

`func (o *Driver) SetEldYmEnabled(v bool)`

SetEldYmEnabled sets EldYmEnabled field to given value.

### HasEldYmEnabled

`func (o *Driver) HasEldYmEnabled() bool`

HasEldYmEnabled returns a boolean if a field has been set.

### GetExternalIds

`func (o *Driver) GetExternalIds() map[string]string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *Driver) GetExternalIdsOk() (*map[string]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *Driver) SetExternalIds(v map[string]string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *Driver) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetId

`func (o *Driver) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Driver) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Driver) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Driver) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeactivated

`func (o *Driver) GetIsDeactivated() bool`

GetIsDeactivated returns the IsDeactivated field if non-nil, zero value otherwise.

### GetIsDeactivatedOk

`func (o *Driver) GetIsDeactivatedOk() (*bool, bool)`

GetIsDeactivatedOk returns a tuple with the IsDeactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeactivated

`func (o *Driver) SetIsDeactivated(v bool)`

SetIsDeactivated sets IsDeactivated field to given value.

### HasIsDeactivated

`func (o *Driver) HasIsDeactivated() bool`

HasIsDeactivated returns a boolean if a field has been set.

### GetLicenseNumber

`func (o *Driver) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *Driver) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *Driver) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *Driver) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### GetLicenseState

`func (o *Driver) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *Driver) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *Driver) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *Driver) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLocale

`func (o *Driver) GetLocale() DriverLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Driver) GetLocaleOk() (*DriverLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Driver) SetLocale(v DriverLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Driver) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *Driver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Driver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Driver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Driver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *Driver) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Driver) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Driver) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Driver) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPhone

`func (o *Driver) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Driver) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Driver) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Driver) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStaticAssignedVehicle

`func (o *Driver) GetStaticAssignedVehicle() DriverStaticAssignedVehicle`

GetStaticAssignedVehicle returns the StaticAssignedVehicle field if non-nil, zero value otherwise.

### GetStaticAssignedVehicleOk

`func (o *Driver) GetStaticAssignedVehicleOk() (*DriverStaticAssignedVehicle, bool)`

GetStaticAssignedVehicleOk returns a tuple with the StaticAssignedVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAssignedVehicle

`func (o *Driver) SetStaticAssignedVehicle(v DriverStaticAssignedVehicle)`

SetStaticAssignedVehicle sets StaticAssignedVehicle field to given value.

### HasStaticAssignedVehicle

`func (o *Driver) HasStaticAssignedVehicle() bool`

HasStaticAssignedVehicle returns a boolean if a field has been set.

### GetTachographCardNumber

`func (o *Driver) GetTachographCardNumber() string`

GetTachographCardNumber returns the TachographCardNumber field if non-nil, zero value otherwise.

### GetTachographCardNumberOk

`func (o *Driver) GetTachographCardNumberOk() (*string, bool)`

GetTachographCardNumberOk returns a tuple with the TachographCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTachographCardNumber

`func (o *Driver) SetTachographCardNumber(v string)`

SetTachographCardNumber sets TachographCardNumber field to given value.

### HasTachographCardNumber

`func (o *Driver) HasTachographCardNumber() bool`

HasTachographCardNumber returns a boolean if a field has been set.

### GetTags

`func (o *Driver) GetTags() []TagTinyResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Driver) GetTagsOk() (*[]TagTinyResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Driver) SetTags(v []TagTinyResponse)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Driver) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimezone

`func (o *Driver) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Driver) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Driver) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Driver) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUpdatedAtTime

`func (o *Driver) GetUpdatedAtTime() time.Time`

GetUpdatedAtTime returns the UpdatedAtTime field if non-nil, zero value otherwise.

### GetUpdatedAtTimeOk

`func (o *Driver) GetUpdatedAtTimeOk() (*time.Time, bool)`

GetUpdatedAtTimeOk returns a tuple with the UpdatedAtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTime

`func (o *Driver) SetUpdatedAtTime(v time.Time)`

SetUpdatedAtTime sets UpdatedAtTime field to given value.

### HasUpdatedAtTime

`func (o *Driver) HasUpdatedAtTime() bool`

HasUpdatedAtTime returns a boolean if a field has been set.

### GetUsername

`func (o *Driver) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Driver) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Driver) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Driver) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVehicleGroupTag

`func (o *Driver) GetVehicleGroupTag() DriverVehicleGroupTag`

GetVehicleGroupTag returns the VehicleGroupTag field if non-nil, zero value otherwise.

### GetVehicleGroupTagOk

`func (o *Driver) GetVehicleGroupTagOk() (*DriverVehicleGroupTag, bool)`

GetVehicleGroupTagOk returns a tuple with the VehicleGroupTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleGroupTag

`func (o *Driver) SetVehicleGroupTag(v DriverVehicleGroupTag)`

SetVehicleGroupTag sets VehicleGroupTag field to given value.

### HasVehicleGroupTag

`func (o *Driver) HasVehicleGroupTag() bool`

HasVehicleGroupTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


