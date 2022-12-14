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

// LicenseGateways struct for LicenseGateways
type LicenseGateways struct {
	AllowLdapGateway *bool `json:"allowLdapGateway,omitempty"`
	AllowKerberosGateway *bool `json:"allowKerberosGateway,omitempty"`
	AllowRadiusGateway *bool `json:"allowRadiusGateway,omitempty"`
}

// NewLicenseGateways instantiates a new LicenseGateways object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseGateways() *LicenseGateways {
	this := LicenseGateways{}
	return &this
}

// NewLicenseGatewaysWithDefaults instantiates a new LicenseGateways object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseGatewaysWithDefaults() *LicenseGateways {
	this := LicenseGateways{}
	return &this
}

// GetAllowLdapGateway returns the AllowLdapGateway field value if set, zero value otherwise.
func (o *LicenseGateways) GetAllowLdapGateway() bool {
	if o == nil || o.AllowLdapGateway == nil {
		var ret bool
		return ret
	}
	return *o.AllowLdapGateway
}

// GetAllowLdapGatewayOk returns a tuple with the AllowLdapGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseGateways) GetAllowLdapGatewayOk() (*bool, bool) {
	if o == nil || o.AllowLdapGateway == nil {
		return nil, false
	}
	return o.AllowLdapGateway, true
}

// HasAllowLdapGateway returns a boolean if a field has been set.
func (o *LicenseGateways) HasAllowLdapGateway() bool {
	if o != nil && o.AllowLdapGateway != nil {
		return true
	}

	return false
}

// SetAllowLdapGateway gets a reference to the given bool and assigns it to the AllowLdapGateway field.
func (o *LicenseGateways) SetAllowLdapGateway(v bool) {
	o.AllowLdapGateway = &v
}

// GetAllowKerberosGateway returns the AllowKerberosGateway field value if set, zero value otherwise.
func (o *LicenseGateways) GetAllowKerberosGateway() bool {
	if o == nil || o.AllowKerberosGateway == nil {
		var ret bool
		return ret
	}
	return *o.AllowKerberosGateway
}

// GetAllowKerberosGatewayOk returns a tuple with the AllowKerberosGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseGateways) GetAllowKerberosGatewayOk() (*bool, bool) {
	if o == nil || o.AllowKerberosGateway == nil {
		return nil, false
	}
	return o.AllowKerberosGateway, true
}

// HasAllowKerberosGateway returns a boolean if a field has been set.
func (o *LicenseGateways) HasAllowKerberosGateway() bool {
	if o != nil && o.AllowKerberosGateway != nil {
		return true
	}

	return false
}

// SetAllowKerberosGateway gets a reference to the given bool and assigns it to the AllowKerberosGateway field.
func (o *LicenseGateways) SetAllowKerberosGateway(v bool) {
	o.AllowKerberosGateway = &v
}

// GetAllowRadiusGateway returns the AllowRadiusGateway field value if set, zero value otherwise.
func (o *LicenseGateways) GetAllowRadiusGateway() bool {
	if o == nil || o.AllowRadiusGateway == nil {
		var ret bool
		return ret
	}
	return *o.AllowRadiusGateway
}

// GetAllowRadiusGatewayOk returns a tuple with the AllowRadiusGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseGateways) GetAllowRadiusGatewayOk() (*bool, bool) {
	if o == nil || o.AllowRadiusGateway == nil {
		return nil, false
	}
	return o.AllowRadiusGateway, true
}

// HasAllowRadiusGateway returns a boolean if a field has been set.
func (o *LicenseGateways) HasAllowRadiusGateway() bool {
	if o != nil && o.AllowRadiusGateway != nil {
		return true
	}

	return false
}

// SetAllowRadiusGateway gets a reference to the given bool and assigns it to the AllowRadiusGateway field.
func (o *LicenseGateways) SetAllowRadiusGateway(v bool) {
	o.AllowRadiusGateway = &v
}

func (o LicenseGateways) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowLdapGateway != nil {
		toSerialize["allowLdapGateway"] = o.AllowLdapGateway
	}
	if o.AllowKerberosGateway != nil {
		toSerialize["allowKerberosGateway"] = o.AllowKerberosGateway
	}
	if o.AllowRadiusGateway != nil {
		toSerialize["allowRadiusGateway"] = o.AllowRadiusGateway
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseGateways struct {
	value *LicenseGateways
	isSet bool
}

func (v NullableLicenseGateways) Get() *LicenseGateways {
	return v.value
}

func (v *NullableLicenseGateways) Set(val *LicenseGateways) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseGateways) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseGateways) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseGateways(val *LicenseGateways) *NullableLicenseGateways {
	return &NullableLicenseGateways{value: val, isSet: true}
}

func (v NullableLicenseGateways) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseGateways) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


