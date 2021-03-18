package policy

// Request struct
type Request struct {
	Action    string                 `json:"action,omitempty"`
	Resource  string                 `json:"resource,omitempty"`
	Condition map[string]interface{} `json:"condition,omitempty"`
}

// NewRequest returns a new request
func NewRequest(action, resource string, conditions map[string]interface{}) Request {
	return Request{
		Action:    action,
		Resource:  resource,
		Condition: conditions,
	}
}
