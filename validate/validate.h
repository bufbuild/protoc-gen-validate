#ifndef _VALIDATE_H
#define _VALIDATE_H

#include <string>

#include <google/protobuf/message.h>
#include <google/protobuf/util/time_util.h>

namespace pgv {
using std::string;

namespace protobuf = google::protobuf;
namespace protobuf_wkt = google::protobuf;

static inline bool IsPrefix(const string& haystack, const string& needle) {
  return haystack.compare(0, needle.size(), needle) == 0;
}

static inline bool IsSuffix(const string& haystack, const string& needle) {
  return (needle.size() <= haystack.size()) &&
    (haystack.compare(haystack.size() - needle.size(), needle.size(), needle) == 0);
}

static inline bool Contains(const string& haystack, const string& needle) {
  return haystack.find(needle) != string::npos;
}

} // namespace pgv

#endif // _VALIDATE_H
