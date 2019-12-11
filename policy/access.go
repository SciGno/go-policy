package policy

import "strings"

// AccessValidator validates all policies against a given request
type AccessValidator struct {
	registry Registry
}

// Validate checks the request against all statements. Effect=Deny statements will be certified first
// followed by Effect=Allow statements
func (p *AccessValidator) Validate(req *Request, policies []Policy) error {
	if len(strings.TrimSpace(req.Action)) == 0 || len(strings.TrimSpace(req.Resource)) == 0 {
		return NewValidationError("missing request parameters")
	}
	for _, pol := range policies {
		if err := p.validateDeny(req, pol.DenyStatements()); err != nil {
			e := err.(*ValidationError)
			e.PolicyID = pol.PolicyID
			return e
		}
	}
	for _, pol := range policies {
		if err := p.validateAllow(req, pol.AllowStatements()); err != nil {
			e := err.(*ValidationError)
			e.PolicyID = pol.PolicyID
			return e
		}
	}
	return nil
}

// ValidateDeny checks the request against all statements. Effect=Deny is assumed
func (p *AccessValidator) validateDeny(req *Request, stmts []Statement) error {
	for _, s := range stmts {
		if err := p.validateyResource(s.Resource, req); err != nil {
			e := err.(*ValidationError)
			e.Statement = s.StatementID
			e.Location = RESOURCE
			return e
		}
		if len(s.Condition) > 0 {
			for name, data := range s.Condition {
				if err := p.validateyConfition(name, data, req); err == nil {
					e := NewValidationError("policy validation failed")
					e.Statement = s.StatementID
					e.Location = CONDITION
					e.Validator = name
					e.Data = data
					return e
				}
			}
		}
		if err := p.validateyAction(s.Action, req); err == nil {
			e := NewValidationError("policy validation failed")
			e.Validator = ACTION
			e.Data = req.Action
			// e := err.(*ValidationError)
			e.Statement = s.StatementID
			e.Location = ACTION
			return e

		}
	}
	return nil
}

// ValidateAllow checks the request against all statements. Effect=Allow is assumed
func (p *AccessValidator) validateAllow(req *Request, stmts []Statement) error {
	for _, s := range stmts {
		if err := p.validateyResource(s.Resource, req); err != nil {
			e := err.(*ValidationError)
			e.Statement = s.StatementID
			e.Location = RESOURCE
			return e
		}
		if len(s.Condition) > 0 {
			for name, data := range s.Condition {
				if err := p.validateyConfition(name, data, req); err != nil {
					e := err.(*ValidationError)
					e.Statement = s.StatementID
					e.Location = CONDITION
					return e
				}
			}
		}
		if err := p.validateyAction(s.Action, req); err != nil {
			e := err.(*ValidationError)
			e.Statement = s.StatementID
			e.Location = ACTION
			return e
		}
	}
	return nil
}

func (p *AccessValidator) validateyAction(actions []string, request *Request) error {
	if validator := p.registry.Action; validator != nil {
		if validator.Validate(actions, request.Action) {
			return nil
		}
	}
	e := NewValidationError("policy validation failed")
	e.Validator = ACTION
	e.Data = actions
	return e
}

func (p *AccessValidator) validateyResource(resources []string, request *Request) error {
	if validator := p.registry.Resource; validator != nil {
		if validator.Validate(resources, request.Resource) {
			return nil
		}
	}
	e := NewValidationError("policy validation failed")
	e.Validator = RESOURCE
	e.Data = resources
	return e
}

func (p *AccessValidator) validateyConfition(name string, data interface{}, request *Request) error {
	if validator, ok := p.registry.Condition[name]; ok { // validator is registered
		if condVal, ok := request.Condition[name]; ok { // validator name in request
			if validator.Validate(data, condVal) {
				return nil
			}
		} else if validator.Validate(data, nil) {
			return nil
		}
	}
	e := NewValidationError("policy validation failed")
	e.Validator = name
	e.Data = data
	return e
}
