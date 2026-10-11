# graphql-civil

[![Test](https://github.com/mattoddie/graphql-civil/actions/workflows/test.yml/badge.svg)](https://github.com/mattoddie/graphql-civil/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mattoddie/graphql-civil.svg)](https://pkg.go.dev/github.com/mattoddie/graphql-civil)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattoddie/graphql-civil)](https://goreportcard.com/report/github.com/mattoddie/graphql-civil)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/mattoddie/graphql-civil/badge)](https://scorecard.dev/viewer/?uri=github.com/mattoddie/graphql-civil)

This package wraps the Google Cloud [civil](https://pkg.go.dev/cloud.google.com/go/civil)
package, adding GraphQL marshal/unmarshal functions for [gqlgen](https://gqlgen.com/).
Use it to expose timezone-free dates, times and datetimes as GraphQL scalars.

Requires Go 1.26 or later.

## Installation

```sh
go get github.com/mattoddie/graphql-civil
```

## Usage

Declare the scalars you need in your schema:

```graphql
scalar Date
scalar DateTime
scalar Time
```

Then map them to this package in gqlgen.yml:

```yaml
models:
  Date:
    model: github.com/mattoddie/graphql-civil.CivilDate
  DateTime:
    model: github.com/mattoddie/graphql-civil.CivilDateTime
  Time:
    model: github.com/mattoddie/graphql-civil.CivilTime
```

Resolvers then use `civil.Date`, `civil.DateTime` and `civil.Time` from
`cloud.google.com/go/civil`.

| Scalar     | Go type          | Example value                     |
| ---------- | ---------------- | --------------------------------- |
| `Date`     | `civil.Date`     | `"2020-11-06"`                    |
| `DateTime` | `civil.DateTime` | `"2020-11-06T12:30:00"`           |
| `Time`     | `civil.Time`     | `"12:30:00"`, `"12:30:00.5"`      |

Times and datetimes accept optional fractional seconds on input, and are
written with nine-digit fractional seconds when they are non-zero.

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md), and report
security issues as described in [SECURITY.md](SECURITY.md).

## License

[MIT](LICENSE)
