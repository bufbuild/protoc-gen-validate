module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0
	github.com/golang/protobuf v1.3.1
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
	github.com/iancoleman/strcase v0.0.0-20180726023541-3605ed457bf7
	google.golang.org/genproto v0.0.0-20200729003335-053ba62fc06f
)

replace github.com/envoyproxy/protoc-gen-validate => ../
