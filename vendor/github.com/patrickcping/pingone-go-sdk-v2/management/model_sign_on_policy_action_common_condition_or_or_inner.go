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

// SignOnPolicyActionCommonConditionOrOrInner - struct for SignOnPolicyActionCommonConditionOrOrInner
type SignOnPolicyActionCommonConditionOrOrInner struct {
	SignOnPolicyActionCommonConditionAggregate *SignOnPolicyActionCommonConditionAggregate
	SignOnPolicyActionCommonConditionAnd       *SignOnPolicyActionCommonConditionAnd
	SignOnPolicyActionCommonConditionNot       *SignOnPolicyActionCommonConditionNot
	SignOnPolicyActionCommonConditionOr        *SignOnPolicyActionCommonConditionOr
}

// SignOnPolicyActionCommonConditionAggregateAsSignOnPolicyActionCommonConditionOrOrInner is a convenience function that returns SignOnPolicyActionCommonConditionAggregate wrapped in SignOnPolicyActionCommonConditionOrOrInner
func SignOnPolicyActionCommonConditionAggregateAsSignOnPolicyActionCommonConditionOrOrInner(v *SignOnPolicyActionCommonConditionAggregate) SignOnPolicyActionCommonConditionOrOrInner {
	return SignOnPolicyActionCommonConditionOrOrInner{
		SignOnPolicyActionCommonConditionAggregate: v,
	}
}

// SignOnPolicyActionCommonConditionAndAsSignOnPolicyActionCommonConditionOrOrInner is a convenience function that returns SignOnPolicyActionCommonConditionAnd wrapped in SignOnPolicyActionCommonConditionOrOrInner
func SignOnPolicyActionCommonConditionAndAsSignOnPolicyActionCommonConditionOrOrInner(v *SignOnPolicyActionCommonConditionAnd) SignOnPolicyActionCommonConditionOrOrInner {
	return SignOnPolicyActionCommonConditionOrOrInner{
		SignOnPolicyActionCommonConditionAnd: v,
	}
}

// SignOnPolicyActionCommonConditionNotAsSignOnPolicyActionCommonConditionOrOrInner is a convenience function that returns SignOnPolicyActionCommonConditionNot wrapped in SignOnPolicyActionCommonConditionOrOrInner
func SignOnPolicyActionCommonConditionNotAsSignOnPolicyActionCommonConditionOrOrInner(v *SignOnPolicyActionCommonConditionNot) SignOnPolicyActionCommonConditionOrOrInner {
	return SignOnPolicyActionCommonConditionOrOrInner{
		SignOnPolicyActionCommonConditionNot: v,
	}
}

// SignOnPolicyActionCommonConditionOrAsSignOnPolicyActionCommonConditionOrOrInner is a convenience function that returns SignOnPolicyActionCommonConditionOr wrapped in SignOnPolicyActionCommonConditionOrOrInner
func SignOnPolicyActionCommonConditionOrAsSignOnPolicyActionCommonConditionOrOrInner(v *SignOnPolicyActionCommonConditionOr) SignOnPolicyActionCommonConditionOrOrInner {
	return SignOnPolicyActionCommonConditionOrOrInner{
		SignOnPolicyActionCommonConditionOr: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SignOnPolicyActionCommonConditionOrOrInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0

	// try to unmarshal data into SignOnPolicyActionCommonConditionAnd
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionAnd)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionAnd, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionAnd)
		if string(jsonSignOnPolicyActionCommonConditionAnd) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionAnd = nil
		} else {
			if dst.SignOnPolicyActionCommonConditionAnd.HasAnd() {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionAnd = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionAnd = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionNot
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionNot)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionNot, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionNot)
		if string(jsonSignOnPolicyActionCommonConditionNot) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionNot = nil
		} else {
			if dst.SignOnPolicyActionCommonConditionNot.HasNot() {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionNot = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionNot = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionOr
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionOr)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionOr, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionOr)
		if string(jsonSignOnPolicyActionCommonConditionOr) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionOr = nil
		} else {
			if dst.SignOnPolicyActionCommonConditionOr.HasOr() {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionOr = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionOr = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionAggregate
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionAggregate)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionAggregate, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionAggregate)
		if string(jsonSignOnPolicyActionCommonConditionAggregate) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionAggregate = nil
		} else {
			match++
		}
	} else {
		dst.SignOnPolicyActionCommonConditionAggregate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SignOnPolicyActionCommonConditionAggregate = nil
		dst.SignOnPolicyActionCommonConditionAnd = nil
		dst.SignOnPolicyActionCommonConditionNot = nil
		dst.SignOnPolicyActionCommonConditionOr = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SignOnPolicyActionCommonConditionOrOrInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SignOnPolicyActionCommonConditionOrOrInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SignOnPolicyActionCommonConditionOrOrInner) MarshalJSON() ([]byte, error) {
	if src.SignOnPolicyActionCommonConditionAggregate != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionAggregate)
	}

	if src.SignOnPolicyActionCommonConditionAnd != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionAnd)
	}

	if src.SignOnPolicyActionCommonConditionNot != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionNot)
	}

	if src.SignOnPolicyActionCommonConditionOr != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionOr)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SignOnPolicyActionCommonConditionOrOrInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SignOnPolicyActionCommonConditionAggregate != nil {
		return obj.SignOnPolicyActionCommonConditionAggregate
	}

	if obj.SignOnPolicyActionCommonConditionAnd != nil {
		return obj.SignOnPolicyActionCommonConditionAnd
	}

	if obj.SignOnPolicyActionCommonConditionNot != nil {
		return obj.SignOnPolicyActionCommonConditionNot
	}

	if obj.SignOnPolicyActionCommonConditionOr != nil {
		return obj.SignOnPolicyActionCommonConditionOr
	}

	// all schemas are nil
	return nil
}

type NullableSignOnPolicyActionCommonConditionOrOrInner struct {
	value *SignOnPolicyActionCommonConditionOrOrInner
	isSet bool
}

func (v NullableSignOnPolicyActionCommonConditionOrOrInner) Get() *SignOnPolicyActionCommonConditionOrOrInner {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonConditionOrOrInner) Set(val *SignOnPolicyActionCommonConditionOrOrInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonConditionOrOrInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonConditionOrOrInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonConditionOrOrInner(val *SignOnPolicyActionCommonConditionOrOrInner) *NullableSignOnPolicyActionCommonConditionOrOrInner {
	return &NullableSignOnPolicyActionCommonConditionOrOrInner{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonConditionOrOrInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonConditionOrOrInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
