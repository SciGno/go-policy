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
                "AfterTime": "12:00"
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
	// else {
	// 	d, err2 := json.MarshalIndent(&p, "", "   ")
	// 	if err2 != nil {
	// 		fmt.Printf("%s\n", err)
	// 	}
	// 	println(string(d))
	// }

	registry := policy.NewRegistry(&policy.DelimitedValidator{}, &policy.ActionValidator{}, map[string]policy.Validator{"AfterTime": &policy.AfterTime{}})
	request := policy.Request{Action: "update", Resource: "us:aws:graphqls", Condition: map[string]interface{}{"AfterTime": "11:00"}}

	fmt.Println(p.Validate(&request, &registry))
}
