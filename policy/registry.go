package policy

// Validator interface validates statement data against request data.
// In the event that a request does not provide data for the Validator,
// a nill value will be passed to the requestData parameter.
type Validator interface {
	Validate(stmntData, requestData interface{}) bool
}

// Registry struct
type Registry struct {
	Action    Validator
	Resource  Validator
	Condition map[string]Validator
}

// NewRegistry func
func NewRegistry(a, b Validator, c map[string]Validator) Registry {
	return Registry{
		Action:    a,
		Resource:  b,
		Condition: c,
	}
}
