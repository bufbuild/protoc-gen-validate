# the name of this package
PKG := $(shell go list .)

.PHONY: build
build: # generates the extension necessary for use in the package
	cd validate && protoc -I . --go_out=. validate.proto
	go install .
	which protoc-gen-validate

.PHONY: bazel
bazel:
	bazel build //testdata/...

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

.PHONY: generate-testdata
generate-testdata: # generates the test protos
	rm -r testdata/generated || true
	mkdir -p testdata/generated/go
	cd testdata/protos && \
	protoc \
		-I . \
		-I ../.. \
		-I $(GOPATH)/src \
		--go_out=":../generated/go" \
		--validate_out="lang=go:../generated/go" \
		`find . -name "*.proto"`
