package policy

// LookupInput is used for retreiving the ValidatorMap with LookupName
// and the Validator with the ValidatorName.
type Input struct {
	Value     string `json:"value"`
	Validator string `json:"validator"`
}

/**
Request is used to for validating the policy statements agains it.

This is an example of a Request:

Request {
	Resource: {
		Value: "res1",
		Validator: "default"
	},
	Action: LookupImput {
		Value: "read",
		Validator: "default"
	},
	Condition: {
		"sourceIp": "10.24.0.23",
		"datetime": "03-28-2021 12:00PM UTC"
	},
}

NOTE:
The resource and the action in the request decide what validator will be used to validate them.
The condition validators are decided by the statement.

**/
type Request struct {
	Resource Input             `json:"resource"`
	Action   Input             `json:"action"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// NewRequest returns a new request
// By default, if resource and action validator names are omited,
// a "default" validator name will be used.
func NewRequest(resource, action Input, metadata map[string]string) Request {

	if len(resource.Validator) == 0 {
		resource.Validator = "default"
	}

	if len(action.Validator) == 0 {
		action.Validator = "default"
	}

	return Request{
		Action:   action,
		Resource: resource,
		Metadata: metadata,
	}
}

func (r *Request) MetaNameExists(metaName string) bool {
	if _, ok := r.Metadata[metaName]; ok {
		return true
	}
	return false
}
