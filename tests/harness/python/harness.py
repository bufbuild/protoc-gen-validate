import sys
import inspect

from tests.harness.harness_pb2 import TestCase, TestResult
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

sys.path.append('/go/src/github.com/envoyproxy/protoc-gen-validate/validate')
from validator import *

class_list = []
for k, v in inspect.getmembers(sys.modules[__name__], inspect.isclass):
    if 'DESCRIPTOR' in dir(v):
        class_list.append(v)

def unpack(message):
    for cls in class_list:
        if message.Is(cls.DESCRIPTOR):
            test_class = cls()
            message.Unpack(test_class)
            return test_class

if __name__ == "__main__":
    message = sys.stdin.read()
    testcase = TestCase()
    testcase.ParseFromString(message)
    test_class = unpack(testcase.message)
    try:
        result = TestResult()
        validate = generate_validate(test_class)
        valid = validate(test_class)
        result.Valid = True
    except ValidationFailed as e:
        result.Valid = False
    except UnimplementedException as e:
        result.Error = False
        result.AllowFailure = True
    sys.stdout.write(result.SerializeToString())
