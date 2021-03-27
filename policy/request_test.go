package policy

import (
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	type args struct {
		resource   string
		action     string
		conditions map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Request
	}{
		{
			"CreateRequestWithCondition",
			args{"res1", "read", map[string]interface{}{"AfterTime": "12:00"}},
			Request{"res1", "read", map[string]interface{}{"AfterTime": "12:00"}},
		},
		{
			"CreateRequestWithoutCondition",
			args{"res1", "read", nil},
			Request{"res1", "read", nil},
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
