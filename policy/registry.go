package policy

// Validator interface validates statement data against request data.
// In the event that a request does not provide data for the Validator,
// a nill value will be passed to the requestData parameter.
type Validator interface {
	Validate(stmntData, requestData interface{}) bool
}

// Registry struct
type Registry struct {
	resource  Validator
	action    Validator
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

// SetResourceValidator raplaces the current resource validator with rv
func (r *Registry) SetResourceValidator(rv Validator) {
	r.resource = rv
}

// GetActionValidator returns an action validator from this registry
func (r *Registry) GetActionValidator() Validator {
	return r.action
}

// SetActionValidator replaces the action validator with av
func (r *Registry) SetActionValidator(av Validator) {
	r.action = av
}

// GetConditionValidator returns a named condition validator from this registry
func (r *Registry) GetConditionValidator(name string) Validator {
	return r.condition[name]
}

// GetConditionValidators returns a new map with all the condition validators from this registry
func (r *Registry) GetConditionValidators() map[string]Validator {
	return r.condition
}

// AddConditionValidator adds a named condition validator to this registry
func (r *Registry) AddConditionValidator(name string, validator Validator) {
	r.condition[name] = validator
}

// RemoveConditionValidator removes a named condition validator from this registry
func (r *Registry) RemoveConditionValidator(name string) {
	delete(r.condition, name)
}
