package glesys

import (
	"context"
	"net"
	"strings"
)

// IPService provides functions to interact with IP addresses
type IPService struct {
	client clientInterface
}

// IP represents an IP address
type IP struct {
	Address         string   `json:"ipaddress"`
	Broadcast       string   `json:"broadcast,omitempty"`
	Cost            IPCost   `json:"cost,omitempty"`
	DataCenter      string   `json:"datacenter,omitempty"`
	Gateway         string   `json:"gateway,omitempty"`
	LockedToAccount string   `json:"lockedtoaccount,omitempty"`
	NameServers     []string `json:"nameservers,omitempty"`
	Netmask         string   `json:"netmask,omitempty"`
	Platforms       []string `json:"platforms,omitempty"`
	Platform        string   `json:"platform,omitempty"`
	PTR             string   `json:"ptr,omitempty"`
	Reserved        string   `json:"reserved,omitempty"`
	ServerID        string   `json:"serverid,omitempty"`
	Version         int      `json:"ipversion,omitempty"`
}

// AvailableIPsParams is used to filter results when listing available IP addresses
type AvailableIPsParams struct {
	DataCenter string `json:"datacenter"`
	Platform   string `json:"platform"`
	Version    int    `json:"ipversion"`
}

// ReservedIPsParams is used to filter results when listing reserved IP addresses
type ReservedIPsParams struct {
	DataCenter string `json:"datacenter,omitempty"`
	Platform   string `json:"platform,omitempty"`
	Version    int    `json:"ipversion,omitempty"`
	Used       string `json:"used,omitempty"`
}

// IPCost is used to show cost details for a IP address
type IPCost struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	TimePeriod string  `json:"timeperiod"`
}

// Available returns a list of IP addresses available for reservation
func (s *IPService) Available(context context.Context, params AvailableIPsParams) (*[]IP, error) {
	data := struct {
		Response struct {
			IPList struct {
				IPAddresses []string
			}
		}
	}{}

	err := s.client.post(context, "ip/listfree", &data, params)
	if err != nil {
		return nil, err
	}

	// Because, inconsistencies...
	ips := make([]IP, 0)
	for _, address := range data.Response.IPList.IPAddresses {
		ips = append(ips, IP{Address: address})
	}
	return &ips, nil
}

// Details about an IP address
func (s *IPService) Details(context context.Context, ipAddress string) (*IP, error) {
	data := struct {
		Response struct {
			Details IP
		}
	}{}
	err := s.client.post(context, "ip/details", &data, map[string]string{"ipaddress": ipAddress})
	return &data.Response.Details, err
}

// IsIPv4 verify that ip is IPv4
func (ip *IP) IsIPv4() bool {
	netAddr := net.ParseIP(ip.Address)
	return netAddr != nil && strings.Contains(ip.Address, ".")
}

// IsIPv6 verify that ip is IPv6
func (ip *IP) IsIPv6() bool {
	netAddr := net.ParseIP(ip.Address)
	return netAddr != nil && strings.Contains(ip.Address, ":")
}

// Release releases a reserved IP address
func (s *IPService) Release(context context.Context, ipAddress string) error {
	return s.client.post(context, "ip/release", nil, map[string]string{"ipaddress": ipAddress})
}

// Reserve reserves an available IP address
func (s *IPService) Reserve(context context.Context, ipAddress string) (*IP, error) {
	data := struct {
		Response struct {
			Details IP
		}
	}{}
	err := s.client.post(context, "ip/take", &data, map[string]string{"ipaddress": ipAddress})
	return &data.Response.Details, err
}

// Reserved returns a list of reserved IP addresses
func (s *IPService) Reserved(context context.Context, params ReservedIPsParams) (*[]IP, error) {
	data := struct {
		Response struct {
			IPList []IP
		}
	}{}
	err := s.client.post(context, "ip/listown", &data, params)
	return &data.Response.IPList, err
}

// SetPTR sets PTR for an IP address
func (s *IPService) SetPTR(context context.Context, ipAddress string, ptrdata string) (*IP, error) {
	data := struct {
		Response struct {
			Details IP
		}
	}{}
	err := s.client.post(context, "ip/setptr", &data, struct {
		Address string `json:"ipaddress"`
		PTR     string `json:"data"`
	}{ipAddress, ptrdata})
	return &data.Response.Details, err
}

// ResetPTR resets PTR for an IP address
func (s *IPService) ResetPTR(context context.Context, ipAddress string) (*IP, error) {
	data := struct {
		Response struct {
			Details IP
		}
	}{}
	err := s.client.post(context, "ip/resetptr", &data, struct {
		Address string `json:"ipaddress"`
	}{ipAddress})
	return &data.Response.Details, err
}
