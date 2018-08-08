#ifndef _VALIDATE_H
#define _VALIDATE_H

#include <functional>
#include <stdexcept>
#include <string>
#include <typeinfo>

#include <google/protobuf/message.h>
#include <google/protobuf/util/time_util.h>

namespace pgv {
using std::string;

namespace protobuf = google::protobuf;
namespace protobuf_wkt = google::protobuf;

class UnimplementedException : public std::runtime_error {
public:
  UnimplementedException() : std::runtime_error("not yet implemented") {}
  // Thrown by C++ validation code that is not yet implemented.
};

using ValidationMsg = std::string;

class BaseValidator {
protected:
  static std::map<size_t, BaseValidator*> validators;
};

std::map<size_t, BaseValidator*> __attribute__((weak)) BaseValidator::validators;

template <typename T>
class Validator : public BaseValidator {
public:
  Validator(std::function<bool(const T&, ValidationMsg*)> check) : check_(check)
  {
    validators[typeid(T).hash_code()] = this;
  }

  static bool CheckMessage(const T& m, ValidationMsg* err)
  {
    auto val = static_cast<Validator<T>*>(validators[typeid(T).hash_code()]);
    if (val) {
      return val->check_(m, err);
    }
    return true;
  }

private:
  std::function<bool(const T&, ValidationMsg*)> check_;
};

static inline std::string String(const ValidationMsg& msg)
{
  return std::string(msg);
}

static inline bool IsPrefix(const string& maybe_prefix, const string& search_in)
{
  return search_in.compare(0, maybe_prefix.size(), maybe_prefix) == 0;
}

static inline bool IsSuffix(const string& maybe_suffix, const string& search_in)
{
  return maybe_suffix.size() <= search_in.size() && search_in.compare(search_in.size() - maybe_suffix.size(), maybe_suffix.size(), maybe_suffix) == 0;
}

static inline bool Contains(const string& search_in, const string& to_find)
{
  return search_in.find(to_find) != string::npos;
}

} // namespace pgv

#endif // _VALIDATE_H
