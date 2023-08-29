package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerDetailsIsLocked(t *testing.T) {
	serverDetails := ServerDetails{}

	assert.Equal(t, false, serverDetails.IsLocked, "should not be locked")

	serverDetails.IsLocked = true
	assert.Equal(t, true, serverDetails.IsLocked, "should be locked")
}

func TestServerDetailsIsRunning(t *testing.T) {
	serverDetails := ServerDetails{}

	assert.Equal(t, false, serverDetails.IsRunning, "should not be running")

	serverDetails.IsRunning = true
	assert.Equal(t, true, serverDetails.IsRunning, "should be running")
}

func TestCreateServerParamsWithDefaults(t *testing.T) {
	params := CreateServerParams{}.WithDefaults()

	assert.Equal(t, 100, params.Bandwidth, "Bandwidth has correct default value")
	assert.Equal(t, 2, params.CPU, "CPU has correct default value")
	assert.Equal(t, "Falkenberg", params.DataCenter, "DataCenter has correct default value")
	assert.Equal(t, "any", params.IPv4, "IPv4 has correct default value")
	assert.Equal(t, "any", params.IPv6, "IPv6 has correct default value")
	assert.Equal(t, 2048, params.Memory, "Memory has correct default value")
	assert.Equal(t, "KVM", params.Platform, "Platform has correct default value")
	assert.Equal(t, 50, params.Storage, "Storage has correct default value")
	assert.Equal(t, "Debian 11 (Bullseye)", params.Template, "Template has correct default value")

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
	assert.Equal(t, "KVM", params.Platform, "Platform has correct default value")
	assert.Equal(t, 50, params.Storage, "Storage has correct default value")
	assert.Equal(t, "Debian 11 (Bullseye)", params.Template, "Template has correct default value")
}

func TestCreateServerParamsWithUsers(t *testing.T) {
	params := CreateServerParams{
		Bandwidth:  100,
		CPU:        2,
		DataCenter: "Falkenberg",
		IPv4:       "any",
		IPv6:       "any",
		Memory:     2048,
		Storage:    20,
		Platform:   "KVM",
		Template:   "ubuntu-18-04",
		Hostname:   "kvmXXXXXXX",
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

func TestServersConsole(t *testing.T) {
	c := &mockClient{body: `{"response": {"console": {
		"host": "None", "password": "", "protocol": "",
		"url": "https://console.example.com/view/abc123456-ff00-aabb-ccdd-xyz987654321"}}} `}
	s := ServerService{client: c}

	console, _ := s.Console(context.Background(), "kvm123456")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/console", c.lastPath, "path used is correct")
	assert.Equal(t, "https://console.example.com/view/abc123456-ff00-aabb-ccdd-xyz987654321", console.URL, "server console url is correct")
}

func TestServersPreviewCloudConfig(t *testing.T) {
	c := &mockClient{body: `{"response":{
			"cloudconfig":{
				"preview": "#cloud-config\nusers:\n    -\n        name: bob\n        shell: /bin/bash\n        lock_passwd: false\n        sudo: 'ALL=(ALL) PASSWD:ALL'\n        passwd: $6$ecb46c3c2a73263f$wYkIrbHQzZ0zZvsb7PxdhIbskjOA4Ti5NnDe7EBBP.1SDAfborckfDcuYsqDmdgbGMFJgBzQMjXgJ4qHbLV5s.\n        ssh_authorized_keys: ['ssh-ed25519 AAAAKEY bob@bob-machine']\nssh_pwauth: false\nchpasswd:\n    expire: false\n",
				"context": {"params": { "foo": "bar", "balloon": 99 },
							"users": [{"username": "bob", "password": "hunter333", "sshKeys": ["ssh-ed25519 AAAAKEY bob@bob-machine"]}]
						   }
			}}}`}
	s := ServerService{client: c}

	cloudConfigParams := map[string]any{"foo": "bar", "balloon": 99}
	users := []User{}
	users = append(users, User{
		Username:   "bob",
		Password:   "hunter333",
		PublicKeys: []string{"ssh-ed25519 AAAAKEY bob@bob-machine"},
	})
	params := PreviewCloudConfigParams{
		CloudConfig:       "## template: glesys\n#cloud-config\n{{>users}}\n",
		CloudConfigParams: cloudConfigParams,
		Users:             users,
	}

	preview, _ := s.PreviewCloudConfig(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/previewcloudconfig", c.lastPath, "path used is correct")
	assert.Equal(t, "bob", preview.Context.Users[0].Username, "Preview contains user")
	assert.Equal(t, float64(99), preview.Context.Params["balloon"], "Preview contains parameter")
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
		"templatename": "Debian 11 64-bit",
		"backup": {"enabled": "yes", "schedules":
			[{"frequency": "daily", "numberofimagestokeep": 1}]}
		} } }`}
	s := ServerService{client: c}

	server, _ := s.Details(context.Background(), "kvm123456")

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/details/serverid/kvm123456/includestate/yes", c.lastPath, "path used is correct")
	assert.Equal(t, "my-server-123", server.Hostname, "server Hostname is correct")
	assert.Equal(t, 100, server.Bandwidth, "server bandwidth is correct")
	assert.Equal(t, "MyServer", server.Description, "server Description is correct")
	assert.Equal(t, "Debian 11 64-bit", server.Template, "server Template is correct")
	assert.Equal(t, "daily", server.Backup.Schedules[0].Frequency, "Backup schedule is daily")
	assert.Equal(t, 1, server.Backup.Schedules[0].Numberofimagestokeep, "Backup images to keep is correct")
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

func TestServersTemplates(t *testing.T) {
	c := &mockClient{body: `{"response":{ "templates": { "KVM": [{"id": "ac7c05f1-4cb6-4330-a0a2-d1f2e6244b21",
     "name": "AlmaLinux 8", "minimumdisksize": 5, "minimummemorysize": 512, "operatingsystem": "linux", "platform": "KVM",
     "instancecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"},
     "licensecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"}, "bootstrapmethod": "CLOUD_INIT"},
    {"id": "2563b4d0-ea80-4aef-8f77-f9b9e479a008", "name": "AlmaLinux 9", "minimumdisksize": 5, "minimummemorysize": 512,
     "operatingsystem": "linux", "platform": "KVM", "instancecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"},
     "licensecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"}, "bootstrapmethod": "CLOUD_INIT"}],
   "VMware": [{"id": "420fe17c-bc03-4b2c-a741-7a790e5f21ad", "name": "Alma Linux 8", "minimumdisksize": 5, "minimummemorysize": 512,
     "operatingsystem": "linux", "platform": "VMware", "instancecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"},
     "licensecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"}, "bootstrapmethod": "DEPLOY_SCRIPT"},
    {"id": "dbbca8a7-1e26-4b76-8bb4-8d56dae59039", "name": "Alma Linux 9", "minimumdisksize": 5, "minimummemorysize": 512,
     "operatingsystem": "linux", "platform": "VMware", "instancecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"},
     "licensecost": {"amount": 0, "currency": "SEK", "timeperiod": "month"}, "bootstrapmethod": "CLOUD_INIT"},
    {"id": "d924551c-0a0d-43ba-abcd-aoeuqwer1234", "name": "Windows Server 2022 Standard LTSC", "minimumdisksize": 30,
     "minimummemorysize": 1024, "operatingsystem": "windows", "platform": "VMware", "instancecost":
     {"amount": 999.10, "currency": "SEK", "timeperiod": "month"}, "licensecost": {"amount": 123.4, "currency": "SEK",
     "timeperiod": "month"}, "bootstrapmethod": "DEPLOY_SCRIPT"}]}}}`}
	s := ServerService{client: c}

	templates, _ := s.Templates(context.Background())

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "server/templates", c.lastPath, "path used is correct")
	assert.Equal(t, 512, templates.KVM[0].MinMemSize, "template minmemsize is correct")
	assert.Equal(t, "CLOUD_INIT", templates.KVM[0].BootstrapMethod, "template bootstrapmethod is correct")
	assert.Equal(t, "Alma Linux 8", templates.VMware[0].Name, "template name is correct")
	assert.Equal(t, "DEPLOY_SCRIPT", templates.VMware[0].BootstrapMethod, "template bootstrapmethod is correct")
	assert.Equal(t, 0.0, templates.VMware[0].LicenseCost.Amount, "template licensecost amount is correct")
	assert.Equal(t, "month", templates.VMware[0].LicenseCost.Timeperiod, "template licensecost timeperiod is correct")
	assert.Equal(t, 999.1, templates.VMware[2].InstanceCost.Amount, "template instancecost amount is correct")
	assert.Equal(t, 123.4, templates.VMware[2].LicenseCost.Amount, "template licensecost amount is correct")
	assert.Equal(t, "month", templates.VMware[2].LicenseCost.Timeperiod, "template licensecost timeperiod is correct")
}

func TestGenerateHostnameReturnsAHostnameInTheCorrectFormat(t *testing.T) {
	hostname := generateHostname()
	assert.Regexp(t, "^\\w+-\\w+-\\d{3}$", hostname, "Hostname is dasherized and contains two words followed by a number")
}
