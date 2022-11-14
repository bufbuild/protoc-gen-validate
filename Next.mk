# Licensed under the Apache License, Version 2.0 (the "License")

# Plugin name.
name := protoc-gen-validate

# Root dir returns absolute path of current directory. It has a trailing "/".
root_dir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Local cache directory.
cache_dir := $(root_dir).cache

# Directory of Go tools.
go_tools_dir := $(cache_dir)/tools/go

# Directory of prepackaged tools (e.g. protoc).
prepackaged_tools_dir := $(cache_dir)/tools/prepackaged

# Currently we resolve Go using `which`. But more sophisticated approach is to use infer GOROOT.
go     := $(shell which go)
goarch := $(shell $(go) env GOARCH)
goexe  := $(shell $(go) env GOEXE)
goos   := $(shell $(go) env GOOS)

# The current binary location for the current runtime (goos, goarch). We install our plugin here.
current_binary_path := build/$(name)_$(goos)_$(goarch)
current_binary      := $(current_binary_path)/$(name)$(goexe)

# This makes sure protoc can access the installed plugins.
export PATH := $(root_dir)$(current_binary_path):$(go_tools_dir)/bin:$(prepackaged_tools_dir)/bin:$(PATH)

# The main generated file.
validate_pb_go := validate/validate.pb.go

# Include versions of tools we build on-demand
include Tools.mk
# This provides the "help" target.
include tools/build/Help.mk
# This sets some required environment variables.
include tools/build/Env.mk

# Path to the installed protocol buffer compiler.
protoc := $(prepackaged_tools_dir)/bin/protoc

# Go based tools.
protoc-gen-go := $(go_tools_dir)/bin/protoc-gen-go

build: $(current_binary) ## Build the plugin

clean: ## Clean all build and test artifacts
	@rm -f $(validate_pb_go)
	@rm -f $(current_binary)

check: ## Verify contents of last commit
	@# Make sure the check-in is clean
	@if [ ! -z "`git status -s`" ]; then \
		echo "The following differences will fail CI until committed:"; \
		git diff --exit-code; \
	fi

# Generate validate/validate.pb.go from validate/validate.proto.
$(validate_pb_go): $(protoc) $(protoc-gen-go) validate/validate.proto
	@$(protoc) -I . --go_opt=paths=source_relative --go_out=. $(filter %.proto,$^)

# Build target for current binary.
build/$(name)_%/$(name)$(goexe): $(validate_pb_go)
	@GOBIN=$(root_dir)$(current_binary_path) $(go) install .

include tools/build/Installer.mk
