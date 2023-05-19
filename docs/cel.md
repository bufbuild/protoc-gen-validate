# Common Expression Language (CEL)

The [Common Expression Language (CEL)](https://github.com/google/cel-spec) is
utilized by `protovalidate` for the formulation of validation rules, serving
both as the basis for the standard constraints as well as custom constraints.

[CEL](https://github.com/google/cel-spec/blob/master/doc/langdef.md) features a
user-friendly syntax, reminiscent of the expressions found in C, C++, Java,
JavaScript, and Go.

## Usage in `protovalidate`

The CEL expressions within protovalidate can expect to have access to the
following features of the CEL language:

- CEL's standard functions, as detailed in the [list of standard definitions](https://github.com/google/cel-spec/blob/master/doc/langdef.md#list-of-standard-definitions)
- CEL's standard [macros](https://github.com/google/cel-spec/blob/v0.8.0/doc/langdef.md#macros)
- CEL's [extended string function library](https://pkg.go.dev/github.com/google/cel-go/ext#Strings)
- `protovalidate`'s [custom functions and overloads](#custom-functions-and-overloads)

`protovalidate` seeks to be as portable as possible. As such, providing further
functionality beyond what's captured here is not recommended and not support by
the library or implementations.

### Custom functions and overloads

Certain pre-existing functionality inherited from `protoc-gen-validate` cannot
be represented in CEL, requiring the usage of custom CEL functions. Below is a
list of new functions as well as overrides to the standard functions included
in `protovalidate`. These functions and overloads are free to use
within [custom constraints](custom-constraints.md).

| Symbol       | Type                                                                                                                                                                                                        | Description                                                                                        |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|
| `contains`   | `bytes.contains(bytes) -> bool`                                                                                                                                                                             | Overload of the CEL standard `contains` to support bytes                                           |
| `endsWith`   | `bytes.endsWith(bytes) -> bool`                                                                                                                                                                             | Overload of the CEL standard `endWith` to support bytes                                            |
| `isEmail`    | `string.isEmail() -> bool`                                                                                                                                                                                  | Test whether the string is a valid email address                                                   |
| `isHostname` | `string.isHostname() -> bool`                                                                                                                                                                               | Test whether the string is a valid hostname                                                        |
| `isIp`       | `string.isIp() -> bool`<br/>`string.isIp(4) -> bool`<br/>`string.isIp(6) -> bool`                                                                                                                           | Test whether the string is a valid IP address, optionally limited to a specific version (v4 or v6) |
| `isUriRef`   | `string.isUriRef() -> bool`                                                                                                                                                                                 | Tests whether the string is a valid (absolute or relative) URI                                     |
| `isUri`      | `string.isUri() -> bool`                                                                                                                                                                                    | Tests whether the string is a valid absolute URI                                                   |
| `now`        | `now() -> google.protobuf.Timestamp`                                                                                                                                                                        | Get the current timestamp                                                                          |
| `startsWith` | `bytes.startsWith(bytes) -> bool`                                                                                                                                                                           | Overload of the CEL standard `startsWith` to support bytes                                         |
| `unique`     | `list(bool).unique() -> bool`<br/>`list(int).unique() -> bool`<br/>`list(uint).unique() -> bool`<br/>`list(double).unique() -> bool`<br/>`list(string).unique() -> bool`<br/>`list(bytes).unique() -> bool` | Test whether the items in the list are all unique.                                                 |