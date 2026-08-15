package is

import "math"

func FloatEqualTo[T Float](value, expected T) bool { return NumberEqualTo(value, expected) }

func FloatGreaterThan[T Float](value, expected T) bool { return NumberGreaterThan(value, expected) }

func FloatGreaterOrEqualTo[T Float](value, expected T) bool {
	return NumberGreaterOrEqualTo(value, expected)
}

func FloatLessThan[T Float](value, expected T) bool { return NumberLessThan(value, expected) }

func FloatLessOrEqualTo[T Float](value, expected T) bool {
	return NumberLessOrEqualTo(value, expected)
}

func FloatBetween[T Float](value, min, max T) bool { return NumberBetween(value, min, max) }

func FloatZero[T Float](value T) bool { return NumberZero(value) }

func FloatPositive[T Float](value T) bool { return NumberGreaterThan(value, 0) }

func FloatNegative[T Float](value T) bool { return NumberLessThan(value, 0) }

func FloatInSlice[T Float](value T, values []T) bool { return NumberInSlice(value, values) }

func FloatPassing[T Float](value T, function func(T) bool) bool {
	return NumberPassing(value, function)
}

func FloatNaN[T Float](value T) bool { return math.IsNaN(float64(value)) }

func FloatInfinite[T Float](value T) bool { return math.IsInf(float64(value), 0) }

func FloatFinite[T Float](value T) bool {
	return !FloatNaN(value) && !FloatInfinite(value)
}

func FloatPEqualTo[T Float](value *T, expected T) bool {
	return NumberPEqualTo(value, expected)
}

func FloatPGreaterThan[T Float](value *T, expected T) bool {
	return NumberPGreaterThan(value, expected)
}

func FloatPGreaterOrEqualTo[T Float](value *T, expected T) bool {
	return NumberPGreaterOrEqualTo(value, expected)
}

func FloatPLessThan[T Float](value *T, expected T) bool {
	return NumberPLessThan(value, expected)
}

func FloatPLessOrEqualTo[T Float](value *T, expected T) bool {
	return NumberPLessOrEqualTo(value, expected)
}

func FloatPBetween[T Float](value *T, min, max T) bool {
	return NumberPBetween(value, min, max)
}

func FloatPZero[T Float](value *T) bool { return NumberPZero(value) }

func FloatPPositive[T Float](value *T) bool { return NumberPGreaterThan(value, 0) }

func FloatPNegative[T Float](value *T) bool { return NumberPLessThan(value, 0) }

func FloatPInSlice[T Float](value *T, values []T) bool { return NumberPInSlice(value, values) }

func FloatPZeroOrNil[T Float](value *T) bool { return NumberPZeroOrNil(value) }

func FloatPNil[T Float](value *T) bool { return NumberPNil(value) }

func FloatPPassing[T Float](value *T, function func(*T) bool) bool {
	return NumberPPassing(value, function)
}

func FloatPNaN[T Float](value *T) bool { return value != nil && FloatNaN(*value) }

func FloatPInfinite[T Float](value *T) bool { return value != nil && FloatInfinite(*value) }

func FloatPFinite[T Float](value *T) bool { return value != nil && FloatFinite(*value) }
