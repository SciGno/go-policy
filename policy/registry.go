package policy

// Validator interface
type Validator interface {
	Validate(stmnt, req interface{}) bool
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
