#ifndef _VALIDATE_H
#define _VALIDATE_H

#include <stdexcept>

#include <google/protobuf/message.h>
#include <google/protobuf/util/time_util.h>

namespace pgv {

namespace protobuf = google::protobuf;
namespace protobuf_wkt = google::protobuf;

class UnimplementedException : public std::runtime_error {
 public:
  UnimplementedException() : std::runtime_error("not yet implemented") {}
  // Thrown by C++ validation code that is not yet implemented.
};

using ValidationMsg = std::string;

static inline std::string String(const ValidationMsg& msg) {
  return std::string(msg);
}

} // namespace pgv

#endif // _VALIDATE_H
