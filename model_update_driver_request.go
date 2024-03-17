/*
Samsara API

Integrate your data with the Samsara API, customize the Samsara experience, and join a community of developers building with Samsara.

API version: 2019-12-12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateDriverRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDriverRequest{}

// UpdateDriverRequest Driver that should be updated.
type UpdateDriverRequest struct {
	CarrierSettings *DriverCarrierSettings `json:"carrierSettings,omitempty"`
	// Flag indicating this driver may use Adverse Weather exemptions in ELD logs.
	EldAdverseWeatherExemptionEnabled *bool `json:"eldAdverseWeatherExemptionEnabled,omitempty"`
	// Flag indicating this driver may use Big Day exemption in ELD logs.
	EldBigDayExemptionEnabled *bool `json:"eldBigDayExemptionEnabled,omitempty"`
	// `0` indicating midnight-to-midnight ELD driving hours, `12` to indicate noon-to-noon driving hours.
	EldDayStartHour *int32 `json:"eldDayStartHour,omitempty"`
	// Flag indicating this driver is exempt from the Electronic Logging Mandate.
	EldExempt *bool `json:"eldExempt,omitempty"`
	// Reason that this driver is exempt from the Electronic Logging Mandate (see eldExempt).
	EldExemptReason *string `json:"eldExemptReason,omitempty"`
	// Flag indicating this driver may select the Personal Conveyance duty status in ELD logs.
	EldPcEnabled *bool `json:"eldPcEnabled,omitempty"`
	// Flag indicating this driver may select the Yard Move duty status in ELD logs.
	EldYmEnabled *bool `json:"eldYmEnabled,omitempty"`
	// The [external IDs](https://developers.samsara.com/docs/external-ids) for the given object.
	ExternalIds *map[string]string `json:"externalIds,omitempty"`
	// Driver's state issued license number. The combination of this number and `licenseState` must be unique.
	LicenseNumber *string `json:"licenseNumber,omitempty"`
	// Abbreviation of state that issued driver's license.
	LicenseState *string `json:"licenseState,omitempty"`
	Locale *DriverLocale `json:"locale,omitempty"`
	// Driver's name.
	Name *string `json:"name,omitempty"`
	// Notes about the driver.
	Notes *string `json:"notes,omitempty"`
	// Password that the driver can use to login to the Samsara driver app.
	Password *string `json:"password,omitempty"`
	// Phone number of the driver.
	Phone *string `json:"phone,omitempty"`
	// ID of vehicle that the driver is permanently assigned to. (uncommon).
	StaticAssignedVehicleId *string `json:"staticAssignedVehicleId,omitempty"`
	// Driver's assigned tachograph card number (Europe specific)
	TachographCardNumber *string `json:"tachographCardNumber,omitempty"`
	// IDs of tags the driver is associated with.
	TagIds []string `json:"tagIds,omitempty"`
	// Home terminal timezone, in order to indicate what time zone should be used to calculate the ELD logs. Driver timezones use [IANA timezone database](https://www.iana.org/time-zones) keys (e.g. `America/Los_Angeles`, `America/New_York`, `Europe/London`, etc.). You can find a mapping of common timezone formats to IANA timezone keys [here](https://unicode.org/cldr/charts/latest/supplemental/zone_tzid.html).
	Timezone *string `json:"timezone,omitempty"`
	// Driver's login username into the driver app. The username may not contain spaces or the '@' symbol. The username must be unique.
	Username *string `json:"username,omitempty"`
	// Tag ID which determines which vehicles a driver will see when selecting vehicles.
	VehicleGroupTagId *string `json:"vehicleGroupTagId,omitempty"`
}

// NewUpdateDriverRequest instantiates a new UpdateDriverRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDriverRequest() *UpdateDriverRequest {
	this := UpdateDriverRequest{}
	var eldAdverseWeatherExemptionEnabled bool = false
	this.EldAdverseWeatherExemptionEnabled = &eldAdverseWeatherExemptionEnabled
	var eldBigDayExemptionEnabled bool = false
	this.EldBigDayExemptionEnabled = &eldBigDayExemptionEnabled
	var eldDayStartHour int32 = 0
	this.EldDayStartHour = &eldDayStartHour
	var eldExempt bool = false
	this.EldExempt = &eldExempt
	var eldPcEnabled bool = false
	this.EldPcEnabled = &eldPcEnabled
	var eldYmEnabled bool = false
	this.EldYmEnabled = &eldYmEnabled
	var timezone string = "America/Los_Angeles"
	this.Timezone = &timezone
	return &this
}

// NewUpdateDriverRequestWithDefaults instantiates a new UpdateDriverRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDriverRequestWithDefaults() *UpdateDriverRequest {
	this := UpdateDriverRequest{}
	var eldAdverseWeatherExemptionEnabled bool = false
	this.EldAdverseWeatherExemptionEnabled = &eldAdverseWeatherExemptionEnabled
	var eldBigDayExemptionEnabled bool = false
	this.EldBigDayExemptionEnabled = &eldBigDayExemptionEnabled
	var eldDayStartHour int32 = 0
	this.EldDayStartHour = &eldDayStartHour
	var eldExempt bool = false
	this.EldExempt = &eldExempt
	var eldPcEnabled bool = false
	this.EldPcEnabled = &eldPcEnabled
	var eldYmEnabled bool = false
	this.EldYmEnabled = &eldYmEnabled
	var timezone string = "America/Los_Angeles"
	this.Timezone = &timezone
	return &this
}

// GetCarrierSettings returns the CarrierSettings field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetCarrierSettings() DriverCarrierSettings {
	if o == nil || IsNil(o.CarrierSettings) {
		var ret DriverCarrierSettings
		return ret
	}
	return *o.CarrierSettings
}

// GetCarrierSettingsOk returns a tuple with the CarrierSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetCarrierSettingsOk() (*DriverCarrierSettings, bool) {
	if o == nil || IsNil(o.CarrierSettings) {
		return nil, false
	}
	return o.CarrierSettings, true
}

// HasCarrierSettings returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasCarrierSettings() bool {
	if o != nil && !IsNil(o.CarrierSettings) {
		return true
	}

	return false
}

// SetCarrierSettings gets a reference to the given DriverCarrierSettings and assigns it to the CarrierSettings field.
func (o *UpdateDriverRequest) SetCarrierSettings(v DriverCarrierSettings) {
	o.CarrierSettings = &v
}

// GetEldAdverseWeatherExemptionEnabled returns the EldAdverseWeatherExemptionEnabled field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetEldAdverseWeatherExemptionEnabled() bool {
	if o == nil || IsNil(o.EldAdverseWeatherExemptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EldAdverseWeatherExemptionEnabled
}

// GetEldAdverseWeatherExemptionEnabledOk returns a tuple with the EldAdverseWeatherExemptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetEldAdverseWeatherExemptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EldAdverseWeatherExemptionEnabled) {
		return nil, false
	}
	return o.EldAdverseWeatherExemptionEnabled, true
}

// HasEldAdverseWeatherExemptionEnabled returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasEldAdverseWeatherExemptionEnabled() bool {
	if o != nil && !IsNil(o.EldAdverseWeatherExemptionEnabled) {
		return true
	}

	return false
}

// SetEldAdverseWeatherExemptionEnabled gets a reference to the given bool and assigns it to the EldAdverseWeatherExemptionEnabled field.
func (o *UpdateDriverRequest) SetEldAdverseWeatherExemptionEnabled(v bool) {
	o.EldAdverseWeatherExemptionEnabled = &v
}

// GetEldBigDayExemptionEnabled returns the EldBigDayExemptionEnabled field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetEldBigDayExemptionEnabled() bool {
	if o == nil || IsNil(o.EldBigDayExemptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EldBigDayExemptionEnabled
}

// GetEldBigDayExemptionEnabledOk returns a tuple with the EldBigDayExemptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetEldBigDayExemptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EldBigDayExemptionEnabled) {
		return nil, false
	}
	return o.EldBigDayExemptionEnabled, true
}

// HasEldBigDayExemptionEnabled returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasEldBigDayExemptionEnabled() bool {
	if o != nil && !IsNil(o.EldBigDayExemptionEnabled) {
		return true
	}

	return false
}

// SetEldBigDayExemptionEnabled gets a reference to the given bool and assigns it to the EldBigDayExemptionEnabled field.
func (o *UpdateDriverRequest) SetEldBigDayExemptionEnabled(v bool) {
	o.EldBigDayExemptionEnabled = &v
}

// GetEldDayStartHour returns the EldDayStartHour field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetEldDayStartHour() int32 {
	if o == nil || IsNil(o.EldDayStartHour) {
		var ret int32
		return ret
	}
	return *o.EldDayStartHour
}

// GetEldDayStartHourOk returns a tuple with the EldDayStartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetEldDayStartHourOk() (*int32, bool) {
	if o == nil || IsNil(o.EldDayStartHour) {
		return nil, false
	}
	return o.EldDayStartHour, true
}

// HasEldDayStartHour returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasEldDayStartHour() bool {
	if o != nil && !IsNil(o.EldDayStartHour) {
		return true
	}

	return false
}

// SetEldDayStartHour gets a reference to the given int32 and assigns it to the EldDayStartHour field.
func (o *UpdateDriverRequest) SetEldDayStartHour(v int32) {
	o.EldDayStartHour = &v
}

// GetEldExempt returns the EldExempt field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetEldExempt() bool {
	if o == nil || IsNil(o.EldExempt) {
		var ret bool
		return ret
	}
	return *o.EldExempt
}

// GetEldExemptOk returns a tuple with the EldExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetEldExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.EldExempt) {
		return nil, false
	}
	return o.EldExempt, true
}

// HasEldExempt returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasEldExempt() bool {
	if o != nil && !IsNil(o.EldExempt) {
		return true
	}

	return false
}

// SetEldExempt gets a reference to the given bool and assigns it to the EldExempt field.
func (o *UpdateDriverRequest) SetEldExempt(v bool) {
	o.EldExempt = &v
}

// GetEldExemptReason returns the EldExemptReason field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetEldExemptReason() string {
	if o == nil || IsNil(o.EldExemptReason) {
		var ret string
		return ret
	}
	return *o.EldExemptReason
}

// GetEldExemptReasonOk returns a tuple with the EldExemptReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetEldExemptReasonOk() (*string, bool) {
	if o == nil || IsNil(o.EldExemptReason) {
		return nil, false
	}
	return o.EldExemptReason, true
}

// HasEldExemptReason returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasEldExemptReason() bool {
	if o != nil && !IsNil(o.EldExemptReason) {
		return true
	}

	return false
}

// SetEldExemptReason gets a reference to the given string and assigns it to the EldExemptReason field.
func (o *UpdateDriverRequest) SetEldExemptReason(v string) {
	o.EldExemptReason = &v
}

// GetEldPcEnabled returns the EldPcEnabled field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetEldPcEnabled() bool {
	if o == nil || IsNil(o.EldPcEnabled) {
		var ret bool
		return ret
	}
	return *o.EldPcEnabled
}

// GetEldPcEnabledOk returns a tuple with the EldPcEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetEldPcEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EldPcEnabled) {
		return nil, false
	}
	return o.EldPcEnabled, true
}

// HasEldPcEnabled returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasEldPcEnabled() bool {
	if o != nil && !IsNil(o.EldPcEnabled) {
		return true
	}

	return false
}

// SetEldPcEnabled gets a reference to the given bool and assigns it to the EldPcEnabled field.
func (o *UpdateDriverRequest) SetEldPcEnabled(v bool) {
	o.EldPcEnabled = &v
}

// GetEldYmEnabled returns the EldYmEnabled field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetEldYmEnabled() bool {
	if o == nil || IsNil(o.EldYmEnabled) {
		var ret bool
		return ret
	}
	return *o.EldYmEnabled
}

// GetEldYmEnabledOk returns a tuple with the EldYmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetEldYmEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EldYmEnabled) {
		return nil, false
	}
	return o.EldYmEnabled, true
}

// HasEldYmEnabled returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasEldYmEnabled() bool {
	if o != nil && !IsNil(o.EldYmEnabled) {
		return true
	}

	return false
}

// SetEldYmEnabled gets a reference to the given bool and assigns it to the EldYmEnabled field.
func (o *UpdateDriverRequest) SetEldYmEnabled(v bool) {
	o.EldYmEnabled = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetExternalIds() map[string]string {
	if o == nil || IsNil(o.ExternalIds) {
		var ret map[string]string
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetExternalIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given map[string]string and assigns it to the ExternalIds field.
func (o *UpdateDriverRequest) SetExternalIds(v map[string]string) {
	o.ExternalIds = &v
}

// GetLicenseNumber returns the LicenseNumber field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetLicenseNumber() string {
	if o == nil || IsNil(o.LicenseNumber) {
		var ret string
		return ret
	}
	return *o.LicenseNumber
}

// GetLicenseNumberOk returns a tuple with the LicenseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetLicenseNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseNumber) {
		return nil, false
	}
	return o.LicenseNumber, true
}

// HasLicenseNumber returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasLicenseNumber() bool {
	if o != nil && !IsNil(o.LicenseNumber) {
		return true
	}

	return false
}

// SetLicenseNumber gets a reference to the given string and assigns it to the LicenseNumber field.
func (o *UpdateDriverRequest) SetLicenseNumber(v string) {
	o.LicenseNumber = &v
}

// GetLicenseState returns the LicenseState field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetLicenseState() string {
	if o == nil || IsNil(o.LicenseState) {
		var ret string
		return ret
	}
	return *o.LicenseState
}

// GetLicenseStateOk returns a tuple with the LicenseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetLicenseStateOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseState) {
		return nil, false
	}
	return o.LicenseState, true
}

// HasLicenseState returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasLicenseState() bool {
	if o != nil && !IsNil(o.LicenseState) {
		return true
	}

	return false
}

// SetLicenseState gets a reference to the given string and assigns it to the LicenseState field.
func (o *UpdateDriverRequest) SetLicenseState(v string) {
	o.LicenseState = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetLocale() DriverLocale {
	if o == nil || IsNil(o.Locale) {
		var ret DriverLocale
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetLocaleOk() (*DriverLocale, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given DriverLocale and assigns it to the Locale field.
func (o *UpdateDriverRequest) SetLocale(v DriverLocale) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDriverRequest) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateDriverRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateDriverRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UpdateDriverRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetStaticAssignedVehicleId returns the StaticAssignedVehicleId field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetStaticAssignedVehicleId() string {
	if o == nil || IsNil(o.StaticAssignedVehicleId) {
		var ret string
		return ret
	}
	return *o.StaticAssignedVehicleId
}

// GetStaticAssignedVehicleIdOk returns a tuple with the StaticAssignedVehicleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetStaticAssignedVehicleIdOk() (*string, bool) {
	if o == nil || IsNil(o.StaticAssignedVehicleId) {
		return nil, false
	}
	return o.StaticAssignedVehicleId, true
}

// HasStaticAssignedVehicleId returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasStaticAssignedVehicleId() bool {
	if o != nil && !IsNil(o.StaticAssignedVehicleId) {
		return true
	}

	return false
}

// SetStaticAssignedVehicleId gets a reference to the given string and assigns it to the StaticAssignedVehicleId field.
func (o *UpdateDriverRequest) SetStaticAssignedVehicleId(v string) {
	o.StaticAssignedVehicleId = &v
}

// GetTachographCardNumber returns the TachographCardNumber field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetTachographCardNumber() string {
	if o == nil || IsNil(o.TachographCardNumber) {
		var ret string
		return ret
	}
	return *o.TachographCardNumber
}

// GetTachographCardNumberOk returns a tuple with the TachographCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetTachographCardNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TachographCardNumber) {
		return nil, false
	}
	return o.TachographCardNumber, true
}

// HasTachographCardNumber returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasTachographCardNumber() bool {
	if o != nil && !IsNil(o.TachographCardNumber) {
		return true
	}

	return false
}

// SetTachographCardNumber gets a reference to the given string and assigns it to the TachographCardNumber field.
func (o *UpdateDriverRequest) SetTachographCardNumber(v string) {
	o.TachographCardNumber = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetTagIds() []string {
	if o == nil || IsNil(o.TagIds) {
		var ret []string
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetTagIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TagIds) {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasTagIds() bool {
	if o != nil && !IsNil(o.TagIds) {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []string and assigns it to the TagIds field.
func (o *UpdateDriverRequest) SetTagIds(v []string) {
	o.TagIds = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UpdateDriverRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateDriverRequest) SetUsername(v string) {
	o.Username = &v
}

// GetVehicleGroupTagId returns the VehicleGroupTagId field value if set, zero value otherwise.
func (o *UpdateDriverRequest) GetVehicleGroupTagId() string {
	if o == nil || IsNil(o.VehicleGroupTagId) {
		var ret string
		return ret
	}
	return *o.VehicleGroupTagId
}

// GetVehicleGroupTagIdOk returns a tuple with the VehicleGroupTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDriverRequest) GetVehicleGroupTagIdOk() (*string, bool) {
	if o == nil || IsNil(o.VehicleGroupTagId) {
		return nil, false
	}
	return o.VehicleGroupTagId, true
}

// HasVehicleGroupTagId returns a boolean if a field has been set.
func (o *UpdateDriverRequest) HasVehicleGroupTagId() bool {
	if o != nil && !IsNil(o.VehicleGroupTagId) {
		return true
	}

	return false
}

// SetVehicleGroupTagId gets a reference to the given string and assigns it to the VehicleGroupTagId field.
func (o *UpdateDriverRequest) SetVehicleGroupTagId(v string) {
	o.VehicleGroupTagId = &v
}

func (o UpdateDriverRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDriverRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierSettings) {
		toSerialize["carrierSettings"] = o.CarrierSettings
	}
	if !IsNil(o.EldAdverseWeatherExemptionEnabled) {
		toSerialize["eldAdverseWeatherExemptionEnabled"] = o.EldAdverseWeatherExemptionEnabled
	}
	if !IsNil(o.EldBigDayExemptionEnabled) {
		toSerialize["eldBigDayExemptionEnabled"] = o.EldBigDayExemptionEnabled
	}
	if !IsNil(o.EldDayStartHour) {
		toSerialize["eldDayStartHour"] = o.EldDayStartHour
	}
	if !IsNil(o.EldExempt) {
		toSerialize["eldExempt"] = o.EldExempt
	}
	if !IsNil(o.EldExemptReason) {
		toSerialize["eldExemptReason"] = o.EldExemptReason
	}
	if !IsNil(o.EldPcEnabled) {
		toSerialize["eldPcEnabled"] = o.EldPcEnabled
	}
	if !IsNil(o.EldYmEnabled) {
		toSerialize["eldYmEnabled"] = o.EldYmEnabled
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !IsNil(o.LicenseNumber) {
		toSerialize["licenseNumber"] = o.LicenseNumber
	}
	if !IsNil(o.LicenseState) {
		toSerialize["licenseState"] = o.LicenseState
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.StaticAssignedVehicleId) {
		toSerialize["staticAssignedVehicleId"] = o.StaticAssignedVehicleId
	}
	if !IsNil(o.TachographCardNumber) {
		toSerialize["tachographCardNumber"] = o.TachographCardNumber
	}
	if !IsNil(o.TagIds) {
		toSerialize["tagIds"] = o.TagIds
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.VehicleGroupTagId) {
		toSerialize["vehicleGroupTagId"] = o.VehicleGroupTagId
	}
	return toSerialize, nil
}

type NullableUpdateDriverRequest struct {
	value *UpdateDriverRequest
	isSet bool
}

func (v NullableUpdateDriverRequest) Get() *UpdateDriverRequest {
	return v.value
}

func (v *NullableUpdateDriverRequest) Set(val *UpdateDriverRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDriverRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDriverRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDriverRequest(val *UpdateDriverRequest) *NullableUpdateDriverRequest {
	return &NullableUpdateDriverRequest{value: val, isSet: true}
}

func (v NullableUpdateDriverRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDriverRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


