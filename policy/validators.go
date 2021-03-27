package policy

import (
	"net"
	"strings"
	"time"
)

// DelimitedValidator returns true if the delimited values are the exactly the same
// and treats an asterisk ( * ) as a wildcard. Default delimiter is a colon ':'
type DelimitedValidator struct{}

// Validate function
func (d *DelimitedValidator) Validate(statementData, requestData interface{}) bool {

	if requestData != nil || statementData != nil {
		statementDataArray := strings.Split(statementData.(string), ":")
		requestDataArray := strings.Split(requestData.(string), ":")

		if len(requestDataArray) < len(statementDataArray) {
			return false
		}

		for i, v := range requestDataArray {
			if i > len(statementDataArray)-1 {
				if statementDataArray[i-1] != "*" {
					return false
				}
				return true
			}

			if v != statementDataArray[i] && statementDataArray[i] != "*" {
				return false
			}
		}
		return true
	}
	return false
}

// ActionValidator returns true statementData equals requestData
// and if statementData is an asterisk '*'.
type ActionValidator struct{}

// Validate function
func (a *ActionValidator) Validate(statementData, requestData interface{}) bool {
	if requestData != nil || statementData != nil {
		for _, v := range statementData.([]string) {
			if v == requestData.(string) || v == "*" {
				return true
			}
		}
	}
	return false
}

// StringMatch returns true if any value in statementData equals any value in requestData
type StringMatch struct{}

// Validate function
func (a *StringMatch) Validate(statementData, requestData interface{}) bool {
	if requestData != nil || statementData != nil {
		if strings.Compare(statementData.(string), requestData.(string)) == 0 {
			return true
		}
	}
	return false
}

// CIDR struct
type CIDR struct{}

// Validate expects an IP address and checks for membership
func (c *CIDR) Validate(statementData, requestData interface{}) bool {
	if requestData != nil || statementData != nil {
		for _, v := range statementData.([]string) {
			if _, ipnetA, _ := net.ParseCIDR(v); ipnetA != nil { // check for CIDR notation
				if requestData != nil {
					ipB := net.ParseIP(requestData.(string))
					if ipnetA.Contains(ipB) {
						return true
					}
				}
			} else if ipA := net.ParseIP(v); ipA != nil { // check for IP
				if requestData != nil {
					ipB := net.ParseIP(requestData.(string))
					return ipA.Equal(ipB)
				}
			}
		}
	}
	return false
}

// AfterTime returns true if time1 happens after time2.
type AfterTime struct{}

// Validate function
func (a *AfterTime) Validate(statementData, requestData interface{}) bool {
	if statementData != nil && requestData != nil {
		if t1, err := time.Parse("15:04", statementData.(string)); err == nil {
			if t2, err := time.Parse("15:04", requestData.(string)); err == nil {
				if t1.UTC().Before(t2.UTC()) {
					return true
				}
			}
		}
	} else if statementData != nil {
		if t1, err := time.Parse("15:04", statementData.(string)); err == nil {
			t2, _ := time.Parse("15:04", time.Now().Format("15:04"))
			if t1.UTC().Before(t2.UTC()) {
				return true
			}
		}
	}
	return false
}

// BeforeTime returns true if time1 happens before time2.
type BeforeTime struct{}

// Validate function
func (a *BeforeTime) Validate(statementData, requestData interface{}) bool {
	if statementData != nil && requestData != nil {
		if t1, err := time.Parse("15:04", statementData.(string)); err == nil {
			if t2, err := time.Parse("15:04", requestData.(string)); err == nil {
				if t1.UTC().After(t2.UTC()) {
					return true
				}
			}
		}
	} else if statementData != nil {
		if t1, err := time.Parse("15:04", statementData.(string)); err == nil {
			t2, _ := time.Parse("15:04", time.Now().Format("15:04"))
			if t1.UTC().After(t2.UTC()) {
				return true
			}
		}
	}
	return false
}

// TimeRanges returns true if time1 is within the range of from and to values.
// The from and to values are inclusive.
type TimeRanges struct{}

// Validate function
func (a *TimeRanges) Validate(statementData, requestData interface{}) bool {
	if requestData != nil && statementData != nil {
		switch stmntTimes := statementData.(type) {
		case map[string]interface{}:
			if requestTime, err := time.Parse("15:04", requestData.(string)); err == nil {
				if from, ok := stmntTimes["from"]; ok {
					if f, err := time.Parse("15:04", from.(string)); err == nil {
						if to, ok := stmntTimes["to"]; ok {
							if t, err := time.Parse("15:04", to.(string)); err == nil {
								if (requestTime.UTC().Equal(f) || requestTime.UTC().After(f)) && (requestTime.UTC().Equal(t) || requestTime.UTC().Before(t)) {
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}
