// +build tools

package main

import (
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow"
	_ "golang.org/x/net/context"
)
