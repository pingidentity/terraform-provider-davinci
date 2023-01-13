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

// GroupMembership struct for GroupMembership
type GroupMembership struct {
	// ID of the group to assign
	Id string `json:"id"`
}

// NewGroupMembership instantiates a new GroupMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMembership(id string) *GroupMembership {
	this := GroupMembership{}
	this.Id = id
	return &this
}

// NewGroupMembershipWithDefaults instantiates a new GroupMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMembershipWithDefaults() *GroupMembership {
	this := GroupMembership{}
	return &this
}

// GetId returns the Id field value
func (o *GroupMembership) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupMembership) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupMembership) SetId(v string) {
	o.Id = v
}

func (o GroupMembership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGroupMembership struct {
	value *GroupMembership
	isSet bool
}

func (v NullableGroupMembership) Get() *GroupMembership {
	return v.value
}

func (v *NullableGroupMembership) Set(val *GroupMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMembership(val *GroupMembership) *NullableGroupMembership {
	return &NullableGroupMembership{value: val, isSet: true}
}

func (v NullableGroupMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


