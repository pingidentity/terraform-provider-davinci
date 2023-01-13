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

// UpdateApplicationRequest - struct for UpdateApplicationRequest
type UpdateApplicationRequest struct {
	ApplicationExternalLink       *ApplicationExternalLink
	ApplicationOIDC               *ApplicationOIDC
	ApplicationPingOnePortal      *ApplicationPingOnePortal
	ApplicationPingOneSelfService *ApplicationPingOneSelfService
	ApplicationSAML               *ApplicationSAML
	ApplicationWSFED              *ApplicationWSFED
}

// ApplicationExternalLinkAsUpdateApplicationRequest is a convenience function that returns ApplicationExternalLink wrapped in UpdateApplicationRequest
func ApplicationExternalLinkAsUpdateApplicationRequest(v *ApplicationExternalLink) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationExternalLink: v,
	}
}

// ApplicationOIDCAsUpdateApplicationRequest is a convenience function that returns ApplicationOIDC wrapped in UpdateApplicationRequest
func ApplicationOIDCAsUpdateApplicationRequest(v *ApplicationOIDC) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationOIDC: v,
	}
}

// ApplicationPingOnePortalAsUpdateApplicationRequest is a convenience function that returns ApplicationPingOnePortal wrapped in UpdateApplicationRequest
func ApplicationPingOnePortalAsUpdateApplicationRequest(v *ApplicationPingOnePortal) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationPingOnePortal: v,
	}
}

// ApplicationPingOneSelfServiceAsUpdateApplicationRequest is a convenience function that returns ApplicationPingOneSelfService wrapped in UpdateApplicationRequest
func ApplicationPingOneSelfServiceAsUpdateApplicationRequest(v *ApplicationPingOneSelfService) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationPingOneSelfService: v,
	}
}

// ApplicationSAMLAsUpdateApplicationRequest is a convenience function that returns ApplicationSAML wrapped in UpdateApplicationRequest
func ApplicationSAMLAsUpdateApplicationRequest(v *ApplicationSAML) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationSAML: v,
	}
}

// ApplicationWSFEDAsUpdateApplicationRequest is a convenience function that returns ApplicationWSFED wrapped in UpdateApplicationRequest
func ApplicationWSFEDAsUpdateApplicationRequest(v *ApplicationWSFED) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationWSFED: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateApplicationRequest) UnmarshalJSON(data []byte) error {

	var common Application

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.ApplicationOIDC = nil
	dst.ApplicationSAML = nil
	dst.ApplicationWSFED = nil
	dst.ApplicationExternalLink = nil
	dst.ApplicationPingOnePortal = nil
	dst.ApplicationPingOneSelfService = nil

	switch common.GetProtocol() {
	case ENUMAPPLICATIONPROTOCOL_OPENID_CONNECT:
		switch common.GetType() {
		case ENUMAPPLICATIONTYPE_PING_ONE_PORTAL:
			if err := json.Unmarshal(data, &dst.ApplicationPingOnePortal); err != nil { // simple model
				return err
			}
		case ENUMAPPLICATIONTYPE_PING_ONE_SELF_SERVICE:
			if err := json.Unmarshal(data, &dst.ApplicationPingOneSelfService); err != nil { // simple model
				return err
			}
		case ENUMAPPLICATIONTYPE_PING_ONE_ADMIN_CONSOLE:
			return fmt.Errorf("PingOne admin console not yet supported in oneOf(EntityArrayEmbeddedApplicationsInner)")
		default:
			if err := json.Unmarshal(data, &dst.ApplicationOIDC); err != nil { // simple model
				return err
			}
		}
	case ENUMAPPLICATIONPROTOCOL_SAML:
		if err := json.Unmarshal(data, &dst.ApplicationSAML); err != nil { // simple model
			return err
		}
	case ENUMAPPLICATIONPROTOCOL_WS_FED:
		if err := json.Unmarshal(data, &dst.ApplicationWSFED); err != nil { // simple model
			return err
		}
	case ENUMAPPLICATIONPROTOCOL_EXTERNAL_LINK:
		if err := json.Unmarshal(data, &dst.ApplicationExternalLink); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateApplicationRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	if src.ApplicationExternalLink != nil {
		return json.Marshal(&src.ApplicationExternalLink)
	}

	if src.ApplicationOIDC != nil {
		return json.Marshal(&src.ApplicationOIDC)
	}

	if src.ApplicationPingOnePortal != nil {
		return json.Marshal(&src.ApplicationPingOnePortal)
	}

	if src.ApplicationPingOneSelfService != nil {
		return json.Marshal(&src.ApplicationPingOneSelfService)
	}

	if src.ApplicationSAML != nil {
		return json.Marshal(&src.ApplicationSAML)
	}

	if src.ApplicationWSFED != nil {
		return json.Marshal(&src.ApplicationWSFED)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateApplicationRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApplicationExternalLink != nil {
		return obj.ApplicationExternalLink
	}

	if obj.ApplicationOIDC != nil {
		return obj.ApplicationOIDC
	}

	if obj.ApplicationPingOnePortal != nil {
		return obj.ApplicationPingOnePortal
	}

	if obj.ApplicationPingOneSelfService != nil {
		return obj.ApplicationPingOneSelfService
	}

	if obj.ApplicationSAML != nil {
		return obj.ApplicationSAML
	}

	if obj.ApplicationWSFED != nil {
		return obj.ApplicationWSFED
	}

	// all schemas are nil
	return nil
}

type NullableUpdateApplicationRequest struct {
	value *UpdateApplicationRequest
	isSet bool
}

func (v NullableUpdateApplicationRequest) Get() *UpdateApplicationRequest {
	return v.value
}

func (v *NullableUpdateApplicationRequest) Set(val *UpdateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplicationRequest(val *UpdateApplicationRequest) *NullableUpdateApplicationRequest {
	return &NullableUpdateApplicationRequest{value: val, isSet: true}
}

func (v NullableUpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
