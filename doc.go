// Package glesys is the official Go client for interacting with the GleSYS API.
//
// Please note that only a subset of features available in the GleSYS API has
// been implemented. We greatly appreciate contributions.
//
// # Getting Started
//
// To get started you need to signup for a GleSYS Cloud account and create an
// API key. Signup is available at https://glesys.com/signup and API keys can be
// created at https://cloud.glesys.com.
//
//	client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")
//
// CL12345 is the key of the Project you want to work with.
//
// To be able to monitor usage and help track down issues, we encourage you to
// provide a user agent string identifying your application or library.
// Recommended syntax is "my-library/version" or "www.example.com".
//
// The different modules of the GleSYS API are available on the client.
// For example:
//
//	client.EmailDomains.Overview(...)
//	client.IPs.List(...)
//	client.NetworkAdapters.Create(...)
//	client.Networks.Create(...)
//	client.Servers.Create(...)
//
// More examples provided below.
package glesys
