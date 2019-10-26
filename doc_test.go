package glesys_test

import (
	"context"
	"fmt"

	glesys "github.com/glesys/glesys-go"
)

func ExampleEmailService_Overview() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: All parameters in EmailOverviewParams are optional and can be omitted.
	overview, _ := client.Emails.Overview(context.Background(), glesys.EmailOverviewParams{
		Filter: "your-domain.com",
		Page:   1,
	})

	fmt.Printf("%#v", overview)
}

func ExampleEmailService_GlobalQuota() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	// NOTE: The GlobalQuota parameter can be omitted to only fetch the current value.
	globalquota, _ := client.Emails.GlobalQuota(context.Background(), glesys.EmailGlobalQuotaParams{
		GlobalQuota: 20480,
	})

	fmt.Println(globalquota.Usage)
	fmt.Println(globalquota.Max)
}

func ExampleIPService_Available() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	ips, _ := client.IPs.Available(context.Background(), glesys.AvailableIPsParams{
		DataCenter: "Falkenberg",
		Platform:   "OpenVZ",
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

	ips, _ := client.IPs.Reserved(context.Background())

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

func ExampleLoadBalancer_AddBackend() {
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

func ExampleLoadBalancer_AddTarget() {
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

func ExampleLoadBalancer_DisableTarget() {
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

func ExampleLoadBalancer_AddFrontend() {
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

func ExampleLoadBalancer_AddCertificate() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	mybase64pem := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCm15Y2VydAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLQpteWtleQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg=="

	client.LoadBalancers.AddCertificate(context.Background(), "lb123456", glesys.AddCertificateParams{
		Name:        "mycert",
		Certificate: mybase64pem,
	})
}

func ExampleLoadBalancer_ListCertificate() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	certlist, _ := client.LoadBalancers.ListCertificates(context.Background(), "lb123456")

	for _, cert := range *certlist {
		fmt.Println("Certificate:", cert)
	}
}

func ExampleLoadBalancer_RemoveCertificate() {
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
		Platform:     "OpenVZ",
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

func ExampleServerService_Destroy() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.Servers.Destroy(context.Background(), "vz12345", glesys.DestroyServerParams{
		KeepIP: true, // KeepIP defaults to false
	})
}

func ExampleServerService_Details() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	server, _ := client.Servers.Details(context.Background(), "vz12345")

	fmt.Println(server.Hostname)
}

func ExampleServerService_Edit() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	server, _ := client.Servers.Edit(context.Background(), "vz12345", glesys.EditServerParams{
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

func ExampleServerService_Start() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.Servers.Start(context.Background(), "vz12345")
}

func ExampleServerService_Stop() {
	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")

	client.Servers.Stop(context.Background(), "vz12345", glesys.StopServerParams{
		Type: "reboot", // Type "soft", "hard" and "reboot" available
	})
}
