package is

import "reflect"

func ComparableEqualTo[T comparable](value, expected T) bool { return value == expected }

func ComparableInSlice[T comparable](value T, values []T) bool {
	for _, candidate := range values {
		if value == candidate {
			return true
		}
	}
	return false
}

func ComparablePEqualTo[T comparable](value *T, expected T) bool {
	return value != nil && ComparableEqualTo(*value, expected)
}

func ComparablePInSlice[T comparable](value *T, values []T) bool {
	return value != nil && ComparableInSlice(*value, values)
}

func ComparablePNil[T comparable](value *T) bool { return value == nil }

// AnyEqualTo uses interface equality and, like Go's == operator, panics if the
// dynamic values are not comparable. It exists for parity with Valgo's
// deprecated Any.EqualTo rule; prefer ComparableEqualTo for new code.
func AnyEqualTo(value, expected any) bool { return value == expected }

// Passing evaluates a caller-provided validation rule.
func Passing[T any](value T, function func(T) bool) bool { return function(value) }

// Nil reports whether value is nil or contains a nil pointer. This preserves
// the nil semantics of Valgo's Any and Typed validators.
func Nil(value any) bool {
	return value == nil ||
		(reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}
