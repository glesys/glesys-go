package glesys

import "context"

// IPService provides functions to interact with IP addresses
type IPService struct {
	client clientInterface
}

// IP represents an IP address
type IP struct {
	Address string `json:"ipaddress"`
	Version int    `json:"version"`
}

// AvailableIPsParams is used to filter results when listing available IP addresses
type AvailableIPsParams struct {
	DataCenter string `json:"datacenter"`
	Platform   string `json:"platform"`
	Version    int    `json:"ipversion"`
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
func (s *IPService) Reserved(context context.Context) (*[]IP, error) {
	data := struct {
		Response struct {
			IPList []IP
		}
	}{}
	err := s.client.get(context, "ip/listown", &data)
	return &data.Response.IPList, err
}
