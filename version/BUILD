load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "version",
    srcs = [
        "version.go",
    ],
    importpath = "github.com/envoyproxy/protoc-gen-validate/version",
    visibility = ["//visibility:public"],
)

alias(
    name = "go_default_library",
    actual = ":version",
    deprecation = "Use :version instead of :go_default_library.  Details about the new naming convention: https://github.com/bazelbuild/bazel-gazelle/pull/863",
    visibility = ["//visibility:public"],
)

