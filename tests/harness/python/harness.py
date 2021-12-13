import sys
import inspect

from python.protoc_gen_validate.validator import validate, ValidationFailed

from tests.harness.harness_pb2 import TestCase, TestResult
from tests.harness.cases.bool_pb2 import *
from tests.harness.cases.bytes_pb2 import *
from tests.harness.cases.enums_pb2 import *
from tests.harness.cases.enums_pb2 import *
from tests.harness.cases.messages_pb2 import *
from tests.harness.cases.numbers_pb2 import *
from tests.harness.cases.oneofs_pb2 import *
from tests.harness.cases.repeated_pb2 import *
from tests.harness.cases.strings_pb2 import *
from tests.harness.cases.maps_pb2 import *
from tests.harness.cases.wkt_any_pb2 import *
from tests.harness.cases.wkt_duration_pb2 import *
from tests.harness.cases.wkt_nested_pb2 import *
from tests.harness.cases.wkt_wrappers_pb2 import *
from tests.harness.cases.wkt_timestamp_pb2 import *
from tests.harness.cases.kitchen_sink_pb2 import *


message_classes = {}
for k, v in inspect.getmembers(sys.modules[__name__], inspect.isclass):
    if 'DESCRIPTOR' in dir(v):
        message_classes[v.DESCRIPTOR.full_name] = v


if __name__ == "__main__":
    read = sys.stdin.buffer.read()

    testcase = TestCase()
    testcase.ParseFromString(read)

    test_class = message_classes[testcase.message.TypeName()]
    test_msg = test_class()
    testcase.message.Unpack(test_msg)

    try:
        result = TestResult()
        valid = validate(test_msg)
        result.Valid = True
    except ValidationFailed as e:
        result.Valid = False
        result.Reasons[:] = [repr(e)]

    sys.stdout = open(sys.stdout.fileno(), mode='w', encoding='utf8')
    sys.stdout.write(result.SerializeToString().decode("utf-8"))
