# gazelle:exclude go
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

load(
    "//bazel:pgv_proto_library.bzl",
    "pgv_go_proto_library",
    "pgv_gogo_proto_library",
    "pgv_cc_proto_library",
)

proto_library(
    name = "embed_proto",
    srcs = [
        "embed.proto",
    ],
    visibility = ["//visibility:public"],
    deps = ["//validate:validate_proto"],
)

pgv_go_proto_library(
    name = "go",
    proto = ":embed_proto",
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/cases/other_package/go",
    deps = [
        "@com_github_golang_protobuf//ptypes:go_default_library",
    ],
)

pgv_gogo_proto_library(
    name = "gogo",
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/cases/other_package/gogo",
    proto = ":embed_proto",
    deps = [
        "@com_github_golang_protobuf//ptypes:go_default_library",
    ],
)

pgv_cc_proto_library(
    name = "cc",
    deps = [":embed_proto"],
    visibility = ["//tests:__subpackages__"],
)
