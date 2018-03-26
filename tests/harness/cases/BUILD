# gazelle:exclude go
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

load(
    "//bazel:pgv_proto_library.bzl",
    "pgv_go_proto_library",
    "pgv_cc_proto_library",
)

pgv_go_proto_library(
    name = "go",
    srcs = [
        "bool.proto",
        "bytes.proto",
        "enums.proto",
        "maps.proto",
        "messages.proto",
        "numbers.proto",
        "oneofs.proto",
        "repeated.proto",
        "strings.proto",
        "wkt_any.proto",
        "wkt_duration.proto",
        "wkt_timestamp.proto",
        "wkt_wrappers.proto",
    ],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/cases/go",
    rules_go_repo_only_for_internal_use = "",
    deps = [
        "//tests/harness/cases/other_package:go",
        "@com_github_golang_protobuf//ptypes/any:go_default_library",
        "@com_github_golang_protobuf//ptypes/duration:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
        "@com_github_golang_protobuf//ptypes/wrappers:go_default_library",
    ],
)

go_library(
    name = "gogo",
    srcs = glob(["gogo/*.go"]),
    visibility = ["//visibility:public"],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/cases/gogo",
    deps = [
        "//tests/harness/gogo:go_default_library",
        "//tests/harness/cases/other_package:gogo",
        "//validate:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/types:go_default_library",
    ],
)

pgv_cc_proto_library(
    name = "cc",
    srcs = [
        "bool.proto",
        "bytes.proto",
        "enums.proto",
        "maps.proto",
        "messages.proto",
        "numbers.proto",
        "oneofs.proto",
        "repeated.proto",
        "strings.proto",
        "wkt_any.proto",
        "wkt_duration.proto",
        "wkt_timestamp.proto",
        "wkt_wrappers.proto",
    ],
    deps = [
        "//tests/harness/cases/other_package:cc",
    ],
    visibility = ["//tests:__subpackages__"],
)
