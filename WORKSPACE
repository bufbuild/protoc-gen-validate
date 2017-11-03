git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "2e319588571f20fdaaf83058b690abd32f596e89",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("//bazel:go_proto_library.bzl", "go_proto_repositories")

go_proto_repositories()

# TODO(htuch): This can switch back to a point release http_archive at the next
# release (> 3.4.1), we need HEAD proto_library support and
# https://github.com/google/protobuf/pull/3761.
http_archive(
    name = "com_google_protobuf",
    strip_prefix = "protobuf-c4f59dcc5c13debc572154c8f636b8a9361aacde",
    sha256 = "5d4551193416861cb81c3bc0a428f22a6878148c57c31fb6f8f2aa4cf27ff635",
    url = "https://github.com/google/protobuf/archive/c4f59dcc5c13debc572154c8f636b8a9361aacde.tar.gz",
)

go_rules_dependencies()
go_register_toolchains()
