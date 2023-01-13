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

// IdentityProviderCommon struct for IdentityProviderCommon
type IdentityProviderCommon struct {
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
}

// NewIdentityProviderCommon instantiates a new IdentityProviderCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderCommon(enabled bool, name string, type_ EnumIdentityProviderExt) *IdentityProviderCommon {
	this := IdentityProviderCommon{}
	this.Enabled = enabled
	this.Name = name
	this.Type = type_
	return &this
}

// NewIdentityProviderCommonWithDefaults instantiates a new IdentityProviderCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderCommonWithDefaults() *IdentityProviderCommon {
	this := IdentityProviderCommon{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetLinks() map[string]interface{} {
	if o == nil || isNil(o.Links) {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Links) {
    return map[string]interface{}{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *IdentityProviderCommon) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityProviderCommon) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *IdentityProviderCommon) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *IdentityProviderCommon) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *IdentityProviderCommon) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetIcon() IdentityProviderCommonIcon {
	if o == nil || isNil(o.Icon) {
		var ret IdentityProviderCommonIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetIconOk() (*IdentityProviderCommonIcon, bool) {
	if o == nil || isNil(o.Icon) {
    return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasIcon() bool {
	if o != nil && !isNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given IdentityProviderCommonIcon and assigns it to the Icon field.
func (o *IdentityProviderCommon) SetIcon(v IdentityProviderCommonIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProviderCommon) SetId(v string) {
	o.Id = &v
}

// GetLoginButtonIcon returns the LoginButtonIcon field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetLoginButtonIcon() IdentityProviderCommonLoginButtonIcon {
	if o == nil || isNil(o.LoginButtonIcon) {
		var ret IdentityProviderCommonLoginButtonIcon
		return ret
	}
	return *o.LoginButtonIcon
}

// GetLoginButtonIconOk returns a tuple with the LoginButtonIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetLoginButtonIconOk() (*IdentityProviderCommonLoginButtonIcon, bool) {
	if o == nil || isNil(o.LoginButtonIcon) {
    return nil, false
	}
	return o.LoginButtonIcon, true
}

// HasLoginButtonIcon returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasLoginButtonIcon() bool {
	if o != nil && !isNil(o.LoginButtonIcon) {
		return true
	}

	return false
}

// SetLoginButtonIcon gets a reference to the given IdentityProviderCommonLoginButtonIcon and assigns it to the LoginButtonIcon field.
func (o *IdentityProviderCommon) SetLoginButtonIcon(v IdentityProviderCommonLoginButtonIcon) {
	o.LoginButtonIcon = &v
}

// GetName returns the Name field value
func (o *IdentityProviderCommon) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityProviderCommon) SetName(v string) {
	o.Name = v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetRegistration() IdentityProviderCommonRegistration {
	if o == nil || isNil(o.Registration) {
		var ret IdentityProviderCommonRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetRegistrationOk() (*IdentityProviderCommonRegistration, bool) {
	if o == nil || isNil(o.Registration) {
    return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasRegistration() bool {
	if o != nil && !isNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given IdentityProviderCommonRegistration and assigns it to the Registration field.
func (o *IdentityProviderCommon) SetRegistration(v IdentityProviderCommonRegistration) {
	o.Registration = &v
}

// GetType returns the Type field value
func (o *IdentityProviderCommon) GetType() EnumIdentityProviderExt {
	if o == nil {
		var ret EnumIdentityProviderExt
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetTypeOk() (*EnumIdentityProviderExt, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdentityProviderCommon) SetType(v EnumIdentityProviderExt) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *IdentityProviderCommon) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityProviderCommon) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommon) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityProviderCommon) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *IdentityProviderCommon) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o IdentityProviderCommon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LoginButtonIcon) {
		toSerialize["loginButtonIcon"] = o.LoginButtonIcon
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProviderCommon struct {
	value *IdentityProviderCommon
	isSet bool
}

func (v NullableIdentityProviderCommon) Get() *IdentityProviderCommon {
	return v.value
}

func (v *NullableIdentityProviderCommon) Set(val *IdentityProviderCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCommon(val *IdentityProviderCommon) *NullableIdentityProviderCommon {
	return &NullableIdentityProviderCommon{value: val, isSet: true}
}

func (v NullableIdentityProviderCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


