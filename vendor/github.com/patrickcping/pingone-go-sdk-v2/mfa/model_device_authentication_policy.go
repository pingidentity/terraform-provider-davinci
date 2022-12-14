/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"time"
)

// DeviceAuthenticationPolicy struct for DeviceAuthenticationPolicy
type DeviceAuthenticationPolicy struct {
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// Device authentication policy's UUID.
	Id *string `json:"id,omitempty"`
	// Device authentication policy's name.
	Name string `json:"name"`
	Sms DeviceAuthenticationPolicyOfflineDevice `json:"sms"`
	Voice DeviceAuthenticationPolicyOfflineDevice `json:"voice"`
	Email DeviceAuthenticationPolicyOfflineDevice `json:"email"`
	Mobile DeviceAuthenticationPolicyMobile `json:"mobile"`
	Totp DeviceAuthenticationPolicyTotp `json:"totp"`
	SecurityKey DeviceAuthenticationPolicyFIDODevice `json:"securityKey"`
	Platform DeviceAuthenticationPolicyFIDODevice `json:"platform"`
	// The default policy for Flow Manager.
	Default bool `json:"default"`
	// Deprecated
	ForSignOnPolicy bool `json:"forSignOnPolicy"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewDeviceAuthenticationPolicy instantiates a new DeviceAuthenticationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicy(name string, sms DeviceAuthenticationPolicyOfflineDevice, voice DeviceAuthenticationPolicyOfflineDevice, email DeviceAuthenticationPolicyOfflineDevice, mobile DeviceAuthenticationPolicyMobile, totp DeviceAuthenticationPolicyTotp, securityKey DeviceAuthenticationPolicyFIDODevice, platform DeviceAuthenticationPolicyFIDODevice, default_ bool, forSignOnPolicy bool) *DeviceAuthenticationPolicy {
	this := DeviceAuthenticationPolicy{}
	this.Name = name
	this.Sms = sms
	this.Voice = voice
	this.Email = email
	this.Mobile = mobile
	this.Totp = totp
	this.SecurityKey = securityKey
	this.Platform = platform
	this.Default = default_
	this.ForSignOnPolicy = forSignOnPolicy
	return &this
}

// NewDeviceAuthenticationPolicyWithDefaults instantiates a new DeviceAuthenticationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyWithDefaults() *DeviceAuthenticationPolicy {
	this := DeviceAuthenticationPolicy{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *DeviceAuthenticationPolicy) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceAuthenticationPolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DeviceAuthenticationPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceAuthenticationPolicy) SetName(v string) {
	o.Name = v
}

// GetSms returns the Sms field value
func (o *DeviceAuthenticationPolicy) GetSms() DeviceAuthenticationPolicyOfflineDevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDevice
		return ret
	}

	return o.Sms
}

// GetSmsOk returns a tuple with the Sms field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetSmsOk() (*DeviceAuthenticationPolicyOfflineDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sms, true
}

// SetSms sets field value
func (o *DeviceAuthenticationPolicy) SetSms(v DeviceAuthenticationPolicyOfflineDevice) {
	o.Sms = v
}

// GetVoice returns the Voice field value
func (o *DeviceAuthenticationPolicy) GetVoice() DeviceAuthenticationPolicyOfflineDevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDevice
		return ret
	}

	return o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetVoiceOk() (*DeviceAuthenticationPolicyOfflineDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Voice, true
}

// SetVoice sets field value
func (o *DeviceAuthenticationPolicy) SetVoice(v DeviceAuthenticationPolicyOfflineDevice) {
	o.Voice = v
}

// GetEmail returns the Email field value
func (o *DeviceAuthenticationPolicy) GetEmail() DeviceAuthenticationPolicyOfflineDevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDevice
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetEmailOk() (*DeviceAuthenticationPolicyOfflineDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *DeviceAuthenticationPolicy) SetEmail(v DeviceAuthenticationPolicyOfflineDevice) {
	o.Email = v
}

// GetMobile returns the Mobile field value
func (o *DeviceAuthenticationPolicy) GetMobile() DeviceAuthenticationPolicyMobile {
	if o == nil {
		var ret DeviceAuthenticationPolicyMobile
		return ret
	}

	return o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetMobileOk() (*DeviceAuthenticationPolicyMobile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mobile, true
}

// SetMobile sets field value
func (o *DeviceAuthenticationPolicy) SetMobile(v DeviceAuthenticationPolicyMobile) {
	o.Mobile = v
}

// GetTotp returns the Totp field value
func (o *DeviceAuthenticationPolicy) GetTotp() DeviceAuthenticationPolicyTotp {
	if o == nil {
		var ret DeviceAuthenticationPolicyTotp
		return ret
	}

	return o.Totp
}

// GetTotpOk returns a tuple with the Totp field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetTotpOk() (*DeviceAuthenticationPolicyTotp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Totp, true
}

// SetTotp sets field value
func (o *DeviceAuthenticationPolicy) SetTotp(v DeviceAuthenticationPolicyTotp) {
	o.Totp = v
}

// GetSecurityKey returns the SecurityKey field value
func (o *DeviceAuthenticationPolicy) GetSecurityKey() DeviceAuthenticationPolicyFIDODevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyFIDODevice
		return ret
	}

	return o.SecurityKey
}

// GetSecurityKeyOk returns a tuple with the SecurityKey field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetSecurityKeyOk() (*DeviceAuthenticationPolicyFIDODevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityKey, true
}

// SetSecurityKey sets field value
func (o *DeviceAuthenticationPolicy) SetSecurityKey(v DeviceAuthenticationPolicyFIDODevice) {
	o.SecurityKey = v
}

// GetPlatform returns the Platform field value
func (o *DeviceAuthenticationPolicy) GetPlatform() DeviceAuthenticationPolicyFIDODevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyFIDODevice
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetPlatformOk() (*DeviceAuthenticationPolicyFIDODevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *DeviceAuthenticationPolicy) SetPlatform(v DeviceAuthenticationPolicyFIDODevice) {
	o.Platform = v
}

// GetDefault returns the Default field value
func (o *DeviceAuthenticationPolicy) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *DeviceAuthenticationPolicy) SetDefault(v bool) {
	o.Default = v
}

// GetForSignOnPolicy returns the ForSignOnPolicy field value
// Deprecated
func (o *DeviceAuthenticationPolicy) GetForSignOnPolicy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForSignOnPolicy
}

// GetForSignOnPolicyOk returns a tuple with the ForSignOnPolicy field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceAuthenticationPolicy) GetForSignOnPolicyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForSignOnPolicy, true
}

// SetForSignOnPolicy sets field value
// Deprecated
func (o *DeviceAuthenticationPolicy) SetForSignOnPolicy(v bool) {
	o.ForSignOnPolicy = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeviceAuthenticationPolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DeviceAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["sms"] = o.Sms
	}
	if true {
		toSerialize["voice"] = o.Voice
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["mobile"] = o.Mobile
	}
	if true {
		toSerialize["totp"] = o.Totp
	}
	if true {
		toSerialize["securityKey"] = o.SecurityKey
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["default"] = o.Default
	}
	if true {
		toSerialize["forSignOnPolicy"] = o.ForSignOnPolicy
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicy struct {
	value *DeviceAuthenticationPolicy
	isSet bool
}

func (v NullableDeviceAuthenticationPolicy) Get() *DeviceAuthenticationPolicy {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicy) Set(val *DeviceAuthenticationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicy(val *DeviceAuthenticationPolicy) *NullableDeviceAuthenticationPolicy {
	return &NullableDeviceAuthenticationPolicy{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


