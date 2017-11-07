# protoc-gen-go parameters for properly generating the import path for PGV
VALIDATE_IMPORT="Mvalidate/validate.proto=github.com/lyft/protoc-gen-validate/validate"

.PHONY: build
build: validate/validate.pb.go
	# generates the PGV binary and installs it into $GOPATH/bin
	go install .

.PHONY: bazel
bazel:
	# generate the PGV plugin with Bazel
	bazel build //tests/...

.PHONY: gazelle
gazelle:
	# runs gazelle against the codebase to generate Bazel BUILD files
	bazel run //:gazelle

.PHONY: lint
lint:
	# lints the package for common code smells
	which golint || go get -u github.com/golang/lint/golint
	test -z "$(gofmt -d -s ./*.go)" || (gofmt -d -s ./*.go && exit 1)
	golint -set_exit_status
	go tool vet -all -shadow -shadowstrict *.go

.PHONY: quick
quick:
	# runs all tests without the race detector or coverage percentage
	go test

.PHONY: tests
tests:
	# runs all tests against the package with race detection and coverage percentage
	go test -race -cover

.PHONY: cover
cover:
	# runs all tests against the package, generating a coverage report and opening it in the browser
	go test -race -covermode=atomic -coverprofile=cover.out
	go tool cover -html cover.out -o cover.html
	open cover.html

.PHONY: harness
harness: tests/harness/harness.pb.go tests/harness/go/go-harness
 	# runs the test harness, validating a series of test cases in all supported languages
	go run ./tests/harness/executor/*.go

.PHONY: bazel-harness
bazel-harness:
	# runs the test harness via bazel
	bazel run //tests/harness/executor

.PHONY: kitchensink
kitchensink:
	# generates the kitchensink test protos
	rm -r tests/kitchensink/go || true
	mkdir -p tests/kitchensink/go
	cd tests/kitchensink && \
	protoc \
		-I . \
		-I ../.. \
		--go_out="${VALIDATE_IMPORT}:./go" \
		--validate_out="lang=go:./go" \
		`find . -name "*.proto"`

.PHONY: testcases
testcases:
	# generate the test harness case protos
	rm -r tests/harness/cases/go || true
	mkdir -p tests/harness/cases/go
	cd tests/harness/cases && \
	protoc \
		-I . \
		-I ../../.. \
		--go_out="${VALIDATE_IMPORT}:./go" \
		--validate_out="lang=go:./go" \
		`find . -name "*.proto"`

validate/validate.pb.go:
	# generates the proto extension in Go
	cd validate && protoc -I . --go_out=. validate.proto && cp github.com/lyft/protoc-gen-validate/validate/validate.pb.go .

tests/harness/harness.pb.go:
	# generates the test harness protos
	cd tests/harness && protoc -I . --go_out=. harness.proto

tests/harness/go/go-harness:
	# generates the go-specific test harness
	go build -o ./tests/harness/go/go-harness ./tests/harness/go

.PHONY: ci
ci: build tests kitchensink testcases harness
