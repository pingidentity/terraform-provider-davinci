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

// Language struct for Language
type Language struct {
	// Specifies whether this language is the default for the environment. This property value must be set to false when creating a language resource. It can be set to true only after the language is enabled and after the localization of an agreement resource is complete when agreements are used for the environment.
	Default bool `json:"default"`
	// Specifies whether this language is enabled for the environment. This property value must be set to false when creating a language.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An ISO standard language code. For more information about standard language codes, see ISO Language Code Table.
	Locale string `json:"locale"`
	// The language name. If omitted, the ISO standard name is used.
	Name *string `json:"name,omitempty"`
	// The time the language resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// Specifies whether this language was added by a customer administrator.
	CustomerAdded *bool `json:"customerAdded,omitempty"`
	// The time the language resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewLanguage instantiates a new Language object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguage(default_ bool, enabled bool, locale string) *Language {
	this := Language{}
	this.Default = default_
	this.Enabled = enabled
	this.Locale = locale
	return &this
}

// NewLanguageWithDefaults instantiates a new Language object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageWithDefaults() *Language {
	this := Language{}
	return &this
}

// GetDefault returns the Default field value
func (o *Language) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *Language) GetDefaultOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *Language) SetDefault(v bool) {
	o.Default = v
}

// GetEnabled returns the Enabled field value
func (o *Language) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Language) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Language) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Language) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Language) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Language) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Language) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Language) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Language) SetId(v string) {
	o.Id = &v
}

// GetLocale returns the Locale field value
func (o *Language) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *Language) GetLocaleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *Language) SetLocale(v string) {
	o.Locale = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Language) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Language) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Language) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Language) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Language) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Language) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCustomerAdded returns the CustomerAdded field value if set, zero value otherwise.
func (o *Language) GetCustomerAdded() bool {
	if o == nil || isNil(o.CustomerAdded) {
		var ret bool
		return ret
	}
	return *o.CustomerAdded
}

// GetCustomerAddedOk returns a tuple with the CustomerAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetCustomerAddedOk() (*bool, bool) {
	if o == nil || isNil(o.CustomerAdded) {
    return nil, false
	}
	return o.CustomerAdded, true
}

// HasCustomerAdded returns a boolean if a field has been set.
func (o *Language) HasCustomerAdded() bool {
	if o != nil && !isNil(o.CustomerAdded) {
		return true
	}

	return false
}

// SetCustomerAdded gets a reference to the given bool and assigns it to the CustomerAdded field.
func (o *Language) SetCustomerAdded(v bool) {
	o.CustomerAdded = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Language) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Language) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Language) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Language) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["default"] = o.Default
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["locale"] = o.Locale
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CustomerAdded) {
		toSerialize["customerAdded"] = o.CustomerAdded
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableLanguage struct {
	value *Language
	isSet bool
}

func (v NullableLanguage) Get() *Language {
	return v.value
}

func (v *NullableLanguage) Set(val *Language) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguage(val *Language) *NullableLanguage {
	return &NullableLanguage{value: val, isSet: true}
}

func (v NullableLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


