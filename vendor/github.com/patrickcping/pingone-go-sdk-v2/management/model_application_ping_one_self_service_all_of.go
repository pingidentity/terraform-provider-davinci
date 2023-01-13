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

// ApplicationPingOneSelfServiceAllOf struct for ApplicationPingOneSelfServiceAllOf
type ApplicationPingOneSelfServiceAllOf struct {
	// If `true`, shows the default theme footer on the self service application. Applies only if `applyDefaultTheme` is also `true`.
	EnableDefaultThemeFooter *bool `json:"enableDefaultThemeFooter,omitempty"`
	// If `true`, applies the default theme to the self service application.
	ApplyDefaultTheme bool `json:"applyDefaultTheme"`
}

// NewApplicationPingOneSelfServiceAllOf instantiates a new ApplicationPingOneSelfServiceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPingOneSelfServiceAllOf(applyDefaultTheme bool) *ApplicationPingOneSelfServiceAllOf {
	this := ApplicationPingOneSelfServiceAllOf{}
	this.ApplyDefaultTheme = applyDefaultTheme
	return &this
}

// NewApplicationPingOneSelfServiceAllOfWithDefaults instantiates a new ApplicationPingOneSelfServiceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPingOneSelfServiceAllOfWithDefaults() *ApplicationPingOneSelfServiceAllOf {
	this := ApplicationPingOneSelfServiceAllOf{}
	return &this
}

// GetEnableDefaultThemeFooter returns the EnableDefaultThemeFooter field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfServiceAllOf) GetEnableDefaultThemeFooter() bool {
	if o == nil || isNil(o.EnableDefaultThemeFooter) {
		var ret bool
		return ret
	}
	return *o.EnableDefaultThemeFooter
}

// GetEnableDefaultThemeFooterOk returns a tuple with the EnableDefaultThemeFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfServiceAllOf) GetEnableDefaultThemeFooterOk() (*bool, bool) {
	if o == nil || isNil(o.EnableDefaultThemeFooter) {
    return nil, false
	}
	return o.EnableDefaultThemeFooter, true
}

// HasEnableDefaultThemeFooter returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfServiceAllOf) HasEnableDefaultThemeFooter() bool {
	if o != nil && !isNil(o.EnableDefaultThemeFooter) {
		return true
	}

	return false
}

// SetEnableDefaultThemeFooter gets a reference to the given bool and assigns it to the EnableDefaultThemeFooter field.
func (o *ApplicationPingOneSelfServiceAllOf) SetEnableDefaultThemeFooter(v bool) {
	o.EnableDefaultThemeFooter = &v
}

// GetApplyDefaultTheme returns the ApplyDefaultTheme field value
func (o *ApplicationPingOneSelfServiceAllOf) GetApplyDefaultTheme() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ApplyDefaultTheme
}

// GetApplyDefaultThemeOk returns a tuple with the ApplyDefaultTheme field value
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfServiceAllOf) GetApplyDefaultThemeOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApplyDefaultTheme, true
}

// SetApplyDefaultTheme sets field value
func (o *ApplicationPingOneSelfServiceAllOf) SetApplyDefaultTheme(v bool) {
	o.ApplyDefaultTheme = v
}

func (o ApplicationPingOneSelfServiceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnableDefaultThemeFooter) {
		toSerialize["enableDefaultThemeFooter"] = o.EnableDefaultThemeFooter
	}
	if true {
		toSerialize["applyDefaultTheme"] = o.ApplyDefaultTheme
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPingOneSelfServiceAllOf struct {
	value *ApplicationPingOneSelfServiceAllOf
	isSet bool
}

func (v NullableApplicationPingOneSelfServiceAllOf) Get() *ApplicationPingOneSelfServiceAllOf {
	return v.value
}

func (v *NullableApplicationPingOneSelfServiceAllOf) Set(val *ApplicationPingOneSelfServiceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPingOneSelfServiceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPingOneSelfServiceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPingOneSelfServiceAllOf(val *ApplicationPingOneSelfServiceAllOf) *NullableApplicationPingOneSelfServiceAllOf {
	return &NullableApplicationPingOneSelfServiceAllOf{value: val, isSet: true}
}

func (v NullableApplicationPingOneSelfServiceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPingOneSelfServiceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


