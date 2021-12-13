module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
	google.golang.org/protobuf v1.27.1
)

replace github.com/envoyproxy/protoc-gen-validate => ../
