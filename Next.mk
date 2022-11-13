# The next Makefile

name := protoc-gen-validate

# Root dir returns absolute path of current directory. It has a trailing "/".
root_dir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Currently we resolve it using which. But more sophisticated approach is to use infer GOROOT.
go     := $(shell which go)
goarch := $(shell $(go) env GOARCH)
goexe  := $(shell $(go) env GOEXE)
goos   := $(shell $(go) env GOOS)

# The current binary info.
current_binary_path := build/$(name)_$(goos)_$(goarch)
current_binary      := $(current_binary_path)/$(name)$(goexe)

# Overridable variables.
# Currently, harness tests only run against C++ and Go implementations.
# Note: we do harness tests for all languages via bazel.
# TODO(dio): Run harness to all supported languages.
HARNESS_LANGUAGES ?= cc go
# Local cache directory.
CACHE_DIR ?= $(root_dir).cache

# Directory of Go tools.
go_tools_dir := $(CACHE_DIR)/tools/go

# Directory of prepackaged tools (e.g. protoc).
prepackaged_tools_dir := $(CACHE_DIR)/tools/prepackaged

export PATH := $(go_tools_dir)/bin:$(prepackaged_tools_dir)/bin:$(root_dir)$(current_binary_path):$(PATH)

include $(root_dir)tools/build/Help.mk
include $(root_dir)tools/build/Env.mk

# Include versions of tools we build or fetch on-demand.
include $(root_dir)Tools.mk

# Prepackaged tools.
protoc := $(prepackaged_tools_dir)/bin/protoc

# Go based tools.
bazel         := $(go_tools_dir)/bin/bazelisk
buildifier    := $(go_tools_dir)/bin/buildifier
protoc-gen-go := $(go_tools_dir)/bin/protoc-gen-go
gosimports    := $(go_tools_dir)/bin/gosimports
golangci-lint := $(or $(shell which golangci-lint),$(go_tools_dir)/bin/golangci-lint)

bazel_files       := WORKSPACE BUILD.bazel $(shell find . \( -name "*.bzl" -or -name "*.bazel" -or -name "BUILD" \) -not -path "./bazel-*" -not -path "./.cache")
nongen_go_sources := $(shell find . -name "*.go" -not -path "*.pb.go" -not -path "*.pb.validate.go" -not -path "./templates/go/file.go" -not -path "./bazel-*" -not -path "./.cache")

# Harness executables.
go_harness   := $(root_dir)tests/harness/go/main/go-harness
cc_harness   := $(root_dir)tests/harness/cc/cc-harness

# List of harness test cases for Go.
tests_harness_cases_go := \
	/harness \
	/harness/cases \
	/harness/cases/other_package \
	/harness/cases/yet_another_package

# The main generated file.
validate_pb_go := validate/validate.pb.go

test: $(bazel) $(tests_harness_cases_go) ## Runs PGV tests
	@$(bazel) test //tests/... --test_output=errors

build: $(current_binary) ## Builds PGV binary

harness: $(go_harness) $(cc_harness) ## Runs PGV harness test
	@$(go) run tests/harness/executor/*.go $(addprefix -,$(HARNESS_LANGUAGES))

# Note: ./templates/go/file.go is ignored.
# TODO(dio): Format *.proto using buf format.
# TODO(dio): Format files outside Go and Python.
format: $(buildifier) $(gosimports) ## Format source code files.
	@$(buildifier) --lint=fix $(bazel_files)
	@$(go) mod tidy
	@$(go)fmt -s -w $(nongen_go_sources)
	@$(gosimports) -local $$(sed -ne 's/^module //gp' go.mod) -w $(nongen_go_sources)
# The following requires "pip3 install -r requirements.txt". For macOS, the default "bin" directory
# most likely is at ~/Library/Python/3.9/bin.
	@flake8 --config=python/setup.cfg python/protoc_gen_validate/validator.py
	@isort --check-only python/protoc_gen_validate/validator.py

# TODO(dio): Lint non-Go files.
lint: .golangci.yml $(golangci-lint) $(nongen_go_sources) ## Lint source code files.
	@$(golangci-lint) run --timeout 5m --config $< ./...

check: ## Check consistency
	@if [ ! -z "`git status -s`" ]; then \
		echo "The following differences will fail CI until committed:"; \
		git diff --exit-code; \
	fi

clean: ## Clean all build and test artifacts
	@rm -f $(validate_pb_go)
	@rm -f $(current_binary) $(shell find tests \( -name "*.pb.go" -or -name "*.pb.validate.go" \))
	@rm -f $(go_harness) $(cc_harness)

bazel-build: $(bazel) ## Builds PGV binary using bazel
	@$(bazel) build //:$(name)
	@mkdir -p $(current_binary_path)
	@cp -f bazel-bin/$(name)_/$(name)$(goexe) $(current_binary)

bazel-build-tests-generation: $(bazel) ## Builds //tests/generation/... using bazel
	@$(bazel) build //tests/generation/...

bazel-test-example-workspace: $(bazel) ## Tests example workspace using bazel
	@cd example-workspace && bazel test //... --test_output=errors

bazel-gazelle: $(bazel) ## Runs gazelle against the codebase to generate Bazel BUILD files
	@$(bazel) run //:gazelle -- update-repos -from_file=go.mod -prune -to_macro=dependencies.bzl%go_third_party
	@$(bazel) run //:gazelle

# Build target for current binary.
build/$(name)_%/$(name)$(goexe): $(validate_pb_go)
	@GOBIN=$(root_dir)$(current_binary_path) $(go) install .

# Generate validate/validate.pb.go from validate/validate.proto.
$(validate_pb_go): $(protoc) $(protoc-gen-go) validate/validate.proto
	@$(protoc) -I . --go_opt=paths=source_relative --go_out=. $(filter %.proto,$^)

# Build go-harnes executable.
$(go_harness): $(tests_harness_cases_go)
	@cd tests/harness/go/main && $(go) build -o $@ .

# Build cc-harnes executable.
$(cc_harness):
	@bazel build //tests/harness/cc:cc-harness
	@cp bazel-bin/tests/harness/cc/cc-harness $@

# Generate all required files for harness tests in Go.
$(tests_harness_cases_go): $(current_binary)
	$(call generate-test-cases-go,tests$@)

# Generates a test-case for Go.
define generate-test-cases-go
	@cd $1 && \
	mkdir -p go && \
	$(protoc) \
		-I . \
		-I $(root_dir) \
		--go_opt=paths=source_relative \
		--go_out=go \
		--validate_opt=paths=source_relative \
		--validate_out=lang=go:go \
		*.proto
endef

include $(root_dir)tools/build/Installer.mk
