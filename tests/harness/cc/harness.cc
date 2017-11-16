#include <iostream>

#include "tests/harness/cases/bool.pb.h"
#include "tests/harness/cases/bool.pb.validate.h"
#include "tests/harness/cases/bytes.pb.h"
#include "tests/harness/cases/enums.pb.h"
#include "tests/harness/cases/maps.pb.h"
#include "tests/harness/cases/messages.pb.h"
#include "tests/harness/cases/numbers.pb.h"
#include "tests/harness/cases/oneofs.pb.h"
#include "tests/harness/cases/repeated.pb.h"
#include "tests/harness/cases/strings.pb.h"
#include "tests/harness/cases/wkt_any.pb.h"
#include "tests/harness/cases/wkt_duration.pb.h"
#include "tests/harness/cases/wkt_timestamp.pb.h"
#include "tests/harness/cases/wkt_wrappers.pb.h"

#include "tests/harness/harness.pb.h"

namespace {

using tests::harness::TestCase;
using tests::harness::TestResult;
using google::protobuf::Any;
using google::protobuf::Message;

std::ostream& operator<<(std::ostream& out, const TestResult& result) {
  out << "valid: " << result.valid() << " reason: '" << result.reason() << "'"
      << std::endl;
  return out;
}

void WriteTestResultAndExit(const TestResult& result) {
  if (!result.SerializeToOstream(&std::cout)) {
    std::cerr << "could not martial response: ";
    std::cerr << result << std::endl;
    exit(EXIT_FAILURE);
  }
  exit(EXIT_SUCCESS);
}

void ExitIfFailed(bool succeeded, const std::string& err_msg) {
  if (succeeded) {
    return;
  }

  TestResult result;
  result.set_error(true);
  result.set_reason(err_msg);
  WriteTestResultAndExit(result);
}

void ValidateOrExit(const std::function<bool(std::string*)>& validate_fn) {
  std::string error_msg;
  TestResult result;

  result.set_valid(validate_fn(&error_msg));
  result.set_reason(std::move(error_msg));
  WriteTestResultAndExit(result);
}

std::function<bool(std::string*)> GetValidationCheck(const Any& msg) {
#define TRY_RETURN_VALIDATE_CALLABLE(cls) \
  if (msg.Is<cls>()) { \
    return [msg] (std::string* err) { \
      cls unpacked; \
      msg.UnpackTo(&unpacked); \
      return Validate(unpacked, err); \
    }; \
  }

  X_TESTS_HARNESS_CASES_BOOL(TRY_RETURN_VALIDATE_CALLABLE)

  // TODO(akonradi) remove this once all C++ validation code is done
  return [](std::string*) { return true; };
}

}  // namespace

int main() {
  TestCase test_case;
  ExitIfFailed(test_case.ParseFromIstream(&std::cin), "failed to parse TestCase");

  auto validate_fn = GetValidationCheck(test_case.message());

  ValidateOrExit(validate_fn);

  return 0;
}
