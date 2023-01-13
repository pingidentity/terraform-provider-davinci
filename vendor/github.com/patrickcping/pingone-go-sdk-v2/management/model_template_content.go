/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// TemplateContent - struct for TemplateContent
type TemplateContent struct {
	TemplateContentEmail *TemplateContentEmail
	TemplateContentPush *TemplateContentPush
	TemplateContentSMS *TemplateContentSMS
	TemplateContentVoice *TemplateContentVoice
}

// TemplateContentEmailAsTemplateContent is a convenience function that returns TemplateContentEmail wrapped in TemplateContent
func TemplateContentEmailAsTemplateContent(v *TemplateContentEmail) TemplateContent {
	return TemplateContent{
		TemplateContentEmail: v,
	}
}

// TemplateContentPushAsTemplateContent is a convenience function that returns TemplateContentPush wrapped in TemplateContent
func TemplateContentPushAsTemplateContent(v *TemplateContentPush) TemplateContent {
	return TemplateContent{
		TemplateContentPush: v,
	}
}

// TemplateContentSMSAsTemplateContent is a convenience function that returns TemplateContentSMS wrapped in TemplateContent
func TemplateContentSMSAsTemplateContent(v *TemplateContentSMS) TemplateContent {
	return TemplateContent{
		TemplateContentSMS: v,
	}
}

// TemplateContentVoiceAsTemplateContent is a convenience function that returns TemplateContentVoice wrapped in TemplateContent
func TemplateContentVoiceAsTemplateContent(v *TemplateContentVoice) TemplateContent {
	return TemplateContent{
		TemplateContentVoice: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TemplateContent) UnmarshalJSON(data []byte) error {

	var common TemplateContentCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.TemplateContentEmail = nil
	dst.TemplateContentPush = nil
	dst.TemplateContentSMS = nil
	dst.TemplateContentVoice = nil

	switch common.GetDeliveryMethod() {
	case ENUMTEMPLATECONTENTDELIVERYMETHOD_EMAIL:
		if err := json.Unmarshal(data, &dst.TemplateContentEmail); err != nil { // simple model
			return err
		}
	case ENUMTEMPLATECONTENTDELIVERYMETHOD_PUSH:
		if err := json.Unmarshal(data, &dst.TemplateContentPush); err != nil { // simple model
			return err
		}
	case ENUMTEMPLATECONTENTDELIVERYMETHOD_SMS:
		if err := json.Unmarshal(data, &dst.TemplateContentSMS); err != nil { // simple model
			return err
		}
	case ENUMTEMPLATECONTENTDELIVERYMETHOD_VOICE:
		if err := json.Unmarshal(data, &dst.TemplateContentVoice); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(TemplateContent)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TemplateContent) MarshalJSON() ([]byte, error) {
	if src.TemplateContentEmail != nil {
		return json.Marshal(&src.TemplateContentEmail)
	}

	if src.TemplateContentPush != nil {
		return json.Marshal(&src.TemplateContentPush)
	}

	if src.TemplateContentSMS != nil {
		return json.Marshal(&src.TemplateContentSMS)
	}

	if src.TemplateContentVoice != nil {
		return json.Marshal(&src.TemplateContentVoice)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TemplateContent) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TemplateContentEmail != nil {
		return obj.TemplateContentEmail
	}

	if obj.TemplateContentPush != nil {
		return obj.TemplateContentPush
	}

	if obj.TemplateContentSMS != nil {
		return obj.TemplateContentSMS
	}

	if obj.TemplateContentVoice != nil {
		return obj.TemplateContentVoice
	}

	// all schemas are nil
	return nil
}

type NullableTemplateContent struct {
	value *TemplateContent
	isSet bool
}

func (v NullableTemplateContent) Get() *TemplateContent {
	return v.value
}

func (v *NullableTemplateContent) Set(val *TemplateContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContent(val *TemplateContent) *NullableTemplateContent {
	return &NullableTemplateContent{value: val, isSet: true}
}

func (v NullableTemplateContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

