load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//bazel:go_proto_library.bzl", "go_google_protobuf")

go_binary(
    name = "protoc-gen-validate",
    embed = [":go_default_library"],
    importpath = "github.com/lyft/protoc-gen-validate",
    visibility = ["//visibility:public"],
)

go_google_protobuf()

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/lyft/protoc-gen-validate",
    visibility = ["//visibility:private"],
    deps = [
        "//module:go_default_library",
        "//vendor/github.com/lyft/protoc-gen-star:go_default_library",
    ],
)
