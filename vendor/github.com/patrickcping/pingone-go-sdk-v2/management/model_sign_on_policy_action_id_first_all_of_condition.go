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

// SignOnPolicyActionIDFirstAllOfCondition struct for SignOnPolicyActionIDFirstAllOfCondition
type SignOnPolicyActionIDFirstAllOfCondition struct {
	Contains string `json:"contains"`
	Value string `json:"value"`
}

// NewSignOnPolicyActionIDFirstAllOfCondition instantiates a new SignOnPolicyActionIDFirstAllOfCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDFirstAllOfCondition(contains string, value string) *SignOnPolicyActionIDFirstAllOfCondition {
	this := SignOnPolicyActionIDFirstAllOfCondition{}
	this.Contains = contains
	this.Value = value
	return &this
}

// NewSignOnPolicyActionIDFirstAllOfConditionWithDefaults instantiates a new SignOnPolicyActionIDFirstAllOfCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDFirstAllOfConditionWithDefaults() *SignOnPolicyActionIDFirstAllOfCondition {
	this := SignOnPolicyActionIDFirstAllOfCondition{}
	return &this
}

// GetContains returns the Contains field value
func (o *SignOnPolicyActionIDFirstAllOfCondition) GetContains() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contains
}

// GetContainsOk returns a tuple with the Contains field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOfCondition) GetContainsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contains, true
}

// SetContains sets field value
func (o *SignOnPolicyActionIDFirstAllOfCondition) SetContains(v string) {
	o.Contains = v
}

// GetValue returns the Value field value
func (o *SignOnPolicyActionIDFirstAllOfCondition) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOfCondition) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SignOnPolicyActionIDFirstAllOfCondition) SetValue(v string) {
	o.Value = v
}

func (o SignOnPolicyActionIDFirstAllOfCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contains"] = o.Contains
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionIDFirstAllOfCondition struct {
	value *SignOnPolicyActionIDFirstAllOfCondition
	isSet bool
}

func (v NullableSignOnPolicyActionIDFirstAllOfCondition) Get() *SignOnPolicyActionIDFirstAllOfCondition {
	return v.value
}

func (v *NullableSignOnPolicyActionIDFirstAllOfCondition) Set(val *SignOnPolicyActionIDFirstAllOfCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDFirstAllOfCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDFirstAllOfCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDFirstAllOfCondition(val *SignOnPolicyActionIDFirstAllOfCondition) *NullableSignOnPolicyActionIDFirstAllOfCondition {
	return &NullableSignOnPolicyActionIDFirstAllOfCondition{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDFirstAllOfCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDFirstAllOfCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


