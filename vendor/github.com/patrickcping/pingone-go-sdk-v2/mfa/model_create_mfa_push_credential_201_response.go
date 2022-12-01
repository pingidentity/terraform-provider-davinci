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

// CreateMFAPushCredential201Response - struct for CreateMFAPushCredential201Response
type CreateMFAPushCredential201Response struct {
	MFAPushCredential     *MFAPushCredential
	MFAPushCredentialAPNS *MFAPushCredentialAPNS
}

// MFAPushCredentialAsCreateMFAPushCredential201Response is a convenience function that returns MFAPushCredential wrapped in CreateMFAPushCredential201Response
func MFAPushCredentialAsCreateMFAPushCredential201Response(v *MFAPushCredential) CreateMFAPushCredential201Response {
	return CreateMFAPushCredential201Response{
		MFAPushCredential: v,
	}
}

// MFAPushCredentialAPNSAsCreateMFAPushCredential201Response is a convenience function that returns MFAPushCredentialAPNS wrapped in CreateMFAPushCredential201Response
func MFAPushCredentialAPNSAsCreateMFAPushCredential201Response(v *MFAPushCredentialAPNS) CreateMFAPushCredential201Response {
	return CreateMFAPushCredential201Response{
		MFAPushCredentialAPNS: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateMFAPushCredential201Response) UnmarshalJSON(data []byte) error {
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
func (src CreateMFAPushCredential201Response) MarshalJSON() ([]byte, error) {
	if src.MFAPushCredential != nil {
		return json.Marshal(&src.MFAPushCredential)
	}

	if src.MFAPushCredentialAPNS != nil {
		return json.Marshal(&src.MFAPushCredentialAPNS)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateMFAPushCredential201Response) GetActualInstance() interface{} {
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

type NullableCreateMFAPushCredential201Response struct {
	value *CreateMFAPushCredential201Response
	isSet bool
}

func (v NullableCreateMFAPushCredential201Response) Get() *CreateMFAPushCredential201Response {
	return v.value
}

func (v *NullableCreateMFAPushCredential201Response) Set(val *CreateMFAPushCredential201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMFAPushCredential201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMFAPushCredential201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMFAPushCredential201Response(val *CreateMFAPushCredential201Response) *NullableCreateMFAPushCredential201Response {
	return &NullableCreateMFAPushCredential201Response{value: val, isSet: true}
}

func (v NullableCreateMFAPushCredential201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMFAPushCredential201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
