module github.com/envoyproxy/protoc-gen-validate

go 1.12

require (
	github.com/iancoleman/strcase v0.2.0
	github.com/lyft/protoc-gen-star v0.6.1
	github.com/spf13/afero v1.9.2 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net v0.1.0
	golang.org/x/tools v0.2.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.28.1
)

retract [v0.6.9, v0.6.12] // Published accidentally
