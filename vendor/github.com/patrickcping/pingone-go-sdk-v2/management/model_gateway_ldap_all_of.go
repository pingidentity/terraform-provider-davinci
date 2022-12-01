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

// GatewayLDAPAllOf struct for GatewayLDAPAllOf
type GatewayLDAPAllOf struct {
	// A string that specifies the distinguished name information to bind to the LDAP database (for example, uid=pingone,dc=example,dc=com).
	BindDN string `json:"bindDN"`
	// A string that specifies the bind password for the LDAP database. This is a required property.
	BindPassword string `json:"bindPassword"`
	ConnectionSecurity *EnumGatewayLDAPSecurity `json:"connectionSecurity,omitempty"`
	Kerberos *GatewayLDAPAllOfKerberos `json:"kerberos,omitempty"`
	// An array of strings that specifies the LDAP server host name and port number (for example, [`ds1.example.com:389`, `ds2.example.com:389`]).
	ServersHostAndPort []string `json:"serversHostAndPort"`
	// An array of the userTypes properties for the users to be provisioned in PingOne. userTypes specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory.
	UserTypes []GatewayLDAPAllOfUserTypes `json:"userTypes,omitempty"`
	// A boolean that specifies whether or not to trust all SSL certificates (defaults to true). If this value is false, TLS certificates are not validated. When the value is set to true, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted.
	ValidateTlsCertificates *bool `json:"validateTlsCertificates,omitempty"`
	Vendor EnumGatewayVendor `json:"vendor"`
	FollowReferrals *bool `json:"followReferrals,omitempty"`
}

// NewGatewayLDAPAllOf instantiates a new GatewayLDAPAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayLDAPAllOf(bindDN string, bindPassword string, serversHostAndPort []string, vendor EnumGatewayVendor) *GatewayLDAPAllOf {
	this := GatewayLDAPAllOf{}
	this.BindDN = bindDN
	this.BindPassword = bindPassword
	this.ServersHostAndPort = serversHostAndPort
	this.Vendor = vendor
	return &this
}

// NewGatewayLDAPAllOfWithDefaults instantiates a new GatewayLDAPAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayLDAPAllOfWithDefaults() *GatewayLDAPAllOf {
	this := GatewayLDAPAllOf{}
	return &this
}

// GetBindDN returns the BindDN field value
func (o *GatewayLDAPAllOf) GetBindDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindDN
}

// GetBindDNOk returns a tuple with the BindDN field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetBindDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindDN, true
}

// SetBindDN sets field value
func (o *GatewayLDAPAllOf) SetBindDN(v string) {
	o.BindDN = v
}

// GetBindPassword returns the BindPassword field value
func (o *GatewayLDAPAllOf) GetBindPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindPassword
}

// GetBindPasswordOk returns a tuple with the BindPassword field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetBindPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindPassword, true
}

// SetBindPassword sets field value
func (o *GatewayLDAPAllOf) SetBindPassword(v string) {
	o.BindPassword = v
}

// GetConnectionSecurity returns the ConnectionSecurity field value if set, zero value otherwise.
func (o *GatewayLDAPAllOf) GetConnectionSecurity() EnumGatewayLDAPSecurity {
	if o == nil || o.ConnectionSecurity == nil {
		var ret EnumGatewayLDAPSecurity
		return ret
	}
	return *o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetConnectionSecurityOk() (*EnumGatewayLDAPSecurity, bool) {
	if o == nil || o.ConnectionSecurity == nil {
		return nil, false
	}
	return o.ConnectionSecurity, true
}

// HasConnectionSecurity returns a boolean if a field has been set.
func (o *GatewayLDAPAllOf) HasConnectionSecurity() bool {
	if o != nil && o.ConnectionSecurity != nil {
		return true
	}

	return false
}

// SetConnectionSecurity gets a reference to the given EnumGatewayLDAPSecurity and assigns it to the ConnectionSecurity field.
func (o *GatewayLDAPAllOf) SetConnectionSecurity(v EnumGatewayLDAPSecurity) {
	o.ConnectionSecurity = &v
}

// GetKerberos returns the Kerberos field value if set, zero value otherwise.
func (o *GatewayLDAPAllOf) GetKerberos() GatewayLDAPAllOfKerberos {
	if o == nil || o.Kerberos == nil {
		var ret GatewayLDAPAllOfKerberos
		return ret
	}
	return *o.Kerberos
}

// GetKerberosOk returns a tuple with the Kerberos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetKerberosOk() (*GatewayLDAPAllOfKerberos, bool) {
	if o == nil || o.Kerberos == nil {
		return nil, false
	}
	return o.Kerberos, true
}

// HasKerberos returns a boolean if a field has been set.
func (o *GatewayLDAPAllOf) HasKerberos() bool {
	if o != nil && o.Kerberos != nil {
		return true
	}

	return false
}

// SetKerberos gets a reference to the given GatewayLDAPAllOfKerberos and assigns it to the Kerberos field.
func (o *GatewayLDAPAllOf) SetKerberos(v GatewayLDAPAllOfKerberos) {
	o.Kerberos = &v
}

// GetServersHostAndPort returns the ServersHostAndPort field value
func (o *GatewayLDAPAllOf) GetServersHostAndPort() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServersHostAndPort
}

// GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetServersHostAndPortOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServersHostAndPort, true
}

// SetServersHostAndPort sets field value
func (o *GatewayLDAPAllOf) SetServersHostAndPort(v []string) {
	o.ServersHostAndPort = v
}

// GetUserTypes returns the UserTypes field value if set, zero value otherwise.
func (o *GatewayLDAPAllOf) GetUserTypes() []GatewayLDAPAllOfUserTypes {
	if o == nil || o.UserTypes == nil {
		var ret []GatewayLDAPAllOfUserTypes
		return ret
	}
	return o.UserTypes
}

// GetUserTypesOk returns a tuple with the UserTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetUserTypesOk() ([]GatewayLDAPAllOfUserTypes, bool) {
	if o == nil || o.UserTypes == nil {
		return nil, false
	}
	return o.UserTypes, true
}

// HasUserTypes returns a boolean if a field has been set.
func (o *GatewayLDAPAllOf) HasUserTypes() bool {
	if o != nil && o.UserTypes != nil {
		return true
	}

	return false
}

// SetUserTypes gets a reference to the given []GatewayLDAPAllOfUserTypes and assigns it to the UserTypes field.
func (o *GatewayLDAPAllOf) SetUserTypes(v []GatewayLDAPAllOfUserTypes) {
	o.UserTypes = v
}

// GetValidateTlsCertificates returns the ValidateTlsCertificates field value if set, zero value otherwise.
func (o *GatewayLDAPAllOf) GetValidateTlsCertificates() bool {
	if o == nil || o.ValidateTlsCertificates == nil {
		var ret bool
		return ret
	}
	return *o.ValidateTlsCertificates
}

// GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetValidateTlsCertificatesOk() (*bool, bool) {
	if o == nil || o.ValidateTlsCertificates == nil {
		return nil, false
	}
	return o.ValidateTlsCertificates, true
}

// HasValidateTlsCertificates returns a boolean if a field has been set.
func (o *GatewayLDAPAllOf) HasValidateTlsCertificates() bool {
	if o != nil && o.ValidateTlsCertificates != nil {
		return true
	}

	return false
}

// SetValidateTlsCertificates gets a reference to the given bool and assigns it to the ValidateTlsCertificates field.
func (o *GatewayLDAPAllOf) SetValidateTlsCertificates(v bool) {
	o.ValidateTlsCertificates = &v
}

// GetVendor returns the Vendor field value
func (o *GatewayLDAPAllOf) GetVendor() EnumGatewayVendor {
	if o == nil {
		var ret EnumGatewayVendor
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetVendorOk() (*EnumGatewayVendor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *GatewayLDAPAllOf) SetVendor(v EnumGatewayVendor) {
	o.Vendor = v
}

// GetFollowReferrals returns the FollowReferrals field value if set, zero value otherwise.
func (o *GatewayLDAPAllOf) GetFollowReferrals() bool {
	if o == nil || o.FollowReferrals == nil {
		var ret bool
		return ret
	}
	return *o.FollowReferrals
}

// GetFollowReferralsOk returns a tuple with the FollowReferrals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOf) GetFollowReferralsOk() (*bool, bool) {
	if o == nil || o.FollowReferrals == nil {
		return nil, false
	}
	return o.FollowReferrals, true
}

// HasFollowReferrals returns a boolean if a field has been set.
func (o *GatewayLDAPAllOf) HasFollowReferrals() bool {
	if o != nil && o.FollowReferrals != nil {
		return true
	}

	return false
}

// SetFollowReferrals gets a reference to the given bool and assigns it to the FollowReferrals field.
func (o *GatewayLDAPAllOf) SetFollowReferrals(v bool) {
	o.FollowReferrals = &v
}

func (o GatewayLDAPAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bindDN"] = o.BindDN
	}
	if true {
		toSerialize["bindPassword"] = o.BindPassword
	}
	if o.ConnectionSecurity != nil {
		toSerialize["connectionSecurity"] = o.ConnectionSecurity
	}
	if o.Kerberos != nil {
		toSerialize["kerberos"] = o.Kerberos
	}
	if true {
		toSerialize["serversHostAndPort"] = o.ServersHostAndPort
	}
	if o.UserTypes != nil {
		toSerialize["userTypes"] = o.UserTypes
	}
	if o.ValidateTlsCertificates != nil {
		toSerialize["validateTlsCertificates"] = o.ValidateTlsCertificates
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	if o.FollowReferrals != nil {
		toSerialize["followReferrals"] = o.FollowReferrals
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayLDAPAllOf struct {
	value *GatewayLDAPAllOf
	isSet bool
}

func (v NullableGatewayLDAPAllOf) Get() *GatewayLDAPAllOf {
	return v.value
}

func (v *NullableGatewayLDAPAllOf) Set(val *GatewayLDAPAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayLDAPAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayLDAPAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayLDAPAllOf(val *GatewayLDAPAllOf) *NullableGatewayLDAPAllOf {
	return &NullableGatewayLDAPAllOf{value: val, isSet: true}
}

func (v NullableGatewayLDAPAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayLDAPAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


