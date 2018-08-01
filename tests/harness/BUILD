# gazelle:exclude harness.pb.go

load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@com_google_protobuf//:protobuf.bzl", "cc_proto_library")

proto_library(
    name = "harness_proto_target",
    srcs = ["harness.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
    ],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "go_default_library",
    proto = ":harness_proto_target",
    importpath = "github.com/lyft/protoc-gen-validate/tests/harness",
    visibility = ["//visibility:public"],
    deps = [
        "@io_bazel_rules_go//proto/wkt:descriptor_go_proto",
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
