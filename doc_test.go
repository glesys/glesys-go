package glesys_test

import (
	"context"
	"fmt"

	glesys "github.com/glesys/glesys-go"
)

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
