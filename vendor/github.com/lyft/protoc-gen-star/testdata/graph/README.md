# AST Graph Test Data

This directory contains various test proto file sets for black-box testing of the AST gatherer `graph`.

Proto files are preprocessed to their descriptors, imported directly into the `ast_test.go` tests, and unmarshaled as a  `DescriptorFileSet`.

## To Generate

From the project root:

```sh
make testdata-graph
```
