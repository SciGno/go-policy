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

const (
	// ALLOW effect for policy statement
	ALLOW Effect = "Allow"
	// DENY effect for policy statement
	DENY Effect = "Deny"
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

// AllowStatements returns all statements with Effect=Allow
func (p *Policy) AllowStatements() []Statement {
	a := []Statement{}
	for _, s := range p.Statement {
		if s.Effect == ALLOW {
			a = append(a, s)
		}
	}
	return a
}

// DenyStatements returns all statements with Effect=Deny
func (p *Policy) DenyStatements() []Statement {
	a := []Statement{}
	for _, s := range p.Statement {
		if s.Effect == DENY {
			a = append(a, s)
		}
	}
	return a
}
