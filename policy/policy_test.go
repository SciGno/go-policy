package policy

// var full = []byte(`
// {
//     "id": "f0dacf4e-c051-475e-b9dc-f86cd2d9539c",
//     "name": "Delegate",
//     "version": "1.0",
//     "statement": [
//         {
//             "sid": "AllPublicationActions",
//             "effect": "Allow",
//             "action": [
//                 {
//                     "Service": [
//                         "publication:*",
//                         "publication:updatePublication"
//                     ]
//                 }
//             ],
//             "resource": [
//                 {
//                     "Cloud": [
//                         "us:aws:graphql",
//                         "us:azure:graphql"
//                     ]
//                 }
//             ],
//             "condition": [
//                 {
//                     "CIDR": [
//                         "192.168.1.0/24",
//                         "10.10.20.12"
//                     ]
//                 },
//                 {
//                     "WithinDate": {
//                         "After": "2019/11/16",
//                         "Before": "2019/11/19"
//                     }
//                 },
//                 {
//                     "BeforeDate": {
//                         "Date": "2019/11/16"
//                     }
//                 },
//                 {
//                     "AfterDate": {
//                         "Date": "2019/11/16"
//                     }
//                 },
//                 {
//                     "WithinTime": {
//                         "After": "02:00",
//                         "Before": "13:00"
//                     }
//                 },
//                 {
//                     "BeforeTime": {
//                         "Time": "02:39"
//                     }
//                 },
//                 {
//                     "AfterTime": {
//                         "Time": "02:39"
//                     }
//                 }
//             ]
//         },
//         {
//             "sid": "NoDelegateActions",
//             "effect": "Deny",
//             "action": [
//                 {
//                     "Service": [
//                         "delegate:assignDelegate",
//                         "delegate:unassignDelegate",
//                         "service:action"
//                     ]
//                 }
//             ],
//             "resource": [
//                 {
//                     "Cloud": [
//                         "us:aws:graphql",
//                         "us:azure:graphql"
//                     ]
//                 }
//             ],
//             "condition": [
//                 {
//                     "WithinTime": {
//                         "After": "02:00",
//                         "Before": "13:00"
//                     }
//                 }
//             ]
//         }
//     ]
// }
// `)

// var simple = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"sid":"Statement1","effect":"Allow","action":[{"Service":["publication:*","publication:updatePublication"]}],"resource":[{"Cloud":["us:aws:graphql","us:azure:graphql"]}],"condition":[{"CIDR":["192.168.1.0/24","10.10.20.12"]}]}]}`)
// var missingACTION = []byte(`{"id": "123","name": "Test","version": "1.0","statement": [{"sid": "Statement1","effect": "Allow","resource": [{"Cloud":["us:aws:graphql","us:azure:graphql"]}],"condition": [{"CIDR": ["192.168.1.0/24","10.10.20.12"]}]}]}`)
// var nullACTION = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"sid":"Statement1","effect":"Allow","action":null,"resource":[{"Cloud":["us:aws:graphql","us:azure:graphql"]}],"condition":[{"CIDR":["192.168.1.0/24","10.10.20.12"]}]}]}`)
// var nullRESOURCE = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"sid":"Statement1","effect":"Allow","action":[{"Service":["publication:*","publication:updatePublication"]}],"resource":null,"condition":[{"CIDR":["192.168.1.0/24","10.10.20.12"]}]}]}`)
// var nullCONDITION = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"sid":"Statement1","effect":"Allow","action":[{"Service":["publication:*","publication:updatePublication"]}],"resource":[{"Cloud":["us:aws:graphql","us:azure:graphql"]}],"condition":null}]}`)
// var emptySID = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"sid":"","effect":"Allow","action":[{"Service":["publication:*","publication:updatePublication"]}],"resource":[{"Cloud":["us:aws:graphql","us:azure:graphql"]}]}]}`)
// var emptyEFFECT = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"sid":"Test","effect":"","action":[{"Service":["publication:*","publication:updatePublication"]}],"resource":[{"Cloud":["us:aws:graphql","us:azure:graphql"]}]}]}`)
// var missingSID = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"effect":"Allow","action":[{"Service":["publication:*","publication:updatePublication"]}],"resource":[{"Cloud":["us:aws:graphql","us:azure:graphql"]}]}]}`)
// var missingEFFECT = []byte(`{"id":"123","name":"Test","version":"1.0","statement":[{"sid":"Testing","action":[{"Service":["publication:*","publication:updatePublication"]}],"resource":[{"Cloud":["us:aws:graphql","us:azure:graphql"]}]}]}`)
// var simpleBadJSON = []byte(`{"id": "123","name": "Test","version": "1.0","statement": }`)
// var simpleWithError = []byte(`{"id":"123","name":"Test","version":"1.0"}`)
// var missingID = []byte(`{"name":"Test","version":"1.0"}`)
// var missingNAME = []byte(`{"id":"123","version":"1.0"}`)
// var missingVERSION = []byte(`{"id":"123","name":"Test"}`)
// var missingSTATEMENT = []byte(`{"id":"123","name":"Test","version":"1.0"}`)

// var a = []action.Action{
// 	&action.ServiceAction{
// 		Actions: []string{"publication:*"},
// 	},
// }

// var r = []resource.Resource{
// 	&resource.StringMatch{
// 		Resources: []string{"us:aws:graphql"},
// 	},
// }

// var c = []condition.Condition{
// 	&condition.CIDR{
// 		CIDR: []string{
// 			"192.168.1.0/24",
// 			"10.10.20.12",
// 		},
// 	},
// }

// var s = []statement.Statement{
// 	{
// 		StatementID: "AllPublicationActions",
// 		Effect:      statement.ALLOW,
// 		Action:      a,
// 		Resource:    r,
// 		Condition:   c,
// 	},
// 	{
// 		StatementID: "NoDelegateActions",
// 		Effect:      statement.DENY,
// 		Action:      a,
// 		Resource:    r,
// 		Condition:   c,
// 	},
// }

// func TestPolicy_UnmarshallJSON(t *testing.T) {
// 	type args struct {
// 		a []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{"full", args{full}, false},
// 		{"simple", args{simple}, false},
// 		{"simpleWithError", args{simpleWithError}, true},
// 		{"simpleBadJSON", args{simpleBadJSON}, true},
// 		{"emptySID", args{emptySID}, true},
// 		{"emptyEFFECT", args{emptyEFFECT}, true},
// 		{"missingSID", args{missingSID}, true},
// 		{"missingEFFECT", args{missingEFFECT}, true},
// 		{"missingACTION", args{missingACTION}, true},
// 		{"nullACTION", args{nullACTION}, true},
// 		{"nullRESOURCE", args{nullRESOURCE}, true},
// 		{"nullCONDITION", args{nullCONDITION}, true},
// 		{"missingID", args{missingID}, true},
// 		{"missingNAME", args{missingNAME}, true},
// 		{"missingVERSION", args{missingVERSION}, true},
// 		{"missingSTATEMENT", args{missingSTATEMENT}, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := &Policy{}
// 			if err := p.UnmarshallJSON(tt.args.a); (err != nil) != tt.wantErr {
// 				t.Errorf("Policy.UnmarshallJSON() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestPolicy_Validate(t *testing.T) {

// 	var AllowStatement = []statement.Statement{
// 		{
// 			StatementID: "AllPublicationActions",
// 			Effect:      statement.ALLOW,
// 			Action:      a,
// 			Resource:    r,
// 			Condition:   c,
// 		},
// 	}

// 	var AllowNoCondition = []statement.Statement{
// 		{
// 			StatementID: "AllPublicationActions",
// 			Effect:      statement.ALLOW,
// 			Action:      a,
// 			Resource:    r,
// 		},
// 	}

// 	var DenyStatement = []statement.Statement{
// 		{
// 			StatementID: "NoPublications",
// 			Effect:      statement.DENY,
// 			Action:      a,
// 			Resource:    r,
// 			Condition:   c,
// 		},
// 	}

// 	var DenyNoConditions = []statement.Statement{
// 		{
// 			StatementID: "NoPublications",
// 			Effect:      statement.DENY,
// 			Action:      a,
// 			Resource:    r,
// 		},
// 	}

// 	r1 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"publication:addPublication"},
// 			},
// 		},
// 		Resource: []resource.Resource{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws:graphql"},
// 			},
// 		},
// 		Condition: []condition.Condition{
// 			&condition.CIDR{
// 				CIDR: []string{
// 					"10.10.20.12",
// 				},
// 			},
// 		},
// 	}

// 	r2 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"delegate:addPublication"},
// 			},
// 		},
// 		Resource: []resource.Resource{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws:graphql"},
// 			},
// 		},
// 	}

// 	r3 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"delegate:addPublication"},
// 			},
// 		},
// 		Resource: []resource.Resource{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws2:graphql"},
// 			},
// 		},
// 	}

// 	r4 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"delegate:addPublication"},
// 			},
// 		},
// 		Resource: []resource.Resource{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws:graphql"},
// 			},
// 		},
// 		Condition: []condition.Condition{
// 			&condition.CIDR{
// 				CIDR: []string{
// 					"10.10.20.123",
// 				},
// 			},
// 		},
// 	}

// 	r5 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"publication:addPublication"},
// 			},
// 		},
// 		Resource: []resource.Resource{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws:graphql"},
// 			},
// 		},
// 	}

// 	r6 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"publication:addPublication"},
// 			},
// 		},
// 		Resource: []resource.Resource{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws2:graphql"},
// 			},
// 		},
// 	}

// 	r7 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"publication:addPublication"},
// 			},
// 		},
// 		Resource: []resource.Resource{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws:graphql"},
// 			},
// 		},
// 		Condition: []condition.Condition{
// 			&condition.CIDR{
// 				CIDR: []string{
// 					"120.10.20.123",
// 				},
// 			},
// 		},
// 	}

// 	type fields struct {
// 		ID        string
// 		Name      string
// 		Version   string
// 		Statement []statement.Statement
// 	}
// 	type args struct {
// 		r Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   bool
// 	}{
// 		{"AllowValidation", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", AllowStatement}, args{r1}, true},
// 		{"AllowValidationNoCondition", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", AllowNoCondition}, args{r1}, true},
// 		{"FailsAction", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", AllowNoCondition}, args{r2}, false},
// 		{"FailsResource", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", AllowStatement}, args{r3}, false},
// 		{"FailsCondition", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", AllowStatement}, args{r4}, false},
// 		{"FailsNoCondition", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", AllowStatement}, args{r5}, false},
// 		{"DenyActions", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", DenyNoConditions}, args{r5}, false},
// 		{"DenyNoCondition", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", DenyStatement}, args{r5}, false},
// 		{"DenyResource", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", DenyStatement}, args{r6}, false},
// 		{"DenyOnConditions", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", DenyStatement}, args{r1}, false},
// 		{"DenyOnAction", fields{"f0dacf4e-c051-475e-b9dc-f86cd2d9539c", "Delegate", "1.0", DenyStatement}, args{r7}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := &Policy{
// 				ID:        tt.fields.ID,
// 				Name:      tt.fields.Name,
// 				Version:   tt.fields.Version,
// 				Statement: tt.fields.Statement,
// 			}
// 			if got := p.Validate(tt.args.r); got != tt.want {
// 				t.Errorf("Policy.Validate() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPolicy_validateResources(t *testing.T) {

// 	res := []Certifier{
// 		&resource.StringMatch{
// 			Resources: []string{"us:aws:graphql"},
// 		},
// 	}

// 	r1 := Request{
// 		Resource: []Certifier{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws:graphql"},
// 			},
// 		},
// 	}

// 	r2 := Request{
// 		Resource: []Certifier{
// 			&resource.StringMatch{
// 				Resources: []string{"us:aws2:graphql"},
// 			},
// 		},
// 	}

// 	type args struct {
// 		resources []Certifier
// 		request   Request
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{"Pass", args{res, r1}, true},
// 		{"Fails", args{res, r2}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := &Policy{}
// 			if got := p.validateResources(tt.args.resources, tt.args.request); got != tt.want {
// 				t.Errorf("Policy.validateResources() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPolicy_validateConfitions(t *testing.T) {

// 	c := []condition.Condition{
// 		&condition.CIDR{
// 			CIDR: []string{
// 				"10.10.20.12",
// 			},
// 		},
// 	}

// 	c1 := []condition.Condition{
// 		&condition.CIDR{
// 			CIDR: []string{
// 				"10.10.20.123",
// 			},
// 		},
// 	}

// 	r1 := Request{
// 		Condition: []condition.Condition{
// 			&condition.CIDR{
// 				CIDR: []string{
// 					"10.10.20.12",
// 				},
// 			},
// 		},
// 	}

// 	type args struct {
// 		conditions []condition.Condition
// 		request    Request
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{"AllowConditionMatch", args{c, r1}, true},
// 		{"AllowConditionFails", args{c1, r1}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := &Policy{}
// 			if got := p.validateConfitions(tt.args.conditions, tt.args.request); got != tt.want {
// 				t.Errorf("Policy.validateConfitions() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPolicy_validateActions(t *testing.T) {

// 	a1 := []action.Action{
// 		&action.ServiceAction{
// 			Actions: []string{"publication:addPublication"},
// 		},
// 	}

// 	r1 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"publication:addPublication"},
// 			},
// 		},
// 	}

// 	r2 := Request{
// 		Action: []action.Action{
// 			&action.ServiceAction{
// 				Actions: []string{"delegate:addDelegate"},
// 			},
// 		},
// 	}

// 	type args struct {
// 		actions []action.Action
// 		request Request
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{"PassAction", args{a1, r1}, true},
// 		{"AllowConditionFails", args{a1, r2}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := &Policy{}
// 			if got := p.validateActions(tt.args.actions, tt.args.request); got != tt.want {
// 				t.Errorf("Policy.validateActions() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
