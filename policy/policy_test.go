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
			args{"123", "Test Policy", "1.0", []Statement{{"Read Only", ALLOW, "res1:a", []string{"read"}, nil}}},
			Policy{
				PolicyID: "123",
				Name:     "Test Policy",
				Version:  "1.0",
				Statement: []Statement{
					{"Read Only", ALLOW, "res1:a", []string{"read"}, nil}},
			},
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

	registry := NewRegistry(
		map[string]ValidatorMap{
			"AfterTime": NewValidatorMap(map[string]Validator{"datetime": &AfterTime{}, "StringMatch": &StringMatch{}}),
		},
	)

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
			"AllowAny",
			fields{
				"123",
				"Test Policy",
				"1.0",
				[]Statement{
					{"", ALLOW, "*", []string{"*"}, nil},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1"},
					Action:   Input{Value: "anyaction"},
					Metadata: map[string]string{
						"datetime": "13:00",
					},
				},
				&registry,
			},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: true,
				StatementResult: StatementResult{
					Match:       true,
					StatementID: "",
					Processed:   Processed(CONDITION),
					Effect:      Effect(ALLOW),
					Resource:    "res1",
					Action:      "anyaction",
					Condition: map[string]string{
						"datetime": "13:00",
					},
				},
			},
		},
		{
			"DenyReadOnAnyResource",
			fields{
				PolicyID: "123",
				Name:     "Test Policy",
				Version:  "1.0",
				Statement: []Statement{
					{"Allow Any", ALLOW, "*", []string{"*"}, nil},
					{"Deny Read", DENY, "*", []string{"read"}, nil},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1"},
					Action:   Input{Value: "read"},
					Metadata: map[string]string{
						"datetime": "13:00",
					},
				},
				&registry,
			},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: false,
				StatementResult: StatementResult{
					Match:       true,
					StatementID: "Deny Read",
					Processed:   Processed(ACTION),
					Effect:      Effect(DENY),
					Resource:    "res1",
					Action:      "read",
				},
			},
		},
		{
			"PassesAllowAnyAfter12:00",
			fields{
				"123",
				"Test Policy",
				"1.0",
				[]Statement{
					{
						StatementID: "Only after 12:00",
						Effect:      ALLOW,
						Resource:    "res1",
						Action:      []string{"read"},
						Condition: map[string]map[string]string{
							"AfterTime": {
								"datetime": "12:00",
							},
						},
					},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1"},
					Action:   Input{Value: "read"},
					Metadata: map[string]string{
						"datetime": "13:00",
					},
				},
				&registry,
			},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: true,
				StatementResult: StatementResult{
					Match:       true,
					StatementID: "Only after 12:00",
					Processed:   Processed(CONDITION),
					Effect:      Effect(ALLOW),
					Resource:    "res1",
					Action:      "read",
					Condition: map[string]string{
						"datetime": "13:00",
					},
				},
			},
		},
		{
			"FailsAllowAnyAfter12:00",
			fields{
				"123",
				"Test Policy",
				"1.0",
				[]Statement{
					{
						StatementID: "Only after 12:00",
						Effect:      ALLOW,
						Resource:    "res1",
						Action:      []string{"read"},
						Condition: map[string]map[string]string{
							"AfterTime": {
								"datetime": "12:00",
							},
						},
					},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1"},
					Action:   Input{Value: "read"},
					Metadata: map[string]string{
						"datetime": "11:00",
					},
				},
				&registry,
			},
			PolicyResult{
				PolicyID:  "123",
				IsAllowed: false,
				StatementResult: StatementResult{
					Match:       false,
					StatementID: "Only after 12:00",
					Processed:   Processed(CONDITION),
					Effect:      Effect(ALLOW),
					Resource:    "res1",
					Action:      "read",
					Condition: map[string]string{
						"datetime": "11:00",
					},
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
				t.Errorf("\nPolicy.Validate() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}
