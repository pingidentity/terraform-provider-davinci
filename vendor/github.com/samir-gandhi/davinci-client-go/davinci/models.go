package davinci

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type ClientInput struct {
	HostURL         string
	PingOneRegion   string
	Username        string
	Password        string
	PingOneSSOEnvId string
}

type Client struct {
	HostURL     string
	HTTPClient  *http.Client
	Token       string
	Auth        AuthStruct
	CompanyID   string
	AuthP1SSO   AuthP1SSO
	AuthRefresh bool
}
type APIClient struct {
	HostURL         string
	PingOneRegion   string
	HTTPClient      *http.Client
	Token           string
	Auth            AuthStruct
	CompanyID       string
	PingOneSSOEnvId string
	AuthRefresh     bool
}

type Params struct {
	Page        string
	Limit       string
	ExtraParams map[string]string
	// TODO: figure out what query is
	// query  string
}

type DvHttpRequest struct {
	Method string
	Url    string
	Body   io.Reader
}

type DvHttpResponse struct {
	Body           []byte
	Headers        http.Header
	StatusCode     int
	Location       *url.URL
	LocationParams url.Values
}

type DvHttpError struct {
	Status int
	Body   string
}

type AuthStruct struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type AuthP1SSO struct {
	PingOneAdminEnvId  string `json:"envId,omitempty"`
	PingOneTargetEnvId string `json:"targetEnvId,omitempty"`
}

type SkSdkToken struct {
	CapabilityName string `json:"capabilityName"`
	AccessToken    string `json:"access_token"`
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	Success        bool   `json:"success"`
}
type LoginResponse struct {
	AccessToken     string     `json:"access_token"`
	TokenType       string     `json:"token_type"`
	MfaRequired     bool       `json:"mfaRequired"`
	Status          string     `json:"status"`
	CustomerID      string     `json:"customerId"`
	AppConfig       bool       `json:"appConfig"`
	SkSdkToken      SkSdkToken `json:"skSdkToken"`
	FlowPolicyID    string     `json:"flowPolicyId"`
	CompanyID       string     `json:"companyId"`
	SelectedCompany string     `json:"selectedCompany"`
}

type Callback struct {
	InteractionID    string `json:"interactionId"`
	CompanyID        string `json:"companyId"`
	ConnectionID     string `json:"connectionId"`
	ConnectorID      string `json:"connectorId"`
	ID               string `json:"id"`
	CapabilityName   string `json:"capabilityName"`
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	IDToken          string `json:"id_token"`
	Success          bool   `json:"success"`
	InteractionToken string `json:"interactionToken"`
}

type AuthResponse struct {
	AccessToken     string `json:"access_token"`
	TokenType       string `json:"token_type"`
	Status          string `json:"status"`
	CustomerID      string `json:"customerId"`
	TryFlowToken    string `json:"tryFlowToken"`
	SelectedCompany string `json:"selectedCompany"`
}

type Companies struct {
	CompanyID string   `json:"companyId"`
	Roles     []string `json:"roles"`
	Name      string   `json:"name"`
	SvgIcon   string   `json:"svgIcon"`
}

type Environments struct {
	CustomerID  string      `json:"customerId"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Email       string      `json:"email"`
	PhoneNumber string      `json:"phoneNumber"`
	CompanyID   string      `json:"companyId"`
	Companies   []Companies `json:"companies"`
	ClientID    string      `json:"clientId,omitempty"`
	CreatedDate int64       `json:"createdDate"`
}

type Environment struct {
	CreatedByCustomerID string `json:"createdByCustomerId"`
	CreatedByCompanyID  string `json:"createdByCompanyId"`
	Name                string `json:"name"`
	CompanyType         string `json:"companyType"`
	EntitlementTemplate string `json:"entitlementTemplate"`
	EntitlementProps    struct {
	} `json:"entitlementProps"`
	SecurityType string `json:"securityType"`
	JwtKeys      struct {
		Jwks struct {
			Keys []struct {
				Kty string `json:"kty"`
				Kid string `json:"kid"`
				N   string `json:"n"`
				E   string `json:"e"`
				Alg string `json:"alg"`
				Use string `json:"use"`
			} `json:"keys"`
		} `json:"jwks"`
	} `json:"jwtKeys"`
	SamlKeys struct {
		PublicKey string `json:"publicKey"`
		Cert      string `json:"cert"`
	} `json:"samlKeys"`
	Properties struct {
		PngIcon struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"pngIcon"`
		SvgIcon struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"svgIcon"`
		SvgViewBox struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"svgViewBox"`
		IconOpacity struct {
			DisplayName          string  `json:"displayName"`
			DataType             string  `json:"dataType"`
			PreferredControlType string  `json:"preferredControlType"`
			Value                float64 `json:"value"`
		} `json:"iconOpacity"`
		BackgroundColor struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"backgroundColor"`
		TextColor struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"textColor"`
		IconColor struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"iconColor"`
		ArcColor struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"arcColor"`
		ArcProgressColor struct {
			DisplayName          string `json:"displayName"`
			DataType             string `json:"dataType"`
			PreferredControlType string `json:"preferredControlType"`
			Value                string `json:"value"`
		} `json:"arcProgressColor"`
	} `json:"properties"`
	CreatedDate int64 `json:"createdDate"`
	Entitlement struct {
		Company struct {
			CreateAdditional bool `json:"createAdditional"`
		} `json:"company"`
		Connectors struct {
			Whitelist []string `json:"whitelist"`
			//TODO
			// Blacklist []string `json:"blacklist"`
			Blacklist interface{} `json:"blacklist"`
		} `json:"connectors"`
		Connections struct {
			Total                              int `json:"total"`
			MaxNumberOfConnectionsPerConnector struct {
				ConnectorID string `json:"connectorId"`
				Total       int    `json:"total"`
			} `json:"maxNumberOfConnectionsPerConnector"`
		} `json:"connections"`
		Flows struct {
			Enabled bool  `json:"enabled"`
			Total   int   `json:"total"`
			Expires int64 `json:"expires"`
		} `json:"flows"`
		Attributes struct {
			Enabled bool `json:"enabled"`
			Total   int  `json:"total"`
		} `json:"attributes"`
		Apps struct {
			Total int `json:"total"`
		} `json:"apps"`
		Users struct {
			Total                   int `json:"total"`
			TotalCredentialsPerUser int `json:"totalCredentialsPerUser"`
		} `json:"users"`
		Expires int64 `json:"expires"`
	} `json:"entitlement"`
	CompanyID string `json:"companyId"`
}

type EnvironmentStats struct {
	TableStats []struct {
		Flows       int `json:"Flows"`
		Connections int `json:"Connections"`
		Apps        int `json:"Apps"`
		Customers   int `json:"Customers"`
		Constructs  int `json:"Constructs"`
		Users       int `json:"Users"`
		Events      int `json:"Events"`
		ID          struct {
			CompanyID string `json:"companyId"`
			Ts        int64  `json:"ts"`
		} `json:"_id"`
	} `json:"tableStats"`
	PopularFlows []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Name     string `json:"name,omitempty"`
	} `json:"popularFlows"`
	RunningFlowsCount []struct {
		KeyAsString time.Time `json:"key_as_string"`
		Key         int64     `json:"key"`
		DocCount    int       `json:"doc_count"`
	} `json:"runningFlowsCount"`
	EventOutcomesCount []interface{} `json:"eventOutcomesCount"`
	AllFlows           []string      `json:"allFlows"`
}

type Message struct {
	Message string `json:"message,omitempty"`
	Value   string `json:"value,omitempty"`
}

type Customer struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Companies []struct {
		CompanyID string   `json:"companyId"`
		Roles     []string `json:"roles"`
	} `json:"companies"`
	CustomerType        string `json:"customerType"`
	CreatedByCustomerID string `json:"createdByCustomerId"`
	CreatedByCompanyID  string `json:"createdByCompanyId"`
	CompanyID           string `json:"companyId"`
	EmailVerified       bool   `json:"emailVerified"`
	CreatedDate         int64  `json:"createdDate"`
	LastLogin           int64  `json:"lastLogin"`
	SkUserID            string `json:"skUserId,omitempty"`
	CustomerID          string `json:"customerId"`
	ClientID            string `json:"clientId,omitempty"`
	PhoneNumber         string `json:"phoneNumber,omitempty"`
	Status              string `json:"status,omitempty"`
	EmailVerifiedDate   int64  `json:"emailVerifiedDate,omitempty"`
}

type Customers struct {
	Customers     []Customer `json:"customers"`
	CustomerCount int        `json:"customerCount"`
}

type CustomerUpdate struct {
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Roles       []string `json:"roles"`
	PhoneNumber string   `json:"phoneNumber"`
}

type CustomerCreate struct {
	Email       string   `json:"email"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Roles       []string `json:"roles"`
	PhoneNumber string   `json:"phoneNumber"`
}

type CreatedCustomer struct {
	Email               string      `json:"email"`
	CompanyID           string      `json:"companyId"`
	ClientID            interface{} `json:"clientId"`
	FirstName           string      `json:"firstName"`
	LastName            string      `json:"lastName"`
	PhoneNumber         interface{} `json:"phoneNumber"`
	CreatedByCustomerID string      `json:"createdByCustomerId"`
	CreatedByCompanyID  string      `json:"createdByCompanyId"`
	EmailVerified       bool        `json:"emailVerified"`
	Companies           []Companies `json:"companies"`
	Salt                string      `json:"salt"`
	HashedPassword      string      `json:"hashedPassword"`
	Status              string      `json:"status"`
	CustomerType        string      `json:"customerType"`
	CreatedDate         int64       `json:"createdDate"`
	EmailVerifiedDate   int64       `json:"emailVerifiedDate"`
	PasswordHistory     []struct {
		HashedPassword string `json:"hashedPassword"`
		Salt           string `json:"salt"`
		AddedDate      int64  `json:"addedDate"`
	} `json:"passwordHistory"`
	SkUserID    string `json:"skUserId"`
	LastLogin   int64  `json:"lastLogin"`
	FailedLogin struct {
		RetryCount           int   `json:"retryCount"`
		FirstFailedTimestamp int64 `json:"firstFailedTimestamp"`
	} `json:"failedLogin"`
	CustomerID string `json:"customerId"`
}

// TODO: Cleanup roles
type Role struct {
	ID struct {
		Name      string `json:"name,omitempty"`
		CompanyID string `json:"companyId,omitempty"`
	} `json:"_id,omitempty"`
	CreatedDate int64  `json:"createdDate,omitempty"`
	Description string `json:"description,omitempty"`
	Policy      []struct {
		Resource string `json:"resource,omitempty"`
		Actions  []struct {
			Action string `json:"action,omitempty"`
			Allow  bool   `json:"allow,omitempty"`
		} `json:"actions,omitempty"`
	} `json:"policy,omitempty"`
}

//TODO Wish this worked
// type Roles []struct{
// 	RolesCreate
// 	RolesCreateResponse
// }

type RoleCreate struct {
	Name string `json:"name"`
}

type RoleCreateResponse struct {
	ID struct {
		Name      string `json:"name,omitempty"`
		CompanyID string `json:"companyId,omitempty"`
	} `json:"_id,omitempty"`
	CreatedDate int64 `json:"createdDate,omitempty"`
}

type RoleUpdate struct {
	Description string `json:"description,omitempty"`
	Policy      []struct {
		Resource string `json:"resource,omitempty"`
		Actions  []struct {
			Action string `json:"action,omitempty"`
			Allow  bool   `json:"allow,omitempty"`
		} `json:"actions,omitempty"`
	} `json:"policy,omitempty"`
}

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

type FlowInfo struct {
	Flow Flow `json:"flowInfo,omitempty"`
}
type FlowsInfo struct {
	Flow []Flow `json:"flowsInfo,omitempty"`
}

type FlowImport struct {
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	FlowInfo        Flow              `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}
type FlowsImport struct {
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	FlowInfo        Flows             `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}

type Flows struct {
	Flow []Flow `json:"flows,omitempty"`
}

type Flow struct {
	CustomerID         string        `json:"customerId,omitempty"`
	FlowStatus         string        `json:"flowStatus,omitempty"`
	CurrentVersion     int           `json:"currentVersion,omitempty"`
	PublishedVersion   int           `json:"publishedVersion,omitempty"`
	Name               string        `json:"name,omitempty"`
	Description        string        `json:"description,omitempty"`
	CreatedDate        int64         `json:"createdDate,omitempty"`
	UpdatedDate        int64         `json:"updatedDate,omitempty"`
	AuthTokenExpireIds []interface{} `json:"authTokenExpireIds,omitempty"`
	DeployedDate       int64         `json:"deployedDate,omitempty"`
	// Edited, removed struct. Staying one level only
	EnabledGraphData     interface{} `json:"enabledGraphData,omitempty"`
	FunctionConnectionID interface{} `json:"functionConnectionId,omitempty"`
	// edited
	InputSchemaCompiled interface{}   `json:"inputSchemaCompiled,omitempty"`
	IsInputSchemaSaved  bool          `json:"isInputSchemaSaved,omitempty"`
	IsOutputSchemaSaved bool          `json:"isOutputSchemaSaved,omitempty"`
	Orx                 string        `json:"orx,omitempty"`
	Settings            interface{}   `json:"settings,omitempty"`
	Timeouts            interface{}   `json:"timeouts,omitempty"`
	FlowID              string        `json:"flowId,omitempty"`
	CompanyID           string        `json:"companyId,omitempty"`
	GraphData           GraphData     `json:"graphData,omitempty"`
	InputSchema         []interface{} `json:"inputSchema,omitempty"`
	OutputSchema        interface{}   `json:"outputSchema,omitempty"`
	FlowColor           string        `json:"flowColor,omitempty"`
	ConnectorIds        []string      `json:"connectorIds,omitempty"`
	SavedDate           int64         `json:"savedDate,omitempty"`
}

// type GraphData struct {
// 	Elements            interface{} `json:"elements"`
// 	Data                interface{} `json:"data"`
// 	ZoomingEnabled      bool        `json:"zoomingEnabled"`
// 	UserZoomingEnabled  bool        `json:"userZoomingEnabled"`
// 	Zoom                int         `json:"zoom"`
// 	MinZoom             float64     `json:"minZoom"`
// 	MaxZoom             float64     `json:"maxZoom"`
// 	PanningEnabled      bool        `json:"panningEnabled"`
// 	UserPanningEnabled  bool        `json:"userPanningEnabled"`
// 	Pan                 interface{} `json:"pan"`
// 	BoxSelectionEnabled bool        `json:"boxSelectionEnabled"`
// 	Renderer            interface{} `json:"renderer"`
// }

type Variable struct {
	Context     string `json:"context,omitempty"`
	CreatedDate int64  `json:"createdDate,omitempty"`
	CustomerID  string `json:"customerId,omitempty"`
	Type        string `json:"type,omitempty"`
	Visibility  string `json:"visibility,omitempty"`
	CompanyID   string `json:"companyId,omitempty"`
	TotalCount  int    `json:"totalCount,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Value       string `json:"value,omitempty"`
	Mutable     bool   `json:"mutable,omitempty"`
	Min         int    `json:"min,omitempty"`
	Max         int    `json:"max,omitempty"`
}

type VariablePayload struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"displayName,omitempty"`
	Context     string `json:"context,omitempty" validate:"oneof=company flowInstance user"`
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	Mutable     bool   `json:"mutable,omitempty"`
	Min         int    `json:"min,omitempty"`
	Max         int    `json:"max,omitempty"`
}
