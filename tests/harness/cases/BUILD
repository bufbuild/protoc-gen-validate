# gazelle:exclude go

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
