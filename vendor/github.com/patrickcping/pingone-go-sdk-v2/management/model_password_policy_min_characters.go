/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

package management

import (
	"encoding/json"
)

// PasswordPolicyMinCharacters A set of key-value pairs where the key is contains all characters that can be included, and the value is the minimum number of times one of the characters must appear in the password. The only allowed key values are `ABCDEFGHIJKLMNOPQRSTUVWXYZ`, `abcdefghijklmnopqrstuvwxyz`, `0123456789`, and `~!@#$%^&*()-_=+[]{}\\|;:,.<>/?`. This property is not enforced when not present.
type PasswordPolicyMinCharacters struct {
	// Count of alphabetical uppercase characters (`ABCDEFGHIJKLMNOPQRSTUVWXYZ`) that should feature in the user's password.
	ABCDEFGHIJKLMNOPQRSTUVWXYZ *int32 `json:"ABCDEFGHIJKLMNOPQRSTUVWXYZ,omitempty"`
	// Count of alphabetical uppercase characters (`abcdefghijklmnopqrstuvwxyz`) that should feature in the user's password.
	Abcdefghijklmnopqrstuvwxyz *int32 `json:"abcdefghijklmnopqrstuvwxyz,omitempty"`
	// Count of numeric characters (`0123456789`) that should feature in the user's password.
	Var0123456789 *int32 `json:"0123456789,omitempty"`
	// Count of special characters (`~!@#$%^&*()-_=+[]{}|;:,.<>/?`) that should feature in the user's password.
	SpecialChar *int32 `json:"specialchar,omitempty"`
}

// NewPasswordPolicyMinCharacters instantiates a new PasswordPolicyMinCharacters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyMinCharacters() *PasswordPolicyMinCharacters {
	this := PasswordPolicyMinCharacters{}
	return &this
}

// NewPasswordPolicyMinCharactersWithDefaults instantiates a new PasswordPolicyMinCharacters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyMinCharactersWithDefaults() *PasswordPolicyMinCharacters {
	this := PasswordPolicyMinCharacters{}
	return &this
}

// GetABCDEFGHIJKLMNOPQRSTUVWXYZ returns the ABCDEFGHIJKLMNOPQRSTUVWXYZ field value if set, zero value otherwise.
func (o *PasswordPolicyMinCharacters) GetABCDEFGHIJKLMNOPQRSTUVWXYZ() int32 {
	if o == nil || isNil(o.ABCDEFGHIJKLMNOPQRSTUVWXYZ) {
		var ret int32
		return ret
	}
	return *o.ABCDEFGHIJKLMNOPQRSTUVWXYZ
}

// GetABCDEFGHIJKLMNOPQRSTUVWXYZOk returns a tuple with the ABCDEFGHIJKLMNOPQRSTUVWXYZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyMinCharacters) GetABCDEFGHIJKLMNOPQRSTUVWXYZOk() (*int32, bool) {
	if o == nil || isNil(o.ABCDEFGHIJKLMNOPQRSTUVWXYZ) {
		return nil, false
	}
	return o.ABCDEFGHIJKLMNOPQRSTUVWXYZ, true
}

// HasABCDEFGHIJKLMNOPQRSTUVWXYZ returns a boolean if a field has been set.
func (o *PasswordPolicyMinCharacters) HasABCDEFGHIJKLMNOPQRSTUVWXYZ() bool {
	if o != nil && !isNil(o.ABCDEFGHIJKLMNOPQRSTUVWXYZ) {
		return true
	}

	return false
}

// SetABCDEFGHIJKLMNOPQRSTUVWXYZ gets a reference to the given int32 and assigns it to the ABCDEFGHIJKLMNOPQRSTUVWXYZ field.
func (o *PasswordPolicyMinCharacters) SetABCDEFGHIJKLMNOPQRSTUVWXYZ(v int32) {
	o.ABCDEFGHIJKLMNOPQRSTUVWXYZ = &v
}

// GetAbcdefghijklmnopqrstuvwxyz returns the Abcdefghijklmnopqrstuvwxyz field value if set, zero value otherwise.
func (o *PasswordPolicyMinCharacters) GetAbcdefghijklmnopqrstuvwxyz() int32 {
	if o == nil || isNil(o.Abcdefghijklmnopqrstuvwxyz) {
		var ret int32
		return ret
	}
	return *o.Abcdefghijklmnopqrstuvwxyz
}

// GetAbcdefghijklmnopqrstuvwxyzOk returns a tuple with the Abcdefghijklmnopqrstuvwxyz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyMinCharacters) GetAbcdefghijklmnopqrstuvwxyzOk() (*int32, bool) {
	if o == nil || isNil(o.Abcdefghijklmnopqrstuvwxyz) {
		return nil, false
	}
	return o.Abcdefghijklmnopqrstuvwxyz, true
}

// HasAbcdefghijklmnopqrstuvwxyz returns a boolean if a field has been set.
func (o *PasswordPolicyMinCharacters) HasAbcdefghijklmnopqrstuvwxyz() bool {
	if o != nil && !isNil(o.Abcdefghijklmnopqrstuvwxyz) {
		return true
	}

	return false
}

// SetAbcdefghijklmnopqrstuvwxyz gets a reference to the given int32 and assigns it to the Abcdefghijklmnopqrstuvwxyz field.
func (o *PasswordPolicyMinCharacters) SetAbcdefghijklmnopqrstuvwxyz(v int32) {
	o.Abcdefghijklmnopqrstuvwxyz = &v
}

// GetVar0123456789 returns the Var0123456789 field value if set, zero value otherwise.
func (o *PasswordPolicyMinCharacters) GetVar0123456789() int32 {
	if o == nil || isNil(o.Var0123456789) {
		var ret int32
		return ret
	}
	return *o.Var0123456789
}

// GetVar0123456789Ok returns a tuple with the Var0123456789 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyMinCharacters) GetVar0123456789Ok() (*int32, bool) {
	if o == nil || isNil(o.Var0123456789) {
		return nil, false
	}
	return o.Var0123456789, true
}

// HasVar0123456789 returns a boolean if a field has been set.
func (o *PasswordPolicyMinCharacters) HasVar0123456789() bool {
	if o != nil && !isNil(o.Var0123456789) {
		return true
	}

	return false
}

// SetVar0123456789 gets a reference to the given int32 and assigns it to the Var0123456789 field.
func (o *PasswordPolicyMinCharacters) SetVar0123456789(v int32) {
	o.Var0123456789 = &v
}

// GetSpecialChar returns the SpecialChar field value if set, zero value otherwise.
func (o *PasswordPolicyMinCharacters) GetSpecialChar() int32 {
	if o == nil || o.SpecialChar == nil {
		var ret int32
		return ret
	}
	return *o.SpecialChar
}

// GetSpecialCharOk returns a tuple with the SpecialChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyMinCharacters) GetSpecialCharOk() (*int32, bool) {
	if o == nil || isNil(o.SpecialChar) {
		return nil, false
	}
	return o.SpecialChar, true
}

// HasSpecialChar returns a boolean if a field has been set.
func (o *PasswordPolicyMinCharacters) HasSpecialChar() bool {
	if o != nil && !isNil(o.SpecialChar) {
		return true
	}

	return false
}

// SetSpecialChar gets a reference to the given int32 and assigns it to the SpecialChar field.
func (o *PasswordPolicyMinCharacters) SetSpecialChar(v int32) {
	o.SpecialChar = &v
}

func (o PasswordPolicyMinCharacters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ABCDEFGHIJKLMNOPQRSTUVWXYZ) {
		toSerialize["ABCDEFGHIJKLMNOPQRSTUVWXYZ"] = o.ABCDEFGHIJKLMNOPQRSTUVWXYZ
	}
	if !isNil(o.Abcdefghijklmnopqrstuvwxyz) {
		toSerialize["abcdefghijklmnopqrstuvwxyz"] = o.Abcdefghijklmnopqrstuvwxyz
	}
	if !isNil(o.Var0123456789) {
		toSerialize["0123456789"] = o.Var0123456789
	}
	if !isNil(o.SpecialChar) {
		toSerialize["~!@#$%^&*()-_=+[]{}|;:,.<>/?"] = o.SpecialChar
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordPolicyMinCharacters struct {
	value *PasswordPolicyMinCharacters
	isSet bool
}

func (v NullablePasswordPolicyMinCharacters) Get() *PasswordPolicyMinCharacters {
	return v.value
}

func (v *NullablePasswordPolicyMinCharacters) Set(val *PasswordPolicyMinCharacters) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyMinCharacters) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyMinCharacters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyMinCharacters(val *PasswordPolicyMinCharacters) *NullablePasswordPolicyMinCharacters {
	return &NullablePasswordPolicyMinCharacters{value: val, isSet: true}
}

func (v NullablePasswordPolicyMinCharacters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyMinCharacters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
