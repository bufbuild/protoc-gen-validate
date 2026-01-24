module github.com/envoyproxy/protoc-gen-validate/tests

go 1.24.1

require (
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	golang.org/x/net v0.49.0
	google.golang.org/protobuf v1.36.11
)

replace github.com/envoyproxy/protoc-gen-validate => ../
