# Contributing Guidelines

Firstly, we want to say thank you for considering contributing
to `protovalidate`. We genuinely appreciate your help. This document aims to
provide some guidelines to make your contribution process straightforward and
meaningful.

## Code of Conduct

We pledge to maintain a welcoming and inclusive community. Please read
our [Code of Conduct](../CODE_OF_CONDUCT.md) before participating.

## How Can I Contribute?

### Reporting Bugs

Bugs are tracked as GitHub issues. If you discover a problem
with `protovalidate`, we want to hear about it. Here's how you can report a bug:

1. __Ensure the bug was not already reported__: Before creating a new issue,
   please do a search
   in [issues](https://github.com/bufbuild/protovalidate/issues) to see if
   the problem has already been reported. If it has and the issue is still open,
   add a comment to the existing issue instead of opening a new one.

2. __Check if the issue is fixed__: Try to reproduce the issue using the
   latest `main` branch to see if the problem has already been fixed. If fixed,
   that's great!

3. __Open new issue__: If the issue has not been reported and has not been
   fixed, then we encourage you to [open a new issue][file-bug].

Remember to detail the steps to reproduce the issue. This information is
invaluable in helping us fix the issue.

Once you've filled in the template, hit "Submit new issue", and we will take
care of the rest. We appreciate your contribution to making `protovalidate`
better!

### Suggesting Enhancements

We welcome ideas for enhancements and new features to improve `protovalidate`.
If you have an idea you'd like to share, if you want to expand language
support,  
please read [the section below](#language-support-requirements) first.

1. __Check if the enhancement is already suggested__: Before creating a new
   issue, please do a search
   in [issues](https://github.com/bufbuild/protovalidate/issues) to see if
   the idea or enhancement has already been suggested. If it has and the issue
   is still open, add a comment to the existing issue instead of opening a new
   one.

2. __Open a new issue__: If your enhancement hasn't been suggested before,
   please [create a new issue][file-feature-request].

3. __Discussion__: Once you've submitted the issue, maintainers or other
   community members might jump in to discuss the enhancement. Be prepared to
   provide more context or insights about your suggestion.

Remember, the goal of suggesting an enhancement is to improve `protovalidate`
for everyone. Every suggestion is valued, and we thank you in advance for your
contribution.

### Pull Requests

For changes, improvements, or fixes, please create a pull request. Make sure
your PR is up-to-date with the main branch. Please write clear and concise
commit messages to help us understand and review your PR.

## Language Support Requirements

We aim for `protovalidate` to support multiple languages, including but not
limited to Go, Java, Python, C++, and Typescript. Here are the requirements for
adding a new language:

1. __Conformance__: Make sure that your language addition passes the conformance
   test suite. This ensures that your addition meets the project's standards and
   behaves as expected.

2. __CEL Interpreter__: Implement a Common Expression Language (CEL) interpreter
   in your chosen language. CEL is a non-Turing complete language that makes it
   easy to write simple expressions, and it's crucial to `protovalidate`.

3. __Custom Function Equivalence__: Ensure that custom functions have equivalent
   behavior across all languages. This uniformity is essential to maintain the
   integrity and consistency of the project. Check out
   the [Custom Functions](../docs/cel.md#custom-library-in-protovalidate) for more

If you are interested in adding a new language to `protovalidate`, please open
an issue to discuss the details and requirements. We will be more than happy to
guide you through the process.

### Minimizing Performance Regression

Performance and efficient resource management are critical aspects
of `protovalidate`. CEL, being non-Turing complete, provides production safety
controls to limit execution time, helping to prevent excessive resource
consumption during evaluation. Here are some guidelines for effectively managing
resource constraints and minimizing performance regressions:

1. __Understanding Resource Constraints__: CEL has resource constraint features
   which provide feedback about expression complexity. These are designed to
   prevent CEL evaluation from consuming excessive resources. One key element is
   the concept of a _cost unit_, an independent measure used for tracking CPU
   utilization regardless of system load or hardware. The cost is deterministic,
   meaning for any CEL expression and input data, the evaluation cost will
   remain the same.
2. __Cost Units and Operations__: Many of CEL's operations have fixed costs.
   Simple operations, such as comparisons (e.g. `<`), have a cost of 1, while
   some, like list literal declarations, have a higher fixed cost of 40 cost
   units. Functions implemented in native code approximate cost based on time
   complexity. For instance, regular expression operations like `match`
   and `find` use an approximated cost
   of `length(regexString)*length(inputString)`, reflecting the worst-case time
   complexity of the operation.
3. __Testing the Overall Cost__: You can use the CEL library to test the overall
   cost of an expression. This can help predict the resources required to
   evaluate an expression and prevent operations that would consume excessive
   resources.
4. __Benchmark and Profile__: Benchmark your changes against the current `main`
   branch to evaluate the performance impact. If a performance regression is
   suspected, profile the code to pinpoint the bottleneck.
5. __Optimize__: Always look for ways to optimize your changes without
   compromising readability, maintainability, or correctness of your code.
6. __Discuss__: If your changes might cause a performance regression or resource
   constraint, but you believe they're still beneficial, discuss this in the
   pull request. Explain why you think the performance regression or resource
   constraint might be acceptable.

By keeping performance and resource management in mind throughout the
development process, we can ensure `protovalidate` remains efficient and
responsive, even as we add new features and fix bugs.

## Questions?

If you have any questions, please don't hesitate to create an issue, and we'll
answer as soon as possible. If your question is regarding a specific issue or
pull request, please link it in your comment.

## Thank You

Again, we appreciate your help and time, and we are excited to see your
contributions!

Remember, you can reach out to us at any time, and we're looking forward to
working together to make `protovalidate` the best it can be.

[file-bug]: https://github.com/bufbuild/protovalidate/issues/new?assignees=&labels=Bug&template=bug_report.md&title=%5BBUG%5D

[file-feature-request]: https://github.com/bufbuild/protovalidate/issues/new?assignees=&labels=Feature&template=feature_request.md&title=%5BFeature+Request%5D

[cel-spec]: https://github.com/google/cel-spec
