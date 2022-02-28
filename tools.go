//go:build tools
// +build tools

package main

import (
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/net/context"
	_ "golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
