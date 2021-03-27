package policy

import (
	"reflect"
	"testing"
)

// import (
// 	"reflect"
// 	"testing"
// )

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
			args{"123", "Test Policy", "1.0", []Statement{{"", ALLOW, "res1:a", []string{"read"}, nil}}},
			NewPolicy("123", "Test Policy", "1.0", []Statement{{"", ALLOW, "res1:a", []string{"read"}, nil}}),
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

// func TestPolicy_ValidateDeny(t *testing.T) {

// 	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

// 	type fields struct {
// 		PolicyID  string
// 		Name      string
// 		Version   string
// 		Statement []Statement
// 	}
// 	type args struct {
// 		request  *Request
// 		registry *Registry
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   bool
// 	}{
// 		{
// 			"WillAllowRead",
// 			fields{"", "", "", []Statement{{"", DENY, []string{"read"}, "*", nil}}},
// 			args{&Request{"anyaction", "res1:a", nil}, &registry},
// 			true,
// 		},
// 		{
// 			"DenyReadOnAnyResource",
// 			fields{"", "", "", []Statement{{"", DENY, []string{"read"}, "*", nil}}},
// 			args{&Request{"read", "res1:a", nil}, &registry},
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := &Policy{
// 				PolicyID:  tt.fields.PolicyID,
// 				Name:      tt.fields.Name,
// 				Version:   tt.fields.Version,
// 				Statement: tt.fields.Statement,
// 			}
// 			if got := p.ValidateDeny(tt.args.request, tt.args.registry); got != tt.want {
// 				t.Errorf("Policy.ValidateDeny() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPolicy_ValidateAllow(t *testing.T) {

// 	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

// 	type fields struct {
// 		PolicyID  string
// 		Name      string
// 		Version   string
// 		Statement []Statement
// 	}
// 	type args struct {
// 		request  *Request
// 		registry *Registry
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   bool
// 	}{
// 		{
// 			"AllowAnyAction",
// 			fields{"", "", "", []Statement{{"", ALLOW, []string{"*"}, "res1:a", nil}}},
// 			args{&Request{"anyaction", "res1:a", nil}, &registry},
// 			true,
// 		},
// 		{
// 			"AllowRead",
// 			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}},
// 			args{&Request{"anyaction", "res1:a", nil}, &registry},
// 			false,
// 		},
// 		{
// 			"DenyReadBecauseOfCondiution",
// 			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", map[string]interface{}{"AfterTime": "12:00"}}}},
// 			args{&Request{"read", "res1:a", map[string]interface{}{"AfterTime": "11:00"}}, &registry},
// 			false,
// 		},
// 		{
// 			"DenyCondiutionNotExists",
// 			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", map[string]interface{}{"AfterTime": "12:00"}}}},
// 			args{&Request{"read", "res1:a", map[string]interface{}{"SomeCondition": "11:00"}}, &registry},
// 			false,
// 		},
// 		{
// 			"AllowReasourceFails",
// 			fields{"", "", "", []Statement{{"", ALLOW, []string{"read"}, "res1:a", nil}}},
// 			args{&Request{"read", "res1:b", nil}, &registry},
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := &Policy{
// 				PolicyID:  tt.fields.PolicyID,
// 				Name:      tt.fields.Name,
// 				Version:   tt.fields.Version,
// 				Statement: tt.fields.Statement,
// 			}
// 			if got := p.ValidateAllow(tt.args.request, tt.args.registry); got != tt.want {
// 				t.Errorf("Policy.ValidateAllow() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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
		want   PolicyResult
	}{
		{
			"AllowAnyNotRead",
			fields{"123", "Test Policy", "1.0", []Statement{{"", ALLOW, "*", []string{"*"}, nil}, {"", DENY, "*", []string{"read"}, nil}}},
			args{&Request{"res1:a", "anyaction", nil}, &registry},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: true,
				StatementResult: StatementResult{
					Match:       true,
					StatementID: "",
					Effect:      Effect(ALLOW),
					Resource:    "res1:a",
					Action:      "anyaction",
				},
			},
		},
		{
			"AllowAnyNotRead2",
			fields{"123", "Test Policy", "1.0", []Statement{{"", ALLOW, "*", []string{"*"}, nil}, {"", DENY, "*", []string{"read"}, nil}}},
			args{&Request{"res1:a", "read", nil}, &registry},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: false,
				StatementResult: StatementResult{
					Match:       true,
					StatementID: "",
					Effect:      Effect(DENY),
					Resource:    "res1:a",
					Action:      "read",
				},
			},
		},
		{
			"DenyRead",
			fields{"123", "Test Policy", "1.0", []Statement{{"", ALLOW, "res1:a", []string{"read"}, nil}, {"", DENY, "res1:a", []string{"read"}, nil}}},
			args{&Request{"res1:a", "read", nil}, &registry},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: false,
				StatementResult: StatementResult{
					Match:       true,
					StatementID: "",
					Effect:      Effect(DENY),
					Resource:    "res1:a",
					Action:      "read",
				},
			},
		},
		{
			"FailOnNoPolicyExists",
			fields{"123", "Test Policy", "1.0", []Statement{{"", ALLOW, "res1:a", []string{"read"}, nil}}},
			args{&Request{"res1:b", "read", nil}, &registry},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: false,
				StatementResult: StatementResult{
					Match:       false,
					StatementID: "",
					Effect:      Effect(ALLOW),
					Resource:    "res1:b",
				},
			},
		},
		{
			"FailOnAction",
			fields{"123", "Test Policy", "1.0", []Statement{{"456", ALLOW, "res1:a", []string{"read"}, nil}}},
			args{&Request{"res1:a", "write", nil}, &registry},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: false,
				StatementResult: StatementResult{
					Match:       false,
					StatementID: "456",
					Effect:      Effect(ALLOW),
					Resource:    "res1:a",
					Action:      "write",
				},
			},
		},
		{
			"AllowAnyAction",
			fields{"123", "Test Policy", "1.0", []Statement{{"456", ALLOW, "res1:a", []string{"*"}, map[string]interface{}{"AfterTime": "12:00"}}}},
			args{&Request{"res1:a", "anyaction", map[string]interface{}{"AfterTime": "13:00"}}, &registry},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: true,
				StatementResult: StatementResult{
					Match:       true,
					StatementID: "456",
					Effect:      Effect(ALLOW),
					Resource:    "res1:a",
					Action:      "anyaction",
					Condition:   map[string]interface{}{"AfterTime": "13:00"},
				},
			},
		},
		{
			"FailsAnyActionOnCondition",
			fields{"123", "Test Policy", "1.0", []Statement{{"456", ALLOW, "res1:a", []string{"*"}, map[string]interface{}{"AfterTime": "12:00"}}}},
			args{&Request{"res1:a", "anyaction", map[string]interface{}{"AfterTime": "11:00"}}, &registry},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: false,
				StatementResult: StatementResult{
					Match:       false,
					StatementID: "456",
					Effect:      Effect(ALLOW),
					Resource:    "res1:a",
					Action:      "anyaction",
					Condition:   map[string]interface{}{"AfterTime": "11:00"},
				},
			},
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
			if got := p.Validate(tt.args.request, tt.args.registry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nPolicy.Validate() = \n%+v, \nwant \n%+v\n", got, tt.want)
			}
		})
	}
}
