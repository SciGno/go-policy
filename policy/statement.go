package policy

// Effect type
type Effect string

const (
	// ALLOW effect for policy statement
	ALLOW Effect = "Allow"
	// DENY effect for policy statement
	DENY Effect = "Deny"
)

// Statement strct
type Statement struct {
	StatementID string                 `json:"sid,omitempty"`
	Effect      Effect                 `json:"effect,omitempty"`
	Action      []string               `json:"action,omitempty"`
	Resource    string                 `json:"resource,omitempty"`
	Condition   map[string]interface{} `json:"condition,omitempty"`
}

// NewStatement returns a new request
func NewStatement(id string, effect Effect, action []string, resource string, conditions map[string]interface{}) Statement {
	return Statement{
		StatementID: id,
		Effect:      effect,
		Action:      action,
		Resource:    resource,
		Condition:   conditions,
	}
}

// Validate this statement against the specified request
func (s *Statement) Validate(request *Request, registry *Registry) bool {

	if !registry.GetResourceValidator().Validate(s.Resource, request.Resource) {
		return false
	}

	if !registry.GetActionValidator().Validate(s.Action, request.Action) {
		return false
	}

	for rn, rc := range request.Condition {
		if cv := registry.GetConditionnValidator(rn); cv != nil {
			if c, ok := s.Condition[rn]; ok {
				if !cv.Validate(c, rc) {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}

// IsAllow returns TRUE if this statement's effect = allow
func (s *Statement) IsAllow() bool {
	return s.Effect == ALLOW
}
