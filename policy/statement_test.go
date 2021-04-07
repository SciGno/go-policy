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
		conditions map[string]map[string]string
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

	registry := NewRegistry(
		map[string]ValidatorMap{
			"AfterTime": NewValidatorMap(map[string]Validator{"datetime": &AfterTime{}, "StringMatch": &StringMatch{}}),
		},
	)

	type fields struct {
		StatementID string
		Effect      Effect
		Resource    string
		Action      []string
		Condition   map[string]map[string]string
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
		//////////////////////////////////////
		// testing resources
		//////////////////////////////////////
		{
			"AllowStetement_ResourceNotExists",
			fields{"", ALLOW, "res1", []string{"read"}, nil},
			args{
				&Request{
					Resource: Input{
						Value:     "res2",
						Validator: "default",
					},
				},
				&registry,
			},
			StatementResult{
				Match:       false,
				StatementID: "",
				Processed:   Processed(RESOURCE),
				Effect:      Effect(ALLOW),
				Resource:    "res2",
			},
		},
		{
			"DenyStetement_ResourceNotExists",
			fields{"", DENY, "res1", []string{"read"}, nil},
			args{&Request{Resource: Input{Value: "res2", Validator: "default"}}, &registry},
			StatementResult{
				Match:       true,
				StatementID: "",
				Processed:   Processed(RESOURCE),
				Effect:      Effect(DENY),
				Resource:    "res2",
			},
		},
		//////////////////////////////////////
		// testing actions without conditions
		//////////////////////////////////////
		{
			"AllowStetement_ActionSucceedes",
			fields{"", ALLOW, "res1", []string{"read"}, nil},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "read", Validator: "default"},
				},
				&registry,
			},
			StatementResult{
				Match:       true,
				StatementID: "",
				Processed:   Processed(CONDITION),
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
			},
		},
		{
			"AllowStetement_ActionFails",
			fields{"", ALLOW, "res1", []string{"read"}, nil},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "write", Validator: "default"},
				},
				&registry,
			},
			StatementResult{
				Match:       false,
				StatementID: "",
				Processed:   Processed(ACTION),
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "write",
			},
		},
		{
			"DenyStetement_ActionIsDenied",
			fields{"", DENY, "res1", []string{"read"}, nil},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "read", Validator: "default"},
				},
				&registry,
			},
			StatementResult{
				Match:       true,
				StatementID: "",
				Processed:   Processed(ACTION),
				Effect:      Effect(DENY),
				Resource:    "res1",
				Action:      "read",
			},
		},
		{
			"DenyStetement_ActionNotInDenyStatement",
			fields{"", DENY, "res1", []string{"read"}, nil},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "write", Validator: "default"},
				},
				&registry,
			},
			StatementResult{
				Match:       false,
				StatementID: "",
				Processed:   Processed(CONDITION),
				Effect:      Effect(DENY),
				Resource:    "res1",
				Action:      "write",
			},
		},

		//////////////////////////////////
		// testing actions with conditions
		//////////////////////////////////
		{
			"StatementHasCondition_RequestHasCondition",
			fields{
				StatementID: "",
				Effect:      ALLOW,
				Resource:    "res1",
				Action:      []string{"read"},
				Condition: map[string]map[string]string{
					"AfterTime": {
						"datetime": "12:00",
					},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "read", Validator: "default"},
					Metadata: map[string]string{
						"datetime": "13:00",
					},
				},
				&registry,
			},
			StatementResult{
				Match:       true,
				StatementID: "",
				Processed:   Processed(CONDITION),
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
				Condition: map[string]string{
					"datetime": "13:00",
				},
			},
		},
		{
			"StatementHasCondition_RequestDoesNotHaveCondition",
			fields{
				StatementID: "",
				Effect:      ALLOW,
				Resource:    "res1",
				Action:      []string{"read"},
				Condition: map[string]map[string]string{
					"AfterTime": {
						"datetime": "12:00",
					},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "read", Validator: "default"},
				},
				&registry,
			},
			StatementResult{
				Match:       false,
				StatementID: "",
				Processed:   Processed(CONDITION),
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
			},
		},
		{
			"AllowStetement_ActionSucceedesWithCondtion",
			fields{
				StatementID: "",
				Effect:      ALLOW,
				Resource:    "res1",
				Action:      []string{"read"},
				Condition: map[string]map[string]string{
					"AfterTime": {
						"datetime": "12:00",
					},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "read", Validator: "default"},
					Metadata: map[string]string{
						"datetime": "13:00",
					},
				},
				&registry,
			},
			StatementResult{
				Match:       true,
				StatementID: "",
				Processed:   Processed(CONDITION),
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
				Condition: map[string]string{
					"datetime": "13:00",
				},
			},
		},
		{
			"AllowStetement_ActionFailsOnCondtion",
			fields{
				StatementID: "",
				Effect:      ALLOW,
				Resource:    "res1",
				Action:      []string{"read"},
				Condition: map[string]map[string]string{
					"AfterTime": {
						"datetime": "12:00",
					},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "read", Validator: "default"},
					Metadata: map[string]string{
						"datetime": "11:00",
					},
				},
				&registry,
			},
			StatementResult{
				Match:       false,
				StatementID: "",
				Processed:   Processed(CONDITION),
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
				Condition: map[string]string{
					"datetime": "11:00",
				},
			},
		},
		{
			"AllowStetement_ConditionNotRegistered",
			fields{
				StatementID: "",
				Effect:      ALLOW,
				Resource:    "res1",
				Action:      []string{"read"},
				Condition: map[string]map[string]string{
					"Other": {
						"datetime": "12:00",
					},
				},
			},
			args{
				&Request{
					Resource: Input{Value: "res1", Validator: "default"},
					Action:   Input{Value: "read", Validator: "default"},
					Metadata: map[string]string{
						"datetime": "13:00",
					},
				},
				&registry,
			},
			StatementResult{
				Match:       false,
				StatementID: "",
				Processed:   Processed(CONDITION),
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      "read",
				Condition: map[string]string{
					"datetime": "13:00",
				},
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
				t.Errorf("Statement.Validate() = %v, want %v", got, tt.want)
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
		Condition   map[string]map[string]string
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

func TestStatement_JSON(t *testing.T) {
	type fields struct {
		StatementID string
		Effect      Effect
		Resource    string
		Action      []string
		Condition   map[string]map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"TestJSON",
			fields{
				StatementID: "Read Only",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      []string{"read"},
			},
			"{\"sid\":\"Read Only\",\"effect\":\"Allow\",\"resource\":\"res1\",\"action\":[\"read\"]}",
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
			if got := s.JSON(); got != tt.want {
				t.Errorf("\nStatement.JSON() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}

func TestStatement_PrettyJSON(t *testing.T) {
	type fields struct {
		StatementID string
		Effect      Effect
		Resource    string
		Action      []string
		Condition   map[string]map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"TestJSON",
			fields{
				StatementID: "Read Only",
				Effect:      Effect(ALLOW),
				Resource:    "res1",
				Action:      []string{"read"},
				Condition: map[string]map[string]string{
					"AfterTime": {
						"datetime": "12:00",
					},
				},
			},
			`{
   "sid": "Read Only",
   "effect": "Allow",
   "resource": "res1",
   "action": [
      "read"
   ],
   "condition": {
      "AfterTime": {
         "datetime": "12:00"
      }
   }
}`,
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
			if got := s.PrettyJSON(); got != tt.want {
				t.Errorf("Statement.PrettyJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
