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

func TestIPsDetails(t *testing.T) {
	c := &mockClient{body: `{ "response": { "details": { "ipaddress": "127.0.0.1",
		"netmask": "None", "broadcast": "None", "gateway": "None", "nameservers": ["127.255.255.1"],
		"platform": "KVM", "platforms": ["KVM"],
		"cost": {"amount":1, "currency": "SEK", "timeperiod": "month"},
		"datacenter": "Falkenberg", "ipversion": 4, "serverid": "kvm123456", "reserved": "yes",
		"lockedtoaccount": "no", "ptr": "1.0.0.127-static.example.com."} } }`}
	s := IPService{client: c}

	ip, _ := s.Details(context.Background(), "127.0.0.1")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/details", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ip).Address, "one ip was returned")
	assert.Equal(t, "KVM", (*ip).Platform, "platform is correct")
	assert.Equal(t, "Falkenberg", (*ip).DataCenter, "datacenter is correct")
	assert.Equal(t, "1.0.0.127-static.example.com.", (*ip).PTR, "ptr is correct")
	assert.Equal(t, 1.00, (*ip).Cost.Amount, "cost amount is correct")
}

func TestIPsReserved(t *testing.T) {
	c := &mockClient{body: `{ "response": { "iplist": [{ "ipaddress": "127.0.0.1", "datacenter": "Falkenberg" },
		{ "ipaddress": "2001:db8::1", "datacenter": "Falkenberg" }] } }`}
	s := IPService{client: c}

	params := ReservedIPsParams{
		DataCenter: "Falkenberg",
	}

	ips, _ := s.Reserved(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/listown", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ips)[0].Address, "one ip was returned")
	assert.Equal(t, "2001:db8::1", (*ips)[1].Address, "one ip was returned")
	assert.Equal(t, "Falkenberg", (*ips)[0].DataCenter, "correct DataCenter was returned")
	assert.Equal(t, "Falkenberg", (*ips)[1].DataCenter, "correct DataCenter was returned")
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

func TestIPs_SetPTR(t *testing.T) {
	c := &mockClient{body: `{ "response": { "details": { "ipaddress": "127.0.0.1",
		"netmask": "None", "broadcast": "None", "gateway": "None", "nameservers": ["127.255.255.1"],
		"platform": "KVM", "platforms": ["KVM"],
		"cost": {"amount":1, "currency": "SEK", "timeperiod": "month"},
		"datacenter": "Falkenberg", "ipversion": 4, "serverid": "kvm123456", "reserved": "yes",
		"lockedtoaccount": "no", "ptr": "ptr.parker.example.com."} } }`}
	s := IPService{client: c}

	ip, _ := s.SetPTR(context.Background(), "127.0.0.1", "ptr.parker.example.com.")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/setptr", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ip).Address, "one ip was returned")
	assert.Equal(t, "ptr.parker.example.com.", (*ip).PTR, "ptr is correct")
}

func TestIPs_ResetPTR(t *testing.T) {
	c := &mockClient{body: `{ "response": { "details": { "ipaddress": "127.0.0.1",
		"netmask": "None", "broadcast": "None", "gateway": "None", "nameservers": ["127.255.255.1"],
		"platform": "KVM", "platforms": ["KVM"],
		"cost": {"amount":1, "currency": "SEK", "timeperiod": "month"},
		"datacenter": "Falkenberg", "ipversion": 4, "serverid": "kvm123456", "reserved": "yes",
		"lockedtoaccount": "no", "ptr": "1-0-0-127-static.glesys.net."} } }`}
	s := IPService{client: c}

	ip, _ := s.ResetPTR(context.Background(), "127.0.0.1")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "ip/resetptr", c.lastPath, "path used is correct")
	assert.Equal(t, "127.0.0.1", (*ip).Address, "one ip was returned")
	assert.Equal(t, "1-0-0-127-static.glesys.net.", (*ip).PTR, "ptr is correct")
}
