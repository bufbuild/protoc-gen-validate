# the name of this package
PKG := $(shell go list .)

VALIDATE_IMPORT="Mvalidate/validate.proto=github.com/lyft/protoc-gen-validate/validate"

.PHONY: build
build: generate-annotations # generates the extension necessary for use in the package
	go install .

.PHONY: bazel
bazel: # generate the PGV plugin with Bazel
	bazel build //testdata/...

.PHONY: gazelle
gazelle: # runs gazelle against the codebase to generate Bazel BUILD files
	bazel run //:gazelle

.PHONY: lint
lint: # lints the package for common code smells
	which golint || go get -u github.com/golang/lint/golint
	test -z "$(gofmt -d -s ./*.go)" || (gofmt -d -s ./*.go && exit 1)
	golint -set_exit_status
	go tool vet -all -shadow -shadowstrict *.go

.PHONY: quick
quick: # runs all tests without the race detector or coverage percentage
	go test

.PHONY: tests
tests: # runs all tests against the package with race detection and coverage percentage
	go test -race -cover

.PHONY: cover
cover: # runs all tests against the package, generating a coverage report and opening it in the browser
	go test -race -covermode=atomic -coverprofile=cover.out
	go tool cover -html cover.out -o cover.html
	open cover.html

.PHONY: harness
harness: harness/go/go-harness # runs the test harness, validating a series of test cases in all supported languages
	go run ./harness/executor/*.go

.PHONY: generate-annotations
generate-annotations: # generates the proto extension in Go
	cd validate && \
	protoc \
		-I . \
		--go_out=. \
		validate.proto

.PHONY: generate-kitchensink
generate-kitchensink: # generates the kitchensink test protos
	rm -r testdata/kitchensink/generated || true
	mkdir -p testdata/kitchensink/generated/go
	cd testdata/kitchensink/protos && \
	protoc \
		-I . \
		-I ../../.. \
		--go_out="${VALIDATE_IMPORT}:../generated/go" \
		--validate_out="lang=go:../generated/go" \
		`find . -name "*.proto"`

.PHONY: generate-harness
generate-harness: # generates the test harness protos
	cd harness && \
	protoc \
		-I . \
		--go_out=. \
		harness.proto

.PHONY: generate-testcases
generate-testcases: # generate the test harness cases
	rm -r harness/cases/go || true
	mkdir -p harness/cases/go
	cd harness/cases && \
	protoc \
		-I . \
		-I ../.. \
		--go_out="${VALIDATE_IMPORT}:./go" \
		--validate_out="lang=go:./go" \
		`find . -name "*.proto"`

harness/go/go-harness:
	go build -o ./harness/go/go-harness ./harness/go
