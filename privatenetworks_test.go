package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateNetworksCreate(t *testing.T) {
	c := &mockClient{body: `{ "response": { "privatenetwork":
		{ "id": "pn-123ab", "name": "mynetwork", "ipv6aggregate": "2001:db8::/48"}}}`}
	n := PrivateNetworkService{client: c}

	networkname := "mynetwork"

	net, _ := n.Create(context.Background(), networkname)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "privatenetwork/create", c.lastPath, "path used is correct")
	assert.Equal(t, "pn-123ab", net.ID, "ID is correct")
	assert.Equal(t, "mynetwork", net.Name, "Name is correct")
	assert.Equal(t, "2001:db8::/48", net.IPv6Aggregate, "IPv6Aggregate is correct")
}

func TestPrivateNetworksDestroy(t *testing.T) {
	c := &mockClient{}
	n := PrivateNetworkService{client: c}

	n.Destroy(context.Background(), "pn-123ab")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "privatenetwork/delete", c.lastPath, "path used is correct")
}

func TestPrivateNetworksDetails(t *testing.T) {
	c := &mockClient{body: `{ "response": { "privatenetwork": {
		"id": "pn-123ab",
		"name": "mynetwork",
		"ipv6aggregate": "2001:db8::/48"
		} } }`}
	n := PrivateNetworkService{client: c}

	net, _ := n.Details(context.Background(), "pn-123ab")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "privatenetwork/details", c.lastPath, "path used is correct")
	assert.Equal(t, "mynetwork", net.Name, "Name is correct")
	assert.Equal(t, "pn-123ab", net.ID, "ID is correct")
	assert.Equal(t, "2001:db8::/48", net.IPv6Aggregate, "IPv6Aggregate is correct")
}

func TestPrivateNetworksList(t *testing.T) {
	c := &mockClient{body: `{ "response": { "privatenetworks": [{
		"id": "pn-123ab",
		"name": "mynetwork",
		"ipv6aggregate": "2001:db8:1::/48"
		}, {
		"id": "pn-456cd",
		"name": "othernet",
		"ipv6aggregate": "2001:db8:2::/48" } ] } }`}

	n := PrivateNetworkService{client: c}

	net, _ := n.List(context.Background())

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "privatenetwork/list", c.lastPath, "path used is correct")
	assert.Equal(t, "mynetwork", (*net)[0].Name, "Name is correct")
	assert.Equal(t, "pn-123ab", (*net)[0].ID, "ID is correct")
	assert.Equal(t, "2001:db8:1::/48", (*net)[0].IPv6Aggregate, "IPv6Aggregate is correct")
	assert.Equal(t, "othernet", (*net)[1].Name, "Name is correct")
	assert.Equal(t, "pn-456cd", (*net)[1].ID, "ID is correct")
	assert.Equal(t, "2001:db8:2::/48", (*net)[1].IPv6Aggregate, "IPv6Aggregate is correct")
}

func TestPrivateNetworksEdit(t *testing.T) {
	c := &mockClient{body: `{ "response": { "privatenetwork":
		{ "name": "newnetwork-1", "id": "pn-123ab"}}}`}

	n := PrivateNetworkService{client: c}

	params := EditPrivateNetworkParams{
		Name: "newnetwork-1",
		ID:   "pn-123ab",
	}

	net, _ := n.Edit(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "privatenetwork/edit", c.lastPath, "path used is correct")
	assert.Equal(t, "pn-123ab", net.ID, "ID is correct")
	assert.Equal(t, "newnetwork-1", net.Name, "Name is correct")
}

func TestPrivateNetworksEstimatedcost(t *testing.T) {
	c := &mockClient{body: `{ "response": { "billing":
		{ "currency": "SEK", "price": 50.0,
		"discount": 12.5, "total": 37.5}}}`}

	n := PrivateNetworkService{client: c}

	netid := "pn-123ab"

	net, _ := n.EstimatedCost(context.Background(), netid)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "privatenetwork/estimatedcost", c.lastPath, "path used is correct")
	assert.Equal(t, 50.0, net.Price, "Price is correct")
	assert.Equal(t, 12.5, net.Discount, "Discount is correct")
	assert.Equal(t, 37.5, net.Total, "Total price is correct")
}

func TestPrivateNetworkSegmentsCreate(t *testing.T) {
	c := &mockClient{body: `{ "response": { "privatenetworksegment":
		{ "id": "266979ab-1e05-4fbc-b9e0-577f31c0d2e9",
		"name": "mysegment", "ipv6subnet": "2001:db8:0::/64",
		"ipv4subnet": "192.0.2.0/24", "datacenter": "dc-fbg1",
		"platform": "kvm"}}}`}
	n := PrivateNetworkService{client: c}

	params := CreatePrivateNetworkSegmentParams{
		Name:       "mysegment",
		Platform:   "kvm",
		Datacenter: "dc-fbg1",
		IPv4Subnet: "192.0.2.0/24",
	}

	segment, _ := n.CreateSegment(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "privatenetwork/createsegment", c.lastPath, "path used is correct")
	assert.Equal(t, "266979ab-1e05-4fbc-b9e0-577f31c0d2e9", segment.ID, "ID is correct")
	assert.Equal(t, "mysegment", segment.Name, "Name is correct")
	assert.Equal(t, "2001:db8:0::/64", segment.IPv6Subnet, "IPv6Subnet is correct")
	assert.Equal(t, "192.0.2.0/24", segment.IPv4Subnet, "IPv4Subnet is correct")
}

func TestPrivateNetworkSegmentsDestroy(t *testing.T) {
	c := &mockClient{}
	n := PrivateNetworkService{client: c}

	n.DestroySegment(context.Background(), "pn-123ab")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "privatenetwork/deletesegment", c.lastPath, "path used is correct")
}

func TestPrivateNetworkSegmentsList(t *testing.T) {
	c := &mockClient{body: `
		{ "response": { "privatenetworksegments": [{
		"id": "fb34a19a-392a-43ec-ab3f-0c5b73ad1234",
		"name": "mysegment",
		"platform": "kvm",
		"datacenter": "dc-fbg1",
		"ipv4subnet": "192.0.2.0/24",
		"ipv6subnet": "2001:db8:0::/64"
		}, {
		"id": "fb34a19a-392a-43ec-ab3f-0c5b73ad5678",
		"name": "othersegment",
		"platform": "kvm",
		"datacenter": "dc-fbg1",
		"ipv4subnet": "192.0.2.0/24",
		"ipv6subnet": "2001:db8:1::/64" } ] } }
		`}

	n := PrivateNetworkService{client: c}

	segments, _ := n.ListSegments(context.Background(), "pn-123ab")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "privatenetwork/listsegments", c.lastPath, "path used is correct")
	assert.Equal(t, "mysegment", (*segments)[0].Name, "Name is correct")
	assert.Equal(t, "fb34a19a-392a-43ec-ab3f-0c5b73ad1234", (*segments)[0].ID, "ID is correct")
	assert.Equal(t, "2001:db8:0::/64", (*segments)[0].IPv6Subnet, "IPv6Aggregate is correct")
	assert.Equal(t, "othersegment", (*segments)[1].Name, "Name is correct")
	assert.Equal(t, "fb34a19a-392a-43ec-ab3f-0c5b73ad5678", (*segments)[1].ID, "ID is correct")
	assert.Equal(t, "kvm", (*segments)[1].Platform, "Platform is correct")
	assert.Equal(t, "2001:db8:1::/64", (*segments)[1].IPv6Subnet, "IPv6Aggregate is correct")
}

func TestPrivateNetworkSegmentsEdit(t *testing.T) {
	c := &mockClient{body: `{ "response": { "privatenetworksegment":
		{ "name": "segmentname-2", "id": "fb34a19a-392a-43ec-ab3f-0c5b73ad1234"}}}`}

	n := PrivateNetworkService{client: c}

	params := EditPrivateNetworkSegmentParams{
		Name: "segmentname-2",
		ID:   "fb34a19a-392a-43ec-ab3f-0c5b73ad1234",
	}

	segment, _ := n.EditSegment(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "privatenetwork/editsegment", c.lastPath, "path used is correct")
	assert.Equal(t, "fb34a19a-392a-43ec-ab3f-0c5b73ad1234", segment.ID, "ID is correct")
	assert.Equal(t, "segmentname-2", segment.Name, "Name is correct")
}
