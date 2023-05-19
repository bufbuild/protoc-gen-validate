# [![The Buf logo](../../.github/buf-logo.svg)][buf] protovalidate

[![Report Card](https://goreportcard.com/badge/github.com/bufbuild/protovalidate)](https://goreportcard.com/report/github.com/bufbuild/protovalidate)
[![GoDoc](https://pkg.go.dev/badge/github.com/bufbuild/protovalidate.svg)](https://pkg.go.dev/github.com/bufbuild/protovalidate)

`protovalidate` is a Go library designed to validate Protobuf messages at
runtime based on user-defined validation constraints. Powered by Google's Common
Expression Language ([CEL](https://github.com/google/cel-spec)), it provides a
flexible and efficient foundation for defining and evaluating custom validation
rules. The primary goal of `protovalidate` is to help developers ensure data
consistency and integrity across the network without requiring generated code.

## Installation

**Requires the `go` toolchain (≥ v1.18)**

To install the package, use the `go get` command from within your Go module:

```shell
go get github.com/bufbuild/protovalidate/go/v2
```

Import the package into your Go project:

```go
import "github.com/bufbuild/protovalidate/go/v2"
```

Remember to always check for the latest version of `protovalidate` on the
project's [GitHub releases page](https://github.com/bufbuild/protovalidate/releases)
to ensure you're using the most up-to-date version.

## Usage

> For API-specific details, you can refer to the
official [pkg.go.dev](https://pkg.go.dev/github.com/bufbuild/protovalidate/go)
documentation which provides in-depth information on the library's API,
functions, types, and source files. This can be particularly useful for
understanding the lower-level workings of `protovalidate` and how to leverage
its full potential.

### Example

```go
package main

import (
  "fmt"
  pb "github.com/path/to/generated/protos"
  "github.com/bufbuild/protovalidate/go/v2"
)

func main() {
  msg := &pb.Person{
    Id:    1000, 
    Email: "example@bufbuild.com", 
    Name:  "Protobuf",
    Home: &example.Person_Location{
      Lat: 37.7, 
      Lng: -122.4,
    },
  }

  v, err := protovalidate.New()
  if err != nil {
    fmt.Println("failed to initialize validator:", err)
  }

  if err = v.Validate(msg); err != nil {
    fmt.Println("validation failed:", err)
  } else {
    fmt.Println("validation succeeded")
  }
}
```

### Implementing validation constraints

Validation constraints are defined directly within `.proto` files.
Documentation for adding constraints can be found in the root [README](../../README.md) and the [comprehensive docs](../../docs).

The `protovalidate` package assumes the constraint extensions are imported into the `protoc-gen-go` generated code via 
`buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go`. 

#### Buf managed mode

If you are using Buf [managed mode](https://buf.build/docs/generate/managed-mode/) to augment Go code generation, ensure that the `protovalidate` module is excluded in your [`buf.gen.yaml`](https://buf.build/docs/configuration/v1/buf-gen-yaml#except):

```yaml
version: v1
# <snip>
managed:
  enabled: true
  go_package_prefix:
    except:
      - buf.build/bufbuild/protovalidate
# <snip>
```

### Lazy mode

`protovalidate` defaults to lazily construct validation logic for Protobuf 
message types the first time they are encountered. A validator's internal 
cache can be pre-warmed with the `WithMessages` or `WithDescriptors` options 
during initialization:

```go
validator, err := protovalidate.New(
  protovalidate.WithMessages(
    &pb.MyFoo{}, 
    &pb.MyBar{}, 
  ),
)
```

Lazy mode requires usage of a mutex to keep the validator thread-safe, which 
results in about 50% of CPU time spent obtaining a read lock. While [performance](#performance)
is sub-microsecond, the mutex overhead can be further reduced by disabling lazy 
mode with the `WithDisableLazy` option. Note that all expected messages must be
provided during initialization of the validator:

```go
validator, err := protovalidate.New(
  protovalidate.WithDisableLazy(true),
  protovalidate.WithMessages(
    &pb.MyFoo{},
    &pb.MyBar{},
  ),
)
```

### Support legacy `protoc-gen-validate` constraints

The protovalidate module comes with a `legacy` package which adds opt-in support
for existing `protoc-gen-validate` constraints. Provide the`legacy.WithLegacySupport` 
option when initializing the validator:

```go
validator, err := protovalidate.New(
  legacy.WithLegacySupport(legacy.ModeMerge),
)
```

`protoc-gen-validate` code generation is **not** used by `protovalidate`. The 
`legacy` package assumes the `protoc-gen-validate` extensions are imported into
the `protoc-gen-go` generated code via `github.com/envoyproxy/protoc-gen-validate/validate`.

A [migration tool](../../tools/migrate/README.md) is also available to incrementally upgrade legacy constraints 
in `.proto` files.

## Performance

[Benchmarks](validator_bench_test.go) are provided to test a variety of use-cases. Generally, after the 
initial cold start, validation on a message is sub-microsecond 
and only allocates in the event of a validation error.

```
[circa 15 May 2023]
goos: darwin
goarch: arm64
pkg: github.com/bufbuild/protovalidate/go/v2
BenchmarkValidator
BenchmarkValidator/ColdStart
BenchmarkValidator/ColdStart-10         	    4372	    276457 ns/op	  470780 B/op	    9255 allocs/op
BenchmarkValidator/Lazy/Valid
BenchmarkValidator/Lazy/Valid-10        	 9022392	     134.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidator/Lazy/Invalid
BenchmarkValidator/Lazy/Invalid-10      	 3416996	     355.9 ns/op	     632 B/op	      14 allocs/op
BenchmarkValidator/Lazy/FailFast
BenchmarkValidator/Lazy/FailFast-10     	 6751131	     172.6 ns/op	     168 B/op	       3 allocs/op
BenchmarkValidator/PreWarmed/Valid
BenchmarkValidator/PreWarmed/Valid-10   	17557560	     69.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidator/PreWarmed/Invalid
BenchmarkValidator/PreWarmed/Invalid-10 	 3621961	     332.9 ns/op	     632 B/op	      14 allocs/op
BenchmarkValidator/PreWarmed/FailFast
BenchmarkValidator/PreWarmed/FailFast-10	13960359	     92.22 ns/op	     168 B/op	       3 allocs/op
PASS
```

[buf]: https://buf.build
