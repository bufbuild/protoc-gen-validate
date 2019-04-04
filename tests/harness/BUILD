load("@io_bazel_rules_go//go:def.bzl", "go_library")

# gazelle:exclude harness.pb.go

load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "harness_proto",
    srcs = ["harness.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:any_proto",
    ],
)

go_proto_library(
    name = "harness_go_proto",
    importpath = "github.com/envoyproxy/protoc-gen-validate/tests/harness/go",
    proto = ":harness_proto",
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "harness_gogo_proto",
    compilers = ["@io_bazel_rules_go//proto:gogo_proto"],
    importpath = "github.com/envoyproxy/protoc-gen-validate/tests/harness/gogo",
    proto = ":harness_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_protobuf//ptypes/any:go_default_library",
        "@com_github_golang_protobuf//ptypes/timestamp:go_default_library",
        "@com_github_golang_protobuf//ptypes/wrappers:go_default_library",
        "@io_bazel_rules_go//proto/wkt:descriptor_go_proto",
    ],
)

cc_proto_library(
    name = "harness_cc_proto",
    visibility = ["//visibility:public"],
    deps = [
        ":harness_proto",
    ],
)

java_proto_library(
    name = "harness_java_proto",
    visibility = ["//visibility:public"],
    deps = [":harness_proto"],
)

go_library(
    name = "go_default_library",
    embed = [":harness_go_proto"],
    importpath = "github.com/envoyproxy/protoc-gen-validate/tests/harness",
    visibility = ["//visibility:public"],
)
