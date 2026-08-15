package is

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func StringEqualTo[T ~string](value, expected T) bool { return ComparableEqualTo(value, expected) }

func StringGreaterThan[T ~string](value, expected T) bool { return value > expected }

func StringGreaterOrEqualTo[T ~string](value, expected T) bool { return value >= expected }

func StringLessThan[T ~string](value, expected T) bool { return value < expected }

func StringLessOrEqualTo[T ~string](value, expected T) bool { return value <= expected }

func StringBetween[T ~string](value, min, max T) bool {
	return StringGreaterOrEqualTo(value, min) && StringLessOrEqualTo(value, max)
}

func StringEmpty[T ~string](value T) bool { return len(value) == 0 }

func StringBlank[T ~string](value T) bool { return len(strings.TrimSpace(string(value))) == 0 }

func StringInSlice[T ~string](value T, values []T) bool {
	return ComparableInSlice(value, values)
}

func StringMatchingTo[T ~string](value T, regex *regexp.Regexp) bool {
	return regex.MatchString(string(value))
}

func StringMaxBytes[T ~string](value T, length int) bool { return len(value) <= length }

func StringMinBytes[T ~string](value T, length int) bool { return len(value) >= length }

func StringByteLength[T ~string](value T, length int) bool { return len(value) == length }

func StringByteLengthBetween[T ~string](value T, min, max int) bool {
	return len(value) >= min && len(value) <= max
}

func StringMaxLength[T ~string](value T, length int) bool {
	return utf8.RuneCountInString(string(value)) <= length
}

func StringMinLength[T ~string](value T, length int) bool {
	return utf8.RuneCountInString(string(value)) >= length
}

func StringLength[T ~string](value T, length int) bool {
	return utf8.RuneCountInString(string(value)) == length
}

func StringLengthBetween[T ~string](value T, min, max int) bool {
	length := utf8.RuneCountInString(string(value))
	return length >= min && length <= max
}

func StringPEqualTo[T ~string](value *T, expected T) bool {
	return ComparablePEqualTo(value, expected)
}

func StringPGreaterThan[T ~string](value *T, expected T) bool {
	return value != nil && StringGreaterThan(*value, expected)
}

func StringPGreaterOrEqualTo[T ~string](value *T, expected T) bool {
	return value != nil && StringGreaterOrEqualTo(*value, expected)
}

func StringPLessThan[T ~string](value *T, expected T) bool {
	return value != nil && StringLessThan(*value, expected)
}

func StringPLessOrEqualTo[T ~string](value *T, expected T) bool {
	return value != nil && StringLessOrEqualTo(*value, expected)
}

func StringPBetween[T ~string](value *T, min, max T) bool {
	return value != nil && StringBetween(*value, min, max)
}

func StringPEmpty[T ~string](value *T) bool { return value != nil && StringEmpty(*value) }

func StringPEmptyOrNil[T ~string](value *T) bool { return StringPNil(value) || StringPEmpty(value) }

func StringPBlank[T ~string](value *T) bool { return value != nil && StringBlank(*value) }

func StringPBlankOrNil[T ~string](value *T) bool { return StringPNil(value) || StringPBlank(value) }

func StringPInSlice[T ~string](value *T, values []T) bool {
	return ComparablePInSlice(value, values)
}

func StringPMatchingTo[T ~string](value *T, regex *regexp.Regexp) bool {
	return value != nil && StringMatchingTo(*value, regex)
}

func StringPMaxBytes[T ~string](value *T, length int) bool {
	return value != nil && StringMaxBytes(*value, length)
}

func StringPMinBytes[T ~string](value *T, length int) bool {
	return value != nil && StringMinBytes(*value, length)
}

func StringPByteLength[T ~string](value *T, length int) bool {
	return value != nil && StringByteLength(*value, length)
}

func StringPByteLengthBetween[T ~string](value *T, min, max int) bool {
	return value != nil && StringByteLengthBetween(*value, min, max)
}

func StringPMaxLength[T ~string](value *T, length int) bool {
	return value != nil && StringMaxLength(*value, length)
}

func StringPMinLength[T ~string](value *T, length int) bool {
	return value != nil && StringMinLength(*value, length)
}

func StringPLength[T ~string](value *T, length int) bool {
	return value != nil && StringLength(*value, length)
}

func StringPLengthBetween[T ~string](value *T, min, max int) bool {
	return value != nil && StringLengthBetween(*value, min, max)
}

func StringPNil[T ~string](value *T) bool { return ComparablePNil(value) }
