package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPsAvailable(t *testing.T) {
	c := &mockClient{body: `{ "response": { "iplist": { "ipaddresses": ["127.0.0.1"] }} }`}
	s := IPService{client: c}

	ips, _ := s.Available(context.Background(), AvailableIPsParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/listfree", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ips)[0].Address, "one ip was returned")
}

func TestIPsIsIPv4(t *testing.T) {
	var ips = []IP{
		{Address: "127.0.0.1"},
		{Address: "300.0.0.1"},
		{Address: "2001:db8::1"},
	}

	assert.Equal(t, true, ips[0].IsIPv4(), "ip is version 4")
	assert.Equal(t, false, ips[1].IsIPv4(), "ip is not version 4")
	assert.Equal(t, false, ips[2].IsIPv4(), "ip is not version 4")
}

func TestIPsIsIPv6(t *testing.T) {
	var ips = []IP{
		{Address: "2001:db8::1"},
		{Address: "::1"},
		{Address: "300.0.0.1"},
		{Address: "::2001::1"},
	}

	assert.Equal(t, true, ips[0].IsIPv6(), "ip is version 6")
	assert.Equal(t, true, ips[1].IsIPv6(), "ip is version 6")
	assert.Equal(t, false, ips[2].IsIPv6(), "ip is not version 6")
	assert.Equal(t, false, ips[3].IsIPv6(), "ip is not version 6")
}

func TestIPsReserved(t *testing.T) {
	c := &mockClient{body: `{ "response": { "iplist": [{ "ipaddress": "127.0.0.1"},
		{ "ipaddress": "2001:db8::1"}] } }`}
	s := IPService{client: c}

	ips, _ := s.Reserved(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/listown", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ips)[0].Address, "one ip was returned")
	assert.Equal(t, "2001:db8::1", (*ips)[1].Address, "one ip was returned")
}

func TestIPsReserve(t *testing.T) {
	c := &mockClient{body: `{ "response": { "details": { "ipaddress": "127.0.0.1" } } }`}
	s := IPService{client: c}

	ip, _ := s.Reserve(context.Background(), "127.0.0.1")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/take", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ip).Address, "one ip was returned")
}

func TestIPService_Release(t *testing.T) {
	c := &mockClient{}
	s := IPService{client: c}

	s.Release(context.Background(), "127.0.0.1")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/release", c.lastPath, "path used is correct")
}
