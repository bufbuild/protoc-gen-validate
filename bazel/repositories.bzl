load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:jvm.bzl", "jvm_maven_import_external")

_DEFAULT_REPOSITORIES = ["https://repo.maven.apache.org/maven2"]

def pgv_dependencies(maven_repos = _DEFAULT_REPOSITORIES):
    if not native.existing_rule("io_bazel_rules_go"):
        http_archive(
            name = "io_bazel_rules_go",
            sha256 = "91585017debb61982f7054c9688857a2ad1fd823fc3f9cb05048b0025c47d023",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
                "https://github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
            ],
        )

    if not native.existing_rule("bazel_gazelle"):
        http_archive(
            name = "bazel_gazelle",
            sha256 = "d3fa66a39028e97d76f9e2db8f1b0c11c099e8e01bf363a923074784e451f809",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
                "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
            ],
        )

    if not native.existing_rule("com_google_protobuf"):
        http_archive(
            name = "com_google_protobuf",
            url = "https://github.com/protocolbuffers/protobuf/archive/v3.15.3.tar.gz",
            sha256 = "b10bf4e2d1a7586f54e64a5d9e7837e5188fc75ae69e36f215eb01def4f9721b",
            strip_prefix = "protobuf-3.15.3",
        )

    # TODO(akonradi): This shouldn't be necessary since the same http_archive block is imported by
    # protobuf_deps from @com_google_protobuf. Investigate why.
    if not native.existing_rule("zlib"):
        http_archive(
            name = "zlib",
            build_file = "@com_google_protobuf//:third_party/zlib.BUILD",
            sha256 = "b3a24de97a8fdbc835b9833169501030b8977031bcb54b3b3ac13740f846ab30",
            strip_prefix = "zlib-1.2.13",
            urls = ["https://zlib.net/fossils/zlib-1.2.13.tar.gz"],
        )

    if not native.existing_rule("bazel_skylib"):
        http_archive(
            name = "bazel_skylib",
            sha256 = "66ffd9315665bfaafc96b52278f57c7e2dd09f5ede279ea6d39b2be471e7e3aa",
            urls = ["https://github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz"],
        )

    if not native.existing_rule("six"):
        http_archive(
            name = "six",
            build_file = "@com_google_protobuf//:third_party/six.BUILD",
            sha256 = "d16a0141ec1a18405cd4ce8b4613101da75da0e9a7aec5bdd4fa804d0e0eba73",
            urls = ["https://pypi.python.org/packages/source/s/six/six-1.12.0.tar.gz"],
        )

    if not native.existing_rule("com_google_re2j"):
        jvm_maven_import_external(
            name = "com_google_re2j",
            artifact = "com.google.re2j:re2j:1.2",
            artifact_sha256 = "e9dc705fd4c570344b54a7146b2e3a819cdc271a29793f4acc1a93b56a388e59",
            server_urls = maven_repos,
        )

    if not native.existing_rule("com_googlesource_code_re2"):
        http_archive(
            name = "com_googlesource_code_re2",
            sha256 = "2e9489a31ae007c81e90e8ec8a15d62d58a9c18d4fd1603f6441ef248556b41f",
            strip_prefix = "re2-2020-07-06",
            urls = ["https://github.com/google/re2/archive/2020-07-06.tar.gz"],
        )

    if not native.existing_rule("com_google_guava"):
        jvm_maven_import_external(
            name = "com_google_guava",
            artifact = "com.google.guava:guava:27.0-jre",
            artifact_sha256 = "63b09db6861011e7fb2481be7790c7fd4b03f0bb884b3de2ecba8823ad19bf3f",
            server_urls = maven_repos,
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
            server_urls = maven_repos,
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
            server_urls = maven_repos,
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
            server_urls = maven_repos,
        )

    if not native.existing_rule("rules_python"):
        http_archive(
            name = "rules_python",
            sha256 = "5868e73107a8e85d8f323806e60cad7283f34b32163ea6ff1020cf27abef6036",
            strip_prefix = "rules_python-0.25.0",
            url = "https://github.com/bazelbuild/rules_python/releases/download/0.25.0/rules_python-0.25.0.tar.gz",
        )

    if not native.existing_rule("rules_proto"):
        http_archive(
            name = "rules_proto",
            sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
            strip_prefix = "rules_proto-5.3.0-21.7",
            urls = [
                "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
            ],
        )
