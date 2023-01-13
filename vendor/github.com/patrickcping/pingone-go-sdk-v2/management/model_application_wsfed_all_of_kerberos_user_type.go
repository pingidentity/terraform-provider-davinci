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

// ApplicationWSFEDAllOfKerberosUserType struct for ApplicationWSFEDAllOfKerberosUserType
type ApplicationWSFEDAllOfKerberosUserType struct {
	// The UUID of a user type in the list of `userTypes` for the LDAP gateway.
	Id *string `json:"id,omitempty"`
}

// NewApplicationWSFEDAllOfKerberosUserType instantiates a new ApplicationWSFEDAllOfKerberosUserType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationWSFEDAllOfKerberosUserType() *ApplicationWSFEDAllOfKerberosUserType {
	this := ApplicationWSFEDAllOfKerberosUserType{}
	return &this
}

// NewApplicationWSFEDAllOfKerberosUserTypeWithDefaults instantiates a new ApplicationWSFEDAllOfKerberosUserType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWSFEDAllOfKerberosUserTypeWithDefaults() *ApplicationWSFEDAllOfKerberosUserType {
	this := ApplicationWSFEDAllOfKerberosUserType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationWSFEDAllOfKerberosUserType) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationWSFEDAllOfKerberosUserType) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationWSFEDAllOfKerberosUserType) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationWSFEDAllOfKerberosUserType) SetId(v string) {
	o.Id = &v
}

func (o ApplicationWSFEDAllOfKerberosUserType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationWSFEDAllOfKerberosUserType struct {
	value *ApplicationWSFEDAllOfKerberosUserType
	isSet bool
}

func (v NullableApplicationWSFEDAllOfKerberosUserType) Get() *ApplicationWSFEDAllOfKerberosUserType {
	return v.value
}

func (v *NullableApplicationWSFEDAllOfKerberosUserType) Set(val *ApplicationWSFEDAllOfKerberosUserType) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationWSFEDAllOfKerberosUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationWSFEDAllOfKerberosUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationWSFEDAllOfKerberosUserType(val *ApplicationWSFEDAllOfKerberosUserType) *NullableApplicationWSFEDAllOfKerberosUserType {
	return &NullableApplicationWSFEDAllOfKerberosUserType{value: val, isSet: true}
}

func (v NullableApplicationWSFEDAllOfKerberosUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationWSFEDAllOfKerberosUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


