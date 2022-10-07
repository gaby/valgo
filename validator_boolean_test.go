package valgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorBoolNot(t *testing.T) {

	v := Is(Bool(true).Not().EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolEqualToWhenIsValid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(true).EqualTo(true))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	v = Is(Bool(false).EqualTo(false))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = Is(Bool(mybool1).EqualTo(mybool2))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolEqualToWhenIsInvalid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(true).EqualTo(false))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 must be equal to \"false\"",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = false

	v = Is(Bool(mybool1).EqualTo(mybool2))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 must be equal to \"false\"",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolTrueWhenIsValid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(true).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = Is(Bool(mybool1).True())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolTrueWhenIsInvalid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(false).True())
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 must be true",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(Bool(mybool1).True())
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 must be true",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolFalseWhenIsValid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(false).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(Bool(mybool1).False())
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolFalseWhenIsInvalid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(true).False())
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 must be false",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true

	v = Is(Bool(mybool1).False())
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 must be false",
		v.Errors()["value_0"].Messages()[0])
}

func TestValidatorBoolPassingWhenIsValid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(true).Passing(func(val bool) bool {
		return val == true
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = true
	var mybool2 MyBool = true

	v = Is(Bool(mybool1).Passing(func(val MyBool) bool {
		return val == mybool2
	}))
	assert.True(t, v.Valid())
	assert.Empty(t, v.Errors())
}

func TestValidatorBoolPassingWhenIsInvalid(t *testing.T) {
	ResetMessages()

	var v *ValidatorGroup

	v = Is(Bool(false).Passing(func(val bool) bool {
		return val == true
	}))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])

	// Custom Type
	type MyBool bool
	var mybool1 MyBool = false

	v = Is(Bool(mybool1).Passing(func(val MyBool) bool {
		return val == true
	}))
	assert.False(t, v.Valid())
	assert.NotEmpty(t, v.Errors())
	assert.Contains(t,
		v.Errors()["value_0"].Messages(),
		"Value 0 is not valid",
		v.Errors()["value_0"].Messages()[0])
}
