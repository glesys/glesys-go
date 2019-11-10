package glesys

import (
	"context"
	"strconv"
	"strings"
)

// EmailService provides functions to interact with the Email api
type EmailService struct {
	client clientInterface
}

type EmailOverviewDomain struct {
	DomainName  string `json:"domainname"`
	DisplayName string `json:"displayname"`
	Accounts    int    `json:"accounts"`
	Aliases     int    `json:"aliases"`
}

type EmailOverviewSummary struct {
	Accounts    int `json:"accounts"`
	MaxAccounts int `json:"maxaccounts"`
	Aliases     int `json:"aliases"`
	MaxAliases  int `json:"maxaliases"`
}

type EmailOverviewMeta struct {
	Page    int `json:"page"`
	Total   int `json:"total"`
	PerPage int `json:"perpage"`
}

// EmailOverview represents the overview of the accounts email setups.
type EmailOverview struct {
	Summary EmailOverviewSummary  `json:"summary"`
	Domains []EmailOverviewDomain `json:"domains"`
	Meta    EmailOverviewMeta     `json:"meta"`
}

// OverviewParams is used for filtering and/or paging on the overview endpoint.
type OverviewParams struct {
	Filter string `json:"filter,omitempty"`
	Page   int    `json:"page,omitempty"`
}

type EmailGlobalQuota struct {
	Usage int `json:"usage"`
	Max   int `json:"max"`
}

// GlobalQuotaParams is used for updating the global quota for email accounts.
type GlobalQuotaParams struct {
	GlobalQuota int `json:"globalquota,omitempty"`
}

type EmailAccountQuota struct {
	Max  int    `json:"max"`
	Unit string `json:"unit"`
}

type EmailAccount struct {
	EmailAccount         string            `json:"emailaccount"`
	DisplayName          string            `json:"displayname"`
	Quota                EmailAccountQuota `json:"quota"`
	AntiSpamLevel        int               `json:"antispamlevel"`
	AntiVirus            string            `json:"antivirus"`
	AutoRespond          string            `json:"autorespond"`
	AutoRespondMessage   string            `json:"autorespondmessage,omitempty"`
	AutoRespondSaveEmail string            `json:"autorespondsaveemail"`
	RejectSpam           string            `json:"rejectspam"`
	Created              string            `json:"created"`
	Modified             string            `json:"modified,omitempty"`
}

type EmailAlias struct {
	EmailAlias  string `json:"emailalias"`
	DisplayName string `json:"displayname"`
	GoTo        string `json:"goto"`
}

type EmailList struct {
	EmailAccounts []EmailAccount `json:"emailaccounts"`
	EmailAliases  []EmailAlias   `json:"emailaliases"`
}

// ListEmailsParams is use for filtering when listing emails for a domain.
type ListEmailsParams struct {
	Filter string `json:"filter,omitempty"`
}

type EditAccountParams struct {
	AntiSpamLevel      int    `json:"antispamlevel,omitempty"`
	AntiVirus          string `json:"antivirus,omitempty"`
	Password           string `json:"password,omitempty"`
	AutoRespond        string `json:"autorespond,omitempty"`
	AutoRespondMessage string `json:"autorespondmessage,omitempty"`
	Quota              int    `json:"quota,omitempty"`
	RejectSpam         string `json:"rejectspam,omitempty"`
}

// Overview fetches a summary of the email accounts and domains on the account.
func (em *EmailService) Overview(context context.Context, params OverviewParams) (*EmailOverview, error) {

	// String builder for creating the suffix based on the page and/or filter.
	var suffixbuilder strings.Builder

	// Only add suffix paths to the builder if the struct fields does not contain the default value.
	if params.Filter != "" {
		suffixbuilder.WriteString("/filter/" + params.Filter)
	}

	if params.Page != 0 {
		suffixbuilder.WriteString("/page/" + strconv.Itoa(params.Page))
	}

	// Extract the string from the builder.
	// If nothing has been added then this will be an empty string.
	pathsuffix := suffixbuilder.String()

	data := struct {
		Response struct {
			Overview EmailOverview
		}
	}{}

	// Append the suffix to the endpoint
	err := em.client.get(context, "email/overview"+pathsuffix, &data)
	return &data.Response.Overview, err
}

// GlobalQuoata enables the user to set and get the global quota.
func (em *EmailService) GlobalQuota(context context.Context, params GlobalQuotaParams) (*EmailGlobalQuota, error) {
	data := struct {
		Response struct {
			GlobalQuota EmailGlobalQuota
		}
	}{}

	err := em.client.post(context, "email/globalquota", &data, struct {
		GlobalQuotaParams
	}{params})

	return &data.Response.GlobalQuota, err
}

func (em *EmailService) List(context context.Context, domain string, params ListEmailsParams) (*EmailList, error) {
	data := struct {
		Response struct {
			List EmailList
		}
	}{}

	err := em.client.post(context, "email/list", &data, struct {
		ListEmailsParams
		DomainName string `json:"domainname"`
	}{params, domain})

	return &data.Response.List, err
}

//EditAccount allows you to Edit an email account's parameters.
func (em *EmailService) EditAccount(context context.Context, emailAccount string, params EditAccountParams) (*EmailAccount, error) {
	data := struct {
		Response struct {
			EmailAccount EmailAccount
		}
	}{}

	err := em.client.post(context, "email/editaccount", &data, struct {
		EditAccountParams
		EmailAccount string `json:"emailaccount"`
	}{params, emailAccount})

	return &data.Response.EmailAccount, err
}
