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

// AgreementLanguageRevision struct for AgreementLanguageRevision
type AgreementLanguageRevision struct {
	Agreement *AgreementLanguageAgreement `json:"agreement,omitempty"`
	// An immutable string that specifies the content type of text. Options are text/html and text/plain, as defined by rfc-6838 and Media Types/text. This attribute is supported in POST requests only.
	ContentType *string `json:"contentType,omitempty"`
	// A date that specifies the start date that the revision is presented to users. This property value can be modified only if the current value is a date that has not already passed. The effective date must be unique for each language agreement, and the property value can be the present date or a future date only.
	EffectiveAt *string `json:"effectiveAt,omitempty"`
	// A read-only string that specifies the revision ID.
	Id *string `json:"id,omitempty"`
	Language *AgreementLanguageRevisionLanguage `json:"language,omitempty"`
	// A date that specifies whether the revision is still valid in the context of all revisions for a language. This property is calculated dynamically at read time, taking into consideration the agreement language, the language enabled property, and the agreement enabled property. When a new revision is added, the notValidAfter property values for all other previous revisions might be impacted. For example, if a new revision becomes effective and it forces reconsent, then all older revisions are no longer valid.
	NotValidAfter *string `json:"notValidAfter,omitempty"`
	// A boolean that specifies whether the user is required to provide consent to the language revision after it becomes effective.
	RequiresReconsent *bool `json:"requiresReconsent,omitempty"`
	// An immutable string that specifies text or HTML for the revision. This attribute is supported in POST requests only. For more information, see contentType.
	Text *string `json:"text,omitempty"`
}

// NewAgreementLanguageRevision instantiates a new AgreementLanguageRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementLanguageRevision() *AgreementLanguageRevision {
	this := AgreementLanguageRevision{}
	return &this
}

// NewAgreementLanguageRevisionWithDefaults instantiates a new AgreementLanguageRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementLanguageRevisionWithDefaults() *AgreementLanguageRevision {
	this := AgreementLanguageRevision{}
	return &this
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetAgreement() AgreementLanguageAgreement {
	if o == nil || o.Agreement == nil {
		var ret AgreementLanguageAgreement
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetAgreementOk() (*AgreementLanguageAgreement, bool) {
	if o == nil || o.Agreement == nil {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasAgreement() bool {
	if o != nil && o.Agreement != nil {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given AgreementLanguageAgreement and assigns it to the Agreement field.
func (o *AgreementLanguageRevision) SetAgreement(v AgreementLanguageAgreement) {
	o.Agreement = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *AgreementLanguageRevision) SetContentType(v string) {
	o.ContentType = &v
}

// GetEffectiveAt returns the EffectiveAt field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetEffectiveAt() string {
	if o == nil || o.EffectiveAt == nil {
		var ret string
		return ret
	}
	return *o.EffectiveAt
}

// GetEffectiveAtOk returns a tuple with the EffectiveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetEffectiveAtOk() (*string, bool) {
	if o == nil || o.EffectiveAt == nil {
		return nil, false
	}
	return o.EffectiveAt, true
}

// HasEffectiveAt returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasEffectiveAt() bool {
	if o != nil && o.EffectiveAt != nil {
		return true
	}

	return false
}

// SetEffectiveAt gets a reference to the given string and assigns it to the EffectiveAt field.
func (o *AgreementLanguageRevision) SetEffectiveAt(v string) {
	o.EffectiveAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgreementLanguageRevision) SetId(v string) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetLanguage() AgreementLanguageRevisionLanguage {
	if o == nil || o.Language == nil {
		var ret AgreementLanguageRevisionLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetLanguageOk() (*AgreementLanguageRevisionLanguage, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given AgreementLanguageRevisionLanguage and assigns it to the Language field.
func (o *AgreementLanguageRevision) SetLanguage(v AgreementLanguageRevisionLanguage) {
	o.Language = &v
}

// GetNotValidAfter returns the NotValidAfter field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetNotValidAfter() string {
	if o == nil || o.NotValidAfter == nil {
		var ret string
		return ret
	}
	return *o.NotValidAfter
}

// GetNotValidAfterOk returns a tuple with the NotValidAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetNotValidAfterOk() (*string, bool) {
	if o == nil || o.NotValidAfter == nil {
		return nil, false
	}
	return o.NotValidAfter, true
}

// HasNotValidAfter returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasNotValidAfter() bool {
	if o != nil && o.NotValidAfter != nil {
		return true
	}

	return false
}

// SetNotValidAfter gets a reference to the given string and assigns it to the NotValidAfter field.
func (o *AgreementLanguageRevision) SetNotValidAfter(v string) {
	o.NotValidAfter = &v
}

// GetRequiresReconsent returns the RequiresReconsent field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetRequiresReconsent() bool {
	if o == nil || o.RequiresReconsent == nil {
		var ret bool
		return ret
	}
	return *o.RequiresReconsent
}

// GetRequiresReconsentOk returns a tuple with the RequiresReconsent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetRequiresReconsentOk() (*bool, bool) {
	if o == nil || o.RequiresReconsent == nil {
		return nil, false
	}
	return o.RequiresReconsent, true
}

// HasRequiresReconsent returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasRequiresReconsent() bool {
	if o != nil && o.RequiresReconsent != nil {
		return true
	}

	return false
}

// SetRequiresReconsent gets a reference to the given bool and assigns it to the RequiresReconsent field.
func (o *AgreementLanguageRevision) SetRequiresReconsent(v bool) {
	o.RequiresReconsent = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AgreementLanguageRevision) SetText(v string) {
	o.Text = &v
}

func (o AgreementLanguageRevision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Agreement != nil {
		toSerialize["agreement"] = o.Agreement
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.EffectiveAt != nil {
		toSerialize["effectiveAt"] = o.EffectiveAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.NotValidAfter != nil {
		toSerialize["notValidAfter"] = o.NotValidAfter
	}
	if o.RequiresReconsent != nil {
		toSerialize["requiresReconsent"] = o.RequiresReconsent
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableAgreementLanguageRevision struct {
	value *AgreementLanguageRevision
	isSet bool
}

func (v NullableAgreementLanguageRevision) Get() *AgreementLanguageRevision {
	return v.value
}

func (v *NullableAgreementLanguageRevision) Set(val *AgreementLanguageRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementLanguageRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementLanguageRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementLanguageRevision(val *AgreementLanguageRevision) *NullableAgreementLanguageRevision {
	return &NullableAgreementLanguageRevision{value: val, isSet: true}
}

func (v NullableAgreementLanguageRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementLanguageRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


