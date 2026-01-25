module github.com/envoyproxy/protoc-gen-validate

go 1.24.1

require (
	github.com/iancoleman/strcase v0.3.0
	github.com/lyft/protoc-gen-star/v2 v2.0.4
	golang.org/x/net v0.49.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/spf13/afero v1.15.0 // indirect
	golang.org/x/mod v0.32.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	golang.org/x/tools v0.41.0 // indirect
)

retract [v0.6.9, v0.6.12] // Published accidentally
