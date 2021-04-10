module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
	google.golang.org/protobuf v1.26.0
)

replace github.com/envoyproxy/protoc-gen-validate => ../
