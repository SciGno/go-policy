package policy

import (
	"fmt"
	"testing"
)

func TestDelimitedValidator_Validate(t *testing.T) {

	type args struct {
		stmntData   interface{}
		requestData interface{}
	}
	tests := []struct {
		name string
		sa   *DelimitedValidator
		args args
		want bool
	}{
		{"EmptyValues", &DelimitedValidator{}, args{}, false},
		{"ExactMatchSingleValue", &DelimitedValidator{}, args{"a", "a"}, true},
		{"MultipleWildcards", &DelimitedValidator{}, args{"a:*:c:*:e", "a:x:c:x:e"}, true},
		{"OmittedValuesMultipleWildcards", &DelimitedValidator{}, args{"a:*:c:*:e", "a::c::e"}, true},
		{"ExactMatch", &DelimitedValidator{}, args{"a:b", "a:b"}, true},
		{"NotMatch", &DelimitedValidator{}, args{"a:*", "x:b"}, false},
		{"MatchWildcard", &DelimitedValidator{}, args{"a:*", "a:b:c"}, true},
		{"RequestTooBig", &DelimitedValidator{}, args{"a:b", "a:b:c"}, false},
		{"RequestTooShort", &DelimitedValidator{}, args{"a:b:c:d", "a:b:c"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sa.Validate(tt.args.stmntData, tt.args.requestData); got != tt.want {
				t.Errorf("DelimitedValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}

}

func BenchmarkDelimitedValidator_Validate(b *testing.B) {
	ds := DelimitedValidator{}

	data := [][]string{
		{"a", "b"},
		{"a:b", "a:b"},
		{"a:*", "a:b"},
		{"a:*", "x:b"},
		{"a:b", "a:b:c"},
		{"a:*", "a:b:c:d:e"},
		{"a:b:c:d", "a:b:c"},
	}

	for i, v := range data {
		b.Run(fmt.Sprintf("%s %d", v, i), func(b *testing.B) {
			for x := 0; x < b.N; x++ {
				ds.Validate(v[0], v[0])
			}
		})
	}
}

func TestActionValidator_Validate(t *testing.T) {
	type args struct {
		statementData interface{}
		requestData   interface{}
	}
	tests := []struct {
		name string
		a    *ActionValidator
		args args
		want bool
	}{
		{"MatchesFirstElement", &ActionValidator{}, args{[]string{"a", "b"}, "a"}, true},
		{"MatchesWildcard", &ActionValidator{}, args{[]string{"*"}, "a"}, true},
		{"NoMatchingElements", &ActionValidator{}, args{[]string{"a", "b"}, "c"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ActionValidator{}
			if got := a.Validate(tt.args.statementData, tt.args.requestData); got != tt.want {
				t.Errorf("ActionValidator.Validate() = %v, want %v", got, tt.want)
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
		{"TestPasses", &StringMatch{}, args{"a", "a"}, true},
		{"TestFails", &StringMatch{}, args{"a", "b"}, false},
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

	cidr := "10.10.10.0/24"
	ip := "10.10.10.1"
	ip2 := "10.20.10.1"

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
		{"TestCIDR", &CIDR{}, args{[]string{cidr}, ip}, true},
		{"TestIP", &CIDR{}, args{[]string{ip}, ip}, true},
		{"CIDRTails", &CIDR{}, args{[]string{cidr}, ip2}, false},
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
