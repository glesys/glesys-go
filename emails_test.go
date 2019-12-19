package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailsOverview(t *testing.T) {
	c := &mockClient{body: `{"response":{"overview":{"summary":{"accounts":0,"maxaccounts":2000,"aliases":0,"maxaliases":10000},"domains":[{"domainname":"example.com","displayname":"example.com","accounts":0,"aliases":0}],"meta":{"page":1,"total":1,"perpage":1}}}}`}
	s := EmailService{client: c}

	emailoverview, _ := s.Overview(context.Background(), OverviewParams{})
	emaildomains := &emailoverview.Domains
	emailsummary := &emailoverview.Summary
	emailmeta := &emailoverview.Meta

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/overview", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DomainName, "domainname is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DisplayName, "displayname is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Accounts, "number of accounts is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Aliases, "number of aliases is correct")
	assert.Equal(t, 0, (*emailsummary).Accounts, "number of summary accounts is correct")
	assert.Equal(t, 2000, (*emailsummary).MaxAccounts, "number of summary maxaccounts is correct")
	assert.Equal(t, 0, (*emailsummary).Aliases, "number of summary aliases is correct")
	assert.Equal(t, 10000, (*emailsummary).MaxAliases, "number of summary maxaliases is correct")
	assert.Equal(t, 1, (*emailmeta).Page, "Page number is correct")
	assert.Equal(t, 1, (*emailmeta).Total, "Total number is correct")
	assert.Equal(t, 1, (*emailmeta).PerPage, "Per page number is correct")
}

func TestEmailsOverviewWithFilterParameter(t *testing.T) {
	c := &mockClient{body: `{"response":{"overview":{"summary":{"accounts":0,"maxaccounts":2000,"aliases":0,"maxaliases":10000},"domains":[{"domainname":"example.com","displayname":"example.com","accounts":0,"aliases":0}],"meta":{"page":1,"total":1,"perpage":1}}}}`}
	s := EmailService{client: c}

	emailoverview, _ := s.Overview(context.Background(), OverviewParams{Filter: "example.com"})
	emaildomains := &emailoverview.Domains
	emailsummary := &emailoverview.Summary
	emailmeta := &emailoverview.Meta

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/overview/filter/example.com", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DomainName, "domainname is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DisplayName, "displayname is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Accounts, "number of accounts is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Aliases, "number of aliases is correct")
	assert.Equal(t, 0, (*emailsummary).Accounts, "number of summary accounts is correct")
	assert.Equal(t, 2000, (*emailsummary).MaxAccounts, "number of summary maxaccounts is correct")
	assert.Equal(t, 0, (*emailsummary).Aliases, "number of summary aliases is correct")
	assert.Equal(t, 10000, (*emailsummary).MaxAliases, "number of summary maxaliases is correct")
	assert.Equal(t, 1, (*emailmeta).Page, "Page number is correct")
	assert.Equal(t, 1, (*emailmeta).Total, "Total number is correct")
	assert.Equal(t, 1, (*emailmeta).PerPage, "Per page number is correct")
}

func TestEmailsOverviewWithPageParameter(t *testing.T) {
	c := &mockClient{body: `{"response":{"overview":{"summary":{"accounts":0,"maxaccounts":2000,"aliases":0,"maxaliases":10000},"domains":[{"domainname":"example.com","displayname":"example.com","accounts":0,"aliases":0}],"meta":{"page":2,"total":31,"perpage":30}}}}`}
	s := EmailService{client: c}
	emailoverview, _ := s.Overview(context.Background(), OverviewParams{Page: 2})
	emaildomains := &emailoverview.Domains
	emailsummary := &emailoverview.Summary
	emailmeta := &emailoverview.Meta

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/overview/page/2", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DomainName, "domainname is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DisplayName, "displayname is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Accounts, "number of accounts is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Aliases, "number of aliases is correct")
	assert.Equal(t, 0, (*emailsummary).Accounts, "number of summary accounts is correct")
	assert.Equal(t, 2000, (*emailsummary).MaxAccounts, "number of summary maxaccounts is correct")
	assert.Equal(t, 0, (*emailsummary).Aliases, "number of summary aliases is correct")
	assert.Equal(t, 10000, (*emailsummary).MaxAliases, "number of summary maxaliases is correct")
	assert.Equal(t, 2, (*emailmeta).Page, "Page number is correct")
	assert.Equal(t, 31, (*emailmeta).Total, "Total number is correct")
	assert.Equal(t, 30, (*emailmeta).PerPage, "Per page number is correct")
}

func TestEmailsOverviewWithFilterAndPageParameter(t *testing.T) {
	c := &mockClient{body: `{"response":{"overview":{"summary":{"accounts":0,"maxaccounts":2000,"aliases":0,"maxaliases":10000},"domains":[{"domainname":"example.com","displayname":"example.com","accounts":0,"aliases":0}],"meta":{"page":1,"total":1,"perpage":1}}}}`}
	s := EmailService{client: c}

	emailoverview, _ := s.Overview(context.Background(), OverviewParams{Filter: "example.com", Page: 1})
	emaildomains := &emailoverview.Domains
	emailsummary := &emailoverview.Summary
	emailmeta := &emailoverview.Meta

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/overview/filter/example.com/page/1", c.lastPath, "path used is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DomainName, "domainname is correct")
	assert.Equal(t, "example.com", (*emaildomains)[0].DisplayName, "displayname is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Accounts, "number of accounts is correct")
	assert.Equal(t, 0, (*emaildomains)[0].Aliases, "number of aliases is correct")
	assert.Equal(t, 0, (*emailsummary).Accounts, "number of summary accounts is correct")
	assert.Equal(t, 2000, (*emailsummary).MaxAccounts, "number of summary maxaccounts is correct")
	assert.Equal(t, 0, (*emailsummary).Aliases, "number of summary aliases is correct")
	assert.Equal(t, 10000, (*emailsummary).MaxAliases, "number of summary maxaliases is correct")
	assert.Equal(t, 1, (*emailmeta).Page, "Page number is correct")
	assert.Equal(t, 1, (*emailmeta).Total, "Total number is correct")
	assert.Equal(t, 1, (*emailmeta).PerPage, "Per page number is correct")
}

func TestEmailsGlobalQuotaWithoutNewParam(t *testing.T) {
	c := &mockClient{body: `{"response":{"globalquota":{"usage":0,"max":10240}}}`}
	s := EmailService{client: c}

	emailglobalquota, _ := s.GlobalQuota(context.Background(), GlobalQuotaParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/globalquota", c.lastPath, "path used is correct")
	assert.Equal(t, 0, emailglobalquota.Usage, "usage number is correct")
	assert.Equal(t, 10240, emailglobalquota.Max, "max number is correct")
}

func TestEmailsGlobalQuotaWithNewParam(t *testing.T) {
	c := &mockClient{body: `{"response":{"globalquota":{"usage":0,"max":20480}}}`}
	s := EmailService{client: c}

	emailglobalquota, _ := s.GlobalQuota(context.Background(), GlobalQuotaParams{GlobalQuota: 20480})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/globalquota", c.lastPath, "path used is correct")
	assert.Equal(t, 0, emailglobalquota.Usage, "usage number is correct")
	assert.Equal(t, 20480, emailglobalquota.Max, "max number is correct")
}

func TestEmailsList(t *testing.T) {
	c := &mockClient{body: `{"response":{"list":{"emailaccounts":[{"emailaccount":"user@example.com","displayname":"user@example.com","quota":{"max":200,"unit":"MB"},"antispamlevel":3,"antivirus":"yes","autorespond":"yes","autorespondmessage":"This is not the account you are looking for.\n\nMove along, move along.","autorespondsaveemail":"yes","rejectspam":"no","created":"2019-10-26T13:07:13+02:00","modified":"2019-10-26T15:38:51+02:00"}],"emailaliases":[{"emailalias":"alias@example.com","displayname":"alias@example.com","goto":"user@example.com"}]}}}`}
	s := EmailService{client: c}

	emaillist, _ := s.List(context.Background(), "example.com", ListEmailsParams{})
	emailaccounts := &emaillist.EmailAccounts
	emailaliases := &emaillist.EmailAliases

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/list", c.lastPath, "path used is correct")

	assert.Equal(t, "user@example.com", (*emailaccounts)[0].EmailAccount, "emailaccount is correct")
	assert.Equal(t, "user@example.com", (*emailaccounts)[0].DisplayName, "displayname is correct")
	assert.Equal(t, 200, (*emailaccounts)[0].Quota.Max, "quota.max is correct")
	assert.Equal(t, "MB", (*emailaccounts)[0].Quota.Unit, "quota.unit is correct")
	assert.Equal(t, 3, (*emailaccounts)[0].AntiSpamLevel, "antispamlevel is correct")
	assert.Equal(t, "yes", (*emailaccounts)[0].AntiVirus, "antivirus is correct")
	assert.Equal(t, "yes", (*emailaccounts)[0].AutoRespond, "autorespond is correct")
	assert.Equal(t, "This is not the account you are looking for.\n\nMove along, move along.", (*emailaccounts)[0].AutoRespondMessage, "autorespondmessage is correct")
	assert.Equal(t, "yes", (*emailaccounts)[0].AutoRespondSaveEmail, "autorespondsaveemail is correct")
	assert.Equal(t, "no", (*emailaccounts)[0].RejectSpam, "rejectspam is correct")
	assert.Equal(t, "2019-10-26T13:07:13+02:00", (*emailaccounts)[0].Created, "created is correct")
	assert.Equal(t, "2019-10-26T15:38:51+02:00", (*emailaccounts)[0].Modified, "modified is correct")

	assert.Equal(t, "alias@example.com", (*emailaliases)[0].EmailAlias, "emailalias is correct")
	assert.Equal(t, "alias@example.com", (*emailaliases)[0].DisplayName, "displayname is correct")
	assert.Equal(t, "user@example.com", (*emailaliases)[0].GoTo, "goto is correct")
}

func TestEmailEditAccount(t *testing.T) {
	c := &mockClient{body: `{"response":{"emailaccount":{"emailaccount":"user@example.com","displayname":"user@example.com","quota":{"max":200,"unit":"MB"},"antispamlevel":3,"antivirus":"yes","autorespond":"yes","autorespondmessage":"This is not the account you are looking for.\n\nMove along, move along.","autorespondsaveemail":"yes","rejectspam":"no","created":"2019-10-26T13:07:13+02:00","modified":"2019-11-10T22:09:14+01:00"}}}`}
	s := EmailService{client: c}

	editaccount, _ := s.EditAccount(context.Background(), "user@example.com", EditAccountParams{Quota: 200})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/editaccount", c.lastPath, "path used is correct")
	assert.Equal(t, "user@example.com", editaccount.EmailAccount, "emailaccount is correct")
	assert.Equal(t, "user@example.com", editaccount.DisplayName, "displayname is correct")
	assert.Equal(t, 200, editaccount.Quota.Max, "quota.max is correct")
	assert.Equal(t, "MB", editaccount.Quota.Unit, "quota.unit is correct")
	assert.Equal(t, 3, editaccount.AntiSpamLevel, "antispamlevel is correct")
	assert.Equal(t, "yes", editaccount.AntiVirus, "antivirus is correct")
	assert.Equal(t, "yes", editaccount.AutoRespond, "autorespond is correct")
	assert.Equal(t, "This is not the account you are looking for.\n\nMove along, move along.", editaccount.AutoRespondMessage, "autorespondmessage is correct")
	assert.Equal(t, "yes", editaccount.AutoRespondSaveEmail, "autorespondsaveemail is correct")
	assert.Equal(t, "no", editaccount.RejectSpam, "rejectspam is correct")
	assert.Equal(t, "2019-10-26T13:07:13+02:00", editaccount.Created, "created is correct")
	assert.Equal(t, "2019-11-10T22:09:14+01:00", editaccount.Modified, "modified is correct")
}

func TestEmailDelete(t *testing.T) {
	c := &mockClient{body: `{"response": {}}`}
	s := EmailService{client: c}

	s.Delete(context.Background(), "user@example.com")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/delete", c.lastPath, "path used is correct")
}

func TestCreateAccount(t *testing.T) {
	c := &mockClient{body: `{"response":{"emailaccount":{"emailaccount":"new_user@example.com","displayname":"new_user@example.com","quota":{"max":200,"unit":"MB"},"antispamlevel":3,"antivirus":"yes","autorespond":"no","autorespondmessage":null,"autorespondsaveemail":"yes","rejectspam":"no","created":"2019-11-29T21:31:28+01:00","modified":null}}}`}

	s := EmailService{client: c}

	createaccount, _ := s.CreateAccount(context.Background(), CreateAccountParams{EmailAccount: "new_user@example.com", Password: "SuperSecretPassword"})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/createaccount", c.lastPath, "path used is correct")
	assert.Equal(t, "new_user@example.com", createaccount.EmailAccount, "emailaccount is correct")
	assert.Equal(t, "new_user@example.com", createaccount.DisplayName, "displayname is correct")
	assert.Equal(t, 200, createaccount.Quota.Max, "quota.max is correct")
	assert.Equal(t, "MB", createaccount.Quota.Unit, "quota.unit is correct")
	assert.Equal(t, 3, createaccount.AntiSpamLevel, "antispamlevel is correct")
	assert.Equal(t, "yes", createaccount.AntiVirus, "antivirus is correct")
	assert.Equal(t, "no", createaccount.AutoRespond, "autorespond is correct")
	assert.Equal(t, "", createaccount.AutoRespondMessage, "autorespondmessage is correct")
	assert.Equal(t, "yes", createaccount.AutoRespondSaveEmail, "autorespondsaveemail is correct")
	assert.Equal(t, "no", createaccount.RejectSpam, "rejectspam is correct")
	assert.Equal(t, "2019-11-29T21:31:28+01:00", createaccount.Created, "created is correct")
	assert.Equal(t, "", createaccount.Modified, "modified is correct")
}

func TestQuota(t *testing.T) {
	c := &mockClient{body: `{"response":{"quota":{"emailaccount":"user@example.com","used":{"amount":0,"unit":"MB"},"total":{"max":400,"unit":"MB"}}}}`}

	s := EmailService{client: c}

	quota, _ := s.Quota(context.Background(), "user@example.com")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/quota", c.lastPath, "path used is correct")
	assert.Equal(t, "user@example.com", quota.EmailAccount, "emailaccount is correct")
	assert.Equal(t, 0, quota.Used.Amount, "used amount is correct")
	assert.Equal(t, "MB", quota.Used.Unit, "used unit is correct")
	assert.Equal(t, 400, quota.Total.Max, "total max is correct")
	assert.Equal(t, "MB", quota.Total.Unit, "total unit is correct")
}

func TestCreateAlias(t *testing.T) {
	c := &mockClient{body: `{"response":{"alias":{"emailalias":"alias@example.com","displayname":"alias@example.com","goto":"user@example.com"}}}`}

	s := EmailService{client: c}

	alias, _ := s.CreateAlias(context.Background(), EmailAliasParams{EmailAlias: "alias@example.com", GoTo: "user@example.com"})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "email/createalias", c.lastPath, "path used is correct")
	assert.Equal(t, "alias@example.com", alias.EmailAlias, "emailalias is correct")
	assert.Equal(t, "alias@example.com", alias.DisplayName, "displayname is correct")
	assert.Equal(t, "user@example.com", alias.GoTo, "goto is correct")
}
