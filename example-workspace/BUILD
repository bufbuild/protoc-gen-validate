# Example C++ binary that uses the generated validation code.
#
# This binary attempts to read files named on the command line as binary protos.
# To try it out, first use Bazel to build this binary:
#
#   bazel build //:example_cc
#
# Now run it on some files. From the repository root directory:
#
#   bazel run //:example_cc -- $(pwd)/example.cc
#
# The binary will fail because example.cc is not a valid textproto. Luckily
# this directory contains two textprotos already: valid.textproto and
# invalid.textproto. From the root directory:
#
#   bazel run //:example_cc -- $(pwd)/valid.textproto
#
# which succeeds, or
#
#   bazel run //:example_cc -- $(pwd)/invalid.textproto
#
# which fails.
cc_binary(
    name = "example_cc",
    srcs = ["example.cc"],
    deps = ["//foo:bar_cc_proto"],
)

# Test that the example textproto inputs evoke the right responses.
[
    sh_test(
        name = "example_cc_test_" + which,
        srcs = ["example_test.sh"],
        args = [
            "$(location :example_cc)",
            str(code),
            "$(location :" + which + ")",
        ],
        data = [
            which,
            ":example_cc",
        ],
    )
    for (which, code) in (
        ("valid.textproto", 0),
        ("invalid.textproto", 1),
    )
]
