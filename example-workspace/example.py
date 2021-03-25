import os
import sys

from google.protobuf import text_format
from protoc_gen_validate import validator
from foo import bar_pb2

EX_FAILURE = 1


def main(filenames):
    if not filenames:
        print("No inputs provided; exiting")
        return os.EX_OK

    success_count = 0
    for filename in filenames:
        bars = bar_pb2.Bars()
        try:
            with open(filename, 'r') as fh:
                text_format.Parse(fh.read(), bars)
        except IOError as error:
            print(
                "Failed to open file '{filename}': {error}".format(
                    filename=filename,
                    error=error,
                ),
                file=sys.stderr,
            )
            return EX_FAILURE
        except text_format.ParseError as error:
            print(dir(bars))
            print(
                "Failed to parse file '{filename}'as a {fullname} textproto: {error}".format(
                    filename=filename,
                    fullname=bars.DESCRIPTOR.name,
                    error=error,
                ),
                file=sys.stderr,
            )
            return EX_FAILURE

        try:
            validator.validate(bars)
        except validator.ValidationFailed as error:
            print(
                "Failed to validate file '{filename}'as a {fullname} textproto: {error}".format(
                    filename=filename,
                    fullname=bars.DESCRIPTOR.name,
                    error=error,
                ),
                file=sys.stderr,
            )
        else:
            print(
                "Successfully validated file '{filename}'as a {fullname} textproto".format(
                    filename=filename,
                    fullname=bars.DESCRIPTOR.name,
                )
            )
            success_count += 1

    failure_count = len(filenames) - success_count
    if failure_count:
        print(
            "Failed to validate {count} file{s}".format(
                count=failure_count,
                s=("s" if failure_count > 1 else ""),
            ),
            file=sys.stderr,
        )
        return EX_FAILURE

    return os.EX_OK


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
