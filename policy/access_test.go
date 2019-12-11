package policy

import (
	"testing"
)

func TestAccessValidator_Validate(t *testing.T) {
	type fields struct {
		registry Registry
	}
	type args struct {
		req      *Request
		policies []Policy
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Pass",
			fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})},
			args{
				&Request{"a", "b", nil},
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "00:00"})}),
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, []string{"x"}, []string{"b"}, map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			false,
		},
		{
			"AllowFails",
			fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})},
			args{
				&Request{"x", "b", nil},
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "00:00"})}),
				},
			},
			true,
		},
		{
			"DenyFails",
			fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})},
			args{
				&Request{"x", "b", nil},
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, []string{"x"}, []string{"b"}, map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			true,
		},
		{
			"BadRequest",
			fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})},
			args{
				&Request{},
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, []string{"x"}, []string{"b"}, map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AccessValidator{
				registry: tt.fields.registry,
			}
			if err := p.Validate(tt.args.req, tt.args.policies); (err != nil) != tt.wantErr {
				t.Errorf("AccessValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccessValidator_validateDeny(t *testing.T) {
	type fields struct {
		registry Registry
	}
	type args struct {
		req   *Request
		stmts []Statement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"AllPass",
			fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})},
			args{
				&Request{"x", "b", nil},
				[]Statement{
					NewStatement("", DENY, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "23:59"})},
			},
			false,
		},
		{"ActionDenied", fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})}, args{&Request{"a", "b", nil}, []Statement{NewStatement("", DENY, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "23:59"})}}, true},
		{"ResourceFails", fields{NewRegistry(&StringMatch{}, &StringMatch{}, nil)}, args{&Request{"a", "x", nil}, []Statement{NewStatement("", DENY, []string{"a"}, []string{"b"}, nil)}}, true},
		{"ConditionDenied", fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})}, args{&Request{"a", "b", nil}, []Statement{NewStatement("", DENY, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "08:00"})}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AccessValidator{
				registry: tt.fields.registry,
			}
			if err := p.validateDeny(tt.args.req, tt.args.stmts); (err != nil) != tt.wantErr {
				t.Errorf("AccessValidator.validateDeny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccessValidator_validateAllow(t *testing.T) {
	type fields struct {
		registry Registry
	}
	type args struct {
		req   *Request
		stmts []Statement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"AllPass", fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})}, args{&Request{"a", "b", nil}, []Statement{NewStatement("", ALLOW, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "00:00"})}}, false},
		{"ResourceFails", fields{NewRegistry(&StringMatch{}, &StringMatch{}, nil)}, args{&Request{"a", "x", nil}, []Statement{NewStatement("", ALLOW, []string{"a"}, []string{"b"}, nil)}}, true},
		{"ConditionFails", fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})}, args{&Request{"a", "b", nil}, []Statement{NewStatement("", ALLOW, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "23:59"})}}, true},
		{"ActionFails", fields{NewRegistry(&StringMatch{}, &StringMatch{}, map[string]Validator{"AfterTime": &AfterTime{}})}, args{&Request{"x", "b", nil}, []Statement{NewStatement("", ALLOW, []string{"a"}, []string{"b"}, map[string]interface{}{"AfterTime": "00:00"})}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AccessValidator{
				registry: tt.fields.registry,
			}
			if err := p.validateAllow(tt.args.req, tt.args.stmts); (err != nil) != tt.wantErr {
				t.Errorf("AccessValidator.validateAllow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccessValidator_validateyConfition(t *testing.T) {
	r1 := NewRequest("publication:addPublication", "us:aws:graphql", map[string]interface{}{"SourceIP": "10.10.20.12"})
	type fields struct {
		registry Registry
	}
	type args struct {
		name    string
		data    interface{}
		request *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"TestIP", fields{NewRegistry(nil, nil, map[string]Validator{"SourceIP": &CIDR{}})}, args{"SourceIP", []string{"10.10.20.12"}, &r1}, false},
		{"IPFails", fields{NewRegistry(nil, nil, map[string]Validator{"SourceIP": &CIDR{}})}, args{"SourceIP", []string{"10.10.20.123"}, &r1}, true},
		{"TestPassNill", fields{NewRegistry(nil, nil, map[string]Validator{"AfterTime": &AfterTime{}})}, args{"AfterTime", "00:00", &r1}, false},
		{"TestPassNill", fields{NewRegistry(nil, nil, map[string]Validator{"AfterTime": &AfterTime{}})}, args{"AfterTime", "23:59", &r1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AccessValidator{
				registry: tt.fields.registry,
			}
			if err := p.validateyConfition(tt.args.name, tt.args.data, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AccessValidator.validateyConfition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccessValidator_validateyResource(t *testing.T) {
	r1 := NewRequest("publication:addPublication", "us:aws:graphql", map[string]interface{}{"SourceIP": "10.10.20.12"})
	type fields struct {
		registry Registry
	}
	type args struct {
		resources []string
		request   *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"TestPass", fields{NewRegistry(nil, &StringMatch{}, map[string]Validator{})}, args{[]string{"us:aws:graphql"}, &r1}, false},
		{"TestFails", fields{NewRegistry(nil, &StringMatch{}, map[string]Validator{})}, args{[]string{"us:aws2:graphql"}, &r1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AccessValidator{
				registry: tt.fields.registry,
			}
			if err := p.validateyResource(tt.args.resources, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AccessValidator.validateyResource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccessValidator_validateyAction(t *testing.T) {
	r1 := NewRequest("publication:addPublication", "us:aws:graphql", map[string]interface{}{"SourceIP": "10.10.20.12"})
	type fields struct {
		registry Registry
	}
	type args struct {
		actions []string
		request *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"TestPass", fields{NewRegistry(&ServiceWithAction{}, nil, map[string]Validator{})}, args{[]string{"publication:*"}, &r1}, false},
		{"TestFails", fields{NewRegistry(&ServiceWithAction{}, nil, map[string]Validator{})}, args{[]string{"publication:deletePub"}, &r1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AccessValidator{
				registry: tt.fields.registry,
			}
			if err := p.validateyAction(tt.args.actions, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AccessValidator.validateyAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
