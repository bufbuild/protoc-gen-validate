load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("@rules_python//python:pip.bzl", "pip_install")

def _pgv_pip_dependencies():
    # This rule translates the specified requirements.in into
    # @pgv_pip_deps//:requirements.bzl.
    # `make python/requirements.generated` must be executed before bazel
    pip_install(
        name = "pgv_pip_deps",
        requirements = "//python:requirements.generated",
    )

def _pgv_go_dependencies():
    go_rules_dependencies()
    go_register_toolchains(
        version = "1.15.6",
    )
    gazelle_dependencies()

def pgv_dependency_imports():
    # Import @com_google_protobuf's dependencies.
    protobuf_deps()

    # Import @pgv_pip_deps defined by pip's requirements.txt.
    _pgv_pip_dependencies()

    # Import rules for the Go compiler.
    _pgv_go_dependencies()

    # Setup rules_proto.
    rules_proto_dependencies()
    rules_proto_toolchains()
