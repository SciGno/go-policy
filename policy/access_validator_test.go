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
						Location:    StatementLocation(ALL),
						StatementID: "",
						Effect:      Effect(ALLOW),
						Resource:    "res1",
						Action:      "read",
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
				t.Errorf("AccessValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
