package policy

// Request struct
type Request struct {
	Resource  string                 `json:"resource,omitempty"`
	Action    string                 `json:"action,omitempty"`
	Condition map[string]interface{} `json:"condition,omitempty"`
}

// NewRequest returns a new request
func NewRequest(resource, action string, conditions map[string]interface{}) Request {
	return Request{
		Action:    action,
		Resource:  resource,
		Condition: conditions,
	}
}
