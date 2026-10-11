# Contributing

Thanks for your interest in improving graphql-civil! Bug reports, fixes and
small improvements are all welcome.

## Reporting bugs and requesting features

Open an [issue](https://github.com/mattoddie/graphql-civil/issues/new/choose)
using one of the templates. For bugs, include the Go and gqlgen versions you
are using and the input that triggers the problem.

Please **do not** report security vulnerabilities in public issues; see
[SECURITY.md](SECURITY.md) instead.

## Development

You need Go 1.26 or later.

```sh
git clone https://github.com/mattoddie/graphql-civil.git
cd graphql-civil
go test ./...
```

Before opening a pull request, make sure the same checks that CI runs pass:

```sh
go mod tidy -diff
go vet ./...
go test -race ./...
golangci-lint run          # https://golangci-lint.run, v2
govulncheck ./...          # go install golang.org/x/vuln/cmd/govulncheck@latest
```

`golangci-lint fmt` applies the configured formatters (gofumpt and goimports).

## Pull requests

1. Fork the repository and create a branch from `main`.
2. Keep changes focused: one fix or feature per pull request.
3. Add or update tests for any behaviour change, and update the README and
   doc comments when the public API changes.
4. Write commit messages with a short imperative subject line (at most 72
   characters) and a body explaining why the change is needed.
5. Open the pull request and fill in the template. CI must pass before it can
   be merged.

## Releases

Releases follow [semantic versioning](https://semver.org/) and are cut by the
maintainer by tagging `main`.

## Code of conduct

This project follows the [Code of Conduct](CODE_OF_CONDUCT.md). By
participating you agree to uphold it.
