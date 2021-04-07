package main

import (
	"encoding/json"
	"fmt"

	"github.com/scigno/go-policy/policy"
)

var pol = []byte(`
{
    "id": "f0dacf4e-c051-475e-b9dc-f86cd2d9539c",
    "name": "Demo Policy",
    "version": "1.0",
    "weight": "1",
    "statement": [
        {
            "sid": "AllActions",
            "effect": "Allow",
            "resource": "us:aws:graphql",
            "action": [
                "delete",
                "update"
            ],
            "condition": {
                "AfterTime": {
                    "datetime": "12:00"
                }
            }
        }
    ]
}
`)

// Chase Fraud Claims: 866-564-2262

func main() {

	p := policy.Policy{}

	if err := json.Unmarshal(pol, &p); err != nil {
		fmt.Printf("%s\n", err)
	}

	registry := policy.NewRegistry(
		map[string]policy.ValidatorMap{
			"AfterTime": policy.NewValidatorMap(map[string]policy.Validator{"datetime": &policy.AfterTime{}, "StringMatch": &policy.StringMatch{}}),
		},
	)
	request := policy.Request{
		Resource: policy.Input{
			Value: "us:aws:graphql",
		},
		Action: policy.Input{
			Value: "update",
		},
		Metadata: map[string]string{
			"datetime": "13:00",
		},
	}

	pr := p.Validate(&request, &registry)
	data, _ := json.MarshalIndent(pr, "", "   ")
	fmt.Println(string(data))
}
