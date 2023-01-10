/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// NotificationsSettingsPhoneDeliverySettings struct for NotificationsSettingsPhoneDeliverySettings
type NotificationsSettingsPhoneDeliverySettings struct {
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Provider EnumNotificationsSettingsPhoneDeliverySettingsProvider `json:"provider"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewNotificationsSettingsPhoneDeliverySettings instantiates a new NotificationsSettingsPhoneDeliverySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsPhoneDeliverySettings(provider EnumNotificationsSettingsPhoneDeliverySettingsProvider) *NotificationsSettingsPhoneDeliverySettings {
	this := NotificationsSettingsPhoneDeliverySettings{}
	this.Provider = provider
	return &this
}

// NewNotificationsSettingsPhoneDeliverySettingsWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsPhoneDeliverySettingsWithDefaults() *NotificationsSettingsPhoneDeliverySettings {
	this := NotificationsSettingsPhoneDeliverySettings{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettings) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationsSettingsPhoneDeliverySettings) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettings) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *NotificationsSettingsPhoneDeliverySettings) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetProvider returns the Provider field value
func (o *NotificationsSettingsPhoneDeliverySettings) GetProvider() EnumNotificationsSettingsPhoneDeliverySettingsProvider {
	if o == nil {
		var ret EnumNotificationsSettingsPhoneDeliverySettingsProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) GetProviderOk() (*EnumNotificationsSettingsPhoneDeliverySettingsProvider, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *NotificationsSettingsPhoneDeliverySettings) SetProvider(v EnumNotificationsSettingsPhoneDeliverySettingsProvider) {
	o.Provider = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettings) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationsSettingsPhoneDeliverySettings) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettings) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettings) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NotificationsSettingsPhoneDeliverySettings) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o NotificationsSettingsPhoneDeliverySettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationsSettingsPhoneDeliverySettings struct {
	value *NotificationsSettingsPhoneDeliverySettings
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettings) Get() *NotificationsSettingsPhoneDeliverySettings {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettings) Set(val *NotificationsSettingsPhoneDeliverySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettings(val *NotificationsSettingsPhoneDeliverySettings) *NullableNotificationsSettingsPhoneDeliverySettings {
	return &NullableNotificationsSettingsPhoneDeliverySettings{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


