package policy

const (
	// ACTION variable
	ACTION = "action"
	// RESOURCE variable
	RESOURCE = "resource"
	// CONDITION variable
	CONDITION = "condition"
	// SID variable
	SID = "sid"
	// EFFECT variable
	EFFECT = "effect"
	// ID variable
	ID = "id"
	// NAME variable
	NAME = "name"
	// VERSION variable
	VERSION = "version"
	// STATEMENT variable
	STATEMENT = "statement"
)

// Policy struct
type Policy struct {
	PolicyID  string      `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Version   string      `json:"version,omitempty"`
	Statement []Statement `json:"statement,omitempty"`
}

// NewPolicy returns a new policy with the given arguments
func NewPolicy(id string, name string, version string, statements []Statement) Policy {
	return Policy{
		PolicyID:  id,
		Name:      name,
		Version:   version,
		Statement: statements,
	}
}

// Validate all statements against a Request
func (p *Policy) Validate(request *Request, registry *Registry) bool {

	if !p.ValidateDeny(request, registry) {
		return false
	}

	if p.ValidateAllow(request, registry) {
		return true
	}

	return false
}

// ValidateDeny checks all Effect=Deny statements in this policy
func (p *Policy) ValidateDeny(request *Request, registry *Registry) bool {

	for _, s := range p.Statement {
		if !s.IsAllow() {
			if s.Validate(request, registry) {
				return false
			}
		}
	}

	return true
}

// ValidateAllow checks all Effect=Allow statements in this policy
func (p *Policy) ValidateAllow(request *Request, registry *Registry) bool {

	for _, s := range p.Statement {
		if s.IsAllow() {
			if s.Validate(request, registry) {
				return true
			}
		}
	}

	return false
}
