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

// AgreementLanguageUserExperience struct for AgreementLanguageUserExperience
type AgreementLanguageUserExperience struct {
	// A string that specifies the text next to the \"accept\" checkbox in the end user interface. Accepted character are unicode letters, combining marks, numeric characters, whitespace, and punctuation characters (regex `^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}]+$`).
	AcceptCheckboxText *string `json:"acceptCheckboxText,omitempty"`
	// A string that specifies the text of the \"continue\" button in the end user interface. Accepted character are unicode letters, combining marks, numeric characters, whitespace, and punctuation characters (regex `^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}]+$`).
	ContinueButtonText *string `json:"continueButtonText,omitempty"`
	// A string that specifies the text of the \"decline\" button in the end user interface. Accepted character are unicode letters, combining marks, numeric characters, whitespace, and punctuation characters (regex `^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}]+$`).
	DeclineButtonText *string `json:"declineButtonText,omitempty"`
}

// NewAgreementLanguageUserExperience instantiates a new AgreementLanguageUserExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementLanguageUserExperience() *AgreementLanguageUserExperience {
	this := AgreementLanguageUserExperience{}
	return &this
}

// NewAgreementLanguageUserExperienceWithDefaults instantiates a new AgreementLanguageUserExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementLanguageUserExperienceWithDefaults() *AgreementLanguageUserExperience {
	this := AgreementLanguageUserExperience{}
	return &this
}

// GetAcceptCheckboxText returns the AcceptCheckboxText field value if set, zero value otherwise.
func (o *AgreementLanguageUserExperience) GetAcceptCheckboxText() string {
	if o == nil || o.AcceptCheckboxText == nil {
		var ret string
		return ret
	}
	return *o.AcceptCheckboxText
}

// GetAcceptCheckboxTextOk returns a tuple with the AcceptCheckboxText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageUserExperience) GetAcceptCheckboxTextOk() (*string, bool) {
	if o == nil || o.AcceptCheckboxText == nil {
		return nil, false
	}
	return o.AcceptCheckboxText, true
}

// HasAcceptCheckboxText returns a boolean if a field has been set.
func (o *AgreementLanguageUserExperience) HasAcceptCheckboxText() bool {
	if o != nil && o.AcceptCheckboxText != nil {
		return true
	}

	return false
}

// SetAcceptCheckboxText gets a reference to the given string and assigns it to the AcceptCheckboxText field.
func (o *AgreementLanguageUserExperience) SetAcceptCheckboxText(v string) {
	o.AcceptCheckboxText = &v
}

// GetContinueButtonText returns the ContinueButtonText field value if set, zero value otherwise.
func (o *AgreementLanguageUserExperience) GetContinueButtonText() string {
	if o == nil || o.ContinueButtonText == nil {
		var ret string
		return ret
	}
	return *o.ContinueButtonText
}

// GetContinueButtonTextOk returns a tuple with the ContinueButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageUserExperience) GetContinueButtonTextOk() (*string, bool) {
	if o == nil || o.ContinueButtonText == nil {
		return nil, false
	}
	return o.ContinueButtonText, true
}

// HasContinueButtonText returns a boolean if a field has been set.
func (o *AgreementLanguageUserExperience) HasContinueButtonText() bool {
	if o != nil && o.ContinueButtonText != nil {
		return true
	}

	return false
}

// SetContinueButtonText gets a reference to the given string and assigns it to the ContinueButtonText field.
func (o *AgreementLanguageUserExperience) SetContinueButtonText(v string) {
	o.ContinueButtonText = &v
}

// GetDeclineButtonText returns the DeclineButtonText field value if set, zero value otherwise.
func (o *AgreementLanguageUserExperience) GetDeclineButtonText() string {
	if o == nil || o.DeclineButtonText == nil {
		var ret string
		return ret
	}
	return *o.DeclineButtonText
}

// GetDeclineButtonTextOk returns a tuple with the DeclineButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageUserExperience) GetDeclineButtonTextOk() (*string, bool) {
	if o == nil || o.DeclineButtonText == nil {
		return nil, false
	}
	return o.DeclineButtonText, true
}

// HasDeclineButtonText returns a boolean if a field has been set.
func (o *AgreementLanguageUserExperience) HasDeclineButtonText() bool {
	if o != nil && o.DeclineButtonText != nil {
		return true
	}

	return false
}

// SetDeclineButtonText gets a reference to the given string and assigns it to the DeclineButtonText field.
func (o *AgreementLanguageUserExperience) SetDeclineButtonText(v string) {
	o.DeclineButtonText = &v
}

func (o AgreementLanguageUserExperience) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptCheckboxText != nil {
		toSerialize["acceptCheckboxText"] = o.AcceptCheckboxText
	}
	if o.ContinueButtonText != nil {
		toSerialize["continueButtonText"] = o.ContinueButtonText
	}
	if o.DeclineButtonText != nil {
		toSerialize["declineButtonText"] = o.DeclineButtonText
	}
	return json.Marshal(toSerialize)
}

type NullableAgreementLanguageUserExperience struct {
	value *AgreementLanguageUserExperience
	isSet bool
}

func (v NullableAgreementLanguageUserExperience) Get() *AgreementLanguageUserExperience {
	return v.value
}

func (v *NullableAgreementLanguageUserExperience) Set(val *AgreementLanguageUserExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementLanguageUserExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementLanguageUserExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementLanguageUserExperience(val *AgreementLanguageUserExperience) *NullableAgreementLanguageUserExperience {
	return &NullableAgreementLanguageUserExperience{value: val, isSet: true}
}

func (v NullableAgreementLanguageUserExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementLanguageUserExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


