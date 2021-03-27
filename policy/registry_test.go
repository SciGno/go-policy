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
			fields{&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}}},
			&DelimitedValidator{},
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
			"GetConditionValidator",
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

func TestRegistry_SetResourceValidator(t *testing.T) {

	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	type args struct {
		rv Validator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"SetResourceValidator",
			fields{&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}}},
			args{&ActionValidator{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			r.SetResourceValidator(tt.args.rv)
			if got := r.GetResourceValidator(); !reflect.DeepEqual(got, tt.args.rv) {
				t.Errorf("Registry.GetResourceValidator() = %v, want %v", got, tt.args.rv)
			}
		})
	}
}

func TestRegistry_SetActionValidator(t *testing.T) {
	type fields struct {
		resource  Validator
		action    Validator
		condition map[string]Validator
	}
	type args struct {
		av Validator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"SetActionValidator",
			fields{&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}}},
			args{&DelimitedValidator{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			r.SetActionValidator(tt.args.av)
			if got := r.GetActionValidator(); !reflect.DeepEqual(got, tt.args.av) {
				t.Errorf("Registry.SetActionValidator() = %v, want %v", got, tt.args.av)
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
		{
			"AddConditionValidator",
			fields{&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}}},
			args{"BeforeTime", &BeforeTime{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			r.AddConditionValidator(tt.args.name, tt.args.validator)
			if got := r.GetConditionValidator(tt.args.name); !reflect.DeepEqual(got, tt.args.validator) {
				t.Errorf("Registry.AddConditionValidator() = %v, want %v", got, tt.args.validator)
			}
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
		{
			"AddConditionValidator",
			fields{&DelimitedValidator{}, &ActionValidator{}, map[string]Validator{"AfterTime": &AfterTime{}, "BeforeTime": &BeforeTime{}}},
			args{"BeforeTime"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				resource:  tt.fields.resource,
				action:    tt.fields.action,
				condition: tt.fields.condition,
			}
			r.RemoveConditionValidator(tt.args.name)
			if got := r.GetConditionValidator(tt.args.name); got != nil {
				t.Errorf("Registry.AddConditionValidator() = %v, want %v", got, nil)
			}
		})
	}
}
