package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkAdapterIsLocked(t *testing.T) {
	networkadapter := NetworkAdapter{}

	assert.Equal(t, false, networkadapter.IsLocked(), "should not be locked")

	networkadapter.State = "locked"
	assert.Equal(t, true, networkadapter.IsLocked(), "should be locked")
}

func TestNetworkAdapterIsReady(t *testing.T) {
	networkadapter := NetworkAdapter{}

	assert.Equal(t, false, networkadapter.IsReady(), "should not be ready")

	networkadapter.State = "ready"
	assert.Equal(t, true, networkadapter.IsReady(), "should be ready")
}

func TestNetworkAdaptersCreate(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkadapter":
	{ "adaptertype": "E1000", "bandwidth": 10, "name": "Network Adapter 2", "networkid": "mynetwork",
	"networkadapterid": "ab12cd34-dcba-0123-abcd-abc123456789", "serverid": "wps123456" }}}`}
	n := NetworkAdapterService{client: c}

	params := CreateNetworkAdapterParams{
		Bandwidth: 10,
		NetworkID: "mynetwork",
		ServerID:  "wps123456",
	}

	networkadapter, _ := n.Create(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "networkadapter/create", c.lastPath, "path used is correct")
	assert.Equal(t, "wps123456", networkadapter.ServerID, "networkadapter ServerID is correct")
	assert.Equal(t, "Network Adapter 2", networkadapter.Name, "networkadapter Name is correct")
	assert.Equal(t, "mynetwork", networkadapter.NetworkID, "networkadapter Description is correct")
}

func TestNetworkAdaptersCreate_KVM(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkadapter":
		{ "bandwidth": 1000, "name": "Adapter Example", "isprimary": false, "isconnected": true,
		"networkid": "79c67265-a9a8-4607-b2b5-7377a6b6ebf7",
		"networkadapterid": "ab12cd34-dcba-0123-abcd-abc123456789",
		"serverid": "kvm123456" }}}`}
	n := NetworkAdapterService{client: c}

	params := CreateNetworkAdapterParams{
		Bandwidth: 1000,
		NetworkID: "79c67265-a9a8-4607-b2b5-7377a6b6ebf7",
		ServerID:  "kvm123456",
		Name:      "Adapter Example",
	}

	networkadapter, _ := n.Create(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "networkadapter/create", c.lastPath, "path used is correct")
	assert.Equal(t, "kvm123456", networkadapter.ServerID, "networkadapter ServerID is correct")
	assert.Equal(t, "Adapter Example", networkadapter.Name, "networkadapter Name is correct")
	assert.Equal(t, true, networkadapter.IsConnected, "networkadapter IsConnected is correct")
	assert.Equal(t, false, networkadapter.IsPrimary, "networkadapter IsPrimary is correct")
	assert.Equal(t, "79c67265-a9a8-4607-b2b5-7377a6b6ebf7", networkadapter.NetworkID, "networkadapter networkID is correct")
}

func TestNetworkAdaptersDestroy(t *testing.T) {
	c := &mockClient{}
	n := NetworkAdapterService{client: c}

	n.Destroy(context.Background(), "ab12cd34-dcba-0123-abcd-abc123456789")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "networkadapter/delete", c.lastPath, "path used is correct")
}

func TestNetworkAdaptersDetails(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkadapter": {
		"networkadapterid": "9ac61694-eb4d-4011-9d10-c395ba5f7269",
		"bandwidth": 100,
		"name": "My Network Adapter",
		"adaptertype": "VMXNET 3",
		"state": "ready",
		"serverid": "wps123456",
		"networkid": "internet-fbg"
		} } }`}
	s := NetworkAdapterService{client: c}

	networkAdapter, _ := s.Details(context.Background(), "9ac61694-eb4d-4011-9d10-c395ba5f7269")

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "networkadapter/details/networkadapterid/9ac61694-eb4d-4011-9d10-c395ba5f7269", c.lastPath, "path used is correct")
	assert.Equal(t, "VMXNET 3", networkAdapter.AdapterType, "network adapter Bandwidth is correct")
	assert.Equal(t, 100, networkAdapter.Bandwidth, "network adapter Bandwidth is correct")
	assert.Equal(t, "9ac61694-eb4d-4011-9d10-c395ba5f7269", networkAdapter.ID, "network adapter Bandwidth is correct")
	assert.Equal(t, "My Network Adapter", networkAdapter.Name, "network adapter Bandwidth is correct")
	assert.Equal(t, "internet-fbg", networkAdapter.NetworkID, "network adapter Bandwidth is correct")
	assert.Equal(t, "wps123456", networkAdapter.ServerID, "network adapter Bandwidth is correct")
	assert.Equal(t, "ready", networkAdapter.State, "network adapter Bandwidth is correct")
}

func TestNetworkAdaptersEdit(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkadapter":
	{ "adaptertype": "E1000", "bandwidth": 100, "name": "Network Adapter 2", "networkid": "mynewnetwork",
	"networkadapterid": "ab12cd34-dcba-0123-abcd-abc123456789", "serverid": "wps123456" }}}`}
	n := NetworkAdapterService{client: c}

	params := EditNetworkAdapterParams{
		Bandwidth: 100,
		NetworkID: "mynewnetwork",
	}

	networkadapter, _ := n.Edit(context.Background(), "ab12cd34-dcba-0123-abcd-abc123456789", params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "networkadapter/edit", c.lastPath, "path used is correct")
	assert.Equal(t, "ab12cd34-dcba-0123-abcd-abc123456789", networkadapter.ID, "networkadapter ID is correct")
	assert.Equal(t, "mynewnetwork", networkadapter.NetworkID, "networkadapter network ID is correct")
	assert.Equal(t, 100, networkadapter.Bandwidth, "networkadapter Bandwidth is correct")
}
