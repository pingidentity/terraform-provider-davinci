/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumCertificateKeyAlgorithm Specifies the key algorithm. Options are `RSA`, `EC`, and `UNKNOWN`.
type EnumCertificateKeyAlgorithm string

// List of EnumCertificateKeyAlgorithm
const (
	ENUMCERTIFICATEKEYALGORITHM_RSA EnumCertificateKeyAlgorithm = "RSA"
	ENUMCERTIFICATEKEYALGORITHM_EC EnumCertificateKeyAlgorithm = "EC"
	ENUMCERTIFICATEKEYALGORITHM_UNKNOWN EnumCertificateKeyAlgorithm = "UNKNOWN"
)

// All allowed values of EnumCertificateKeyAlgorithm enum
var AllowedEnumCertificateKeyAlgorithmEnumValues = []EnumCertificateKeyAlgorithm{
	"RSA",
	"EC",
	"UNKNOWN",
}

func (v *EnumCertificateKeyAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCertificateKeyAlgorithm(value)
	for _, existing := range AllowedEnumCertificateKeyAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCertificateKeyAlgorithm", value)
}

// NewEnumCertificateKeyAlgorithmFromValue returns a pointer to a valid EnumCertificateKeyAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCertificateKeyAlgorithmFromValue(v string) (*EnumCertificateKeyAlgorithm, error) {
	ev := EnumCertificateKeyAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCertificateKeyAlgorithm: valid values are %v", v, AllowedEnumCertificateKeyAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCertificateKeyAlgorithm) IsValid() bool {
	for _, existing := range AllowedEnumCertificateKeyAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCertificateKeyAlgorithm value
func (v EnumCertificateKeyAlgorithm) Ptr() *EnumCertificateKeyAlgorithm {
	return &v
}

type NullableEnumCertificateKeyAlgorithm struct {
	value *EnumCertificateKeyAlgorithm
	isSet bool
}

func (v NullableEnumCertificateKeyAlgorithm) Get() *EnumCertificateKeyAlgorithm {
	return v.value
}

func (v *NullableEnumCertificateKeyAlgorithm) Set(val *EnumCertificateKeyAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCertificateKeyAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCertificateKeyAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCertificateKeyAlgorithm(val *EnumCertificateKeyAlgorithm) *NullableEnumCertificateKeyAlgorithm {
	return &NullableEnumCertificateKeyAlgorithm{value: val, isSet: true}
}

func (v NullableEnumCertificateKeyAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCertificateKeyAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
