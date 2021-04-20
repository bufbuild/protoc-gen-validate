load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "checker.go",
        "validate.go",
    ],
    importpath = "github.com/envoyproxy/protoc-gen-validate/module",
    visibility = ["//visibility:public"],
    deps = [
        "//templates:go_default_library",
        "//templates/java:go_default_library",
        "//validate:go_default_library",  # keep
        "@com_github_lyft_protoc_gen_star//:go_default_library",
        "@com_github_lyft_protoc_gen_star//lang/go:go_default_library",
        "@org_golang_google_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//types/known/durationpb:go_default_library",
        "@org_golang_google_protobuf//types/known/timestamppb:go_default_library",
    ],
)
