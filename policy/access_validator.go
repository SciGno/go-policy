package policy

// AccessValidator validates all policies against a given request
type AccessValidator struct {
	registry Registry
}

// Validate checks the request against all statements. Effect=Deny statements will be validated first
// followed by Effect=Allow statements
func (v *AccessValidator) Validate(request *Request, policies []Policy) bool {

	for _, policy := range policies {
		if !policy.ValidateDeny(request, &v.registry) {
			return false
		}
	}
	for _, policy := range policies {
		if policy.ValidateAllow(request, &v.registry) {
			return true
		}
	}

	return false
}
