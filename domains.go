package glesys

import (
	"context"
)

// DomainService provides functions to interact with Domains
type DomainService struct {
	client clientInterface
}

// Domain represents a domain
type Domain struct {
	Name                  string        `json:"domainname"`
	Available             bool          `json:"available,omitempty"`
	CreateTime            string        `json:"createtime,omitempty"`
	DisplayName           string        `json:"displayname,omitempty"`
	Expire                int           `json:"expire,omitempty"`
	Minimum               int           `json:"minimum,omitempty"`
	Prices                []DomainPrice `json:"prices,omitempty"`
	PrimaryNameServer     string        `json:"primarynameserver,omitempty"`
	RecordCount           int           `json:"recordcount,omitempty"`
	Refresh               int           `json:"refresh,omitempty"`
	RegistrarInfo         RegistrarInfo `json:"registrarinfo,omitempty"`
	ResponsiblePerson     string        `json:"responsibleperson,omitempty"`
	Retry                 int           `json:"retry,omitempty"`
	TTL                   int           `json:"ttl,omitempty"`
	UsingGlesysNameserver string        `json:"usingglesysnameserver,omitempty"`
}

type DomainPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Years    int    `json:"years"`
}

// AddDomainParams - used for adding existing domains to GleSYS
// use CreateRecords = false to not create additional records.
type AddDomainParams struct {
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

// EditDomainParams - used when editing domain parameters
type EditDomainParams struct {
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

// RegisterDomainParams - parameters used when registering a domain
type RegisterDomainParams struct {
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

// DeleteDomainParams - parameters for deleting a domain from the dns system.
// Set ForceDeleteEmail to true, to delete a domain AND email accounts for the domain.
type DeleteDomainParams struct {
	Name             string `json:"domainname"`
	ForceDeleteEmail string `json:"forcedeleteemail,omitempty"`
}

// RenewDomainParams - parameters to send when renewing a domain.
type RenewDomainParams struct {
	Name     string `json:"domainname"`
	NumYears int    `json:"numyears"`
}

// DomainRecord - data in the domain
type DomainRecord struct {
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
func (s *DomainService) Available(context context.Context, search string) (*[]Domain, error) {
	data := struct {
		Response struct {
			Domain []Domain
		}
	}{}
	err := s.client.post(context, "domain/available", &data, search)
	return &data.Response.Domain, err
}

// AddDomain - add an existing domain to your GleSYS account
func (s *DomainService) AddDomain(context context.Context, params AddDomainParams) (*Domain, error) {
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/add", &data, params)
	return &data.Response.Domain, err
}

// Delete - deletes a domain from the dns system
func (s *DomainService) Delete(context context.Context, params DeleteDomainParams) error {
	return s.client.post(context, "domain/delete", nil, params)
}

// Details - return details about the domain
func (s *DomainService) Details(context context.Context, domainname string) (*Domain, error) {
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/details", &data, struct {
		Name string `json:"domainname"`
	}{domainname})
	return &data.Response.Domain, err
}

// Edit - edit domain parameters
func (s *DomainService) Edit(context context.Context, params EditDomainParams) (*Domain, error) {
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/edit", &data, params)
	return &data.Response.Domain, err
}

// List - return a list of all domains in your account
func (s *DomainService) List(context context.Context) (*[]Domain, error) {
	data := struct {
		Response struct {
			Domains []Domain
		}
	}{}
	err := s.client.get(context, "domain/list", &data)
	return &data.Response.Domains, err
}

// Register - Register a domain
func (s *DomainService) Register(context context.Context, params RegisterDomainParams) (*Domain, error) {
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/register", &data, params)
	return &data.Response.Domain, err
}

// Renew - Renew a domain
func (s *DomainService) Renew(context context.Context, params RenewDomainParams) (*Domain, error) {
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/renew", &data, params)
	return &data.Response.Domain, err
}

// Transfer - Transfer a domain
func (s *DomainService) Transfer(context context.Context, params RegisterDomainParams) (*Domain, error) {
	data := struct {
		Response struct {
			Domain Domain
		}
	}{}
	err := s.client.post(context, "domain/transfer", &data, params)
	return &data.Response.Domain, err
}

// ListRecords - return a list of all records for domain
func (s *DomainService) ListRecords(context context.Context, domainname string) (*[]DomainRecord, error) {
	data := struct {
		Response struct {
			Records []DomainRecord
		}
	}{}
	err := s.client.post(context, "domain/listrecords", &data, struct {
		Name string `json:"domainname"`
	}{domainname})
	return &data.Response.Records, err
}

// AddRecord - add a domain record
func (s *DomainService) AddRecord(context context.Context, params AddRecordParams) (*DomainRecord, error) {
	data := struct {
		Response struct {
			Record DomainRecord
		}
	}{}
	err := s.client.post(context, "domain/addrecord", &data, params)
	return &data.Response.Record, err
}

// UpdateRecord - update a domain record
func (s *DomainService) UpdateRecord(context context.Context, params UpdateRecordParams) (*DomainRecord, error) {
	data := struct {
		Response struct {
			Record DomainRecord
		}
	}{}
	err := s.client.post(context, "domain/updaterecord", &data, params)
	return &data.Response.Record, err
}

// DeleteRecord deletes a record
func (s *DomainService) DeleteRecord(context context.Context, recordID int) error {
	return s.client.post(context, "domain/deleterecord", nil, struct {
		RecordID int `json:"recordid"`
	}{recordID})
}

// ChangeNameservers - change the nameservers for domain
func (s *DomainService) ChangeNameservers(context context.Context, params ChangeNameserverParams) error {
	return s.client.post(context, "domain/changenameservers", nil, params)
}
