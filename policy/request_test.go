package policy

import (
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	type args struct {
		resource   Input
		action     Input
		conditions map[string]string
	}
	tests := []struct {
		name string
		args args
		want Request
	}{
		{
			"New",
			args{
				Input{Value: "res1", Validator: "default"},
				Input{Value: "read", Validator: "default"},
				nil,
			},
			Request{
				Input{Value: "res1", Validator: "default"},
				Input{Value: "read", Validator: "default"},
				nil,
			},
		},
		{
			"NewWithoutValidatorName",
			args{
				Input{Value: "res1"},
				Input{Value: "read"},
				nil,
			},
			Request{
				Input{Value: "res1", Validator: "default"},
				Input{Value: "read", Validator: "default"},
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequest(tt.args.resource, tt.args.action, tt.args.conditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_MetaNameExists(t *testing.T) {
	type fields struct {
		Resource Input
		Action   Input
		Metadata map[string]string
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
			"MetaNameExists",
			fields{
				Input{Value: "res1", Validator: "default"},
				Input{Value: "read", Validator: "default"},
				map[string]string{
					"datetime": "13:00",
				},
			},
			args{"datetime"},
			true,
		},
		{
			"MetaNameDoesNoExist",
			fields{
				Input{Value: "res1", Validator: "default"},
				Input{Value: "read", Validator: "default"},
				map[string]string{
					"datetime": "13:00",
				},
			},
			args{"other"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Resource: tt.fields.Resource,
				Action:   tt.fields.Action,
				Metadata: tt.fields.Metadata,
			}
			if got := r.MetaNameExists(tt.args.metaName); got != tt.want {
				t.Errorf("Request.MetaNameExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
