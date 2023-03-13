module github.com/envoyproxy/protoc-gen-validate

go 1.17

require (
	github.com/iancoleman/strcase v0.2.0
	github.com/lyft/protoc-gen-star/v2 v2.0.1
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net v0.7.0
	golang.org/x/tools v0.4.0
	google.golang.org/protobuf v1.29.0
)

require (
	github.com/spf13/afero v1.9.2 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)

retract [v0.6.9, v0.6.12] // Published accidentally
