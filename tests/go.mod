module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.13
	golang.org/x/net v0.1.0
	google.golang.org/protobuf v1.28.1
)

replace github.com/envoyproxy/protoc-gen-validate => ../
