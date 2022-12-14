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

// EmailDomain struct for EmailDomain
type EmailDomain struct {
	// A string that specifies the auto-generated ID of the email domain.
	Id *string `json:"id,omitempty"`
	// A string that specifies the resource name, which must be provided and must be unique within an environment (for example, auth.shopco.com). This is a required property. Wildcards are NOT supported.
	DomainName string `json:"domainName"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
}

// NewEmailDomain instantiates a new EmailDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomain(domainName string) *EmailDomain {
	this := EmailDomain{}
	this.DomainName = domainName
	return &this
}

// NewEmailDomainWithDefaults instantiates a new EmailDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainWithDefaults() *EmailDomain {
	this := EmailDomain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailDomain) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailDomain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailDomain) SetId(v string) {
	o.Id = &v
}

// GetDomainName returns the DomainName field value
func (o *EmailDomain) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *EmailDomain) SetDomainName(v string) {
	o.DomainName = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EmailDomain) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EmailDomain) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *EmailDomain) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

func (o EmailDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["domainName"] = o.DomainName
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableEmailDomain struct {
	value *EmailDomain
	isSet bool
}

func (v NullableEmailDomain) Get() *EmailDomain {
	return v.value
}

func (v *NullableEmailDomain) Set(val *EmailDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomain(val *EmailDomain) *NullableEmailDomain {
	return &NullableEmailDomain{value: val, isSet: true}
}

func (v NullableEmailDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


