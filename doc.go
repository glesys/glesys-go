// Package glesys is the official Go client for interacting with the GleSYS API.
//
// Please note that only a subset of features available in the GleSYS API has
// been implemented. We greatly appreciate contributions.
//
// Getting Started
//
// To get started you need to signup for an account to GleSYS Cloud and create
// an API key. Signup is available at https://glesys.com/signup and API keys can
// be created at https://customer.glesys.com.
//
//     client := glesys.NewClient("CL12345", "your-api-key")
//
// CL12345 is the key of the Project you want to work with.
//
// The different modules of the GleSYS API are available on the client.
// For example:
//
//     client.IPs.List(...)
//     client.Servers.Create(...)
//
// More examples provided below.
package glesys
