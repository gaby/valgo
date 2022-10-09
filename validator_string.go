package valgo

import (
	"regexp"
	"strings"
)

func isStringEqualTo[T ~string](v0 T, v1 T) bool {
	return v0 == v1
}
func isStringGreaterThan[T ~string](v0 T, v1 T) bool {
	return v0 > v1
}
func isStringGreaterOrEqualTo[T ~string](v0 T, v1 T) bool {
	return v0 >= v1
}
func isStringLessThan[T ~string](v0 T, v1 T) bool {
	return v0 < v1
}
func isStringLessOrEqualTo[T ~string](v0 T, v1 T) bool {
	return v0 <= v1
}
func isStringBetween[T ~string](v T, min T, max T) bool {
	return v >= min && v <= max
}
func isStringEmpty[T ~string](v T) bool {
	return len(v) == 0
}
func isStringBlank[T ~string](v T) bool {
	return len(strings.TrimSpace(string(v))) == 0
}
func isStringInSlice[T ~string](v T, slice []T) bool {
	for _, _v := range slice {
		if v == _v {
			return true
		}
	}
	return false
}
func isStringMatchingTo[T ~string](v T, regex *regexp.Regexp) bool {
	return regex.MatchString(string(v))
}
func isStringMaxLength[T ~string](v T, length int) bool {
	return len(v) <= length
}
func isStringMinLength[T ~string](v T, length int) bool {
	return len(v) >= length
}
func isStringLength[T ~string](v T, length int) bool {
	return len(v) == length
}
func isStringLengthBetween[T ~string](v T, min int, max int) bool {
	return len(v) >= min && len(v) <= max
}

type ValidatorString[T ~string] struct {
	context *ValidatorContext
}

// Receives a string value to validate.
//
// The value also can be a custom boolean type such as `type Status string;`
//
// Optionally, the function can receive a name and title, in that order,
// to be used in the error messages. A `value_%N`` pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0. When the name is
// provided but not the title, then the name is humanized to be used as the
// title as well; for example the name `phone_number` will be humanized as
// `Phone Number`

func String[T ~string](value T, nameAndTitle ...string) *ValidatorString[T] {
	return &ValidatorString[T]{context: NewContext(value, nameAndTitle...)}
}

// This function returns the context for the Valgo Validator session's
// validator. The function should not be called unless you are creating a custom
// validator by extending this validator.
func (validator *ValidatorString[T]) Context() *ValidatorContext {
	return validator.context
}

// Reverse the logical value associated to the next validation function.
// For example:
//
//	// It will return false because Not() inverts to Blank()
//	Is(v.String("").Not().Blank()).Valid()
func (validator *ValidatorString[T]) Not() *ValidatorString[T] {
	validator.context.Not()

	return validator
}

// Validate if a string value is equal to another.
// For example:
//
//	status := "running"
//	Is(v.String(status).Equal("running"))
func (validator *ValidatorString[T]) EqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyEqualTo, value, template...)

	return validator
}

// Validate if a string value is greater than to another. This function
// internally uses the golang `>` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).GreaterThan("ba"))
func (validator *ValidatorString[T]) GreaterThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringGreaterThan(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterThan, value, template...)

	return validator
}

// Validate if a string value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	section := "bc"
//	Is(v.String(section).GreaterOrEqualTo("bc"))

func (validator *ValidatorString[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringGreaterOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyGreaterOrEqualTo, value, template...)

	return validator
}

// Validate if a string value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).LessThan("bc"))
func (validator *ValidatorString[T]) LessThan(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringLessThan(validator.context.Value().(T), value)
		},
		ErrorKeyLessThan, value, template...)

	return validator
}

// Validate if a string value is less or equal to another. This function
// internally uses the golang `<=` operator to compare two strings.
// For example:
//
//	section := "bc"
//	Is(v.String(section).LessOrEqualTo("bc"))
func (validator *ValidatorString[T]) LessOrEqualTo(value T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringLessOrEqualTo(validator.context.Value().(T), value)
		},
		ErrorKeyLessOrEqualTo, value, template...)

	return validator
}

// Validate if a string value is empty. Empty will be false if the length
// of the string is greater than zero, even if the string has only spaces.
// For checking if the string has only spaces, uses the function `Blank()`
// instead.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be false
func (validator *ValidatorString[T]) Empty(template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			return isStringEmpty(validator.context.Value().(T))
		},
		ErrorKeyEmpty, template...)

	return validator
}

// Validate if a string value is blank. Blank will be true if the length
// of the string is zero or if the string only has spaces.
// instead.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be true
func (validator *ValidatorString[T]) Blank(template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			return isStringBlank(validator.context.Value().(T))
		},
		ErrorKeyBlank, template...)

	return validator
}

// Validate if a string value pass a custom function.
// For example:
//
//	status := ""
//	Is(v.String(status).Passing((v string) bool {
//		return v == getNewStatus()
//	})
func (validator *ValidatorString[T]) Passing(function func(v0 T) bool, template ...string) *ValidatorString[T] {
	validator.context.Add(
		func() bool {
			return function(validator.context.Value().(T))
		},
		ErrorKeyPassing, template...)

	return validator
}

// Validate if a string is present in an string slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.String(status).InSlice(validStatus))
func (validator *ValidatorString[T]) InSlice(slice []T, template ...string) *ValidatorString[T] {
	validator.context.AddWithValue(
		func() bool {
			return isStringInSlice(validator.context.Value().(T), slice)
		},
		ErrorKeyInSlice, validator.context.Value(), template...)

	return validator
}

// Validate if a string match a regular expression.
// For example:
//
//	status := "pre-approved"
//	regex, _ := regexp.Compile("pre-.+")
//	Is(v.String(status).MatchingTo(regex))
func (validator *ValidatorString[T]) MatchingTo(regex *regexp.Regexp, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringMatchingTo(validator.context.Value().(T), regex)
		},
		ErrorKeyMatchingTo,
		map[string]any{"title": validator.context.title, "regexp": regex},
		template...)

	return validator
}

// Validate the maximum length of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MaxLength(6))
func (validator *ValidatorString[T]) MaxLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringMaxLength(validator.context.Value().(T), length)
		},
		ErrorKeyMaxLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// Validate the minimum length of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MinLength(6))
func (validator *ValidatorString[T]) MinLength(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringMinLength(validator.context.Value().(T), length)
		},
		ErrorKeyMinLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// Validate the length of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).Length(6))
func (validator *ValidatorString[T]) Length(length int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringLength(validator.context.Value().(T), length)
		},
		ErrorKeyLength,
		map[string]any{"title": validator.context.title, "length": length},
		template...)

	return validator
}

// Validate if the length a string is in a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).LengthBetween(2,6))
func (validator *ValidatorString[T]) LengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringLengthBetween(validator.context.Value().(T), min, max)
		},
		ErrorKeyLengthBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}

// Validate if the value of a string is in a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).Between(2,6))
func (validator *ValidatorString[T]) Between(min T, max T, template ...string) *ValidatorString[T] {
	validator.context.AddWithParams(
		func() bool {
			return isStringBetween(validator.context.Value().(T), min, max)
		},
		ErrorKeyBetween,
		map[string]any{"title": validator.context.title, "min": min, "max": max},
		template...)

	return validator
}
