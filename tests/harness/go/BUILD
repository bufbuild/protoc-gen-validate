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

genrule(
    name = "go-harness",
    srcs = [":go"],
    outs = ["go-harness"],
    cmd = "cp $(SRCS) $@",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "go",
    embed = [":go_default_library"],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/go",
    visibility = ["//visibility:public"],
)
