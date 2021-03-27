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

	policyResult := p.ValidateDeny(request, registry)
	if !policyResult.IsAllowed {
		return policyResult
	}

	policyResult = p.ValidateAllow(request, registry)
	if policyResult.IsAllowed {
		return policyResult
	}

	return policyResult
}

// ValidateDeny checks all Effect=Deny statements in this policy
func (p *Policy) ValidateDeny(request *Request, registry *Registry) PolicyResult {

	for _, statement := range p.Statement {
		if !statement.IsAllow() {
			statementResult := statement.Validate(request, registry)
			policyResult := PolicyResult{p.PolicyID, true, statementResult}
			if statementResult.Match {
				policyResult.IsAllowed = false
				policyResult.StatementResult = statementResult
				return policyResult
			}
		}
	}

	return PolicyResult{p.PolicyID, true, StatementResult{}}
}

// ValidateAllow checks all Effect=Allow statements in this policy
func (p *Policy) ValidateAllow(request *Request, registry *Registry) PolicyResult {

	policyResult := PolicyResult{p.PolicyID, false, StatementResult{}}

	for _, statement := range p.Statement {
		if statement.IsAllow() {
			statementResult := statement.Validate(request, registry)
			policyResult.StatementResult = statementResult
			if statementResult.Match {
				policyResult.IsAllowed = true
				return policyResult
			}
		}
	}

	return policyResult
}
