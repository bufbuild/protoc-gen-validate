load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library", "go_prefix")
load("//bazel:go_proto_library.bzl", "go_google_protobuf")

go_prefix("github.com/lyft/protoc-gen-validate")

gazelle(
    name = "gazelle",
    external = "vendored",
)

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
