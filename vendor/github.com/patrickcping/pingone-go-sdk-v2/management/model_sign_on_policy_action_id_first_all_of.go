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

// SignOnPolicyActionIDFirstAllOf struct for SignOnPolicyActionIDFirstAllOf
type SignOnPolicyActionIDFirstAllOf struct {
	// The list of IDP discovery rules that are evaluated in order when no user is associated with the user identifier. The maximum number of rules is 100. The condition on which this identity provider is used to authenticate the user is expressed using the PingOne policy condition language
	DiscoveryRules []SignOnPolicyActionIDFirstAllOfDiscoveryRules `json:"discoveryRules,omitempty"`
	// A boolean that if set to true and if the user's account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented.
	EnforceLockoutForIdentityProviders *bool `json:"enforceLockoutForIdentityProviders,omitempty"`
	Recovery *SignOnPolicyActionLoginAllOfRecovery `json:"recovery,omitempty"`
	Registration *SignOnPolicyActionLoginAllOfRegistration `json:"registration,omitempty"`
	// An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow.
	SocialProviders []SignOnPolicyActionLoginAllOfSocialProviders `json:"socialProviders,omitempty"`
}

// NewSignOnPolicyActionIDFirstAllOf instantiates a new SignOnPolicyActionIDFirstAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDFirstAllOf() *SignOnPolicyActionIDFirstAllOf {
	this := SignOnPolicyActionIDFirstAllOf{}
	return &this
}

// NewSignOnPolicyActionIDFirstAllOfWithDefaults instantiates a new SignOnPolicyActionIDFirstAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDFirstAllOfWithDefaults() *SignOnPolicyActionIDFirstAllOf {
	this := SignOnPolicyActionIDFirstAllOf{}
	return &this
}

// GetDiscoveryRules returns the DiscoveryRules field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirstAllOf) GetDiscoveryRules() []SignOnPolicyActionIDFirstAllOfDiscoveryRules {
	if o == nil || isNil(o.DiscoveryRules) {
		var ret []SignOnPolicyActionIDFirstAllOfDiscoveryRules
		return ret
	}
	return o.DiscoveryRules
}

// GetDiscoveryRulesOk returns a tuple with the DiscoveryRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOf) GetDiscoveryRulesOk() ([]SignOnPolicyActionIDFirstAllOfDiscoveryRules, bool) {
	if o == nil || isNil(o.DiscoveryRules) {
    return nil, false
	}
	return o.DiscoveryRules, true
}

// HasDiscoveryRules returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirstAllOf) HasDiscoveryRules() bool {
	if o != nil && !isNil(o.DiscoveryRules) {
		return true
	}

	return false
}

// SetDiscoveryRules gets a reference to the given []SignOnPolicyActionIDFirstAllOfDiscoveryRules and assigns it to the DiscoveryRules field.
func (o *SignOnPolicyActionIDFirstAllOf) SetDiscoveryRules(v []SignOnPolicyActionIDFirstAllOfDiscoveryRules) {
	o.DiscoveryRules = v
}

// GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirstAllOf) GetEnforceLockoutForIdentityProviders() bool {
	if o == nil || isNil(o.EnforceLockoutForIdentityProviders) {
		var ret bool
		return ret
	}
	return *o.EnforceLockoutForIdentityProviders
}

// GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOf) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool) {
	if o == nil || isNil(o.EnforceLockoutForIdentityProviders) {
    return nil, false
	}
	return o.EnforceLockoutForIdentityProviders, true
}

// HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirstAllOf) HasEnforceLockoutForIdentityProviders() bool {
	if o != nil && !isNil(o.EnforceLockoutForIdentityProviders) {
		return true
	}

	return false
}

// SetEnforceLockoutForIdentityProviders gets a reference to the given bool and assigns it to the EnforceLockoutForIdentityProviders field.
func (o *SignOnPolicyActionIDFirstAllOf) SetEnforceLockoutForIdentityProviders(v bool) {
	o.EnforceLockoutForIdentityProviders = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirstAllOf) GetRecovery() SignOnPolicyActionLoginAllOfRecovery {
	if o == nil || isNil(o.Recovery) {
		var ret SignOnPolicyActionLoginAllOfRecovery
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOf) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool) {
	if o == nil || isNil(o.Recovery) {
    return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirstAllOf) HasRecovery() bool {
	if o != nil && !isNil(o.Recovery) {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given SignOnPolicyActionLoginAllOfRecovery and assigns it to the Recovery field.
func (o *SignOnPolicyActionIDFirstAllOf) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery) {
	o.Recovery = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirstAllOf) GetRegistration() SignOnPolicyActionLoginAllOfRegistration {
	if o == nil || isNil(o.Registration) {
		var ret SignOnPolicyActionLoginAllOfRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOf) GetRegistrationOk() (*SignOnPolicyActionLoginAllOfRegistration, bool) {
	if o == nil || isNil(o.Registration) {
    return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirstAllOf) HasRegistration() bool {
	if o != nil && !isNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionLoginAllOfRegistration and assigns it to the Registration field.
func (o *SignOnPolicyActionIDFirstAllOf) SetRegistration(v SignOnPolicyActionLoginAllOfRegistration) {
	o.Registration = &v
}

// GetSocialProviders returns the SocialProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDFirstAllOf) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders {
	if o == nil || isNil(o.SocialProviders) {
		var ret []SignOnPolicyActionLoginAllOfSocialProviders
		return ret
	}
	return o.SocialProviders
}

// GetSocialProvidersOk returns a tuple with the SocialProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstAllOf) GetSocialProvidersOk() ([]SignOnPolicyActionLoginAllOfSocialProviders, bool) {
	if o == nil || isNil(o.SocialProviders) {
    return nil, false
	}
	return o.SocialProviders, true
}

// HasSocialProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDFirstAllOf) HasSocialProviders() bool {
	if o != nil && !isNil(o.SocialProviders) {
		return true
	}

	return false
}

// SetSocialProviders gets a reference to the given []SignOnPolicyActionLoginAllOfSocialProviders and assigns it to the SocialProviders field.
func (o *SignOnPolicyActionIDFirstAllOf) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders) {
	o.SocialProviders = v
}

func (o SignOnPolicyActionIDFirstAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DiscoveryRules) {
		toSerialize["discoveryRules"] = o.DiscoveryRules
	}
	if !isNil(o.EnforceLockoutForIdentityProviders) {
		toSerialize["enforceLockoutForIdentityProviders"] = o.EnforceLockoutForIdentityProviders
	}
	if !isNil(o.Recovery) {
		toSerialize["recovery"] = o.Recovery
	}
	if !isNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	if !isNil(o.SocialProviders) {
		toSerialize["socialProviders"] = o.SocialProviders
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionIDFirstAllOf struct {
	value *SignOnPolicyActionIDFirstAllOf
	isSet bool
}

func (v NullableSignOnPolicyActionIDFirstAllOf) Get() *SignOnPolicyActionIDFirstAllOf {
	return v.value
}

func (v *NullableSignOnPolicyActionIDFirstAllOf) Set(val *SignOnPolicyActionIDFirstAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDFirstAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDFirstAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDFirstAllOf(val *SignOnPolicyActionIDFirstAllOf) *NullableSignOnPolicyActionIDFirstAllOf {
	return &NullableSignOnPolicyActionIDFirstAllOf{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDFirstAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDFirstAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


