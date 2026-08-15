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

// Passing evaluates a caller-provided validation rule.
func Passing[T any](value T, function func(T) bool) bool { return function(value) }

// Nil reports whether value is nil or contains a nil pointer. This preserves
// the nil semantics of Valgo's Any and Typed validators.
func Nil(value any) bool {
	return value == nil ||
		(reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}
