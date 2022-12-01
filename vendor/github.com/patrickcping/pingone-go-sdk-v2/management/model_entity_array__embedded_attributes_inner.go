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

// EntityArrayEmbeddedAttributesInner - struct for EntityArrayEmbeddedAttributesInner
type EntityArrayEmbeddedAttributesInner struct {
	ApplicationAttributeMapping *ApplicationAttributeMapping
	IdentityProviderAttribute   *IdentityProviderAttribute
	ResourceAttribute           *ResourceAttribute
	SchemaAttribute             *SchemaAttribute
}

// ApplicationAttributeMappingAsEntityArrayEmbeddedAttributesInner is a convenience function that returns ApplicationAttributeMapping wrapped in EntityArrayEmbeddedAttributesInner
func ApplicationAttributeMappingAsEntityArrayEmbeddedAttributesInner(v *ApplicationAttributeMapping) EntityArrayEmbeddedAttributesInner {
	return EntityArrayEmbeddedAttributesInner{
		ApplicationAttributeMapping: v,
	}
}

// IdentityProviderAttributeAsEntityArrayEmbeddedAttributesInner is a convenience function that returns IdentityProviderAttribute wrapped in EntityArrayEmbeddedAttributesInner
func IdentityProviderAttributeAsEntityArrayEmbeddedAttributesInner(v *IdentityProviderAttribute) EntityArrayEmbeddedAttributesInner {
	return EntityArrayEmbeddedAttributesInner{
		IdentityProviderAttribute: v,
	}
}

// ResourceAttributeAsEntityArrayEmbeddedAttributesInner is a convenience function that returns ResourceAttribute wrapped in EntityArrayEmbeddedAttributesInner
func ResourceAttributeAsEntityArrayEmbeddedAttributesInner(v *ResourceAttribute) EntityArrayEmbeddedAttributesInner {
	return EntityArrayEmbeddedAttributesInner{
		ResourceAttribute: v,
	}
}

// SchemaAttributeAsEntityArrayEmbeddedAttributesInner is a convenience function that returns SchemaAttribute wrapped in EntityArrayEmbeddedAttributesInner
func SchemaAttributeAsEntityArrayEmbeddedAttributesInner(v *SchemaAttribute) EntityArrayEmbeddedAttributesInner {
	return EntityArrayEmbeddedAttributesInner{
		SchemaAttribute: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EntityArrayEmbeddedAttributesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApplicationAttributeMapping
	err = json.Unmarshal(data, &dst.ApplicationAttributeMapping)
	if err == nil {
		jsonApplicationAttributeMapping, _ := json.Marshal(dst.ApplicationAttributeMapping)
		if string(jsonApplicationAttributeMapping) == "{}" { // empty struct
			dst.ApplicationAttributeMapping = nil
		} else {
			if dst.ApplicationAttributeMapping.HasApplication() {
				match++
			} else {
				dst.ApplicationAttributeMapping = nil
			}
		}
	} else {
		dst.ApplicationAttributeMapping = nil
	}

	// try to unmarshal data into IdentityProviderAttribute
	err = json.Unmarshal(data, &dst.IdentityProviderAttribute)
	if err == nil {
		jsonIdentityProviderAttribute, _ := json.Marshal(dst.IdentityProviderAttribute)
		if string(jsonIdentityProviderAttribute) == "{}" { // empty struct
			dst.IdentityProviderAttribute = nil
		} else {
			if dst.IdentityProviderAttribute.HasIdentityProvider() {
				match++
			} else {
				dst.IdentityProviderAttribute = nil
			}
		}
	} else {
		dst.IdentityProviderAttribute = nil
	}

	// try to unmarshal data into ResourceAttribute
	err = json.Unmarshal(data, &dst.ResourceAttribute)
	if err == nil {
		jsonResourceAttribute, _ := json.Marshal(dst.ResourceAttribute)
		if string(jsonResourceAttribute) == "{}" { // empty struct
			dst.ResourceAttribute = nil
		} else {
			if dst.ResourceAttribute.HasResource() {
				match++
			} else {
				dst.ResourceAttribute = nil
			}
		}
	} else {
		dst.ResourceAttribute = nil
	}

	// try to unmarshal data into SchemaAttribute
	err = json.Unmarshal(data, &dst.SchemaAttribute)
	if err == nil {
		jsonSchemaAttribute, _ := json.Marshal(dst.SchemaAttribute)
		if string(jsonSchemaAttribute) == "{}" { // empty struct
			dst.SchemaAttribute = nil
		} else {
			if dst.SchemaAttribute.HasSchema() {
				match++
			} else {
				dst.SchemaAttribute = nil
			}
		}
	} else {
		dst.SchemaAttribute = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplicationAttributeMapping = nil
		dst.IdentityProviderAttribute = nil
		dst.ResourceAttribute = nil
		dst.SchemaAttribute = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(EntityArrayEmbeddedAttributesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(EntityArrayEmbeddedAttributesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntityArrayEmbeddedAttributesInner) MarshalJSON() ([]byte, error) {
	if src.ApplicationAttributeMapping != nil {
		return json.Marshal(&src.ApplicationAttributeMapping)
	}

	if src.IdentityProviderAttribute != nil {
		return json.Marshal(&src.IdentityProviderAttribute)
	}

	if src.ResourceAttribute != nil {
		return json.Marshal(&src.ResourceAttribute)
	}

	if src.SchemaAttribute != nil {
		return json.Marshal(&src.SchemaAttribute)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntityArrayEmbeddedAttributesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApplicationAttributeMapping != nil {
		return obj.ApplicationAttributeMapping
	}

	if obj.IdentityProviderAttribute != nil {
		return obj.IdentityProviderAttribute
	}

	if obj.ResourceAttribute != nil {
		return obj.ResourceAttribute
	}

	if obj.SchemaAttribute != nil {
		return obj.SchemaAttribute
	}

	// all schemas are nil
	return nil
}

type NullableEntityArrayEmbeddedAttributesInner struct {
	value *EntityArrayEmbeddedAttributesInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedAttributesInner) Get() *EntityArrayEmbeddedAttributesInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedAttributesInner) Set(val *EntityArrayEmbeddedAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedAttributesInner(val *EntityArrayEmbeddedAttributesInner) *NullableEntityArrayEmbeddedAttributesInner {
	return &NullableEntityArrayEmbeddedAttributesInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
