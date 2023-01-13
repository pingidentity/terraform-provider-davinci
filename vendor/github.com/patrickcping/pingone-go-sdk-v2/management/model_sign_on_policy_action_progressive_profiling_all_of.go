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

// SignOnPolicyActionProgressiveProfilingAllOf struct for SignOnPolicyActionProgressiveProfilingAllOf
type SignOnPolicyActionProgressiveProfilingAllOf struct {
	Attributes []SignOnPolicyActionProgressiveProfilingAllOfAttributes `json:"attributes"`
	// A boolean that specifies whether the progressive profiling action will not be executed if another progressive profiling action has already been executed during the flow. This property is required.
	PreventMultiplePromptsPerFlow bool `json:"preventMultiplePromptsPerFlow"`
	// An integer that specifies how often to prompt the user to provide profile data for the configured attributes for which they do not have values. This property is required.
	PromptIntervalSeconds int32 `json:"promptIntervalSeconds"`
	// A string that specifies text to display to the user when prompting for attribute values. This property is required.
	PromptText string `json:"promptText"`
}

// NewSignOnPolicyActionProgressiveProfilingAllOf instantiates a new SignOnPolicyActionProgressiveProfilingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionProgressiveProfilingAllOf(attributes []SignOnPolicyActionProgressiveProfilingAllOfAttributes, preventMultiplePromptsPerFlow bool, promptIntervalSeconds int32, promptText string) *SignOnPolicyActionProgressiveProfilingAllOf {
	this := SignOnPolicyActionProgressiveProfilingAllOf{}
	this.Attributes = attributes
	this.PreventMultiplePromptsPerFlow = preventMultiplePromptsPerFlow
	this.PromptIntervalSeconds = promptIntervalSeconds
	this.PromptText = promptText
	return &this
}

// NewSignOnPolicyActionProgressiveProfilingAllOfWithDefaults instantiates a new SignOnPolicyActionProgressiveProfilingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionProgressiveProfilingAllOfWithDefaults() *SignOnPolicyActionProgressiveProfilingAllOf {
	this := SignOnPolicyActionProgressiveProfilingAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetAttributes() []SignOnPolicyActionProgressiveProfilingAllOfAttributes {
	if o == nil {
		var ret []SignOnPolicyActionProgressiveProfilingAllOfAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetAttributesOk() ([]SignOnPolicyActionProgressiveProfilingAllOfAttributes, bool) {
	if o == nil {
    return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetAttributes(v []SignOnPolicyActionProgressiveProfilingAllOfAttributes) {
	o.Attributes = v
}

// GetPreventMultiplePromptsPerFlow returns the PreventMultiplePromptsPerFlow field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPreventMultiplePromptsPerFlow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PreventMultiplePromptsPerFlow
}

// GetPreventMultiplePromptsPerFlowOk returns a tuple with the PreventMultiplePromptsPerFlow field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPreventMultiplePromptsPerFlowOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreventMultiplePromptsPerFlow, true
}

// SetPreventMultiplePromptsPerFlow sets field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetPreventMultiplePromptsPerFlow(v bool) {
	o.PreventMultiplePromptsPerFlow = v
}

// GetPromptIntervalSeconds returns the PromptIntervalSeconds field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptIntervalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PromptIntervalSeconds
}

// GetPromptIntervalSecondsOk returns a tuple with the PromptIntervalSeconds field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptIntervalSecondsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PromptIntervalSeconds, true
}

// SetPromptIntervalSeconds sets field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetPromptIntervalSeconds(v int32) {
	o.PromptIntervalSeconds = v
}

// GetPromptText returns the PromptText field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PromptText
}

// GetPromptTextOk returns a tuple with the PromptText field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptTextOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PromptText, true
}

// SetPromptText sets field value
func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetPromptText(v string) {
	o.PromptText = v
}

func (o SignOnPolicyActionProgressiveProfilingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["preventMultiplePromptsPerFlow"] = o.PreventMultiplePromptsPerFlow
	}
	if true {
		toSerialize["promptIntervalSeconds"] = o.PromptIntervalSeconds
	}
	if true {
		toSerialize["promptText"] = o.PromptText
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionProgressiveProfilingAllOf struct {
	value *SignOnPolicyActionProgressiveProfilingAllOf
	isSet bool
}

func (v NullableSignOnPolicyActionProgressiveProfilingAllOf) Get() *SignOnPolicyActionProgressiveProfilingAllOf {
	return v.value
}

func (v *NullableSignOnPolicyActionProgressiveProfilingAllOf) Set(val *SignOnPolicyActionProgressiveProfilingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionProgressiveProfilingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionProgressiveProfilingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionProgressiveProfilingAllOf(val *SignOnPolicyActionProgressiveProfilingAllOf) *NullableSignOnPolicyActionProgressiveProfilingAllOf {
	return &NullableSignOnPolicyActionProgressiveProfilingAllOf{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionProgressiveProfilingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionProgressiveProfilingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


