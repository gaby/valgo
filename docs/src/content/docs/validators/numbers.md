---
title: Number Validators for Go
description: Validate Go numeric values and pointers with Valgo generic validators, comparison rules, range checks, and type-specific helpers.
---

Valgo provides a generic validator for all supported numeric types and
family-specific constructors that preserve the exact input type.

## Number

`Number()` accepts any value whose underlying type is a signed integer,
unsigned integer, `float32`, or `float64`.

```go
v.Is(v.Number(10).EqualTo(10))
v.Is(v.Number(11).GreaterThan(10))
v.Is(v.Number(11).Between(10, 12)) // inclusive
v.Is(v.Number(0).Zero())
v.Is(v.Number(20).InSlice([]int{10, 20, 30}))
```

`Number()` provides equality, ordering, inclusive `Between()`, `Zero()`,
`InSlice()`, and `Passing()`. Use `NumberP()` for pointers; it additionally
provides `Nil()` and `ZeroOrNil()`.

```go
var amount *int
v.Is(v.NumberP(amount, "amount").Nil())
```

## Signed integers

`Int()` accepts every signed integer type: `int`, `int8`, `int16`, `int32`,
`int64`, and custom types based on them. Use `IntP()` for pointers. `Rune()`
and `RuneP()` remain available when the `int32` value is semantically a rune.

```go
v.Is(v.Int(10).Positive())
v.Is(v.Int(int16(10)).GreaterThan(int16(5)))
v.Is(v.Int(int64(-1)).Negative())

age := int16(21)
v.Is(v.IntP(&age).GreaterOrEqualTo(int16(18)))
```

Signed integer validators add `Positive()` and `Negative()` to the common
numeric rules. The width-specific constructors `Int8()`, `Int16()`, `Int32()`,
`Int64()` and their pointer forms are deprecated aliases; use `Int()` or
`IntP()` in new code.

## Unsigned integers

`Uint()` accepts every unsigned integer type: `uint`, `uint8`, `uint16`,
`uint32`, `uint64`, and custom types based on them. Use `UintP()` for pointers.
`Byte()` and `ByteP()` remain available when the `uint8` value is semantically
a byte.

```go
v.Is(v.Uint(uint64(10)).LessOrEqualTo(uint64(10)))
v.Is(v.Byte(byte(1)).GreaterThan(byte(0)))
```

The width-specific constructors `Uint8()`, `Uint16()`, `Uint32()`, `Uint64()`
and their pointer forms are deprecated aliases; use `Uint()` or `UintP()` in
new code.

## Floating-point numbers

`Float()` accepts `float32`, `float64`, and custom types based on them. Use
`FloatP()` for pointers.

```go
v.Is(v.Float(float32(1.5)).GreaterThan(float32(1.0)))
v.Is(v.Float(3.14).Finite())
```

Float validators add `Positive()`, `Negative()`, `NaN()`, `Infinite()`, and
`Finite()` to the common numeric rules. `Float32()`, `Float64()`, `Float32P()`,
and `Float64P()` are deprecated aliases; use `Float()` or `FloatP()` in new
code.

Use the [`is` predicate package](/validators/predicates/) when you need the
same numeric checks as standalone boolean functions without a validation
session or error messages.
