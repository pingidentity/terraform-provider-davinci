/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// LicenseEnvironments struct for LicenseEnvironments
type LicenseEnvironments struct {
	AllowAddResources *bool `json:"allowAddResources,omitempty"`
	// A boolean that specifies whether the license supports creation of application connections in the specified environment.
	AllowConnections *bool `json:"allowConnections,omitempty"`
	// A read-only boolean that specifies whether the license supports creation of a custom domain in the specified environment.
	AllowCustomDomain *bool `json:"allowCustomDomain,omitempty"`
	// A read-only boolean that specifies whether the license supports using custom schema attributes in the specified environment.
	AllowCustomSchema *bool `json:"allowCustomSchema,omitempty"`
	// A read-only boolean that specifies whether production environments are allowed.
	AllowProduction *bool `json:"allowProduction,omitempty"`
	// A read-only integer that specifies the maximum number of environments allowed.
	Max *int32 `json:"max,omitempty"`
	Regions []EnumRegionCodeLicense `json:"regions,omitempty"`
}

// NewLicenseEnvironments instantiates a new LicenseEnvironments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseEnvironments() *LicenseEnvironments {
	this := LicenseEnvironments{}
	return &this
}

// NewLicenseEnvironmentsWithDefaults instantiates a new LicenseEnvironments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseEnvironmentsWithDefaults() *LicenseEnvironments {
	this := LicenseEnvironments{}
	return &this
}

// GetAllowAddResources returns the AllowAddResources field value if set, zero value otherwise.
func (o *LicenseEnvironments) GetAllowAddResources() bool {
	if o == nil || isNil(o.AllowAddResources) {
		var ret bool
		return ret
	}
	return *o.AllowAddResources
}

// GetAllowAddResourcesOk returns a tuple with the AllowAddResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEnvironments) GetAllowAddResourcesOk() (*bool, bool) {
	if o == nil || isNil(o.AllowAddResources) {
    return nil, false
	}
	return o.AllowAddResources, true
}

// HasAllowAddResources returns a boolean if a field has been set.
func (o *LicenseEnvironments) HasAllowAddResources() bool {
	if o != nil && !isNil(o.AllowAddResources) {
		return true
	}

	return false
}

// SetAllowAddResources gets a reference to the given bool and assigns it to the AllowAddResources field.
func (o *LicenseEnvironments) SetAllowAddResources(v bool) {
	o.AllowAddResources = &v
}

// GetAllowConnections returns the AllowConnections field value if set, zero value otherwise.
func (o *LicenseEnvironments) GetAllowConnections() bool {
	if o == nil || isNil(o.AllowConnections) {
		var ret bool
		return ret
	}
	return *o.AllowConnections
}

// GetAllowConnectionsOk returns a tuple with the AllowConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEnvironments) GetAllowConnectionsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowConnections) {
    return nil, false
	}
	return o.AllowConnections, true
}

// HasAllowConnections returns a boolean if a field has been set.
func (o *LicenseEnvironments) HasAllowConnections() bool {
	if o != nil && !isNil(o.AllowConnections) {
		return true
	}

	return false
}

// SetAllowConnections gets a reference to the given bool and assigns it to the AllowConnections field.
func (o *LicenseEnvironments) SetAllowConnections(v bool) {
	o.AllowConnections = &v
}

// GetAllowCustomDomain returns the AllowCustomDomain field value if set, zero value otherwise.
func (o *LicenseEnvironments) GetAllowCustomDomain() bool {
	if o == nil || isNil(o.AllowCustomDomain) {
		var ret bool
		return ret
	}
	return *o.AllowCustomDomain
}

// GetAllowCustomDomainOk returns a tuple with the AllowCustomDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEnvironments) GetAllowCustomDomainOk() (*bool, bool) {
	if o == nil || isNil(o.AllowCustomDomain) {
    return nil, false
	}
	return o.AllowCustomDomain, true
}

// HasAllowCustomDomain returns a boolean if a field has been set.
func (o *LicenseEnvironments) HasAllowCustomDomain() bool {
	if o != nil && !isNil(o.AllowCustomDomain) {
		return true
	}

	return false
}

// SetAllowCustomDomain gets a reference to the given bool and assigns it to the AllowCustomDomain field.
func (o *LicenseEnvironments) SetAllowCustomDomain(v bool) {
	o.AllowCustomDomain = &v
}

// GetAllowCustomSchema returns the AllowCustomSchema field value if set, zero value otherwise.
func (o *LicenseEnvironments) GetAllowCustomSchema() bool {
	if o == nil || isNil(o.AllowCustomSchema) {
		var ret bool
		return ret
	}
	return *o.AllowCustomSchema
}

// GetAllowCustomSchemaOk returns a tuple with the AllowCustomSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEnvironments) GetAllowCustomSchemaOk() (*bool, bool) {
	if o == nil || isNil(o.AllowCustomSchema) {
    return nil, false
	}
	return o.AllowCustomSchema, true
}

// HasAllowCustomSchema returns a boolean if a field has been set.
func (o *LicenseEnvironments) HasAllowCustomSchema() bool {
	if o != nil && !isNil(o.AllowCustomSchema) {
		return true
	}

	return false
}

// SetAllowCustomSchema gets a reference to the given bool and assigns it to the AllowCustomSchema field.
func (o *LicenseEnvironments) SetAllowCustomSchema(v bool) {
	o.AllowCustomSchema = &v
}

// GetAllowProduction returns the AllowProduction field value if set, zero value otherwise.
func (o *LicenseEnvironments) GetAllowProduction() bool {
	if o == nil || isNil(o.AllowProduction) {
		var ret bool
		return ret
	}
	return *o.AllowProduction
}

// GetAllowProductionOk returns a tuple with the AllowProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEnvironments) GetAllowProductionOk() (*bool, bool) {
	if o == nil || isNil(o.AllowProduction) {
    return nil, false
	}
	return o.AllowProduction, true
}

// HasAllowProduction returns a boolean if a field has been set.
func (o *LicenseEnvironments) HasAllowProduction() bool {
	if o != nil && !isNil(o.AllowProduction) {
		return true
	}

	return false
}

// SetAllowProduction gets a reference to the given bool and assigns it to the AllowProduction field.
func (o *LicenseEnvironments) SetAllowProduction(v bool) {
	o.AllowProduction = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *LicenseEnvironments) GetMax() int32 {
	if o == nil || isNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEnvironments) GetMaxOk() (*int32, bool) {
	if o == nil || isNil(o.Max) {
    return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *LicenseEnvironments) HasMax() bool {
	if o != nil && !isNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *LicenseEnvironments) SetMax(v int32) {
	o.Max = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *LicenseEnvironments) GetRegions() []EnumRegionCodeLicense {
	if o == nil || isNil(o.Regions) {
		var ret []EnumRegionCodeLicense
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseEnvironments) GetRegionsOk() ([]EnumRegionCodeLicense, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *LicenseEnvironments) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []EnumRegionCodeLicense and assigns it to the Regions field.
func (o *LicenseEnvironments) SetRegions(v []EnumRegionCodeLicense) {
	o.Regions = v
}

func (o LicenseEnvironments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowAddResources) {
		toSerialize["allowAddResources"] = o.AllowAddResources
	}
	if !isNil(o.AllowConnections) {
		toSerialize["allowConnections"] = o.AllowConnections
	}
	if !isNil(o.AllowCustomDomain) {
		toSerialize["allowCustomDomain"] = o.AllowCustomDomain
	}
	if !isNil(o.AllowCustomSchema) {
		toSerialize["allowCustomSchema"] = o.AllowCustomSchema
	}
	if !isNil(o.AllowProduction) {
		toSerialize["allowProduction"] = o.AllowProduction
	}
	if !isNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseEnvironments struct {
	value *LicenseEnvironments
	isSet bool
}

func (v NullableLicenseEnvironments) Get() *LicenseEnvironments {
	return v.value
}

func (v *NullableLicenseEnvironments) Set(val *LicenseEnvironments) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseEnvironments) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseEnvironments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseEnvironments(val *LicenseEnvironments) *NullableLicenseEnvironments {
	return &NullableLicenseEnvironments{value: val, isSet: true}
}

func (v NullableLicenseEnvironments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseEnvironments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


