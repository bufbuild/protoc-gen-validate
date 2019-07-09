load("@my_deps//:requirements.bzl", "requirement")
load("@io_bazel_rules_python//python:python.bzl",  "py_binary")


py_binary(
    name = "python-harness",
    srcs = ["harness.py"],
    deps = [
        "//tests/harness:python-harness-proto",
        "//tests/harness/cases:python",
           requirement("validate-email"),
           requirement("ipaddress")],
    main = "harness.py",
    visibility = ["//tests/harness:__subpackages__"],
    python_version = "PY2",
)
