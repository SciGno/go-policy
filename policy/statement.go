package policy

import (
	"encoding/json"
)

// StatementResult struct
type StatementResult struct {
	Match       bool                   `json:"match"`
	StatementID string                 `json:"statement_id"`
	Effect      Effect                 `json:"effect"`
	Resource    string                 `json:"resource,omitempty"`
	Action      string                 `json:"action,omitempty"`
	Condition   map[string]interface{} `json:"condition,omitempty"`
}

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

	statementResutl := StatementResult{}
	statementResutl.StatementID = s.StatementID
	statementResutl.Effect = s.Effect
	statementResutl.Match = true

	statementResutl.Resource = request.Resource
	if !registry.GetResourceValidator().Validate(s.Resource, request.Resource) {
		statementResutl.Match = false
		return statementResutl
	}

	statementResutl.Action = request.Action
	if !registry.GetActionValidator().Validate(s.Action, request.Action) {
		statementResutl.Match = false
		return statementResutl
	}

	statementResutl.Condition = request.Condition
	for reqCondName, reqCond := range request.Condition {
		if cv := registry.GetConditionValidator(reqCondName); cv != nil {
			if c, ok := s.Condition[reqCondName]; ok {
				if !cv.Validate(c, reqCond) {
					statementResutl.Match = false
					return statementResutl
				}
			}
		} else {
			statementResutl.Match = false
			return statementResutl
		}
	}

	return statementResutl
}

// IsAllow returns TRUE if this statement's effect = allow
func (s *Statement) IsAllow() bool {
	return s.Effect == Effect(ALLOW)
}

func (s *Statement) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}

func (s *Statement) PrettyJSON() string {
	data, _ := json.MarshalIndent(s, "", "   ")
	return string(data)
}
