package policy

import (
	"reflect"
	"testing"
)

func TestStatement_Validate(t *testing.T) {

	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

	type fields struct {
		StatementID string
		Effect      Effect
		Action      []string
		Resource    string
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
		want   bool
	}{
		{
			"Passes_Action",
			fields{"", ALLOW, []string{"read"}, "res1", nil},
			args{&Request{"read", "res1", nil}, &registry},
			true,
		},
		{
			"Fails_Action",
			fields{"", ALLOW, []string{"read"}, "res1", nil},
			args{&Request{"write", "res1", nil}, &registry},
			false,
		},
		{
			"Passes_Resource",
			fields{"", ALLOW, []string{"read"}, "res1", nil},
			args{&Request{"read", "res1", nil}, &registry},
			true,
		},
		{
			"Fails_Resource",
			fields{"", ALLOW, []string{"read"}, "res1", nil},
			args{&Request{"read", "res2", nil}, &registry},
			false,
		},
		{
			"Passes_Condition",
			fields{"", ALLOW, []string{"read"}, "res1", map[string]interface{}{"AfterTime": "12:00"}},
			args{&Request{"read", "res1", map[string]interface{}{"AfterTime": "13:00"}}, &registry},
			true,
		},
		{
			"Fails_Condition",
			fields{"", ALLOW, []string{"read"}, "res1", map[string]interface{}{"AfterTime": "12:00"}},
			args{&Request{"read", "res1", map[string]interface{}{"AfterTime": "11:00"}}, &registry},
			false,
		},
		{
			"Fails_ConditionDoesNotExist",
			fields{"", ALLOW, []string{"read"}, "res1", map[string]interface{}{"AfterTime": "12:00"}},
			args{&Request{"read", "res1", map[string]interface{}{"NotExist": "11:00"}}, &registry},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statement{
				StatementID: tt.fields.StatementID,
				Effect:      tt.fields.Effect,
				Action:      tt.fields.Action,
				Resource:    tt.fields.Resource,
				Condition:   tt.fields.Condition,
			}
			got := s.Validate(tt.args.request, tt.args.registry)
			if got != tt.want {
				t.Errorf("Statement.Validate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatement(t *testing.T) {
	type args struct {
		id         string
		effect     Effect
		action     []string
		resource   string
		conditions map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Statement
	}{
		{
			"CreateStatement",
			args{"", ALLOW, []string{"read"}, "res1", nil},
			NewStatement("", ALLOW, []string{"read"}, "res1", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatement(tt.args.id, tt.args.effect, tt.args.action, tt.args.resource, tt.args.conditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
