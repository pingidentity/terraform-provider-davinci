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

// SignOnPolicyActionLoginAllOfRegistrationPopulation struct for SignOnPolicyActionLoginAllOfRegistrationPopulation
type SignOnPolicyActionLoginAllOfRegistrationPopulation struct {
	// A string that specifies the population ID associated with the newly registered user.
	Id string `json:"id"`
}

// NewSignOnPolicyActionLoginAllOfRegistrationPopulation instantiates a new SignOnPolicyActionLoginAllOfRegistrationPopulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLoginAllOfRegistrationPopulation(id string) *SignOnPolicyActionLoginAllOfRegistrationPopulation {
	this := SignOnPolicyActionLoginAllOfRegistrationPopulation{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionLoginAllOfRegistrationPopulationWithDefaults instantiates a new SignOnPolicyActionLoginAllOfRegistrationPopulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginAllOfRegistrationPopulationWithDefaults() *SignOnPolicyActionLoginAllOfRegistrationPopulation {
	this := SignOnPolicyActionLoginAllOfRegistrationPopulation{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionLoginAllOfRegistrationPopulation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfRegistrationPopulation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionLoginAllOfRegistrationPopulation) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionLoginAllOfRegistrationPopulation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionLoginAllOfRegistrationPopulation struct {
	value *SignOnPolicyActionLoginAllOfRegistrationPopulation
	isSet bool
}

func (v NullableSignOnPolicyActionLoginAllOfRegistrationPopulation) Get() *SignOnPolicyActionLoginAllOfRegistrationPopulation {
	return v.value
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistrationPopulation) Set(val *SignOnPolicyActionLoginAllOfRegistrationPopulation) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLoginAllOfRegistrationPopulation) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistrationPopulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLoginAllOfRegistrationPopulation(val *SignOnPolicyActionLoginAllOfRegistrationPopulation) *NullableSignOnPolicyActionLoginAllOfRegistrationPopulation {
	return &NullableSignOnPolicyActionLoginAllOfRegistrationPopulation{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLoginAllOfRegistrationPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistrationPopulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


