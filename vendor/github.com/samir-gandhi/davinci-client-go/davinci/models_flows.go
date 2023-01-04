package davinci

//	type ShowContinueButton struct {
//		Value bool `json:"value,omitempty"`
//	}
type SubFlowValue struct {
	Label string `json:"label,omitempty" mapstructure:"label"`
	Value string `json:"value,omitempty" mapstructure:"value"`
}
type SubFlowID struct {
	Value SubFlowValue `json:"value,omitempty" mapstructure:"value"`
}
type SubFlowVersionID struct {
	Value string `json:"value,omitempty" mapstructure:"value"`
}

// Used for type assertion on Properties of a Node Data
type SubFlowProperties struct {
	SubFlowID        SubFlowID        `json:"subFlowId,omitempty" mapstructure:"subFlowId"`
	SubFlowVersionID SubFlowVersionID `json:"subFlowVersionId,omitempty" mapstructure:"subFlowVersionId"`
}

type Data struct {
	ID                 string     `json:"id,omitempty"`
	NodeType           string     `json:"nodeType,omitempty"`
	ConnectionID       string     `json:"connectionId,omitempty"`
	ConnectorID        string     `json:"connectorId,omitempty"`
	Name               string     `json:"name,omitempty"`
	Label              string     `json:"label,omitempty"`
	Status             string     `json:"status,omitempty"`
	CapabilityName     string     `json:"capabilityName,omitempty"`
	Type               string     `json:"type,omitempty"`
	Properties         Properties `json:"properties,omitempty"`
	Source             string     `json:"source,omitempty"`
	Target             string     `json:"target,omitempty"`
	MultiValueSourceId string     `json:"multiValueSourceId,omitempty"`
}

type NodeData struct {
	ID             string     `json:"id,omitempty"`
	NodeType       string     `json:"nodeType,omitempty"`
	ConnectionID   string     `json:"connectionId,omitempty"`
	ConnectorID    string     `json:"connectorId,omitempty"`
	Name           string     `json:"name,omitempty"`
	Label          string     `json:"label,omitempty"`
	Status         string     `json:"status,omitempty"`
	CapabilityName string     `json:"capabilityName,omitempty"`
	Type           string     `json:"type,omitempty"`
	Properties     Properties `json:"properties,omitempty"`
	Source         string     `json:"source,omitempty"`
	Target         string     `json:"target,omitempty"`
}

type EdgeData struct {
	ID                 string     `json:"id,omitempty"`
	NodeType           string     `json:"nodeType,omitempty"`
	ConnectionID       string     `json:"connectionId,omitempty"`
	ConnectorID        string     `json:"connectorId,omitempty"`
	Name               string     `json:"name,omitempty"`
	Label              string     `json:"label,omitempty"`
	Status             string     `json:"status,omitempty"`
	CapabilityName     string     `json:"capabilityName,omitempty"`
	Type               string     `json:"type,omitempty"`
	Properties         Properties `json:"properties,omitempty"`
	Source             string     `json:"source,omitempty"`
	Target             string     `json:"target,omitempty"`
	MultiValueSourceId string     `json:"multiValueSourceId,omitempty"`
}

type Position struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type Nodes struct {
	Data       NodeData `json:"data,omitempty"`
	Position   Position `json:"position,omitempty"`
	Group      string   `json:"group"`
	Removed    bool     `json:"removed"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbable  bool     `json:"grabbable"`
	Pannable   bool     `json:"pannable"`
	Classes    string   `json:"classes"`
}

type Edges struct {
	Data       EdgeData `json:"data,omitempty"`
	Position   Position `json:"position,omitempty"`
	Group      string   `json:"group"`
	Removed    bool     `json:"removed"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbable  bool     `json:"grabbable"`
	Pannable   bool     `json:"pannable"`
	Classes    string   `json:"classes"`
}

type Elements struct {
	Nodes []Nodes `json:"nodes,omitempty"`
	Edges []Edges `json:"edges,omitempty"`
}

type Pan struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type Renderer struct {
	Name string `json:"name,omitempty"`
}

type GraphData struct {
	Elements            Elements `json:"elements,omitempty"`
	Data                Data     `json:"data,omitempty"`
	ZoomingEnabled      bool     `json:"zoomingEnabled,omitempty"`
	UserZoomingEnabled  bool     `json:"userZoomingEnabled,omitempty"`
	Zoom                int      `json:"zoom,omitempty"`
	MinZoom             float64  `json:"minZoom,omitempty"`
	MaxZoom             float64  `json:"maxZoom,omitempty"`
	PanningEnabled      bool     `json:"panningEnabled,omitempty"`
	UserPanningEnabled  bool     `json:"userPanningEnabled,omitempty"`
	Pan                 Pan      `json:"pan,omitempty"`
	BoxSelectionEnabled bool     `json:"boxSelectionEnabled,omitempty"`
	Renderer            Renderer `json:"renderer,omitempty"`
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
