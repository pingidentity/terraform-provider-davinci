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

// IdentityProviderAttribute struct for IdentityProviderAttribute
type IdentityProviderAttribute struct {
	MappingType *EnumIdentityProviderAttributeMappingType `json:"mappingType,omitempty"`
	// The user attribute, which is unique per provider. The attribute must not be defined as read only from the user schema or of type COMPLEX based on the user schema. Valid examples username, and name.first. The following attributes may not be used account, id, created, updated, lifecycle, mfaEnabled, and enabled.
	Name string `json:"name"`
	// A placeholder referring to the attribute (or attributes) from the provider. Placeholders must be valid for the attributes returned by the IdP type and use the ${} syntax (for example, username=`${email}`). For SAML, any placeholder is acceptable, and it is mapped against the attributes available in the SAML assertion after authentication. The ${samlAssertion.subject} placeholder is a special reserved placeholder used to refer to the subject name ID in the SAML assertion response.
	Value string `json:"value"`
	Update EnumIdentityProviderAttributeMappingUpdate `json:"update"`
	// The unique identifier for the resource.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	IdentityProvider *IdentityProviderAttributeIdentityProvider `json:"identityProvider,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewIdentityProviderAttribute instantiates a new IdentityProviderAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderAttribute(name string, value string, update EnumIdentityProviderAttributeMappingUpdate) *IdentityProviderAttribute {
	this := IdentityProviderAttribute{}
	this.Name = name
	this.Value = value
	this.Update = update
	return &this
}

// NewIdentityProviderAttributeWithDefaults instantiates a new IdentityProviderAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderAttributeWithDefaults() *IdentityProviderAttribute {
	this := IdentityProviderAttribute{}
	return &this
}

// GetMappingType returns the MappingType field value if set, zero value otherwise.
func (o *IdentityProviderAttribute) GetMappingType() EnumIdentityProviderAttributeMappingType {
	if o == nil || o.MappingType == nil {
		var ret EnumIdentityProviderAttributeMappingType
		return ret
	}
	return *o.MappingType
}

// GetMappingTypeOk returns a tuple with the MappingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetMappingTypeOk() (*EnumIdentityProviderAttributeMappingType, bool) {
	if o == nil || o.MappingType == nil {
		return nil, false
	}
	return o.MappingType, true
}

// HasMappingType returns a boolean if a field has been set.
func (o *IdentityProviderAttribute) HasMappingType() bool {
	if o != nil && o.MappingType != nil {
		return true
	}

	return false
}

// SetMappingType gets a reference to the given EnumIdentityProviderAttributeMappingType and assigns it to the MappingType field.
func (o *IdentityProviderAttribute) SetMappingType(v EnumIdentityProviderAttributeMappingType) {
	o.MappingType = &v
}

// GetName returns the Name field value
func (o *IdentityProviderAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityProviderAttribute) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *IdentityProviderAttribute) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IdentityProviderAttribute) SetValue(v string) {
	o.Value = v
}

// GetUpdate returns the Update field value
func (o *IdentityProviderAttribute) GetUpdate() EnumIdentityProviderAttributeMappingUpdate {
	if o == nil {
		var ret EnumIdentityProviderAttributeMappingUpdate
		return ret
	}

	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetUpdateOk() (*EnumIdentityProviderAttributeMappingUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Update, true
}

// SetUpdate sets field value
func (o *IdentityProviderAttribute) SetUpdate(v EnumIdentityProviderAttributeMappingUpdate) {
	o.Update = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderAttribute) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderAttribute) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProviderAttribute) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *IdentityProviderAttribute) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *IdentityProviderAttribute) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *IdentityProviderAttribute) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise.
func (o *IdentityProviderAttribute) GetIdentityProvider() IdentityProviderAttributeIdentityProvider {
	if o == nil || o.IdentityProvider == nil {
		var ret IdentityProviderAttributeIdentityProvider
		return ret
	}
	return *o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetIdentityProviderOk() (*IdentityProviderAttributeIdentityProvider, bool) {
	if o == nil || o.IdentityProvider == nil {
		return nil, false
	}
	return o.IdentityProvider, true
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *IdentityProviderAttribute) HasIdentityProvider() bool {
	if o != nil && o.IdentityProvider != nil {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given IdentityProviderAttributeIdentityProvider and assigns it to the IdentityProvider field.
func (o *IdentityProviderAttribute) SetIdentityProvider(v IdentityProviderAttributeIdentityProvider) {
	o.IdentityProvider = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityProviderAttribute) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityProviderAttribute) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *IdentityProviderAttribute) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityProviderAttribute) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttribute) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityProviderAttribute) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *IdentityProviderAttribute) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o IdentityProviderAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MappingType != nil {
		toSerialize["mappingType"] = o.MappingType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["update"] = o.Update
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.IdentityProvider != nil {
		toSerialize["identityProvider"] = o.IdentityProvider
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProviderAttribute struct {
	value *IdentityProviderAttribute
	isSet bool
}

func (v NullableIdentityProviderAttribute) Get() *IdentityProviderAttribute {
	return v.value
}

func (v *NullableIdentityProviderAttribute) Set(val *IdentityProviderAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderAttribute(val *IdentityProviderAttribute) *NullableIdentityProviderAttribute {
	return &NullableIdentityProviderAttribute{value: val, isSet: true}
}

func (v NullableIdentityProviderAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


