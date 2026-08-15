package is

func NumberEqualTo[T Number](value, expected T) bool { return ComparableEqualTo(value, expected) }

func NumberGreaterThan[T Number](value, expected T) bool { return value > expected }

func NumberGreaterOrEqualTo[T Number](value, expected T) bool { return value >= expected }

func NumberLessThan[T Number](value, expected T) bool { return value < expected }

func NumberLessOrEqualTo[T Number](value, expected T) bool { return value <= expected }

func NumberBetween[T Number](value, min, max T) bool {
	return NumberGreaterOrEqualTo(value, min) && NumberLessOrEqualTo(value, max)
}

func NumberZero[T Number](value T) bool { return NumberEqualTo(value, 0) }

func NumberInSlice[T Number](value T, values []T) bool { return ComparableInSlice(value, values) }

func NumberPassing[T Number](value T, function func(T) bool) bool {
	return Passing(value, function)
}

func NumberPEqualTo[T Number](value *T, expected T) bool {
	return ComparablePEqualTo(value, expected)
}

func NumberPGreaterThan[T Number](value *T, expected T) bool {
	return value != nil && NumberGreaterThan(*value, expected)
}

func NumberPGreaterOrEqualTo[T Number](value *T, expected T) bool {
	return value != nil && NumberGreaterOrEqualTo(*value, expected)
}

func NumberPLessThan[T Number](value *T, expected T) bool {
	return value != nil && NumberLessThan(*value, expected)
}

func NumberPLessOrEqualTo[T Number](value *T, expected T) bool {
	return value != nil && NumberLessOrEqualTo(*value, expected)
}

func NumberPBetween[T Number](value *T, min, max T) bool {
	return NumberPGreaterOrEqualTo(value, min) && NumberPLessOrEqualTo(value, max)
}

func NumberPZero[T Number](value *T) bool { return NumberPEqualTo(value, 0) }

func NumberPZeroOrNil[T Number](value *T) bool { return NumberPNil(value) || NumberPZero(value) }

func NumberPInSlice[T Number](value *T, values []T) bool {
	return ComparablePInSlice(value, values)
}

func NumberPPassing[T Number](value *T, function func(*T) bool) bool {
	return Passing(value, function)
}

func NumberPNil[T Number](value *T) bool { return ComparablePNil(value) }
