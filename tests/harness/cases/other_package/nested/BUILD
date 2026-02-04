load("@com_google_protobuf//bazel:proto_library.bzl", "proto_library")
load("@com_google_protobuf//bazel:py_proto_library.bzl", "py_proto_library")
load("@rules_java//java:defs.bzl", "java_proto_library")
load(
    "//bazel:pgv_proto_library.bzl",
    "pgv_cc_proto_library",
    "pgv_go_proto_library",
    "pgv_java_proto_library",
)

proto_library(
    name = "enum_proto",
    srcs = [
        "enum.proto",
    ],
    visibility = ["//visibility:public"],
)

pgv_go_proto_library(
    name = "go",
    importpath = "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/nested/go",
    proto = ":enum_proto",
)

pgv_cc_proto_library(
    name = "cc",
    visibility = ["//tests:__subpackages__"],
    deps = [":enum_proto"],
)

java_proto_library(
    name = "enum_java_proto",
    visibility = ["//visibility:public"],
    deps = [":enum_proto"],
)

pgv_java_proto_library(
    name = "java",
    java_deps = [":enum_java_proto"],
    visibility = ["//visibility:public"],
    deps = [":enum_proto"],
)

py_proto_library(
    name = "enum_python_proto",
    visibility = ["//visibility:public"],
    deps = [":enum_proto"],
)
