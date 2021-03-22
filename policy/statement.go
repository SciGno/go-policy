package policy

import (
	"encoding/json"
	"fmt"
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

	ve := ValidationEvent{}
	ve.StatementID = s.StatementID
	ve.Effect = s.Effect
	ve.Type = ValidationResultType(VALIDATION)

	ve.Resource = request.Resource
	if !registry.GetResourceValidator().Validate(s.Resource, request.Resource) {
		ve.Result = ValidationResult(DENIED)
		ve.Location = ValidationLocation(RESOURCE)
		fmt.Println(ve.PrettyJSON())
		return false
	}

	ve.Action = request.Action
	if !registry.GetActionValidator().Validate(s.Action, request.Action) {
		ve.Result = ValidationResult(DENIED)
		ve.Location = ValidationLocation(ACTION)
		fmt.Println(ve.PrettyJSON())
		return false
	}

	ve.Condition = request.Condition
	for reqCondName, reqCond := range request.Condition {
		if cv := registry.GetConditionValidator(reqCondName); cv != nil {
			if c, ok := s.Condition[reqCondName]; ok {
				if !cv.Validate(c, reqCond) {
					ve.Result = ValidationResult(DENIED)
					ve.Location = ValidationLocation(CONDITION)
					fmt.Println(ve.PrettyJSON())
					return false
				}
			}
		} else {
			ve.Type = ValidationResultType(ERROR)
			ve.Result = ValidationResult(NO_CONDITION)
			ve.Location = ValidationLocation(CONDITION)
			fmt.Println(ve.PrettyJSON())
			return false
		}
	}

	fmt.Println(ve.PrettyJSON())

	return true
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
