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
	# golint -set_exit_status
	go tool vet -all -shadow -shadowstrict *.go

.PHONY: quick
quick:
	# runs all tests without the race detector or coverage percentage
	go test

.PHONY: tests
tests:
	# runs all tests against the package with race detection and coverage percentage
	go test -race -cover
	# tests validate proto generation
	bazel build //validate:go_default_library \
		&& diff $$(bazel info bazel-genfiles)/validate/validate.pb.go validate/validate.pb.go

.PHONY: cover
cover:
	# runs all tests against the package, generating a coverage report and opening it in the browser
	go test -race -covermode=atomic -coverprofile=cover.out
	go tool cover -html cover.out -o cover.html
	open cover.html

.PHONY: harness
harness: tests/harness/harness.pb.go tests/harness/go/go-harness tests/harness/cc/cc-harness
 	# runs the test harness, validating a series of test cases in all supported languages
	go run ./tests/harness/executor/*.go

.PHONY: bazel-harness
bazel-harness:
	# runs the test harness via bazel
	bazel run //tests/harness/executor:executor

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
	cd tests/kitchensink/go && go build .

.PHONY: testcases
testcases:
	# generate the test harness case protos
	rm -r tests/harness/cases/go || true
	mkdir tests/harness/cases/go
	rm -r tests/harness/cases/other_package/go || true
	mkdir tests/harness/cases/other_package/go
	# protoc-gen-go makes us go a package at a time
	cd tests/harness/cases/other_package && \
	protoc \
		-I . \
		-I ../../../.. \
		--go_out="${VALIDATE_IMPORT}:./go" \
		--validate_out="lang=go:./go" \
		./*.proto
	cd tests/harness/cases && \
	protoc \
		-I . \
		-I ../../.. \
		--go_out="Mtests/harness/cases/other_package/embed.proto=github.com/lyft/protoc-gen-validate/tests/harness/cases/other_package/go,${VALIDATE_IMPORT}:./go" \
		--validate_out="lang=go:./go" \
		./*.proto

.PHONY: update-vendor
update-vendor:
	# updates the vendored dependencies using the Go Dep tool
	dep ensure -update
	$(MAKE) gazelle

tests/harness/harness.pb.go:
	# generates the test harness protos
	cd tests/harness && protoc -I . --go_out=. harness.proto

tests/harness/go/go-harness:
	# generates the go-specific test harness
	go build -o ./tests/harness/go/go-harness ./tests/harness/go

tests/harness/cc/cc-harness: tests/harness/cc/harness.cc
	# generates the C++-specific test harness
	# use bazel which knows how to pull in the C++ common proto libraries
	bazel build //tests/harness/cc:cc-harness
	cp bazel-bin/tests/harness/cc/cc-harness $@
	chmod 0755 $@

.PHONY: ci
ci: lint build tests kitchensink testcases harness bazel-harness
