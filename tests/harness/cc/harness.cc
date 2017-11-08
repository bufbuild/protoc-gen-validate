#include <iostream>

#include "tests/harness/cases/bool.pb.h"
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

using tests::harness::TestCase;
using tests::harness::TestResult;
using google::protobuf::Any;
using google::protobuf::Message;

std::ostream& operator<<(std::ostream& out, const TestResult& result) {
  out << "valid: " << result.valid() << " reason: '" << result.reason() << "'"
      << std::endl;
  return out;
}

void resp(const TestResult& result) {
  if (!result.SerializeToOstream(&std::cout)) {
    std::cerr << "could not martial response: ";
    std::cerr << result << std::endl;
    exit(1);
  }
  exit(0);
}

void check_err(bool succeeded, const std::string& err_msg) {
  if (succeeded) {
    return;
  }

  TestResult result;
  result.set_error(true);
  result.set_reason(err_msg);
  resp(result);
}

void check_valid(const std::function<bool(std::string*)>& validate_fn) {
  std::string error_msg;
  TestResult result;

  result.set_valid(validate_fn(&error_msg));
  result.set_reason(std::move(error_msg));
  resp(result);
}

std::pair<std::unique_ptr<Message>, std::function<bool(std::string*)>> get_validate(const Any& msg) {
  // TODO(akonradi) remove this once all C++ validation code is done
  auto default_validate = [](std::string*) { return true; };

  // TODO(akonradi) use Any::UnpackTo to unpack messages
  return {nullptr, default_validate};
}

int main() {
  TestCase test_case;
  check_err(test_case.ParseFromIstream(&std::cin), "failed to parse TestCase");

  auto validate_pair = get_validate(test_case.message());
  auto message = std::move(validate_pair.first);
  auto validate_fn = std::move(validate_pair.second);

  check_valid(validate_fn);

  return 0;
}
