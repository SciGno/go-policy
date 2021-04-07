package policy

import (
	"encoding/json"
)

// StatementResult struct
type StatementResult struct {
	Match       bool              `json:"match"`
	Processed   Processed         `json:"processed"`
	StatementID string            `json:"statement_id"`
	Effect      Effect            `json:"effect"`
	Resource    string            `json:"resource,omitempty"`
	Action      string            `json:"action,omitempty"`
	Condition   map[string]string `json:"condition,omitempty"`
}

// Effect type
type Effect string

const (
	ALLOW Effect = "Allow"
	DENY  Effect = "Deny"
)

/**
Processed shows what was ccompleted in the validation process.
The validation process is executed in the following order:
First resource, followed by action and finally conditions.

If Processed = CONDITION, then the validation process completed successfuly.

NOTE:
A condition should never be used in a DENY statement in order to prevent
undesireble results.  Use an ALLOW statement to accomplish the desireble result.

Example:
   Requirement:
     DENY action on a resource after 6:00PM to 9:00AM.

   Statement should be:
     ALLOW action on a resource from 9:00PM to 6:00PM.

Since action can only be performed by ALLOW, a DENY statement should only
be used as an exeption to an ALLOW statement.

So the following DENY statement will never validate the condition
because the statement will be denied by matching the action.

Statement {
	StatementID: "Denied Statement",
	Effect: 	 DENY,
	Resource:	"res1",
	Action: 	"read",
	Condition: 	map[string]interface{}{"AfterTime": "12:00"},
}

Request {
	Resource: "res1",
	Action: "read",
	Condition: map[string]interface{}{"AfterTime": "13:00"}
}
**/
type Processed string

const (
	RESOURCE  Processed = "resource"
	ACTION    Processed = "action"
	CONDITION Processed = "condition"
)

type ConditionMap struct {
	Condition map[string]string
}

// Statement strct
type Statement struct {
	StatementID string                       `json:"sid"`
	Effect      Effect                       `json:"effect"`
	Resource    string                       `json:"resource"`
	Action      []string                     `json:"action"`
	Condition   map[string]map[string]string `json:"condition,omitempty"`
}

// NewStatement returns a new request
func NewStatement(id string, effect Effect, resource string, action []string, conditions map[string]map[string]string) Statement {
	return Statement{
		StatementID: id,
		Effect:      effect,
		Resource:    resource,
		Action:      action,
		Condition:   conditions,
	}
}

// Validate this statement against the specified request
// "default" validators will be used if a resource or action validator names
// are not provided in the request
func (s *Statement) Validate(request *Request, registry *Registry) StatementResult {

	statementResutl := StatementResult{}
	statementResutl.StatementID = s.StatementID
	statementResutl.Effect = s.Effect
	statementResutl.Match = true // Matching fails validation in a DENY statement

	statementResutl.Resource = request.Resource.Value
	statementResutl.Processed = Processed(RESOURCE)

	if len(request.Resource.Validator) == 0 {
		request.Resource.Validator = "default"
	}

	if len(request.Action.Validator) == 0 {
		request.Action.Validator = "default"
	}

	if !registry.GetValidator(string(RESOURCE), request.Resource.Validator).Validate(s.Resource, request.Resource.Value) {
		if s.IsAllow() { // not matching fails validation in an ALLOW statement
			statementResutl.Match = false
		}
		return statementResutl
	}

	// The resource exists.  We need to validate the action.
	statementResutl.Action = request.Action.Value
	statementResutl.Processed = Processed(ACTION)
	if registry.GetValidator(string(ACTION), request.Action.Validator).Validate(s.Action, request.Action.Value) {
		if !s.IsAllow() {
			return statementResutl
		}
	} else if s.IsAllow() {
		statementResutl.Match = false
		return statementResutl
	} else {
		statementResutl.Match = false
	}

	statementResutl.Condition = request.Metadata
	statementResutl.Processed = Processed(CONDITION)
	// if the statement contains conditions then the request MUST contain metadata that
	// can be validated against those conditions.
	for metaName, condition := range s.Condition {
		if registry.MetaNameExists(metaName) { // validator is registered
			for conditionKey, conditionValue := range condition {
				if !request.MetaNameExists(conditionKey) {
					statementResutl.Match = false
					return statementResutl
				}
				if !registry.GetValidator(metaName, conditionKey).Validate(conditionValue, request.Metadata[conditionKey]) {
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
