# Vaku

[![Vaku](www/assets/logo-vaku-sm.png?raw=true)](www/assets/logo-vaku-sm.png "Vaku")

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lingrino/vaku/vaku)](https://pkg.go.dev/github.com/lingrino/vaku/vaku)
[![goreportcard](https://goreportcard.com/badge/github.com/lingrino/vaku)](https://goreportcard.com/report/github.com/lingrino/vaku)
[![Maintainability](https://api.codeclimate.com/v1/badges/db6951b0aa53becf8c92/maintainability)](https://codeclimate.com/github/lingrino/vaku/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/db6951b0aa53becf8c92/test_coverage)](https://codeclimate.com/github/lingrino/vaku/test_coverage)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/65802905eb8148e2ae9ae4c909673ee2)](https://www.codacy.com/manual/lingrino/vaku)
[![Codebeat badge](https://codebeat.co/badges/f6dfd08e-97c5-4afd-9dd0-64cf0a5d03a8)](https://codebeat.co/projects/github-com-lingrino-vaku-main)

A CLI and API for running path and folder based operations on Vault k/v engines.

## Installation

### Homebrew

```shell
brew install lingrino/tap/vaku
```

### Scoop

```shell
scoop bucket add vaku https://github.com/lingrino/scoop-vaku.git
scoop install vaku
```

### Docker

```shell
docker run lingrino/vaku --help
```

### Binary

Download the latest binary or deb/rpm for your os/arch from the [releases page](https://github.com/lingrino/vaku/releases).

## Usage

Usage docs here.

## API

Documentation for the Vaku API is on [pkg.go.dev](https://pkg.go.dev/github.com/lingrino/vaku/vaku).

## Contributing

Suggestions and [contributions](https://github.com/lingrino/.github/blob/HEAD/CONTRIBUTING.md) of all kinds are welcome! If there is functionality you would like to see in Vaku please open an issue or pull request and I will be sure to address it.

## Tests

Vaku is well tested and uses only the standard go testing tools.

```shell
$ go test -cover -race ./...
ok  github.com/lingrino/vaku/vaku (cached) coverage: 100.0% of statements
```
