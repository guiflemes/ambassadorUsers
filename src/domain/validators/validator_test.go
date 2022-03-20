package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type item struct {
	Title  string `validate:"required"`
	Email  string `validate:"required,email"`
	Street string `validate:"required,gt=2"`
}

func TestValidateStruct(t *testing.T) {
	assert := assert.New(t)
	validator := NewValidator()

	type errorTestCase struct {
		description string
		input       *item
		expectError *ErrorStruct
	}

	for _, scenario := range []errorTestCase{
		{
			description: "email error",
			input: &item{
				Title:  "title",
				Email:  "invalid email",
				Street: "Street",
			},
			expectError: &ErrorStruct{
				"Email",
				"invalid email",
			},
		},

		{
			description: "required field error (Title)",
			input: &item{
				Title:  "",
				Email:  "any_valid_email@gmail.com",
				Street: "Street",
			},
			expectError: &ErrorStruct{
				"Title",
				"this field is required",
			},
		},

		{
			description: "len smaller than expected (Street)",
			input: &item{
				Title:  "title",
				Email:  "any_valid_email@gmail.com",
				Street: "ST",
			},
			expectError: &ErrorStruct{
				"Street",
				"the length should be greater than 2",
			},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			actual := validator.ValidateStruct(scenario.input)
			assert.Equal(scenario.expectError, actual)
		})

	}

}
