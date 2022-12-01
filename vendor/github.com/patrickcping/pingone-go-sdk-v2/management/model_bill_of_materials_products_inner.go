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

// BillOfMaterialsProductsInner struct for BillOfMaterialsProductsInner
type BillOfMaterialsProductsInner struct {
	// A string that specifies the BOM ID
	Id *string `json:"id,omitempty"`
	Type EnumProductType `json:"type"`
	// A string that specifies the description of the product or standalone service
	Description *string `json:"description,omitempty"`
	Console *BillOfMaterialsProductsInnerConsole `json:"console,omitempty"`
	Deployment *BillOfMaterialsProductsInnerDeployment `json:"deployment,omitempty"`
	// Optional array of custom bookmarks. Maximum of five bookmarks per product.
	Bookmarks []BillOfMaterialsProductsInnerBookmarksInner `json:"bookmarks,omitempty"`
}

// NewBillOfMaterialsProductsInner instantiates a new BillOfMaterialsProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillOfMaterialsProductsInner(type_ EnumProductType) *BillOfMaterialsProductsInner {
	this := BillOfMaterialsProductsInner{}
	this.Type = type_
	return &this
}

// NewBillOfMaterialsProductsInnerWithDefaults instantiates a new BillOfMaterialsProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillOfMaterialsProductsInnerWithDefaults() *BillOfMaterialsProductsInner {
	this := BillOfMaterialsProductsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillOfMaterialsProductsInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillOfMaterialsProductsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillOfMaterialsProductsInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *BillOfMaterialsProductsInner) GetType() EnumProductType {
	if o == nil {
		var ret EnumProductType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInner) GetTypeOk() (*EnumProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BillOfMaterialsProductsInner) SetType(v EnumProductType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillOfMaterialsProductsInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillOfMaterialsProductsInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillOfMaterialsProductsInner) SetDescription(v string) {
	o.Description = &v
}

// GetConsole returns the Console field value if set, zero value otherwise.
func (o *BillOfMaterialsProductsInner) GetConsole() BillOfMaterialsProductsInnerConsole {
	if o == nil || o.Console == nil {
		var ret BillOfMaterialsProductsInnerConsole
		return ret
	}
	return *o.Console
}

// GetConsoleOk returns a tuple with the Console field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInner) GetConsoleOk() (*BillOfMaterialsProductsInnerConsole, bool) {
	if o == nil || o.Console == nil {
		return nil, false
	}
	return o.Console, true
}

// HasConsole returns a boolean if a field has been set.
func (o *BillOfMaterialsProductsInner) HasConsole() bool {
	if o != nil && o.Console != nil {
		return true
	}

	return false
}

// SetConsole gets a reference to the given BillOfMaterialsProductsInnerConsole and assigns it to the Console field.
func (o *BillOfMaterialsProductsInner) SetConsole(v BillOfMaterialsProductsInnerConsole) {
	o.Console = &v
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *BillOfMaterialsProductsInner) GetDeployment() BillOfMaterialsProductsInnerDeployment {
	if o == nil || o.Deployment == nil {
		var ret BillOfMaterialsProductsInnerDeployment
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInner) GetDeploymentOk() (*BillOfMaterialsProductsInnerDeployment, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *BillOfMaterialsProductsInner) HasDeployment() bool {
	if o != nil && o.Deployment != nil {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given BillOfMaterialsProductsInnerDeployment and assigns it to the Deployment field.
func (o *BillOfMaterialsProductsInner) SetDeployment(v BillOfMaterialsProductsInnerDeployment) {
	o.Deployment = &v
}

// GetBookmarks returns the Bookmarks field value if set, zero value otherwise.
func (o *BillOfMaterialsProductsInner) GetBookmarks() []BillOfMaterialsProductsInnerBookmarksInner {
	if o == nil || o.Bookmarks == nil {
		var ret []BillOfMaterialsProductsInnerBookmarksInner
		return ret
	}
	return o.Bookmarks
}

// GetBookmarksOk returns a tuple with the Bookmarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfMaterialsProductsInner) GetBookmarksOk() ([]BillOfMaterialsProductsInnerBookmarksInner, bool) {
	if o == nil || o.Bookmarks == nil {
		return nil, false
	}
	return o.Bookmarks, true
}

// HasBookmarks returns a boolean if a field has been set.
func (o *BillOfMaterialsProductsInner) HasBookmarks() bool {
	if o != nil && o.Bookmarks != nil {
		return true
	}

	return false
}

// SetBookmarks gets a reference to the given []BillOfMaterialsProductsInnerBookmarksInner and assigns it to the Bookmarks field.
func (o *BillOfMaterialsProductsInner) SetBookmarks(v []BillOfMaterialsProductsInnerBookmarksInner) {
	o.Bookmarks = v
}

func (o BillOfMaterialsProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Console != nil {
		toSerialize["console"] = o.Console
	}
	if o.Deployment != nil {
		toSerialize["deployment"] = o.Deployment
	}
	if o.Bookmarks != nil {
		toSerialize["bookmarks"] = o.Bookmarks
	}
	return json.Marshal(toSerialize)
}

type NullableBillOfMaterialsProductsInner struct {
	value *BillOfMaterialsProductsInner
	isSet bool
}

func (v NullableBillOfMaterialsProductsInner) Get() *BillOfMaterialsProductsInner {
	return v.value
}

func (v *NullableBillOfMaterialsProductsInner) Set(val *BillOfMaterialsProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBillOfMaterialsProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBillOfMaterialsProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillOfMaterialsProductsInner(val *BillOfMaterialsProductsInner) *NullableBillOfMaterialsProductsInner {
	return &NullableBillOfMaterialsProductsInner{value: val, isSet: true}
}

func (v NullableBillOfMaterialsProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillOfMaterialsProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


