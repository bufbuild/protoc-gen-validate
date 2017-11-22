# gazelle:exclude harness.pb.go

load("//bazel:go_proto_library.bzl", "go_proto_library")
load("@com_google_protobuf//:protobuf.bzl", "cc_proto_library")

go_proto_library(
    name = "go_default_library",
    srcs = ["harness.proto"],
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness",
    protoc = "@com_google_protobuf//:protoc",
    rules_go_repo_only_for_internal_use = "",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_protobuf//protoc-gen-go/descriptor:go_default_library",
        "@com_github_golang_protobuf//ptypes/any:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
        "@com_github_golang_protobuf//ptypes/wrappers:go_default_library",
    ],
)

cc_proto_library(
    name = "harness_proto",
    srcs = ["harness.proto"],
    default_runtime = "@com_google_protobuf//:protobuf",
    protoc = "@com_google_protobuf//:protoc",
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:cc_wkt_protos",
    ],
)
