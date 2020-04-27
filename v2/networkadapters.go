package glesys

import (
	"context"
	"fmt"
)

// NetworkAdapterService provides functions to interact with Networks
type NetworkAdapterService struct {
	client clientInterface
}

// NetworkAdapter represents a networkadapter
type NetworkAdapter struct {
	AdapterType string `json:"adaptertype,omitempty"`
	Bandwidth   int    `json:"bandwidth"`
	ID          string `json:"networkadapterid"`
	Name        string `json:"name"`
	NetworkID   string `json:"networkid"`
	ServerID    string `json:"serverid"`
	State       string `json:"state"`
}

// IsLocked returns true if the network adapter is currently locked, false otherwise
func (na *NetworkAdapter) IsLocked() bool {
	return na.State == "locked"
}

// IsReady returns true if the network adapter is currently ready, false otherwise
func (na *NetworkAdapter) IsReady() bool {
	return na.State == "ready"
}

// CreateNetworkAdapterParams is used when creating a new network adapter
type CreateNetworkAdapterParams struct {
	AdapterType string `json:"adaptertype,omitempty"`
	Bandwidth   int    `json:"bandwidth,omitempty"`
	NetworkID   string `json:"networkid,omitempty"`
	ServerID    string `json:"serverid"`
}

// EditNetworkAdapterParams is used when editing an existing network adapter
type EditNetworkAdapterParams struct {
	Bandwidth int    `json:"bandwidth,omitempty"`
	NetworkID string `json:"networkid,omitempty"`
}

// Create creates a new NetworkAdapter
func (s *NetworkAdapterService) Create(context context.Context, params CreateNetworkAdapterParams) (*NetworkAdapter, error) {
	data := struct {
		Response struct {
			NetworkAdapter NetworkAdapter
		}
	}{}
	err := s.client.post(context, "networkadapter/create", &data, params)
	return &data.Response.NetworkAdapter, err
}

// Details returns detailed information about a NetworkAdapter
func (s *NetworkAdapterService) Details(context context.Context, networkAdapterID string) (*NetworkAdapter, error) {
	data := struct {
		Response struct {
			NetworkAdapter NetworkAdapter
		}
	}{}
	err := s.client.get(context, fmt.Sprintf("networkadapter/details/networkadapterid/%s", networkAdapterID), &data)
	return &data.Response.NetworkAdapter, err
}

// Destroy deletes a NetworkAdapter
func (s *NetworkAdapterService) Destroy(context context.Context, networkAdapterID string) error {
	return s.client.post(context, "networkadapter/delete", nil, struct {
		NetworkAdapterID string `json:"networkadapterid"`
	}{networkAdapterID})
}

// Edit modifies a NetworkAdapter
func (s *NetworkAdapterService) Edit(context context.Context, networkAdapterID string, params EditNetworkAdapterParams) (*NetworkAdapter, error) {
	data := struct {
		Response struct {
			NetworkAdapter NetworkAdapter
		}
	}{}
	err := s.client.post(context, "networkadapter/edit", &data, struct {
		EditNetworkAdapterParams
		NetworkAdapterID string `json:"networkadapterid"`
	}{params, networkAdapterID})
	return &data.Response.NetworkAdapter, err
}
