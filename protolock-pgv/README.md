# Protolock-PGV

Protolock-PGV is a backwards compatibility verifier for protoc-gen-validate implemented as a
[Protolock](https://github.com/nilslice/protolock) plugin.

Protolock ensures that protos remain backwards compatable as they change and evolve. Protolock-PGV extends Protolock's
change detection so it understands PGV validation annotations. Working together, Protolock and Protolock-PGV ensure
your protos remain backward compatible and valid throughtout their evolution.

## Installation

Installing PGV can currently only be done from source:

```shell
# fetches this repo into $GOPATH
go get -d github.com/envoyproxy/protoc-gen-validate

# installs Protolock-PGV into $GOPATH/bin
make build
```

## Usage

Protolock-PGV requires [Protolock](https://github.com/nilslice/protolock) to also be installed.

```shell
> protolock status --plugins=protolock-pgv
```

## Rules Enforced

Protolock-PGV will alert you with a Protolock warning if any of the following changes are detected:

* A new PGV validation rule is added
* A PGV rule is changed
