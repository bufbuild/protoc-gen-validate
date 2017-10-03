git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "2e319588571f20fdaaf83058b690abd32f596e89",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("//bazel:go_proto_library.bzl", "go_proto_repositories")

go_proto_repositories()

git_repository(
    name = "protobuf_bzl",
    # v3.4.0
    commit = "80a37e0782d2d702d52234b62dd4b9ec74fd2c95",
    remote = "https://github.com/google/protobuf.git",
)

go_rules_dependencies()
go_register_toolchains()
