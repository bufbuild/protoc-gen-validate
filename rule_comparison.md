# Constraint Rule Comparison
## Global
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| disabled               |✅|✅|✅|✅|

## Numerics
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| const                  |✅|✅|✅|✅|
| lt/lte/gt/gte          |✅|✅|✅|✅|
| in/not_in              |✅|✅|✅|✅|

## Bools
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| const                  |✅|✅|✅|✅|

## Strings
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| const                  |✅|✅|✅|✅|
| len/min\_len/max_len   |✅|✅|❌|✅|
| min\_bytes/max\_bytes  |✅|✅|✅|✅|
| pattern                |✅|✅|❌|✅|
| prefix/suffix/contains |✅|✅|✅|✅|
| in/not_in              |✅|✅|✅|✅|
| email                  |✅|✅|❌|✅|
| hostname               |✅|✅|✅|✅|
| address                |✅|✅|✅|✅|
| ip                     |✅|✅|✅|✅|
| ipv4                   |✅|✅|✅|✅|
| ipv6                   |✅|✅|✅|✅|
| uri                    |✅|✅|✅|✅|
| uri_ref                |✅|✅|✅|✅|

## Bytes
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| const                  |✅|✅|✅|✅|
| len/min\_len/max_len   |✅|✅|✅|✅|
| pattern                |✅|✅|✅|✅|
| prefix/suffix/contains |✅|✅|✅|✅|
| in/not_in              |✅|✅|✅|✅|
| ip                     |✅|✅|❌|✅|
| ipv4                   |✅|✅|❌|✅|
| ipv6                   |✅|✅|❌|✅|

## Enums
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| const                  |✅|✅|✅|✅|
| defined_only           |✅|✅|✅|✅|
| in/not_in              |✅|✅|✅|✅|

## Messages
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| skip                   |✅|✅|✅|✅|
| required               |✅|✅|✅|✅|

## Repeated
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| min\_items/max_items   |✅|✅|✅|✅|
| unique                 |✅|✅|✅|✅|
| items                  |✅|✅|❌|✅|

## Maps
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| min\_pairs/max_pairs   |✅|✅|❌|✅|
| no_sparse              |✅|✅|❌|❌|
| keys                   |✅|✅|❌|✅|
| values                 |✅|✅|❌|✅|

## OneOf
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| required               |✅|✅|✅|✅|

## WKT Scalar Value Wrappers
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| wrapper validation     |✅|✅|✅|✅|

## WKT Any
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| required               |✅|✅|✅|✅|
| in/not_in              |✅|✅|✅|✅|

## WKT Duration
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| required               |✅|✅|✅|✅|
| const                  |✅|✅|✅|✅|
| lt/lte/gt/gte          |✅|✅|✅|✅|
| in/not_in              |✅|✅|✅|✅|

## WKT Timestamp
| Constraint Rule | Go | GoGo | C++ | Java |
| ---| :---: | :---: | :---: | :---: |
| required               |✅|✅|❌|✅|
| const                  |✅|✅|❌|✅|
| lt/lte/gt/gte          |✅|✅|❌|✅|
| lt_now/gt_now          |✅|✅|❌|✅|
| within                 |✅|✅|❌|✅|
