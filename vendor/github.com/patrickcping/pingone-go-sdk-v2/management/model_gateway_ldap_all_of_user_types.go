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

// GatewayLDAPAllOfUserTypes struct for GatewayLDAPAllOfUserTypes
type GatewayLDAPAllOfUserTypes struct {
	// Defaults to false if this property isn't specified in the request. If false, the user cannot change the password in the remote LDAP directory. In this case, operations for forgotten passwords or resetting of passwords are not available to a user referencing this gateway.
	AllowPasswordChanges *bool `json:"allowPasswordChanges,omitempty"`
	// The UUID of the user type. This correlates to the password.external.gateway.userType.id User property.
	Id *string `json:"id,omitempty"`
	// The name of the user type.
	Name string `json:"name"`
	NewUserLookup *GatewayLDAPAllOfNewUserLookup `json:"newUserLookup,omitempty"`
	// A map of key/value entries used to persist the external LDAP directory attributes.
	OrderedCorrelationAttributes []string `json:"orderedCorrelationAttributes"`
	PasswordAuthority EnumGatewayPasswordAuthority `json:"passwordAuthority"`
	// The LDAP base domain name (DN) for this user type.
	SearchBaseDn string `json:"searchBaseDn"`
}

// NewGatewayLDAPAllOfUserTypes instantiates a new GatewayLDAPAllOfUserTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayLDAPAllOfUserTypes(name string, orderedCorrelationAttributes []string, passwordAuthority EnumGatewayPasswordAuthority, searchBaseDn string) *GatewayLDAPAllOfUserTypes {
	this := GatewayLDAPAllOfUserTypes{}
	this.Name = name
	this.OrderedCorrelationAttributes = orderedCorrelationAttributes
	this.PasswordAuthority = passwordAuthority
	this.SearchBaseDn = searchBaseDn
	return &this
}

// NewGatewayLDAPAllOfUserTypesWithDefaults instantiates a new GatewayLDAPAllOfUserTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayLDAPAllOfUserTypesWithDefaults() *GatewayLDAPAllOfUserTypes {
	this := GatewayLDAPAllOfUserTypes{}
	return &this
}

// GetAllowPasswordChanges returns the AllowPasswordChanges field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfUserTypes) GetAllowPasswordChanges() bool {
	if o == nil || o.AllowPasswordChanges == nil {
		var ret bool
		return ret
	}
	return *o.AllowPasswordChanges
}

// GetAllowPasswordChangesOk returns a tuple with the AllowPasswordChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfUserTypes) GetAllowPasswordChangesOk() (*bool, bool) {
	if o == nil || o.AllowPasswordChanges == nil {
		return nil, false
	}
	return o.AllowPasswordChanges, true
}

// HasAllowPasswordChanges returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfUserTypes) HasAllowPasswordChanges() bool {
	if o != nil && o.AllowPasswordChanges != nil {
		return true
	}

	return false
}

// SetAllowPasswordChanges gets a reference to the given bool and assigns it to the AllowPasswordChanges field.
func (o *GatewayLDAPAllOfUserTypes) SetAllowPasswordChanges(v bool) {
	o.AllowPasswordChanges = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfUserTypes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfUserTypes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfUserTypes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayLDAPAllOfUserTypes) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *GatewayLDAPAllOfUserTypes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfUserTypes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayLDAPAllOfUserTypes) SetName(v string) {
	o.Name = v
}

// GetNewUserLookup returns the NewUserLookup field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfUserTypes) GetNewUserLookup() GatewayLDAPAllOfNewUserLookup {
	if o == nil || o.NewUserLookup == nil {
		var ret GatewayLDAPAllOfNewUserLookup
		return ret
	}
	return *o.NewUserLookup
}

// GetNewUserLookupOk returns a tuple with the NewUserLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfUserTypes) GetNewUserLookupOk() (*GatewayLDAPAllOfNewUserLookup, bool) {
	if o == nil || o.NewUserLookup == nil {
		return nil, false
	}
	return o.NewUserLookup, true
}

// HasNewUserLookup returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfUserTypes) HasNewUserLookup() bool {
	if o != nil && o.NewUserLookup != nil {
		return true
	}

	return false
}

// SetNewUserLookup gets a reference to the given GatewayLDAPAllOfNewUserLookup and assigns it to the NewUserLookup field.
func (o *GatewayLDAPAllOfUserTypes) SetNewUserLookup(v GatewayLDAPAllOfNewUserLookup) {
	o.NewUserLookup = &v
}

// GetOrderedCorrelationAttributes returns the OrderedCorrelationAttributes field value
func (o *GatewayLDAPAllOfUserTypes) GetOrderedCorrelationAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrderedCorrelationAttributes
}

// GetOrderedCorrelationAttributesOk returns a tuple with the OrderedCorrelationAttributes field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfUserTypes) GetOrderedCorrelationAttributesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderedCorrelationAttributes, true
}

// SetOrderedCorrelationAttributes sets field value
func (o *GatewayLDAPAllOfUserTypes) SetOrderedCorrelationAttributes(v []string) {
	o.OrderedCorrelationAttributes = v
}

// GetPasswordAuthority returns the PasswordAuthority field value
func (o *GatewayLDAPAllOfUserTypes) GetPasswordAuthority() EnumGatewayPasswordAuthority {
	if o == nil {
		var ret EnumGatewayPasswordAuthority
		return ret
	}

	return o.PasswordAuthority
}

// GetPasswordAuthorityOk returns a tuple with the PasswordAuthority field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfUserTypes) GetPasswordAuthorityOk() (*EnumGatewayPasswordAuthority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordAuthority, true
}

// SetPasswordAuthority sets field value
func (o *GatewayLDAPAllOfUserTypes) SetPasswordAuthority(v EnumGatewayPasswordAuthority) {
	o.PasswordAuthority = v
}

// GetSearchBaseDn returns the SearchBaseDn field value
func (o *GatewayLDAPAllOfUserTypes) GetSearchBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchBaseDn
}

// GetSearchBaseDnOk returns a tuple with the SearchBaseDn field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfUserTypes) GetSearchBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchBaseDn, true
}

// SetSearchBaseDn sets field value
func (o *GatewayLDAPAllOfUserTypes) SetSearchBaseDn(v string) {
	o.SearchBaseDn = v
}

func (o GatewayLDAPAllOfUserTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowPasswordChanges != nil {
		toSerialize["allowPasswordChanges"] = o.AllowPasswordChanges
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewUserLookup != nil {
		toSerialize["newUserLookup"] = o.NewUserLookup
	}
	if true {
		toSerialize["orderedCorrelationAttributes"] = o.OrderedCorrelationAttributes
	}
	if true {
		toSerialize["passwordAuthority"] = o.PasswordAuthority
	}
	if true {
		toSerialize["searchBaseDn"] = o.SearchBaseDn
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayLDAPAllOfUserTypes struct {
	value *GatewayLDAPAllOfUserTypes
	isSet bool
}

func (v NullableGatewayLDAPAllOfUserTypes) Get() *GatewayLDAPAllOfUserTypes {
	return v.value
}

func (v *NullableGatewayLDAPAllOfUserTypes) Set(val *GatewayLDAPAllOfUserTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayLDAPAllOfUserTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayLDAPAllOfUserTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayLDAPAllOfUserTypes(val *GatewayLDAPAllOfUserTypes) *NullableGatewayLDAPAllOfUserTypes {
	return &NullableGatewayLDAPAllOfUserTypes{value: val, isSet: true}
}

func (v NullableGatewayLDAPAllOfUserTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayLDAPAllOfUserTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


