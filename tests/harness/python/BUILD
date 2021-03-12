load("@rules_python//python:defs.bzl", "py_binary")
load("@pgv_pip_deps//:requirements.bzl", "requirement")

py_binary(
    name = "python-harness",
    srcs = ["harness.py"],
    main = "harness.py",
    srcs_version = "PY3",
    visibility = ["//visibility:public"],
    deps = [
        # ensures we test with the package's own protobuf runtime specified in setup.cfg
        # and not the one riding on the py_proto_library dependencies
        requirement("protobuf"),
        "//tests/harness:harness_py_proto",
        "//tests/harness/cases:cases_py_proto",
        "//python:validator_py"
    ]
)
