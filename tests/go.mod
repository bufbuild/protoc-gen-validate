module github.com/envoyproxy/protoc-gen-validate/tests

go 1.14

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.1
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
)

replace github.com/envoyproxy/protoc-gen-validate => ../
