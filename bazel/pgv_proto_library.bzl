load(":go_proto_library.bzl", "go_proto_library")
load(":protobuf.bzl", "cc_proto_library")

def pgv_go_proto_library(name, srcs = None, deps = [], **kwargs):
    go_proto_library(name,
                     srcs,
                     deps = ["//validate:go_default_library"] + deps,
                     protoc = "@protobuf_bzl//:protoc",
                     validate = 1,
                     **kwargs)

def pgv_cc_proto_library(name, srcs = None, deps = [], **kwargs):
    cc_proto_library(name,
                     srcs,
                     protoc = "@protobuf_bzl//:protoc",
                     default_runtime = "@protobuf_bzl//:protobuf",
                     deps = ["//validate:validate_cc"] + deps,
                     validate = 1,
                     **kwargs)
