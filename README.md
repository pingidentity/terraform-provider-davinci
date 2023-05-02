# Terraform Provider Davinci (Terraform Plugin SDK)

This repository is for the PingOne DaVinci [Terraform](https://www.terraform.io) provider.

 - Resources and Datasources are found in: `internal/service/davinci`

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.19

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Find readable documentation on registry.terraform.io

## Developing the Provider

See [development-environment](./contributing/development-environment.md)