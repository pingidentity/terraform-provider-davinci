/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

package mfa

import (
	"encoding/json"
	"fmt"
)

// UpdateMFAPushCredentialRequest - struct for UpdateMFAPushCredentialRequest
type UpdateMFAPushCredentialRequest struct {
	MFAPushCredential     *MFAPushCredential
	MFAPushCredentialAPNS *MFAPushCredentialAPNS
}

// MFAPushCredentialAsUpdateMFAPushCredentialRequest is a convenience function that returns MFAPushCredential wrapped in UpdateMFAPushCredentialRequest
func MFAPushCredentialAsUpdateMFAPushCredentialRequest(v *MFAPushCredential) UpdateMFAPushCredentialRequest {
	return UpdateMFAPushCredentialRequest{
		MFAPushCredential: v,
	}
}

// MFAPushCredentialAPNSAsUpdateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialAPNS wrapped in UpdateMFAPushCredentialRequest
func MFAPushCredentialAPNSAsUpdateMFAPushCredentialRequest(v *MFAPushCredentialAPNS) UpdateMFAPushCredentialRequest {
	return UpdateMFAPushCredentialRequest{
		MFAPushCredentialAPNS: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateMFAPushCredentialRequest) UnmarshalJSON(data []byte) error {

	var common MFAPushCredential

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.MFAPushCredential = nil
	dst.MFAPushCredentialAPNS = nil

	switch common.GetType() {
	case ENUMMFAPUSHCREDENTIALATTRTYPE_FCM:
		if err := json.Unmarshal(data, &dst.MFAPushCredential); err != nil { // simple model
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_APNS:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialAPNS); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateMFAPushCredentialRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	if src.MFAPushCredential != nil {
		return json.Marshal(&src.MFAPushCredential)
	}

	if src.MFAPushCredentialAPNS != nil {
		return json.Marshal(&src.MFAPushCredentialAPNS)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateMFAPushCredentialRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MFAPushCredential != nil {
		return obj.MFAPushCredential
	}

	if obj.MFAPushCredentialAPNS != nil {
		return obj.MFAPushCredentialAPNS
	}

	// all schemas are nil
	return nil
}

type NullableUpdateMFAPushCredentialRequest struct {
	value *UpdateMFAPushCredentialRequest
	isSet bool
}

func (v NullableUpdateMFAPushCredentialRequest) Get() *UpdateMFAPushCredentialRequest {
	return v.value
}

func (v *NullableUpdateMFAPushCredentialRequest) Set(val *UpdateMFAPushCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMFAPushCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMFAPushCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMFAPushCredentialRequest(val *UpdateMFAPushCredentialRequest) *NullableUpdateMFAPushCredentialRequest {
	return &NullableUpdateMFAPushCredentialRequest{value: val, isSet: true}
}

func (v NullableUpdateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMFAPushCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
