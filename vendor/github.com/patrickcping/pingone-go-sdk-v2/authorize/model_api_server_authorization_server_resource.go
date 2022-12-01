/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2022-09-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// APIServerAuthorizationServerResource The resource defines the characteristics of the OAuth 2.0 access tokens used to get access to the APIs on the API server such as the audience and scopes.
type APIServerAuthorizationServerResource struct {
	// A string that specifies the UUID of the custom PingOne resource. This property must identify a PingOne resource with a type property value of CUSTOM.
	Id string `json:"id"`
}

// NewAPIServerAuthorizationServerResource instantiates a new APIServerAuthorizationServerResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerAuthorizationServerResource(id string) *APIServerAuthorizationServerResource {
	this := APIServerAuthorizationServerResource{}
	this.Id = id
	return &this
}

// NewAPIServerAuthorizationServerResourceWithDefaults instantiates a new APIServerAuthorizationServerResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerAuthorizationServerResourceWithDefaults() *APIServerAuthorizationServerResource {
	this := APIServerAuthorizationServerResource{}
	return &this
}

// GetId returns the Id field value
func (o *APIServerAuthorizationServerResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *APIServerAuthorizationServerResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *APIServerAuthorizationServerResource) SetId(v string) {
	o.Id = v
}

func (o APIServerAuthorizationServerResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAPIServerAuthorizationServerResource struct {
	value *APIServerAuthorizationServerResource
	isSet bool
}

func (v NullableAPIServerAuthorizationServerResource) Get() *APIServerAuthorizationServerResource {
	return v.value
}

func (v *NullableAPIServerAuthorizationServerResource) Set(val *APIServerAuthorizationServerResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerAuthorizationServerResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerAuthorizationServerResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerAuthorizationServerResource(val *APIServerAuthorizationServerResource) *NullableAPIServerAuthorizationServerResource {
	return &NullableAPIServerAuthorizationServerResource{value: val, isSet: true}
}

func (v NullableAPIServerAuthorizationServerResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerAuthorizationServerResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

