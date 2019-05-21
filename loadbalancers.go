package glesys

import (
	"context"
	"fmt"
)

// LoadbalancerService provides functions to interact with Loadbalancers
type LoadbalancerService struct {
	client clientInterface
}

// Loadbalancer represents a loadbalancer
type Loadbalancer struct {
	DataCenter string `json:"datacenter"`
	ID         string `json:"loadbalancerid"`
	Name       string `json:"name"`
}

type LoadbalancerDetails struct {
	BackendsList  []LBBackend      `json:"backends"`
	Blacklists    []string         `json:"blacklist"`
	DataCenter    string           `json:"datacenter"`
	FrontendsList []LBFrontend     `json:"frontends"`
	ID            string           `json:"loadbalancerid"`
	IPList        []LoadbalancerIP `json:"ipaddress"`
	Name          string           `json:"name"`
}

type LoadbalancerIP struct {
	Address         string `json:"ipaddress"`
	Cost            int    `json:"cost"`
	LockedToAccount bool   `json:"lockedtoaccount"`
	Version         int    `json:"version"`
}

// Certificate associated with the loadbalancer
type Certificate struct {
	Name []string `json:"certificate"`
}

// CreateLoadbalancerParams is used when creating a new loadbalancer
type CreateLoadbalancerParams struct {
	DataCenter string `json:"datacenter"`
	IPv4       string `json:"ip,omitempty"`
	IPv6       string `json:"ipv6,omitempty"`
	Name       string `json:"name"`
}

// EditLoadbalancerParams is used when editing a loadbalancer
type EditLoadbalancerParams struct {
	Name string `json:"name"`
}

// LBBackend represents a Loadbalancer Backend
type LBBackend struct {
	ConnectTimeout int      `json:"connecttimeout"`
	Mode           string   `json:"mode"`
	Name           string   `json:"name"`
	Targets        []Target `json:"targets"`
}

// AddBackendParams used when creating backends
type AddBackendParams struct {
	ConnectTimeout  int    `json:"connecttimeout,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Name            string `json:"name"`
	ResponseTimeout int    `json:"responsetimeout,omitempty"`
	Stickysessions  string `json:"stickysessions,omitempty"`
}

// EditBackendParams used to edit a backend
type EditBackendParams struct {
	ConnectTimeout  int    `json:"connecttimeout,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Name            string `json:"name"`
	ResponseTimeout int    `json:"responsetimeout,omitempty"`
	Stickysessions  string `json:"stickysessions,omitempty"`
}

// RemoveBackendParams used when removing
type RemoveBackendParams struct {
	Name string `json:"backendname"`
}

// AddCertificateParams
type AddCertificateParams struct {
	Name        string `json:"certificatename"`
	Certificate string `json:"certificate"`
}

// LBFrontend represents a Loadbalancer Frontend
type LBFrontend struct {
	Backend        string `json:"backend"`
	ClientTimeout  int    `json:"clienttimeout"`
	MaxConnections int    `json:"maxconnections"`
	Name           string `json:"name"`
	Port           int    `json:"port"`
	Status         string `json:"status"`
	SSLCertificate string `json:"sslcertificate"`
}

// AddFrontendParams used when creating frontends
type AddFrontendParams struct {
	Backend        string `json:"backendname"`
	ClientTimeout  int    `json:"clienttimeout,omitempty"`
	MaxConnections int    `json:"maxconnections,omitempty"`
	Name           string `json:"name"`
	Port           int    `json:"port"`
	SSLCertificate string `json:"sslcertificate,omitempty"`
}

// EditFrontendParams used to edit a frontend
type EditFrontendParams struct {
	ClientTimeout  int    `json:"clienttimeout,omitempty"`
	MaxConnections int    `json:"maxconnections,omitempty"`
	Name           string `json:"name"`
	Port           int    `json:"port,omitempty"`
	SSLCertificate string `json:"sslcertificate,omitempty"`
}

// RemoveFrontendParams used when removing
type RemoveFrontendParams struct {
	Name string `json:"backendname"`
}

// Targets used in backends
type Target struct {
	Enabled  bool   `json:"enabled"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Status   string `json:"status"`
	TargetIP string `json:"ipaddress"`
	Weight   int    `json:"weight"`
}

// AddTargetParams used when creating targets
type AddTargetParams struct {
	Backend  string `json:"backendname"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	TargetIP string `json:"ipaddress"`
	Weight   int    `json:"weight"`
}

// EditTargetParams used when editing targets
type EditTargetParams struct {
	Backend  string `json:"backendname"`
	Name     string `json:"name"`
	Port     int    `json:"port,omitempty"`
	TargetIP string `json:"ipaddress,omitempty"`
	Weight   int    `json:"weight,omitempty"`
}

// RemoveTargetParams used when removing targets
type RemoveTargetParams struct {
	Backend string `json:"backendname"`
	Name    string `json:"name"`
}

// BlacklistParams set prefix to add/delete
type BlacklistParams struct {
	Prefix string `json:"prefix"`
}

// ToggleTargetParams used when enabling/disabling targets
type ToggleTargetParams struct {
	Backend string `json:"backendname"`
	Name    string `json:"targetname"`
}

// Create creates a new network
func (lb *LoadbalancerService) Create(context context.Context, params CreateLoadbalancerParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/create", &data, params)
	return &data.Response.Loadbalancer, err
}

// Destroy deletes a loadbalancer
func (lb *LoadbalancerService) Destroy(context context.Context, loadbalancerID string) error {
	return lb.client.post(context, "loadbalancer/destroy", nil, struct {
		LoadbalancerID string `json:"loadbalancerid"`
	}{loadbalancerID})
}

// Details returns a detailed information about one loadbalancer
func (lb *LoadbalancerService) Details(context context.Context, loadbalancerID string) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.get(context, fmt.Sprintf("loadbalancer/details/loadbalancerid/%s", loadbalancerID), &data)
	return &data.Response.Loadbalancer, err
}

// Edit edits a loadbalancer
func (lb *LoadbalancerService) Edit(context context.Context, loadbalancerID string, params EditLoadbalancerParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/edit", &data, struct {
		EditLoadbalancerParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// List returns a list of loadbalancers
func (lb *LoadbalancerService) List(context context.Context) (*[]Loadbalancer, error) {
	data := struct {
		Response struct {
			Loadbalancers []Loadbalancer
		}
	}{}
	err := lb.client.get(context, "loadbalancer/list", &data)
	return &data.Response.Loadbalancers, err
}

// AddBackend creates a new backend used by the loadbalancer specified
func (lb *LoadbalancerService) AddBackend(context context.Context, loadbalancerID string, params AddBackendParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addbackend", &data, struct {
		AddBackendParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// EditBackend edits a Backend
func (lb *LoadbalancerService) EditBackend(context context.Context, loadbalancerID string, params EditBackendParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/editbackend", &data, struct {
		EditBackendParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// RemoveBackend deletes a backend
func (lb *LoadbalancerService) RemoveBackend(context context.Context, loadbalancerID string, params RemoveBackendParams) error {
	return lb.client.post(context, "loadbalancer/removebackend", nil, struct {
		RemoveBackendParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// AddFrontend creates a new frontend used by the loadbalancer specified
func (lb *LoadbalancerService) AddFrontend(context context.Context, loadbalancerID string, params AddFrontendParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addfrontend", &data, struct {
		AddFrontendParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// EditFrontend edits a frontend
func (lb *LoadbalancerService) EditFrontend(context context.Context, loadbalancerID string, params EditFrontendParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/editfrontend", &data, struct {
		EditFrontendParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// RemoveFrontend deletes a frontend
func (lb *LoadbalancerService) RemoveFrontend(context context.Context, loadbalancerID string, params RemoveFrontendParams) error {
	return lb.client.post(context, "loadbalancer/removefrontend", nil, struct {
		RemoveFrontendParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// AddCertificate adds a certificate to the loadbalancer specified
func (lb *LoadbalancerService) AddCertificate(context context.Context, loadbalancerID string, params AddCertificateParams) error {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	return lb.client.post(context, "loadbalancer/addcertificate", &data, struct {
		AddCertificateParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// ListCertificate
func (lb *LoadbalancerService) ListCertificate(context context.Context, loadbalancerID string) (*[]string, error) {
	data := struct {
		Response struct {
			Certificates []string `json:"certificate"` // TODO cleanup and use the Certificate struct
		}
	}{}
	err := lb.client.post(context, "loadbalancer/listcertificate", &data, struct {
		LoadbalancerID string `json:"loadbalancerid"`
	}{loadbalancerID})

	return &data.Response.Certificates, err
}

// RemoveCertificate deletes a certificate from the loadbalancer
func (lb *LoadbalancerService) RemoveCertificate(context context.Context, loadbalancerID string, params string) error {
	return lb.client.post(context, "loadbalancer/removecertificate", nil, struct {
		CertificateName string `json:"certificatename"`
		LoadbalancerID  string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// AddTarget adds a target to the backend specified
func (lb *LoadbalancerService) AddTarget(context context.Context, loadbalancerID string, params AddTargetParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addtarget", &data, struct {
		AddTargetParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// EditTarget edits a target
func (lb *LoadbalancerService) EditTarget(context context.Context, loadbalancerID string, params EditTargetParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/edittarget", &data, struct {
		EditTargetParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// Enable a target
func (lb *LoadbalancerService) EnableTarget(context context.Context, loadbalancerID string, params ToggleTargetParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/enabletarget", &data, struct {
		ToggleTargetParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// Disable a target
func (lb *LoadbalancerService) DisableTarget(context context.Context, loadbalancerID string, params ToggleTargetParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/disabletarget", &data, struct {
		ToggleTargetParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// RemoveTarget deletes a target
func (lb *LoadbalancerService) RemoveTarget(context context.Context, loadbalancerID string, params RemoveTargetParams) error {
	return lb.client.post(context, "loadbalancer/removetarget", nil, struct {
		RemoveTargetParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// Addtoblacklist adds a prefix to loadbalancer blacklist
func (lb *LoadbalancerService) Addtoblacklist(context context.Context, loadbalancerID string, params BlacklistParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addtoblacklist", &data, struct {
		BlacklistParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}

// Removefromblacklist deletes a prefix from loadbalancer blacklist
func (lb *LoadbalancerService) Removefromblacklist(context context.Context, loadbalancerID string, params BlacklistParams) (*LoadbalancerDetails, error) {
	data := struct {
		Response struct {
			Loadbalancer LoadbalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/removefromblacklist", &data, struct {
		BlacklistParams
		LoadbalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.Loadbalancer, err
}
