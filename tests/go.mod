module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	golang.org/x/net v0.5.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.28.1
)

replace github.com/envoyproxy/protoc-gen-validate => ../
