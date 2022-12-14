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

// AgreementLanguageRevisionLanguage struct for AgreementLanguageRevisionLanguage
type AgreementLanguageRevisionLanguage struct {
	// A relationship that specifies the language resource associated with this revision.
	Id *string `json:"id,omitempty"`
}

// NewAgreementLanguageRevisionLanguage instantiates a new AgreementLanguageRevisionLanguage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementLanguageRevisionLanguage() *AgreementLanguageRevisionLanguage {
	this := AgreementLanguageRevisionLanguage{}
	return &this
}

// NewAgreementLanguageRevisionLanguageWithDefaults instantiates a new AgreementLanguageRevisionLanguage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementLanguageRevisionLanguageWithDefaults() *AgreementLanguageRevisionLanguage {
	this := AgreementLanguageRevisionLanguage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgreementLanguageRevisionLanguage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevisionLanguage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgreementLanguageRevisionLanguage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgreementLanguageRevisionLanguage) SetId(v string) {
	o.Id = &v
}

func (o AgreementLanguageRevisionLanguage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAgreementLanguageRevisionLanguage struct {
	value *AgreementLanguageRevisionLanguage
	isSet bool
}

func (v NullableAgreementLanguageRevisionLanguage) Get() *AgreementLanguageRevisionLanguage {
	return v.value
}

func (v *NullableAgreementLanguageRevisionLanguage) Set(val *AgreementLanguageRevisionLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementLanguageRevisionLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementLanguageRevisionLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementLanguageRevisionLanguage(val *AgreementLanguageRevisionLanguage) *NullableAgreementLanguageRevisionLanguage {
	return &NullableAgreementLanguageRevisionLanguage{value: val, isSet: true}
}

func (v NullableAgreementLanguageRevisionLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementLanguageRevisionLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


