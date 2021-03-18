package policy

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
	FAIL    ValidationResult = 0
	SUCCESS ValidationResult = 1
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

func (s *ValidationEvent) JSON(message string) *ValidationEvent {
	return s
}
