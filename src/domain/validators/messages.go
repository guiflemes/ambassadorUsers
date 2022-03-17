package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type requiredMessageError struct{}

func (re *requiredMessageError) message(fe validator.FieldError) string {
	return "this field is required"
}

type emailMessageError struct{}

func (e *emailMessageError) message(fe validator.FieldError) string {
	return "invalid email"
}

type minMessageError struct{}

func (e *minMessageError) message(fe validator.FieldError) string {
	return fmt.Sprintf("the length should be > %s", fe.Param())
}
