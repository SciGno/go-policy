package policy

import "testing"

func TestServiceWithAction_Validate(t *testing.T) {
	type args struct {
		stmntData   interface{}
		requestData interface{}
	}
	tests := []struct {
		name string
		sa   *ServiceWithAction
		args args
		want bool
	}{
		{"TestPasses", &ServiceWithAction{}, args{[]string{"a:*"}, "a:b"}, true},
		{"TestFails", &ServiceWithAction{}, args{[]string{"a:*"}, "x:b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := &ServiceWithAction{}
			if got := sa.Validate(tt.args.stmntData, tt.args.requestData); got != tt.want {
				t.Errorf("ServiceWithAction.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringMatch_Validate(t *testing.T) {
	type args struct {
		stmntData   interface{}
		requestData interface{}
	}
	tests := []struct {
		name string
		a    *StringMatch
		args args
		want bool
	}{
		{"TestPasses", &StringMatch{}, args{[]string{"a"}, "a"}, true},
		{"TestFails", &StringMatch{}, args{[]string{"a"}, "b"}, false},
		{"TestFails", &StringMatch{}, args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &StringMatch{}
			if got := a.Validate(tt.args.stmntData, tt.args.requestData); got != tt.want {
				t.Errorf("StringMatch.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCIDR_Validate(t *testing.T) {
	type args struct {
		stmntData   interface{}
		requestData interface{}
	}
	tests := []struct {
		name string
		c    *CIDR
		args args
		want bool
	}{
		{"TestCIDR", &CIDR{}, args{[]string{"10.10.10.0/24"}, "10.10.10.1"}, true},
		{"TestIP", &CIDR{}, args{[]string{"10.10.10.1"}, "10.10.10.1"}, true},
		{"CIDRTails", &CIDR{}, args{[]string{"10.10.10.0/24"}, "10.20.10.1"}, false},
		{"TestFails", &CIDR{}, args{[]string{"a"}, "a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CIDR{}
			if got := c.Validate(tt.args.stmntData, tt.args.requestData); got != tt.want {
				t.Errorf("CIDR.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfterTime_Validate(t *testing.T) {
	type args struct {
		stmntData   interface{}
		requestData interface{}
	}
	tests := []struct {
		name string
		a    *AfterTime
		args args
		want bool
	}{
		{"TestPasses", &AfterTime{}, args{"12:00", "13:00"}, true},
		{"TestNil", &AfterTime{}, args{"00:00", nil}, true},
		{"TestNilStmnt", &AfterTime{}, args{nil, "12:00"}, false},
		{"TestAllNill", &AfterTime{}, args{nil, nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AfterTime{}
			if got := a.Validate(tt.args.stmntData, tt.args.requestData); got != tt.want {
				t.Errorf("AfterTime.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBeforeTime_Validate(t *testing.T) {
	type args struct {
		stmntData   interface{}
		requestData interface{}
	}
	tests := []struct {
		name string
		a    *BeforeTime
		args args
		want bool
	}{
		{"TestPasses", &BeforeTime{}, args{"13:00", "12:00"}, true},
		{"TestNil", &BeforeTime{}, args{"23:59", nil}, true},
		{"TestNilStmnt", &BeforeTime{}, args{nil, "12:00"}, false},
		{"TestAllNill", &BeforeTime{}, args{nil, nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &BeforeTime{}
			if got := a.Validate(tt.args.stmntData, tt.args.requestData); got != tt.want {
				t.Errorf("BeforeTime.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeRanges_Validate(t *testing.T) {
	type args struct {
		stmntData   interface{}
		requestData interface{}
	}
	tests := []struct {
		name string
		a    *TimeRanges
		args args
		want bool
	}{
		{"TestPasses", &TimeRanges{}, args{map[string]interface{}{"from": "12:00", "to": "23:59"}, "15:00"}, true},
		{"TestNilReq", &TimeRanges{}, args{map[string]interface{}{"from": "12:00", "to": "23:59"}, nil}, false},
		{"TestEmptyMap", &TimeRanges{}, args{map[string]interface{}{}, nil}, false},
		{"TestNilBoth", &TimeRanges{}, args{nil, nil}, false},
		{"TestWrongType", &TimeRanges{}, args{[]string{"12:00", "23:59"}, nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &TimeRanges{}
			if got := a.Validate(tt.args.stmntData, tt.args.requestData); got != tt.want {
				t.Errorf("TimeRanges.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
