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

// NotificationsSettingsPhoneDeliverySettingsCustomAllOf struct for NotificationsSettingsPhoneDeliverySettingsCustomAllOf
type NotificationsSettingsPhoneDeliverySettingsCustomAllOf struct {
	// The customer provider's name.
	Name *string `json:"name,omitempty"`
	Requests NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests `json:"requests"`
	Authentication NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication `json:"authentication"`
}

// NewNotificationsSettingsPhoneDeliverySettingsCustomAllOf instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOf(requests NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests, authentication NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) *NotificationsSettingsPhoneDeliverySettingsCustomAllOf {
	this := NotificationsSettingsPhoneDeliverySettingsCustomAllOf{}
	this.Requests = requests
	this.Authentication = authentication
	return &this
}

// NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomAllOf {
	this := NotificationsSettingsPhoneDeliverySettingsCustomAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) SetName(v string) {
	o.Name = &v
}

// GetRequests returns the Requests field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetRequests() NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests {
	if o == nil {
		var ret NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetRequestsOk() (*NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) SetRequests(v NotificationsSettingsPhoneDeliverySettingsCustomAllOfRequests) {
	o.Requests = v
}

// GetAuthentication returns the Authentication field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetAuthentication() NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication {
	if o == nil {
		var ret NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication
		return ret
	}

	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetAuthenticationOk() (*NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) SetAuthentication(v NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) {
	o.Authentication = v
}

func (o NotificationsSettingsPhoneDeliverySettingsCustomAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["requests"] = o.Requests
	}
	if true {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf struct {
	value *NotificationsSettingsPhoneDeliverySettingsCustomAllOf
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf) Get() *NotificationsSettingsPhoneDeliverySettingsCustomAllOf {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf) Set(val *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf(val *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf {
	return &NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


