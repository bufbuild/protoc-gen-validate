# gazelle:exclude go

load(
    "//bazel:pgv_proto_library.bzl",
    "pgv_go_proto_library",
    "pgv_cc_proto_library",
)

pgv_go_proto_library(
    name = "go",
    srcs = [
        "embed.proto",
    ],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness/cases/other_package/go",
    rules_go_repo_only_for_internal_use = "",
)

pgv_cc_proto_library(
    name = "cc",
    srcs = [
        "embed.proto",
    ],
    visibility = ["//tests:__subpackages__"],
)
