/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

package management

import (
	"encoding/json"
	"math/big"
	"time"
)

// Certificate struct for Certificate
type Certificate struct {
	Algorithm EnumCertificateKeyAlgorithm `json:"algorithm"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Specifies whether this is the default key for the specified environment.
	Default     *bool              `json:"default,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The time the key resource expires.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// Specifies the distinguished name of the certificate issuer.
	IssuerDN *string `json:"issuerDN,omitempty"`
	// Specifies the key length. For RSA keys, options are 2048, 3072, and 7680. For elliptical curve (EC) keys, options are 224, 256, and 384.
	KeyLength int32 `json:"keyLength"`
	// Specifies the resource name.
	Name         string              `json:"name"`
	Organization *ObjectOrganization `json:"organization,omitempty"`
	// Specifies the serial number of the key or certificate.
	SerialNumber       *big.Int                             `json:"serialNumber,omitempty"`
	SignatureAlgorithm EnumCertificateKeySignagureAlgorithm `json:"signatureAlgorithm"`
	// The time the validity period starts.
	StartsAt *time.Time                `json:"startsAt,omitempty"`
	Status   *EnumCertificateKeyStatus `json:"status,omitempty"`
	// Specifies the distinguished name of the subject being secured.
	SubjectDN string                      `json:"subjectDN"`
	UsageType EnumCertificateKeyUsageType `json:"usageType"`
	// Specifies the number of days the key is valid.
	ValidityPeriod int32 `json:"validityPeriod"`
}

// NewCertificate instantiates a new Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificate(algorithm EnumCertificateKeyAlgorithm, keyLength int32, name string, signatureAlgorithm EnumCertificateKeySignagureAlgorithm, subjectDN string, usageType EnumCertificateKeyUsageType, validityPeriod int32) *Certificate {
	this := Certificate{}
	this.Algorithm = algorithm
	this.KeyLength = keyLength
	this.Name = name
	this.SignatureAlgorithm = signatureAlgorithm
	this.SubjectDN = subjectDN
	this.UsageType = usageType
	this.ValidityPeriod = validityPeriod
	return &this
}

// NewCertificateWithDefaults instantiates a new Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateWithDefaults() *Certificate {
	this := Certificate{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *Certificate) GetAlgorithm() EnumCertificateKeyAlgorithm {
	if o == nil {
		var ret EnumCertificateKeyAlgorithm
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetAlgorithmOk() (*EnumCertificateKeyAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *Certificate) SetAlgorithm(v EnumCertificateKeyAlgorithm) {
	o.Algorithm = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Certificate) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Certificate) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Certificate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Certificate) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Certificate) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *Certificate) SetDefault(v bool) {
	o.Default = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Certificate) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Certificate) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Certificate) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Certificate) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Certificate) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *Certificate) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Certificate) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Certificate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Certificate) SetId(v string) {
	o.Id = &v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise.
func (o *Certificate) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN) {
		var ret string
		return ret
	}
	return *o.IssuerDN
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetIssuerDNOk() (*string, bool) {
	if o == nil || isNil(o.IssuerDN) {
		return nil, false
	}
	return o.IssuerDN, true
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *Certificate) HasIssuerDN() bool {
	if o != nil && !isNil(o.IssuerDN) {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given string and assigns it to the IssuerDN field.
func (o *Certificate) SetIssuerDN(v string) {
	o.IssuerDN = &v
}

// GetKeyLength returns the KeyLength field value
func (o *Certificate) GetKeyLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KeyLength
}

// GetKeyLengthOk returns a tuple with the KeyLength field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetKeyLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyLength, true
}

// SetKeyLength sets field value
func (o *Certificate) SetKeyLength(v int32) {
	o.KeyLength = v
}

// GetName returns the Name field value
func (o *Certificate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Certificate) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Certificate) GetOrganization() ObjectOrganization {
	if o == nil || isNil(o.Organization) {
		var ret ObjectOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetOrganizationOk() (*ObjectOrganization, bool) {
	if o == nil || isNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Certificate) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given ObjectOrganization and assigns it to the Organization field.
func (o *Certificate) SetOrganization(v ObjectOrganization) {
	o.Organization = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Certificate) GetSerialNumber() big.Int {
	if o == nil || isNil(o.SerialNumber) {
		var ret big.Int
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetSerialNumberOk() (*big.Int, bool) {
	if o == nil || isNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Certificate) HasSerialNumber() bool {
	if o != nil && !isNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given big.Int and assigns it to the SerialNumber field.
func (o *Certificate) SetSerialNumber(v big.Int) {
	o.SerialNumber = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value
func (o *Certificate) GetSignatureAlgorithm() EnumCertificateKeySignagureAlgorithm {
	if o == nil {
		var ret EnumCertificateKeySignagureAlgorithm
		return ret
	}

	return o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetSignatureAlgorithmOk() (*EnumCertificateKeySignagureAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureAlgorithm, true
}

// SetSignatureAlgorithm sets field value
func (o *Certificate) SetSignatureAlgorithm(v EnumCertificateKeySignagureAlgorithm) {
	o.SignatureAlgorithm = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *Certificate) GetStartsAt() time.Time {
	if o == nil || isNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *Certificate) HasStartsAt() bool {
	if o != nil && !isNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *Certificate) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Certificate) GetStatus() EnumCertificateKeyStatus {
	if o == nil || isNil(o.Status) {
		var ret EnumCertificateKeyStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetStatusOk() (*EnumCertificateKeyStatus, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Certificate) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EnumCertificateKeyStatus and assigns it to the Status field.
func (o *Certificate) SetStatus(v EnumCertificateKeyStatus) {
	o.Status = &v
}

// GetSubjectDN returns the SubjectDN field value
func (o *Certificate) GetSubjectDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetSubjectDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectDN, true
}

// SetSubjectDN sets field value
func (o *Certificate) SetSubjectDN(v string) {
	o.SubjectDN = v
}

// GetUsageType returns the UsageType field value
func (o *Certificate) GetUsageType() EnumCertificateKeyUsageType {
	if o == nil {
		var ret EnumCertificateKeyUsageType
		return ret
	}

	return o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetUsageTypeOk() (*EnumCertificateKeyUsageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageType, true
}

// SetUsageType sets field value
func (o *Certificate) SetUsageType(v EnumCertificateKeyUsageType) {
	o.UsageType = v
}

// GetValidityPeriod returns the ValidityPeriod field value
func (o *Certificate) GetValidityPeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetValidityPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidityPeriod, true
}

// SetValidityPeriod sets field value
func (o *Certificate) SetValidityPeriod(v int32) {
	o.ValidityPeriod = v
}

func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IssuerDN) {
		toSerialize["issuerDN"] = o.IssuerDN
	}
	if true {
		toSerialize["keyLength"] = o.KeyLength
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if true {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	if !isNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subjectDN"] = o.SubjectDN
	}
	if true {
		toSerialize["usageType"] = o.UsageType
	}
	if true {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	return json.Marshal(toSerialize)
}

type NullableCertificate struct {
	value *Certificate
	isSet bool
}

func (v NullableCertificate) Get() *Certificate {
	return v.value
}

func (v *NullableCertificate) Set(val *Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificate(val *Certificate) *NullableCertificate {
	return &NullableCertificate{value: val, isSet: true}
}

func (v NullableCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
