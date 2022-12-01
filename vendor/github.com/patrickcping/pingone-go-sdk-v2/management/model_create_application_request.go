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

// CreateApplicationRequest - struct for CreateApplicationRequest
type CreateApplicationRequest struct {
	ApplicationExternalLink *ApplicationExternalLink
	ApplicationOIDC         *ApplicationOIDC
	ApplicationSAML         *ApplicationSAML
}

// ApplicationExternalLinkAsCreateApplicationRequest is a convenience function that returns ApplicationExternalLink wrapped in CreateApplicationRequest
func ApplicationExternalLinkAsCreateApplicationRequest(v *ApplicationExternalLink) CreateApplicationRequest {
	return CreateApplicationRequest{
		ApplicationExternalLink: v,
	}
}

// ApplicationOIDCAsCreateApplicationRequest is a convenience function that returns ApplicationOIDC wrapped in CreateApplicationRequest
func ApplicationOIDCAsCreateApplicationRequest(v *ApplicationOIDC) CreateApplicationRequest {
	return CreateApplicationRequest{
		ApplicationOIDC: v,
	}
}

// ApplicationSAMLAsCreateApplicationRequest is a convenience function that returns ApplicationSAML wrapped in CreateApplicationRequest
func ApplicationSAMLAsCreateApplicationRequest(v *ApplicationSAML) CreateApplicationRequest {
	return CreateApplicationRequest{
		ApplicationSAML: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateApplicationRequest) UnmarshalJSON(data []byte) error {

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
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateApplicationRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateApplicationRequest) MarshalJSON() ([]byte, error) {
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
func (obj *CreateApplicationRequest) GetActualInstance() interface{} {
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

type NullableCreateApplicationRequest struct {
	value *CreateApplicationRequest
	isSet bool
}

func (v NullableCreateApplicationRequest) Get() *CreateApplicationRequest {
	return v.value
}

func (v *NullableCreateApplicationRequest) Set(val *CreateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApplicationRequest(val *CreateApplicationRequest) *NullableCreateApplicationRequest {
	return &NullableCreateApplicationRequest{value: val, isSet: true}
}

func (v NullableCreateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
