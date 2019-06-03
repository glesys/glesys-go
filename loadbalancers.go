package glesys

import (
	"context"
	"fmt"
)

// LoadBalancerService provides functions to interact with LoadBalancers
type LoadBalancerService struct {
	client clientInterface
}

// LoadBalancer represents a loadbalancer
type LoadBalancer struct {
	DataCenter string `json:"datacenter"`
	ID         string `json:"loadbalancerid"`
	Name       string `json:"name"`
}

type LoadBalancerDetails struct {
	BackendsList  []LoadBalancerBackend  `json:"backends"`
	Blacklists    []string               `json:"blacklist"`
	DataCenter    string                 `json:"datacenter"`
	FrontendsList []LoadBalancerFrontend `json:"frontends"`
	ID            string                 `json:"loadbalancerid"`
	IPList        []LoadBalancerIP       `json:"ipaddress"`
	Name          string                 `json:"name"`
}

type LoadBalancerIP struct {
	Address         string `json:"ipaddress"`
	Cost            int    `json:"cost"`
	LockedToAccount bool   `json:"lockedtoaccount"`
	Version         int    `json:"version"`
}

// CreateLoadBalancerParams is used when creating a new loadbalancer
type CreateLoadBalancerParams struct {
	DataCenter string `json:"datacenter"`
	IPv4       string `json:"ip,omitempty"`
	IPv6       string `json:"ipv6,omitempty"`
	Name       string `json:"name"`
}

// EditLoadBalancerParams is used when editing a loadbalancer
type EditLoadBalancerParams struct {
	Name string `json:"name"`
}

// LoadBalancerBackend represents a LoadBalancer Backend
type LoadBalancerBackend struct {
	ConnectTimeout  int      `json:"connecttimeout"`
	Mode            string   `json:"mode"`
	Name            string   `json:"name"`
	ResponseTimeout int      `json:"responsetimeout"`
	Status          string   `json:"status"`
	StickySession   string   `json:"stickysessions"`
	Targets         []Target `json:"targets"`
}

// AddBackendParams used when creating backends
type AddBackendParams struct {
	ConnectTimeout  int    `json:"connecttimeout,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Name            string `json:"name"`
	ResponseTimeout int    `json:"responsetimeout,omitempty"`
	StickySession   string `json:"stickysession,omitempty"`
}

// EditBackendParams used to edit a backend
type EditBackendParams struct {
	ConnectTimeout  int    `json:"connecttimeout,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Name            string `json:"backendname"`
	ResponseTimeout int    `json:"responsetimeout,omitempty"`
	StickySession   string `json:"stickysession,omitempty"`
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

// LoadBalancerFrontend represents a LoadBalancer Frontend
type LoadBalancerFrontend struct {
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
	Name           string `json:"frontendname"`
	Port           int    `json:"port,omitempty"`
	SSLCertificate string `json:"sslcertificate,omitempty"`
}

// RemoveFrontendParams used when removing
type RemoveFrontendParams struct {
	Name string `json:"frontendname"`
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
	Name     string `json:"targetname"`
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

// Create creates a new loadbalancer
func (lb *LoadBalancerService) Create(context context.Context, params CreateLoadBalancerParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/create", &data, params)
	return &data.Response.LoadBalancer, err
}

// Destroy deletes a loadbalancer
func (lb *LoadBalancerService) Destroy(context context.Context, loadbalancerID string) error {
	return lb.client.post(context, "loadbalancer/destroy", nil, struct {
		LoadBalancerID string `json:"loadbalancerid"`
	}{loadbalancerID})
}

// Details returns a detailed information about one loadbalancer
func (lb *LoadBalancerService) Details(context context.Context, loadbalancerID string) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.get(context, fmt.Sprintf("loadbalancer/details/loadbalancerid/%s", loadbalancerID), &data)
	return &data.Response.LoadBalancer, err
}

// Edit edits a loadbalancer
func (lb *LoadBalancerService) Edit(context context.Context, loadbalancerID string, params EditLoadBalancerParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/edit", &data, struct {
		EditLoadBalancerParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// List returns a list of loadbalancers
func (lb *LoadBalancerService) List(context context.Context) (*[]LoadBalancer, error) {
	data := struct {
		Response struct {
			LoadBalancers []LoadBalancer
		}
	}{}
	err := lb.client.get(context, "loadbalancer/list", &data)
	return &data.Response.LoadBalancers, err
}

// AddBackend creates a new backend used by the loadbalancer specified
func (lb *LoadBalancerService) AddBackend(context context.Context, loadbalancerID string, params AddBackendParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addbackend", &data, struct {
		AddBackendParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// EditBackend edits a Backend
func (lb *LoadBalancerService) EditBackend(context context.Context, loadbalancerID string, params EditBackendParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/editbackend", &data, struct {
		EditBackendParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// RemoveBackend deletes a backend
func (lb *LoadBalancerService) RemoveBackend(context context.Context, loadbalancerID string, params RemoveBackendParams) error {
	return lb.client.post(context, "loadbalancer/removebackend", nil, struct {
		RemoveBackendParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// AddFrontend creates a new frontend used by the loadbalancer specified
func (lb *LoadBalancerService) AddFrontend(context context.Context, loadbalancerID string, params AddFrontendParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addfrontend", &data, struct {
		AddFrontendParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// EditFrontend edits a frontend
func (lb *LoadBalancerService) EditFrontend(context context.Context, loadbalancerID string, params EditFrontendParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/editfrontend", &data, struct {
		EditFrontendParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// RemoveFrontend deletes a frontend
func (lb *LoadBalancerService) RemoveFrontend(context context.Context, loadbalancerID string, params RemoveFrontendParams) error {
	return lb.client.post(context, "loadbalancer/removefrontend", nil, struct {
		RemoveFrontendParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// AddCertificate adds a certificate to the loadbalancer specified
func (lb *LoadBalancerService) AddCertificate(context context.Context, loadbalancerID string, params AddCertificateParams) error {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	return lb.client.post(context, "loadbalancer/addcertificate", &data, struct {
		AddCertificateParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// ListCertificate list certificates for the LoadBalancer
func (lb *LoadBalancerService) ListCertificate(context context.Context, loadbalancerID string) (*[]string, error) {
	data := struct {
		Response struct {
			Certificates []string `json:"certificate"` // TODO cleanup and use the Certificate struct
		}
	}{}
	err := lb.client.post(context, "loadbalancer/listcertificate", &data, struct {
		LoadBalancerID string `json:"loadbalancerid"`
	}{loadbalancerID})

	return &data.Response.Certificates, err
}

// RemoveCertificate deletes a certificate from the loadbalancer
func (lb *LoadBalancerService) RemoveCertificate(context context.Context, loadbalancerID string, params string) error {
	return lb.client.post(context, "loadbalancer/removecertificate", nil, struct {
		CertificateName string `json:"certificatename"`
		LoadBalancerID  string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// AddTarget adds a target to the backend specified
func (lb *LoadBalancerService) AddTarget(context context.Context, loadbalancerID string, params AddTargetParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addtarget", &data, struct {
		AddTargetParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// EditTarget edits a target for the specified backend
func (lb *LoadBalancerService) EditTarget(context context.Context, loadbalancerID string, params EditTargetParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/edittarget", &data, struct {
		EditTargetParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// EnableTarget enables a target for the specified LoadBalancerBackend
func (lb *LoadBalancerService) EnableTarget(context context.Context, loadbalancerID string, params ToggleTargetParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/enabletarget", &data, struct {
		ToggleTargetParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// DisableTarget disables the specified target for the LoadBalancerBackend
func (lb *LoadBalancerService) DisableTarget(context context.Context, loadbalancerID string, params ToggleTargetParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/disabletarget", &data, struct {
		ToggleTargetParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// RemoveTarget deletes a target from the specified LoadBalancerBackend
func (lb *LoadBalancerService) RemoveTarget(context context.Context, loadbalancerID string, params RemoveTargetParams) error {
	return lb.client.post(context, "loadbalancer/removetarget", nil, struct {
		RemoveTargetParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
}

// AddToBlacklist adds a prefix to loadbalancer blacklist
func (lb *LoadBalancerService) AddToBlacklist(context context.Context, loadbalancerID string, params BlacklistParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/addtoblacklist", &data, struct {
		BlacklistParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}

// RemoveFromBlacklist deletes a prefix from the LoadBalancer blacklist
func (lb *LoadBalancerService) RemoveFromBlacklist(context context.Context, loadbalancerID string, params BlacklistParams) (*LoadBalancerDetails, error) {
	data := struct {
		Response struct {
			LoadBalancer LoadBalancerDetails
		}
	}{}
	err := lb.client.post(context, "loadbalancer/removefromblacklist", &data, struct {
		BlacklistParams
		LoadBalancerID string `json:"loadbalancerid"`
	}{params, loadbalancerID})
	return &data.Response.LoadBalancer, err
}
