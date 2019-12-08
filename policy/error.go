package policy

// ValidationError struct
type ValidationError struct {
	PolicyID  string      `json:"id,omitempty"`
	Statement string      `json:"statement,omitempty"`
	Location  string      `json:"location,omitempty"`
	Validator string      `json:"validator,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	message   string
}

// NewValidationError returns a new ValidationError that is initialized with the given string
func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		message: message,
	}
}

func (v *ValidationError) Error() string {
	return v.message
}
