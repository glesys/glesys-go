package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerDetailsIsLocked(t *testing.T) {
	serverDetails := ServerDetails{}

	assert.Equal(t, false, serverDetails.IsLocked(), "should not be locked")

	serverDetails.State = "locked"
	assert.Equal(t, true, serverDetails.IsLocked(), "should be locked")
}

func TestServerDetailsIsRunning(t *testing.T) {
	serverDetails := ServerDetails{}

	assert.Equal(t, false, serverDetails.IsRunning(), "should not be running")

	serverDetails.State = "running"
	assert.Equal(t, true, serverDetails.IsRunning(), "should be running")
}

func TestCreateServerParamsWithDefaults(t *testing.T) {
	params := CreateServerParams{}.WithDefaults()

	assert.Equal(t, 100, params.Bandwidth, "Bandwidth has correct default value")
	assert.Equal(t, 2, params.CPU, "CPU has correct default value")
	assert.Equal(t, "Falkenberg", params.DataCenter, "DataCenter has correct default value")
	assert.Equal(t, "any", params.IPv4, "IPv4 has correct default value")
	assert.Equal(t, "any", params.IPv6, "IPv6 has correct default value")
	assert.Equal(t, 2048, params.Memory, "Memory has correct default value")
	assert.Equal(t, "OpenVZ", params.Platform, "Platform has correct default value")
	assert.Equal(t, 50, params.Storage, "Storage has correct default value")
	assert.Equal(t, "Debian 8 64-bit", params.Template, "Template has correct default value")

	assert.NotEmpty(t, params.Hostname, "Hostname has a default value")
}

func TestCreateServerParamsCustomWithDefaults(t *testing.T) {
	params := CreateServerParams{
		DataCenter: "Stockholm",
		Memory:     4096,
	}.WithDefaults()

	assert.Equal(t, 100, params.Bandwidth, "Bandwidth has correct default value")
	assert.Equal(t, 2, params.CPU, "CPU has correct default value")
	assert.Equal(t, "Stockholm", params.DataCenter, "DataCenter has correct custom value")
	assert.Equal(t, "any", params.IPv4, "IPv4 has correct default value")
	assert.Equal(t, "any", params.IPv6, "IPv6 has correct default value")
	assert.Equal(t, 4096, params.Memory, "Memory has correct custom value")
	assert.Equal(t, "OpenVZ", params.Platform, "Platform has correct default value")
	assert.Equal(t, 50, params.Storage, "Storage has correct default value")
	assert.Equal(t, "Debian 8 64-bit", params.Template, "Template has correct default value")
}

func TestCreateServerParamsWithUsers(t *testing.T) {
	params := CreateServerParams{
		Bandwidth: 100,
		CPU: 2,
		DataCenter: "Falkenberg",
		IPv4: "any",
		IPv6: "any",
		Memory: 2048,
		Storage: 20,
		Platform: "KVM",
		Template: "ubuntu-18-04",
		Hostname: "kvmXXXXXXX",

	}.WithUser("glesys", []string{"ssh-rsa"}, "password")

	users := []User{{"glesys",
		 []string{"ssh-rsa"},
		 "password",
	}}

	assert.Equal(t, 100, params.Bandwidth, "Bandwidth has correct default value")
	assert.Equal(t, 2, params.CPU, "CPU has correct default value")
	assert.Equal(t, "Falkenberg", params.DataCenter, "DataCenter has correct default value")
	assert.Equal(t, "any", params.IPv4, "IPv4 has correct default value")
	assert.Equal(t, "any", params.IPv6, "IPv6 has correct default value")
	assert.Equal(t, 2048, params.Memory, "Memory has correct default value")
	assert.Equal(t, "KVM", params.Platform, "Platform has correct default value")
	assert.Equal(t, 20, params.Storage, "Storage has correct default value")
	assert.Equal(t, "ubuntu-18-04", params.Template, "Template has correct default value")
	assert.Equal(t, users, params.Users, "Users has correct default value")

	assert.NotEmpty(t, params.Hostname, "Hostname has a default value")
}

func TestServersCreate(t *testing.T) {
	c := &mockClient{body: `{ "response": { "server": { "serverid": "vz12345" } } }`}
	s := ServerService{client: c}

	server, _ := s.Create(context.Background(), CreateServerParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/create", c.lastPath, "path used is correct")
	assert.Equal(t, "vz12345", server.ID, "server ID is correct")
}

func TestServersDestroy(t *testing.T) {
	c := &mockClient{}
	s := ServerService{client: c}

	s.Destroy(context.Background(), "vz123456", DestroyServerParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/destroy", c.lastPath, "path used is correct")
}

func TestServersDetails(t *testing.T) {
	c := &mockClient{body: `{ "response": { "server": { "hostname": "my-server-123",
		"bandwidth": 100,
		"description": "MyServer",
		"templatename": "Debian 8 64-bit"
		} } }`}
	s := ServerService{client: c}

	server, _ := s.Details(context.Background(), "vz123456")

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/details/serverid/vz123456/includestate/yes", c.lastPath, "path used is correct")
	assert.Equal(t, "my-server-123", server.Hostname, "server Hostname is correct")
	assert.Equal(t, 100, server.Bandwidth, "server bandwidth is correct")
	assert.Equal(t, "MyServer", server.Description, "server Description is correct")
	assert.Equal(t, "Debian 8 64-bit", server.Template, "server Template is correct")
}

func TestServersEdit(t *testing.T) {
	c := &mockClient{}
	s := ServerService{client: c}

	s.Edit(context.Background(), "vz123456", EditServerParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/edit", c.lastPath, "path used is correct")
}

func TestServersList(t *testing.T) {
	c := &mockClient{body: `{ "response": { "servers": [{ "serverid": "vz12345" }] } }`}
	s := ServerService{client: c}

	servers, _ := s.List(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/list", c.lastPath, "path used is correct")
	assert.Equal(t, "vz12345", (*servers)[0].ID, "one server was returned")
}

func TestServersStart(t *testing.T) {
	c := &mockClient{}
	s := ServerService{client: c}

	s.Start(context.Background(), "vz123456")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/start", c.lastPath, "path used is correct")
}

func TestServersStop(t *testing.T) {
	c := &mockClient{}
	s := ServerService{client: c}

	s.Stop(context.Background(), "vz123456", StopServerParams{})

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/stop", c.lastPath, "path used is correct")
}

func TestGenerateHostnameReturnsAHostnameInTheCorrectFormat(t *testing.T) {
	hostname := generateHostname()
	assert.Regexp(t, "^\\w+-\\w+-\\d{3}$", hostname, "Hostname is dasherized and contains two words followed by a number")
}
