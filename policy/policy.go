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

// // SetID sets id to id
// func (p *Policy) SetID(id string) {
// 	p.PolicyID = id
// }

// // SetName sets name to n
// func (p *Policy) SetName(n string) {
// 	p.Name = n
// }

// // SetVersion sets version to v
// func (p *Policy) SetVersion(v string) {
// 	p.Version = v
// }

// // AddStatements adds a Statement to the array
// func (p *Policy) AddStatements(s ...Statement) error {
// 	if len(s) > 0 {
// 		for _, v := range s {
// 			switch v.Effect {
// 			case ALLOW:
// 				p.allow = append(p.allow, &v)
// 			case DENY:
// 				p.deny = append(p.deny, &v)
// 			default:
// 				return errors.New("invalid effect")
// 			}
// 		}
// 	}
// 	return nil
// }

// // MarshalJSON returns the JSON encoding of p
// func (p *Policy) MarshalJSON() ([]byte, error) {
// 	tmp := struct {
// 		ID        string      `json:"id"`
// 		Name      string      `json:"name"`
// 		Version   string      `json:"version"`
// 		Statement []Statement `json:"statement"`
// 	}{
// 		ID:        p.id,
// 		Name:      p.name,
// 		Version:   p.version,
// 		Statement: nil,
// 	}

// 	for _, v := range p.deny {
// 		tmp.Statement = append(tmp.Statement, *v)
// 	}

// 	for _, v := range p.allow {
// 		tmp.Statement = append(tmp.Statement, *v)
// 	}

// 	return json.Marshal(&tmp)
// }

// // UnmarshalJSON parses the JSON-encoded data and stores the result in the Policy
// func (p *Policy) UnmarshalJSON(a []byte) error {

// 	o := struct {
// 		ID         string        `json:"id"`
// 		Name       string        `json:"name"`
// 		Version    string        `json:"version"`
// 		Statements []interface{} `json:"statement"`
// 	}{}

// 	if err := json.Unmarshal(a, &o); err == nil {

// 		if len(o.ID) == 0 {
// 			return errors.New("missing field 'id'")
// 		}

// 		if len(o.Name) == 0 {
// 			return errors.New("missing field 'name'")
// 		}

// 		if len(o.Version) == 0 {
// 			return errors.New("missing field 'version'")
// 		}

// 		if len(o.Statements) == 0 {
// 			return errors.New("missing field 'statement'")
// 		}

// 		p.id = o.ID
// 		p.name = o.Name
// 		p.version = o.Version

// 		for _, aVal := range o.Statements {
// 			newStatement := Statement{"", "", []string{}, []string{}, map[string]interface{}{}}

// 			m := aVal.(map[string]interface{})
// 			if _, ok := m[SID]; !ok {
// 				return errors.New("missing field " + SID)
// 			}
// 			if _, ok := m[EFFECT]; !ok {
// 				return errors.New("missing field " + EFFECT)
// 			}
// 			if _, ok := m[ACTION]; !ok {
// 				return errors.New("missing field " + ACTION)
// 			}
// 			for name, val := range aVal.(map[string]interface{}) {
// 				switch name {
// 				case SID:
// 					if len(val.(string)) == 0 {
// 						return errors.New("field 'sid' is empty")
// 					}
// 					newStatement.StatementID = val.(string)
// 				case EFFECT:
// 					if len(val.(string)) == 0 {
// 						return errors.New("field 'effect' is empty")
// 					}
// 					newStatement.Effect = Effect(val.(string))
// 				case ACTION:
// 					if val != nil {
// 						a := []string{}
// 						for _, s := range val.([]interface{}) {
// 							a = append(a, s.(string))
// 						}
// 						newStatement.Action = a
// 					} else {
// 						return errors.New("action specified but found nil")
// 					}
// 				case RESOURCE:
// 					if val != nil {
// 						a := []string{}
// 						for _, s := range val.([]interface{}) {
// 							a = append(a, s.(string))
// 						}
// 						newStatement.Resource = a
// 					} else {
// 						return errors.New("resource specified but found nil")
// 					}
// 				case CONDITION:
// 					if val != nil {
// 						newStatement.Condition = val.(map[string]interface{})
// 					} else {
// 						return errors.New("condition specified but found nil")
// 					}
// 				}
// 			}
// 			p.AddStatements(newStatement)
// 		}
// 	} else {
// 		return err
// 	}

// 	return nil
// }
