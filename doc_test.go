package glesys_test

import (
	"context"
	"fmt"

	glesys "github.com/glesys/glesys-go/v8"
)

func ExampleUserService_DoOTPLogin() {
	userAgent := "MyGleSYSClient 0.0.1"
	login := glesys.NewLogin(userAgent)

	loginDetails, err := login.Users.DoOTPLogin(
		context.Background(),
		"user@example.com",
		"VerySecretPassword123",
		"abc123-otpstring")

	if err != nil {
		fmt.Printf("Error logging in %s\n", err)
	}

	// Set the temporary key to the Login object
	login.Username = loginDetails.Username
	login.APIKey = loginDetails.APIKey

	// Now you can run user specific calls to the api
	// list projects for organization 12345
	projects, err := login.Users.ListCustomerProjects(context.Background(), "12345")
	if err != nil {
		fmt.Printf("Error listing projects %s\n", err)
	}
	for _, project := range *projects {
		fmt.Printf("Name: %s, ProjectID: %s\n", project.Name, project.Accountname)
	}
}

func ExampleEmailDomainService_Overview() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: All parameters in OverviewParams are optional and can be omitted.
	overview, _ := client.EmailDomains.Overview(context.Background(), glesys.OverviewParams{
		Filter: "example.com",
		Page:   1,
	})

	fmt.Printf("%#v", overview)
}

func ExampleEmailDomainService_List() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: The filter in ListEmailsParams is optional and can be omitted.
	list, _ := client.EmailDomains.List(context.Background(), "example.com", glesys.ListEmailsParams{
		Filter: "user@example.com",
	})

	fmt.Printf("%#v\n", list)
}

func ExampleEmailDomainService_EditAccount() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: All parameters are optional and can be omitted.
	editaccount, _ := client.EmailDomains.EditAccount(context.Background(), "user@example.com", glesys.EditAccountParams{
		AntiSpamLevel:      3,
		AntiVirus:          "yes",
		AutoRespond:        "yes",
		AutoRespondMessage: "Your Automatic Response",
		QuotaInGiB:         10,
		RejectSpam:         "yes",
	})

	fmt.Printf("%#v\n", editaccount)
}

func ExampleEmailDomainService_Delete() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: The email parameter can be both an account and an alias.
	client.EmailDomains.Delete(context.Background(), "user@example.com")
}

func ExampleEmailDomainService_CreateAccount() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: All parameters except for EmailAccount and Password are optional and can be omitted.
	createaccount, _ := client.EmailDomains.CreateAccount(context.Background(), glesys.CreateAccountParams{
		EmailAccount:       "new_user@example.com",
		Password:           "SuperSecretPassword",
		AntiSpamLevel:      3,
		AntiVirus:          "yes",
		AutoRespond:        "yes",
		AutoRespondMessage: "Your Automatic Response",
		QuotaInGiB:         10,
		RejectSpam:         "yes",
	})

	fmt.Printf("%#v\n", createaccount)
}

func ExampleEmailQuota() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	quota, _ := client.EmailDomains.Quota(context.Background(), "user@example.com")

	fmt.Printf("%#v\n", quota)
}

func ExampleEmailDomainService_CreateAlias() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	alias, _ := client.EmailDomains.CreateAlias(context.Background(), glesys.EmailAliasParams{
		EmailAlias: "alias@example.com",
		GoTo:       "user@example.com",
	})

	fmt.Printf("%#v\n", alias)
}

func ExampleEmailDomainService_EditAlias() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	alias, _ := client.EmailDomains.EditAlias(context.Background(), glesys.EmailAliasParams{
		EmailAlias: "alias@example.com",
		GoTo:       "another_user@example.com",
	})

	fmt.Printf("%#v\n", alias)
}

func ExampleEmailDomainService_Costs() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	costs, _ := client.EmailDomains.Costs(context.Background())

	fmt.Printf("%#v\n", costs)
}

func ExampleIPService_Available() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	ips, _ := client.IPs.Available(context.Background(), glesys.AvailableIPsParams{
		DataCenter: "Falkenberg",
		Platform:   "KVM",
		Version:    4,
	})

	for _, ip := range *ips {
		fmt.Println(ip.Address)
	}
}

func ExampleIPService_Release() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.IPs.Release(context.Background(), "1.2.3.4")
}

func ExampleIPService_Reserve() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	ip, _ := client.IPs.Reserve(context.Background(), "1.2.3.4")

	fmt.Println(ip.Address)
}

func ExampleIPService_Reserved() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	ips, _ := client.IPs.Reserved(context.Background(), glesys.ReservedIPsParams{})

	for _, ip := range *ips {
		fmt.Println(ip.Address)
	}
}

func ExampleLoadBalancerService_Create() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	loadbalancer, _ := client.LoadBalancers.Create(context.Background(),
		glesys.CreateLoadBalancerParams{
			DataCenter: "Falkenberg",
			Name:       "mylb-1",
		})

	fmt.Println(loadbalancer.ID)
}

func ExampleLoadBalancerService_Destroy() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.LoadBalancers.Destroy(context.Background(), "lb123456")
}

func ExampleLoadBalancerService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	loadbalancer, _ := client.LoadBalancers.Details(context.Background(), "lb123456")

	fmt.Println(loadbalancer.Name)
}

func ExampleLoadBalancerService_AddBackend() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	loadbalancer, _ := client.LoadBalancers.AddBackend(context.Background(), "lb123456",
		glesys.AddBackendParams{
			Name: "mywebbackend",
			Mode: "http",
		})

	// print the name of all backends for the LoadBalancer
	for i := range (*loadbalancer).BackendsList {
		be := (*loadbalancer).BackendsList[i]
		fmt.Println("Name:", be.Name)
	}
}

func ExampleLoadBalancerService_AddTarget() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	loadbalancer, _ := client.LoadBalancers.AddTarget(context.Background(), "lb123456",
		glesys.AddTargetParams{
			Backend:  "mywebbackend",
			Name:     "web01",
			Port:     8080,
			TargetIP: "172.17.0.10",
			Weight:   5,
		})

	for i := range (*loadbalancer).BackendsList {
		be := (*loadbalancer).BackendsList[i]
		for k := range be.Targets {
			fmt.Printf("Backend: %s, Target: %s\n", be.Name, be.Targets[k].Name)
		}
	}
}

func ExampleLoadBalancerService_DisableTarget() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	loadbalancer, _ := client.LoadBalancers.DisableTarget(context.Background(), "lb123456",
		glesys.ToggleTargetParams{
			Backend: "mywebbackend",
			Name:    "web01",
		})

	for i := range (*loadbalancer).BackendsList {
		be := (*loadbalancer).BackendsList[i]
		for k := range be.Targets {
			fmt.Printf("Backend: %s, Target: %s, Status: %s\n", be.Name, be.Targets[k].Name, be.Targets[k].Status)
		}
	}
}

func ExampleLoadBalancerService_AddFrontend() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	loadbalancer, _ := client.LoadBalancers.AddFrontend(context.Background(), "lb123456",
		glesys.AddFrontendParams{
			Name:           "mywebfrontend",
			Backend:        "mywebbackend",
			Port:           80,
			ClientTimeout:  4000,
			MaxConnections: 1000,
		})

	// print the name of all frontends for the LoadBalancer
	for i := range (*loadbalancer).FrontendsList {
		fe := (*loadbalancer).FrontendsList[i]
		fmt.Println("Name:", fe.Name)
	}
}

func ExampleLoadBalancerService_AddCertificate() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	mybase64pem := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCm15Y2VydAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLQpteWtleQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg=="

	client.LoadBalancers.AddCertificate(context.Background(), "lb123456", glesys.AddCertificateParams{
		Name:        "mycert",
		Certificate: mybase64pem,
	})
}

func ExampleLoadBalancerService_ListCertificates() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	certlist, _ := client.LoadBalancers.ListCertificates(context.Background(), "lb123456")

	for _, cert := range *certlist {
		fmt.Println("Certificate:", cert)
	}
}

func ExampleLoadBalancerService_RemoveCertificate() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	err := client.LoadBalancers.RemoveCertificate(context.Background(), "lb123456", "mycert")

	if err != nil {
		fmt.Printf("Error removing certificate: %s\n", err)
	}
}

func ExampleNetworkAdapterService_Create() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: All parameters except ServerID are optional, the values shown below
	// are defaults and can be omitted.

	networkAdapter, _ := client.NetworkAdapters.Create(context.Background(), glesys.CreateNetworkAdapterParams{
		AdapterType: "VMXNET 3", // "E1000" also available
		Bandwidth:   100,
		NetworkID:   "internet-fbg",
		ServerID:    "wps123456",
	})

	fmt.Println(networkAdapter.ID)
}

func ExampleNetworkAdapterService_Destroy() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.NetworkAdapters.Destroy(context.Background(), "f590b422-453c-4fc4-99e7-af2b72a60f63")
}

func ExampleNetworkAdapterService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	networkAdapter, _ := client.NetworkAdapters.Details(context.Background(), "f590b422-453c-4fc4-99e7-af2b72a60f63")

	fmt.Println(networkAdapter.Name)
}

func ExampleNetworkAdapterService_Edit() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: Changes are not reflected immediately.

	networkAdapter, _ := client.NetworkAdapters.Edit(context.Background(), "f590b422-453c-4fc4-99e7-af2b72a60f63", glesys.EditNetworkAdapterParams{
		Bandwidth: 200,
		NetworkID: "vl12345",
	})

	fmt.Println(networkAdapter.Name)
}

func ExampleNetworkService_Create() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	network, _ := client.Networks.Create(context.Background(), glesys.CreateNetworkParams{
		DataCenter:  "Falkenberg",
		Description: "My Network",
	})

	fmt.Println(network.ID)
}

func ExampleNetworkService_Destroy() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.Networks.Destroy(context.Background(), "vl123456")
}

func ExampleNetworkService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	network, _ := client.Networks.Details(context.Background(), "vl123456")

	fmt.Println(network.Description)
}

func ExampleNetworkService_Edit() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	network, _ := client.Networks.Edit(context.Background(), "vl123456", glesys.EditNetworkParams{
		Description: "My Private Network",
	})

	fmt.Println(network.Description)
}

func ExampleNetworkService_List() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	networks, _ := client.Networks.List(context.Background())

	for _, network := range *networks {
		fmt.Println(network.ID)
	}
}

func ExampleNetworkCircuitService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	ic, _ := client.NetworkCircuits.Details(context.Background(), "ic123456")

	fmt.Println(ic.Type)
}

func ExampleNetworkCircuitService_List() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	ics, _ := client.NetworkCircuits.List(context.Background())

	for _, ic := range *ics {
		fmt.Println(ic.ID)
		fmt.Println(ic.Type)
	}
}

func ExamplePrivateNetworkService_Create() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	privnet, _ := client.PrivateNetworks.Create(context.Background(), "my-network")

	fmt.Printf("ID: %s - Name: %s - IPv6: %s\n", privnet.ID, privnet.Name, privnet.IPv6Aggregate)
}

func ExamplePrivateNetworkService_Destroy() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	err := client.PrivateNetworks.Destroy(context.Background(), "pn-123ab")

	if err != nil {
		fmt.Printf("Cannot remove private network. Error: %s", err)
	}
}

func ExamplePrivateNetworkService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	privnet, _ := client.PrivateNetworks.Details(context.Background(), "pn-123ab")

	fmt.Printf("ID: %s - Name: %s - IPv6: %s\n", privnet.ID, privnet.Name, privnet.IPv6Aggregate)
}

func ExamplePrivateNetworkService_EstimatedCost() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	netCost, _ := client.PrivateNetworks.EstimatedCost(context.Background(), "pn-123ab")

	fmt.Printf("Price: %.2f%s - Discount: %.2f - Total: %.2f\n", netCost.Price, netCost.Currency, netCost.Discount, netCost.Total)
}

func ExamplePrivateNetworkService_Edit() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.EditPrivateNetworkParams{
		ID:   "pn-123ab",
		Name: "new-network-name",
	}

	privnet, _ := client.PrivateNetworks.Edit(context.Background(), params)

	fmt.Printf("ID: %s - Name: %s - IPv6: %s\n", privnet.ID, privnet.Name, privnet.IPv6Aggregate)
}

func ExamplePrivateNetworkService_List() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	privnet, _ := client.PrivateNetworks.List(context.Background())

	for _, net := range *privnet {
		fmt.Printf("ID: %s - Name: %s - IPv6: %s\n", net.ID, net.Name, net.IPv6Aggregate)
	}
}

func ExamplePrivateNetworkService_CreateSegment() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.CreatePrivateNetworkSegmentParams{
		PrivateNetworkID: "pn-123ab",
		Datacenter:       "dc-fbg1",
		Name:             "kvm-segment",
		Platform:         "kvm",
		IPv4Subnet:       "192.0.2.0/24",
	}

	segment, _ := client.PrivateNetworks.CreateSegment(context.Background(), params)

	fmt.Printf("ID: %s - Name: %s\nIPv6: %s\nIPv4: %s\nPlatform: %s\nDatacenter: %s\n",
		segment.ID, segment.Name, segment.IPv6Subnet, segment.IPv4Subnet, segment.Platform, segment.Datacenter)
}

func ExamplePrivateNetworkService_EditSegment() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.EditPrivateNetworkSegmentParams{
		ID:   "6f5cb761-163e-4a5f-b9da-98acbe68e28d",
		Name: "updated-name",
	}

	segment, _ := client.PrivateNetworks.EditSegment(context.Background(), params)

	fmt.Printf("ID: %s - Name: %s\nIPv6: %s\nIPv4: %s\nPlatform: %s\nDatacenter: %s\n",
		segment.ID, segment.Name, segment.IPv6Subnet, segment.IPv4Subnet, segment.Platform, segment.Datacenter)
}

func ExamplePrivateNetworkService_ListSegments() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	segments, _ := client.PrivateNetworks.ListSegments(context.Background(), "pn-123ab")

	for _, s := range *segments {
		fmt.Printf("ID: %s - Name: %s\nIPv6: %s\nIPv4: %s\nPlatform: %s\nDatacenter: %s\n",
			s.ID, s.Name, s.IPv6Subnet, s.IPv4Subnet, s.Platform, s.Datacenter)
	}
}

func ExamplePrivateNetworkService_DestroySegment() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	err := client.PrivateNetworks.DestroySegment(context.Background(), "6f5cb761-163e-4a5f-b9da-98acbe68e28d")

	if err != nil {
		fmt.Printf("Cannot remove private network segment. Error: %s", err)
	}
}

func ExampleServerService_Create() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	server, _ := client.Servers.Create(context.Background(), glesys.CreateServerParams{
		Bandwidth:    100,
		CampaignCode: "",
		CPU:          2,
		DataCenter:   "Falkenberg",
		Description:  "",
		Hostname:     "my-hostname",
		IPv4:         "any",
		IPv6:         "any",
		Memory:       2048,
		Password:     "...",
		Platform:     "VMware",
		PublicKey:    "...",
		Storage:      50,
		Template:     "Debian 8 64-bit",
	})

	fmt.Println(server.ID)

	// NOTE: You can also use the WithDefaults() to provide defaults values for
	// all required parameters except Password (or PublicKey)

	server2, _ := client.Servers.Create(context.Background(), glesys.CreateServerParams{
		Password: "...",
	}.WithDefaults())

	fmt.Println(server2.ID)
}

func ExampleServerService_Console() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	console, _ := client.Servers.Console(context.Background(), "kvm12345")

	fmt.Println(console.URL)
}

func ExampleServerService_Destroy() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.Servers.Destroy(context.Background(), "kvm12345", glesys.DestroyServerParams{
		KeepIP: true, // KeepIP defaults to false
	})
}

func ExampleServerService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	server, _ := client.Servers.Details(context.Background(), "kvm12345")

	fmt.Println(server.Hostname)
}

func ExampleServerService_Edit() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	server, _ := client.Servers.Edit(context.Background(), "kvm12345", glesys.EditServerParams{
		Bandwidth:   100,
		CPU:         4,
		Description: "Web Server",
		Hostname:    "example.com",
		Memory:      4096,
		Storage:     250,
	})

	fmt.Println(server.ID)
}

func ExampleServerService_List() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	servers, _ := client.Servers.List(context.Background())

	for _, server := range *servers {
		fmt.Println(server.Hostname)
	}
}

func ExampleServerService_ListISOs() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	isos, _ := client.Servers.ListISOs(context.Background(), "wps123456")

	for _, iso := range *isos {
		fmt.Println(iso)
	}
}

func ExampleServerService_MountISO() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	serverDetail, _ := client.Servers.MountISO(context.Background(), "wps123456", "OpenBSD/7.4/amd64/install74.iso")

	fmt.Printf("ISO Mounted: %s\n", serverDetail.ISOFile)
}

func ExampleServerService_NetworkAdapters() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	na, _ := client.Servers.NetworkAdapters(context.Background(), "wps123456")

	for _, n := range *na {
		fmt.Println(n.Name)
		fmt.Println(n.ID)
	}
}

func ExampleServerService_Start() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.Servers.Start(context.Background(), "kvm12345")
}

func ExampleServerService_Stop() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.Servers.Stop(context.Background(), "kvm12345", glesys.StopServerParams{
		Type: "reboot", // Type "soft", "hard" and "reboot" available
	})
}

func ExampleServerService_Templates() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	templates, _ := client.Servers.Templates(context.Background())

	for _, template := range templates.KVM {
		fmt.Println(template.Name)
	}
}

func ExampleServerService_PreviewCloudConfig() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	cloudconfig := "## template: glesys\n#cloud-config\n{{>users}}"
	cloudConfigParams := map[string]any{"foo": "bar", "balloon": 99}
	users := []glesys.User{}
	users = append(users, glesys.User{
		Username:   "bob",
		Password:   "hunter2!",
		PublicKeys: []string{"ssh-ed25519 AAAAC3NKEY bob@bob-machine"},
	})

	preview, _ := client.Servers.PreviewCloudConfig(context.Background(), glesys.PreviewCloudConfigParams{
		CloudConfig:       cloudconfig,
		CloudConfigParams: cloudConfigParams,
		Users:             users,
	})

	fmt.Println(preview.Context.Users[0].Username)
	fmt.Printf("Number of balloons: %f\n", preview.Context.Params["balloon"])
}

func ExampleServerDisksService_Create() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.CreateServerDiskParams{
		ServerID:  "wps123456",
		SizeInGIB: 125,
		Name:      "Extra disk",
		Type:      "silver", // [gold|silver]
	}

	_, _ = client.ServerDisks.Create(context.Background(), params)

	server, _ := client.Servers.Details(context.Background(), "wps123456")

	for _, disk := range server.AdditionalDisks {
		fmt.Printf("Disk%d: %s - %d - Type: %s\n", disk.SCSIID, disk.Name, disk.SizeInGIB, disk.Type)
	}
}

func ExampleServerDisksService_Delete() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	diskid := "abc1234-zxy987-cccbbbaaa"
	err := client.ServerDisks.Delete(context.Background(), diskid)

	if err != nil {
		fmt.Println("Error deleting disk", err)
	}
}

func ExampleServerDisksService_UpdateName() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.EditServerDiskParams{
		ID:   "abc1234-zxy987-cccbbbaaa",
		Name: "new name",
	}

	_, err := client.ServerDisks.UpdateName(context.Background(), params)

	if err != nil {
		fmt.Println("Error renaming disk", err)
	}
}

func ExampleServerDisksService_Reconfigure() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.EditServerDiskParams{
		ID:        "abc1234-zxy987-cccbbbaaa",
		SizeInGIB: 120,
	}

	_, err := client.ServerDisks.Reconfigure(context.Background(), params)

	if err != nil {
		fmt.Println("Error updating disk size", err)
	}
}

func ExampleServerDisksService_Limits() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	serverid := "wps123456"
	limits, _ := client.ServerDisks.Limits(context.Background(), serverid)

	fmt.Printf("Limits (%s): Min size: %d, Max size %d, Current number of disks: %d, Max number of disks: %d\n",
		serverid,
		limits.MinSizeInGIB,
		limits.MaxSizeInGIB,
		limits.CurrentNumDisks,
		limits.MaxNumDisks,
	)
}

func ExampleObjectStorageService_CreateInstance() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.CreateObjectStorageInstanceParams{
		DataCenter:  "dc-sto1",
		Description: "My ObjectStorage",
	}

	instance, _ := client.ObjectStorages.CreateInstance(context.Background(), params)

	fmt.Println(instance.InstanceID)
}

func ExampleObjectStorageService_InstanceDetails() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	instance, _ := client.ObjectStorages.InstanceDetails(context.Background(), "os-ab123")

	fmt.Println(instance.InstanceID)
}

func ExampleObjectStorageService_DeleteInstance() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	err := client.ObjectStorages.DeleteInstance(context.Background(), "os-ab123")

	if err != nil {
		fmt.Printf("Error removing objectstorage instance: %s", err)
	}
}

func ExampleObjectStorageService_ListInstances() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	instances, _ := client.ObjectStorages.ListInstances(context.Background())

	for _, instance := range *instances {
		fmt.Println(instance.InstanceID)
	}
}

func ExampleObjectStorageService_EditInstance() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.EditObjectStorageInstanceParams{
		InstanceID:  "os-ab123",
		Description: "My ObjectStorage New",
	}

	instance, _ := client.ObjectStorages.EditInstance(context.Background(), params)

	fmt.Println(instance.Description)
}

func ExampleObjectStorageService_CreateCredential() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.CreateObjectStorageCredentialParams{
		InstanceID:  "os-ab123",
		Description: "New Key 1",
	}

	credential, _ := client.ObjectStorages.CreateCredential(context.Background(), params)

	fmt.Println(credential.AccessKey)
}

func ExampleObjectStorageService_DeleteCredential() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	params := glesys.DeleteObjectStorageCredentialParams{
		InstanceID:   "os-ab123",
		CredentialID: "16df46b3-b2f0-471b-81bf-56c26fff7c4d",
	}

	err := client.ObjectStorages.DeleteCredential(context.Background(), params)

	if err != nil {
		fmt.Printf("Error removing objectstorage credential: %s", err)
	}
}

func ExampleDatabaseService_Create() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	details, _ := client.Databases.Create(context.Background(), glesys.CreateDatabaseParams{
		PlanKey:       "plan-1core-4gib-25gib",
		Engine:        "mysql",
		EngineVersion: "8.0",
		DataCenterKey: "",
		Name:          "My Database",
	})
	fmt.Printf("%#v\n", details)
}

func ExampleDatabaseService_Delete() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	err := client.Databases.Delete(context.Background(), "db-12345")

	if err != nil {
		fmt.Printf("Error removing database: %s", err)
	}
}

func ExampleDatabaseService_UpdateAllowlist() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	plans, _ := client.Databases.UpdateAllowlist(context.Background(), glesys.UpdateAllowlistParams{
		ID:        "db-12345",
		AllowList: []string{"127.0.0.1", "127.0.0.2"},
	})

	fmt.Printf("%#v\n", plans)
}

func ExampleDatabaseService_List() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	list, _ := client.Databases.List(context.Background())

	fmt.Printf("%#v\n", list)
}

func ExampleDatabaseService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	plans, _ := client.Databases.UpdateAllowlist(context.Background(), glesys.UpdateAllowlistParams{
		ID:        "db-12345",
		AllowList: []string{"127.0.0.1", "127.0.0.2"},
	})

	fmt.Printf("%#v\n", plans)
}

func ExampleDatabaseService_ConnectionString() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	connectionString, _ := client.Databases.ConnectionString(context.Background(), "db-12345")

	fmt.Printf("%#v\n", connectionString)
}

func ExampleDatabaseService_ListPlans() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	plans, _ := client.Databases.ListPlans(context.Background())

	fmt.Printf("%#v\n", plans)
}

func ExampleDatabaseService_EstimatedCost() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	billing, _ := client.Databases.EstimatedCost(context.Background(), glesys.EstimatedCostParams{
		ID:      "db-12346",
		PlanKey: "plan-1core-4gib-25gib",
	})

	fmt.Printf("%#v\n", billing)
}
