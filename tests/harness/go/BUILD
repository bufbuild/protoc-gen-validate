load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["harness.go"],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/go",
    visibility = ["//visibility:private"],
    deps = [
        "//tests/harness:go_default_library",
        "//tests/harness/cases:go",
        "//tests/harness/cases/other_package:go",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes:go_default_library",
    ],
)

go_binary(
    name = "go-harness-binary",
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/go",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

# todo: pass binary location to executor
# locations of binaries should not be depended on
# (see https://github.com/bazelbuild/rules_go#how-do-i-access-go-binary-executables-from-go-test)
# this is a temporary workaround to allow the tests to pass on rules_go 0.9.0
genrule(
    name = "go-harness",
    srcs = [":go-harness-binary"],
    outs = ['go-harness'],
    cmd = "cp $(location :go-harness-binary) $@",
    visibility = ["//visibility:public"],
)
