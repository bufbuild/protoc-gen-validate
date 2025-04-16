module github.com/envoyproxy/protoc-gen-validate

go 1.24

require (
	github.com/iancoleman/strcase v0.3.0
	github.com/lyft/protoc-gen-star/v2 v2.0.4
	golang.org/x/net v0.39.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/spf13/afero v1.14.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
)

retract [v0.6.9, v0.6.12] // Published accidentally
