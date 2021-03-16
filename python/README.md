# Protoc-gen-validate (PGV)
While protocol buffers effectively guarantee the types of structured data, 
they cannot enforce semantic rules for values. This package is a python implementation
of [protoc-gen-validate][pgv-home], which allows for runtime validation of various 
semantic assertions expressed as annotations on the protobuf schema. The syntax for all available annotations is
in `validate.proto`. Implemented Python annotations are listed in the [rules comparison][rules-comparison].

### Example
```python3
from entities_pb2 import Person
from protoc_gen_validate.validator import validate, ValidationFailed

p = Person(first_name="Foo", last_name="Bar", age=42)
try:
    validate(p)
except ValidationFailed as err:
    print(err)
```

[pgv-home]: https://github.com/envoyproxy/protoc-gen-validate
[rules-comparison]: https://github.com/envoyproxy/protoc-gen-validate/blob/main/rule_comparison.md