load("@rules_python//python:defs.bzl", "py_library")
load("@pgv_pip_deps//:requirements.bzl", "all_requirements")

py_library(
    name = "validator_py",
    srcs = glob(["**/*.py"]),
    deps = all_requirements,
    visibility = ["//visibility:public"],
)
