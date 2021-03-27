package policy

import (
	"reflect"
	"testing"
)

func TestAccessValidator_Validate(t *testing.T) {

	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})

	type fields struct {
		registry Registry
	}
	type args struct {
		request  *Request
		policies []Policy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ValidationResult
	}{
		{
			"Pass",
			fields{registry},
			args{
				&Request{Resource: "res1", Action: "read", Condition: nil},
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, "res1", []string{"read"}, map[string]interface{}{"AfterTime": "00:00"})}),
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, "res1", []string{"write"}, map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			ValidationResult{
				PolicyResult{
					PolicyID:  "123",
					IsAllowed: true,
					StatementResult: StatementResult{
						Match:       true,
						StatementID: "",
						Effect:      Effect(ALLOW),
						Resource:    "res1",
						Action:      "read",
					},
				},
			},
		},
		{
			"Fail",
			fields{registry},
			args{
				&Request{Resource: "res1", Action: "write", Condition: nil},
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, "res1", []string{"read"}, map[string]interface{}{"AfterTime": "00:00"})}),
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, "res1", []string{"write"}, map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			ValidationResult{
				PolicyResult{
					PolicyID:  "123",
					IsAllowed: false,
					StatementResult: StatementResult{
						Match:       true,
						StatementID: "",
						Effect:      Effect(DENY),
						Resource:    "res1",
						Action:      "write",
					},
				},
			},
		},
		{
			"FailNoPolicyExists",
			fields{registry},
			args{
				&Request{Resource: "res10", Action: "write", Condition: nil},
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, "res1", []string{"read"}, map[string]interface{}{"AfterTime": "00:00"})}),
				},
			},
			ValidationResult{
				PolicyResult{
					PolicyID:  "",
					IsAllowed: false,
					StatementResult: StatementResult{
						Match:       false,
						StatementID: "",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &AccessValidator{
				registry: tt.fields.registry,
			}
			if got := v.Validate(tt.args.request, tt.args.policies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nAccessValidator.Validate() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}

func TestValidationResult_JSON(t *testing.T) {
	type fields struct {
		PolicyResult PolicyResult
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"TestJSON",
			fields{
				PolicyResult: PolicyResult{
					PolicyID:  "123",
					IsAllowed: true,
					StatementResult: StatementResult{
						Match:       true,
						StatementID: "",
						Effect:      Effect(ALLOW),
						Resource:    "res1",
						Action:      "read",
					},
				},
			},
			"{\"policy_id\":\"123\",\"is_allowed\":true,\"statement_result\":{\"match\":true,\"statement_id\":\"\",\"effect\":\"Allow\",\"resource\":\"res1\",\"action\":\"read\"}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ValidationResult{
				PolicyResult: tt.fields.PolicyResult,
			}
			if got := v.JSON(); got != tt.want {
				t.Errorf("\nValidationResult.JSON() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}

func TestValidationResult_PrettyJSON(t *testing.T) {
	type fields struct {
		PolicyResult PolicyResult
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"TestJSON",
			fields{
				PolicyResult: PolicyResult{
					PolicyID:  "123",
					IsAllowed: true,
					StatementResult: StatementResult{
						Match:       true,
						StatementID: "456",
						Effect:      Effect(ALLOW),
						Resource:    "res1",
						Action:      "read",
					},
				},
			},
			`{
   "policy_id": "123",
   "is_allowed": true,
   "statement_result": {
      "match": true,
      "statement_id": "456",
      "effect": "Allow",
      "resource": "res1",
      "action": "read"
   }
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ValidationResult{
				PolicyResult: tt.fields.PolicyResult,
			}
			if got := v.PrettyJSON(); got != tt.want {
				t.Errorf("\nValidationResult.PrettyJSON() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}
