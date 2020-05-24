module github.com/envoyproxy/protoc-gen-validate/tests

go 1.14

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0
	github.com/golang/protobuf v1.4.2
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/protobuf v1.23.0
)

replace github.com/envoyproxy/protoc-gen-validate => ./..
