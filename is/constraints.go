// Package is exposes Valgo's validation rules as stateless boolean functions.
//
// The functions in this package contain no validation context, error-message,
// or localization behavior. They are suitable for use from ordinary Go code
// and as building blocks for other validation libraries.
package is

// Number covers the numeric types supported by Valgo's Number validators.
type Number interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64 |
		~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64 |
		~float32 |
		~float64
}

// Int covers the signed integer types supported by Valgo's Int validators.
type Int interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64
}

// Uint covers the unsigned integer types supported by Valgo's Uint validators.
type Uint interface {
	~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64
}

// Float covers the floating-point types supported by Valgo.
type Float interface {
	~float32 | ~float64
}
