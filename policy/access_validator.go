package policy

import (
	"encoding/json"
)

// ValidationResult struct
type ValidationResult struct {
	PolicyResult
}

func (v *ValidationResult) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}

func (v *ValidationResult) PrettyJSON() string {
	data, _ := json.MarshalIndent(v, "", "   ")
	return string(data)
}

// AccessValidator validates all policies against a given request
type AccessValidator struct {
	registry Registry
}

// Validate checks the request against all statements. Effect=Deny statements will be validated first
// followed by Effect=Allow statements
func (v *AccessValidator) Validate(request *Request, policies []Policy) ValidationResult {

	for _, policy := range policies {
		if pr := policy.ValidateDeny(request, &v.registry); !pr.IsAllowed {
			return ValidationResult{pr}
		}
	}
	for _, policy := range policies {
		if pr := policy.ValidateAllow(request, &v.registry); pr.IsAllowed {
			return ValidationResult{pr}
		}
	}

	return ValidationResult{PolicyResult{IsAllowed: false}}

}
