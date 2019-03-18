py_library(
    name = "python-harness-proto",
    srcs = ["harness_pb2.py"],
    deps = ["@com_google_protobuf//:protobuf_python"],
)

py_binary(
    name = "python-harness",
    srcs = ["harness.py"],
    deps = ["python-harness-proto", ":pgv_python_proto_library"],
    main = "harness.py",
    visibility = ["//tests/harness:__subpackages__"],
)