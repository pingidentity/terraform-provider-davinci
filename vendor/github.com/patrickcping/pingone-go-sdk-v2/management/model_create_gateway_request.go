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

// CreateGatewayRequest - struct for CreateGatewayRequest
type CreateGatewayRequest struct {
	Gateway     *Gateway
	GatewayLDAP *GatewayLDAP
}

// GatewayAsCreateGatewayRequest is a convenience function that returns Gateway wrapped in CreateGatewayRequest
func GatewayAsCreateGatewayRequest(v *Gateway) CreateGatewayRequest {
	return CreateGatewayRequest{
		Gateway: v,
	}
}

// GatewayLDAPAsCreateGatewayRequest is a convenience function that returns GatewayLDAP wrapped in CreateGatewayRequest
func GatewayLDAPAsCreateGatewayRequest(v *GatewayLDAP) CreateGatewayRequest {
	return CreateGatewayRequest{
		GatewayLDAP: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateGatewayRequest) UnmarshalJSON(data []byte) error {

	var common Gateway

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.Gateway = nil
	dst.GatewayLDAP = nil

	switch common.GetType() {
	case ENUMGATEWAYTYPE_LDAP:
		if err := json.Unmarshal(data, &dst.GatewayLDAP); err != nil { // simple model
			return err
		}
	case ENUMGATEWAYTYPE_PING_FEDERATE:
		if err := json.Unmarshal(data, &dst.Gateway); err != nil { // simple model
			return err
		}
	case ENUMGATEWAYTYPE_API_GATEWAY_INTEGRATION:
		if err := json.Unmarshal(data, &dst.Gateway); err != nil { // simple model
			return err
		}
	case ENUMGATEWAYTYPE_PING_INTELLIGENCE:
		if err := json.Unmarshal(data, &dst.Gateway); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateGatewayRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateGatewayRequest) MarshalJSON() ([]byte, error) {
	if src.Gateway != nil {
		return json.Marshal(&src.Gateway)
	}

	if src.GatewayLDAP != nil {
		return json.Marshal(&src.GatewayLDAP)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateGatewayRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Gateway != nil {
		return obj.Gateway
	}

	if obj.GatewayLDAP != nil {
		return obj.GatewayLDAP
	}

	// all schemas are nil
	return nil
}

type NullableCreateGatewayRequest struct {
	value *CreateGatewayRequest
	isSet bool
}

func (v NullableCreateGatewayRequest) Get() *CreateGatewayRequest {
	return v.value
}

func (v *NullableCreateGatewayRequest) Set(val *CreateGatewayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGatewayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGatewayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGatewayRequest(val *CreateGatewayRequest) *NullableCreateGatewayRequest {
	return &NullableCreateGatewayRequest{value: val, isSet: true}
}

func (v NullableCreateGatewayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGatewayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
