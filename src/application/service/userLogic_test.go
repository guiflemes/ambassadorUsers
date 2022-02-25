package service

import (
	"testing"
	"users/src/application/port/in"
	"users/src/domain"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) GetAll() (domain.UsersList, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(domain.UsersList), args.Error(1)
}

func (mock *MockUserRepository) GetBy(filter map[string]interface{}) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}
func (mock *MockUserRepository) Store(data *domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}
func (mock *MockUserRepository) Update(data *domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}
func (mock *MockUserRepository) Delete(id string) error {
	args := mock.Called()
	return args.Error(0)
}

func testGetAllOk(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockUserRepository)

	mockUsers := domain.UsersList{
		{
			Id:           "378927492",
			FirstName:    "first_name",
			LastName:     "last_name",
			Email:        "email@email.com",
			Password:     "password",
			IsAmbassador: true,
		},
	}

	mockRepo.On("GetAll").Return(mockUsers, nil)

	userService := NewUserLogic(mockRepo)

	expectedResult := []*in.UserRespBody{
		{
			Id:        "378927492",
			FirstName: "first_name",
			LastName:  "last_name",
			Email:     "email@email.com",
		},
	}

	want, wantErr := userService.GetAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(want, expectedResult)
	assert.Nil(wantErr)

}

func testGetAllError(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockUserRepository)

	mockRepo.On("GetAll").Return(domain.UsersList{}, errors.New("error"))

	userService := NewUserLogic(mockRepo)

	want, wantErr := userService.GetAll()

	mockRepo.AssertExpectations(t)

	assert.Nil(want)
	assert.Error(wantErr)

}

func TestUserService(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"GetAllOk":    testGetAllOk,
		"GetAllError": testGetAllError,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
