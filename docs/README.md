# Documentation

This section of the documentation is dedicated to the `protovalidate` runtime
library. It serves as a comprehensive guide to help you understand and use the
library effectively. Here you will find everything you need to learn about the
Common Expression Language (CEL) and the constraint systems provided by the
library.

- [CEL Overview](cel.md): This document provides an in-depth understanding of
  the Common Expression Language (CEL). It is an excellent starting point if you
  are new to CEL or need a refresher on its core concepts.

- [Constraints](constraints.md): This section explores the concept of
  constraints in `protovalidate`. You'll learn about their purpose, how they
  work,
  and how they can be applied in your project.

    - [Custom Constraints](custom-constraints.md): This section explains how to
      define and apply your own custom constraints in `protovalidate`. This can
      be
      particularly useful when the standard constraints do not fit the specific
      requirements of your project.

    - [Standard Constraints](standard-constraints.md): Here, you will learn
      about the pre-defined constraints available in the `protovalidate`
      library.
      This guide will help you understand when and how to use these standard
      constraints effectively.

- [Errors](errors.md): This section explains the error system in `protovalidate`
  and provides guidance on how to handle them effectively.

## Tools

The tools section introduces optional build-time convenience tools that can be
used to optimize the development process with `protovalidate`. These tools are
designed to help you implement new language support and assist in migrating your
existing projects to `protovalidate`.

- [Conformance](conformance.md): This document is dedicated to explaining the
  Conformance tool. Learn how to use this tool to ensure your project aligns
  with the `protovalidate`'s rules and constraints effectively.

- [Migrate](migrate.md): If you're planning to migrate your existing project to
  `protovalidate`, this guide is for you. The Migrate tool is designed to help
  you transition smoothly, minimizing any potential disruption to your project's
  development.
