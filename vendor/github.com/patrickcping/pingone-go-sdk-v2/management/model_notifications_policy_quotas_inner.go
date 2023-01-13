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

// NotificationsPolicyQuotasInner struct for NotificationsPolicyQuotasInner
type NotificationsPolicyQuotasInner struct {
	Type EnumNotificationsPolicyQuotaItemType `json:"type"`
	// The delivery methods for which the limit is being defined. Currently, the content of the array must be `SMS`, `Voice`. This means that the combined total of SMS and Voice notifications must be below the limit defined.
	DeliveryMethods []EnumNotificationsPolicyQuotaDeliveryMethods `json:"deliveryMethods"`
	// The maximum number of notifications allowed per day.
	Total *int32 `json:"total,omitempty"`
	// The maximum number of notifications that can be received and responded to each day. Used in conjunction with unclaimed in place of the single field total.
	Claimed *int32 `json:"claimed,omitempty"`
	// The maximum number of notifications that can be received and not responded to each day. Used in conjunction with claimed in place of the single field total.
	Unclaimed *int32 `json:"unclaimed,omitempty"`
}

// NewNotificationsPolicyQuotasInner instantiates a new NotificationsPolicyQuotasInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsPolicyQuotasInner(type_ EnumNotificationsPolicyQuotaItemType, deliveryMethods []EnumNotificationsPolicyQuotaDeliveryMethods) *NotificationsPolicyQuotasInner {
	this := NotificationsPolicyQuotasInner{}
	this.Type = type_
	this.DeliveryMethods = deliveryMethods
	return &this
}

// NewNotificationsPolicyQuotasInnerWithDefaults instantiates a new NotificationsPolicyQuotasInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsPolicyQuotasInnerWithDefaults() *NotificationsPolicyQuotasInner {
	this := NotificationsPolicyQuotasInner{}
	return &this
}

// GetType returns the Type field value
func (o *NotificationsPolicyQuotasInner) GetType() EnumNotificationsPolicyQuotaItemType {
	if o == nil {
		var ret EnumNotificationsPolicyQuotaItemType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyQuotasInner) GetTypeOk() (*EnumNotificationsPolicyQuotaItemType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationsPolicyQuotasInner) SetType(v EnumNotificationsPolicyQuotaItemType) {
	o.Type = v
}

// GetDeliveryMethods returns the DeliveryMethods field value
func (o *NotificationsPolicyQuotasInner) GetDeliveryMethods() []EnumNotificationsPolicyQuotaDeliveryMethods {
	if o == nil {
		var ret []EnumNotificationsPolicyQuotaDeliveryMethods
		return ret
	}

	return o.DeliveryMethods
}

// GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field value
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyQuotasInner) GetDeliveryMethodsOk() ([]EnumNotificationsPolicyQuotaDeliveryMethods, bool) {
	if o == nil {
    return nil, false
	}
	return o.DeliveryMethods, true
}

// SetDeliveryMethods sets field value
func (o *NotificationsPolicyQuotasInner) SetDeliveryMethods(v []EnumNotificationsPolicyQuotaDeliveryMethods) {
	o.DeliveryMethods = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NotificationsPolicyQuotasInner) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyQuotasInner) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NotificationsPolicyQuotasInner) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *NotificationsPolicyQuotasInner) SetTotal(v int32) {
	o.Total = &v
}

// GetClaimed returns the Claimed field value if set, zero value otherwise.
func (o *NotificationsPolicyQuotasInner) GetClaimed() int32 {
	if o == nil || isNil(o.Claimed) {
		var ret int32
		return ret
	}
	return *o.Claimed
}

// GetClaimedOk returns a tuple with the Claimed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyQuotasInner) GetClaimedOk() (*int32, bool) {
	if o == nil || isNil(o.Claimed) {
    return nil, false
	}
	return o.Claimed, true
}

// HasClaimed returns a boolean if a field has been set.
func (o *NotificationsPolicyQuotasInner) HasClaimed() bool {
	if o != nil && !isNil(o.Claimed) {
		return true
	}

	return false
}

// SetClaimed gets a reference to the given int32 and assigns it to the Claimed field.
func (o *NotificationsPolicyQuotasInner) SetClaimed(v int32) {
	o.Claimed = &v
}

// GetUnclaimed returns the Unclaimed field value if set, zero value otherwise.
func (o *NotificationsPolicyQuotasInner) GetUnclaimed() int32 {
	if o == nil || isNil(o.Unclaimed) {
		var ret int32
		return ret
	}
	return *o.Unclaimed
}

// GetUnclaimedOk returns a tuple with the Unclaimed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyQuotasInner) GetUnclaimedOk() (*int32, bool) {
	if o == nil || isNil(o.Unclaimed) {
    return nil, false
	}
	return o.Unclaimed, true
}

// HasUnclaimed returns a boolean if a field has been set.
func (o *NotificationsPolicyQuotasInner) HasUnclaimed() bool {
	if o != nil && !isNil(o.Unclaimed) {
		return true
	}

	return false
}

// SetUnclaimed gets a reference to the given int32 and assigns it to the Unclaimed field.
func (o *NotificationsPolicyQuotasInner) SetUnclaimed(v int32) {
	o.Unclaimed = &v
}

func (o NotificationsPolicyQuotasInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["deliveryMethods"] = o.DeliveryMethods
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Claimed) {
		toSerialize["claimed"] = o.Claimed
	}
	if !isNil(o.Unclaimed) {
		toSerialize["unclaimed"] = o.Unclaimed
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationsPolicyQuotasInner struct {
	value *NotificationsPolicyQuotasInner
	isSet bool
}

func (v NullableNotificationsPolicyQuotasInner) Get() *NotificationsPolicyQuotasInner {
	return v.value
}

func (v *NullableNotificationsPolicyQuotasInner) Set(val *NotificationsPolicyQuotasInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsPolicyQuotasInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsPolicyQuotasInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsPolicyQuotasInner(val *NotificationsPolicyQuotasInner) *NullableNotificationsPolicyQuotasInner {
	return &NullableNotificationsPolicyQuotasInner{value: val, isSet: true}
}

func (v NullableNotificationsPolicyQuotasInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsPolicyQuotasInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

