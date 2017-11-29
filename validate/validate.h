#ifndef _VALIDATE_H
#define _VALIDATE_H

#include <google/protobuf/message.h>

namespace pgv {
namespace validate {

template<class T>
struct MessageValidator {
  static bool Check(const T& m, std::string* err) {
    // do nothing by default
    return true;
  }
};

} // namespace validate
} // namespace pgv

#endif // _VALIDATE_H
