package validators

import (
	"github.com/go-playground/validator/v10"
)

var tags = []errorTag{
	errorTag{"required", &requiredMessageError{}},
	errorTag{"email", &emailMessageError{}},
	errorTag{"min", &minMessageError{}},
}

type errorTag struct {
	name       string
	fieldError errorMessage
}

type errorMessage interface {
	message(validator.FieldError) string
}
