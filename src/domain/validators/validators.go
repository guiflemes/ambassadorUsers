package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

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

func (e *ErrorStruct) Error() string {
	return e.Message
}

func (v *Validator) ValidateStruct(any interface{}) error {
	err := validate.Struct(any)

	if err != nil {
		fe := err.(validator.ValidationErrors)[0]
		return &ErrorStruct{fe.Field(), v.msgForTag(fe)}
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

// func (v *Validator) ValidateAny(any interface{}) []ErrorStruct {
// 	validate = validator.New()
// 	err := validate.Struct(any)

// 	if err != nil {
// 		ve := err.(validator.ValidationErrors)
// 		out := make([]ErrorStruct, len(ve))
// 		for _, fe := range ve {
// 			out = append(out, ErrorStruct{fe.Field(), v.msgForTag(fe)})
// 		}
// 		return out
// 	}

// 	return nil
// }
