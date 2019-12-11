package policy

import (
	"reflect"
	"testing"
)

func TestNewValidationError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *ValidationError
	}{
		{"Test1", args{"message"}, NewValidationError("message")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidationError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidationError_Error(t *testing.T) {
	type fields struct {
		PolicyID  string
		Statement string
		Location  string
		Validator string
		Data      interface{}
		message   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"TestErrorValue", fields{"123", "a", "b", "c", "data", "message"}, "message"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ValidationError{
				PolicyID:  tt.fields.PolicyID,
				Statement: tt.fields.Statement,
				Location:  tt.fields.Location,
				Validator: tt.fields.Validator,
				Data:      tt.fields.Data,
				message:   tt.fields.message,
			}
			if got := v.Error(); got != tt.want {
				t.Errorf("ValidationError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
