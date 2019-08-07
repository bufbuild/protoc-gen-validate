load("@pgv_pip_deps//:requirements.bzl", "requirement")
load("@io_bazel_rules_python//python:python.bzl",  "py_binary")


py_binary(
    name = "python-harness",
    srcs = ["harness.py"],
    deps = [
        "//tests/harness:python-harness-proto",
        "//tests/harness/cases:python",
       requirement("validate-email"),
       requirement("ipaddress"),
       requirement("Jinja2"),
       requirement("MarkupSafe"),],
    main = "harness.py",
    visibility = ["//tests/harness:__subpackages__"],
#   Python Version is set by default to PY2AND3.
)
