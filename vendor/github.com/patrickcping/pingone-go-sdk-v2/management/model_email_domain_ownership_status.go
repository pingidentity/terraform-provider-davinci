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

// EmailDomainOwnershipStatus struct for EmailDomainOwnershipStatus
type EmailDomainOwnershipStatus struct {
	// A string that specifies the type of DNS record, with the value \"TXT\".
	Type *string `json:"type,omitempty"`
	// The regions collection specifies the properties for the 4 AWS SES regions that are used for sending email for the environment. The regions are determined by the geography where this environment was provisioned (North America, Canada, Europe & Asia-Pacific).
	Regions []EmailDomainOwnershipStatusRegionsInner `json:"regions,omitempty"`
}

// NewEmailDomainOwnershipStatus instantiates a new EmailDomainOwnershipStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainOwnershipStatus() *EmailDomainOwnershipStatus {
	this := EmailDomainOwnershipStatus{}
	return &this
}

// NewEmailDomainOwnershipStatusWithDefaults instantiates a new EmailDomainOwnershipStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainOwnershipStatusWithDefaults() *EmailDomainOwnershipStatus {
	this := EmailDomainOwnershipStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EmailDomainOwnershipStatus) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainOwnershipStatus) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EmailDomainOwnershipStatus) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EmailDomainOwnershipStatus) SetType(v string) {
	o.Type = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *EmailDomainOwnershipStatus) GetRegions() []EmailDomainOwnershipStatusRegionsInner {
	if o == nil || isNil(o.Regions) {
		var ret []EmailDomainOwnershipStatusRegionsInner
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainOwnershipStatus) GetRegionsOk() ([]EmailDomainOwnershipStatusRegionsInner, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *EmailDomainOwnershipStatus) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []EmailDomainOwnershipStatusRegionsInner and assigns it to the Regions field.
func (o *EmailDomainOwnershipStatus) SetRegions(v []EmailDomainOwnershipStatusRegionsInner) {
	o.Regions = v
}

func (o EmailDomainOwnershipStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableEmailDomainOwnershipStatus struct {
	value *EmailDomainOwnershipStatus
	isSet bool
}

func (v NullableEmailDomainOwnershipStatus) Get() *EmailDomainOwnershipStatus {
	return v.value
}

func (v *NullableEmailDomainOwnershipStatus) Set(val *EmailDomainOwnershipStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainOwnershipStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainOwnershipStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainOwnershipStatus(val *EmailDomainOwnershipStatus) *NullableEmailDomainOwnershipStatus {
	return &NullableEmailDomainOwnershipStatus{value: val, isSet: true}
}

func (v NullableEmailDomainOwnershipStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainOwnershipStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


