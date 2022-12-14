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

// IdentityProviderCommonRegistrationPopulation struct for IdentityProviderCommonRegistrationPopulation
type IdentityProviderCommonRegistrationPopulation struct {
	// An external IdP to use as authoritative. Setting this attribute gives management of linked users to the IdP and also triggers just-in-time provisioning of new users. These users are created in the population indicated with registration.population.id.
	Id *string `json:"id,omitempty"`
}

// NewIdentityProviderCommonRegistrationPopulation instantiates a new IdentityProviderCommonRegistrationPopulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderCommonRegistrationPopulation() *IdentityProviderCommonRegistrationPopulation {
	this := IdentityProviderCommonRegistrationPopulation{}
	return &this
}

// NewIdentityProviderCommonRegistrationPopulationWithDefaults instantiates a new IdentityProviderCommonRegistrationPopulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderCommonRegistrationPopulationWithDefaults() *IdentityProviderCommonRegistrationPopulation {
	this := IdentityProviderCommonRegistrationPopulation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderCommonRegistrationPopulation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommonRegistrationPopulation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderCommonRegistrationPopulation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProviderCommonRegistrationPopulation) SetId(v string) {
	o.Id = &v
}

func (o IdentityProviderCommonRegistrationPopulation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProviderCommonRegistrationPopulation struct {
	value *IdentityProviderCommonRegistrationPopulation
	isSet bool
}

func (v NullableIdentityProviderCommonRegistrationPopulation) Get() *IdentityProviderCommonRegistrationPopulation {
	return v.value
}

func (v *NullableIdentityProviderCommonRegistrationPopulation) Set(val *IdentityProviderCommonRegistrationPopulation) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCommonRegistrationPopulation) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCommonRegistrationPopulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCommonRegistrationPopulation(val *IdentityProviderCommonRegistrationPopulation) *NullableIdentityProviderCommonRegistrationPopulation {
	return &NullableIdentityProviderCommonRegistrationPopulation{value: val, isSet: true}
}

func (v NullableIdentityProviderCommonRegistrationPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCommonRegistrationPopulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


