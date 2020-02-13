package glesys

import (
	"context"
)

// DnsDomainService provides functions to interact withDnsDomains
type DnsDomainService struct {
	client clientInterface
}

//DnsDomain represents a domain
type DnsDomain struct {
	Name                  string           `json:"domainname"`
	Available             bool             `json:"available,omitempty"`
	CreateTime            string           `json:"createtime,omitempty"`
	DisplayName           string           `json:"displayname,omitempty"`
	Expire                int              `json:"expire,omitempty"`
	Minimum               int              `json:"minimum,omitempty"`
	Prices                []DnsDomainPrice `json:"prices,omitempty"`
	PrimaryNameServer     string           `json:"primarynameserver,omitempty"`
	RecordCount           int              `json:"recordcount,omitempty"`
	Refresh               int              `json:"refresh,omitempty"`
	RegistrarInfo         RegistrarInfo    `json:"registrarinfo,omitempty"`
	ResponsiblePerson     string           `json:"responsibleperson,omitempty"`
	Retry                 int              `json:"retry,omitempty"`
	TTL                   int              `json:"ttl,omitempty"`
	UsingGlesysNameserver string           `json:"usingglesysnameserver,omitempty"`
}

type DnsDomainPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Years    int    `json:"years"`
}

// AddDnsDomainParams - used for adding existing domains to GleSYS
// use CreateRecords = false to not create additional records.
type AddDnsDomainParams struct {
	Name              string `json:"domainname"`
	CreateRecords     bool   `json:"createrecords,omitempty"`
	Expire            int    `json:"expire,omitempty"`
	Minimum           int    `json:"minimum,omitempty"`
	PrimaryNameServer string `json:"primarynameserver,omitempty"`
	Refresh           int    `json:"refresh,omitempty"`
	ResponsiblePerson string `json:"responsibleperson,omitempty"`
	Retry             int    `json:"retry,omitempty"`
	TTL               int    `json:"ttl,omitempty"`
}

// EditDnsDomainParams - used when editing domain parameters
type EditDnsDomainParams struct {
	Name              string `json:"domainname"`
	Expire            int    `json:"expire,omitempty"`
	Minimum           int    `json:"minimum,omitempty"`
	PrimaryNameServer string `json:"primarynameserver,omitempty"`
	Refresh           int    `json:"refresh,omitempty"`
	ResponsiblePerson string `json:"responsibleperson,omitempty"`
	Retry             int    `json:"retry,omitempty"`
	TTL               int    `json:"ttl,omitempty"`
}

// RegistrarInfo contains information about the registrar for the domain
type RegistrarInfo struct {
	AutoRenew        string `json:"autorenew"`
	State            string `json:"state"`
	StateDescription string `json:"statedescription,omitempty"`
	Expire           string `json:"expire,omitempty"`
	TLD              string `json:"tld,omitempty"`
	InvoiceNumber    string `json:"invoicenumber,omitempty"`
}

// RegisterDnsDomainParams - parameters used when registering a domain
type RegisterDnsDomainParams struct {
	Name               string `json:"domainname"`
	Address            string `json:"address"`
	City               string `json:"city"`
	Country            string `json:"country"`
	Email              string `json:"email"`
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	OrganizationNumber int    `json:"organizationnumber"`
	Organization       string `json:"organization"`
	PhoneNumber        string `json:"phonenumber"`
	ZipCode            string `json:"zipcode"`

	FaxNumber string `json:"fax,omitempty"`
	NumYears  int    `json:"numyears,omitempty"`
}

// DeleteDnsDomainParams - parameters for deleting a domain from the dns system.
// Set ForceDeleteEmail to true, to delete a domain AND email accounts for the domain.
type DeleteDnsDomainParams struct {
	Name             string `json:"domainname"`
	ForceDeleteEmail string `json:"forcedeleteemail,omitempty"`
}

// RenewDnsDomainParams - parameters to send when renewing a domain.
type RenewDnsDomainParams struct {
	Name     string `json:"domainname"`
	NumYears int    `json:"numyears"`
}

// SetAutoRenewParams - parameters to send for renewing a domain automatically.
type SetAutoRenewParams struct {
	Name         string `json:"domainname"`
	SetAutoRenew string `json:"setautorenew"`
}

// DnsDomainRecord - data in the domain
type DnsDomainRecord struct {
	DomainName string `json:"domainname"`
	Data       string `json:"data"`
	Host       string `json:"host"`
	RecordID   int    `json:"recordid"`
	TTL        int    `json:"ttl"`
	Type       string `json:"type"`
}

// AddRecordParams - parameters for updating domain records
type AddRecordParams struct {
	DomainName string `json:"domainname"`
	Data       string `json:"data"`
	Host       string `json:"host"`
	Type       string `json:"type"`
	TTL        int    `json:"ttl,omitempty"`
}

// UpdateRecordParams - parameters for updating domain records
type UpdateRecordParams struct {
	RecordID int    `json:"recordid"`
	Data     string `json:"data,omitempty"`
	Host     string `json:"host,omitempty"`
	Type     string `json:"type,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
}

// ChangeNameserverParams - parameters for updating the nameservers for domain
type ChangeNameserverParams struct {
	DomainName string `json:"domainname"`
	NS1        string `json:"NS1"`
	NS2        string `json:"NS2"`
	NS3        string `json:"NS3,omitempty"`
	NS4        string `json:"NS4,omitempty"`
}

// Available - checks if the domain is available
func (s *DnsDomainService) Available(context context.Context, search string) (*[]DnsDomain, error) {
	data := struct {
		Response struct {
			Domain []DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/available", &data, search)
	return &data.Response.Domain, err
}

// AddDnsDomain - add an existing domain to your GleSYS account
func (s *DnsDomainService) AddDnsDomain(context context.Context, params AddDnsDomainParams) (*DnsDomain, error) {
	data := struct {
		Response struct {
			Domain DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/add", &data, params)
	return &data.Response.Domain, err
}

// Delete - deletes a domain from the dns system
func (s *DnsDomainService) Delete(context context.Context, params DeleteDnsDomainParams) error {
	return s.client.post(context, "domain/delete", nil, params)
}

// Details - return details about the domain
func (s *DnsDomainService) Details(context context.Context, domainname string) (*DnsDomain, error) {
	data := struct {
		Response struct {
			Domain DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/details", &data, struct {
		Name string `json:"domainname"`
	}{domainname})
	return &data.Response.Domain, err
}

// Edit - edit domain parameters
func (s *DnsDomainService) Edit(context context.Context, params EditDnsDomainParams) (*DnsDomain, error) {
	data := struct {
		Response struct {
			Domain DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/edit", &data, params)
	return &data.Response.Domain, err
}

// List - return a list of all domains in your account
func (s *DnsDomainService) List(context context.Context) (*[]DnsDomain, error) {
	data := struct {
		Response struct {
			Domains []DnsDomain
		}
	}{}
	err := s.client.get(context, "domain/list", &data)
	return &data.Response.Domains, err
}

// Register - Register a domain
func (s *DnsDomainService) Register(context context.Context, params RegisterDnsDomainParams) (*DnsDomain, error) {
	data := struct {
		Response struct {
			Domain DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/register", &data, params)
	return &data.Response.Domain, err
}

// Renew - Renew a domain
func (s *DnsDomainService) Renew(context context.Context, params RenewDnsDomainParams) (*DnsDomain, error) {
	data := struct {
		Response struct {
			Domain DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/renew", &data, params)
	return &data.Response.Domain, err
}

// SetAutoRenew - Set a domain to renew automatically
func (s *DnsDomainService) SetAutoRenew(context context.Context, params SetAutoRenewParams) (*DnsDomain, error) {
	data := struct {
		Response struct {
			Domain DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/setautorenew", &data, params)
	return &data.Response.Domain, err
}

// Transfer - Transfer a domain
func (s *DnsDomainService) Transfer(context context.Context, params RegisterDnsDomainParams) (*DnsDomain, error) {
	data := struct {
		Response struct {
			Domain DnsDomain
		}
	}{}
	err := s.client.post(context, "domain/transfer", &data, params)
	return &data.Response.Domain, err
}

// ListRecords - return a list of all records for domain
func (s *DnsDomainService) ListRecords(context context.Context, domainname string) (*[]DnsDomainRecord, error) {
	data := struct {
		Response struct {
			Records []DnsDomainRecord
		}
	}{}
	err := s.client.post(context, "domain/listrecords", &data, struct {
		Name string `json:"domainname"`
	}{domainname})
	return &data.Response.Records, err
}

// AddRecord - add a domain record
func (s *DnsDomainService) AddRecord(context context.Context, params AddRecordParams) (*DnsDomainRecord, error) {
	data := struct {
		Response struct {
			Record DnsDomainRecord
		}
	}{}
	err := s.client.post(context, "domain/addrecord", &data, params)
	return &data.Response.Record, err
}

// UpdateRecord - update a domain record
func (s *DnsDomainService) UpdateRecord(context context.Context, params UpdateRecordParams) (*DnsDomainRecord, error) {
	data := struct {
		Response struct {
			Record DnsDomainRecord
		}
	}{}
	err := s.client.post(context, "domain/updaterecord", &data, params)
	return &data.Response.Record, err
}

// DeleteRecord deletes a record
func (s *DnsDomainService) DeleteRecord(context context.Context, recordID int) error {
	return s.client.post(context, "domain/deleterecord", nil, struct {
		RecordID int `json:"recordid"`
	}{recordID})
}

// ChangeNameservers - change the nameservers for domain
func (s *DnsDomainService) ChangeNameservers(context context.Context, params ChangeNameserverParams) error {
	return s.client.post(context, "domain/changenameservers", nil, params)
}
