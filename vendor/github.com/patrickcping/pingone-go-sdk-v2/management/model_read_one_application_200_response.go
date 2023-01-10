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

// ReadOneApplication200Response - struct for ReadOneApplication200Response
type ReadOneApplication200Response struct {
	ApplicationExternalLink       *ApplicationExternalLink
	ApplicationOIDC               *ApplicationOIDC
	ApplicationPingOnePortal      *ApplicationPingOnePortal
	ApplicationPingOneSelfService *ApplicationPingOneSelfService
	ApplicationSAML               *ApplicationSAML
	ApplicationWSFED              *ApplicationWSFED
}

// ApplicationExternalLinkAsReadOneApplication200Response is a convenience function that returns ApplicationExternalLink wrapped in ReadOneApplication200Response
func ApplicationExternalLinkAsReadOneApplication200Response(v *ApplicationExternalLink) ReadOneApplication200Response {
	return ReadOneApplication200Response{
		ApplicationExternalLink: v,
	}
}

// ApplicationOIDCAsReadOneApplication200Response is a convenience function that returns ApplicationOIDC wrapped in ReadOneApplication200Response
func ApplicationOIDCAsReadOneApplication200Response(v *ApplicationOIDC) ReadOneApplication200Response {
	return ReadOneApplication200Response{
		ApplicationOIDC: v,
	}
}

// ApplicationPingOnePortalAsReadOneApplication200Response is a convenience function that returns ApplicationPingOnePortal wrapped in ReadOneApplication200Response
func ApplicationPingOnePortalAsReadOneApplication200Response(v *ApplicationPingOnePortal) ReadOneApplication200Response {
	return ReadOneApplication200Response{
		ApplicationPingOnePortal: v,
	}
}

// ApplicationPingOneSelfServiceAsReadOneApplication200Response is a convenience function that returns ApplicationPingOneSelfService wrapped in ReadOneApplication200Response
func ApplicationPingOneSelfServiceAsReadOneApplication200Response(v *ApplicationPingOneSelfService) ReadOneApplication200Response {
	return ReadOneApplication200Response{
		ApplicationPingOneSelfService: v,
	}
}

// ApplicationSAMLAsReadOneApplication200Response is a convenience function that returns ApplicationSAML wrapped in ReadOneApplication200Response
func ApplicationSAMLAsReadOneApplication200Response(v *ApplicationSAML) ReadOneApplication200Response {
	return ReadOneApplication200Response{
		ApplicationSAML: v,
	}
}

// ApplicationWSFEDAsReadOneApplication200Response is a convenience function that returns ApplicationWSFED wrapped in ReadOneApplication200Response
func ApplicationWSFEDAsReadOneApplication200Response(v *ApplicationWSFED) ReadOneApplication200Response {
	return ReadOneApplication200Response{
		ApplicationWSFED: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReadOneApplication200Response) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("Data failed to match schemas in oneOf(EntityArrayEmbeddedApplicationsInner)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReadOneApplication200Response) MarshalJSON() ([]byte, error) {
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
func (obj *ReadOneApplication200Response) GetActualInstance() interface{} {
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

type NullableReadOneApplication200Response struct {
	value *ReadOneApplication200Response
	isSet bool
}

func (v NullableReadOneApplication200Response) Get() *ReadOneApplication200Response {
	return v.value
}

func (v *NullableReadOneApplication200Response) Set(val *ReadOneApplication200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOneApplication200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOneApplication200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOneApplication200Response(val *ReadOneApplication200Response) *NullableReadOneApplication200Response {
	return &NullableReadOneApplication200Response{value: val, isSet: true}
}

func (v NullableReadOneApplication200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOneApplication200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
