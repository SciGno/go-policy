package policy

// PolicyEvent struct
type PolicyResult struct {
	PolicyID        string          `json:"policy_id"`
	IsAllowed       bool            `json:"is_allowed"`
	StatementResult StatementResult `json:"statement_result"`
}

const (
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
func (p *Policy) Validate(request *Request, registry *Registry) PolicyResult {

	if pr := p.ValidateDeny(request, registry); !pr.IsAllowed {
		return pr
	}

	if pr := p.ValidateAllow(request, registry); pr.IsAllowed {
		return pr
	}

	return PolicyResult{p.PolicyID, false, StatementResult{}}
}

// ValidateDeny checks all Effect=Deny statements in this policy
func (p *Policy) ValidateDeny(request *Request, registry *Registry) PolicyResult {

	pr := PolicyResult{p.PolicyID, true, StatementResult{}}

	for _, s := range p.Statement {
		if !s.IsAllow() {
			sr := s.Validate(request, registry)
			if sr.Match {
				pr.IsAllowed = false
				pr.StatementResult = sr
				return pr
			}
		}
	}

	return pr
}

// ValidateAllow checks all Effect=Allow statements in this policy
func (p *Policy) ValidateAllow(request *Request, registry *Registry) PolicyResult {

	pr := PolicyResult{p.PolicyID, false, StatementResult{}}

	for _, s := range p.Statement {
		if s.IsAllow() {
			sr := s.Validate(request, registry)
			if sr.Match {
				pr.IsAllowed = true
				pr.StatementResult = sr
				return pr
			}
		}
	}

	return pr
}
