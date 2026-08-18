---
title: Stateless Go Validation Predicates
description: Use Valgo validation rules as standalone, type-safe boolean
  functions without validation sessions, error messages, or localization.
slug: 0.9.0/validators/predicates
---

The `github.com/cohesivestack/valgo/is` package exposes Valgo's built-in rule
logic as stateless boolean functions. Use it in ordinary Go control flow, as a
building block for application-specific predicates, or anywhere a `bool` is
more useful than a `Validation` result.

```go
import "github.com/cohesivestack/valgo/is"

func validProduct(name string, stock int16) bool {
  return !is.StringBlank(name) &&
    is.StringLengthBetween(name, 2, 80) &&
    is.IntGreaterOrEqualTo(stock, int16(0))
}
```

Predicates do not store a field name, apply `Not()` or OR grouping, create
errors, format messages, or use a locale. Use the validators in the main
`valgo` package when you need those features.

## Reuse in compatible validators

The package is also intended to be a shared rule layer for third-party
validators that integrate with or aim to remain compatible with Valgo. Those
validators can reuse the same Go predicate functions while providing their own
validation context, chaining API, error format, or localization behavior.

Reusing these predicates avoids copying details such as inclusive range
boundaries, rune-versus-byte length, and pointer nil semantics. It also keeps a
compatible validator's underlying rule behavior aligned with Valgo when that
compatibility is required.

## Naming and rule families

Functions begin with their rule family and end with the rule name. For
example, the validator expression `v.String(name).Not().Blank()` corresponds
to `!is.StringBlank(name)`.

| Family | Available value predicates |
| --- | --- |
| `String` | `EqualTo`, ordering, inclusive `Between`, `Empty`, `Blank`, `InSlice`, `MatchingTo`, byte-length rules, and rune-length rules |
| `Number` | `EqualTo`, ordering, inclusive `Between`, `Zero`, `InSlice`, and `Passing` |
| `Int` | Number rules plus `Positive` and `Negative` |
| `Uint` | Number rules for unsigned integers |
| `Float` | Number rules plus `Positive`, `Negative`, `NaN`, `Infinite`, and `Finite` |
| `Bool` | `EqualTo`, `True`, `False`, and `InSlice` |
| `Time` | `EqualTo`, `After`, `AfterOrEqualTo`, `Before`, `BeforeOrEqualTo`, inclusive `Between`, `Zero`, and `InSlice` |
| `Comparable` | `EqualTo` and `InSlice` |
| General | `Passing` and `Nil` |

String byte functions are `StringMaxBytes`, `StringMinBytes`,
`StringByteLength`, and `StringByteLengthBetween`. Character-counting
functions are `StringMaxLength`, `StringMinLength`, `StringLength`, and
`StringLengthBetween`; they count UTF-8 runes just like the corresponding
validator methods.

The exported `Number`, `Int`, `Uint`, and `Float` constraints are available for
generic helpers. For example, a helper that spans every supported numeric type
can use `is.Number`:

```go
func nonNegative[T is.Number](value T) bool {
  return is.NumberGreaterOrEqualTo(value, 0)
}
```

## Pointer predicates

Pointer variants insert `P` between the family and rule names:
`StringPBlank`, `IntPPositive`, `FloatPFinite`, and so on. A nil pointer fails
ordinary pointer predicates. Explicit nil-accepting predicates include:

* `StringPEmptyOrNil` and `StringPBlankOrNil`
* `NumberPZeroOrNil`, `IntPZeroOrNil`, `UintPZeroOrNil`, and
  `FloatPZeroOrNil`
* `BoolPFalseOrNil`
* `TimePNilOrZero`

Each pointer family also provides a `PNil` function, such as `StringPNil` or
`TimePNil`.

```go
var nickname *string

is.StringPNil(nickname)        // true
is.StringPBlank(nickname)      // false
is.StringPBlankOrNil(nickname) // true
```

`NumberPPassing`, `IntPPassing`, `UintPPassing`, and `FloatPPassing` pass the
pointer itself to the callback, including when it is nil. This lets the
callback define its own nil semantics.

## Predicates with validator errors

An application predicate can also be supplied to a validator's `Passing()`
method when you need Valgo to record a message:

```go
func validCode(code string) bool {
  return is.StringLength(code, 8) && !is.StringBlank(code)
}

val := v.Is(
  v.String(input.Code, "code").
    Passing(validCode, "{{title}} must contain exactly 8 non-blank characters"),
)
```

The built-in validator methods and the `is` package share the same predicate
implementations, so their underlying rule semantics stay aligned.
