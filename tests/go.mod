module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.9.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.30.0
)

replace github.com/envoyproxy/protoc-gen-validate => ../
