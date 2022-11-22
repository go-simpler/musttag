# musttag

[![ci](https://github.com/junk1tm/musttag/actions/workflows/go.yml/badge.svg)](https://github.com/junk1tm/musttag/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/github.com/junk1tm/musttag.svg)](https://pkg.go.dev/github.com/junk1tm/musttag)
[![report](https://goreportcard.com/badge/github.com/junk1tm/musttag)](https://goreportcard.com/report/github.com/junk1tm/musttag)
[![codecov](https://codecov.io/gh/junk1tm/musttag/branch/main/graph/badge.svg)](https://codecov.io/gh/junk1tm/musttag)

A Go linter that enforces field tags in (un)marshaled structs

## 📌 About

`musttag` checks if struct fields used in `Marshal`/`Unmarshal` are annotated with the relevant tag:

```go
// BAD:
var user struct {
	Name string
}
data, err := json.Marshal(user)

// GOOD:
var user struct {
	Name string `json:"name"`
}
data, err := json.Marshal(user)
```

The rational from [Uber Style Guide][1]:

> The serialized form of the structure is a contract between different systems.
> Changes to the structure of the serialized form--including field names--break this contract.
> Specifying field names inside tags makes the contract explicit,
> and it guards against accidentally breaking the contract by refactoring or renaming fields.

Currently `musttag` supports only the `json` tag, see [Roadmap](#-roadmap) for future plans.

## 📦 Install

```shell
go install github.com/junk1tm/musttag/cmd/musttag
```

## 📋 Usage

With `go vet`:

```shell
go vet -vettool=$(which musttag) ./...
```

## 📅 Roadmap

* Support `xml`
* Support `yaml`
* Support `toml`
* Support `mapstructure`
* Support custom tags via config

[1]: https://github.com/uber-go/guide/blob/master/style.md#use-field-tags-in-marshaled-structs
