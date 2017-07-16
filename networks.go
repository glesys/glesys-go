package glesys

import "context"

// NetworkService provides functions to interact with Networks
type NetworkService struct {
	client clientInterface
}

// Network represents a network
type NetworkDetails struct {
	Datacenter  string `json:"datacenter"`
	Description string `json:"description"`
	ID          string `json:"networkid"`
	Public      string `json:"public"`
}

type CreateNetworkParams struct {
	Datacenter  string `json:"datacenter"`
	Description string `json:"description"`
}

type EditNetworkParams struct {
	Description string `json:"description"`
}

// Create creates a new network
func (s *NetworkService) Create(context context.Context, params CreateNetworkParams) (*NetworkDetails, error) {
	data := struct {
		Response struct {
			Network NetworkDetails
		}
	}{}
	err := s.client.post(context, "network/create", &data, params)
	return &data.Response.Network, err
}

// Details returns detailed information about one network
func (s *NetworkService) Details(context context.Context, networkID string) (*NetworkDetails, error) {
	data := struct {
		Response struct {
			Network NetworkDetails
		}
	}{}
	err := s.client.post(context, "network/details", &data, struct {
		NetworkID string `json:"networkid"`
	}{networkID})
	return &data.Response.Network, err
}

// Destroy deletes a network
func (s *NetworkService) Destroy(context context.Context, networkID string) error {
	return s.client.post(context, "network/delete", nil, struct {
		NetworkID string `json:"networkid"`
	}{networkID})
}

// Edit modifies a network
func (s *NetworkService) Edit(context context.Context, networkID string, params EditNetworkParams) (*NetworkDetails, error) {
	data := struct {
		Response struct {
			Network NetworkDetails
		}
	}{}
	err := s.client.post(context, "network/edit", &data, struct {
		EditNetworkParams
		NetworkID string `json:"networkid"`
	}{params, networkID})
	return &data.Response.Network, err
}

// List returns a list of Networks available under your account
func (s *NetworkService) List(context context.Context) (*[]NetworkDetails, error) {
	data := struct {
		Response struct {
			Networks []NetworkDetails
		}
	}{}

	err := s.client.get(context, "network/list", &data)
	return &data.Response.Networks, err
}
