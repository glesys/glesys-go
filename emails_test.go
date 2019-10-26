package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailsOverview(t *testing.T) {
	c := &mockClient{body: `{"response":{"overview":{"summary":{"accounts":0,"maxaccounts":2000,"aliases":0,"maxaliases":10000},"domains":[{"domainname":"example.com","displayname":"example.com","accounts":0,"aliases":0}],"meta":{"page":1,"total":1,"perpage":1}}}}`}
	s := EmailService{client: c}

	emailoverview, _ := s.Overview(context.Background(), EmailOverviewParams{})
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

	emailoverview, _ := s.Overview(context.Background(), EmailOverviewParams{Filter: "example.com"})
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
	emailoverview, _ := s.Overview(context.Background(), EmailOverviewParams{Page: 2})
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

	emailoverview, _ := s.Overview(context.Background(), EmailOverviewParams{Filter: "example.com", Page: 1})
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
