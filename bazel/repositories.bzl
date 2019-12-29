load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:jvm.bzl", "jvm_maven_import_external")

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
        jvm_maven_import_external(
            name = "com_google_re2j",
            artifact = "com.google.re2j:re2j:1.2",
	    artifact_sha256 = "e9dc705fd4c570344b54a7146b2e3a819cdc271a29793f4acc1a93b56a388e59",
            server_urls = ["http://central.maven.org/maven2"],
        )

    if not native.existing_rule("com_googlesource_code_re2"):
        http_archive(
            name = "com_googlesource_code_re2",
            sha256 = "38bc0426ee15b5ed67957017fd18201965df0721327be13f60496f2b356e3e01",
            strip_prefix = "re2-2019-08-01",
            urls = ["https://github.com/google/re2/archive/2019-08-01.tar.gz"],
        )

    if not native.existing_rule("com_google_guava"):
        jvm_maven_import_external(
            name = "com_google_guava",
            artifact = "com.google.guava:guava:27.0-jre",
	    artifact_sha256 = "63b09db6861011e7fb2481be7790c7fd4b03f0bb884b3de2ecba8823ad19bf3f",
            server_urls = ["http://central.maven.org/maven2"],
        )

    if not native.existing_rule("guava"):
        native.bind(
            name = "guava",
            actual = "@com_google_guava//jar",
        )

    if not native.existing_rule("com_google_gson"):
        jvm_maven_import_external(
            name = "com_google_gson",
            artifact = "com.google.code.gson:gson:2.8.5",
	    artifact_sha256 = "233a0149fc365c9f6edbd683cfe266b19bdc773be98eabdaf6b3c924b48e7d81",
            server_urls = ["http://central.maven.org/maven2"],
        )

    if not native.existing_rule("gson"):
        native.bind(
            name = "gson",
            actual = "@com_google_gson//jar",
        )

    if not native.existing_rule("error_prone_annotations_maven"):
        jvm_maven_import_external(
            name = "error_prone_annotations_maven",
            artifact = "com.google.errorprone:error_prone_annotations:2.3.2",
	    artifact_sha256 = "357cd6cfb067c969226c442451502aee13800a24e950fdfde77bcdb4565a668d",
            server_urls = ["http://central.maven.org/maven2"],
        )

    if not native.existing_rule("error_prone_annotations"):
        native.bind(
            name = "error_prone_annotations",
            actual = "@error_prone_annotations_maven//jar",
        )

    if not native.existing_rule("org_apache_commons_validator"):
        jvm_maven_import_external(
            name = "org_apache_commons_validator",
            artifact = "commons-validator:commons-validator:1.6",
	    artifact_sha256 = "bd62795d7068a69cbea333f6dbf9c9c1a6ad7521443fb57202a44874f240ba25",
            server_urls = ["http://central.maven.org/maven2"],
        )

    if not native.existing_rule("io_bazel_rules_python"):
        git_repository(
            name = "io_bazel_rules_python",
            remote = "https://github.com/bazelbuild/rules_python.git",
            commit = "fdbb17a4118a1728d19e638a5291b4c4266ea5b8",
            shallow_since = "1557865590 -0400",
        )

    if not native.existing_rule("rules_proto"):
        git_repository(
            name = "rules_proto",
            remote = "https://github.com/bazelbuild/rules_proto.git",
            commit = "2c0468366367d7ed97a1f702f9cd7155ab3f73c5",
            shallow_since = "1575470667 +0100"
        )
