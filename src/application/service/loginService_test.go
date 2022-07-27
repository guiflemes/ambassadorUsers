package service

import (
	"errors"
	"testing"
	"users/src/application/port/in"
	"users/src/domain"
	"users/src/utils"

	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLoginRepository struct {
	mock.Mock
}

func (mock *MockLoginRepository) GetBy(context.Context, map[string]interface{}) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

func testLoginServiceAuthenticateOk(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockLoginRepository)
	user := domain.User{
		Id:        "3242425",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "any_emai@gmail.com",
		Password:  "anypass",
		IsActive:  true,
	}

	expectedResult := in.NewUserRespBody(&user)

	mockRepo.On("GetBy").Return(&user, nil)

	testLoginService := NewLoginService(mockRepo, func(password, userpass string) bool { return true })
	ctx := context.Background()

	want, want1, wantErr := testLoginService.Authenticate(ctx, "any_emai@gmail.com", "anypass")

	mockRepo.AssertExpectations(t)

	assert.True(want)
	assert.Equal(expectedResult, want1)
	assert.Nil(wantErr)
}

func testLoginServiceAuthenticateFailed(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockLoginRepository)
	user := domain.User{}

	mockRepo.On("GetBy").Return(&user, errors.New("failed to authenticate"))

	testLoginService := NewLoginService(mockRepo, func(password, userpass string) bool { return false })
	ctx := context.Background()

	want, want1, wantErr := testLoginService.Authenticate(ctx, "any_emai@gmail.com", "anypass")

	mockRepo.AssertExpectations(t)

	assert.False(want)
	assert.Nil(want1)
	assert.Error(wantErr)
}

func testLoginInvalidParameters(t *testing.T) {
	assert := assert.New(t)
	user := domain.User{}

	type testCase struct {
		description string
		email       string
		password    string
	}

	for _, scenario := range []testCase{
		{
			description: "empty email",
			email:       "",
			password:    "anypass",
		},
		{
			description: "empty password",
			email:       "myemail@email.com",
			password:    "",
		},
		{
			description: "email and password empty",
			email:       " ",
			password:    "",
		},
	} {
		t.Run(scenario.description, func(*testing.T) {

			mockRepo := new(MockLoginRepository)
			mockRepo.On("GetBy").Return(&user, errors.New("failed to authenticate"))
			testLoginService := NewLoginService(mockRepo, func(password, userpass string) bool { return false })
			ctx := context.Background()

			want, want1, wantErr := testLoginService.Authenticate(ctx, scenario.email, scenario.password)

			assert.False(want)
			assert.Nil(want1)
			assert.ErrorContains(wantErr, utils.ErrInvalidParameter.Error())
		})
	}

}

func TestLoginServiceAuthenticate(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Authenticate Ok":     testLoginServiceAuthenticateOk,
		"Authenticate Failed": testLoginServiceAuthenticateFailed,
		"InvalidParameters":   testLoginInvalidParameters,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
