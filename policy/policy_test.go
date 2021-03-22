package policy

import (
	"reflect"
	"testing"
)

func TestNewPolicy(t *testing.T) {
	type args struct {
		id         string
		name       string
		version    string
		statements []Statement
	}
	tests := []struct {
		name string
		args args
		want Policy
	}{
		{
			"PassNew",
			args{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}},
			NewPolicy("", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPolicy(tt.args.id, tt.args.name, tt.args.version, tt.args.statements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_Validate(t *testing.T) {

	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

	type fields struct {
		PolicyID  string
		Name      string
		Version   string
		Statement []Statement
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
			"AllowAnyNotRead",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"*"}, "*", nil}, {"", DENY, []string{"read"}, "*", nil}}},
			args{&Request{"anyaction", "res1:a", nil}, &registry},
			true,
		},
		{
			"AllowAnyNotRead2",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"*"}, "*", nil}, {"", DENY, []string{"read"}, "*", nil}}},
			args{&Request{"read", "res1:a", nil}, &registry},
			false,
		},
		{
			"DenyRead",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}, {"", DENY, []string{"read"}, "res1:a", nil}}},
			args{&Request{"read", "res1:a", nil}, &registry},
			false,
		},
		{
			"FailOnNoPolicyExists",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}},
			args{&Request{"read", "res1:b", nil}, &registry},
			false,
		},
		{
			"FailOnAction",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}},
			args{&Request{"write", "res1:a", nil}, &registry},
			false,
		},
		{
			"AllowAnyAction",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"*"}, "res1:a", nil}}},
			args{&Request{"anyaction", "res1:a", nil}, &registry},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Policy{
				PolicyID:  tt.fields.PolicyID,
				Name:      tt.fields.Name,
				Version:   tt.fields.Version,
				Statement: tt.fields.Statement,
			}
			got := p.Validate(tt.args.request, tt.args.registry)
			if got != tt.want {
				t.Errorf("Policy.Validate() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPolicy_ValidateDeny(t *testing.T) {

	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

	type fields struct {
		PolicyID  string
		Name      string
		Version   string
		Statement []Statement
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
			"WillAllowRead",
			fields{"", "", "", []Statement{{"", DENY, []string{"read"}, "*", nil}}},
			args{&Request{"anyaction", "res1:a", nil}, &registry},
			true,
		},
		{
			"DenyReadOnAnyResource",
			fields{"", "", "", []Statement{{"", DENY, []string{"read"}, "*", nil}}},
			args{&Request{"read", "res1:a", nil}, &registry},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Policy{
				PolicyID:  tt.fields.PolicyID,
				Name:      tt.fields.Name,
				Version:   tt.fields.Version,
				Statement: tt.fields.Statement,
			}
			if got := p.ValidateDeny(tt.args.request, tt.args.registry); got != tt.want {
				t.Errorf("Policy.ValidateDeny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicy_ValidateAllow(t *testing.T) {

	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

	type fields struct {
		PolicyID  string
		Name      string
		Version   string
		Statement []Statement
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
			"AllowAnyAction",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"*"}, "res1:a", nil}}},
			args{&Request{"anyaction", "res1:a", nil}, &registry},
			true,
		},
		{
			"AllowRead",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}},
			args{&Request{"anyaction", "res1:a", nil}, &registry},
			false,
		},
		{
			"DenyReadBecauseOfCondiution",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", map[string]interface{}{"AfterTime": "12:00"}}}},
			args{&Request{"read", "res1:a", map[string]interface{}{"AfterTime": "11:00"}}, &registry},
			false,
		},
		{
			"DenyCondiutionNotExists",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", map[string]interface{}{"AfterTime": "12:00"}}}},
			args{&Request{"read", "res1:a", map[string]interface{}{"SomeCondition": "11:00"}}, &registry},
			false,
		},
		{
			"AllowReasourceFails",
			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}},
			args{&Request{"read", "res1:b", nil}, &registry},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Policy{
				PolicyID:  tt.fields.PolicyID,
				Name:      tt.fields.Name,
				Version:   tt.fields.Version,
				Statement: tt.fields.Statement,
			}
			if got := p.ValidateAllow(tt.args.request, tt.args.registry); got != tt.want {
				t.Errorf("Policy.ValidateAllow() = %v, want %v", got, tt.want)
			}
		})
	}
}
