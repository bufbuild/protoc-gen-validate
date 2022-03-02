#include "tests/harness/cases/bool.pb.h"
#include "tests/harness/cc/other.pb.h"
#include "validate/validate.h"

int main() {
  tests::harness::cc::Foo foo;

  // This does not have an associated validator but should still pass.
  std::string err;
  if (!pgv::BaseValidator::AbstractCheckMessage(foo, &err)) {
    EXIT_FAILURE;
  }

  tests::harness::cases::BoolConstTrue bool_const_true;
  bool_const_true.set_val(false);
  if (pgv::BaseValidator::AbstractCheckMessage(foo, &err)) {
    EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
