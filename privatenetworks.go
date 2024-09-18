package glesys

import (
	"context"
)

// PrivateNetworkService provides functions to interact with PrivateNetworks
type PrivateNetworkService struct {
	client clientInterface
}

// PrivateNetwork represents a privatenetwork
type PrivateNetwork struct {
	ID            string `json:"id"`
	IPv6Aggregate string `json:"ipv6aggregate"`
	Name          string `json:"name"`
}

// PrivateNetworkBilling
type PrivateNetworkBilling struct {
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"` // discount in $Currency
	Total    float64 `json:"total"`
}

// EditPrivateNetworkParams is used when editing an existing private network
type EditPrivateNetworkParams struct {
	ID   string `json:"privatenetworkid"`
	Name string `json:"name,omitempty"`
}

// PrivateNetworkSegment represents a segment as part of a PrivateNetwork
type PrivateNetworkSegment struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IPv4Subnet string `json:"ipv4subnet"`
	IPv6Subnet string `json:"ipv6subnet"`
	Platform   string `json:"platform"`
	Datacenter string `json:"datacenter"`
}

// CreatePrivateNetworkSegmentParams is used when creating Segments in a PrivateNetwork
type CreatePrivateNetworkSegmentParams struct {
	PrivateNetworkID string `json:"privatenetworkid"`
	Datacenter       string `json:"datacenter"`
	IPv4Subnet       string `json:"ipv4subnet"` // x.y.z.a/netmask
	Name             string `json:"name"`
	Platform         string `json:"platform"`
}

// EditPrivateNetworkSegmentParams is used when editing an existing segment
type EditPrivateNetworkSegmentParams struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// Create creates a new PrivateNetwork
func (s *PrivateNetworkService) Create(context context.Context, name string) (*PrivateNetwork, error) {
	data := struct {
		Response struct {
			PrivateNetwork PrivateNetwork
		}
	}{}
	err := s.client.post(context, "privatenetwork/create", &data, struct {
		Name string `json:"name"`
	}{name})
	return &data.Response.PrivateNetwork, err
}

// Details returns detailed information about a PrivateNetwork
func (s *PrivateNetworkService) Details(context context.Context, privateNetworkID string) (*PrivateNetwork, error) {
	data := struct {
		Response struct {
			PrivateNetwork PrivateNetwork
		}
	}{}
	err := s.client.post(context, "privatenetwork/details", &data, struct {
		PrivateNetworkID string `json:"privatenetworkid"`
	}{privateNetworkID})
	return &data.Response.PrivateNetwork, err
}

// List returns detailed information about a PrivateNetwork
func (s *PrivateNetworkService) List(context context.Context) (*[]PrivateNetwork, error) {
	data := struct {
		Response struct {
			PrivateNetworks []PrivateNetwork
		}
	}{}
	err := s.client.post(context, "privatenetwork/list", &data, nil)
	return &data.Response.PrivateNetworks, err
}

// Destroy deletes a PrivateNetwork
func (s *PrivateNetworkService) Destroy(context context.Context, privateNetworkID string) error {
	return s.client.post(context, "privatenetwork/delete", nil, struct {
		PrivateNetworkID string `json:"privatenetworkid"`
	}{privateNetworkID})
}

// Edit modifies a PrivateNetwork
func (s *PrivateNetworkService) Edit(context context.Context, params EditPrivateNetworkParams) (*PrivateNetwork, error) {
	data := struct {
		Response struct {
			PrivateNetwork PrivateNetwork
		}
	}{}
	err := s.client.post(context, "privatenetwork/edit", &data, struct {
		EditPrivateNetworkParams
	}{params})
	return &data.Response.PrivateNetwork, err
}

// EstimatedCost returns billing information about a PrivateNetwork
func (s *PrivateNetworkService) EstimatedCost(context context.Context, privateNetworkID string) (*PrivateNetworkBilling, error) {
	data := struct {
		Response struct {
			Billing PrivateNetworkBilling
		}
	}{}
	err := s.client.post(context, "privatenetwork/estimatedcost", &data, struct {
		ID string `json:"privatenetworkid,omitempty"`
	}{privateNetworkID})
	return &data.Response.Billing, err
}

// Create creates a new PrivateNetworkSegment
func (s *PrivateNetworkService) CreateSegment(context context.Context, params CreatePrivateNetworkSegmentParams) (*PrivateNetworkSegment, error) {
	data := struct {
		Response struct {
			PrivateNetworkSegment PrivateNetworkSegment
		}
	}{}
	err := s.client.post(context, "privatenetwork/createsegment", &data, params)
	return &data.Response.PrivateNetworkSegment, err
}

// Edit modifies a new PrivateNetworkSegment
func (s *PrivateNetworkService) EditSegment(context context.Context, params EditPrivateNetworkSegmentParams) (*PrivateNetworkSegment, error) {
	data := struct {
		Response struct {
			PrivateNetworkSegment PrivateNetworkSegment
		}
	}{}
	err := s.client.post(context, "privatenetwork/editsegment", &data, params)
	return &data.Response.PrivateNetworkSegment, err
}

// ListSegments returns detailed information about a PrivateNetwork
func (s *PrivateNetworkService) ListSegments(context context.Context, privatenetworkid string) (*[]PrivateNetworkSegment, error) {
	data := struct {
		Response struct {
			PrivateNetworkSegments []PrivateNetworkSegment
		}
	}{}
	err := s.client.post(context, "privatenetwork/listsegments", &data, struct {
		PrivateNetworkID string `json:"privatenetworkid"`
	}{privatenetworkid})
	return &data.Response.PrivateNetworkSegments, err
}

// DestroySegment deletes a PrivateNetworkSegment
func (s *PrivateNetworkService) DestroySegment(context context.Context, id string) error {
	return s.client.post(context, "privatenetwork/deletesegment", nil, struct {
		ID string `json:"id"`
	}{id})
}
