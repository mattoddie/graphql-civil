# Civil

This package wraps google cloud civil package adding GraphQL marshal/unmarshal methods for [gqlgen](https://gqlgen.com/).

## Usage

Add the following to gqlgen.yml

```yaml
models:
  Date:
    model: github.com/mattoddie/graphql-civil.CivilDate
```
