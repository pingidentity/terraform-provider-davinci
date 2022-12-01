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

// SignOnPolicyAction - struct for SignOnPolicyAction
type SignOnPolicyAction struct {
	SignOnPolicyActionAgreement                  *SignOnPolicyActionAgreement
	SignOnPolicyActionCommon                     *SignOnPolicyActionCommon
	SignOnPolicyActionIDFirst                    *SignOnPolicyActionIDFirst
	SignOnPolicyActionIDP                        *SignOnPolicyActionIDP
	SignOnPolicyActionLogin                      *SignOnPolicyActionLogin
	SignOnPolicyActionMFA                        *SignOnPolicyActionMFA
	SignOnPolicyActionPingIDWinLoginPasswordless *SignOnPolicyActionPingIDWinLoginPasswordless
	SignOnPolicyActionProgressiveProfiling       *SignOnPolicyActionProgressiveProfiling
}

// SignOnPolicyActionAgreementAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionAgreement wrapped in SignOnPolicyAction
func SignOnPolicyActionAgreementAsSignOnPolicyAction(v *SignOnPolicyActionAgreement) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionAgreement: v,
	}
}

// SignOnPolicyActionCommonAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionCommon wrapped in SignOnPolicyAction
func SignOnPolicyActionCommonAsSignOnPolicyAction(v *SignOnPolicyActionCommon) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionCommon: v,
	}
}

// SignOnPolicyActionIDFirstAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionIDFirst wrapped in SignOnPolicyAction
func SignOnPolicyActionIDFirstAsSignOnPolicyAction(v *SignOnPolicyActionIDFirst) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionIDFirst: v,
	}
}

// SignOnPolicyActionIDPAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionIDP wrapped in SignOnPolicyAction
func SignOnPolicyActionIDPAsSignOnPolicyAction(v *SignOnPolicyActionIDP) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionIDP: v,
	}
}

// SignOnPolicyActionLoginAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionLogin wrapped in SignOnPolicyAction
func SignOnPolicyActionLoginAsSignOnPolicyAction(v *SignOnPolicyActionLogin) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionLogin: v,
	}
}

// SignOnPolicyActionMFAAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionMFA wrapped in SignOnPolicyAction
func SignOnPolicyActionMFAAsSignOnPolicyAction(v *SignOnPolicyActionMFA) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionMFA: v,
	}
}

// SignOnPolicyActionPingIDWinLoginPasswordlessAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionPingIDWinLoginPasswordless wrapped in SignOnPolicyAction
func SignOnPolicyActionPingIDWinLoginPasswordlessAsSignOnPolicyAction(v *SignOnPolicyActionPingIDWinLoginPasswordless) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionPingIDWinLoginPasswordless: v,
	}
}

// SignOnPolicyActionProgressiveProfilingAsSignOnPolicyAction is a convenience function that returns SignOnPolicyActionProgressiveProfiling wrapped in SignOnPolicyAction
func SignOnPolicyActionProgressiveProfilingAsSignOnPolicyAction(v *SignOnPolicyActionProgressiveProfiling) SignOnPolicyAction {
	return SignOnPolicyAction{
		SignOnPolicyActionProgressiveProfiling: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SignOnPolicyAction) UnmarshalJSON(data []byte) error {
	var common SignOnPolicyActionCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.SignOnPolicyActionAgreement = nil
	dst.SignOnPolicyActionIDFirst = nil
	dst.SignOnPolicyActionIDP = nil
	dst.SignOnPolicyActionLogin = nil
	dst.SignOnPolicyActionMFA = nil
	dst.SignOnPolicyActionProgressiveProfiling = nil
	dst.SignOnPolicyActionPingIDWinLoginPasswordless = nil
	dst.SignOnPolicyActionCommon = nil

	switch common.GetType() {
	case ENUMSIGNONPOLICYTYPE_LOGIN:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionLogin); err != nil { // simple model
			return err
		}
	case ENUMSIGNONPOLICYTYPE_MULTI_FACTOR_AUTHENTICATION:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionMFA); err != nil { // simple model
			return err
		}
	case ENUMSIGNONPOLICYTYPE_IDENTIFIER_FIRST:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionIDFirst); err != nil { // simple model
			return err
		}
	case ENUMSIGNONPOLICYTYPE_IDENTITY_PROVIDER:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionIDP); err != nil { // simple model
			return err
		}
	case ENUMSIGNONPOLICYTYPE_PROGRESSIVE_PROFILING:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionProgressiveProfiling); err != nil { // simple model
			return err
		}
	case ENUMSIGNONPOLICYTYPE_AGREEMENT:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionAgreement); err != nil { // simple model
			return err
		}
	case ENUMSIGNONPOLICYTYPE_PINGID_WINLOGIN_PASSWORDLESS_AUTHENTICATION:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionPingIDWinLoginPasswordless); err != nil { // simple model
			return err
		}
	case ENUMSIGNONPOLICYTYPE_PINGID_AUTHENTICATION:
		if err := json.Unmarshal(data, &dst.SignOnPolicyActionCommon); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(SignOnPolicyAction)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SignOnPolicyAction) MarshalJSON() ([]byte, error) {
	if src.SignOnPolicyActionAgreement != nil {
		return json.Marshal(&src.SignOnPolicyActionAgreement)
	}

	if src.SignOnPolicyActionCommon != nil {
		return json.Marshal(&src.SignOnPolicyActionCommon)
	}

	if src.SignOnPolicyActionIDFirst != nil {
		return json.Marshal(&src.SignOnPolicyActionIDFirst)
	}

	if src.SignOnPolicyActionIDP != nil {
		return json.Marshal(&src.SignOnPolicyActionIDP)
	}

	if src.SignOnPolicyActionLogin != nil {
		return json.Marshal(&src.SignOnPolicyActionLogin)
	}

	if src.SignOnPolicyActionMFA != nil {
		return json.Marshal(&src.SignOnPolicyActionMFA)
	}

	if src.SignOnPolicyActionPingIDWinLoginPasswordless != nil {
		return json.Marshal(&src.SignOnPolicyActionPingIDWinLoginPasswordless)
	}

	if src.SignOnPolicyActionProgressiveProfiling != nil {
		return json.Marshal(&src.SignOnPolicyActionProgressiveProfiling)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SignOnPolicyAction) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SignOnPolicyActionAgreement != nil {
		return obj.SignOnPolicyActionAgreement
	}

	if obj.SignOnPolicyActionCommon != nil {
		return obj.SignOnPolicyActionCommon
	}

	if obj.SignOnPolicyActionIDFirst != nil {
		return obj.SignOnPolicyActionIDFirst
	}

	if obj.SignOnPolicyActionIDP != nil {
		return obj.SignOnPolicyActionIDP
	}

	if obj.SignOnPolicyActionLogin != nil {
		return obj.SignOnPolicyActionLogin
	}

	if obj.SignOnPolicyActionMFA != nil {
		return obj.SignOnPolicyActionMFA
	}

	if obj.SignOnPolicyActionPingIDWinLoginPasswordless != nil {
		return obj.SignOnPolicyActionPingIDWinLoginPasswordless
	}

	if obj.SignOnPolicyActionProgressiveProfiling != nil {
		return obj.SignOnPolicyActionProgressiveProfiling
	}

	// all schemas are nil
	return nil
}

type NullableSignOnPolicyAction struct {
	value *SignOnPolicyAction
	isSet bool
}

func (v NullableSignOnPolicyAction) Get() *SignOnPolicyAction {
	return v.value
}

func (v *NullableSignOnPolicyAction) Set(val *SignOnPolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyAction(val *SignOnPolicyAction) *NullableSignOnPolicyAction {
	return &NullableSignOnPolicyAction{value: val, isSet: true}
}

func (v NullableSignOnPolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
