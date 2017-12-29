package valgo

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasttemplate"
)

type Validator struct {
	currentValue interface{}
	currentTitle string
	currentName  string
	currentValid bool
	currentIndex int
	currentError *Error

	_locale locale
	valid   bool
	errors  []Error
}

func (validator *Validator) Is(value interface{}) *Validator {
	validator.currentIndex += 1
	validator.currentValue = value
	validator.currentValid = true
	validator.currentName = fmt.Sprintf("value%v", validator.currentIndex)
	validator.currentTitle = validator.currentName

	return validator
}

func (validator *Validator) Named(name string) *Validator {
	validator.currentName = name

	return validator
}

func (validator *Validator) Titled(title string) *Validator {
	validator.currentTitle = title

	return validator
}

func (validator *Validator) Valid() bool {
	return validator.valid
}

func (validator *Validator) Errors() []Error {
	return validator.errors
}

func (validator *Validator) ensureString() string {
	cv := validator.currentValue

	switch v := validator.currentValue.(type) {
	case uint8, uint16, uint32, uint64:
		return strconv.FormatUint(cv.(uint64), 10)
	case int8, int16, int32, int64:
		return strconv.FormatInt(cv.(int64), 10)
	case float32, float64:
		return strconv.FormatFloat(cv.(float64), 'f', -1, 64)
	case string:
		return cv.(string)
	default:
		fmt.Printf("unexpected type %T", v)
		return ""
	}
}

func (validator *Validator) invalidate(key string, values map[string]interface{}) {
	templateString := validator._locale.messages[key]
	template := fasttemplate.New(templateString, "{{", "}}")
	message := template.ExecuteString(values)

	if validator.currentError == nil {
		validator.currentError = &Error{
			Name:  validator.currentName,
			Title: validator.currentTitle,
			Value: validator.currentValue,
		}

		validator.currentError.Messages = []string{message}
		validator.currentValid = false
		validator.valid = false

		if validator.errors == nil {
			validator.errors = []Error{*validator.currentError}
		} else {
			validator.errors = append(validator.errors, *validator.currentError)
		}
	} else {
		validator.currentError.Messages = append(
			validator.currentError.Messages, message)
	}
}
