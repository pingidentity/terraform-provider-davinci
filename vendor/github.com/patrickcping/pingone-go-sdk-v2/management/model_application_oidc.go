/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// ApplicationOIDC struct for ApplicationOIDC
type ApplicationOIDC struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	AccessControl *ApplicationAccessControl `json:"accessControl,omitempty"`
	// A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request.
	AssignActorRoles *bool `json:"assignActorRoles,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A string that specifies the description of the application.
	Description *string `json:"description,omitempty"`
	// A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Icon *ApplicationIcon `json:"icon,omitempty"`
	// A string that specifies the application ID.
	Id *string `json:"id,omitempty"`
	// A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains.
	LoginPageUrl *string `json:"loginPageUrl,omitempty"`
	// A string that specifies the name of the application. This is a required property.
	Name string `json:"name"`
	Protocol EnumApplicationProtocol `json:"protocol"`
	// An array that specifies the list of labels associated with the application. Options are PING_FED_CONNECTION_INTEGRATION.
	Tags []EnumApplicationTags `json:"tags,omitempty"`
	Type EnumApplicationType `json:"type"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A boolean that specifies whether the request query parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed.
	SupportUnsignedRequestObject *bool `json:"supportUnsignedRequestObject,omitempty"`
	Mobile *ApplicationOIDCAllOfMobile `json:"mobile,omitempty"`
	// A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.
	BundleId *string `json:"bundleId,omitempty"`
	// A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.
	PackageName *string `json:"packageName,omitempty"`
	Kerberos *ApplicationOIDCAllOfKerberos `json:"kerberos,omitempty"`
	// A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS.
	GrantTypes []EnumApplicationOIDCGrantType `json:"grantTypes"`
	// A string that specifies the custom home page URL for the application.
	HomePageUrl *string `json:"homePageUrl,omitempty"`
	PkceEnforcement *EnumApplicationOIDCPKCEOption `json:"pkceEnforcement,omitempty"`
	// A string that specifies the URLs that the browser can be redirected to after logout.
	PostLogoutRedirectUris []string `json:"postLogoutRedirectUris,omitempty"`
	// A string that specifies the callback URI for the authentication response.
	RedirectUris []string `json:"redirectUris,omitempty"`
	// An integer that specifies the lifetime in seconds of the refresh token. If a value is not provided, the default value is 2592000, or 30 days. Valid values are between 60 and 2147483647. If the refreshTokenRollingDuration property is specified for the application, then this property must be less than or equal to the value of refreshTokenRollingDuration. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token.
	RefreshTokenDuration *int32 `json:"refreshTokenDuration,omitempty"`
	// An integer that specifies the number of seconds a refresh token can be exchanged before re-authentication is required. If a value is not provided, the refresh token is valid forever. Valid values are between 60 and 2147483647. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token.
	RefreshTokenRollingDuration *int32 `json:"refreshTokenRollingDuration,omitempty"`
	// A string that specifies the code or token type returned by an authorization request. Options are TOKEN, ID_TOKEN, and CODE. Note that CODE cannot be used in an authorization request with TOKEN or ID_TOKEN because PingOne does not currently support OIDC hybrid flows.
	ResponseTypes []EnumApplicationOIDCResponseType `json:"responseTypes,omitempty"`
	TokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod `json:"tokenEndpointAuthMethod"`
}

// NewApplicationOIDC instantiates a new ApplicationOIDC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDC(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod) *ApplicationOIDC {
	this := ApplicationOIDC{}
	this.Enabled = enabled
	this.Name = name
	this.Protocol = protocol
	this.Type = type_
	this.GrantTypes = grantTypes
	this.TokenEndpointAuthMethod = tokenEndpointAuthMethod
	return &this
}

// NewApplicationOIDCWithDefaults instantiates a new ApplicationOIDC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCWithDefaults() *ApplicationOIDC {
	this := ApplicationOIDC{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *ApplicationOIDC) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetAccessControl() ApplicationAccessControl {
	if o == nil || o.AccessControl == nil {
		var ret ApplicationAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetAccessControlOk() (*ApplicationAccessControl, bool) {
	if o == nil || o.AccessControl == nil {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasAccessControl() bool {
	if o != nil && o.AccessControl != nil {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given ApplicationAccessControl and assigns it to the AccessControl field.
func (o *ApplicationOIDC) SetAccessControl(v ApplicationAccessControl) {
	o.AccessControl = &v
}

// GetAssignActorRoles returns the AssignActorRoles field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetAssignActorRoles() bool {
	if o == nil || o.AssignActorRoles == nil {
		var ret bool
		return ret
	}
	return *o.AssignActorRoles
}

// GetAssignActorRolesOk returns a tuple with the AssignActorRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetAssignActorRolesOk() (*bool, bool) {
	if o == nil || o.AssignActorRoles == nil {
		return nil, false
	}
	return o.AssignActorRoles, true
}

// HasAssignActorRoles returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasAssignActorRoles() bool {
	if o != nil && o.AssignActorRoles != nil {
		return true
	}

	return false
}

// SetAssignActorRoles gets a reference to the given bool and assigns it to the AssignActorRoles field.
func (o *ApplicationOIDC) SetAssignActorRoles(v bool) {
	o.AssignActorRoles = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApplicationOIDC) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationOIDC) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ApplicationOIDC) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApplicationOIDC) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ApplicationOIDC) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetIcon() ApplicationIcon {
	if o == nil || o.Icon == nil {
		var ret ApplicationIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetIconOk() (*ApplicationIcon, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ApplicationIcon and assigns it to the Icon field.
func (o *ApplicationOIDC) SetIcon(v ApplicationIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationOIDC) SetId(v string) {
	o.Id = &v
}

// GetLoginPageUrl returns the LoginPageUrl field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetLoginPageUrl() string {
	if o == nil || o.LoginPageUrl == nil {
		var ret string
		return ret
	}
	return *o.LoginPageUrl
}

// GetLoginPageUrlOk returns a tuple with the LoginPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetLoginPageUrlOk() (*string, bool) {
	if o == nil || o.LoginPageUrl == nil {
		return nil, false
	}
	return o.LoginPageUrl, true
}

// HasLoginPageUrl returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasLoginPageUrl() bool {
	if o != nil && o.LoginPageUrl != nil {
		return true
	}

	return false
}

// SetLoginPageUrl gets a reference to the given string and assigns it to the LoginPageUrl field.
func (o *ApplicationOIDC) SetLoginPageUrl(v string) {
	o.LoginPageUrl = &v
}

// GetName returns the Name field value
func (o *ApplicationOIDC) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationOIDC) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value
func (o *ApplicationOIDC) GetProtocol() EnumApplicationProtocol {
	if o == nil {
		var ret EnumApplicationProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetProtocolOk() (*EnumApplicationProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *ApplicationOIDC) SetProtocol(v EnumApplicationProtocol) {
	o.Protocol = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetTags() []EnumApplicationTags {
	if o == nil || o.Tags == nil {
		var ret []EnumApplicationTags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetTagsOk() ([]EnumApplicationTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []EnumApplicationTags and assigns it to the Tags field.
func (o *ApplicationOIDC) SetTags(v []EnumApplicationTags) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *ApplicationOIDC) GetType() EnumApplicationType {
	if o == nil {
		var ret EnumApplicationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetTypeOk() (*EnumApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationOIDC) SetType(v EnumApplicationType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApplicationOIDC) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetSupportUnsignedRequestObject() bool {
	if o == nil || o.SupportUnsignedRequestObject == nil {
		var ret bool
		return ret
	}
	return *o.SupportUnsignedRequestObject
}

// GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetSupportUnsignedRequestObjectOk() (*bool, bool) {
	if o == nil || o.SupportUnsignedRequestObject == nil {
		return nil, false
	}
	return o.SupportUnsignedRequestObject, true
}

// HasSupportUnsignedRequestObject returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasSupportUnsignedRequestObject() bool {
	if o != nil && o.SupportUnsignedRequestObject != nil {
		return true
	}

	return false
}

// SetSupportUnsignedRequestObject gets a reference to the given bool and assigns it to the SupportUnsignedRequestObject field.
func (o *ApplicationOIDC) SetSupportUnsignedRequestObject(v bool) {
	o.SupportUnsignedRequestObject = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetMobile() ApplicationOIDCAllOfMobile {
	if o == nil || o.Mobile == nil {
		var ret ApplicationOIDCAllOfMobile
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool) {
	if o == nil || o.Mobile == nil {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasMobile() bool {
	if o != nil && o.Mobile != nil {
		return true
	}

	return false
}

// SetMobile gets a reference to the given ApplicationOIDCAllOfMobile and assigns it to the Mobile field.
func (o *ApplicationOIDC) SetMobile(v ApplicationOIDCAllOfMobile) {
	o.Mobile = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetBundleId() string {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetBundleIdOk() (*string, bool) {
	if o == nil || o.BundleId == nil {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasBundleId() bool {
	if o != nil && o.BundleId != nil {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *ApplicationOIDC) SetBundleId(v string) {
	o.BundleId = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *ApplicationOIDC) SetPackageName(v string) {
	o.PackageName = &v
}

// GetKerberos returns the Kerberos field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetKerberos() ApplicationOIDCAllOfKerberos {
	if o == nil || o.Kerberos == nil {
		var ret ApplicationOIDCAllOfKerberos
		return ret
	}
	return *o.Kerberos
}

// GetKerberosOk returns a tuple with the Kerberos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetKerberosOk() (*ApplicationOIDCAllOfKerberos, bool) {
	if o == nil || o.Kerberos == nil {
		return nil, false
	}
	return o.Kerberos, true
}

// HasKerberos returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasKerberos() bool {
	if o != nil && o.Kerberos != nil {
		return true
	}

	return false
}

// SetKerberos gets a reference to the given ApplicationOIDCAllOfKerberos and assigns it to the Kerberos field.
func (o *ApplicationOIDC) SetKerberos(v ApplicationOIDCAllOfKerberos) {
	o.Kerberos = &v
}

// GetGrantTypes returns the GrantTypes field value
func (o *ApplicationOIDC) GetGrantTypes() []EnumApplicationOIDCGrantType {
	if o == nil {
		var ret []EnumApplicationOIDCGrantType
		return ret
	}

	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetGrantTypesOk() ([]EnumApplicationOIDCGrantType, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// SetGrantTypes sets field value
func (o *ApplicationOIDC) SetGrantTypes(v []EnumApplicationOIDCGrantType) {
	o.GrantTypes = v
}

// GetHomePageUrl returns the HomePageUrl field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetHomePageUrl() string {
	if o == nil || o.HomePageUrl == nil {
		var ret string
		return ret
	}
	return *o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetHomePageUrlOk() (*string, bool) {
	if o == nil || o.HomePageUrl == nil {
		return nil, false
	}
	return o.HomePageUrl, true
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasHomePageUrl() bool {
	if o != nil && o.HomePageUrl != nil {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given string and assigns it to the HomePageUrl field.
func (o *ApplicationOIDC) SetHomePageUrl(v string) {
	o.HomePageUrl = &v
}

// GetPkceEnforcement returns the PkceEnforcement field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetPkceEnforcement() EnumApplicationOIDCPKCEOption {
	if o == nil || o.PkceEnforcement == nil {
		var ret EnumApplicationOIDCPKCEOption
		return ret
	}
	return *o.PkceEnforcement
}

// GetPkceEnforcementOk returns a tuple with the PkceEnforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool) {
	if o == nil || o.PkceEnforcement == nil {
		return nil, false
	}
	return o.PkceEnforcement, true
}

// HasPkceEnforcement returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasPkceEnforcement() bool {
	if o != nil && o.PkceEnforcement != nil {
		return true
	}

	return false
}

// SetPkceEnforcement gets a reference to the given EnumApplicationOIDCPKCEOption and assigns it to the PkceEnforcement field.
func (o *ApplicationOIDC) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption) {
	o.PkceEnforcement = &v
}

// GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetPostLogoutRedirectUris() []string {
	if o == nil || o.PostLogoutRedirectUris == nil {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || o.PostLogoutRedirectUris == nil {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasPostLogoutRedirectUris() bool {
	if o != nil && o.PostLogoutRedirectUris != nil {
		return true
	}

	return false
}

// SetPostLogoutRedirectUris gets a reference to the given []string and assigns it to the PostLogoutRedirectUris field.
func (o *ApplicationOIDC) SetPostLogoutRedirectUris(v []string) {
	o.PostLogoutRedirectUris = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *ApplicationOIDC) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetRefreshTokenDuration() int32 {
	if o == nil || o.RefreshTokenDuration == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetRefreshTokenDurationOk() (*int32, bool) {
	if o == nil || o.RefreshTokenDuration == nil {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasRefreshTokenDuration() bool {
	if o != nil && o.RefreshTokenDuration != nil {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int32 and assigns it to the RefreshTokenDuration field.
func (o *ApplicationOIDC) SetRefreshTokenDuration(v int32) {
	o.RefreshTokenDuration = &v
}

// GetRefreshTokenRollingDuration returns the RefreshTokenRollingDuration field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetRefreshTokenRollingDuration() int32 {
	if o == nil || o.RefreshTokenRollingDuration == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokenRollingDuration
}

// GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetRefreshTokenRollingDurationOk() (*int32, bool) {
	if o == nil || o.RefreshTokenRollingDuration == nil {
		return nil, false
	}
	return o.RefreshTokenRollingDuration, true
}

// HasRefreshTokenRollingDuration returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasRefreshTokenRollingDuration() bool {
	if o != nil && o.RefreshTokenRollingDuration != nil {
		return true
	}

	return false
}

// SetRefreshTokenRollingDuration gets a reference to the given int32 and assigns it to the RefreshTokenRollingDuration field.
func (o *ApplicationOIDC) SetRefreshTokenRollingDuration(v int32) {
	o.RefreshTokenRollingDuration = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *ApplicationOIDC) GetResponseTypes() []EnumApplicationOIDCResponseType {
	if o == nil || o.ResponseTypes == nil {
		var ret []EnumApplicationOIDCResponseType
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetResponseTypesOk() ([]EnumApplicationOIDCResponseType, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *ApplicationOIDC) HasResponseTypes() bool {
	if o != nil && o.ResponseTypes != nil {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []EnumApplicationOIDCResponseType and assigns it to the ResponseTypes field.
func (o *ApplicationOIDC) SetResponseTypes(v []EnumApplicationOIDCResponseType) {
	o.ResponseTypes = v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value
func (o *ApplicationOIDC) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod {
	if o == nil {
		var ret EnumApplicationOIDCTokenAuthMethod
		return ret
	}

	return o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDC) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpointAuthMethod, true
}

// SetTokenEndpointAuthMethod sets field value
func (o *ApplicationOIDC) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod) {
	o.TokenEndpointAuthMethod = v
}

func (o ApplicationOIDC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.AccessControl != nil {
		toSerialize["accessControl"] = o.AccessControl
	}
	if o.AssignActorRoles != nil {
		toSerialize["assignActorRoles"] = o.AssignActorRoles
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LoginPageUrl != nil {
		toSerialize["loginPageUrl"] = o.LoginPageUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.SupportUnsignedRequestObject != nil {
		toSerialize["supportUnsignedRequestObject"] = o.SupportUnsignedRequestObject
	}
	if o.Mobile != nil {
		toSerialize["mobile"] = o.Mobile
	}
	if o.BundleId != nil {
		toSerialize["bundleId"] = o.BundleId
	}
	if o.PackageName != nil {
		toSerialize["packageName"] = o.PackageName
	}
	if o.Kerberos != nil {
		toSerialize["kerberos"] = o.Kerberos
	}
	if true {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if o.HomePageUrl != nil {
		toSerialize["homePageUrl"] = o.HomePageUrl
	}
	if o.PkceEnforcement != nil {
		toSerialize["pkceEnforcement"] = o.PkceEnforcement
	}
	if o.PostLogoutRedirectUris != nil {
		toSerialize["postLogoutRedirectUris"] = o.PostLogoutRedirectUris
	}
	if o.RedirectUris != nil {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if o.RefreshTokenDuration != nil {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if o.RefreshTokenRollingDuration != nil {
		toSerialize["refreshTokenRollingDuration"] = o.RefreshTokenRollingDuration
	}
	if o.ResponseTypes != nil {
		toSerialize["responseTypes"] = o.ResponseTypes
	}
	if true {
		toSerialize["tokenEndpointAuthMethod"] = o.TokenEndpointAuthMethod
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationOIDC struct {
	value *ApplicationOIDC
	isSet bool
}

func (v NullableApplicationOIDC) Get() *ApplicationOIDC {
	return v.value
}

func (v *NullableApplicationOIDC) Set(val *ApplicationOIDC) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOIDC) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOIDC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOIDC(val *ApplicationOIDC) *NullableApplicationOIDC {
	return &NullableApplicationOIDC{value: val, isSet: true}
}

func (v NullableApplicationOIDC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOIDC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


