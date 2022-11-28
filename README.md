# musttag

[![ci](https://github.com/junk1tm/musttag/actions/workflows/go.yml/badge.svg)](https://github.com/junk1tm/musttag/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/github.com/junk1tm/musttag.svg)](https://pkg.go.dev/github.com/junk1tm/musttag)
[![report](https://goreportcard.com/badge/github.com/junk1tm/musttag)](https://goreportcard.com/report/github.com/junk1tm/musttag)
[![codecov](https://codecov.io/gh/junk1tm/musttag/branch/main/graph/badge.svg)](https://codecov.io/gh/junk1tm/musttag)

A Go linter that enforces field tags in (un)marshaled structs

## 📌 About

`musttag` checks that exported fields of a struct passed to a `Marshal`-like function are annotated with the relevant tag:

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
> Changes to the structure of the serialized form, including field names, break this contract.
> Specifying field names inside tags makes the contract explicit,
> and it guards against accidentally breaking the contract by refactoring or renaming fields.

## 🚀 Features

`musttag` supports these packages out of the box:

* `encoding/json`
* `encoding/xml`
* `gopkg.in/yaml.v3`
* `github.com/BurntSushi/toml`
* `github.com/mitchellh/mapstructure`
* ...and any custom one

## 📦 Install

```shell
go install github.com/junk1tm/musttag/cmd/musttag@latest
```

## 📋 Usage

```shell
musttag ./...
```

With `go vet`:

```shell
go vet -vettool=$(which musttag) ./...
```

[1]: https://github.com/uber-go/guide/blob/master/style.md#use-field-tags-in-marshaled-structs
