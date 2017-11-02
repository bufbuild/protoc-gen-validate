load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cases.go",
        "executor.go",
        "harness.go",
        "worker.go",
    ],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/executor",
    visibility = ["//visibility:private"],
    deps = [
        "//tests/harness:go_default_library",
        "//tests/harness/cases:go",
        "//tests/harness/go:go-harness",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "@com_github_golang_protobuf//ptypes/any:go_default_library",
        "@com_github_golang_protobuf//ptypes/duration:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
        "@com_github_golang_protobuf//ptypes/wrappers:go_default_library",
    ],
)

go_binary(
    name = "executor",
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/executor",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
