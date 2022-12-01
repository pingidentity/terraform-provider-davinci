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

// CreateMFAPushCredentialRequest - struct for CreateMFAPushCredentialRequest
type CreateMFAPushCredentialRequest struct {
	MFAPushCredential     *MFAPushCredential
	MFAPushCredentialAPNS *MFAPushCredentialAPNS
}

// MFAPushCredentialAsCreateMFAPushCredentialRequest is a convenience function that returns MFAPushCredential wrapped in CreateMFAPushCredentialRequest
func MFAPushCredentialAsCreateMFAPushCredentialRequest(v *MFAPushCredential) CreateMFAPushCredentialRequest {
	return CreateMFAPushCredentialRequest{
		MFAPushCredential: v,
	}
}

// MFAPushCredentialAPNSAsCreateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialAPNS wrapped in CreateMFAPushCredentialRequest
func MFAPushCredentialAPNSAsCreateMFAPushCredentialRequest(v *MFAPushCredentialAPNS) CreateMFAPushCredentialRequest {
	return CreateMFAPushCredentialRequest{
		MFAPushCredentialAPNS: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateMFAPushCredentialRequest) UnmarshalJSON(data []byte) error {

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
		return fmt.Errorf("Data failed to match schemas in anyOf(CreateMFAPushCredential201Response)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	if src.MFAPushCredential != nil {
		return json.Marshal(&src.MFAPushCredential)
	}

	if src.MFAPushCredentialAPNS != nil {
		return json.Marshal(&src.MFAPushCredentialAPNS)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateMFAPushCredentialRequest) GetActualInstance() interface{} {
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

type NullableCreateMFAPushCredentialRequest struct {
	value *CreateMFAPushCredentialRequest
	isSet bool
}

func (v NullableCreateMFAPushCredentialRequest) Get() *CreateMFAPushCredentialRequest {
	return v.value
}

func (v *NullableCreateMFAPushCredentialRequest) Set(val *CreateMFAPushCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMFAPushCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMFAPushCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMFAPushCredentialRequest(val *CreateMFAPushCredentialRequest) *NullableCreateMFAPushCredentialRequest {
	return &NullableCreateMFAPushCredentialRequest{value: val, isSet: true}
}

func (v NullableCreateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMFAPushCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
