package service

import (
	"errors"
	"testing"
	"users/src/application/port/in"
	"users/src/domain"

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

type mockStoreService struct{ mock.Mock }

func (mock *mockStoreService) Store(*domain.User) (*in.UserRespBody, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*in.UserRespBody), args.Error(1)
}

type mockUpdateService struct{ mock.Mock }

func (mock *mockUpdateService) Update(data *domain.User) (*in.UserRespBody, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*in.UserRespBody), args.Error(1)
}

type mockDeleteService struct{ mock.Mock }

func (mock *mockDeleteService) Delete(id string) error {
	args := mock.Called()
	return args.Error(0)
}

type mockGetService struct{ mock.Mock }

func (mock *mockGetService) GetAll() ([]*in.UserRespBody, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]*in.UserRespBody), args.Error(1)
}
func (mock *mockGetService) GetById(id string) (*in.UserRespBody, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*in.UserRespBody), args.Error(1)
}
func (mock *mockGetService) GetByEmail(email string) (bool, *in.UserRespBody, error) {
	args := mock.Called()
	result := args.Get(1)
	return args.Bool(0), result.(*in.UserRespBody), args.Error(2)
}

func testGetAllOk(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mockGetService{}

	expectedResult := []*in.UserRespBody{
		{
			Id:        "378927492",
			FirstName: "first_name",
			LastName:  "last_name",
			Email:     "email@email.com",
		},
	}

	mockRepo.On("GetAll").Return(expectedResult, nil)

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockRepo,
	}

	want, wantErr := userService.GetAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(want, expectedResult)
	assert.Nil(wantErr)

}

func testGetAllError(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mockGetService{}

	var expectedResult []*in.UserRespBody

	mockRepo.On("GetAll").Return(expectedResult, errors.New("error"))

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockRepo,
	}

	want, wantErr := userService.GetAll()

	mockRepo.AssertExpectations(t)

	assert.Nil(want)
	assert.Error(wantErr)

}

func testGetByIdOk(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mockGetService{}

	expectedResult := &in.UserRespBody{
		Id:        "378927492",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email@email.com",
	}

	mockRepo.On("GetById").Return(expectedResult, nil)

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockRepo,
	}

	want, wantErr := userService.GetById("id")

	assert.Equal(want, expectedResult)
	assert.Nil(wantErr)

}

func testGetByIdError(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mockGetService{}

	var expectedResult *in.UserRespBody

	mockRepo.On("GetById").Return(expectedResult, errors.New("Any Error"))

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockRepo,
	}

	want, wantErr := userService.GetById("id")

	assert.Nil(want)
	assert.Error(wantErr)

}

// func testUpdateOk(t *testing.T) {
// 	assert := assert.New(t)
// 	mockRepo := new(MockUserRepository)

// 	mockDomain := domain.User{
// 		Id:        "id",
// 		FirstName: "first_name",
// 		LastName:  "last_name",
// 		Email:     "email",
// 		Password:  "anypass",
// 		IsActive:  true,
// 	}

// 	mockRepo.On("Update").Return(&mockDomain, nil)

// 	userService := NewUserService(mockRepo)

// 	reqBody := in.UserUpdateReq{
// 		Id:        "anyID",
// 		FirstName: "first_name",
// 		LastName:  "last_name",
// 		Email:     "email",
// 	}
// 	want, wantErr := userService.Update(&reqBody)

// 	expectedResult := in.NewUserRespBody(&mockDomain)

// 	assert.Equal(want, expectedResult)
// 	assert.Nil(wantErr)

// }

// func testUpdateError(t *testing.T) {
// 	assert := assert.New(t)
// 	mockRepo := new(MockUserRepository)

// 	mockDomain := domain.User{}
// 	mockRepo.On("Update").Return(&mockDomain, errors.New("any error"))

// 	userService := NewUserService(mockRepo)

// 	reqBody := in.UserUpdateReq{
// 		Id:        "anyID",
// 		FirstName: "first_name",
// 		LastName:  "last_name",
// 		Email:     "email",
// 	}

// 	want, wantErr := userService.Update(&reqBody)

// 	assert.Nil(want)
// 	assert.Error(wantErr)

// }

func TestUserService(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"GetAllOk":     testGetAllOk,
		"GetAllError":  testGetAllError,
		"GetByIdOk":    testGetByIdOk,
		"GetByIdError": testGetByIdError,
		// "UpdateOk":     testUpdateOk,
		// "UpdateError":  testUpdateError,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
