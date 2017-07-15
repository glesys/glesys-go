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

func TestIPsReserved(t *testing.T) {
	c := &mockClient{body: `{ "response": { "iplist": [{ "ipaddress": "127.0.0.1", "version": 4 },
		{ "ipaddress": "2001:db8::1", "version": 6 }] } }`}
	s := IPService{client: c}

	ips, _ := s.Reserved(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/listown", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ips)[0].Address, "one ip was returned")
	assert.Equal(t, 4, (*ips)[0].Version, "IPv4 address")
	assert.Equal(t, "2001:db8::1", (*ips)[1].Address, "one ip was returned")
	assert.Equal(t, 6, (*ips)[1].Version, "IPv6 address")
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
