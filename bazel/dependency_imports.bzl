load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("@rules_python//python:pip.bzl", "pip_install")

def _pgv_pip_dependencies():
    # This rule translates the specified requirements.in (which must be same as install_requires from setup.cfg)
    # into @pgv_pip_deps//:requirements.bzl.
    pip_install(
        name = "pgv_pip_deps",
        requirements = "@com_envoyproxy_protoc_gen_validate//python:requirements.in",
    )

def _pgv_go_dependencies():
    go_rules_dependencies()
    go_register_toolchains(
        version = "1.19.7",
    )
    gazelle_dependencies()

def pgv_dependency_imports():
    # Import @com_google_protobuf's dependencies.
    protobuf_deps()

    # Import @pgv_pip_deps defined by python/requirements.in.
    _pgv_pip_dependencies()

    # Import rules for the Go compiler.
    _pgv_go_dependencies()

    # Setup rules_proto.
    rules_proto_dependencies()
    rules_proto_toolchains()
