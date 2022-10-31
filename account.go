package glesys

import "context"

// AccountService provides functions to interact with Accounts
type AccountService struct {
	client clientInterface
}

// AccountContact
type AccountContact struct {
	Companyname      string `json:"companyname"`
	Contactperson    string `json:"contactperson"`
	Invoicereference string `json:"invoicereference"`
	Address          string `json:"address"`
	Zipcode          string `json:"zipcode"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Nationalidnumber string `json:"nationalidnumber"`
	Phonenumber      string `json:"phonenumber"`
	Email            string `json:"email"`
}

// AccountInvoice
type AccountInvoice struct {
	Invoiceaddress string `json:"invoiceaddress"`
	Invoicezipcode string `json:"invoicezipcode"`
	Invoicecity    string `json:"invoicecity"`
	Invoicecountry string `json:"invoicecountry"`
}

// AccountInvoiceSettings
type AccountInvoiceSettings struct {
	Separateinvoiceaddress string `json:"separateinvoiceaddress"`
}

// AccountContactinfo
type AccountContactinfo struct {
	Contact  AccountContact         `json:"contact"`
	Invoice  AccountInvoice         `json:"invoice"`
	Settings AccountInvoiceSettings `json:"settings"`
}

// AccountCustomer contains customer contact information.
type AccountCustomer struct {
	Customernumber int                `json:"customernumber"`
	Contactinfo    AccountContactinfo `json:"contactinfo"`
}

// AccountServicelimits
type AccountServicelimits struct {
	Maxnumservers              int64 `json:"maxnumservers"`
	Maxnumipv4                 int64 `json:"maxnumipv4"`
	Maxnumipv6                 int64 `json:"maxnumipv6"`
	Maxarchivetotalsize        int64 `json:"maxarchivetotalsize"`
	Maxnumemailaccounts        int64 `json:"maxnumemailaccounts"`
	Emailglobalquota           int   `json:"emailglobalquota"`
	Domainregistrationsallowed int64 `json:"domainregistrationsallowed"`
	Domaintransfersallowed     int64 `json:"domaintransfersallowed"`
}

// Services contains information about enabled services for the account.
type AccountServices struct {
	Server   string `json:"server"`
	IP       string `json:"ip"`
	Domain   string `json:"domain"`
	Archive  string `json:"archive"`
	Email    string `json:"email"`
	Invoice  string `json:"invoice"`
	Customer string `json:"customer"`
	Account  string `json:"account"`
	API      string `json:"api"`
	Vpn      string `json:"vpn"`
}

// Accountinfo contains information about the account.
type Accountinfo struct {
	Accountname    string               `json:"accountname"`
	Customernumber int                  `json:"customernumber"`
	Customer       AccountCustomer      `json:"customer"`
	Currency       string               `json:"currency"`
	Unpaidinvoices string               `json:"unpaidinvoices"`
	Apienabled     string               `json:"apienabled"`
	Servicelimits  AccountServicelimits `json:"servicelimits"`
	Services       AccountServices      `json:"services"`
}

// Info about the current Account
func (s *AccountService) Info(context context.Context) (*Accountinfo, error) {
	data := struct {
		Response struct {
			Accountinfo Accountinfo
		}
	}{}
	err := s.client.post(context, "account/info", &data, nil)
	return &data.Response.Accountinfo, err
}
