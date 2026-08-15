package is

func IntEqualTo[T Int](value, expected T) bool { return NumberEqualTo(value, expected) }

func IntGreaterThan[T Int](value, expected T) bool { return NumberGreaterThan(value, expected) }

func IntGreaterOrEqualTo[T Int](value, expected T) bool {
	return NumberGreaterOrEqualTo(value, expected)
}

func IntLessThan[T Int](value, expected T) bool { return NumberLessThan(value, expected) }

func IntLessOrEqualTo[T Int](value, expected T) bool {
	return NumberLessOrEqualTo(value, expected)
}

func IntBetween[T Int](value, min, max T) bool { return NumberBetween(value, min, max) }

func IntZero[T Int](value T) bool { return NumberZero(value) }

func IntPositive[T Int](value T) bool { return NumberGreaterThan(value, 0) }

func IntNegative[T Int](value T) bool { return NumberLessThan(value, 0) }

func IntInSlice[T Int](value T, values []T) bool { return NumberInSlice(value, values) }

func IntPassing[T Int](value T, function func(T) bool) bool {
	return NumberPassing(value, function)
}

func IntPEqualTo[T Int](value *T, expected T) bool {
	return NumberPEqualTo(value, expected)
}

func IntPGreaterThan[T Int](value *T, expected T) bool {
	return NumberPGreaterThan(value, expected)
}

func IntPGreaterOrEqualTo[T Int](value *T, expected T) bool {
	return NumberPGreaterOrEqualTo(value, expected)
}

func IntPLessThan[T Int](value *T, expected T) bool {
	return NumberPLessThan(value, expected)
}

func IntPLessOrEqualTo[T Int](value *T, expected T) bool {
	return NumberPLessOrEqualTo(value, expected)
}

func IntPBetween[T Int](value *T, min, max T) bool {
	return NumberPBetween(value, min, max)
}

func IntPZero[T Int](value *T) bool { return NumberPZero(value) }

func IntPPositive[T Int](value *T) bool { return NumberPGreaterThan(value, 0) }

func IntPNegative[T Int](value *T) bool { return NumberPLessThan(value, 0) }

func IntPInSlice[T Int](value *T, values []T) bool { return NumberPInSlice(value, values) }

func IntPZeroOrNil[T Int](value *T) bool { return NumberPZeroOrNil(value) }

func IntPNil[T Int](value *T) bool { return NumberPNil(value) }

func IntPPassing[T Int](value *T, function func(*T) bool) bool {
	return NumberPPassing(value, function)
}
