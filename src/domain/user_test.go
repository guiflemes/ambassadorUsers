package domain

import (
	"errors"
	"reflect"
	"testing"
	"users/src/domain/validators"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {

	defer func() {
		Validator = validators.NewValidator().ValidateStruct
	}()

	assert := assert.New(t)

	type calls struct {
		methodName1 string
		methodName2 string
	}

	type testCase struct {
		description   string
		input         *User
		expectedBool  bool
		expectedError error
		calls         calls
	}

	for _, scenario := range []testCase{
		{
			description: "error",
			input: &User{
				Id:        "489284932",
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "Email",
				Password:  "Pass",
				IsActive:  true,
			},
			expectedBool:  false,
			expectedError: errors.New("any error"),
			calls:         calls{"False", "Error"},
		},
		{
			description: "ok",
			input: &User{
				Id:        "489284932",
				FirstName: "FirstName",
				LastName:  "LastName",
				Email:     "email@gmail.com",
				Password:  "Pass",
				IsActive:  true,
			},
			expectedBool:  true,
			expectedError: nil,
			calls:         calls{"True", "Nil"},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			Validator = func(any interface{}) error {
				return scenario.expectedError
			}

			is, err := scenario.input.IsValid()

			reflect.ValueOf(assert).MethodByName(scenario.calls.methodName1).Call(
				[]reflect.Value{reflect.ValueOf(&is).Elem()})

			reflect.ValueOf(assert).MethodByName(scenario.calls.methodName2).Call(
				[]reflect.Value{reflect.ValueOf(&err).Elem()})

		})

	}
}
