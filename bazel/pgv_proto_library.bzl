load(":go_proto_library.bzl", "go_proto_library")
load(":protobuf.bzl", "cc_proto_library")

def pgv_go_proto_library(name, srcs = None, deps = [], **kwargs):
    go_proto_library(name,
                     srcs,
                     deps = ["//validate:go_default_library"] + deps,
                     protoc = "@com_google_protobuf//:protoc",
                     visibility = ["//visibility:public"],
                     validate = 1,
                     **kwargs)

def pgv_cc_proto_library(name, srcs = None, deps = [], **kwargs):
    cc_proto_library(name,
                     srcs,
                     protoc = "@com_google_protobuf//:protoc",
                     default_runtime = "@com_google_protobuf//:protobuf",
                     deps = ["//validate:validate_cc"] + deps,
                     visibility = ["//visibility:public"],
                     validate = 1,
                     **kwargs)
