package glesys

import "context"

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

// Parameters to use when creating a NetworkAdapter
type CreateNetworkAdapterParams struct {
	AdapterType string `json:"adaptertype,omitempty"`
	Bandwidth   int    `json:"bandwidth,omitempty"`
	NetworkID   string `json:"networkid,omitempty"`
	ServerID    string `json:"serverid"`
}

// Parameters to use when modifying a NetworkAdapter
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
func (s *NetworkAdapterService) Details(context context.Context, NetworkAdapterID string) (*NetworkAdapter, error) {
	data := struct {
		Response struct {
			NetworkAdapter NetworkAdapter
		}
	}{}
	err := s.client.post(context, "networkadapter/details", &data, struct {
		NetworkAdapterID string `json:"networkadapterid"`
	}{NetworkAdapterID})
	return &data.Response.NetworkAdapter, err
}

// Destroy deletes a NetworkAdapter
func (s *NetworkAdapterService) Destroy(context context.Context, NetworkAdapterID string) error {
	return s.client.post(context, "networkadapter/delete", nil, struct {
		NetworkAdapterID string `json:"networkadapterid"`
	}{NetworkAdapterID})
}

// Edit modifies a NetworkAdapter
func (s *NetworkAdapterService) Edit(context context.Context, NetworkAdapterID string, params EditNetworkAdapterParams) (*NetworkAdapter, error) {
	data := struct {
		Response struct {
			NetworkAdapter NetworkAdapter
		}
	}{}
	err := s.client.post(context, "networkadapter/edit", &data, struct {
		EditNetworkAdapterParams
		NetworkAdapterID string `json:"networkadapterid"`
	}{params, NetworkAdapterID})
	return &data.Response.NetworkAdapter, err
}

// IsLocked returns true if the networkadapter is currently locked, false otherwise
func (na *NetworkAdapter) IsLocked() bool {
	return na.State == "locked"
}

// IsReady returns true if the networkadapter is currently ready, false otherwise
func (na *NetworkAdapter) IsReady() bool {
	return na.State == "ready"
}
