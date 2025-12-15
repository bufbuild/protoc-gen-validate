module github.com/envoyproxy/protoc-gen-validate/tests

go 1.24.1

require (
	golang.org/x/net v0.48.0
	google.golang.org/protobuf v1.36.11
)

replace github.com/envoyproxy/protoc-gen-validate => ../
