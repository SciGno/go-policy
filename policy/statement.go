package policy

// Effect type
type Effect string

// Statement strct
type Statement struct {
	StatementID string                 `json:"sid,omitempty"`
	Effect      Effect                 `json:"effect,omitempty"`
	Action      []string               `json:"action,omitempty"`
	Resource    []string               `json:"resource,omitempty"`
	Condition   map[string]interface{} `json:"condition,omitempty"`
}

// // Statements type
// type Statements []Statement

// func (s Statements) Len() int           { return len(s) }
// func (s Statements) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s Statements) Less(i, j int) bool { return s[i].Effect < s[j].Effect }

// NewStatement returns a new request
func NewStatement(id string, effect Effect, action []string, resource []string, conditions map[string]interface{}) Statement {
	return Statement{
		StatementID: id,
		Effect:      effect,
		Action:      action,
		Resource:    resource,
		Condition:   conditions,
	}
}
