# [![](./.github/buf-logo.svg)][buf] protoc-gen-validate (PGV)

![License](https://img.shields.io/github/license/bufbuild/protoc-gen-validate?color=blue)
![Release](https://img.shields.io/github/v/release/bufbuild/protoc-gen-validate?include_prereleases)
![Slack](https://img.shields.io/badge/slack-buf-%23e01563)

📣**_Update: The next generation of `protoc-gen-validate`, now called
`protovalidate`, is available in beta for Golang, Python, Java, and C++!
We're hard at work on a TypeScript implementations as well. To learn more, check out
our [blog post][pv-announce]. We value your input in refining our products, so
don't hesitate to share your feedback on `protovalidate`._**

> ❓ Looking for protoc-gen-validate docs? Go [here](docs.md).

`protovalidate` is a series of libraries designed to validate Protobuf messages at
runtime based on user-defined validation rules. Powered by Google's Common
Expression Language ([CEL][cel-spec]), it provides a
flexible and efficient foundation for defining and evaluating custom validation
rules. The primary goal of `protovalidate` is to help developers ensure data
consistency and integrity across the network without requiring generated code.

Runtime implementations of `protovalidate` can be found in their own repositories:

- Go: [`protovalidate-go`][pv-go] (beta release)
- C++: [`protovalidate-cc`][pv-cc] (beta release)
- Java: [`protovalidate-java`][pv-java] (beta release)
- Python: [`protovalidate-python`][pv-python] (beta release)
- TypeScript: `protovalidate-ts` (coming soon)

[buf]:             https://buf.build
[protoc-source]:   https://github.com/google/protobuf
[protoc-releases]: https://github.com/google/protobuf/releases
[pg*]:             https://github.com/bufbuild/protoc-gen-star
[re2]:             https://github.com/google/re2/wiki/Syntax
[wkts]:            https://developers.google.com/protocol-buffers/docs/reference/google.protobuf
[pv]:              https://github.com/bufbuild/protovalidate
[pv-go]:           https://github.com/bufbuild/protovalidate-go
[pv-java]:           https://github.com/bufbuild/protovalidate-java
[pv-cc]:           https://github.com/bufbuild/protovalidate-cc
[pv-python]:       https://github.com/bufbuild/protovalidate-python
[pv-announce]:     https://buf.build/blog/protoc-gen-validate-v1-and-v2/
[cel-spec]:     https://github.com/google/cel-spec
