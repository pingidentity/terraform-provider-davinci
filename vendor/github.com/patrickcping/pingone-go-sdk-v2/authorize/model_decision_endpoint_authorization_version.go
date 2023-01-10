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

// DecisionEndpointAuthorizationVersion struct for DecisionEndpointAuthorizationVersion
type DecisionEndpointAuthorizationVersion struct {
	// A string that specifies the ID of the Authorization Version deployed to this endpoint. Versioning allows independent development and deployment of policies. If omitted, the endpoint always uses the latest policy version available from the policy editor service.
	Id *string `json:"id,omitempty"`
	// A string that specifies the request URL for the authorization version endpoint.
	Href *string `json:"href,omitempty"`
	// A string that specifies the title for the authorization version response.
	Title *string `json:"title,omitempty"`
	// A string that specifies the content type for the authorization version response.
	Type *string `json:"type,omitempty"`
}

// NewDecisionEndpointAuthorizationVersion instantiates a new DecisionEndpointAuthorizationVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionEndpointAuthorizationVersion() *DecisionEndpointAuthorizationVersion {
	this := DecisionEndpointAuthorizationVersion{}
	return &this
}

// NewDecisionEndpointAuthorizationVersionWithDefaults instantiates a new DecisionEndpointAuthorizationVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionEndpointAuthorizationVersionWithDefaults() *DecisionEndpointAuthorizationVersion {
	this := DecisionEndpointAuthorizationVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DecisionEndpointAuthorizationVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpointAuthorizationVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DecisionEndpointAuthorizationVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DecisionEndpointAuthorizationVersion) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DecisionEndpointAuthorizationVersion) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpointAuthorizationVersion) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
    return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DecisionEndpointAuthorizationVersion) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DecisionEndpointAuthorizationVersion) SetHref(v string) {
	o.Href = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DecisionEndpointAuthorizationVersion) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpointAuthorizationVersion) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DecisionEndpointAuthorizationVersion) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DecisionEndpointAuthorizationVersion) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DecisionEndpointAuthorizationVersion) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpointAuthorizationVersion) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DecisionEndpointAuthorizationVersion) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DecisionEndpointAuthorizationVersion) SetType(v string) {
	o.Type = &v
}

func (o DecisionEndpointAuthorizationVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDecisionEndpointAuthorizationVersion struct {
	value *DecisionEndpointAuthorizationVersion
	isSet bool
}

func (v NullableDecisionEndpointAuthorizationVersion) Get() *DecisionEndpointAuthorizationVersion {
	return v.value
}

func (v *NullableDecisionEndpointAuthorizationVersion) Set(val *DecisionEndpointAuthorizationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionEndpointAuthorizationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionEndpointAuthorizationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionEndpointAuthorizationVersion(val *DecisionEndpointAuthorizationVersion) *NullableDecisionEndpointAuthorizationVersion {
	return &NullableDecisionEndpointAuthorizationVersion{value: val, isSet: true}
}

func (v NullableDecisionEndpointAuthorizationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionEndpointAuthorizationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


