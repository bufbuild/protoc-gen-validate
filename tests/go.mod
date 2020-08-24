module github.com/envoyproxy/protoc-gen-validate/tests

go 1.14

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.4.2
	github.com/iancoleman/strcase v0.0.0-20180726023541-3605ed457bf7
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
	google.golang.org/genproto v0.0.0-20200729003335-053ba62fc06f
	google.golang.org/protobuf v1.25.0
)

replace github.com/envoyproxy/protoc-gen-validate => ../
