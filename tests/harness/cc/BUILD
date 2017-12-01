cc_binary(
    name = "cc-harness",
    srcs = ["harness.cc"],
    visibility = ["//visibility:public"],
    # These ensure that we are at least compatible with what Envoy is expecting.
    copts = [
        "-Wall",
        "-Wextra",
        "-Werror",
        "-Wnon-virtual-dtor",
        "-Woverloaded-virtual",
        "-Wold-style-cast",
        "-std=c++14",
    ],
    deps = [
        "//tests/harness:harness_proto",
        "//tests/harness/cases:cc",
    ],
)
