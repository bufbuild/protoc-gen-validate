load("@rules_python//python:defs.bzl", "py_binary")
load("@pgv_pip_deps//:requirements.bzl", "requirement")

# Adding `requirement("protobuf")` first ensures
# we test with the protobuf runtime installed by pip_install (which reads requirements.in).
# Because the other dependencies are created with py_proto_library, they inevitably come with their own runtime.
# Note that passing `default_runtime=""` to py_proto_library() won't solve that problem
# given that we also include @com_google_protobuf//:protobuf_python for Well-Known type support.
#
# This is a bit clunky and will hopefully be resolved if ever there's a canonical Bazel rule for python generation.
# See this blog for context: https://thethoughtfulkoala.com/posts/2020/05/08/py-protobuf-bazel.html
py_binary(
    name = "python-harness",
    srcs = ["harness.py"],
    main = "harness.py",
    visibility = ["//visibility:public"],
    deps = [
        requirement("protobuf"),
        "//tests/harness:harness_py_proto",
        "//tests/harness/cases:cases_py_proto",
        "//python:validator_py"
    ]
)
