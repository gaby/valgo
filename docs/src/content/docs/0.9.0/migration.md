---
title: Valgo Migration Notes
description: Review Valgo v0 migration notes, breaking changes, and validation
  behavior updates for Go applications.
slug: 0.9.0/migration
---

Valgo is pre-v1.0, so breaking changes can happen.

## v0.9.0 - Stateless predicates and generalized numeric constructors

**Generalized numeric constructors**

Valgo originally introduced a separate constructor for each numeric width when
Go did not support generics. At the time, constructors such as `Int8()`,
`Int16()`, `Float32()`, and `Float64()` were a practical way to provide a
type-safe API for each built-in numeric type.

Go has supported generics since Go 1.18, and Valgo's numeric validator types
have been generic over an entire numeric family since v0.7. The specialized
constructors were retained for compatibility, but they no longer overcome a
language limitation on the Go versions Valgo supports. For example,
`ValidatorInt[T]` already supports every signed integer type, so requiring
`Int16()` instead of `Int()` for an `int16` value adds another API name without
providing different rules or stronger type safety.

In v0.9.0, each primary constructor now uses the same constraint as its
validator family:

* `Int()` and `IntP()` accept `int`, `int8`, `int16`, `int32`, and `int64`.
* `Uint()` and `UintP()` accept `uint`, `uint8`, `uint16`, `uint32`, and `uint64`.
* `Float()` and `FloatP()` accept `float32` and `float64`.

Callers can therefore choose a constructor by the validation behavior they
need—signed integer, unsigned integer, or floating point—instead of by storage
width.

Go infers and preserves the exact input type, including user-defined types, so
the broader constructors do not weaken type safety:

```go
type Attempts int8

attempts := Attempts(3)
limit := uint32(100)
ratio := float32(0.75)

v.Int(attempts).Between(0, 5)
v.Uint(limit).GreaterThan(0)
v.Float(ratio).Between(0, 1)
```

The specialized constructors now only repeat the underlying storage width;
they no longer provide the value they did before generics. Keeping every
width-specific name as a first-class API would make the package, documentation,
and future extensions larger without providing additional type safety,
validation behavior, or rules. They are therefore deprecated compatibility
aliases:

* `Int8`, `Int16`, `Int32`, `Int64` and their `P` forms
* `Uint8`, `Uint16`, `Uint32`, `Uint64` and their `P` forms
* `Float32`, `Float64` and their `P` forms

Existing code continues to work, but new code should use the primary family
constructors. Migration changes only the constructor name; the inferred value
type and validator rules stay the same:

```go
v.Int16(value)   // Deprecated
v.Int(value)     // Preferred

v.Float64P(&rate) // Deprecated
v.FloatP(&rate)   // Preferred
```

`Rune`/`RuneP` and `Byte`/`ByteP` remain supported because those names express
domain meaning, not merely storage width: a rune is a Unicode code point and a
byte is a unit of binary data.

**Built-in rule logic functions**

Valgo now exposes its built-in rule logic as ordinary boolean functions in
`github.com/cohesivestack/valgo/is`. These predicates do not create validation
contexts or localized errors. Existing validator methods keep their behavior
and now use the same predicates internally. See [Stateless
Predicates](/0.9.0/validators/predicates/) for the API and examples.

## v0.8.1 - Shorter string length validator names

String length validators now have shorter preferred names:

* `ByteLength` replaces `OfByteLength`
* `ByteLengthBetween` replaces `OfByteLengthBetween`
* `Length` replaces `OfLength`
* `LengthBetween` replaces `OfLengthBetween`

The `Of*` methods remain available as deprecated aliases in v0.8.1 for
compatibility. They will be removed in v1.0.

## v0.8.0 - Go version and validation flow

* Valgo v0.8 is tested with Go 1.23 and later. Using one of these versions is
  recommended.
* When all `Or()` alternatives fail, v0.8 produces one localized message that
  joins the alternatives. OR grouping and precedence remain consistent with
  v0.7.
* `OrElse()` adds a short-circuiting alternative that skips the rest of the
  validator chain when its left side succeeds.
* `PathValid()`, `AllValid()`, and `AnyValid()` query recorded validation
  results. `IsValid()` remains available but is deprecated in favor of
  `PathValid()`.
* The `If*Valid` and `When*Valid` methods conditionally merge validations or
  execute callbacks based on those recorded results.
* Custom validators can supply missing message entries through
  `ValidatorContext.WithLocaleFallback()`.

## v0.7.0 - Numeric validators switched to generics

Numeric validators are now generic per family:

* `ValidatorInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64]`
* `ValidatorUint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64]`
* `ValidatorFloat[T ~float32 | ~float64]`

Most call sites remain the same (constructors like `v.Int16(...)` still exist). You mainly need to update declared types.

## v0.6.0 - String length counts runes (characters)

String length validators now measure **runes** instead of bytes:

* A multi-byte UTF-8 character counts as 1.
* Use explicit byte-length validators if you need `len(s)` semantics.

Byte-based validators:

* `MaxBytes`
* `MinBytes`
* `OfByteLength`
* `OfByteLengthBetween`
