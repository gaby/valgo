package is

func UintEqualTo[T Uint](value, expected T) bool { return NumberEqualTo(value, expected) }

func UintGreaterThan[T Uint](value, expected T) bool { return NumberGreaterThan(value, expected) }

func UintGreaterOrEqualTo[T Uint](value, expected T) bool {
	return NumberGreaterOrEqualTo(value, expected)
}

func UintLessThan[T Uint](value, expected T) bool { return NumberLessThan(value, expected) }

func UintLessOrEqualTo[T Uint](value, expected T) bool {
	return NumberLessOrEqualTo(value, expected)
}

func UintBetween[T Uint](value, min, max T) bool { return NumberBetween(value, min, max) }

func UintZero[T Uint](value T) bool { return NumberZero(value) }

func UintInSlice[T Uint](value T, values []T) bool { return NumberInSlice(value, values) }

func UintPassing[T Uint](value T, function func(T) bool) bool {
	return NumberPassing(value, function)
}

func UintPEqualTo[T Uint](value *T, expected T) bool {
	return NumberPEqualTo(value, expected)
}

func UintPGreaterThan[T Uint](value *T, expected T) bool {
	return NumberPGreaterThan(value, expected)
}

func UintPGreaterOrEqualTo[T Uint](value *T, expected T) bool {
	return NumberPGreaterOrEqualTo(value, expected)
}

func UintPLessThan[T Uint](value *T, expected T) bool {
	return NumberPLessThan(value, expected)
}

func UintPLessOrEqualTo[T Uint](value *T, expected T) bool {
	return NumberPLessOrEqualTo(value, expected)
}

func UintPBetween[T Uint](value *T, min, max T) bool {
	return NumberPBetween(value, min, max)
}

func UintPZero[T Uint](value *T) bool { return NumberPZero(value) }

func UintPInSlice[T Uint](value *T, values []T) bool { return NumberPInSlice(value, values) }

func UintPZeroOrNil[T Uint](value *T) bool { return NumberPZeroOrNil(value) }

func UintPNil[T Uint](value *T) bool { return NumberPNil(value) }

func UintPPassing[T Uint](value *T, function func(*T) bool) bool {
	return NumberPPassing(value, function)
}
