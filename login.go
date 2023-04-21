package glesys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Login is used for login data
type Login struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient httpClientInterface
	Accounts   []Customer
	Customers  []Customer
	APIKey     string
	Username   string

	Users *UserService
}

type UserService struct {
	client clientInterface
}

// LoginParams are used when calling user/login
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Otp      string `json:"otp,omitempty"`
}

// LoginDetailsResponse represents the result of a successful login.
type LoginDetailsResponse struct {
	Username  string     `json:"username"`
	APIKey    string     `json:"apikey"`
	Accounts  []Customer `json:"accounts,omitempty"`
	Customers []Customer `json:"customers,omitempty"`
}

// Customer represents a customer/organization
type Customer struct {
	CustomerNumber string   `json:"customernumber"`
	Description    string   `json:"description,omitempty"`
	Roles          []string `json:"roles"`
}

// CustomerProject describes a project in GleSYS
type CustomerProject struct {
	Accountname    string `json:"accountname,omitempty"`
	Name           string `json:"name,omitempty"`
	Color          string `json:"color,omitempty"`
	Currency       string `json:"currency,omitempty"`
	Lockedreason   string `json:"lockedreason,omitempty"`
	Access         string `json:"access,omitempty"`
	Customernumber int    `json:"customernumber,omitempty"`
}

type UserOrganization struct {
	ID                                int      `json:"id,omitempty"`
	State                             string   `json:"state,omitempty"`
	FeatureFlags                      []string `json:"featureflags,omitempty"`
	Verification                      string   `json:"verification,omitempty"`
	Type                              string   `json:"type,omitempty"`
	IsOwner                           string   `json:"isowner,omitempty"`
	ContactPerson                     string   `json:"contactperson,omitempty"`
	NationalIdnumber                  string   `json:"nationalidnumber,omitempty"`
	Name                              string   `json:"name,omitempty"`
	Address                           string   `json:"address,omitempty"`
	City                              string   `json:"city,omitempty"`
	ZipCode                           string   `json:"zipcode,omitempty"`
	Country                           string   `json:"country,omitempty"`
	SeparateInvoiceAddress            string   `json:"separateinvoiceaddress,omitempty"`
	InvoiceReference                  string   `json:"invoicereference,omitempty"`
	InvoiceAddress                    string   `json:"invoiceaddress,omitempty"`
	InvoiceCity                       string   `json:"invoicecity,omitempty"`
	InvoiceZipcode                    string   `json:"invoicezipcode,omitempty"`
	InvoiceCountry                    string   `json:"invoicecountry,omitempty"`
	Email                             string   `json:"email,omitempty"`
	Phonenumber                       string   `json:"phonenumber,omitempty"`
	Invoicedeliverymethod             string   `json:"invoicedeliverymethod,omitempty"`
	Billingemailoverrides             []string `json:"billingemailoverrides,omitempty"`
	CurrentBillingEmails              []string `json:"currentbillingemails,omitempty"`
	CanAccessBilling                  string   `json:"canaccessbilling,omitempty"`
	ExpiredInvoices                   string   `json:"expiredinvoices,omitempty"`
	UnpaidInvoices                    string   `json:"unpaidinvoices,omitempty"`
	VatNumber                         string   `json:"vatnumber,omitempty"`
	PaymentMethod                     string   `json:"paymentmethod,omitempty"`
	PaymentCard                       int      `json:"paymentcard,omitempty"`
	AllowedPaymentMethods             []string `json:"allowedpaymentmethods,omitempty"`
	PaymentTermsnetdays               int      `json:"paymenttermsnetdays,omitempty"`
	CanPayInvoicesManuallyUsingPaypal bool     `json:"canpayinvoicesmanuallyusingpaypal,omitempty"`
	SlaCost                           int      `json:"slacost,omitempty"`
	SlaLevel                          string   `json:"slalevel,omitempty"`
	SlaPincode                        string   `json:"slapincode,omitempty"`
	SlaPhonenumber                    string   `json:"slaphonenumber,omitempty"`
}

func (l *Login) do(request *http.Request, v interface{}) error {
	response, err := l.httpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return handleResponseError(response)
	}

	return parseResponseBody(response, v)
}

func (l *Login) get(ctx context.Context, path string, v interface{}) error {
	request, err := l.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return err
	}
	return l.do(request, v)
}

func (l *Login) post(ctx context.Context, path string, v interface{}, params interface{}) error {
	request, err := l.newRequest(ctx, "POST", path, params)
	if err != nil {
		return err
	}
	return l.do(request, v)
}

func (l *Login) newRequest(ctx context.Context, method, path string, params interface{}) (*http.Request, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	if l.BaseURL != nil {
		u = l.BaseURL.ResolveReference(u)
	}

	buffer := new(bytes.Buffer)

	if params != nil {
		err = json.NewEncoder(buffer).Encode(params)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequestWithContext(ctx, method, u.String(), buffer)
	if err != nil {
		return nil, err
	}

	userAgent := strings.TrimSpace(fmt.Sprintf("%s glesys-go/%s", l.UserAgent, version))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", userAgent)
	request.SetBasicAuth(l.Username, l.APIKey)

	return request, nil
}

// SetBaseURL can be used to set a custom BaseURL
func (l *Login) SetBaseURL(bu string) error {
	url, err := url.Parse(bu)
	if err != nil {
		return err
	}
	l.BaseURL = url
	return nil
}

func NewLogin(useragent string) *Login {
	BaseURL, _ := url.Parse("https://api.glesys.com")

	l := &Login{
		BaseURL:    BaseURL,
		httpClient: http.DefaultClient,
		UserAgent:  strings.TrimSpace(fmt.Sprintf("%s glesys-go/%s", useragent, version)),
	}

	l.Users = &UserService{client: l}

	return l
}

func (l *UserService) DoOTPLogin(ctx context.Context, username, password, otp string) (*LoginDetailsResponse, error) {
	params := &LoginParams{
		username,
		password,
		otp,
	}

	data := struct {
		Response struct {
			Login LoginDetailsResponse
		}
	}{}
	err := l.client.post(ctx, "user/login", &data, params)

	return &data.Response.Login, err
}

func (l *UserService) ListOrganizations(ctx context.Context) (*[]UserOrganization, error) {
	data := struct {
		Response struct {
			Organizations []UserOrganization
		}
	}{}

	err := l.client.post(ctx, "user/listorganizations", &data, nil)

	return &data.Response.Organizations, err
}

func (l *UserService) ListCustomerProjects(ctx context.Context, organizationNumber string) (*[]CustomerProject, error) {
	data := struct {
		Response struct {
			Projects []CustomerProject
		}
	}{}

	err := l.client.post(ctx, "customer/listprojects", &data, struct {
		Organizationnumber string `json:"organizationnumber"`
	}{organizationNumber})

	return &data.Response.Projects, err
}
