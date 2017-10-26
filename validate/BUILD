# gazelle:exclude validate.pb.go

load("//bazel:go_proto_library.bzl", "go_proto_library")
load("//bazel:protobuf.bzl", "cc_proto_library")

cc_proto_library(
  name = "validate_cc",
  srcs = ["validate.proto"],
  protoc = "@protobuf_bzl//:protoc",
  default_runtime = "@protobuf_bzl//:protobuf",
  deps = ["@protobuf_bzl//:cc_wkt_protos"],
  visibility = ["//visibility:public"],
)

go_proto_library(
    name = "go_default_library",
    srcs = ["validate.proto"],
    protoc = "@protobuf_bzl//:protoc",
    rules_go_repo_only_for_internal_use = "",
    deps = [
        "@com_github_golang_protobuf//ptypes/duration:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
        "@com_github_golang_protobuf//protoc-gen-go/descriptor:go_default_library",
    ],
    importpath = "github.com/lyft/protoc-gen-validate/validate",
    visibility = ["//visibility:public"],
)
