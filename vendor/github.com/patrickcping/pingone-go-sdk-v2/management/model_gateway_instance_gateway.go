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

// GatewayInstanceGateway struct for GatewayInstanceGateway
type GatewayInstanceGateway struct {
	// A string that specifies the gateway ID.
	Id *string `json:"id,omitempty"`
}

// NewGatewayInstanceGateway instantiates a new GatewayInstanceGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayInstanceGateway() *GatewayInstanceGateway {
	this := GatewayInstanceGateway{}
	return &this
}

// NewGatewayInstanceGatewayWithDefaults instantiates a new GatewayInstanceGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayInstanceGatewayWithDefaults() *GatewayInstanceGateway {
	this := GatewayInstanceGateway{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayInstanceGateway) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayInstanceGateway) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayInstanceGateway) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayInstanceGateway) SetId(v string) {
	o.Id = &v
}

func (o GatewayInstanceGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayInstanceGateway struct {
	value *GatewayInstanceGateway
	isSet bool
}

func (v NullableGatewayInstanceGateway) Get() *GatewayInstanceGateway {
	return v.value
}

func (v *NullableGatewayInstanceGateway) Set(val *GatewayInstanceGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayInstanceGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayInstanceGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayInstanceGateway(val *GatewayInstanceGateway) *NullableGatewayInstanceGateway {
	return &NullableGatewayInstanceGateway{value: val, isSet: true}
}

func (v NullableGatewayInstanceGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayInstanceGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


