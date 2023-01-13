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

// SignOnPolicyActionAgreementAllOfAgreement The relationship to the agreement to which the user must consent. The agreement must exist and be enabled. An agreement cannot be disabed if an action uses it. An enabled agreement must always support the default language. This property is required.
type SignOnPolicyActionAgreementAllOfAgreement struct {
	// A string that specifies the ID of the agreement to which the user must consent. This property is required.
	Id string `json:"id"`
}

// NewSignOnPolicyActionAgreementAllOfAgreement instantiates a new SignOnPolicyActionAgreementAllOfAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionAgreementAllOfAgreement(id string) *SignOnPolicyActionAgreementAllOfAgreement {
	this := SignOnPolicyActionAgreementAllOfAgreement{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionAgreementAllOfAgreementWithDefaults instantiates a new SignOnPolicyActionAgreementAllOfAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionAgreementAllOfAgreementWithDefaults() *SignOnPolicyActionAgreementAllOfAgreement {
	this := SignOnPolicyActionAgreementAllOfAgreement{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionAgreementAllOfAgreement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreementAllOfAgreement) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionAgreementAllOfAgreement) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionAgreementAllOfAgreement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionAgreementAllOfAgreement struct {
	value *SignOnPolicyActionAgreementAllOfAgreement
	isSet bool
}

func (v NullableSignOnPolicyActionAgreementAllOfAgreement) Get() *SignOnPolicyActionAgreementAllOfAgreement {
	return v.value
}

func (v *NullableSignOnPolicyActionAgreementAllOfAgreement) Set(val *SignOnPolicyActionAgreementAllOfAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionAgreementAllOfAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionAgreementAllOfAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionAgreementAllOfAgreement(val *SignOnPolicyActionAgreementAllOfAgreement) *NullableSignOnPolicyActionAgreementAllOfAgreement {
	return &NullableSignOnPolicyActionAgreementAllOfAgreement{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionAgreementAllOfAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionAgreementAllOfAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


