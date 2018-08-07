load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@io_bazel_rules_go//proto:compiler.bzl", "go_proto_compiler")
load(":protobuf.bzl", "cc_proto_gen_validate")

def pgv_go_proto_library(name, proto = None, deps = [], **kwargs):
    go_proto_compiler(
        name = "pgv_plugin",
        suffix = ".pb.validate.go",
        valid_archive = False,
        plugin = "//:protoc-gen-validate",
        options = ["lang=go"],
    )

    go_proto_library(name = name,
                     proto = proto,
                     deps = ["//validate:go_default_library"] + deps,
                     compilers = ["@io_bazel_rules_go//proto:go_proto", "pgv_plugin"],
                     visibility = ["//visibility:public"],
                     **kwargs)

def pgv_cc_proto_library(
        name,
        deps=[],
        cc_deps=[],
        **kargs):
    """Bazel rule to create a C++ protobuf validation library from proto source files
    Args:
      name: the name of the pgv_cc_proto_library.
      deps: proto_library rules that contains the necessary .proto files. Note that this
             must include @com_lyft_protoc_gen_validate//validate:validate_proto
      cc_deps: C++ dependencies of the protos being compiled. Likely cc_proto_library or pgv_cc_proto_library
      **kargs: other keyword arguments that are passed to cc_library.
    """

    cc_proto_gen_validate(
        name=name+"_validate",
        deps=deps,
    )

    native.cc_library(
        name=name,
        hdrs=[":"+name+"_validate"],
        srcs=[":"+name+"_validate"],
        deps= cc_deps + [
            "@com_lyft_protoc_gen_validate//validate:cc_validate",
            "@com_lyft_protoc_gen_validate//validate:validate_cc",
            "@com_google_protobuf//:protobuf",
        ],
        alwayslink=1,
        **kargs)

