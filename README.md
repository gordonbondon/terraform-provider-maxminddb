# terraform-provider-maxminddb

Terraform provider that helps working with [MaxMind GeoIP2 Databases](https://www.maxmind.com/en/geoip2-databases)

## Using the Provider

For installation instructions and resource documentation check [provider page on terraform registry](https://registry.terraform.io/providers/gordonbondon/maxminddb/latest).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) below).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

### Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.14.x
-	[Go](https://golang.org/doc/install) >= 1.15

### Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `build` command:
```sh
$ go build .
```

To check locally build provider, update [`dev_overrides` cli config](https://www.terraform.io/docs/commands/cli-config.html#development-overrides-for-provider-developers):

```hcl
provider_installation {
  dev_overrides {
    "gordonbondon/maxminddb" = "/path/to/repo/terraform-provider-maxminddb"
  }

  direct {}
}
```

### Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Developing the Provider

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests may create real resources, and often cost money to run.

```sh
$ make testacc
```

### Releasing the Provider

```shell
$ fingerprint=$(gpg --with-colons --list-key <key@email> | awk -F: '$1 == "fpr" {print $10;}' | head -n 1)
$ export GPG_FINGERPRINT="${fingerprint}"
$ export GITHUB_TOKEN=xxx
$ git tag v0.x.x -s
$ goreleaser release --rm-dist
```
