package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkAdapterCreate(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkadapter":
	{ "adaptertype": "E1000", "bandwidth": 10, "name": "Network Adapter 2", "networkid": "mynetwork",
	"networkadapterid": "ab12cd34-dcba-0123-abcd-abc123456789", "serverid": "vz123456" }}}`}
	n := NetworkAdapterService{client: c}

	params := CreateNetworkAdapterParams{
		Bandwidth: 10,
		NetworkID: "mynetwork",
		ServerID:  "vz123456",
	}

	networkadapter, _ := n.Create(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "networkadapter/create", c.lastPath, "path used is correct")
	assert.Equal(t, "vz123456", networkadapter.ServerID, "networkadapter ServerID is correct")
	assert.Equal(t, "Network Adapter 2", networkadapter.Name, "networkadapter Name is correct")
	assert.Equal(t, "mynetwork", networkadapter.NetworkID, "networkadapter Description is correct")

}

func TestNetworkAdapterDestroy(t *testing.T) {
	c := &mockClient{}
	n := NetworkAdapterService{client: c}

	n.Destroy(context.Background(), "ab12cd34-dcba-0123-abcd-abc123456789")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "networkadapter/delete", c.lastPath, "path used is correct")
}

func TestNetworkAdapterEdit(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkadapter":
	{ "adaptertype": "E1000", "bandwidth": 100, "name": "Network Adapter 2", "networkid": "mynewnetwork",
	"networkadapterid": "ab12cd34-dcba-0123-abcd-abc123456789", "serverid": "vz123456" }}}`}
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
