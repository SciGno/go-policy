package policy

// Validator interface validates statement data against request data.
// In the event that a request does not provide data for the Validator,
// a nill value will be passed to the requestData parameter.
type Validator interface {
	Validate(stmntData, requestData interface{}) bool
}

// Registry struct
type Registry struct {
	action    Validator
	resource  Validator
	condition map[string]Validator
}

// NewRegistry func
func NewRegistry(resourceValidator, actionValidator Validator, conditionMap map[string]Validator) Registry {
	return Registry{
		resource:  resourceValidator,
		action:    actionValidator,
		condition: conditionMap,
	}
}

// GetResourceValidator returns a resource validator from this registry
func (r *Registry) GetResourceValidator() Validator {
	return r.resource
}

// GetActionValidator returns an action validator from this registry
func (r *Registry) GetActionValidator() Validator {
	return r.action
}

// GetConditionnValidator returns a named condition validator from this registry
func (r *Registry) GetConditionnValidator(name string) Validator {
	return r.condition[name]
}

// AddConditionnValidator adds a named condition validator to this registry
func (r *Registry) AddConditionnValidator(name string, validator Validator) {
	r.condition[name] = validator
}

// RemoveConditionnValidator removes a named condition validator from this registry
func (r *Registry) RemoveConditionnValidator(name string) {
	delete(r.condition, name)
}
