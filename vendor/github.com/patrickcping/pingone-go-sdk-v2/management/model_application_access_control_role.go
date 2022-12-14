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

// ApplicationAccessControlRole struct for ApplicationAccessControlRole
type ApplicationAccessControlRole struct {
	Type EnumApplicationAccessControlType `json:"type"`
}

// NewApplicationAccessControlRole instantiates a new ApplicationAccessControlRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAccessControlRole(type_ EnumApplicationAccessControlType) *ApplicationAccessControlRole {
	this := ApplicationAccessControlRole{}
	this.Type = type_
	return &this
}

// NewApplicationAccessControlRoleWithDefaults instantiates a new ApplicationAccessControlRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAccessControlRoleWithDefaults() *ApplicationAccessControlRole {
	this := ApplicationAccessControlRole{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationAccessControlRole) GetType() EnumApplicationAccessControlType {
	if o == nil {
		var ret EnumApplicationAccessControlType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationAccessControlRole) GetTypeOk() (*EnumApplicationAccessControlType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationAccessControlRole) SetType(v EnumApplicationAccessControlType) {
	o.Type = v
}

func (o ApplicationAccessControlRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationAccessControlRole struct {
	value *ApplicationAccessControlRole
	isSet bool
}

func (v NullableApplicationAccessControlRole) Get() *ApplicationAccessControlRole {
	return v.value
}

func (v *NullableApplicationAccessControlRole) Set(val *ApplicationAccessControlRole) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAccessControlRole) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAccessControlRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAccessControlRole(val *ApplicationAccessControlRole) *NullableApplicationAccessControlRole {
	return &NullableApplicationAccessControlRole{value: val, isSet: true}
}

func (v NullableApplicationAccessControlRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAccessControlRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


