/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

package management

import (
	"encoding/json"
	"fmt"
)

// IdentityProvider - struct for IdentityProvider
type IdentityProvider struct {
	IdentityProviderApple                *IdentityProviderApple
	IdentityProviderClientIDClientSecret *IdentityProviderClientIDClientSecret
	IdentityProviderFacebook             *IdentityProviderFacebook
	IdentityProviderOIDC                 *IdentityProviderOIDC
	IdentityProviderPaypal               *IdentityProviderPaypal
	IdentityProviderSAML                 *IdentityProviderSAML
}

// IdentityProviderAppleAsIdentityProvider is a convenience function that returns IdentityProviderApple wrapped in IdentityProvider
func IdentityProviderAppleAsIdentityProvider(v *IdentityProviderApple) IdentityProvider {
	return IdentityProvider{
		IdentityProviderApple: v,
	}
}

// IdentityProviderClientIDClientSecretAsIdentityProvider is a convenience function that returns IdentityProviderClientIDClientSecret wrapped in IdentityProvider
func IdentityProviderClientIDClientSecretAsIdentityProvider(v *IdentityProviderClientIDClientSecret) IdentityProvider {
	return IdentityProvider{
		IdentityProviderClientIDClientSecret: v,
	}
}

// IdentityProviderFacebookAsIdentityProvider is a convenience function that returns IdentityProviderFacebook wrapped in IdentityProvider
func IdentityProviderFacebookAsIdentityProvider(v *IdentityProviderFacebook) IdentityProvider {
	return IdentityProvider{
		IdentityProviderFacebook: v,
	}
}

// IdentityProviderOIDCAsIdentityProvider is a convenience function that returns IdentityProviderOIDC wrapped in IdentityProvider
func IdentityProviderOIDCAsIdentityProvider(v *IdentityProviderOIDC) IdentityProvider {
	return IdentityProvider{
		IdentityProviderOIDC: v,
	}
}

// IdentityProviderPaypalAsIdentityProvider is a convenience function that returns IdentityProviderPaypal wrapped in IdentityProvider
func IdentityProviderPaypalAsIdentityProvider(v *IdentityProviderPaypal) IdentityProvider {
	return IdentityProvider{
		IdentityProviderPaypal: v,
	}
}

// IdentityProviderSAMLAsIdentityProvider is a convenience function that returns IdentityProviderSAML wrapped in IdentityProvider
func IdentityProviderSAMLAsIdentityProvider(v *IdentityProviderSAML) IdentityProvider {
	return IdentityProvider{
		IdentityProviderSAML: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IdentityProvider) UnmarshalJSON(data []byte) error {

	var common IdentityProviderCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.IdentityProviderApple = nil
	dst.IdentityProviderClientIDClientSecret = nil
	dst.IdentityProviderFacebook = nil
	dst.IdentityProviderOIDC = nil
	dst.IdentityProviderPaypal = nil
	dst.IdentityProviderSAML = nil

	switch common.GetType() {
	case ENUMIDENTITYPROVIDEREXT_FACEBOOK:
		if err := json.Unmarshal(data, &dst.IdentityProviderFacebook); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_GOOGLE:
		if err := json.Unmarshal(data, &dst.IdentityProviderClientIDClientSecret); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_LINKEDIN:
		if err := json.Unmarshal(data, &dst.IdentityProviderClientIDClientSecret); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_APPLE:
		if err := json.Unmarshal(data, &dst.IdentityProviderApple); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_TWITTER:
		if err := json.Unmarshal(data, &dst.IdentityProviderClientIDClientSecret); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_AMAZON:
		if err := json.Unmarshal(data, &dst.IdentityProviderClientIDClientSecret); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_YAHOO:
		if err := json.Unmarshal(data, &dst.IdentityProviderClientIDClientSecret); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_MICROSOFT:
		if err := json.Unmarshal(data, &dst.IdentityProviderClientIDClientSecret); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_PAYPAL:
		if err := json.Unmarshal(data, &dst.IdentityProviderPaypal); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_GITHUB:
		if err := json.Unmarshal(data, &dst.IdentityProviderClientIDClientSecret); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_OPENID_CONNECT:
		if err := json.Unmarshal(data, &dst.IdentityProviderOIDC); err != nil { // simple model
			return err
		}
	case ENUMIDENTITYPROVIDEREXT_SAML:
		if err := json.Unmarshal(data, &dst.IdentityProviderSAML); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(IdentityProvider)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IdentityProvider) MarshalJSON() ([]byte, error) {
	if src.IdentityProviderApple != nil {
		return json.Marshal(&src.IdentityProviderApple)
	}

	if src.IdentityProviderClientIDClientSecret != nil {
		return json.Marshal(&src.IdentityProviderClientIDClientSecret)
	}

	if src.IdentityProviderFacebook != nil {
		return json.Marshal(&src.IdentityProviderFacebook)
	}

	if src.IdentityProviderOIDC != nil {
		return json.Marshal(&src.IdentityProviderOIDC)
	}

	if src.IdentityProviderPaypal != nil {
		return json.Marshal(&src.IdentityProviderPaypal)
	}

	if src.IdentityProviderSAML != nil {
		return json.Marshal(&src.IdentityProviderSAML)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IdentityProvider) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IdentityProviderApple != nil {
		return obj.IdentityProviderApple
	}

	if obj.IdentityProviderClientIDClientSecret != nil {
		return obj.IdentityProviderClientIDClientSecret
	}

	if obj.IdentityProviderFacebook != nil {
		return obj.IdentityProviderFacebook
	}

	if obj.IdentityProviderOIDC != nil {
		return obj.IdentityProviderOIDC
	}

	if obj.IdentityProviderPaypal != nil {
		return obj.IdentityProviderPaypal
	}

	if obj.IdentityProviderSAML != nil {
		return obj.IdentityProviderSAML
	}

	// all schemas are nil
	return nil
}

type NullableIdentityProvider struct {
	value *IdentityProvider
	isSet bool
}

func (v NullableIdentityProvider) Get() *IdentityProvider {
	return v.value
}

func (v *NullableIdentityProvider) Set(val *IdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvider(val *IdentityProvider) *NullableIdentityProvider {
	return &NullableIdentityProvider{value: val, isSet: true}
}

func (v NullableIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
