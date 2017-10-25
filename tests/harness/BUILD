# gazelle:exclude harness.pb.go

load("//bazel:go_proto_library.bzl", "go_proto_library")
load("//bazel:protobuf.bzl", "cc_proto_library")

go_proto_library(
    name = "go_default_library",
    srcs = ["harness.proto"],
    protoc = "@protobuf_bzl//:protoc",
    rules_go_repo_only_for_internal_use = "",
    deps = [
        "@com_github_golang_protobuf//ptypes/any:go_default_library",
        "@com_github_golang_protobuf//protoc-gen-go/descriptor:go_default_library",
    ],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness",
    visibility = ["//visibility:public"],
)
