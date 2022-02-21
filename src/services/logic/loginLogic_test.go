package logic

import (
	"testing"
	"users/src/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Authenticate(email string, password string) (bool, *domain.User, error) {
	args := mock.Called()
	result := args.Get(1)
	return args.Bool(0), result.(*domain.User), args.Error(2)
}

func Test_loginService_Authenticate(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockRepository)
	user := domain.User{
		Id:           "3242425",
		FirstName:    "first_name",
		LastName:     "last_name",
		Email:        "any_emai@gmail.com",
		Password:     "anypass",
		IsAmbassador: true,
	}

	mockRepo.On("Authenticate").Return(true, &user, nil)

	testLoginService := NewLoginLogic(mockRepo)

	want, want1, wantErr := testLoginService.Authenticate("any_emai@gmail.com", "anypass")

	mockRepo.AssertExpectations(t)

	assert.True(want)
	assert.Same(want1, &user)
	assert.Nil(wantErr)
}
