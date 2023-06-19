module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	golang.org/x/net v0.11.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.30.0
)

replace github.com/envoyproxy/protoc-gen-validate => ../
