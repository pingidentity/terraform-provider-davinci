/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

package model

import (
	"encoding/json"
)

// P1Error struct for P1Error
type P1Error struct {
	// A unique identifier that is stored in log files and always included in an error response. This value can be used to track the error received by the client, with server-side activity included for troubleshooting purposes.
	Id *string `json:"id,omitempty"`
	// A general fault code which the client must handle to provide all exception handling routines and to localize messages for users. This code is common across all PingOne services and is human readable (such as a defined constant rather than a number).
	Code *string `json:"code,omitempty"`
	// A short description of the error. This message is intended to assist with debugging and is returned in English only.
	Message *string `json:"message,omitempty"`
	// Additional details about the error. Optional information to help resolve the error and to display to users.
	Details []P1ErrorDetailsInner `json:"details,omitempty"`
}

// NewP1Error instantiates a new P1Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP1Error() *P1Error {
	this := P1Error{}
	return &this
}

// NewP1ErrorWithDefaults instantiates a new P1Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP1ErrorWithDefaults() *P1Error {
	this := P1Error{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *P1Error) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1Error) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *P1Error) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *P1Error) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *P1Error) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1Error) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *P1Error) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *P1Error) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *P1Error) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1Error) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *P1Error) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *P1Error) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *P1Error) GetDetails() []P1ErrorDetailsInner {
	if o == nil || o.Details == nil {
		var ret []P1ErrorDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1Error) GetDetailsOk() ([]P1ErrorDetailsInner, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *P1Error) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []P1ErrorDetailsInner and assigns it to the Details field.
func (o *P1Error) SetDetails(v []P1ErrorDetailsInner) {
	o.Details = v
}

func (o P1Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableP1Error struct {
	value *P1Error
	isSet bool
}

func (v NullableP1Error) Get() *P1Error {
	return v.value
}

func (v *NullableP1Error) Set(val *P1Error) {
	v.value = val
	v.isSet = true
}

func (v NullableP1Error) IsSet() bool {
	return v.isSet
}

func (v *NullableP1Error) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP1Error(val *P1Error) *NullableP1Error {
	return &NullableP1Error{value: val, isSet: true}
}

func (v NullableP1Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP1Error) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
