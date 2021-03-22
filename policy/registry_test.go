package policy

import (
	"reflect"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	type args struct {
		resourceValidator Validator
		actionValidator   Validator
		conditionMap      map[string]Validator
	}
	tests := []struct {
		name string
		args args
		want Registry
	}{
		{
			"",
			args{&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}}},
			NewRegistry(&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegistry(tt.args.resourceValidator, tt.args.actionValidator, tt.args.conditionMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_GetResourceValidator(t *testing.T) {

	rv := DelimitedValidator{}

	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	tests := []struct {
		name   string
		fields fields
		want   Validator
	}{
		{
			"AfterTime",
			fields{&rv, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}}},
			&rv,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			if got := r.GetResourceValidator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registry.GetResourceValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_GetActionValidator(t *testing.T) {
	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	tests := []struct {
		name   string
		fields fields
		want   Validator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			if got := r.GetActionValidator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registry.GetActionValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_GetConditionValidator(t *testing.T) {
	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Validator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			if got := r.GetConditionValidator(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registry.GetConditionValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_GetConditionValidators(t *testing.T) {

	cv := map[string]Validator{"AfterTime": &AfterTime{}}

	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]Validator
	}{
		{
			"",
			fields{&DelimitedValidator{}, &ActionValidator{}, cv},
			cv,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			if got := r.GetConditionValidators(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registry.GetConditionValidators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_AddConditionValidator(t *testing.T) {
	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	type args struct {
		name      string
		validator Validator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			r.AddConditionValidator(tt.args.name, tt.args.validator)
		})
	}
}

func TestRegistry_RemoveConditionValidator(t *testing.T) {
	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			r.RemoveConditionValidator(tt.args.name)
		})
	}
}
