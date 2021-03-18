package policy

import (
	"encoding/json"
	"fmt"
)

const (
	MISSING_REQ_PARAM = "missing request parameter"
	NO_MATCH          = "no match"
	ERROR             = "error"
	INFO              = "info"
	TRACE             = "trace"
	ALL               = "all"
	NONE              = "none"
)

type ValidationResult int

const (
	SUCCESS ValidationResult = 1
	FAIL    ValidationResult = 0
)

// ValidationEvent struct
type ValidationEvent struct {
	Type        string           `json:"type"`
	Result      ValidationResult `json:"result"`
	PolicyID    string           `json:"id,omitempty"`
	StatementID string           `json:"statement_id,omitempty"`
	Effect      Effect           `json:"effect,omitempty"`
	Action      string           `json:"action,omitempty"`
	Resource    string           `json:"resource,omitempty"`
	Condition   []string         `json:"condition,omitempty"`
}

func (s *ValidationEvent) JSON() string {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(data)
}

func (s *ValidationEvent) PrettyJSON() string {
	data, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(data)
}
