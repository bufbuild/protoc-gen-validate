load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

#gazelle:exclude go/

load(
    "//bazel:pgv_proto_library.bzl",
    "pgv_cc_proto_library",
    "pgv_go_proto_library",
)

pgv_go_proto_library(
    name = "fixed32_go",
    srcs = ["fixed32.proto"],
)

pgv_cc_proto_library(
    name = "fixed32_cc",
    srcs = ["fixed32.proto"],
)

proto_library(
    name = "tests_kitchensink_proto",
    srcs = [
        "any.proto",
        "bool.proto",
        "bytes.proto",
        "disabled.proto",
        "duration.proto",
        "enum.proto",
        "fixed32.proto",
        "fixed64.proto",
        "float.proto",
        "int32.proto",
        "int64.proto",
        "map.proto",
        "message.proto",
        "oneof.proto",
        "repeated.proto",
        "sfixed32.proto",
        "sfixed64.proto",
        "sint32.proto",
        "sint64.proto",
        "string.proto",
        "timestamp.proto",
        "uint32.proto",
        "uint64.proto",
        "wrapper.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//validate:validate_proto",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@com_google_protobuf//:wrappers_proto",
    ],
)

go_proto_library(
    name = "tests_kitchensink_go_proto",
    importpath = "github.com/lyft/protoc-gen-validate/tests/kitchensink",
    proto = ":tests_kitchensink_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//validate:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/any:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/duration:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/wrappers:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    embed = [":tests_kitchensink_go_proto"],
    importpath = "github.com/lyft/protoc-gen-validate/tests/kitchensink",
    visibility = ["//visibility:public"],
)
