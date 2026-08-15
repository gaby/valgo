package is

import "time"

func TimeEqualTo(value, expected time.Time) bool { return value.Equal(expected) }

func TimeAfter(value, expected time.Time) bool { return value.After(expected) }

func TimeAfterOrEqualTo(value, expected time.Time) bool {
	return TimeAfter(value, expected) || TimeEqualTo(value, expected)
}

func TimeBefore(value, expected time.Time) bool { return value.Before(expected) }

func TimeBeforeOrEqualTo(value, expected time.Time) bool {
	return TimeBefore(value, expected) || TimeEqualTo(value, expected)
}

func TimeBetween(value, min, max time.Time) bool {
	return TimeAfterOrEqualTo(value, min) && TimeBeforeOrEqualTo(value, max)
}

func TimeZero(value time.Time) bool { return value.IsZero() }

func TimeInSlice(value time.Time, values []time.Time) bool {
	for _, candidate := range values {
		if value.Equal(candidate) {
			return true
		}
	}
	return false
}

func TimePEqualTo(value *time.Time, expected time.Time) bool {
	return value != nil && TimeEqualTo(*value, expected)
}

func TimePAfter(value *time.Time, expected time.Time) bool {
	return value != nil && TimeAfter(*value, expected)
}

func TimePAfterOrEqualTo(value *time.Time, expected time.Time) bool {
	return value != nil && TimeAfterOrEqualTo(*value, expected)
}

func TimePBefore(value *time.Time, expected time.Time) bool {
	return value != nil && TimeBefore(*value, expected)
}

func TimePBeforeOrEqualTo(value *time.Time, expected time.Time) bool {
	return value != nil && TimeBeforeOrEqualTo(*value, expected)
}

func TimePBetween(value *time.Time, min, max time.Time) bool {
	return value != nil && TimeBetween(*value, min, max)
}

func TimePZero(value *time.Time) bool { return value != nil && TimeZero(*value) }

func TimePNilOrZero(value *time.Time) bool { return TimePNil(value) || TimePZero(value) }

func TimePInSlice(value *time.Time, values []time.Time) bool {
	return value != nil && TimeInSlice(*value, values)
}

func TimePNil(value *time.Time) bool { return value == nil }
