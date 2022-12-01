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

// CreateGateway201Response - struct for CreateGateway201Response
type CreateGateway201Response struct {
	Gateway     *Gateway
	GatewayLDAP *GatewayLDAP
}

// GatewayAsCreateGateway201Response is a convenience function that returns Gateway wrapped in CreateGateway201Response
func GatewayAsCreateGateway201Response(v *Gateway) CreateGateway201Response {
	return CreateGateway201Response{
		Gateway: v,
	}
}

// GatewayLDAPAsCreateGateway201Response is a convenience function that returns GatewayLDAP wrapped in CreateGateway201Response
func GatewayLDAPAsCreateGateway201Response(v *GatewayLDAP) CreateGateway201Response {
	return CreateGateway201Response{
		GatewayLDAP: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateGateway201Response) UnmarshalJSON(data []byte) error {

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
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateGateway201Response)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateGateway201Response) MarshalJSON() ([]byte, error) {
	if src.Gateway != nil {
		return json.Marshal(&src.Gateway)
	}

	if src.GatewayLDAP != nil {
		return json.Marshal(&src.GatewayLDAP)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateGateway201Response) GetActualInstance() interface{} {
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

type NullableCreateGateway201Response struct {
	value *CreateGateway201Response
	isSet bool
}

func (v NullableCreateGateway201Response) Get() *CreateGateway201Response {
	return v.value
}

func (v *NullableCreateGateway201Response) Set(val *CreateGateway201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGateway201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGateway201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGateway201Response(val *CreateGateway201Response) *NullableCreateGateway201Response {
	return &NullableCreateGateway201Response{value: val, isSet: true}
}

func (v NullableCreateGateway201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGateway201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
