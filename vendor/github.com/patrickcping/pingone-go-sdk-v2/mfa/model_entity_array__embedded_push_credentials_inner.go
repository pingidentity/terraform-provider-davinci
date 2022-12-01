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

// EntityArrayEmbeddedPushCredentialsInner struct for EntityArrayEmbeddedPushCredentialsInner
type EntityArrayEmbeddedPushCredentialsInner struct {
	MFAPushCredential     *MFAPushCredential
	MFAPushCredentialAPNS *MFAPushCredentialAPNS
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EntityArrayEmbeddedPushCredentialsInner) UnmarshalJSON(data []byte) error {

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
		return fmt.Errorf("Data failed to match schemas in anyOf(EntityArrayEmbeddedPushCredentialsInner)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EntityArrayEmbeddedPushCredentialsInner) MarshalJSON() ([]byte, error) {
	if src.MFAPushCredential != nil {
		return json.Marshal(&src.MFAPushCredential)
	}

	if src.MFAPushCredentialAPNS != nil {
		return json.Marshal(&src.MFAPushCredentialAPNS)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEntityArrayEmbeddedPushCredentialsInner struct {
	value *EntityArrayEmbeddedPushCredentialsInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedPushCredentialsInner) Get() *EntityArrayEmbeddedPushCredentialsInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedPushCredentialsInner) Set(val *EntityArrayEmbeddedPushCredentialsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedPushCredentialsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedPushCredentialsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedPushCredentialsInner(val *EntityArrayEmbeddedPushCredentialsInner) *NullableEntityArrayEmbeddedPushCredentialsInner {
	return &NullableEntityArrayEmbeddedPushCredentialsInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedPushCredentialsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedPushCredentialsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
