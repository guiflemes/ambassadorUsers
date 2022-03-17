package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Validator struct {
	Tags []errorTag
}

func NewValidator() *Validator {
	return &Validator{
		Tags: tags,
	}
}

type ErrorStruct struct {
	Param   string
	Message string
}

func (v *Validator) ValidateAny(any interface{}) []ErrorStruct {
	validate = validator.New()
	err := validate.Struct(any)

	if err != nil {
		ve := err.(validator.ValidationErrors)
		out := make([]ErrorStruct, len(ve))
		for _, fe := range ve {
			out = append(out, ErrorStruct{fe.Field(), v.msgForTag(fe)})
		}
		return out
	}

	return nil
}

func (v *Validator) msgForTag(fe validator.FieldError) string {
	for _, tag := range v.Tags {
		if tag.name == fe.Tag() {
			return tag.fieldError.message(fe)
		}
	}
	return fe.Error()

}
