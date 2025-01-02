module github.com/envoyproxy/protoc-gen-validate

go 1.22.0

toolchain go1.23.1

require (
	github.com/iancoleman/strcase v0.3.0
	github.com/lyft/protoc-gen-star/v2 v2.0.4-0.20230330145011-496ad1ac90a4
	golang.org/x/net v0.33.0
	google.golang.org/protobuf v1.36.1
)

require (
	github.com/spf13/afero v1.11.0 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.28.0 // indirect
)

retract [v0.6.9, v0.6.12] // Published accidentally
