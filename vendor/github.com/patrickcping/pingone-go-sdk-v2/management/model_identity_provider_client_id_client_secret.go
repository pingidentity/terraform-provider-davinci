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

// IdentityProviderClientIDClientSecret struct for IdentityProviderClientIDClientSecret
type IdentityProviderClientIDClientSecret struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	// The description of the IdP.
	Description *string `json:"description,omitempty"`
	// The current enabled state of the IdP.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Icon *IdentityProviderCommonIcon `json:"icon,omitempty"`
	// The resource ID.
	Id *string `json:"id,omitempty"`
	LoginButtonIcon *IdentityProviderCommonLoginButtonIcon `json:"loginButtonIcon,omitempty"`
	// The name of the IdP.
	Name string `json:"name"`
	Registration *IdentityProviderCommonRegistration `json:"registration,omitempty"`
	Type EnumIdentityProviderExt `json:"type"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// A string that specifies the application ID from the provider. This is a required property.
	ClientId string `json:"clientId"`
	// A string that specifies the application secret from the provider. This is a required property.
	ClientSecret string `json:"clientSecret"`
}

// NewIdentityProviderClientIDClientSecret instantiates a new IdentityProviderClientIDClientSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderClientIDClientSecret(enabled bool, name string, type_ EnumIdentityProviderExt, clientId string, clientSecret string) *IdentityProviderClientIDClientSecret {
	this := IdentityProviderClientIDClientSecret{}
	this.Enabled = enabled
	this.Name = name
	this.Type = type_
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewIdentityProviderClientIDClientSecretWithDefaults instantiates a new IdentityProviderClientIDClientSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderClientIDClientSecretWithDefaults() *IdentityProviderClientIDClientSecret {
	this := IdentityProviderClientIDClientSecret{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *IdentityProviderClientIDClientSecret) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityProviderClientIDClientSecret) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *IdentityProviderClientIDClientSecret) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *IdentityProviderClientIDClientSecret) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *IdentityProviderClientIDClientSecret) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetIcon() IdentityProviderCommonIcon {
	if o == nil || o.Icon == nil {
		var ret IdentityProviderCommonIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetIconOk() (*IdentityProviderCommonIcon, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given IdentityProviderCommonIcon and assigns it to the Icon field.
func (o *IdentityProviderClientIDClientSecret) SetIcon(v IdentityProviderCommonIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProviderClientIDClientSecret) SetId(v string) {
	o.Id = &v
}

// GetLoginButtonIcon returns the LoginButtonIcon field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetLoginButtonIcon() IdentityProviderCommonLoginButtonIcon {
	if o == nil || o.LoginButtonIcon == nil {
		var ret IdentityProviderCommonLoginButtonIcon
		return ret
	}
	return *o.LoginButtonIcon
}

// GetLoginButtonIconOk returns a tuple with the LoginButtonIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetLoginButtonIconOk() (*IdentityProviderCommonLoginButtonIcon, bool) {
	if o == nil || o.LoginButtonIcon == nil {
		return nil, false
	}
	return o.LoginButtonIcon, true
}

// HasLoginButtonIcon returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasLoginButtonIcon() bool {
	if o != nil && o.LoginButtonIcon != nil {
		return true
	}

	return false
}

// SetLoginButtonIcon gets a reference to the given IdentityProviderCommonLoginButtonIcon and assigns it to the LoginButtonIcon field.
func (o *IdentityProviderClientIDClientSecret) SetLoginButtonIcon(v IdentityProviderCommonLoginButtonIcon) {
	o.LoginButtonIcon = &v
}

// GetName returns the Name field value
func (o *IdentityProviderClientIDClientSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityProviderClientIDClientSecret) SetName(v string) {
	o.Name = v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetRegistration() IdentityProviderCommonRegistration {
	if o == nil || o.Registration == nil {
		var ret IdentityProviderCommonRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetRegistrationOk() (*IdentityProviderCommonRegistration, bool) {
	if o == nil || o.Registration == nil {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasRegistration() bool {
	if o != nil && o.Registration != nil {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given IdentityProviderCommonRegistration and assigns it to the Registration field.
func (o *IdentityProviderClientIDClientSecret) SetRegistration(v IdentityProviderCommonRegistration) {
	o.Registration = &v
}

// GetType returns the Type field value
func (o *IdentityProviderClientIDClientSecret) GetType() EnumIdentityProviderExt {
	if o == nil {
		var ret EnumIdentityProviderExt
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetTypeOk() (*EnumIdentityProviderExt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdentityProviderClientIDClientSecret) SetType(v EnumIdentityProviderExt) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *IdentityProviderClientIDClientSecret) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityProviderClientIDClientSecret) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityProviderClientIDClientSecret) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *IdentityProviderClientIDClientSecret) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetClientId returns the ClientId field value
func (o *IdentityProviderClientIDClientSecret) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *IdentityProviderClientIDClientSecret) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *IdentityProviderClientIDClientSecret) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderClientIDClientSecret) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *IdentityProviderClientIDClientSecret) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o IdentityProviderClientIDClientSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LoginButtonIcon != nil {
		toSerialize["loginButtonIcon"] = o.LoginButtonIcon
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Registration != nil {
		toSerialize["registration"] = o.Registration
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProviderClientIDClientSecret struct {
	value *IdentityProviderClientIDClientSecret
	isSet bool
}

func (v NullableIdentityProviderClientIDClientSecret) Get() *IdentityProviderClientIDClientSecret {
	return v.value
}

func (v *NullableIdentityProviderClientIDClientSecret) Set(val *IdentityProviderClientIDClientSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderClientIDClientSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderClientIDClientSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderClientIDClientSecret(val *IdentityProviderClientIDClientSecret) *NullableIdentityProviderClientIDClientSecret {
	return &NullableIdentityProviderClientIDClientSecret{value: val, isSet: true}
}

func (v NullableIdentityProviderClientIDClientSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderClientIDClientSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


