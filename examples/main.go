package main

import (
	"encoding/json"
	"fmt"

	"github.com/scigno/go-policy/policy"
)

var pol = []byte(`
{
    "id": "f0dacf4e-c051-475e-b9dc-f86cd2d9539c",
    "name": "Delegate",
    "version": "1.0",
    "weight": "1",
    "statement": [
        {
            "sid": "AllPublicationActions",
            "effect": "Allow",
            "action": [
                "publication:*",
                "publication:updatePublication"
            ],
            "resource": [
                "us:aws:graphql",
                "us:azure:graphql"
            ],
            "condition": {
                "CIDR": [
                    "192.168.1.0/24",
                    "10.10.20.12"
                ],
                "DateRanges": [
                    {
                        "After": "2019/10/16",
                        "Before": "2019/11/12"
                    },
                    {
                        "After": "2019/11/15",
                        "Before": "2019/11/31"
                    }
                ],
                "BeforeDate": {
                    "Date": "2019/11/16"
                },
                "AfterDate": {
                    "Date": "2019/11/16"
                },
                "TimeRanges": [
                    {
                        "After": "02:00",
                        "Before": "08:00"
                    },
                    {
                        "After": "17:00",
                        "Before": "20:00"
                    }
                ],
                "BeforeTime": {
                    "Time": "02:39"
                },
                "AfterTime": {
                    "Time": "02:39"
                }
            }
        },
        {
            "sid": "NoDelegateActions",
            "effect": "Deny",
            "action": [
                "delegate:assignDelegate",
                "delegate:unassignDelegate",
                "service:action"
            ],
            "resource": [
                "us:aws:graphql",
                "us:azure:graphql"
            ],
            "condition": {
                "WithinTime": {
                    "After": "02:00",
                    "Before": "13:00"
                }
            }
        }
    ]
}
`)

func main() {
	p := policy.Policy{}

	if err := json.Unmarshal(pol, &p); err == nil {
		fmt.Printf("%+v\n", p)
		d, err2 := json.MarshalIndent(&p, "", " ")
		if err2 != nil {
			fmt.Printf("%s\n", err)
		}
		println(string(d))
	} else {
		fmt.Printf("%s\n", err)
	}

}
