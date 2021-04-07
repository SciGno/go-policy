package policy

import (
	"reflect"
	"testing"
)

func TestValidatorMap_GetValidator(t *testing.T) {
	type fields struct {
		validators map[string]Validator
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
		{
			"GetStringMatchValidator",
			fields{
				map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}},
			},
			args{"StringMatch"},
			&StringMatch{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &ValidatorMap{
				validators: tt.fields.validators,
			}
			if got := vm.GetValidator(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidatorMap.GetValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatorMap_GetValidators(t *testing.T) {
	type fields struct {
		validators map[string]Validator
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]Validator
	}{
		{
			"GetAllValidators",
			fields{
				map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}},
			},
			map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &ValidatorMap{
				validators: tt.fields.validators,
			}
			if got := vm.GetValidators(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidatorMap.GetValidators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatorMap_RemoveValidator(t *testing.T) {
	type fields struct {
		validators map[string]Validator
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
			"GetAllValidators",
			fields{
				map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}},
			},
			args{"StringMatch"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &ValidatorMap{
				validators: tt.fields.validators,
			}
			vm.RemoveValidator(tt.args.name)
			if got := vm.GetValidator(tt.args.name); !reflect.DeepEqual(got, nil) {
				t.Errorf("ValidatorMap.GetValidator() = %v, want %v", got, tt.args.name)
			}
		})
	}
}

func TestValidatorMap_AddValidator(t *testing.T) {
	type fields struct {
		validators map[string]Validator
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
			"AddStringMatchValidator",
			fields{
				map[string]Validator{},
			},
			args{"StringMatch", &StringMatch{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &ValidatorMap{
				validators: tt.fields.validators,
			}
			vm.AddValidator(tt.args.name, tt.args.validator)
			if got := vm.GetValidator(tt.args.name); !reflect.DeepEqual(got, tt.args.validator) {
				t.Errorf("ValidatorMap.GetValidator() = %v, want %v", got, tt.args.name)
			}
		})
	}
}

func TestNewRegistry(t *testing.T) {
	type args struct {
		validatorMap map[string]ValidatorMap
	}
	tests := []struct {
		name string
		args args
		want Registry
	}{
		{
			"NewRegistry",
			args{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"DelimitedValidator": &DelimitedValidator{}, "StringMatch": &StringMatch{}}),
					"action":   NewValidatorMap(map[string]Validator{"ActionValidator": &ActionValidator{}, "StringMatch": &StringMatch{}}),
					"sourceIP": NewValidatorMap(map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}}),
				},
			},
			NewRegistry(
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"DelimitedValidator": &DelimitedValidator{}, "StringMatch": &StringMatch{}}),
					"action":   NewValidatorMap(map[string]Validator{"ActionValidator": &ActionValidator{}, "StringMatch": &StringMatch{}}),
					"sourceIP": NewValidatorMap(map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}}),
				},
			),
		},
		{
			"MissingResourceValidatorMap",
			args{
				map[string]ValidatorMap{
					"action": NewValidatorMap(map[string]Validator{"default": &ActionValidator{}}),
				},
			},
			Registry{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"default": &DelimitedValidator{}}),
					"action":   NewValidatorMap(map[string]Validator{"default": &ActionValidator{}}),
				},
			},
		},
		{
			"MissingActionValidatorMap",
			args{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"default": &DelimitedValidator{}}),
				},
			},
			Registry{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"default": &DelimitedValidator{}}),
					"action":   NewValidatorMap(map[string]Validator{"default": &ActionValidator{}}),
				},
			},
		},
		{
			"MissingMap",
			args{nil},
			Registry{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"default": &DelimitedValidator{}}),
					"action":   NewValidatorMap(map[string]Validator{"default": &ActionValidator{}}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegistry(tt.args.validatorMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_AddValidatorMap(t *testing.T) {
	type fields struct {
		maps map[string]ValidatorMap
	}
	type args struct {
		metaName     string
		validatorMap ValidatorMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"AddValidatorMap",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"DelimitedValidator": &DelimitedValidator{}}),
				},
			},
			args{
				"other",
				ValidatorMap{map[string]Validator{"default": &DelimitedValidator{}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				maps: tt.fields.maps,
			}
			r.AddValidatorMap(tt.args.metaName, tt.args.validatorMap)
			if got := r.GetValidatorMap(tt.args.metaName); !reflect.DeepEqual(got, tt.args.validatorMap) {
				t.Errorf("ValidatorMap.GetValidators() = %v, want %v", got, tt.fields.maps[tt.args.metaName])
			}
		})
	}
}

func TestRegistry_GetValidatorMap(t *testing.T) {
	type fields struct {
		maps map[string]ValidatorMap
	}
	type args struct {
		metaName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ValidatorMap
	}{
		{
			"GetsValidatorMap",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"DelimitedValidator": &DelimitedValidator{}, "StringMatch": &StringMatch{}}),
					"action":   NewValidatorMap(map[string]Validator{"ActionValidator": &ActionValidator{}, "StringMatch": &StringMatch{}}),
					"sourceIP": NewValidatorMap(map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}}),
				},
			},
			args{"resource"},
			NewValidatorMap(map[string]Validator{"DelimitedValidator": &DelimitedValidator{}, "StringMatch": &StringMatch{}}),
		},
		{
			"GetsEmptyValidatorMap",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"DelimitedValidator": &DelimitedValidator{}, "StringMatch": &StringMatch{}}),
					"action":   NewValidatorMap(map[string]Validator{"ActionValidator": &ActionValidator{}, "StringMatch": &StringMatch{}}),
					"sourceIP": NewValidatorMap(map[string]Validator{"CIDR": &CIDR{}, "StringMatch": &StringMatch{}}),
				},
			},
			args{"resources"},
			ValidatorMap{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				maps: tt.fields.maps,
			}
			if got := r.GetValidatorMap(tt.args.metaName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registry.GetValidatorMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_RemoveValidatorMap(t *testing.T) {
	type fields struct {
		maps map[string]ValidatorMap
	}
	type args struct {
		metaName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ValidatorMap
	}{
		{
			"RemovesValidatorMap",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"default": &DelimitedValidator{}}),
					"action":   NewValidatorMap(map[string]Validator{"ActionValidator": &ActionValidator{}, "StringMatch": &StringMatch{}}),
				},
			},
			args{"resource"},
			ValidatorMap{
				map[string]Validator{"default": &DelimitedValidator{}},
			},
		},
		{
			"RemovesValidatorMap",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"default": &DelimitedValidator{}}),
					"action":   NewValidatorMap(map[string]Validator{"ActionValidator": &ActionValidator{}, "StringMatch": &StringMatch{}}),
				},
			},
			args{"resources"},
			ValidatorMap{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				maps: tt.fields.maps,
			}
			if got := r.RemoveValidatorMap(tt.args.metaName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registry.RemoveValidatorMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_AddValidator(t *testing.T) {
	type fields struct {
		maps map[string]ValidatorMap
	}
	type args struct {
		metaName      string
		validatorName string
		validator     Validator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"AddValidator",
			fields{map[string]ValidatorMap{}},
			args{"resource", "match", &StringMatch{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				maps: tt.fields.maps,
			}
			r.AddValidator(tt.args.metaName, tt.args.validatorName, tt.args.validator)
			if got := r.GetValidator(tt.args.metaName, tt.args.validatorName); !reflect.DeepEqual(got, tt.args.validator) {
				t.Errorf("Registry.RemoveValidatorMap() = %v, want %v", got, tt.args.validator)
			}
		})
	}
}

func TestRegistry_RemoveValidator(t *testing.T) {
	type fields struct {
		maps map[string]ValidatorMap
	}
	type args struct {
		metaName      string
		validatorName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Validator
	}{
		{
			"RemoveValidator",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"ActionValidator": &ActionValidator{}, "match": &StringMatch{}}),
				},
			},
			args{"resource", "match"},
			&StringMatch{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				maps: tt.fields.maps,
			}
			if got := r.RemoveValidator(tt.args.metaName, tt.args.validatorName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registry.RemoveValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_MetaNameExists(t *testing.T) {
	type fields struct {
		maps map[string]ValidatorMap
	}
	type args struct {
		metaName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"MetaName_Exists",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"match": &StringMatch{}}),
				},
			},
			args{"resource"},
			true,
		},
		{
			"MetaName_Does_Not_Exist",
			fields{
				map[string]ValidatorMap{
					"resource": NewValidatorMap(map[string]Validator{"match": &StringMatch{}}),
				},
			},
			args{"resources"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registry{
				maps: tt.fields.maps,
			}
			if got := r.MetaNameExists(tt.args.metaName); got != tt.want {
				t.Errorf("Registry.MetaNameExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
