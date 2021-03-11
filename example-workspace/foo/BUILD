load("@com_envoyproxy_protoc_gen_validate//bazel:pgv_proto_library.bzl", "pgv_cc_proto_library")

package(
    default_visibility = ["//visibility:public"],
)

proto_library(
    name = "bar_proto",
    srcs = ["bar.proto"],
    deps = ["@com_envoyproxy_protoc_gen_validate//validate:validate_proto"],
)

pgv_cc_proto_library(
    name = "bar_cc_proto",
    deps = [":bar_proto"],
)
