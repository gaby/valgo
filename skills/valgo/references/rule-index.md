# Validator selection and rule orientation

Use this compact index to choose where to look. It reflects the Valgo v0.9.0
source when this reference was written; confirm constructors and methods in the
consumer's effective Valgo source before generating code.

## Choose the narrowest constructor

| Go value | Value constructor | Pointer constructor | Notes |
| --- | --- | --- | --- |
| `string` or a defined `~string` type | `String` | `StringP` | Use rune-length rules for characters and byte-length rules for payload size. |
| `bool` or a defined `~bool` type | `Bool` | `BoolP` | Prefer over `Comparable` or `Typed`. |
| `int`, `int8`, `int16`, `int32`, `int64` | `Int` | `IntP` | Width-specific constructors are deprecated in current source. Use `Rune`/`RuneP` when an `int32` is semantically a rune. |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `Uint` | `UintP` | Width-specific constructors are deprecated in current source. Use `Byte`/`ByteP` when a `uint8` is semantically a byte. |
| `float32`, `float64` | `Float` | `FloatP` | Width-specific constructors are deprecated in current source. Includes finite/NaN/infinite rules. |
| A generic value spanning numeric families | `Number` | `NumberP` | Use only when the code intentionally accepts Valgo's `TypeNumber` family. |
| `time.Time` | `Time` | `TimeP` | Uses time-specific ordering rules. |
| Another `comparable` domain type | `Comparable` | `ComparableP` | Type-safe equality and membership; no ordering rules. |
| Another statically typed value | `Typed` | — | Use for a typed `Passing` predicate or nil check. |
| A genuinely dynamic `any` value | `Any` | — | Last resort when the compile-time type is unavailable. |

Valgo has no container validator dedicated to slices or maps in this version.
Validate slice elements with `InRow` or `InCell`. Use a typed `Passing`
predicate only when validating a collection as a whole is the intended rule.

The generalized `Int`, `Uint`, and `Float` constructors and the `is` predicate
subpackage start in v0.9.0. For v0.8.1 and earlier, inspect source and retain
the matching width-specific constructor.

## Find the relevant rule family

- `String`/`StringP`: equality and ordering; `Empty`, `Blank`; byte rules
  `MaxBytes`, `MinBytes`, `ByteLength`, `ByteLengthBetween`; rune rules
  `MaxLength`, `MinLength`, `Length`, `LengthBetween`; `InSlice`,
  `MatchingTo`, and `Passing`. Pointer-only rules include `Nil`,
  `EmptyOrNil`, and `BlankOrNil`.
- Numeric families: `EqualTo`, ordering, inclusive `Between`, `Zero`,
  `InSlice`, and `Passing`. Signed integers add `Positive` and `Negative`.
  Floats add `Positive`, `Negative`, `NaN`, `Infinite`, and `Finite`.
  Pointer variants add `Nil` and `ZeroOrNil`.
- `Bool`/`BoolP`: `EqualTo`, `True`, `False`, `InSlice`, and `Passing`.
  The pointer validator adds `Nil` and `FalseOrNil`.
- `Time`/`TimeP`: `EqualTo`, `After`, `AfterOrEqualTo`, `Before`,
  `BeforeOrEqualTo`, inclusive `Between`, `Zero`, `InSlice`, and `Passing`.
  The pointer validator adds `Nil` and `NilOrZero`.
- `Comparable`/`ComparableP`: `EqualTo`, `InSlice`, and `Passing`; the pointer
  validator adds `Nil`.
- `Typed`: `Passing` and `Nil`.
- `Any`: `EqualTo`, `Passing`, and `Nil`.

All current built-in validators expose `Not()`, `Or()`, and `OrElse()`, but
`OrElse()` is unavailable before v0.8. Most rule methods accept an optional
final custom message template; verify the signature instead of assuming.

## Keep sessions and rules distinct

Use package functions or `*Validation` methods to assemble a result:

- Start: `Is`, `Check`, `New`, `In`, `InRow`, `InCell`
- Add/combine: `Is`, `Check`, `In`, `InRow`, `InCell`, `Merge`
- Conditional: `If`, `When`, `Do`; v0.8 also adds the `If*Valid` and
  `When*Valid` families
- Query: `Valid`; version-dependent `IsValid`, `PathValid`, `AllValid`,
  `AnyValid`
- Errors: `AddErrorMessage`, `ToError`, `ToValgoError`

Do not pass a `*Validation` where a `Validator` is required or chain a session
method on a validator. `New` accepts optional `Options`, not validators.
