load("@rules_python//python:defs.bzl", "py_library")
load("@pgv_pip_deps//:requirements.bzl", "all_requirements")

exports_files([
    "requirements.in",
    "setup.cfg",
])

py_library(
    name = "validator_py",
    srcs = glob(["**/*.py"]),
    visibility = ["//visibility:public"],
    deps = all_requirements,
)
