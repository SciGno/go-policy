package policy

import "testing"

func TestAccessValidator_Validate(t *testing.T) {

	registry := NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}})
	request := NewRequest("a", "b", nil)

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
		want   bool
	}{
		{
			"Pass",
			fields{registry},
			args{
				&request,
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, []string{"a"}, "b", map[string]interface{}{"AfterTime": "00:00"})}),
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, []string{"x"}, "b", map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			true,
		},
		{
			"Pass",
			fields{registry},
			args{
				&request,
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, []string{"a"}, "b", map[string]interface{}{"AfterTime": "00:00"})}),
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, []string{"a"}, "b", map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			false,
		},
		{
			"Pass",
			fields{registry},
			args{
				&request,
				[]Policy{
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", ALLOW, []string{"c"}, "b", map[string]interface{}{"AfterTime": "00:00"})}),
					NewPolicy("123", "PolicyName", "1.0", []Statement{NewStatement("", DENY, []string{"e"}, "b", map[string]interface{}{"AfterTime": "23:59"})}),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &AccessValidator{
				registry: tt.fields.registry,
			}
			if got := v.Validate(tt.args.request, tt.args.policies); got != tt.want {
				t.Errorf("AccessValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
