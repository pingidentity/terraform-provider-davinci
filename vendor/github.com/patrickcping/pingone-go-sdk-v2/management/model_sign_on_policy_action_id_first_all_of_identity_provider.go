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

// SignOnPolicyActionIDFirstAllOfIdentityProvider struct for SignOnPolicyActionIDFirstAllOfIdentityProvider
type SignOnPolicyActionIDFirstAllOfIdentityProvider struct {
	// A string that specifies the identity provider that will be used to authenticate the user if the condition is matched.
	Id string `json:"id"`
}

// NewSignOnPolicyActionIDFirstAllOfIdentityProvider instantiates a new SignOnPolicyActionIDFirstAllOfIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDFirstAllOfIdentityProvider(id string) *SignOnPolicyActionIDFirstAllOfIdentityProvider {
	this := SignOnPolicyActionIDFirstAllOfIdentityProvider{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionIDFirstAllOfIdentityProviderWithDefaults instantiates a new SignOnPolicyActionIDFirstAllOfIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDFirstAllOfIdentityProviderWithDefaults() *SignOnPolicyActionIDFirstAllOfIdentityProvider {
	this := SignOnPolicyActionIDFirstAllOfIdentityProvider{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionIDFirstAllOfIdentityProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOfIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionIDFirstAllOfIdentityProvider) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionIDFirstAllOfIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionIDFirstAllOfIdentityProvider struct {
	value *SignOnPolicyActionIDFirstAllOfIdentityProvider
	isSet bool
}

func (v NullableSignOnPolicyActionIDFirstAllOfIdentityProvider) Get() *SignOnPolicyActionIDFirstAllOfIdentityProvider {
	return v.value
}

func (v *NullableSignOnPolicyActionIDFirstAllOfIdentityProvider) Set(val *SignOnPolicyActionIDFirstAllOfIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDFirstAllOfIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDFirstAllOfIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDFirstAllOfIdentityProvider(val *SignOnPolicyActionIDFirstAllOfIdentityProvider) *NullableSignOnPolicyActionIDFirstAllOfIdentityProvider {
	return &NullableSignOnPolicyActionIDFirstAllOfIdentityProvider{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDFirstAllOfIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDFirstAllOfIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


