# [![](./.github/buf-logo.svg)][buf] protovalidate

[![CI](https://github.com/bufbuild/protovalidate/actions/workflows/ci.yaml/badge.svg?branch=main)][ci]
[![Conformance](https://github.com/bufbuild/protovalidate/actions/workflows/conformance.yaml/badge.svg?branch=main)][conformance]
[![Slack](https://img.shields.io/badge/Slack-Buf-%23e01563)][slack]
[![BSR](https://img.shields.io/badge/BSR-Module-0C65EC)][buf-mod]

**_Update: The protoc-gen-validate v2.0, now called `protovalidate`, is 
available in beta for Golang! We're hard at work on TypeScript, C++, Java, 
and Python as well. To learn more, check out 
our [blog post](https://buf.build/blog/protoc-gen-validate-v1-and-v2/). We 
value your input in refining our products, so don't hesitate to share your 
feedback on `protovalidate`. Join us on [Slack][slack] to talk more._**

`protovalidate` is a polyglot library designed to validate Protobuf messages at
runtime based on user-defined validation rules. Powered by Google's Common
Expression Language ([CEL][cel-spec]), it provides a
flexible and efficient foundation for defining and evaluating custom validation
rules. The primary goal of `protovalidate` is to help developers ensure data
consistency and integrity across the network without requiring generated code.

> ❓ Looking for `protoc-gen-validate`? It's been moved to the [v1 branch][pgv].

## Usage

### Import protovalidate

To define constraints within your Protobuf messages,
import `buf/validate/validate.proto` into your `.proto` files:

```protobuf
syntax = "proto3";

package my.package;

import "buf/validate/validate.proto";
```

#### Build with [`buf`][buf]

Add a dependency on [`buf.build/bufbuild/protovalidate`][buf-mod] to your
module's [`buf.yaml`][buf-deps]:

```yaml
version: v1
# <snip>
deps:
  - buf.build/bufbuild/protovalidate
# <snip>
```

After modifying your `buf.yaml`, don't forget to run `buf mod update` to ensure
your dependencies are up-to-date.

#### Build with `protoc`

Add an import path (`-I` flag) pointing to the contents of the `proto/validate`
directory to your invocation of `protoc`:

```shell
protoc \ 
  -I ./vendor/protovalidate/proto/validate \
  # <snip>
```

### Implementing validation constraints

Validation constraints can be enforced using the `buf.validate` Protobuf package. The rules are specified directly in the `.proto` files.

Let's consider a few examples:

1. **Scalar field validation:** For a basic `User` message, we can enforce constraints such as a minimum length for the user's name.

   ```protobuf
   syntax = "proto3";
   
   import "buf/validate/validate.proto";
   
   message User {
     // User's name, must be at least 1 character long.
     string name = 1 [(buf.validate.field).string.min_len = 1];
   }
   ```

2. **Map field validation:** For a `Product` message with a map of item quantities, we can ensure that all quantities are positive.

   ```protobuf
   syntax = "proto3";
   
   import "buf/validate/validate.proto";
   
   message Product {
     // Map of item quantities, all quantities must be positive.
     map<string, int32> item_quantities = 1 [(buf.validate.field).map.values.int32.gt = 0];
   }
   ```

3. **Well-known type (WKT) validation:** For the `User` message, we can add a constraint to ensure the `created_at` timestamp is in the past.

   ```protobuf
   syntax = "proto3";
   
   import "google/protobuf/timestamp.proto";
   import "buf/validate/validate.proto";
   
   message User {
     // User's creation date must be in the past.
     google.protobuf.Timestamp created_at = 1 [(buf.validate.field).timestamp.lt_now = true];
   }
   ```

For more advanced or custom constraints, `protovalidate` allows for CEL expressions that can incorporate information across fields.

1. **Field-level expressions:** We can enforce that a products' `price`, sent as a string, includes a currency symbol like "$" or "£". We want to ensure that the price is positive and the currency symbol is valid.

   ```protobuf
   syntax = "proto3";
   
   import "buf/validate/validate.proto";
   
   message Product {
     string price = 1 [(buf.validate.field).cel = {
       id: "product.price",
       message: "Price must be positive and include a valid currency symbol ($ or £)",
       expression: "(this.startsWith('$') || this.startsWith('£')) && double(this.substring(1)) > 0"
     }];
   }
   ```

2. **Message-level expressions:** For a `Transaction` message, we can use a message-level CEL expression to ensure that the `delivery_date` is always after the `purchase_date`.

   ```protobuf
   syntax = "proto3";
   
   import "google/type/date.proto";
   import "buf/validate/validate.proto";
   
   message Transaction {
     google.type.Date purchase_date = 1;
     google.type.Date delivery_date = 2;
     
     option (buf.validate.message).cel = {
       id: "transaction.delivery_date",
       message: "Delivery date must be after purchase date",
       expression: "this.delivery_date > this.purchase_date"
     };
   }
   ```

3. **Producing an error message in the expression:** We can produce custom error messages directly in the CEL expressions. In this example, if the `age` is less than 18, the CEL expression will evaluate to the error message string.

   ```protobuf
   syntax = "proto3";
   
   import "buf/validate/validate.proto";
   
   message User {
     int32 age = 1 [(buf.validate.field).cel = {
       id: "user.age",
       expression: "this.age < 18 ? 'User must be at least 18 years old': ''"
     }];
   }
   ```

## Documentation

`protovalidate` provides a robust framework for validating Protobuf messages by
enforcing standard and custom constraints on various data types, and offering
detailed error information when validation violations occur. For a detailed
overview of all its components, the supported constraints, and how to use them
effectively, please refer to our [comprehensive documentation](docs/README.md).
The key components include:

- [**Standard Constraints**](https://github.com/bufbuild/protovalidate/blob/main/docs/standard-constraints.md): `protovalidate`
  supports a wide range of standard
  constraints for all field types as well as special functionality for the
  Protobuf Well-Known-Types. You can apply these constraints to your Protobuf
  messages to ensure they meet certain common conditions.

- [**Custom Constraints**](https://github.com/bufbuild/protovalidate/blob/main/docs/custom-constraints.md): With Google's Common
  Expression Language (CEL),
  `protovalidate` allows you to create complex, custom constraints to
  handle unique validation scenarios that aren't covered by the standard
  constraints at both the field and message level.

- [**Error Handling**](https://github.com/bufbuild/protovalidate/blob/main/docs/README.md#errors): When a violation
  occurs, `protovalidate`provides
  detailed error information to help you quickly identify the source and fix for
  an issue.

### Language-specific support & documentation

- [x] [Go](go/v2/) (Beta Release)
- [ ] Typescript (coming soon)
- [ ] C++ (coming soon)
- [ ] Java (coming soon)
- [ ] Python (coming soon)

### Migrating from protoc-gen-validate

`protovalidate`'s constraints very closely emulate those
in `protoc-gen-validate` to ensure an easy transition for developers. To
migrate from `protoc-gen-validate` to `protovalidate`, use the
provided [migration tool](https://github.com/bufbuild/protovalidate/blob/main/tools/protovalidate-migrate) to
incrementally upgrade your `.proto` files.

## Community

For help and discussion around Protobuf, best practices, and more, join us
on the Buf [Slack][slack].

## Ecosystem

- [Buf][buf]
- [CEL Go][cel-go]
- [CEL Spec][cel-spec]

## Status: BETA

This library is currently in a beta status. While it's functional and provides
the features described in the documentation, it may still contain some bugs,
incomplete features, or performance issues. We encourage users to test the
library and provide feedback on any issues they encounter.

Please note the following when using this beta version:

- **API stability**: The API might undergo changes as we continue to refine and
  improve the library. Make sure to keep up-to-date with the latest releases and
  check the [change log][changelog] for any breaking changes.
- **Bug reporting**: If you encounter any bugs,
  please [file a GitHub issue][file-bug].
  Provide a detailed description of the issue, steps to reproduce it, and
  any relevant system information or error messages.
- **Feature requests**: We welcome feature requests and suggestions for
  improvements. Please [file a GitHub issue][file-feature-request] to request a
  new feature or enhancement.
- **Support and feedback**: For questions, support, or general feedback, feel
  free to reach out to the maintainers via
  our [community Slack channel][slack].

We appreciate your interest in this library and your assistance in making it
better. Your contributions will help us advance the project from beta to a
stable release more quickly.

## Legal

Offered under the [Apache 2 license][license].

[buf]: https://buf.build
[cel-go]: https://github.com/google/cel-go
[cel-spec]: https://github.com/google/cel-spec
[changelog]: https://github.com/bufbuild/protovalidate/blob/main/CHANGELOG.md
[ci]: https://github.com/bufbuild/protovalidate/actions/workflows/ci.yaml
[conformance]: https://github.com/bufbuild/protovalidate/actions/workflows/conformance.yaml
[file-bug]: https://github.com/bufbuild/protovalidate/issues/new?assignees=&labels=Bug&template=bug_report.md&title=%5BBUG%5D
[file-feature-request]: https://github.com/bufbuild/protovalidate/issues/new?assignees=&labels=Feature&template=feature_request.md&title=%5BFeature+Request%5D
[license]: https://github.com/bufbuild/protovalidate/blob/main/LICENSE
[slack]: https://buf.build/links/slack
[pgv]: https://github.com/bufbuild/protoc-gen-validate/tree/v1
[buf-deps]: https://buf.build/docs/configuration/v1/buf-yaml/#deps
[buf-mod]: https://buf.build/bufbuild/protovalidate
