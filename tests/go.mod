module github.com/envoyproxy/protoc-gen-validate/tests

go 1.19

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b
	google.golang.org/protobuf v1.28.1
)

replace github.com/envoyproxy/protoc-gen-validate => ../
