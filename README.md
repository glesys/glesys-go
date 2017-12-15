[![Build Status](https://travis-ci.org/glesys/glesys-go.svg?branch=master)](https://travis-ci.org/glesys/glesys-go)

# glesys-go

This is the official client library for interacting with the
[GleSYS API](https://github.com/GleSYS/API/).

## Requirements

- Go 1.7 or higher (relies on [context](https://golang.org/pkg/context/))

## Getting Started

#### Installation

```shell
go get github.com/glesys/glesys-go
```

#### Authentication

To use the glesys-go library you need a GleSYS Cloud account and a valid API
key. You can sign up for an account at https://glesys.com/signup. After signing
up visit https://customer.glesys.com to create an API key for your Project.

#### Set up a Client

```go
client := glesys.NewClient("CL12345", "your-api-key", "my-application/0.0.1")
```

#### Create a Server

```go
// Create a Server
server, err := client.Servers.Create(context.Background(), glesys.CreateServerParams{Password: "..."}.WithDefaults())
```

#### List all Servers

```go
// List all Servers
servers, err := client.Servers.List(context.Background())
```

#### User agent

To be able to monitor usage and help track down issues, we encourage you to
provide a user agent string identifying your application or library. Recommended
syntax is `my-library/version` or `www.example.com`.

#### Context

glesys-go uses Go's [context](https://golang.org/pkg/context) library to handle
timeouts and deadlines. All functions making HTTP requests requires a `context`
argument.

### Documentation

Full documentation is available at
https://godoc.org/github.com/glesys/glesys-go.

## Contribute

#### We love Pull Requests ♥

1. Fork the repo.
2. Make sure to run the tests to verify that you're starting with a clean slate.
3. Add a test for your change, make sure it fails. Refactoring existing code or
   improving documentation does not require new tests.
4. Make the changes and ensure the test pass.
5. Commit your changes, push to your fork and submit a Pull Request.

#### Syntax

Please use the formatting provided by [gofmt](https://golang.org/cmd/gofmt).

## License

The contents of this repository are distributed under the MIT license, see [LICENSE](LICENSE).
