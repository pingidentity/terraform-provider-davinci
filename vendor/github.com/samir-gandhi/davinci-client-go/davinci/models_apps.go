package davinci

type Apps struct {
	Apps []App `json:"apps,omitempty"`
}
type APIKeys struct {
	Prod string `json:"prod,omitempty"`
	Test string `json:"test,omitempty"`
}
type Metadata struct {
	RpName string `json:"rpName,omitempty"`
}
type UserPools struct {
	ConnectionID string `json:"connectionId,omitempty"`
	ConnectorID  string `json:"connectorId,omitempty"`
}
type Values struct {
	Enabled       bool          `json:"enabled,omitempty"`
	ClientSecret  string        `json:"clientSecret,omitempty"`
	RedirectUris  []string      `json:"redirectUris,omitempty"`
	LogoutUris    []interface{} `json:"logoutUris,omitempty"`
	AllowedScopes []string      `json:"allowedScopes,omitempty"`
	AllowedGrants []string      `json:"allowedGrants,omitempty"`
}
type ReadApp struct {
	App App `json:"app"`
}
type App struct {
	CompanyID     string        `json:"companyId,omitempty"`
	Name          string        `json:"name"`
	CustomerID    string        `json:"customerId,omitempty"`
	APIKeys       *APIKeys      `json:"apiKeys,omitempty"`
	Metadata      *Metadata     `json:"metadata,omitempty"`
	UserPools     []UserPools   `json:"userPools,omitempty"`
	Oauth         *Oauth        `json:"oauth,omitempty"`
	Saml          *Saml         `json:"saml,omitempty"`
	Flows         []interface{} `json:"flows,omitempty"`
	Policies      []Policy      `json:"policies,omitempty"`
	CreatedDate   int64         `json:"createdDate,omitempty"`
	APIKeyEnabled bool          `json:"apiKeyEnabled,omitempty"`
	AppID         string        `json:"appId,omitempty"`
	UserPortal    *UserPortal   `json:"userPortal,omitempty"`
}

type UserPortal struct {
	Values *UserPortalValues `json:"values"`
}
type UserPortalValues struct {
	UpTitle                 string `json:"upTitle"`
	AddAuthMethodTitle      string `json:"addAuthMethodTitle"`
	FlowTimeoutInSeconds    int    `json:"flowTimeoutInSeconds"`
	CredentialPageTitle     string `json:"credentialPageTitle"`
	CredentialPageSubTitle  string `json:"credentialPageSubTitle"`
	ShowUserInfo            bool   `json:"showUserInfo"`
	ShowMfaButton           bool   `json:"showMfaButton"`
	ShowVariables           bool   `json:"showVariables"`
	ShowLogoutButton        bool   `json:"showLogoutButton"`
	NameAuthMethodTitle     string `json:"nameAuthMethodTitle"`
	NameConfirmButtonText   string `json:"nameConfirmButtonText"`
	UpdateMessage           string `json:"updateMessage"`
	UpdateBodyMessage       string `json:"updateBodyMessage"`
	RemoveAuthMethodTitle   string `json:"removeAuthMethodTitle"`
	RemoveMessage           string `json:"removeMessage"`
	RemoveBodyMessage       string `json:"removeBodyMessage"`
	RemoveConfirmButtonText string `json:"removeConfirmButtonText"`
	RemoveCancelButtonText  string `json:"removeCancelButtonText"`
}

type Oauth struct {
	Enabled bool         `json:"enabled,omitempty"`
	Values  *OauthValues `json:"values,omitempty"`
}

type OauthValues struct {
	Enabled                    bool     `json:"enabled,omitempty"`
	ClientSecret               string   `json:"clientSecret,omitempty"`
	RedirectUris               []string `json:"redirectUris,omitempty"`
	LogoutUris                 []string `json:"logoutUris,omitempty"`
	AllowedScopes              []string `json:"allowedScopes,omitempty"`
	AllowedGrants              []string `json:"allowedGrants,omitempty"`
	EnforceSignedRequestOpenid bool     `json:"enforceSignedRequestOpenid,omitempty"`
	SpjwksUrl                  string   `json:"spjwksUrl,omitempty"`
	SpJwksOpenid               string   `json:"spJwksOpenid,omitempty"`
}

type Saml struct {
	Values *SamlValues `json:"values,omitempty"`
}

type SamlValues struct {
	Enabled              bool   `json:"enabled,omitempty"`
	RedirectURI          string `json:"redirectUri,omitempty"`
	Audience             string `json:"audience,omitempty"`
	EnforceSignedRequest bool   `json:"enforceSignedRequest,omitempty"`
	SpCert               string `json:"spCert,omitempty"`
}

type PolicyFlow struct {
	FlowID       string   `json:"flowId,omitempty"`
	VersionID    int      `json:"versionId,omitempty"`
	Weight       int      `json:"weight,omitempty"`
	SuccessNodes []string `json:"successNodes,omitempty"`
}
type Policy struct {
	PolicyFlows []PolicyFlow `json:"flows,omitempty"`
	Name        string       `json:"name,omitempty"`
	Status      string       `json:"status,omitempty"`
	PolicyID    string       `json:"policyId,omitempty"`
	CreatedDate int64        `json:"createdDate,omitempty"`
}

type AppUpdate struct {
	Name          string        `json:"name"`
	Oauth         *OauthUpdate  `json:"oauth,omitempty"`
	Saml          *SamlUpdate   `json:"saml,omitempty"`
	Flows         []interface{} `json:"flows,omitempty"`
	Policies      []Policy      `json:"policies,omitempty"`
	APIKeyEnabled bool          `json:"apiKeyEnabled,omitempty"`
	AppID         string        `json:"appId,omitempty"`
	UserPortal    *UserPortal   `json:"userPortal,omitempty"`
}

type OauthUpdate struct {
	Enabled bool               `json:"enabled,omitempty"`
	Values  *OauthValuesUpdate `json:"values,omitempty"`
}

type OauthValuesUpdate struct {
	Enabled                    bool     `json:"enabled,omitempty"`
	RedirectUris               []string `json:"redirectUris,omitempty"`
	LogoutUris                 []string `json:"logoutUris,omitempty"`
	AllowedScopes              []string `json:"allowedScopes,omitempty"`
	AllowedGrants              []string `json:"allowedGrants,omitempty"`
	EnforceSignedRequestOpenid bool     `json:"enforceSignedRequestOpenid,omitempty"`
	SpjwksUrl                  string   `json:"spjwksUrl,omitempty"`
	SpJwksOpenid               string   `json:"spJwksOpenid,omitempty"`
}

type SamlUpdate struct {
	Values *SamlValuesUpdate `json:"values,omitempty"`
}

type SamlValuesUpdate struct {
	Enabled              bool   `json:"enabled,omitempty"`
	RedirectURI          string `json:"redirectUri,omitempty"`
	Audience             string `json:"audience,omitempty"`
	EnforceSignedRequest bool   `json:"enforceSignedRequest,omitempty"`
	SpCert               string `json:"spCert,omitempty"`
}
