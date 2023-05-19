module github.com/bufbuild/protovalidate/go/v2

go 1.18

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.30.0-20230518171452-0a85abcdd9c8.1
	github.com/envoyproxy/protoc-gen-validate v1.0.1
	github.com/google/cel-go v0.15.2
	github.com/stretchr/testify v1.8.2
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230512164433-5d1fd1a340c9 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go => ./internal/gen
