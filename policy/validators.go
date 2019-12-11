package policy

import (
	"net"
	"strings"
	"time"
)

// ServiceWithAction first check for a match on the service portion of the string
// then on the action portion of the string.  It will only validate aginst
// an asterisk ( * ) or if the actions in the request are the exactly the same.
type ServiceWithAction struct{}

// Validate function
func (sa *ServiceWithAction) Validate(stmntData, requestData interface{}) bool {
	if requestData != nil || stmntData != nil {
		for _, act := range stmntData.([]string) {
			aTmp := strings.Split(act, ":")
			rTmp := strings.Split(requestData.(string), ":")
			if aTmp[0] == "*" || aTmp[0] == rTmp[0] {
				if aTmp[1] == "*" || aTmp[1] == rTmp[1] {
					return true
				}
			}
		}
	}
	return false
}

// StringMatch returns true if any value in stmntData equals any value in requestData
type StringMatch struct{}

// Validate function
func (a *StringMatch) Validate(stmntData, requestData interface{}) bool {
	if requestData != nil || stmntData != nil {
		for _, s := range stmntData.([]string) {
			if s == requestData {
				return true
			}
		}
	}
	return false
}

// CIDR struct
type CIDR struct{}

// Validate expects an IP address and checks for membership
func (c *CIDR) Validate(stmntData, requestData interface{}) bool {
	if requestData != nil || stmntData != nil {
		for _, v := range stmntData.([]string) {
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
func (a *AfterTime) Validate(stmntData, requestData interface{}) bool {
	if stmntData != nil && requestData != nil {
		if t1, err := time.Parse("15:04", stmntData.(string)); err == nil {
			if t2, err := time.Parse("15:04", requestData.(string)); err == nil {
				if t1.UTC().Before(t2.UTC()) {
					return true
				}
			}
		}
	} else if stmntData != nil {
		if t1, err := time.Parse("15:04", stmntData.(string)); err == nil {
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
func (a *BeforeTime) Validate(stmntData, requestData interface{}) bool {
	if stmntData != nil && requestData != nil {
		if t1, err := time.Parse("15:04", stmntData.(string)); err == nil {
			if t2, err := time.Parse("15:04", requestData.(string)); err == nil {
				if t1.UTC().After(t2.UTC()) {
					return true
				}
			}
		}
	} else if stmntData != nil {
		if t1, err := time.Parse("15:04", stmntData.(string)); err == nil {
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
func (a *TimeRanges) Validate(stmntData, requestData interface{}) bool {
	if requestData != nil && stmntData != nil {
		switch stmntTimes := stmntData.(type) {
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
