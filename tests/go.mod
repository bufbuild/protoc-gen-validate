module github.com/envoyproxy/protoc-gen-validate/tests

go 1.21.1

require (
	golang.org/x/net v0.19.0
	google.golang.org/protobuf v1.32.0
)

require golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect

replace github.com/envoyproxy/protoc-gen-validate => ../
