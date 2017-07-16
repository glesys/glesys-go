package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworksCreate(t *testing.T) {
	c := &mockClient{body: `{ "response": { "network":
		{ "datacenter": "Falkenberg", "description": "mynetwork", "networkid": "vl123456" }}}`}
	n := NetworkService{client: c}

	params := CreateNetworkParams{
		Datacenter:  "Falkenberg",
		Description: "mynetwork",
	}

	network, _ := n.Create(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "network/create", c.lastPath, "path used is correct")
	assert.Equal(t, "vl123456", network.ID, "network ID is correct")
	assert.Equal(t, "Falkenberg", network.Datacenter, "network Datacenter is correct")
	assert.Equal(t, "mynetwork", network.Description, "network Description is correct")

}

func TestNetworksDestroy(t *testing.T) {
	c := &mockClient{}
	n := NetworkService{client: c}

	n.Destroy(context.Background(), "vl123456")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "network/delete", c.lastPath, "path used is correct")
}

func TestNetworksEdit(t *testing.T) {
	c := &mockClient{body: `{ "response": { "network":
		{ "datacenter": "Falkenberg", "description": "mynewnetwork", "networkid": "vl123456" }}}`}
	n := NetworkService{client: c}

	params := EditNetworkParams{
		Description: "mynetwork",
	}

	network, _ := n.Edit(context.Background(), "vl123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "network/edit", c.lastPath, "path used is correct")
	assert.Equal(t, "vl123456", network.ID, "network ID is correct")
	assert.Equal(t, "Falkenberg", network.Datacenter, "network Datacenter is correct")
	assert.Equal(t, "mynewnetwork", network.Description, "network Description is correct")
}

func TestNetworksList(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networks":
	[{ "datacenter": "Falkenberg", "description": "Internet", "networkid": "internet-fbg", "public": "yes"}] } }`}
	n := NetworkService{client: c}

	networks, _ := n.List(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "network/list", c.lastPath, "path used is correct")
	assert.Equal(t, "Falkenberg", (*networks)[0].Datacenter, "network Datacenter is correct")
	assert.Equal(t, "yes", (*networks)[0].Public, "network is public")
	assert.Equal(t, "Internet", (*networks)[0].Description, "network Description is correct")
	assert.Equal(t, "internet-fbg", (*networks)[0].ID, "network ID is correct")

}
