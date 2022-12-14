/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// CustomDomainCertificate An object that specifies information about the SSL certificate used by this custom domain. If this property is not present, it indicates that an SSL certificate has not been setup for this custom domain.
type CustomDomainCertificate struct {
	// The time when the certificate expires.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewCustomDomainCertificate instantiates a new CustomDomainCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomainCertificate() *CustomDomainCertificate {
	this := CustomDomainCertificate{}
	return &this
}

// NewCustomDomainCertificateWithDefaults instantiates a new CustomDomainCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainCertificateWithDefaults() *CustomDomainCertificate {
	this := CustomDomainCertificate{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CustomDomainCertificate) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomainCertificate) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CustomDomainCertificate) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CustomDomainCertificate) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o CustomDomainCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableCustomDomainCertificate struct {
	value *CustomDomainCertificate
	isSet bool
}

func (v NullableCustomDomainCertificate) Get() *CustomDomainCertificate {
	return v.value
}

func (v *NullableCustomDomainCertificate) Set(val *CustomDomainCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomainCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomainCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomainCertificate(val *CustomDomainCertificate) *NullableCustomDomainCertificate {
	return &NullableCustomDomainCertificate{value: val, isSet: true}
}

func (v NullableCustomDomainCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomainCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


