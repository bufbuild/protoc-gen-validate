# protoc-gen-debug

This plugin can be used to create test files containing the entire encoded CodeGeneratorRequest passed from a protoc execution. This is useful for testing plugins programmatically without having to run protoc. For an example usage, check out [`ast_test.go`](../ast_test.go) in the project root as well as [`testdata/graph`](../testdata/graph) for the test cases. 

Executing the plugin will place a `code_generator_request.pb.bin` file in the specified output location which can be fed directly into a PG* plugin via the `ProtocInput` init option.

## Installation

For a local install:

```bash
make bin/protoc-gen-debug
```

For a global install into `$GOPATH/bin`:

```bash
go install github.com/lyft/protoc-gen-star/protoc-gen-debug
```

## Usage

To create the `code_generator_request.pb.bin` file for all protos in the current directory:

```bash
protoc \
  --plugin=protoc-gen-debug=path/to/protoc-gen-debug \
  --debug_out=".:." \
  *.proto
```

To use the `code_generator_request.pb.bin` in PG*:

```go
func TestModule(t *testing.T) {
  req, err := os.Open("./code_generator_request.pb.bin")
  if err != nil {
    t.Fatal(err)
  }
  
  fs := afero.NewMemMapFs()
  res := &bytes.Buffer{}
  
  pgs.Init(
    pgs.ProtocInput(req),  // use the pre-generated request
    pgs.ProtocOutput(res), // capture CodeGeneratorResponse
    pgs.FileSystem(fs),    // capture any custom files written directly to disk
  ).RegisterModule(&MyModule{}).Render()
  
  // check res and the fs for output
}
```
