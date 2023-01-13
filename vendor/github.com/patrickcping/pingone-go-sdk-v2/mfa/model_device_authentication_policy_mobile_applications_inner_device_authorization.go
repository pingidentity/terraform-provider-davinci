/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization struct for DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization
type DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization struct {
	// Specifies the enabled or disabled state of automatic MFA for native devices paired with the user, for the specified application.
	Enabled bool `json:"enabled"`
	ExtraVerification *EnumMFADevicePolicyMobileExtraVerification `json:"extraVerification,omitempty"`
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization(enabled bool) *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization{}
	this.Enabled = enabled
	return &this
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorizationWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorizationWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExtraVerification returns the ExtraVerification field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetExtraVerification() EnumMFADevicePolicyMobileExtraVerification {
	if o == nil || isNil(o.ExtraVerification) {
		var ret EnumMFADevicePolicyMobileExtraVerification
		return ret
	}
	return *o.ExtraVerification
}

// GetExtraVerificationOk returns a tuple with the ExtraVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetExtraVerificationOk() (*EnumMFADevicePolicyMobileExtraVerification, bool) {
	if o == nil || isNil(o.ExtraVerification) {
    return nil, false
	}
	return o.ExtraVerification, true
}

// HasExtraVerification returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) HasExtraVerification() bool {
	if o != nil && !isNil(o.ExtraVerification) {
		return true
	}

	return false
}

// SetExtraVerification gets a reference to the given EnumMFADevicePolicyMobileExtraVerification and assigns it to the ExtraVerification field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) SetExtraVerification(v EnumMFADevicePolicyMobileExtraVerification) {
	o.ExtraVerification = &v
}

func (o DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.ExtraVerification) {
		toSerialize["extraVerification"] = o.ExtraVerification
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization struct {
	value *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) Get() *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) Set(val *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization(val *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) *NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization {
	return &NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


