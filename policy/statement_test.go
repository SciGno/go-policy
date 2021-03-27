package policy

import (
	"reflect"
	"testing"
)

func TestNewStatement(t *testing.T) {
	type args struct {
		id         string
		effect     Effect
		resource   string
		action     []string
		conditions map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Statement
	}{
		{
			"CreateStatement",
			args{"", ALLOW, "res1", []string{"read"}, nil},
			NewStatement("", ALLOW, "res1", []string{"read"}, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatement(tt.args.id, tt.args.effect, tt.args.resource, tt.args.action, tt.args.conditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatement_Validate(t *testing.T) {

	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

	type fields struct {
		StatementID string
		Effect      Effect
		Resource    string
		Action      []string
		Condition   map[string]interface{}
	}
	type args struct {
		request  *Request
		registry *Registry
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   StatementResult
	}{
		{
			"Passes_Action",
			fields{"", ALLOW, "res1", []string{"read"}, nil},
			args{&Request{Resource: "res1", Action: "read", Condition: nil}, &registry},
			StatementResult{
				Match:       true,
				StatementID: "",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
			},
		},
		{
			"Fails_Action",
			fields{"", ALLOW, "res1", []string{"read,update"}, nil},
			args{&Request{Resource: "res1", Action: "write", Condition: nil}, &registry},
			StatementResult{
				Match:       false,
				StatementID: "",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "write",
			},
		},
		{
			"Passes_Resource",
			fields{"", ALLOW, "res1", []string{"read"}, nil},
			args{&Request{"res1", "read", nil}, &registry},
			StatementResult{
				Match:       true,
				StatementID: "",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
			},
		},
		{
			"Passes_Condition",
			fields{"", ALLOW, "res1", []string{"read"}, map[string]interface{}{"AfterTime": "12:00"}},
			args{&Request{Resource: "res1", Action: "read", Condition: map[string]interface{}{"AfterTime": "13:00"}}, &registry},
			StatementResult{
				Match:       true,
				StatementID: "",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
				Condition:   map[string]interface{}{"AfterTime": "13:00"},
			},
		},
		{
			"Fails_Resource",
			fields{"", ALLOW, "res1", []string{"read"}, nil},
			args{&Request{"res2", "read", nil}, &registry},
			StatementResult{
				Match:       false,
				StatementID: "",
				Effect:      Effect(ALLOW),
				Resource:    "res2",
			},
		},
		{
			"Fails_Condition",
			fields{"", ALLOW, "res1", []string{"read"}, map[string]interface{}{"AfterTime": "12:00"}},
			args{&Request{"res1", "read", map[string]interface{}{"AfterTime": "11:00"}}, &registry},
			StatementResult{
				Match:       false,
				StatementID: "",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
				Condition:   map[string]interface{}{"AfterTime": "11:00"},
			},
		},
		{
			"Fails_ConditionDoesNotExist",
			fields{"", ALLOW, "res1", []string{"read"}, map[string]interface{}{"AfterTime": "12:00"}},
			args{&Request{"res1", "read", map[string]interface{}{"NotExist": "11:00"}}, &registry},
			StatementResult{
				Match:       false,
				StatementID: "",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
				Condition:   map[string]interface{}{"NotExist": "11:00"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statement{
				StatementID: tt.fields.StatementID,
				Effect:      tt.fields.Effect,
				Resource:    tt.fields.Resource,
				Action:      tt.fields.Action,
				Condition:   tt.fields.Condition,
			}
			if got := s.Validate(tt.args.request, tt.args.registry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nStatement.Validate() = \n%+v, \nwant \n%+v", got, tt.want)
			}
		})
	}
}

func TestStatement_IsAllow(t *testing.T) {
	type fields struct {
		StatementID string
		Effect      Effect
		Resource    string
		Action      []string
		Condition   map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"IsAllow",
			fields{"", ALLOW, "res1", []string{"read,update"}, nil},
			true,
		},
		{
			"IsDeny",
			fields{"", DENY, "res1", []string{"read,update"}, nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statement{
				StatementID: tt.fields.StatementID,
				Effect:      tt.fields.Effect,
				Resource:    tt.fields.Resource,
				Action:      tt.fields.Action,
				Condition:   tt.fields.Condition,
			}
			if got := s.IsAllow(); got != tt.want {
				t.Errorf("Statement.IsAllow() = %v, want %v", got, tt.want)
			}
		})
	}
}
