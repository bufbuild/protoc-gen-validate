# Errors

## Violation Errors

The `Violation` message represents a single violation of a constraint. It has
three fields:

- `field_path`: a machine-readable path to the field that failed validation.
- `constraint_id`: the id of the constraint that failed validation.
- `message`: a human-readable error message for this violation.

### Violations

The `Violations` message is used to collect all violations that occur during
constraint validation. It has one field:

- `violations`: a repeated field of `Violation` messages. `Violations` are
  returned by `protovalidate` when a message fails to fulfill the requirements
  of the constraint validation.

In summary, the `Constraints` and `Violations` messages provide a way to define
and enforce validation constraints using the Common Expression Language (CEL)
syntax. Standard constraints are defined using the `Constraints` message and
provide a set of predefined validation rules that can be applied to various.
These messages can be used to ensure that the data conforms to the
specified rules and to report any violations that occur during validation.

## Compilation Errors

Since the constraints are interpreted at runtime, compilation errors occur
during this period. These errors are most commonly due to syntax errors or
incorrect usage of the Common Expression Language (CEL) in defining the
constraints. Compilation errors might also arise from referencing a non-existent
field or incorrect usage of the `protovalidate` API.

## Runtime Errors

Although CEL expressions are type-checked, it is still possible for unexpected
types or invalid field accesses to result in runtime errors. For instance,
attempting to access a field that doesn't exist on a given message, or applying
a numeric operation to a string, can lead to such errors.

Another source of runtime errors is the violation of constraints. For example,
if a message fails to fulfill the requirements of the constraint
validation, `protovalidate` will return a `Violations` message, which is a type
of runtime error.

## Next steps

- [Standard constraints](standard-constraints.md)
- [Custom constraints](custom-constraints.md)
