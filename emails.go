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

type EmailAccountQuotaUsed struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

type EmailQuota struct {
	EmailAccount string                `json:"emailaccount"`
	Used         EmailAccountQuotaUsed `json:"used"`
	Total        EmailAccountQuota     `json:"total"`
}

// ListEmailsParams is used for filtering when listing emails for a domain.
type ListEmailsParams struct {
	Filter string `json:"filter,omitempty"`
}

// EditAccountParams is used for updating the different values on an email account.
type EditAccountParams struct {
	AntiSpamLevel      int    `json:"antispamlevel,omitempty"`
	AntiVirus          string `json:"antivirus,omitempty"`
	Password           string `json:"password,omitempty"`
	AutoRespond        string `json:"autorespond,omitempty"`
	AutoRespondMessage string `json:"autorespondmessage,omitempty"`
	Quota              int    `json:"quota,omitempty"`
	RejectSpam         string `json:"rejectspam,omitempty"`
}

// CreateAccountParams is used for creating new email accounts.
type CreateAccountParams struct {
	EmailAccount       string `json:"emailaccount"`
	Password           string `json:"password"`
	AntiSpamLevel      int    `json:"antispamlevel,omitempty"`
	AntiVirus          string `json:"antivirus,omitempty"`
	AutoRespond        string `json:"autorespond,omitempty"`
	AutoRespondMessage string `json:"autorespondmessage,omitempty"`
	Quota              int    `json:"quota,omitempty"`
	RejectSpam         string `json:"rejectspam,omitempty"`
}

// EmailAliasParams is used for creating new email aliases as well as editing already existing ones.
type EmailAliasParams struct {
	EmailAlias string `json:"emailalias"`
	GoTo       string `json:"goto"`
}

type EmailCostEntity struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type EmailCostsEntry struct {
	Amount int             `json:"amount"`
	Cost   EmailCostEntity `json:"cost"`
}

type EmailQuotaPricelist struct {
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	Unit       string `json:"unit"`
	FreeAmount int    `json:"freeamount"`
}

type EmailAccountsPricelist struct {
	Amount     int    `json:"amount"`
	Currency   string `json:"currency"`
	Unit       string `json:"unit"`
	FreeAmount int    `json:"freeamount"`
}

type EmailCostsContainer struct {
	Quota    EmailCostsEntry `json:"quota"`
	Accounts EmailCostsEntry `json:"accounts"`
}

type EmailPricelistContainer struct {
	Quota    EmailQuotaPricelist    `json:"quota"`
	Accounts EmailAccountsPricelist `json:"accounts"`
}

type EmailCosts struct {
	Costs     EmailCostsContainer     `json:"costs"`
	PriceList EmailPricelistContainer `json:"pricelist"`
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

// List Gets a list of all accounts and aliases of a domain with full details.
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

// EditAccount allows you to Edit an email account's parameters.
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

// Delete allows you to remove an email account or alias.
func (em *EmailService) Delete(context context.Context, email string) error {
	return em.client.post(context, "email/delete", nil, struct {
		Email string `json:"email"`
	}{email})
}

// CreateAccount allows you to create an email account.
func (em *EmailService) CreateAccount(context context.Context, params CreateAccountParams) (*EmailAccount, error) {
	data := struct {
		Response struct {
			EmailAccount EmailAccount
		}
	}{}

	err := em.client.post(context, "email/createaccount", &data, struct {
		CreateAccountParams
	}{params})

	return &data.Response.EmailAccount, err
}

// Quota returns the used and total quota for an email account.
func (em *EmailService) Quota(context context.Context, emailaccount string) (*EmailQuota, error) {
	data := struct {
		Response struct {
			Quota EmailQuota
		}
	}{}

	err := em.client.post(context, "email/quota", &data, struct {
		EmailAccount string `json:"emailaccount"`
	}{emailaccount})

	return &data.Response.Quota, err
}

// CreateAlias sets up a new alias.
func (em *EmailService) CreateAlias(context context.Context, params EmailAliasParams) (*EmailAlias, error) {
	data := struct {
		Response struct {
			Alias EmailAlias
		}
	}{}

	err := em.client.post(context, "email/createalias", &data, struct {
		EmailAliasParams
	}{params})

	return &data.Response.Alias, err
}

// EditAlias updates an already existing alias.
func (em *EmailService) EditAlias(context context.Context, params EmailAliasParams) (*EmailAlias, error) {
	data := struct {
		Response struct {
			Alias EmailAlias
		}
	}{}

	err := em.client.post(context, "email/editalias", &data, struct {
		EmailAliasParams
	}{params})

	return &data.Response.Alias, err
}

// Costs returns the email related costs and the current pricelist.
func (em *EmailService) Costs(context context.Context) (*EmailCosts, error) {
	data := struct {
		Response struct {
			Costs     EmailCostsContainer
			PriceList EmailPricelistContainer
		}
	}{}

	err := em.client.get(context, "email/costs", &data)

	return &EmailCosts{Costs: data.Response.Costs, PriceList: data.Response.PriceList}, err
}
