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

The parsers have fuzz tests, which also run weekly in CI. To fuzz one
locally:

```sh
go test -run '^$' -fuzz '^FuzzCivilDate$' -fuzztime 1m .
```

If fuzzing finds a failure, the failing input is saved under `testdata/fuzz`;
commit it with the fix so it becomes a regression test.

## Pull requests

1. Fork the repository and create a branch from `main`.
2. Keep changes focused: one fix or feature per pull request.
3. Add or update tests for any behaviour change, and update the README and
   doc comments when the public API changes.
4. Write commit messages and the pull request title using
   [Conventional Commits](#commit-messages).
5. Open the pull request and fill in the template. CI must pass before it can
   be merged.

### Commit messages

Commit messages and pull request titles follow
[Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/).
Pull requests are squash-merged, so the title becomes the commit on `main`;
CI checks that it matches.

```text
<type>(<optional scope>)<optional !>: <description>
```

For example `fix(time): reject fractional seconds beyond nanoseconds`,
`feat: add MarshalCivilDateTime`, `docs: explain the gqlgen scalar mapping`
or `build(deps): bump github.com/99designs/gqlgen to v0.17.96`.

- **Types:** `feat`, `fix`, `docs`, `refactor`, `perf`, `test`, `build`
  (dependencies and toolchain), `ci` (workflows), `chore` (anything else),
  `revert` and `style`.
- **Description:** imperative mood, lower-case first word, no trailing full
  stop. Keep the whole subject to 72 characters or fewer.
- **Breaking changes:** add `!` after the type or scope, for example
  `feat(date)!: …`, and a `BREAKING CHANGE: …` footer in the body (or the pull
  request description).
- **Body:** explain why when it isn't obvious, and reference issues with
  `Fixes #123`.

## Releases

Releases follow [semantic versioning](https://semver.org/) and are cut by the
maintainer by tagging `main`.

## Code of conduct

This project follows the [Code of Conduct](CODE_OF_CONDUCT.md). By
participating you agree to uphold it.
