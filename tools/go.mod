module github.com/bufbuild/protovalidate/tools

go 1.18

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.30.0-20230518171452-0a85abcdd9c8.1
	github.com/bufbuild/protocompile v0.5.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.2
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc
	golang.org/x/sync v0.2.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go => ../go/v2/internal/gen
