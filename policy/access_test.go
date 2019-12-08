package policy

import "testing"

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
		{"TestIP", fields{NewRegistry(nil, nil, map[string]Validator{"SourceIP": &CIDR{}})}, args{"SourceIP", []interface{}{"10.10.20.12"}, &r1}, false},
		{"IPFails", fields{NewRegistry(nil, nil, map[string]Validator{"SourceIP": &CIDR{}})}, args{"SourceIP", []interface{}{"10.10.20.123"}, &r1}, true},
		{"TestPassNill", fields{NewRegistry(nil, nil, map[string]Validator{"AfterTime": &AfterTime{}})}, args{"AfterTime", "23:59", &r1}, false},
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
