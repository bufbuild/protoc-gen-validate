workspace(name = "com_lyft_protoc_gen_validate")

load('@bazel_tools//tools/build_defs/repo:http.bzl', 'http_archive')

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.14.0/rules_go-0.14.0.tar.gz",
    sha256 = "5756a4ad75b3703eb68249d50e23f5d64eaf1593e886b9aa931aa6e938c4e301",
)
http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz",
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)
# TODO: use released version of protobuf that includes commit
# fa252ec2a54acb24ddc87d48fed1ecfd458445fd. This works around the issue
# described here: https://github.com/google/protobuf/pull/5024
http_archive(
    name = "com_google_protobuf",
    url = "https://github.com/google/protobuf/archive/fa252ec2a54acb24ddc87d48fed1ecfd458445fd.tar.gz",
    sha256 = "3d610ac90f8fa16e12490088605c248b85fdaf23114ce4b3605cdf81f7823604",
    strip_prefix = "protobuf-fa252ec2a54acb24ddc87d48fed1ecfd458445fd",
)
http_archive(
    name = "bazel_skylib",
    url = "https://github.com/bazelbuild/bazel-skylib/archive/0.5.0.tar.gz",
    sha256 = "b5f6abe419da897b7901f90cbab08af958b97a8f3575b0d3dd062ac7ce78541f",
    strip_prefix = "bazel-skylib-0.5.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

bind(
    name = "six",
    actual = "@six_archive//:six",
)

new_http_archive(
    name = "six_archive",
    build_file = "@com_google_protobuf//:six.BUILD",
    sha256 = "105f8d68616f8248e24bf0e9372ef04d3cc10104f1980f54d57b2ce73a5ad56a",
    url = "https://pypi.python.org/packages/source/s/six/six-1.10.0.tar.gz#md5=34eed507548117b2ab523ab14b2f8b55",
)
