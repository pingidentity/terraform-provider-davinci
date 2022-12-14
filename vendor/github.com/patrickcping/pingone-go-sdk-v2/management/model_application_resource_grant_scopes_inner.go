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

// ApplicationResourceGrantScopesInner struct for ApplicationResourceGrantScopesInner
type ApplicationResourceGrantScopesInner struct {
	// id A array that specifies the IDs of the scopes associated with this grant. This is a required property.
	Id string `json:"id"`
}

// NewApplicationResourceGrantScopesInner instantiates a new ApplicationResourceGrantScopesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResourceGrantScopesInner(id string) *ApplicationResourceGrantScopesInner {
	this := ApplicationResourceGrantScopesInner{}
	this.Id = id
	return &this
}

// NewApplicationResourceGrantScopesInnerWithDefaults instantiates a new ApplicationResourceGrantScopesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourceGrantScopesInnerWithDefaults() *ApplicationResourceGrantScopesInner {
	this := ApplicationResourceGrantScopesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationResourceGrantScopesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrantScopesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationResourceGrantScopesInner) SetId(v string) {
	o.Id = v
}

func (o ApplicationResourceGrantScopesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResourceGrantScopesInner struct {
	value *ApplicationResourceGrantScopesInner
	isSet bool
}

func (v NullableApplicationResourceGrantScopesInner) Get() *ApplicationResourceGrantScopesInner {
	return v.value
}

func (v *NullableApplicationResourceGrantScopesInner) Set(val *ApplicationResourceGrantScopesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResourceGrantScopesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResourceGrantScopesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResourceGrantScopesInner(val *ApplicationResourceGrantScopesInner) *NullableApplicationResourceGrantScopesInner {
	return &NullableApplicationResourceGrantScopesInner{value: val, isSet: true}
}

func (v NullableApplicationResourceGrantScopesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResourceGrantScopesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


