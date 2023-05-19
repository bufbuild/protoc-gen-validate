# Go conformance executor

This binary is the [conformance testing executor](../../../../../tools/protovalidate-conformance/README.md) for the Go implementation. From the root of the project, the Go conformance tests can be executed with make:

```shell
make conformance-go # runs all conformance tests

make conformance-go ARGS='-suite uint64' # pass flags to the conformance harness
```
