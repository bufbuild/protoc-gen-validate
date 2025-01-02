load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:jvm.bzl", "jvm_maven_import_external")

_DEFAULT_REPOSITORIES = ["https://repo.maven.apache.org/maven2"]

def pgv_dependencies(maven_repos = _DEFAULT_REPOSITORIES):
    if not native.existing_rule("io_bazel_rules_go"):
        http_archive(
            name = "io_bazel_rules_go",
            sha256 = "33acc4ae0f70502db4b893c9fc1dd7a9bf998c23e7ff2c4517741d4049a976f8",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.48.0/rules_go-v0.48.0.zip",
                "https://github.com/bazelbuild/rules_go/releases/download/v0.48.0/rules_go-v0.48.0.zip",
            ],
        )

    if not native.existing_rule("rules_python"):
        http_archive(
            name = "rules_python",
            sha256 = "4912ced70dc1a2a8e4b86cec233b192ca053e82bc72d877b98e126156e8f228d",
            strip_prefix = "rules_python-0.32.2",
            url = "https://github.com/bazelbuild/rules_python/releases/download/0.32.2/rules_python-0.32.2.tar.gz",
        )

    if not native.existing_rule("rules_cc"):
        http_archive(
            name = "rules_cc",
            sha256 = "4b12149a041ddfb8306a8fd0e904e39d673552ce82e4296e96fac9cbf0780e59",
            strip_prefix = "rules_cc-0.1.0",
            url = "https://github.com/bazelbuild/rules_cc/releases/download/0.1.0/rules_cc-0.1.0.tar.gz",
        )

    if not native.existing_rule("bazel_features"):
        http_archive(
            name = "bazel_features",
            sha256 = "8b1c9b7558498000f5adebbc584b7bf15b6b2bf181448a66f6b2fc5b4c84231c",
            strip_prefix = "bazel_features-1.23.0",
            url = "https://github.com/bazel-contrib/bazel_features/releases/download/v1.23.0/bazel_features-v1.23.0.tar.gz",
        )

    if not native.existing_rule("rules_license"):
        http_archive(
            name = "rules_license",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_license/releases/download/1.0.0/rules_license-1.0.0.tar.gz",
                "https://github.com/bazelbuild/rules_license/releases/download/1.0.0/rules_license-1.0.0.tar.gz",
            ],
            sha256 = "26d4021f6898e23b82ef953078389dd49ac2b5618ac564ade4ef87cced147b38",
        )

    if not native.existing_rule("rules_jvm_external"):
        RULES_JVM_EXTERNAL_TAG = "6.6"
        RULES_JVM_EXTERNAL_SHA = "3afe5195069bd379373528899c03a3072f568d33bd96fe037bd43b1f590535e7"
        http_archive(
            name = "rules_jvm_external",
            strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
            sha256 = RULES_JVM_EXTERNAL_SHA,
            url = "https://github.com/bazel-contrib/rules_jvm_external/releases/download/%s/rules_jvm_external-%s.tar.gz" % (RULES_JVM_EXTERNAL_TAG, RULES_JVM_EXTERNAL_TAG),
        )

    if not native.existing_rule("rules_java"):
        http_archive(
            name = "rules_java",
            urls = [
                "https://github.com/bazelbuild/rules_java/releases/download/8.6.3/rules_java-8.6.3.tar.gz",
            ],
            sha256 = "6d8c6d5cd86fed031ee48424f238fa35f33abc9921fd97dd4ae1119a29fc807f",
        )

    if not native.existing_rule("bazel_gazelle"):
        http_archive(
            name = "bazel_gazelle",
            sha256 = "d76bf7a60fd8b050444090dfa2837a4eaf9829e1165618ee35dceca5cbdf58d5",
            url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.37.0/bazel-gazelle-v0.37.0.tar.gz",
        )

    if not native.existing_rule("com_google_protobuf"):
        http_archive(
            name = "com_google_protobuf",
            url = "https://github.com/protocolbuffers/protobuf/releases/download/v29.2/protobuf-29.2.tar.gz",
            sha256 = "63150aba23f7a90fd7d87bdf514e459dd5fe7023fdde01b56ac53335df64d4bd",
            strip_prefix = "protobuf-29.2",
        )

    # NOTE: Even though this is not directly used, rules_go requires this still.
    if not native.existing_rule("rules_proto"):
        http_archive(
            name = "rules_proto",
            sha256 = "14a225870ab4e91869652cfd69ef2028277fc1dc4910d65d353b62d6e0ae21f4",
            strip_prefix = "rules_proto-7.1.0",
            url = "https://github.com/bazelbuild/rules_proto/releases/download/7.1.0/rules_proto-7.1.0.tar.gz",
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
            sha256 = "bc283cdfcd526a52c3201279cda4bc298652efa898b10b4db0837dc51652756f",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.7.1/bazel-skylib-1.7.1.tar.gz",
                "https://github.com/bazelbuild/bazel-skylib/releases/download/1.7.1/bazel-skylib-1.7.1.tar.gz",
            ],
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

    if not native.existing_rule("org_apache_commons_validator"):
        jvm_maven_import_external(
            name = "org_apache_commons_validator",
            artifact = "commons-validator:commons-validator:1.6",
            artifact_sha256 = "bd62795d7068a69cbea333f6dbf9c9c1a6ad7521443fb57202a44874f240ba25",
            server_urls = maven_repos,
        )
