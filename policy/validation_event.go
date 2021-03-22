package policy

import (
	"encoding/json"
	"fmt"
)

type ValidationResult int

const (
	ALLOWED ValidationResult = iota
	DENIED
	FAILED
	NO_CONDITION
)

type ValidationResultType int

const (
	VALIDATION ValidationResultType = iota
	INFO
	ERROR
)

type ValidationLocation int

const (
	RESOURCE ValidationLocation = iota
	ACTION
	CONDITION
)

// ValidationEvent struct
type ValidationEvent struct {
	Type        ValidationResultType   `json:"type"`
	Result      ValidationResult       `json:"result"`
	Location    ValidationLocation     `json:"location"`
	PolicyID    string                 `json:"policy_id"`
	StatementID string                 `json:"statement_id"`
	Effect      Effect                 `json:"effect"`
	Action      string                 `json:"action,omitempty"`
	Resource    string                 `json:"resource,omitempty"`
	Condition   map[string]interface{} `json:"condition,omitempty"`
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
