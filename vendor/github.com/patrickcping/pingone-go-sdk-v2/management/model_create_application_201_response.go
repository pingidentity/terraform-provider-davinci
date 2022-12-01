/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

package management

import (
	"encoding/json"
	"fmt"
)

// CreateApplication201Response - struct for CreateApplication201Response
type CreateApplication201Response struct {
	ApplicationExternalLink *ApplicationExternalLink
	ApplicationOIDC         *ApplicationOIDC
	ApplicationSAML         *ApplicationSAML
}

// ApplicationExternalLinkAsCreateApplication201Response is a convenience function that returns ApplicationExternalLink wrapped in CreateApplication201Response
func ApplicationExternalLinkAsCreateApplication201Response(v *ApplicationExternalLink) CreateApplication201Response {
	return CreateApplication201Response{
		ApplicationExternalLink: v,
	}
}

// ApplicationOIDCAsCreateApplication201Response is a convenience function that returns ApplicationOIDC wrapped in CreateApplication201Response
func ApplicationOIDCAsCreateApplication201Response(v *ApplicationOIDC) CreateApplication201Response {
	return CreateApplication201Response{
		ApplicationOIDC: v,
	}
}

// ApplicationSAMLAsCreateApplication201Response is a convenience function that returns ApplicationSAML wrapped in CreateApplication201Response
func ApplicationSAMLAsCreateApplication201Response(v *ApplicationSAML) CreateApplication201Response {
	return CreateApplication201Response{
		ApplicationSAML: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateApplication201Response) UnmarshalJSON(data []byte) error {

	var common Application

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.ApplicationOIDC = nil
	dst.ApplicationSAML = nil
	dst.ApplicationExternalLink = nil

	switch common.GetProtocol() {
	case ENUMAPPLICATIONPROTOCOL_OPENID_CONNECT:
		if err := json.Unmarshal(data, &dst.ApplicationOIDC); err != nil { // simple model
			return err
		}
	case ENUMAPPLICATIONPROTOCOL_SAML:
		if err := json.Unmarshal(data, &dst.ApplicationSAML); err != nil { // simple model
			return err
		}
	case ENUMAPPLICATIONPROTOCOL_EXTERNAL_LINK:
		if err := json.Unmarshal(data, &dst.ApplicationExternalLink); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateApplication201Response)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateApplication201Response) MarshalJSON() ([]byte, error) {
	if src.ApplicationExternalLink != nil {
		return json.Marshal(&src.ApplicationExternalLink)
	}

	if src.ApplicationOIDC != nil {
		return json.Marshal(&src.ApplicationOIDC)
	}

	if src.ApplicationSAML != nil {
		return json.Marshal(&src.ApplicationSAML)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateApplication201Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApplicationExternalLink != nil {
		return obj.ApplicationExternalLink
	}

	if obj.ApplicationOIDC != nil {
		return obj.ApplicationOIDC
	}

	if obj.ApplicationSAML != nil {
		return obj.ApplicationSAML
	}

	// all schemas are nil
	return nil
}

type NullableCreateApplication201Response struct {
	value *CreateApplication201Response
	isSet bool
}

func (v NullableCreateApplication201Response) Get() *CreateApplication201Response {
	return v.value
}

func (v *NullableCreateApplication201Response) Set(val *CreateApplication201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApplication201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApplication201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApplication201Response(val *CreateApplication201Response) *NullableCreateApplication201Response {
	return &NullableCreateApplication201Response{value: val, isSet: true}
}

func (v NullableCreateApplication201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApplication201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
