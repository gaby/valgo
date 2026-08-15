package is

func BoolEqualTo[T ~bool](value, expected T) bool { return ComparableEqualTo(value, expected) }

func BoolTrue[T ~bool](value T) bool { return bool(value) }

func BoolFalse[T ~bool](value T) bool { return !BoolTrue(value) }

func BoolInSlice[T ~bool](value T, values []T) bool { return ComparableInSlice(value, values) }

func BoolPEqualTo[T ~bool](value *T, expected T) bool {
	return ComparablePEqualTo(value, expected)
}

func BoolPTrue[T ~bool](value *T) bool { return value != nil && BoolTrue(*value) }

func BoolPFalse[T ~bool](value *T) bool { return value != nil && BoolFalse(*value) }

func BoolPFalseOrNil[T ~bool](value *T) bool { return BoolPNil(value) || BoolPFalse(value) }

func BoolPInSlice[T ~bool](value *T, values []T) bool {
	return ComparablePInSlice(value, values)
}

func BoolPNil[T ~bool](value *T) bool { return ComparablePNil(value) }
