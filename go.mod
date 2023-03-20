module github.com/envoyproxy/protoc-gen-validate

go 1.20

require (
	github.com/iancoleman/strcase v0.2.0
	github.com/lyft/protoc-gen-star/v2 v2.0.1
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net v0.8.0
	golang.org/x/tools v0.7.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/spf13/afero v1.3.3 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)

retract [v0.6.9, v0.6.12] // Published accidentally
