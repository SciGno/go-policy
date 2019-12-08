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
	for _, act := range stmntData.([]string) {
		aTmp := strings.Split(act, ":")
		for _, v := range requestData.([]string) {
			rTmp := strings.Split(v, ":")
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
	for _, s := range stmntData.([]string) {
		for _, v := range requestData.([]string) {
			if s == v {
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
	for _, v := range stmntData.([]interface{}) {
		if _, ipnetA, _ := net.ParseCIDR(v.(string)); ipnetA != nil { // check for CIDR notation
			if requestData != nil {
				ipB := net.ParseIP(requestData.(string))
				if ipnetA.Contains(ipB) {
					return true
				}
			}
		} else if ipA := net.ParseIP(v.(string)); ipA != nil { // check for IP
			if requestData != nil {
				ipB := net.ParseIP(requestData.(string))
				return ipA.Equal(ipB)
			}
		}
	}
	return false
}

// AfterTime returns true if time1 happens after time2.
type AfterTime struct{}

// Validate function
func (a *AfterTime) Validate(stmntData, requestData interface{}) bool {
	if requestData != nil {
		if t1, err := time.Parse("03:04", stmntData.(string)); err != nil {
			if t2, err := time.Parse("03:04", requestData.(string)); err != nil {
				if t1.After(t2) {
					return true
				}
			}
		}
	} else if t1, err := time.Parse("03:04", stmntData.(string)); err != nil {
		t2, _ := time.Parse("03:04", time.Now().Format("03:04"))
		if t1.After(t2) {
			return true
		}
	}
	return false
}

// BeforeTime returns true if time1 happens before time2.
type BeforeTime struct{}

// Validate function
func (a *BeforeTime) Validate(stmntData, requestData interface{}) bool {
	if requestData != nil {
		if t1, err := time.Parse("03:04", stmntData.(string)); err != nil {
			if t2, err := time.Parse("03:04", requestData.(string)); err != nil {
				if t1.Before(t2) {
					return true
				}
			}
		}
	} else if t1, err := time.Parse("03:04", stmntData.(string)); err != nil {
		t2, _ := time.Parse("03:04", time.Now().Format("03:04"))
		if t1.Before(t2) {
			return true
		}
	}
	return false
}
