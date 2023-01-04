package davinci

// Representation of an instantiated connector
// All fields are set with omitempty as nil values aren't allowed
// to various requests requiring specific fields.
type Connection struct {
	CustomerID   string     `json:"customerId,omitempty"`
	ConnectorID  string     `json:"connectorId,omitempty"`
	Name         string     `json:"name,omitempty"`
	CreatedDate  int64      `json:"createdDate,omitempty"`
	Properties   Properties `json:"properties,omitempty"`
	UpdatedDate  int64      `json:"updatedDate,omitempty"`
	ConnectionID string     `json:"connectionId,omitempty"`
	CompanyID    string     `json:"companyId,omitempty"`
}

type Properties map[string]interface{}

// CustomAuth is a field of the Properties struct
// but is kept separate because the structure is known.
// This is used to unmarshal the properties["customAuth"] field of a connection response.
type CustomAuth struct {
	Properties CustomAuthProperties `json:"properties,omitempty"`
}

type CustomAuthProperties struct {
	ProviderName                  ProviderName                  `json:"providerName,omitempty"`
	AuthTypeDropdown              AuthTypeDropdown              `json:"authTypeDropdown,omitempty"`
	SkRedirectURI                 SkRedirectURI                 `json:"skRedirectUri,omitempty"`
	IssuerURL                     IssuerURL                     `json:"issuerUrl,omitempty"`
	AuthorizationEndpoint         AuthorizationEndpoint         `json:"authorizationEndpoint,omitempty"`
	TokenEndpoint                 TokenEndpoint                 `json:"tokenEndpoint,omitempty"`
	BearerToken                   BearerToken                   `json:"bearerToken,omitempty"`
	UserInfoEndpoint              UserInfoEndpoint              `json:"userInfoEndpoint,omitempty"`
	ClientID                      ClientID                      `json:"clientId,omitempty"`
	ClientSecret                  ClientSecret                  `json:"clientSecret,omitempty"`
	Scope                         Scope                         `json:"scope,omitempty"`
	Code                          Code                          `json:"code,omitempty"`
	UserConnectorAttributeMapping UserConnectorAttributeMapping `json:"userConnectorAttributeMapping,omitempty"`
	CustomAttributes              CustomAttributes              `json:"customAttributes,omitempty"`
	ReturnToURL                   ReturnToURL                   `json:"returnToUrl,omitempty"`
}

type ProviderName struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Required             bool   `json:"required,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
	Value                string `json:"value,omitempty"`
}
type Options struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type AuthTypeDropdown struct {
	DisplayName          string    `json:"displayName,omitempty"`
	PreferredControlType string    `json:"preferredControlType,omitempty"`
	Required             bool      `json:"required,omitempty"`
	Options              []Options `json:"options,omitempty"`
	Enum                 []string  `json:"enum,omitempty"`
	Value                string    `json:"value,omitempty"`
	Placeholder          string    `json:"placeholder,omitempty"`
}
type SkRedirectURI struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Disabled             bool   `json:"disabled,omitempty"`
	InitializeValue      string `json:"initializeValue,omitempty"`
	CopyToClip           bool   `json:"copyToClip,omitempty"`
}
type IssuerURL struct {
	PreferredControlType string `json:"preferredControlType,omitempty"`
	DisplayName          string `json:"displayName,omitempty"`
	Info                 string `json:"info,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type AuthorizationEndpoint struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Required             bool   `json:"required,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type TokenEndpoint struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Required             bool   `json:"required,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type BearerToken struct {
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Type                 string `json:"type,omitempty"`
	DisplayName          string `json:"displayName,omitempty"`
	Info                 string `json:"info,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type UserInfoEndpoint struct {
	DisplayName          string   `json:"displayName,omitempty"`
	PreferredControlType string   `json:"preferredControlType,omitempty"`
	Required             bool     `json:"required,omitempty"`
	Value                []string `json:"value,omitempty"`
	Placeholder          []string `json:"placeholder,omitempty"`
}
type ClientID struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Required             bool   `json:"required,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type ClientSecret struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	HashedVisibility     bool   `json:"hashedVisibility,omitempty"`
	Required             bool   `json:"required,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type Scope struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Required             bool   `json:"required,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type Code struct {
	DisplayName          string `json:"displayName,omitempty"`
	Info                 string `json:"info,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Language             string `json:"language,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
type Username struct {
	Value1 string `json:"value1,omitempty"`
}
type FirstName struct {
	Value1 string `json:"value1,omitempty"`
}
type LastName struct {
	Value1 string `json:"value1,omitempty"`
}
type Name struct {
	Value1 string `json:"value1,omitempty"`
}
type Email struct {
	Value1 string `json:"value1,omitempty"`
}
type Mapping struct {
	Username  Username  `json:"username,omitempty"`
	FirstName FirstName `json:"firstName,omitempty"`
	LastName  LastName  `json:"lastName,omitempty"`
	Name      Name      `json:"name,omitempty"`
	Email     Email     `json:"email,omitempty"`
}
type UserConnectorAttributeMappingValue struct {
	UserPoolConnectionID string  `json:"userPoolConnectionId,omitempty"`
	Mapping              Mapping `json:"mapping,omitempty"`
}
type UserConnectorAttributeMappingPlaceholder struct {
	UserPoolConnectionID string  `json:"userPoolConnectionId,omitempty"`
	Mapping              Mapping `json:"mapping,omitempty"`
}
type UserConnectorAttributeMapping struct {
	Type                 string                                   `json:"type,omitempty"`
	DisplayName          interface{}                              `json:"displayName,omitempty"`
	PreferredControlType string                                   `json:"preferredControlType,omitempty"`
	NewMappingAllowed    bool                                     `json:"newMappingAllowed,omitempty"`
	Title1               interface{}                              `json:"title1,omitempty"`
	Title2               interface{}                              `json:"title2,omitempty"`
	Sections             []string                                 `json:"sections,omitempty"`
	Value                UserConnectorAttributeMappingValue       `json:"value,omitempty"`
	Placeholder          UserConnectorAttributeMappingPlaceholder `json:"placeholder,omitempty"`
}
type CustomAttributesValue struct {
	Name          string      `json:"name,omitempty"`
	Description   string      `json:"description,omitempty"`
	Type          string      `json:"type,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	MinLength     string      `json:"minLength,omitempty"`
	MaxLength     string      `json:"maxLength,omitempty"`
	Required      bool        `json:"required,omitempty"`
	AttributeType string      `json:"attributeType,omitempty"`
}
type Placeholder struct {
	Name          string      `json:"name,omitempty"`
	Description   string      `json:"description,omitempty"`
	Type          string      `json:"type,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	MinLength     string      `json:"minLength,omitempty"`
	MaxLength     string      `json:"maxLength,omitempty"`
	Required      bool        `json:"required,omitempty"`
	AttributeType string      `json:"attributeType,omitempty"`
}
type CustomAttributes struct {
	Type                 string                  `json:"type,omitempty"`
	DisplayName          string                  `json:"displayName,omitempty"`
	PreferredControlType string                  `json:"preferredControlType,omitempty"`
	Info                 string                  `json:"info,omitempty"`
	Sections             []string                `json:"sections,omitempty"`
	Value                []CustomAttributesValue `json:"value,omitempty"`
	Placeholder          []Placeholder           `json:"placeholder,omitempty"`
}
type ReturnToURL struct {
	DisplayName          string `json:"displayName,omitempty"`
	PreferredControlType string `json:"preferredControlType,omitempty"`
	Info                 string `json:"info,omitempty"`
	Value                string `json:"value,omitempty"`
	Placeholder          string `json:"placeholder,omitempty"`
}
