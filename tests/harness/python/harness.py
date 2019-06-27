import sys
import inspect

from tests.harness.harness_pb2 import TestCase, TestResult

from tests.harness.cases.bool_pb_validate import *
from tests.harness.cases.bytes_pb_validate import *
from tests.harness.cases.enums_pb_validate import *
from tests.harness.cases.messages_pb_validate import *
from tests.harness.cases.numbers_pb_validate import *
from tests.harness.cases.oneofs_pb_validate import *
from tests.harness.cases.repeated_pb_validate import *
from tests.harness.cases.strings_pb_validate import *
from tests.harness.cases.maps_pb_validate import *
from tests.harness.cases.wkt_any_pb_validate import *
from tests.harness.cases.wkt_duration_pb_validate import *
from tests.harness.cases.wkt_wrappers_pb_validate import *
from tests.harness.cases.wkt_timestamp_pb_validate import *

from tests.harness.cases.bool_pb2 import *
from tests.harness.cases.bytes_pb2 import *
from tests.harness.cases.enums_pb2 import *
from tests.harness.cases.messages_pb2 import *
from tests.harness.cases.numbers_pb2 import *
from tests.harness.cases.oneofs_pb2 import *
from tests.harness.cases.repeated_pb2 import *
from tests.harness.cases.strings_pb2 import *
from tests.harness.cases.maps_pb2 import *
from tests.harness.cases.wkt_any_pb2 import *
from tests.harness.cases.wkt_duration_pb2 import *
from tests.harness.cases.wkt_wrappers_pb2 import *
from tests.harness.cases.wkt_timestamp_pb2 import *

class_list = {}
for k, v in inspect.getmembers(sys.modules[__name__], inspect.isclass):
    for m, n in inspect.getmembers(sys.modules[__name__], inspect.isfunction):
        if m == "validate_" + k:
            class_list[v] = n

def unpack(message):
    for cls, validate in class_list.items():
        if message.Is(cls.DESCRIPTOR):
            test_class = cls()
            message.Unpack(test_class)
            return test_class, validate

if __name__ == "__main__":
    message = sys.stdin.read()
    testcase = TestCase()
    testcase.ParseFromString(message)
    test_class, validate = unpack(testcase.message)
    try:
        result = TestResult()
        result.Valid, _ = validate(test_class)
    except Exception as e:
        if e.name == "UnimplementedException":
            result.Error = False
            result.AllowFailure = True
    sys.stdout.write(result.SerializeToString())
