package glesys

import (
	"context"
	"strconv"
	"strings"
)

// EmailDomainService provides functions to interact with the Email api
type EmailDomainService struct {
	client clientInterface
}

// EmailOverviewDomain represents a single email domain
type EmailOverviewDomain struct {
	DomainName  string `json:"domainname"`
	DisplayName string `json:"displayname"`
	Accounts    int    `json:"accounts"`
	Aliases     int    `json:"aliases"`
}

// EmailOverviewSummary represents limits for the current project
type EmailOverviewSummary struct {
	Accounts    int `json:"accounts"`
	MaxAccounts int `json:"maxaccounts"`
	Aliases     int `json:"aliases"`
	MaxAliases  int `json:"maxaliases"`
}

// EmailOverviewMeta is used to paginate the results
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

// EmailAccount represents a single email account
type EmailAccount struct {
	EmailAccount         string `json:"emailaccount"`
	DisplayName          string `json:"displayname"`
	Password             string `json:"password,omitempty"`
	QuotaInGiB           int    `json:"quotaingib"`
	AntiSpamLevel        int    `json:"antispamlevel"`
	AntiVirus            string `json:"antivirus"`
	AutoRespond          string `json:"autorespond"`
	AutoRespondMessage   string `json:"autorespondmessage,omitempty"`
	AutoRespondSaveEmail string `json:"autorespondsaveemail"`
	RejectSpam           string `json:"rejectspam"`
	Created              string `json:"created"`
	Modified             string `json:"modified,omitempty"`
}

// EmailAlias represents a single email alias
type EmailAlias struct {
	EmailAlias  string `json:"emailalias"`
	DisplayName string `json:"displayname"`
	GoTo        string `json:"goto"`
}

// EmailList holds arrays of email accounts and aliases
type EmailList struct {
	EmailAccounts []EmailAccount `json:"emailaccounts"`
	EmailAliases  []EmailAlias   `json:"emailaliases"`
}

// EmailQuota represents a quota object for a single email account
type EmailQuota struct {
	EmailAccount string `json:"emailaccount"`
	UsedInMiB    int    `json:"usedquotainmib"`
	QuotaInGiB   int    `json:"quotaingib"`
}

// ListEmailsParams is used for filtering when listing emails for a domain.
type ListEmailsParams struct {
	Filter string `json:"filter,omitempty"`
}

// EditAccountParams is used for updating the different values on an email account.
type EditAccountParams struct {
	AntiSpamLevel      int    `json:"antispamlevel,omitempty"`
	AntiVirus          string `json:"antivirus,omitempty"`
	AutoRespond        string `json:"autorespond,omitempty"`
	AutoRespondMessage string `json:"autorespondmessage,omitempty"`
	QuotaInGiB         int    `json:"quotaingib,omitempty"`
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
	QuotaInGiB         int    `json:"quotaingib,omitempty"`
	RejectSpam         string `json:"rejectspam,omitempty"`
}

// EmailAliasParams is used for creating new email aliases as well as editing already existing ones.
type EmailAliasParams struct {
	EmailAlias string `json:"emailalias"`
	GoTo       string `json:"goto"`
}

// EmailCostEntity represents the amount and currency or a cost
type EmailCostEntity struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// EmailCostsEntry represents a cost object
type EmailCostsEntry struct {
	Amount float64         `json:"amount"`
	Cost   EmailCostEntity `json:"cost"`
}

// EmailQuotaPricelist is the pricelist for quota options
type EmailQuotaPricelist struct {
	Amount     string  `json:"amount"`
	Currency   string  `json:"currency"`
	Unit       string  `json:"unit"`
	FreeAmount float64 `json:"freeamount"`
}

// EmailAccountsPricelist is the pricelist for email
type EmailAccountsPricelist struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Unit       string  `json:"unit"`
	FreeAmount float64 `json:"freeamount"`
}

// EmailCostsContainer contains quota and accounts for email costs
type EmailCostsContainer struct {
	Quota    EmailCostsEntry `json:"quota"`
	Accounts EmailCostsEntry `json:"accounts"`
}

// EmailPricelistContainer contains quota and accounts for pricelist
type EmailPricelistContainer struct {
	Quota    EmailQuotaPricelist    `json:"quota"`
	Accounts EmailAccountsPricelist `json:"accounts"`
}

// EmailCosts represents a email cost object
type EmailCosts struct {
	Costs     EmailCostsContainer     `json:"costs"`
	PriceList EmailPricelistContainer `json:"pricelist"`
}

// Overview fetches a summary of the email accounts and domains on the account.
func (em *EmailDomainService) Overview(context context.Context, params OverviewParams) (*EmailOverview, error) {

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

// List Gets a list of all accounts and aliases of a domain with full details.
func (em *EmailDomainService) List(context context.Context, domain string, params ListEmailsParams) (*EmailList, error) {
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
func (em *EmailDomainService) EditAccount(context context.Context, emailAccount string, params EditAccountParams) (*EmailAccount, error) {
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
func (em *EmailDomainService) Delete(context context.Context, email string) error {
	return em.client.post(context, "email/delete", nil, struct {
		Email string `json:"email"`
	}{email})
}

// CreateAccount allows you to create an email account.
func (em *EmailDomainService) CreateAccount(context context.Context, params CreateAccountParams) (*EmailAccount, error) {
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
func (em *EmailDomainService) Quota(context context.Context, emailaccount string) (*EmailQuota, error) {
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
func (em *EmailDomainService) CreateAlias(context context.Context, params EmailAliasParams) (*EmailAlias, error) {
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
func (em *EmailDomainService) EditAlias(context context.Context, params EmailAliasParams) (*EmailAlias, error) {
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
func (em *EmailDomainService) Costs(context context.Context) (*EmailCosts, error) {
	data := struct {
		Response struct {
			Costs     EmailCostsContainer
			PriceList EmailPricelistContainer
		}
	}{}

	err := em.client.get(context, "email/costs", &data)

	return &EmailCosts{Costs: data.Response.Costs, PriceList: data.Response.PriceList}, err
}
