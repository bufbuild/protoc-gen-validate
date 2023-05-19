# Custom constraints

Custom constraints in `protovalidate` are field-level and message-level
options that define validation rules and constraints for individual fields
within a Protobuf message. These options ensure that the data being passed
through your gRPC services adheres to the expected format, range, or other
requirements specific to your application.

Field-level constraints apply to specific data types within a message and can
include requirements such as minimum and maximum values, regular expressions,
and custom validation functions. Message-level constraints, on the other hand,
apply to the entire message, allowing for more complex validation scenarios that
depend on the relationships between fields or specific combinations of field
values.

- [Constraints](#constraints)
    - [Constraint Type](#constraint)
    - [Specifying validation rules with Constraint](#specifying-validation-rules-with-constraint)
- [Field-level constraints](#field-level-constraints)
    - [Importing required dependencies](#importing-required-dependencies)
    - [Defining custom constraints](#defining-custom-constraints)
    - [Message-level constraints](#message-level-constraints)
    - [Importing required dependencies](#importing-required-dependencies)
    - [Defining MessageConstraints](#defining-messageconstraints)

## Constraints

The `Constraint` message is used to define validation rules for messages or
fields in a message. The constraints are then evaluated using
the `protovalidate` tool to ensure that the data conforms to the specified
rules. If any violations occur during validation, they are reported using
the `Violations` message. The `Violation` message provides detailed information
about the violated constraint, including the field path, the constraint ID, and
the error message.

### Constraint

The `Constraint` message defines a validation constraint using the Common
Expression Language (CEL) syntax. It has three fields:

- `id`: a machine-readable name for this expression that should be unique to its
  scope (either a message or field).
- `message`: a human-readable error message for this expression if it evaluates
  to false. If this message is non-empty, strings from expression evaluation are
  ignored.
- `expression`: a CEL expression that performs a validation. The resolved type
  of the expression must be either a boolean or string value. The expression is
  considered rejected (i.e., validation failed) if the expression evaluates to
  false or a non-empty string.

View Documentation [here](https://buf.build/bufbuild/protovalidate/docs/main:buf.validate#buf.validate.Constraint)

### Specifying validation rules with Constraint

The expression must evaluate to a boolean or string value, and if it evaluates
to false or a non-empty string, the validation is considered to have failed. To
define validation rules, add `Constraint` messages to the `cel` field:

```protobuf
message Example{
  option (buf.validate.message).cel = {
    id: "message_expression_nested",
    message: "a must be greater than b",
    expression: "this.a > this.b"
  };

  int32 a = 1;
  int32 b = 2;
}
```

To create a validation rule, define a Constraint message with a unique
id, an appropriate error message, and a valid CEL expression. The validation
will fail if the expression evaluates to false or a non-empty string.

In addition to the libraries from the CEL community, protovalidate incorporates
CEL libraries that are universally accessible wherever CEL is employed within
protovalidate. Read more [here](cel.md#custom-functions-and-overloads).

## Field-level constraints

Field-level constraints provide a way to apply validation rules for an
individual field within a message. These constraints use CEL expressions to
define more complex validation scenarios.

Here's an example of a field-level custom constraint:

```protobuf
message Product {
  double price = 1 [(buf.validate.field).cel = {
    id: "product.price_range",
    message: "Price must be between 0 and 1000",
    expression: "this >= 0 && this <= 1000",
  }];
}
```

In this example, the field-level constraint ensures that the `price` value is
between `0` and `1000` for a valid `Product` message.

### Defining custom constraints

The `Constraint` message type represents an individual validation rule. It
contains three fields: `id`, `message`, and `expression`. The `id` is a unique,
machine-readable name for the validation rule. The message field stores a
human-readable error message to display when the validation rule fails. The
expression field holds a CEL expression that performs the validation:

```protobuf
message Constraint {
  string id = 1;
  string message = 2;
  string expression = 3;
}
```

To create a validation rule, define a `Constraint` message with a unique
`id`, an appropriate error `message`, and a valid CEL `expression`. The
validation will fail if the expression evaluates to `false` or a non-empty
string.

```protobuf
message Example{
  int32 a = 1 [(buf.validate.field).cel = {
    id: "gt_10"
    message: "a must be greater than 10"
    expression: "this > 10"
  }];
}
```

The schema also defines various rules for scalar, complex, and well-known field
types, which can be used to enforce specific constraints on each field type.

"Variables" are what the inputs to a CEL program are called. For any
protovalidate field that incorporates CEL, the accessible variables for that
field are defined is `this` and is always made available and refers to the
current state of the data being validated by the CEL expression.

---

### Message-level constraints

Message-level constraints provide a way to apply validation rules that depend on
the relationship between fields or specific combinations of field values within
a message. These constraints use CEL expressions to define more complex
validation scenarios.

Here's an example of a message-level constraint:

```protobuf
message Event {
  int64 start_time = 1;
  int64 end_time = 2;
  option (buf.validate.message).cel = {
    id: "event.start_time_before_end_time",
    message: "Start time must be before end time",
    expression: "this.start_time < this.end_time",
  };
}
```

In this example, the message-level constraint ensures that the `start_time`
value is less than the `end_time` value for a valid `Event` message.

By leveraging both field-level and message-level constraints, `protovalidate`
enables developers to create a comprehensive set of validation rules that ensure
the integrity and consistency of their data across gRPC services.

### Defining MessageConstraints

The `MessageConstraints` message defines custom
validation rules for a Protobuf message using CEL expressions. The
disabled field is an optional boolean value that, if set to true, nullifies
all validation rules for the message and its associated fields. The
Constraints field is a repeated field of Constraint messages, which specify
the validation rules to be applied to the message.

```protobuf
message MessageConstraints {
  ...
  // Constraints specifies the validation rules to be applied to this message.
  repeated Constraint cel = 3;
  ...
}
```

## Related docs

- [Standard constraints](standard-constraints.md)
- [Overview of Common Expression Language (CEL)](./cel.md)
