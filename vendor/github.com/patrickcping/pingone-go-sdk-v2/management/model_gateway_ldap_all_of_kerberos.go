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

// GatewayLDAPAllOfKerberos Contains the Kerberos authentication settings. Set this to null to disable Kerberos authentication.
type GatewayLDAPAllOfKerberos struct {
	// The password for the Kerberos service account.
	ServiceAccountPassword *string `json:"serviceAccountPassword,omitempty"`
	// The Kerberos service account user principal name (for example, `username@domain.com`).
	ServiceAccountUserPrincipalName string `json:"serviceAccountUserPrincipalName"`
	// The number of minutes for which the previous credentials are persisted.
	MinutesToRetainPreviousCredentials *int32 `json:"minutesToRetainPreviousCredentials,omitempty"`
}

// NewGatewayLDAPAllOfKerberos instantiates a new GatewayLDAPAllOfKerberos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayLDAPAllOfKerberos(serviceAccountUserPrincipalName string) *GatewayLDAPAllOfKerberos {
	this := GatewayLDAPAllOfKerberos{}
	this.ServiceAccountUserPrincipalName = serviceAccountUserPrincipalName
	return &this
}

// NewGatewayLDAPAllOfKerberosWithDefaults instantiates a new GatewayLDAPAllOfKerberos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayLDAPAllOfKerberosWithDefaults() *GatewayLDAPAllOfKerberos {
	this := GatewayLDAPAllOfKerberos{}
	return &this
}

// GetServiceAccountPassword returns the ServiceAccountPassword field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfKerberos) GetServiceAccountPassword() string {
	if o == nil || o.ServiceAccountPassword == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountPassword
}

// GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfKerberos) GetServiceAccountPasswordOk() (*string, bool) {
	if o == nil || o.ServiceAccountPassword == nil {
		return nil, false
	}
	return o.ServiceAccountPassword, true
}

// HasServiceAccountPassword returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfKerberos) HasServiceAccountPassword() bool {
	if o != nil && o.ServiceAccountPassword != nil {
		return true
	}

	return false
}

// SetServiceAccountPassword gets a reference to the given string and assigns it to the ServiceAccountPassword field.
func (o *GatewayLDAPAllOfKerberos) SetServiceAccountPassword(v string) {
	o.ServiceAccountPassword = &v
}

// GetServiceAccountUserPrincipalName returns the ServiceAccountUserPrincipalName field value
func (o *GatewayLDAPAllOfKerberos) GetServiceAccountUserPrincipalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAccountUserPrincipalName
}

// GetServiceAccountUserPrincipalNameOk returns a tuple with the ServiceAccountUserPrincipalName field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfKerberos) GetServiceAccountUserPrincipalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountUserPrincipalName, true
}

// SetServiceAccountUserPrincipalName sets field value
func (o *GatewayLDAPAllOfKerberos) SetServiceAccountUserPrincipalName(v string) {
	o.ServiceAccountUserPrincipalName = v
}

// GetMinutesToRetainPreviousCredentials returns the MinutesToRetainPreviousCredentials field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfKerberos) GetMinutesToRetainPreviousCredentials() int32 {
	if o == nil || o.MinutesToRetainPreviousCredentials == nil {
		var ret int32
		return ret
	}
	return *o.MinutesToRetainPreviousCredentials
}

// GetMinutesToRetainPreviousCredentialsOk returns a tuple with the MinutesToRetainPreviousCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfKerberos) GetMinutesToRetainPreviousCredentialsOk() (*int32, bool) {
	if o == nil || o.MinutesToRetainPreviousCredentials == nil {
		return nil, false
	}
	return o.MinutesToRetainPreviousCredentials, true
}

// HasMinutesToRetainPreviousCredentials returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfKerberos) HasMinutesToRetainPreviousCredentials() bool {
	if o != nil && o.MinutesToRetainPreviousCredentials != nil {
		return true
	}

	return false
}

// SetMinutesToRetainPreviousCredentials gets a reference to the given int32 and assigns it to the MinutesToRetainPreviousCredentials field.
func (o *GatewayLDAPAllOfKerberos) SetMinutesToRetainPreviousCredentials(v int32) {
	o.MinutesToRetainPreviousCredentials = &v
}

func (o GatewayLDAPAllOfKerberos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceAccountPassword != nil {
		toSerialize["serviceAccountPassword"] = o.ServiceAccountPassword
	}
	if true {
		toSerialize["serviceAccountUserPrincipalName"] = o.ServiceAccountUserPrincipalName
	}
	if o.MinutesToRetainPreviousCredentials != nil {
		toSerialize["minutesToRetainPreviousCredentials"] = o.MinutesToRetainPreviousCredentials
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayLDAPAllOfKerberos struct {
	value *GatewayLDAPAllOfKerberos
	isSet bool
}

func (v NullableGatewayLDAPAllOfKerberos) Get() *GatewayLDAPAllOfKerberos {
	return v.value
}

func (v *NullableGatewayLDAPAllOfKerberos) Set(val *GatewayLDAPAllOfKerberos) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayLDAPAllOfKerberos) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayLDAPAllOfKerberos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayLDAPAllOfKerberos(val *GatewayLDAPAllOfKerberos) *NullableGatewayLDAPAllOfKerberos {
	return &NullableGatewayLDAPAllOfKerberos{value: val, isSet: true}
}

func (v NullableGatewayLDAPAllOfKerberos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayLDAPAllOfKerberos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


