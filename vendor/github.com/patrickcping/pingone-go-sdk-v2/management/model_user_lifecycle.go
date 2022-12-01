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

// UserLifecycle struct for UserLifecycle
type UserLifecycle struct {
	Status *EnumUserLifecycleStatus `json:"status,omitempty"`
}

// NewUserLifecycle instantiates a new UserLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLifecycle() *UserLifecycle {
	this := UserLifecycle{}
	return &this
}

// NewUserLifecycleWithDefaults instantiates a new UserLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLifecycleWithDefaults() *UserLifecycle {
	this := UserLifecycle{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserLifecycle) GetStatus() EnumUserLifecycleStatus {
	if o == nil || o.Status == nil {
		var ret EnumUserLifecycleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLifecycle) GetStatusOk() (*EnumUserLifecycleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserLifecycle) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EnumUserLifecycleStatus and assigns it to the Status field.
func (o *UserLifecycle) SetStatus(v EnumUserLifecycleStatus) {
	o.Status = &v
}

func (o UserLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUserLifecycle struct {
	value *UserLifecycle
	isSet bool
}

func (v NullableUserLifecycle) Get() *UserLifecycle {
	return v.value
}

func (v *NullableUserLifecycle) Set(val *UserLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLifecycle(val *UserLifecycle) *NullableUserLifecycle {
	return &NullableUserLifecycle{value: val, isSet: true}
}

func (v NullableUserLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

