package policy

import (
	"encoding/json"
	"fmt"
)

// StatementResult struct
type StatementResult struct {
	Match       bool                   `json:"match"`
	Location    StatementLocation      `json:"location"`
	StatementID string                 `json:"statement_id"`
	Effect      Effect                 `json:"effect"`
	Resource    string                 `json:"resource,omitempty"`
	Action      string                 `json:"action,omitempty"`
	Condition   map[string]interface{} `json:"condition,omitempty"`
}

type StatementLocation int

const (
	ALL StatementLocation = iota
	RESOURCE
	ACTION
	CONDITION
)

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
	Resource    string                 `json:"resource,omitempty"`
	Action      []string               `json:"action,omitempty"`
	Condition   map[string]interface{} `json:"condition,omitempty"`
}

// NewStatement returns a new request
func NewStatement(id string, effect Effect, resource string, action []string, conditions map[string]interface{}) Statement {
	return Statement{
		StatementID: id,
		Effect:      effect,
		Resource:    resource,
		Action:      action,
		Condition:   conditions,
	}
}

// Validate this statement against the specified request
func (s *Statement) Validate(request *Request, registry *Registry) StatementResult {

	sr := StatementResult{}
	sr.StatementID = s.StatementID
	sr.Effect = s.Effect
	sr.Match = true
	sr.Location = StatementLocation(ALL)

	sr.Resource = request.Resource
	if !registry.GetResourceValidator().Validate(s.Resource, request.Resource) {
		sr.Match = false
		sr.Location = StatementLocation(RESOURCE)
		return sr
	}

	sr.Action = request.Action
	if !registry.GetActionValidator().Validate(s.Action, request.Action) {
		sr.Match = false
		sr.Location = StatementLocation(ACTION)
		return sr
	}

	sr.Condition = request.Condition
	for reqCondName, reqCond := range request.Condition {
		if cv := registry.GetConditionValidator(reqCondName); cv != nil {
			if c, ok := s.Condition[reqCondName]; ok {
				if !cv.Validate(c, reqCond) {
					sr.Match = false
					sr.Location = StatementLocation(CONDITION)
					return sr
				}
			}
		} else {
			sr.Match = false
			sr.Location = StatementLocation(CONDITION)
			return sr
		}
	}
	return sr
}

// IsAllow returns TRUE if this statement's effect = allow
func (s *Statement) IsAllow() bool {
	return s.Effect == Effect(ALLOW)
}

func (s *Statement) JSON() string {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(data)
}

func (s *Statement) PrettyJSON() string {
	data, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(data)
}
