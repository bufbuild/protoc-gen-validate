# Migration guide

Migrating from `protoc-gen-validate` to `protovalidate` should be safe,
incremental, and relatively painless, but it will still require a handful of
operations to get there. To ease this burden, we've provided this documentation
and a migration tool to simplify the process.

## Migration process

1. **[OPTIONAL] Use `protovalidate` in legacy mode.** If supported, an
   implementation may provide optional "legacy support" for the
   `protoc-gen-validate` annotations. Using `protovalidate` in this way can
   allow immediate utilization of the new library without having to make any
   Protobuf changes. As an example, Go's `protovalidate` supports opting into
   [legacy support][go-legacy].

2. **[OPTIONAL] Format the `.proto` files.** The migration tool does a
   best-effort to preserve the formatting of the original file, but in some
   circumstances may produce valid yet ill-formatted code. To reduce noise
   related to reformatting, we recommend running a formatter (such as
   [`buf format`][format]) over the corpus of `.proto` files and committing
   these changes to make them consistent before proceeding.

3. **Run the migration tool.** This first run will add matching `protovalidate`
   annotations alongside any existing `protoc-gen-validate` ones. Optionally run
   the formatter again after the migration tool to clean up any strange output.
   ```shell
   go run ./tools/protovalidate-migrate -w /path/to/protos
   ```

4. **Use `protovalidate`.** Update code that uses `protoc-gen-validate` code
   generation to consume the `protovalidate` library instead (or
   simultaneously).

5. **[OPTIONAL] Update migrated annotations.** Rerunning the migration tool is a
   no-op for any `.proto` files that import `protovalidate`. To have the tool
   replace any existing `protovalidate` annotations, run it with the
   `-replace-protovalidate` flag. This will ensure the annotations match
   `protoc-gen-validate` annotations.
   ```shell
   go run ./tools/protovalidate-migrate -w --replace-protovalidate /path/to/protos
   ```

6. **Remove `protoc-gen-validate` annotations.** Once ready to make the switch,
   and references to the `protoc-gen-validate` generated code has been removed,
   the migration tool can be run again with the `-remove-legacy` flag to remove
   the legacy annotations from the `.proto` files.
   ```shell
   go run ./tools/protovalidate-migrate -w --remove-legacy /path/to/protos
   ```

## Migration Tool

```bash
go run ./tools/protovalidate-migrate <flags> /path/to/proto
```

#### Flags

| Flag                      | Description                                                                                | Default Value                   |
|---------------------------|--------------------------------------------------------------------------------------------|---------------------------------|
| `--verbose`, `-v`         | Enables verbose logging.                                                                   | `false`                         |
| `--write`, `-w`           | Overwrites target files in-place.                                                          | `false`                         |
| `--output`, `-o`          | Writes output to the given path.                                                           | None                            |
| `--legacy-import`         | Specifies a `protoc-gen-validate` proto import path.                                       | `"validate/validate.proto"`     |
| `--protovalidate-import`  | Specifies a `protovalidate` proto import path.                                             | `"buf/validate/validate.proto"` |
| `--remove-legacy`         | Allows the program to remove `protoc-gen-validate` options.                                | `false`                         |
| `--replace-protovalidate` | Replaces `protovalidate` options to match `protoc-gen-validate` options (only if present). | `false`                         |

#### Notes

- If neither `-w` nor `-o` is specified, the modified proto(s) are emitted to
  stdout.

## Standard Constraint Changes

As part of the migration process, please note the following changes to the
standard constraints between `protoc-gen-validate` and `protovalidate`:

### Message Constraints

- All message-level constraints have moved into a single
  option: `(buf.validate.message)`.

- **`(validate.ignored)`**: removed. `protovalidate` does not generate code,
  and as a result, generation does not need to be skipped.

- **`(validate.disabled)`**: moved to `(buf.validate.message).disabled`.

### Oneof Constraints

- All oneof-level constraints have moved into a single
  option: `(buf.validate.oneof)`.

- **`(validate.required)`**: moved to `(buf.validate.oneof).required`.

### Field Constraints

- All field-level constraints have moved into a single
  option: `(buf.validate.field)`.

- **`(validate.rules).<TYPE>.required`**: moved
  to `(buf.validate.field).required`.

- **`(validate.rules).message.skip`**: moved to `(buf.validate.field).skipped`.

- **`(validate.rules).<TYPE>.ignore_empty`**: moved
  to `(buf.validate.field).ignore_empty`.

- **`(validate.rules).map.no_sparse`**: removed. The original rule accommodated
  for a situation that was only possible in the Go code generation exclusively.
  `protovalidate` now treats a sparse map value as an empty value which matches
  the semantics of the Go Protobuf runtime.

---

## Migrating Manually

While the migration tool is designed to simplify the process of migrating
from `protoc-gen-validate` to `protovalidate`, there may be cases where manual
migration is required or preferred. The following steps outline the process of
manual migration:

1. **Understand the changes:** The first step to manual migration is
   understanding the changes between `protoc-gen-validate` and `protovalidate`.
   Review the ["Standard Constraint Changes"](#standard-constraint-changes)
   section of this guide to understand how constraints have changed.

2. **Update the imports:** Replace the import of `validate/validate.proto`
   with `buf/validate/validate.proto` in your `.proto` files.

3. **Update message-level constraints:** Replace `(validate.ignored)`
   and `(validate.disabled)` with the new `(buf.validate.message)` option as
   described in the ["Message Constraints"](#message-constraints) section.

4. **Update oneof constraints:** Replace `(validate.required)` for oneof fields
   with the new `(buf.validate.oneof)` option as described in
   the ["Oneof Constraints"](#oneof-constraints) section.

5. **Update field-level constraints:** Replace all field-level constraints,
   including `(validate.rules).<TYPE>.required`, `(validate.rules).message.skip`,
   and `(validate.rules).<TYPE>.ignore_empty`, with the
   new `(buf.validate.field)` option as described in
   the ["Field Constraints"](#field-constraints) section.

6. **Remove unnecessary constraints:** Remove
   the `(validate.rules).map.no_sparse` constraint as it's not supported
   in `protovalidate`.

7. **Test and validate:** After performing the above steps, test your Protobuf
   code to ensure it's functioning as expected. Review any warnings or errors,
   and make corrections as necessary.

[go-legacy]: /go/README.md#support-legacy-protoc-gen-validate-constraints

[format]: https://buf.build/docs/format/style/
