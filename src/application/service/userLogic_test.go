package service

import (
	"errors"
	"reflect"
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
	mockServ := &mockGetService{}

	var expectedResult *in.UserRespBody

	mockServ.On("GetById").Return(expectedResult, errors.New("Any Error"))

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockServ,
	}

	want, wantErr := userService.GetById("id")

	assert.Nil(want)
	assert.Error(wantErr)

}

func TestStore(t *testing.T) {
	assert := assert.New(t)
	mockGetSrv := &mockGetService{}
	mockStoreSrv := &mockStoreService{}

	userService := &userService{
		storeService:  mockStoreSrv,
		updateService: nil,
		deleteService: nil,
		getService:    mockGetSrv,
	}

	type calls struct {
		methodWant      string
		methodWantError string
	}

	type getEmailResults struct {
		exists bool
		resp   *in.UserRespBody
		err    error
	}

	type testCase struct {
		description           string
		expectedResult        *in.UserRespBody
		expectedError         error
		calls                 calls
		getEmailResults       getEmailResults
		reflectExpectedResult []reflect.Value
	}

	expetedResultOk1 := &in.UserRespBody{
		Id:        "378927492",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email@email.com",
	}

	for _, scenario := range []testCase{
		{
			description:           "store ok",
			expectedResult:        expetedResultOk1,
			expectedError:         nil,
			calls:                 calls{"Equal", "Nil"},
			getEmailResults:       getEmailResults{false, &in.UserRespBody{}, nil},
			reflectExpectedResult: []reflect.Value{reflect.ValueOf(&expetedResultOk1).Elem()},
		},

		{
			description:           "store error",
			expectedResult:        &in.UserRespBody{},
			expectedError:         errors.New("any Error"),
			calls:                 calls{"Nil", "Error"},
			getEmailResults:       getEmailResults{true, &in.UserRespBody{}, errors.New("any Error")},
			reflectExpectedResult: []reflect.Value{},
		},
	} {
		t.Run(scenario.description, func(*testing.T) {

			mockGetSrv.On("GetByEmail").Return(
				scenario.getEmailResults.exists,
				scenario.getEmailResults.resp,
				scenario.getEmailResults.err,
			)

			mockStoreSrv.On("Store").Return(
				scenario.expectedResult,
				scenario.expectedError,
			)

			req := &in.UserReqBody{
				FirstName: "first_name",
				LastName:  "last_name",
				Email:     "email@email.com",
				Password:  "pass123",
			}

			want, wantError := userService.Store(req)

			scenario.reflectExpectedResult = append(scenario.reflectExpectedResult, reflect.ValueOf(&want).Elem())
			reflect.ValueOf(assert).MethodByName(scenario.calls.methodWant).Call(scenario.reflectExpectedResult)

			reflect.ValueOf(assert).MethodByName(scenario.calls.methodWantError).Call(
				[]reflect.Value{reflect.ValueOf(&wantError).Elem()})

		})
	}
}

func TestUserService(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"GetAllOk":     testGetAllOk,
		"GetAllError":  testGetAllError,
		"GetByIdOk":    testGetByIdOk,
		"GetByIdError": testGetByIdError,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
