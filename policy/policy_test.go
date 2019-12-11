package policy

import (
	"reflect"
	"testing"
)

func TestPolicy_AllowStatements(t *testing.T) {
	type fields struct {
		PolicyID  string
		Name      string
		Version   string
		Statement []Statement
	}
	tests := []struct {
		name   string
		fields fields
		want   []Statement
	}{
		{"GetALLOW", fields{"123", "name", "1.0", []Statement{NewStatement("other", ALLOW, []string{}, []string{}, map[string]interface{}{})}}, []Statement{NewStatement("other", ALLOW, []string{}, []string{}, map[string]interface{}{})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Policy{
				PolicyID:  tt.fields.PolicyID,
				Name:      tt.fields.Name,
				Version:   tt.fields.Version,
				Statement: tt.fields.Statement,
			}
			if got := p.AllowStatements(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Policy.AllowStatements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_DenyStatements(t *testing.T) {
	type fields struct {
		PolicyID  string
		Name      string
		Version   string
		Statement []Statement
	}
	tests := []struct {
		name   string
		fields fields
		want   []Statement
	}{
		{"GetDENY", fields{"123", "name", "1.0", []Statement{NewStatement("other", DENY, []string{}, []string{}, map[string]interface{}{})}}, []Statement{NewStatement("other", DENY, []string{}, []string{}, map[string]interface{}{})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Policy{
				PolicyID:  tt.fields.PolicyID,
				Name:      tt.fields.Name,
				Version:   tt.fields.Version,
				Statement: tt.fields.Statement,
			}
			if got := p.DenyStatements(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Policy.DenyStatements() = %v, want %v", got, tt.want)
			}
		})
	}
}
