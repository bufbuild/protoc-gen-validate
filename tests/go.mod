module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	google.golang.org/protobuf v1.27.1
)

replace github.com/envoyproxy/protoc-gen-validate => ../
