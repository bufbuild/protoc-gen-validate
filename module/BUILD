load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "module",
    srcs = [
        "checker.go",
        "validate.go",
    ],
    importpath = "github.com/envoyproxy/protoc-gen-validate/module",
    visibility = ["//visibility:public"],
    deps = [
        "//templates",
        "//templates/java",
        "//validate",
        "@com_github_lyft_protoc_gen_star//:protoc-gen-star",
        "@com_github_lyft_protoc_gen_star//lang/go",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

alias(
    name = "go_default_library",
    actual = ":module",
    visibility = ["//visibility:public"],
)
