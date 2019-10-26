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

type EmailOverview struct {
	Summary EmailOverviewSummary  `json:"summary"`
	Domains []EmailOverviewDomain `json:"domains"`
	Meta    EmailOverviewMeta     `json:"meta"`
}

// OverviewParams is used for Filtering and/or Paging on the overview endpoint.
type OverviewParams struct {
	Filter string `json:"filter,omitempty"`
	Page   int    `json:"page,omitempty"`
}

type EmailGlobalQuota struct {
	Usage int `json:"usage"`
	Max   int `json:"max"`
}

type GlobalQuotaParams struct {
	GlobalQuota int `json:"globalquota,omitempty"`
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
