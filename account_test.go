package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountsInfo(t *testing.T) {
	c := &mockClient{body: `{"response": {
  "accountinfo": {"accountname": "cl12345", "customernumber": 1234, "customer": {"customernumber": 1234,
    "contactinfo": {"contact": {"companyname": "User Bobsson", "contactperson": "", "invoicereference": "",
      "address": "Street 1", "zipcode": "123 45", "city": "Halmstad", "country": "SWEDEN", "nationalidnumber": "12345",
      "phonenumber": "", "email": "email@example.com"}, "invoice": {"invoiceaddress": "street 8", "invoicezipcode": "123 45",
      "invoicecity": "city", "invoicecountry": "SWEDEN"},
     "settings": {"separateinvoiceaddress": "no"}}},
   "currency": "SEK",
   "unpaidinvoices": "no",
   "apienabled": "yes",
   "servicelimits": {"maxnumservers": 9223372036854775807, "maxnumipv4": 9223372036854775807, "maxnumipv6": 9223372036854775807,
    "maxarchivetotalsize": 9223372036854775807, "maxnumemailaccounts": 9223372036854775807, "emailglobalquota": 3072000,
    "domainregistrationsallowed": 9223372036854775807, "domaintransfersallowed": 9223372036854775807},
   "services": {"server": "yes", "ip": "yes", "domain": "yes", "archive": "yes", "email": "yes", "invoice": "yes",
    "customer": "yes", "account": "yes", "api": "yes", "vpn": "yes"}}}}`}
	s := AccountService{client: c}

	account, _ := s.Info(context.Background())

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "account/info", c.lastPath, "path used is correct")
	assert.Equal(t, "cl12345", account.Accountname, "Accountname is correct")
	assert.Equal(t, 1234, account.Customer.Customernumber, "Accountname is correct")
}
