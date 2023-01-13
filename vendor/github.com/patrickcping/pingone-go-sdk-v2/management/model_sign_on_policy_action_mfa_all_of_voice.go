/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionMFAAllOfVoice Specifies MFA through voice messaging options.
type SignOnPolicyActionMFAAllOfVoice struct {
	// A boolean that specifies the enabled/disabled state of the MFA through voice message action. The default is disabled when creating a new policy. When enabled, it allows users to receive the one-time password and authenticate through a voice message.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSignOnPolicyActionMFAAllOfVoice instantiates a new SignOnPolicyActionMFAAllOfVoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAAllOfVoice() *SignOnPolicyActionMFAAllOfVoice {
	this := SignOnPolicyActionMFAAllOfVoice{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSignOnPolicyActionMFAAllOfVoiceWithDefaults instantiates a new SignOnPolicyActionMFAAllOfVoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAAllOfVoiceWithDefaults() *SignOnPolicyActionMFAAllOfVoice {
	this := SignOnPolicyActionMFAAllOfVoice{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAAllOfVoice) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAllOfVoice) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAAllOfVoice) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFAAllOfVoice) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SignOnPolicyActionMFAAllOfVoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFAAllOfVoice struct {
	value *SignOnPolicyActionMFAAllOfVoice
	isSet bool
}

func (v NullableSignOnPolicyActionMFAAllOfVoice) Get() *SignOnPolicyActionMFAAllOfVoice {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAAllOfVoice) Set(val *SignOnPolicyActionMFAAllOfVoice) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAAllOfVoice) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAAllOfVoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAAllOfVoice(val *SignOnPolicyActionMFAAllOfVoice) *NullableSignOnPolicyActionMFAAllOfVoice {
	return &NullableSignOnPolicyActionMFAAllOfVoice{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAAllOfVoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAAllOfVoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


