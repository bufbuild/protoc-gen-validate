load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["harness.go"],
    importpath = "github.com/envoyproxy/protoc-gen-validate/tests/harness/gogo/main",
    visibility = ["//visibility:private"],
    deps = [
        "//tests/harness:harness_gogo_proto",
        "//tests/harness/cases:gogo",
        "//tests/harness/cases/other_package:gogo",
        "@com_github_gogo_protobuf//proto:go_default_library",
        "@com_github_gogo_protobuf//types:go_default_library",
    ],
)

genrule(
    name = "go-harness-bin",
    srcs = [":main"],
    outs = ["go-harness"],
    cmd = "cp $(SRCS) $@",
    visibility = ["//visibility:public"],
)

genrule(
    name = "go-harness-exe",
    srcs = [":main"],
    outs = ["go-harness.exe"],
    cmd = "cp $(SRCS) $@",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "main",
    embed = [":go_default_library"],
    importpath = "github.com/envoyproxy/protoc-gen-validate/tests/harness/gogo/main",
    visibility = ["//visibility:public"],
)
