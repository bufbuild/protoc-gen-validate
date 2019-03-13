load("@com_google_protobuf//:protobuf.bzl", "py_proto_library")

py_proto_library(
    name = "python-proto",
    protos = ["//tests/harness/harness.proto"],
)

py_binary(
    name = "python-harness",
    srcs = ["harness.py"],
    main = "harness.py",
    visibility = ["//tests/harness:__subpackages__"],
    deps=[
        "python-proto",
    ],
)
