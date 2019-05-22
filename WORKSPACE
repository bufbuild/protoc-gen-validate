workspace(name = "com_envoyproxy_protoc_gen_validate")

load('@bazel_tools//tools/build_defs/repo:http.bzl', 'http_archive')

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
)
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
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

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

bind(
    name = "six",
    actual = "@six_archive//:six",
)

http_archive(
    name = "six_archive",
    build_file = "@com_google_protobuf//:six.BUILD",
    sha256 = "105f8d68616f8248e24bf0e9372ef04d3cc10104f1980f54d57b2ce73a5ad56a",
    url = "https://pypi.python.org/packages/source/s/six/six-1.10.0.tar.gz#md5=34eed507548117b2ab523ab14b2f8b55",
)

maven_jar(
    name="com_google_re2j",
    artifact = "com.google.re2j:re2j:1.2",
)

maven_jar(
    name = "com_google_guava",
    artifact="com.google.guava:guava:27.0-jre",
)
bind(
    name = "guava",
    actual="@com_google_guava//jar",
)

maven_jar(
    name = "com_google_gson",
    artifact = "com.google.code.gson:gson:2.8.5"
)
bind(
    name = "gson",
    actual = "@com_google_gson//jar",
)

maven_jar(
    name = "org_apache_commons_validator",
    artifact = "commons-validator:commons-validator:1.6"
)
