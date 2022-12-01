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

// EntityArrayEmbeddedLanguagesInner - struct for EntityArrayEmbeddedLanguagesInner
type EntityArrayEmbeddedLanguagesInner struct {
	AgreementLanguage *AgreementLanguage
	Language          *Language
}

// AgreementLanguageAsEntityArrayEmbeddedLanguagesInner is a convenience function that returns AgreementLanguage wrapped in EntityArrayEmbeddedLanguagesInner
func AgreementLanguageAsEntityArrayEmbeddedLanguagesInner(v *AgreementLanguage) EntityArrayEmbeddedLanguagesInner {
	return EntityArrayEmbeddedLanguagesInner{
		AgreementLanguage: v,
	}
}

// LanguageAsEntityArrayEmbeddedLanguagesInner is a convenience function that returns Language wrapped in EntityArrayEmbeddedLanguagesInner
func LanguageAsEntityArrayEmbeddedLanguagesInner(v *Language) EntityArrayEmbeddedLanguagesInner {
	return EntityArrayEmbeddedLanguagesInner{
		Language: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EntityArrayEmbeddedLanguagesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0

	// try to unmarshal data into AgreementLanguage
	err = json.Unmarshal(data, &dst.AgreementLanguage)
	if err == nil {
		jsonAgreementLanguage, _ := json.Marshal(dst.AgreementLanguage)
		if string(jsonAgreementLanguage) == "{}" { // empty struct
			dst.AgreementLanguage = nil
		} else {
			if dst.AgreementLanguage.HasAgreement() {
				match++
			} else {
				dst.AgreementLanguage = nil
			}
		}
	} else {
		dst.AgreementLanguage = nil
	}

	// try to unmarshal data into Language
	err = json.Unmarshal(data, &dst.Language)
	if err == nil {
		jsonLanguage, _ := json.Marshal(dst.Language)
		if string(jsonLanguage) == "{}" { // empty struct
			dst.Language = nil
		} else {
			if dst.Language.HasCustomerAdded() {
				match++
			} else {
				dst.Language = nil
			}
		}
	} else {
		dst.Language = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AgreementLanguage = nil
		dst.Language = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(EntityArrayEmbeddedLanguagesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(EntityArrayEmbeddedLanguagesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntityArrayEmbeddedLanguagesInner) MarshalJSON() ([]byte, error) {
	if src.AgreementLanguage != nil {
		return json.Marshal(&src.AgreementLanguage)
	}

	if src.Language != nil {
		return json.Marshal(&src.Language)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntityArrayEmbeddedLanguagesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AgreementLanguage != nil {
		return obj.AgreementLanguage
	}

	if obj.Language != nil {
		return obj.Language
	}

	// all schemas are nil
	return nil
}

type NullableEntityArrayEmbeddedLanguagesInner struct {
	value *EntityArrayEmbeddedLanguagesInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedLanguagesInner) Get() *EntityArrayEmbeddedLanguagesInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedLanguagesInner) Set(val *EntityArrayEmbeddedLanguagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedLanguagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedLanguagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedLanguagesInner(val *EntityArrayEmbeddedLanguagesInner) *NullableEntityArrayEmbeddedLanguagesInner {
	return &NullableEntityArrayEmbeddedLanguagesInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedLanguagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedLanguagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
