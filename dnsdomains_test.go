package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDnsDomainsAdd(t *testing.T) {
	c := &mockClient{body: `{"response": {"domain": {"domainname": "example.com",
	  "createtime": "2019-07-02T21:55:18+02:00", "displayname": "example.com",
	  "recordcount": 9, "registrarinfo": "None", "usingglesysnameserver": "no"}}}`}
	d := DNSDomainService{client: c}

	params := AddDNSDomainParams{
		Name:          "example.com",
		CreateRecords: "no",
	}

	domain, _ := d.AddDNSDomain(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/add", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", domain.Name, "Domain name is correct")
	assert.Equal(t, 9, domain.RecordCount, "Domain has the correct number of records")
	assert.Equal(t, "no", domain.UsingGlesysNameserver, "Domain is not using glesys nameservers")
}

func TestDnsDomainsDeleteDomain(t *testing.T) {
	c := &mockClient{}
	d := DNSDomainService{client: c}

	params := DeleteDNSDomainParams{
		Name: "example.com",
	}

	d.Delete(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/delete", c.lastPath, "path used is correct")
}

func TestDnsDomainsEdit(t *testing.T) {
	c := &mockClient{body: `{"response": { "domain": {"domainname": "example.com",
          "createtime": "2010-07-13T11:13:50+02:00", "displayname": "example.com",
          "recordcount": 9, "usingglesysnameserver": "yes",
          "primarynameserver": "ns1.namesystem.se.",
          "responsibleperson": "registry.glesys.se.",
          "ttl": 3600, "refresh": 10800, "retry": 2400, "expire": 1814400, "minimum": 10800,
          "contactinfo": "None", "registrarinfo": {"state": "OK", "statedescription": "", "expire": "2038-01-19",
            "autorenew": "yes", "tld": "com", "invoicenumber": "None"}}}}`}
	d := DNSDomainService{client: c}

	params := EditDNSDomainParams{
		Name:  "example.com",
		Retry: 2400,
	}

	domain, _ := d.Edit(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/edit", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", domain.Name, "Domain name is correct")
	assert.Equal(t, 2400, domain.Retry, "Domain Retry correct")
}

func TestDnsDomainsAvailable(t *testing.T) {
	c := &mockClient{body: `{"response": {"domain": [ {"domainname": "example.com",
	   "available": true,
	   "prices": [{"amount": 123, "currency": "SEK", "years": 1}, {"amount": 1230, "currency": "SEK", "years": 10}]
           }]}}`}
	d := DNSDomainService{client: c}

	domains, _ := d.Available(context.Background(), "example.com")

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/available", c.lastPath, "path used is correct")
	assert.Equal(t, true, (*domains)[0].Available, "Domain is available")
	assert.Equal(t, 1230.00, (*domains)[0].Prices[1].Amount, "Domain amount is correct")
}

func TestDnsDomainsDetails(t *testing.T) {
	c := &mockClient{body: `{"response": { "domain": {"domainname": "example.com",
          "createtime": "2010-07-13T11:13:50+02:00", "displayname": "example.com",
          "recordcount": 9, "usingglesysnameserver": "yes",
          "primarynameserver": "ns1.namesystem.se.",
          "responsibleperson": "registry.glesys.se.",
          "ttl": 3600, "refresh": 10800, "retry": 2700, "expire": 1814400, "minimum": 10800,
          "contactinfo": "None", "registrarinfo": {"state": "OK", "statedescription": "", "expire": "2038-01-19",
            "autorenew": "yes", "tld": "com", "invoicenumber": "None"}}}}`}
	d := DNSDomainService{client: c}

	domain, _ := d.Details(context.Background(), "example.com")

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/details", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", domain.Name, "Domain name is correct")
	assert.Equal(t, 3600, domain.TTL, "Domain has the correct TTL")
	assert.Equal(t, "yes", domain.UsingGlesysNameserver, "Domain is using glesys nameservers")
}

func TestDnsDomainsList(t *testing.T) {
	c := &mockClient{body: `{"response": { "domains": [{"domainname": "example.com",
          "createtime": "2010-07-13T11:13:50+02:00", "displayname": "example.com",
          "recordcount": 4, "registrarinfo": {"state": "OK", "statedescription": "", "expire": "2038-01-19",
          "autorenew": "yes", "tld": "com", "invoicenumber": "None"}}]}}`}

	d := DNSDomainService{client: c}

	domains, _ := d.List(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/list", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", (*domains)[0].Name, "Domain name is correct")
	assert.Equal(t, 4, (*domains)[0].RecordCount, "record count correct")
	assert.Equal(t, "yes", (*domains)[0].RegistrarInfo.AutoRenew, "Domain AutoRenew is set")
}

func TestDnsDomainsRegister(t *testing.T) {
	c := &mockClient{body: `{"response": { "domain": {"domainname": "example.com",
          "createtime": "2010-07-13T11:13:50+02:00", "displayname": "example.com",
          "recordcount": 4, "registrarinfo": {"state": "REGISTER", "statedescription": "", "expire": "2038-01-19",
          "autorenew": "yes", "tld": "com", "invoicenumber": "None"}}}}`}

	d := DNSDomainService{client: c}
	params := RegisterDNSDomainParams{
		Name:         "example.com",
		Firstname:    "Alice",
		Lastname:     "Smith",
		Email:        "alice@example.com",
		Address:      "Badhusvägen 45",
		City:         "Falkenberg",
		ZipCode:      "31132",
		Country:      "SE",
		Organization: "Internetz",
		NationalID:   13337,
	}

	domain, _ := d.Register(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/register", c.lastPath, "path used is correct")
	assert.Equal(t, "REGISTER", domain.RegistrarInfo.State, "Domain is in correct state")
}

func TestDnsDomainsRenew(t *testing.T) {
	c := &mockClient{body: `{"response": { "domain": {"domainname": "example.com",
          "createtime": "2010-07-13T11:13:50+02:00", "displayname": "example.com",
          "recordcount": 4, "registrarinfo": {"state": "RENEW", "statedescription": "", "expire": "2038-01-19",
          "autorenew": "yes", "tld": "com", "invoicenumber": "None"}}}}`}

	d := DNSDomainService{client: c}
	params := RenewDNSDomainParams{
		Name:     "example.com",
		NumYears: 1,
	}

	domain, _ := d.Renew(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/renew", c.lastPath, "path used is correct")
	assert.Equal(t, "RENEW", domain.RegistrarInfo.State, "Domain is in correct state")
}

func TestDnsDomainsSetAutoRenew(t *testing.T) {
	c := &mockClient{body: `{"response": { "domain": {"domainname": "example.com",
          "createtime": "2010-07-13T11:13:50+02:00", "displayname": "example.com",
          "recordcount": 4, "registrarinfo": {"state": "RENEW", "statedescription": "", "expire": "2038-01-19",
          "autorenew": "yes", "tld": "com", "invoicenumber": "None"}}}}`}

	d := DNSDomainService{client: c}
	params := SetAutoRenewParams{
		Name:         "example.com",
		SetAutoRenew: "yes",
	}

	domain, _ := d.SetAutoRenew(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/setautorenew", c.lastPath, "path used is correct")
	assert.Equal(t, "yes", domain.RegistrarInfo.AutoRenew, "Domain is set to renew automatically.")
}

func TestDnsDomainsTransfer(t *testing.T) {
	c := &mockClient{body: `{"response": { "domain": {"domainname": "example.com",
          "createtime": "2010-07-13T11:13:50+02:00", "displayname": "example.com",
          "recordcount": 4, "registrarinfo": {"state": "TRANSFER", "statedescription": "", "expire": "2038-01-19",
          "autorenew": "yes", "tld": "com", "invoicenumber": "None"}}}}`}

	d := DNSDomainService{client: c}
	params := RegisterDNSDomainParams{
		Name:         "example.com",
		Firstname:    "Alice",
		Lastname:     "Smith",
		Email:        "alice@example.com",
		Address:      "Badhusvägen 45",
		City:         "Falkenberg",
		ZipCode:      "31132",
		Country:      "SE",
		Organization: "Internetz",
		NationalID:   13337,
	}

	domain, _ := d.Transfer(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/transfer", c.lastPath, "path used is correct")
	assert.Equal(t, "TRANSFER", domain.RegistrarInfo.State, "Domain is in correct state")
}

func TestDnsDomainsAddRecord(t *testing.T) {
	c := &mockClient{body: `{"response": { "record":
          {"recordid": 1234569, "domainname": "example.com", "host": "test", "type": "A", "data": "127.0.0.1", "ttl": 3600}
	}}`}

	params := AddRecordParams{
		DomainName: "example.com",
		Host:       "test",
		Data:       "127.0.0.1",
		Type:       "A",
	}

	d := DNSDomainService{client: c}

	record, _ := d.AddRecord(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/addrecord", c.lastPath, "path used is correct")
	assert.Equal(t, "test", (*record).Host, "Record host is correct")
	assert.Equal(t, "127.0.0.1", (*record).Data, "Record data is correct")
}

func TestDnsDomainsListRecords(t *testing.T) {
	c := &mockClient{body: `{"response": { "records": [
	  {"recordid": 1234567, "domainname": "example.com", "host": "www", "type": "A", "data": "127.0.0.1", "ttl": 3600},
          {"recordid": 1234568, "domainname": "example.com", "host": "mail", "type": "A", "data": "127.0.0.3", "ttl": 3600}
	]}}`}

	d := DNSDomainService{client: c}

	records, _ := d.ListRecords(context.Background(), "example.com")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/listrecords", c.lastPath, "path used is correct")
	assert.Equal(t, "www", (*records)[0].Host, "Record host is correct")
	assert.Equal(t, "127.0.0.3", (*records)[1].Data, "Record data is correct")
}

func TestDnsDomainsUpdateRecord(t *testing.T) {
	c := &mockClient{body: `{"response": { "record":
          {"recordid": 1234567, "domainname": "example.com", "host": "mail", "type": "A", "data": "127.0.0.3", "ttl": 3600}
	}}`}

	params := UpdateRecordParams{
		RecordID: 1234567,
		Data:     "127.0.0.3",
	}

	d := DNSDomainService{client: c}

	record, _ := d.UpdateRecord(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/updaterecord", c.lastPath, "path used is correct")
	assert.Equal(t, "mail", (*record).Host, "Record host is correct")
	assert.Equal(t, "127.0.0.3", (*record).Data, "Record data is correct")
}

func TestDnsDomainsDeleteRecord(t *testing.T) {
	c := &mockClient{}
	d := DNSDomainService{client: c}

	d.DeleteRecord(context.Background(), 1234567)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/deleterecord", c.lastPath, "path used is correct")
}

func TestDnsDomainsChangeNameservers(t *testing.T) {
	c := &mockClient{}
	d := DNSDomainService{client: c}

	params := ChangeNameserverParams{
		DomainName: "example.com",
		NS1:        "ns1.namesystem.se.",
		NS2:        "ns2.example.com.",
	}

	d.ChangeNameservers(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/changenameservers", c.lastPath, "path used is correct")
}

func TestDnsDomainsGenerateAuthCode(t *testing.T) {
	c := &mockClient{body: `{"response": { "authcode": "abcxy123-=%" }}`}
	d := DNSDomainService{client: c}

	authcode, _ := d.GenerateAuthCode(context.Background(), "example.com")

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/generateauthcode", c.lastPath, "path used is correct")
	assert.Equal(t, "abcxy123-=%", authcode, "correct return data")
}
