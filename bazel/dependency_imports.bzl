load("@bazel_features//:deps.bzl", "bazel_features_deps")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@rules_java//java:rules_java_deps.bzl", "rules_java_dependencies")
load("@rules_jvm_external//:repositories.bzl", "rules_jvm_external_deps")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies")
load("@rules_proto//proto:toolchains.bzl", "rules_proto_toolchains")
load("@rules_python//python:pip.bzl", "pip_parse")

def _pgv_pip_dependencies():
    # This rule translates the specified requirements.in (which must be same as install_requires from setup.cfg)
    # into @pgv_pip_deps//:requirements.bzl.
    pip_parse(
        name = "pgv_pip_deps",
        requirements_lock = "@com_envoyproxy_protoc_gen_validate//python:requirements.txt",
    )

def _pgv_go_dependencies():
    go_rules_dependencies()
    go_register_toolchains(
        version = "1.21.1",
    )
    gazelle_dependencies()

def pgv_dependency_imports():
    bazel_skylib_workspace()
    bazel_features_deps()

    rules_java_dependencies()

    rules_jvm_external_deps()

    # Import @com_google_protobuf's dependencies.
    protobuf_deps()

    # Import @pgv_pip_deps defined by python/requirements.in.
    _pgv_pip_dependencies()

    # Import rules for the Go compiler.
    _pgv_go_dependencies()

    # Setup rules_proto.
    rules_proto_dependencies()
    rules_proto_toolchains()
