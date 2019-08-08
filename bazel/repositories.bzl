load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

def pgv_dependencies():
    if not native.existing_rule("io_bazel_rules_go"):
        http_archive(
            name = "io_bazel_rules_go",
            urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
            sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
        )

    if not native.existing_rule("bazel_gazelle"):
        http_archive(
            name = "bazel_gazelle",
            urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
            sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
        )

    if not native.existing_rule("com_google_protobuf"):
        http_archive(
            name = "com_google_protobuf",
            url = "https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protobuf-all-3.9.1.tar.gz",
            sha256 = "3040a5b946d9df7aa89c0bf6981330bf92b7844fd90e71b61da0c721e421a421",
            strip_prefix = "protobuf-3.9.1",
        )

    # TODO(akonradi): This shouldn't be necesary since the same http_archive block is imported by
    # protobuf_deps from @com_google_protobuf. Investigate why.
    if not native.existing_rule("zlib"):
        http_archive(
            name = "zlib",
            build_file = "@com_google_protobuf//:third_party/zlib.BUILD",
            sha256 = "c3e5e9fdd5004dcb542feda5ee4f0ff0744628baf8ed2dd5d66f8ca1197cb1a1",
            strip_prefix = "zlib-1.2.11",
            urls = ["https://zlib.net/zlib-1.2.11.tar.gz"],
        )

    if not native.existing_rule("bazel_skylib"):
        http_archive(
            name = "bazel_skylib",
            url = "https://github.com/bazelbuild/bazel-skylib/archive/0.5.0.tar.gz",
            sha256 = "b5f6abe419da897b7901f90cbab08af958b97a8f3575b0d3dd062ac7ce78541f",
            strip_prefix = "bazel-skylib-0.5.0",
        )

    if not native.existing_rule("six"):
        native.bind(
            name = "six",
            actual = "@six_archive//:six",
        )

    if not native.existing_rule("six_archive"):
        http_archive(
            name = "six_archive",
            build_file = "@com_google_protobuf//:six.BUILD",
            sha256 = "105f8d68616f8248e24bf0e9372ef04d3cc10104f1980f54d57b2ce73a5ad56a",
            url = "https://pypi.python.org/packages/source/s/six/six-1.10.0.tar.gz#md5=34eed507548117b2ab523ab14b2f8b55",
        )

    if not native.existing_rule("com_google_re2j"):
        native.maven_jar(
            name = "com_google_re2j",
            artifact = "com.google.re2j:re2j:1.2",
        )

    if not native.existing_rule("com_google_guava"):
        native.maven_jar(
            name = "com_google_guava",
            artifact = "com.google.guava:guava:27.0-jre",
        )

    if not native.existing_rule("guava"):
        native.bind(
            name = "guava",
            actual = "@com_google_guava//jar",
        )

    if not native.existing_rule("com_google_gson"):
        native.maven_jar(
            name = "com_google_gson",
            artifact = "com.google.code.gson:gson:2.8.5",
        )

    if not native.existing_rule("gson"):
        native.bind(
            name = "gson",
            actual = "@com_google_gson//jar",
        )

    if not native.existing_rule("error_prone_annotations_maven"):
        native.maven_jar(
            name = "error_prone_annotations_maven",
            artifact = "com.google.errorprone:error_prone_annotations:2.3.2",
        )

    if not native.existing_rule("error_prone_annotations"):
        native.bind(
            name = "error_prone_annotations",
            actual = "@error_prone_annotations_maven//jar",
        )

    if not native.existing_rule("org_apache_commons_validator"):
        native.maven_jar(
            name = "org_apache_commons_validator",
            artifact = "commons-validator:commons-validator:1.6",
        )

    if not native.existing_rule("io_bazel_rules_python"):
        git_repository(
            name = "io_bazel_rules_python",
            remote = "https://github.com/bazelbuild/rules_python.git",
            commit = "fdbb17a4118a1728d19e638a5291b4c4266ea5b8",
            shallow_since = "1557865590 -0400",
        )
