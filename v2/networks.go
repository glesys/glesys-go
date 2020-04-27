package glesys

import (
	"context"
	"fmt"
)

// NetworkService provides functions to interact with Networks
type NetworkService struct {
	client clientInterface
}

// Network represents a network
type Network struct {
	DataCenter  string `json:"datacenter"`
	Description string `json:"description"`
	ID          string `json:"networkid"`
	Public      string `json:"public"`
}

// CreateNetworkParams is used when creating a new network
type CreateNetworkParams struct {
	DataCenter  string `json:"datacenter"`
	Description string `json:"description"`
}

// EditNetworkParams is used when editing an existing network
type EditNetworkParams struct {
	Description string `json:"description"`
}

// Create creates a new network
func (s *NetworkService) Create(context context.Context, params CreateNetworkParams) (*Network, error) {
	data := struct {
		Response struct {
			Network Network
		}
	}{}
	err := s.client.post(context, "network/create", &data, params)
	return &data.Response.Network, err
}

// Details returns detailed information about one network
func (s *NetworkService) Details(context context.Context, networkID string) (*Network, error) {
	data := struct {
		Response struct {
			Network Network
		}
	}{}
	err := s.client.get(context, fmt.Sprintf("network/details/networkid/%s", networkID), &data)
	return &data.Response.Network, err
}

// Destroy deletes a network
func (s *NetworkService) Destroy(context context.Context, networkID string) error {
	return s.client.post(context, "network/delete", nil, struct {
		NetworkID string `json:"networkid"`
	}{networkID})
}

// Edit modifies a network
func (s *NetworkService) Edit(context context.Context, networkID string, params EditNetworkParams) (*Network, error) {
	data := struct {
		Response struct {
			Network Network
		}
	}{}
	err := s.client.post(context, "network/edit", &data, struct {
		EditNetworkParams
		NetworkID string `json:"networkid"`
	}{params, networkID})
	return &data.Response.Network, err
}

// IsPublic return true if network is public
func (s *Network) IsPublic() bool {
	return s.Public == "yes"
}

// List returns a list of Networks available under your account
func (s *NetworkService) List(context context.Context) (*[]Network, error) {
	data := struct {
		Response struct {
			Networks []Network
		}
	}{}

	err := s.client.get(context, "network/list", &data)
	return &data.Response.Networks, err
}
