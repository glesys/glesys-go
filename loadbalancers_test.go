package glesys

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadbalancersCreate(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer":
	{ "backends": [], "datacenter": "Falkenberg", "frontends": [],
	"ipaddress": [{"ipaddress": "192.168.0.1", "version": 4}],
	"name": "myloadbalancer", "loadbalancerid": "lb123456" }}}`}
	lb := LoadbalancerService{client: c}

	params := CreateLoadbalancerParams{
		DataCenter: "Falkenberg",
		IPv4:       "192.168.0.1",
		Name:       "myloadbalancer",
	}

	loadbalancer, _ := lb.Create(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "loadbalancer/create", c.lastPath, "path used is correct")
	assert.Equal(t, "lb123456", loadbalancer.ID, "loadbalancer ID is correct")
	assert.Equal(t, "Falkenberg", loadbalancer.DataCenter, "loadbalancer DataCenter is correct")
	assert.Equal(t, "myloadbalancer", loadbalancer.Name, "loadbalancer Name is correct")
	assert.Equal(t, "192.168.0.1", loadbalancer.IPList[0].Address, "loadbalancer ip is correct")
}

func TestLoadbalancersDestroy(t *testing.T) {
	c := &mockClient{}
	lb := LoadbalancerService{client: c}

	lb.Destroy(context.Background(), "lb123456", DestroyLoadbalancerParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/destroy", c.lastPath, "path used is correct")
}

func TestLoadbalancersDetails(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer":
	{ "backends": [], "datacenter": "Falkenberg", "frontends": [],
	"ipaddress": [{"ipaddress": "192.168.0.1", "version": 4}],
	"name": "myloadbalancer", "loadbalancerid": "lb123456" }}}`}
	lb := LoadbalancerService{client: c}

	loadbalancer, _ := lb.Details(context.Background(), "lb123456")

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/details/loadbalancerid/lb123456", c.lastPath, "path used is correct")
	assert.Equal(t, "myloadbalancer", loadbalancer.Name, "loadbalancer Name is correct")
	assert.Equal(t, "Falkenberg", loadbalancer.DataCenter, "loadbalancer DataCenter is correct")
}

func TestLoadlancersEdit(t *testing.T) {
	c := &mockClient{}
	lb := LoadbalancerService{client: c}

	lb.Edit(context.Background(), "lb123456", EditLoadbalancerParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/edit", c.lastPath, "path used is correct")
}

func TestLoadbalancersList(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancers": [
		{ "backends": [], "datacenter": "Falkenberg", "frontends": [],
		"ipaddress": [{"ipaddress": "192.168.0.1", "version": 4}],
		"name": "myloadbalancer", "loadbalancerid": "lb123456" }]
		}}`}
	lb := LoadbalancerService{client: c}

	loadbalancers, _ := lb.List(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/list", c.lastPath, "path used is correct")
	assert.Equal(t, "myloadbalancer", (*loadbalancers)[0].Name, "loadbalancer Name is correct")
	assert.Equal(t, "Falkenberg", (*loadbalancers)[0].DataCenter, "loadbalancer DataCenter is correct")
}

func TestLoadbalancersAddBackend(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"backends": [{"connecttimeout": 4000, "mode": "tcp", "name": "mybackend",
		"stickysessions": "no", "targets": [] }],
		"loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}
	params := AddBackendParams{
		Name:           "mybackend",
		Mode:           "tcp",
		Stickysessions: "no",
	}

	loadbalancer, _ := lb.AddBackend(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/addbackend", c.lastPath, "path used is correct")
	assert.Equal(t, "mybackend", loadbalancer.BackendsList[0].Name, "backend name is correct")
	assert.Equal(t, "tcp", loadbalancer.BackendsList[0].Mode, "backend mode is correct")
}

func TestLoadbalancersEditBackend(t *testing.T) {
	c := &mockClient{}
	lb := LoadbalancerService{client: c}

	params := EditBackendParams{
		Name: "mybackend",
	}

	lb.EditBackend(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/editbackend", c.lastPath, "path used is correct")
}

func TestLoadbalancersRemoveBackend(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"backends": [], "loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}
	params := RemoveBackendParams{
		Name: "mybackend",
	}

	lb.RemoveBackend(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/removebackend", c.lastPath, "path used is correct")
}

func TestLoadbalancersAddFrontend(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"backends": [{"connecttimeout": 4000, "mode": "tcp", "name": "mybackend",
		"stickysessions": "no", "targets": [] }],
		"frontends": [{"backend": "mybackend", "name": "myfrontend", "port": 8080}],
		"loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}
	params := AddFrontendParams{
		Backend: "mybackend",
		Name:    "myfrontend",
		Port:    8080,
	}

	loadbalancer, _ := lb.AddFrontend(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/addfrontend", c.lastPath, "path used is correct")
	assert.Equal(t, "myfrontend", loadbalancer.FrontendsList[0].Name, "Frontend name is correct")
	assert.Equal(t, 8080, loadbalancer.FrontendsList[0].Port, "Frontend port is correct")
}

func TestLoadbalancersEditFrontend(t *testing.T) {
	c := &mockClient{}
	lb := LoadbalancerService{client: c}

	params := EditFrontendParams{
		Name: "myfrontend",
	}

	lb.EditFrontend(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/editfrontend", c.lastPath, "path used is correct")
}

func TestLoadbalancersRemoveFrontend(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"frontends": [], "loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}
	params := RemoveFrontendParams{
		Name: "myfrontend",
	}

	lb.RemoveFrontend(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/removefrontend", c.lastPath, "path used is correct")
}

func TestLoadbalancersAddTarget(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"backends": [{"connecttimeout": 4000, "mode": "tcp", "name": "mybackend", "stickysessions": "no",
			"targets": [{"ipaddress": "8.8.8.8", "name": "mytarget", "port": 8080, "status": "DOWN", "weight": 10}]
			}],
		"loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}

	params := AddTargetParams{
		Backend:  "mybackend",
		Name:     "mytarget",
		Port:     8080,
		TargetIP: "8.8.8.8",
		Weight:   10,
	}

	loadbalancer, _ := lb.AddTarget(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/addtarget", c.lastPath, "path used is correct")
	assert.Equal(t, "mytarget", loadbalancer.BackendsList[0].Targets[0].Name, "Target name is correct")
	assert.Equal(t, 8080, loadbalancer.BackendsList[0].Targets[0].Port, "Target port is correct")
	assert.Equal(t, 10, loadbalancer.BackendsList[0].Targets[0].Weight, "Target weight is correct")
}

func TestLoadbalancersEditTarget(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"backends": [{"connecttimeout": 4000, "mode": "tcp", "name": "mybackend", "stickysessions": "no",
			"targets": [{"ipaddress": "8.8.8.8", "name": "mytarget", "port": 8080, "status": "DOWN", "weight": 10}]
			}],
		"loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}

	params := EditTargetParams{
		Backend:  "mybackend",
		Name:     "mytarget",
		Port:     8080,
		TargetIP: "8.8.8.8",
		Weight:   10,
	}

	lb.EditTarget(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/edittarget", c.lastPath, "path used is correct")
}

func TestLoadbalancersEnableTarget(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"backends": [{"connecttimeout": 4000, "mode": "tcp", "name": "mybackend", "stickysessions": "no",
			"targets": [{"ipaddress": "8.8.8.8", "name": "mytarget", "port": 8080, "status": "MAINT", "weight": 10}]
			}],
		"loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}

	params := ToggleTargetParams{
		Backend: "mybackend",
		Name:    "mytarget",
	}

	lb.EnableTarget(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/enabletarget", c.lastPath, "path used is correct")
}

func TestLoadbalancersDisableTarget(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer": {
		"backends": [{"connecttimeout": 4000, "mode": "tcp", "name": "mybackend", "stickysessions": "no",
			"targets": [{"ipaddress": "8.8.8.8", "name": "mytarget", "port": 8080, "status": "MAINT", "weight": 10}]
			}],
		"loadbalancerid": "lb123456"}}}`}

	lb := LoadbalancerService{client: c}

	params := ToggleTargetParams{
		Backend: "mybackend",
		Name:    "mytarget",
	}

	lb.DisableTarget(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/disabletarget", c.lastPath, "path used is correct")
}

func TestLoadbalancersRemoveTarget(t *testing.T) {
	c := &mockClient{}
	lb := LoadbalancerService{client: c}

	params := RemoveTargetParams{
		Backend: "mybackend",
		Name:    "mytarget",
	}

	lb.RemoveTarget(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/removetarget", c.lastPath, "path used is correct")
}

func TestLoadbalancersAddtoblacklist(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancer":
	{ "backends": [], "blacklist": ["10.0.0.10/32"], "datacenter": "Falkenberg", "frontends": [],
	"ipaddress": [{"ipaddress": "192.168.0.1", "version": 4}],
	"name": "myloadbalancer", "loadbalancerid": "lb123456" }}}`}

	lb := LoadbalancerService{client: c}

	params := BlacklistParams{
		Prefix: "10.0.0.10/32",
	}

	lbd, _ := lb.Addtoblacklist(context.Background(), "lb123456", params)
	myprefix := strings.Join(lbd.Blacklists, " ")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/addtoblacklist", c.lastPath, "path used is correct")
	assert.Equal(t, "10.0.0.10/32", myprefix, "prefix set correct")
}

func TestLoadbalancersRemovefromblacklist(t *testing.T) {
	c := &mockClient{body: `{ "response": { "loadbalancers": {
		"backends": [{'name': 'my-backend'}], "blacklist": [],
		"name": "myloadbalancer", "loadbalancerid": "lb123456" }
	}}`}

	lb := LoadbalancerService{client: c}

	params := BlacklistParams{
		Prefix: "10.0.0.10/32",
	}

	lbd, _ := lb.Removefromblacklist(context.Background(), "lb123456", params)
	myprefix := strings.Join(lbd.Blacklists, " ")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/removefromblacklist", c.lastPath, "path used is correct")
	assert.Equal(t, "", myprefix, "prefix correctly absent")
}

func TestLoadbalancersAddCertificate(t *testing.T) {
	c := &mockClient{}

	lb := LoadbalancerService{client: c}

	params := AddCertificateParams{
		Name:        "mycert",
		Certificate: "ABC123==",
	}

	lb.AddCertificate(context.Background(), "lb123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/addcertificate", c.lastPath, "path used is correct")
}

func TestLoadbalancersListCertificate(t *testing.T) {
	c := &mockClient{}

	lb := LoadbalancerService{client: c}

	lb.ListCertificate(context.Background(), "lb123456")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/listcertificate", c.lastPath, "path used is correct")
}

func TestLoadbalancersRemoveCertificate(t *testing.T) {
	c := &mockClient{}

	lb := LoadbalancerService{client: c}

	lb.RemoveCertificate(context.Background(), "lb123456", "mycert")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "loadbalancer/removecertificate", c.lastPath, "path used is correct")
}
