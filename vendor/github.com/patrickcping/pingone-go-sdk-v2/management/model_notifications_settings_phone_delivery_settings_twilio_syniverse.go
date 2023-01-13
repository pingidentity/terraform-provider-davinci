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

// NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse struct for NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse
type NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse struct {
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Provider EnumNotificationsSettingsPhoneDeliverySettingsProvider `json:"provider"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The public ID of the Twilio account. Relevant to Twilio only. 
	Sid string `json:"sid"`
	// The secret key of the Twilio or Syniverse account.
	AuthToken string `json:"authToken"`
	// A collection of Twilio or Syniverse numbers to use when sending a notification.
	Numbers []NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers `json:"numbers"`
}

// NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse(provider EnumNotificationsSettingsPhoneDeliverySettingsProvider, sid string, authToken string, numbers []NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse {
	this := NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse{}
	this.Provider = provider
	this.Sid = sid
	this.AuthToken = authToken
	this.Numbers = numbers
	return &this
}

// NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseWithDefaults() *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse {
	this := NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetProvider returns the Provider field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetProvider() EnumNotificationsSettingsPhoneDeliverySettingsProvider {
	if o == nil {
		var ret EnumNotificationsSettingsPhoneDeliverySettingsProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetProviderOk() (*EnumNotificationsSettingsPhoneDeliverySettingsProvider, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetProvider(v EnumNotificationsSettingsPhoneDeliverySettingsProvider) {
	o.Provider = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSid returns the Sid field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetSid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sid
}

// GetSidOk returns a tuple with the Sid field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetSidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sid, true
}

// SetSid sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetSid(v string) {
	o.Sid = v
}

// GetAuthToken returns the AuthToken field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetAuthTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthToken, true
}

// SetAuthToken sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetAuthToken(v string) {
	o.AuthToken = v
}

// GetNumbers returns the Numbers field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetNumbers() []NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers {
	if o == nil {
		var ret []NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetNumbersOk() ([]NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers, bool) {
	if o == nil {
    return nil, false
	}
	return o.Numbers, true
}

// SetNumbers sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetNumbers(v []NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) {
	o.Numbers = v
}

func (o NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["sid"] = o.Sid
	}
	if true {
		toSerialize["authToken"] = o.AuthToken
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse struct {
	value *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) Get() *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) Set(val *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse(val *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse {
	return &NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

