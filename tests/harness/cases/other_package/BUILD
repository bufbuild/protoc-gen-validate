# gazelle:exclude go
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

load(
    "//bazel:pgv_proto_library.bzl",
    "pgv_go_proto_library",
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

go_library(
    name = "gogo",
    srcs = glob(["gogo/*.go"]),
    visibility = ["//visibility:public"],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/cases/other_package/gogo",
    deps = [
        "//tests/harness/gogo:go_default_library",
        "//validate:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/types:go_default_library",
    ],
)

pgv_cc_proto_library(
    name = "cc",
    deps = [":embed_proto"],
    visibility = ["//tests:__subpackages__"],
)
