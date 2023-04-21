package glesys

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	c := &mockClient{body: `{ "response": { "login": { "username": "alice@example.com", "apikey": "abc-123-xyz" } } }`}
	l := UserService{client: c}

	loginDetails, error := l.DoOTPLogin(context.Background(), "alice@example.com", "SuperSecretPassword123", "cccc123xyzotpstring")
	if error != nil {
		fmt.Printf("Error on login %s", error)
	}

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "user/login", c.lastPath, "path used is correct")
	assert.Equal(t, "alice@example.com", loginDetails.Username, "Username is correct")
	assert.Equal(t, "abc-123-xyz", loginDetails.APIKey, "APIKEY is correct")
}

func TestUserListOrganizations(t *testing.T) {
	c := &mockClient{body: `{ "response": { "organizations": [{
		"id": 1337,
		"state": "active",
		"featureflags": [],
		"verification": "verified",
		"type": "personal",
		"isowner": "yes",
		"contactperson": "",
		"nationalidnumber": "123456",
		"name": "Alice Smith",
		"address": "Kanslistvägen 12",
		"city": "Falkenberg",
		"zipcode": "123 45",
		"country": "SWEDEN",
		"separateinvoiceaddress": "no",
		"invoicereference": "test",
		"invoiceaddress": "Kanslistvägen 12",
		"invoicecity": "Falkenberg",
		"invoicezipcode": "123 45",
		"invoicecountry": "SWEDEN",
		"email": "alice@example.com",
		"phonenumber": "",
		"invoicedeliverymethod": "email",
		"billingemailoverrides": ["invoices@example.com"],
		"currentbillingemails": ["invoices@example.com"],
		"canaccessbilling": "yes",
		"expiredinvoices": "no",
		"unpaidinvoices": "no",
		"vatnumber": "",
		"paymentmethod": "invoice",
		"paymentcard": "",
		"allowedpaymentmethods": ["invoice", "card"],
		"paymenttermsnetdays": 20,
		"canpayinvoicesmanuallyusingpaypal": true,
		"slacost": "",
		"slalevel": "base",
		"slapincode": "1234",
		"slaphonenumber": "+461234567890"
		}] } }`}

	l := UserService{client: c}

	orgs, _ := l.ListOrganizations(context.Background())

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "user/listorganizations", c.lastPath, "path used is correct")
	assert.Equal(t, 1337, (*orgs)[0].ID, "Organization ID is correct")
	assert.Equal(t, "personal", (*orgs)[0].Type, "Organization Type is correct")
}

func TestCustomerListProjects(t *testing.T) {
	c := &mockClient{body: `{ "response": {
	"projects": [{"accountname": "cl98765",
	    "name": "prod",
	    "color": "silver",
	    "currency": "SEK",
	    "lockedreason": "",
	    "access": "full",
	    "customernumber": 1337},
	   {"accountname": "cl123456",
	    "name": "dev",
	    "color": "sandybrown",
	    "currency": "SEK",
	    "lockedreason": "",
	    "access": "full",
	    "customernumber": 1337}]}}
	`}
	l := UserService{client: c}

	projects, _ := l.ListCustomerProjects(context.Background(), "1337")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "customer/listprojects", c.lastPath, "path used is correct")
	assert.Equal(t, "prod", (*projects)[0].Name, "project name is correct")
	assert.Equal(t, "dev", (*projects)[1].Name, "project name is correct")
	assert.Equal(t, "silver", (*projects)[0].Color, "project color is correct")
	assert.Equal(t, 1337, (*projects)[0].Customernumber, "project customernumber is correct")
}
