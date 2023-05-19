# Constraints

In the context of `protovalidate`, "constraints" refer to the conditions or
requirements that the data fields within a Protobuf message must meet. These
constraints are defined by developers using the provided custom option
annotations to help ensure that the data is semantically valid.

Constraints can be applied to Protobuf messages, their fields, and oneof unions.
In `protovalidate`, constraints are defined using the provided syntax or
custom expressions within the Protobuf message definition. At runtime, when the
Protobuf message is passed to the library, it checks whether the input data
meets the specified constraints. If the data fails to meet the constraints, a
series of validation
violations are raised in an error, indicating the issue with the input data.
Constraints fall into 2 categories:

## Standard constraints

These options supplied by `protovalidate` are used
to define validation rules and constraints for individual fields in a Protobuf
message. These options include specifications for field types, minimum and
maximum values, regular expressions, and custom validation functions.

To understand all of the standard constraints available in `protovalidate`, see
the [Standard Constraints](standard-constraints.md) page.

## Custom constraints

These field-level and message-level options are defined by the user and used to
define custom validation rules and constraints in a Protobuf message. They
are defined using the Common Expression Language (CEL) and enable complex rules
and constraints that aren't supported by the standard constraints. They can
also reference other fields in the message.

To understand how to define custom constraints in `protovalidate`, see the
[Custom constraints](custom-constraints.md) page.

## Next steps

- [Standard constraints](standard-constraints.md)
- [Custom constraints](custom-constraints.md)
