package service

import (
	"errors"
	"testing"
	"users/src/application/port/in"
	"users/src/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStoreService struct{ mock.Mock }

func (mock *mockStoreService) Store(*domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

type mockUpdateService struct{ mock.Mock }

func (mock *mockUpdateService) Update(data *domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

type mockDeleteService struct{ mock.Mock }

func (mock *mockDeleteService) Delete(id string) error {
	args := mock.Called()
	return args.Error(0)
}

type mockGetService struct{ mock.Mock }

func (mock *mockGetService) GetAll() (domain.UsersList, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(domain.UsersList), args.Error(1)
}
func (mock *mockGetService) GetById(id string) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}
func (mock *mockGetService) GetByEmail(email string) (bool, *domain.User, error) {
	args := mock.Called()
	result := args.Get(1)
	return args.Bool(0), result.(*domain.User), args.Error(2)
}

func testGetAllOk(t *testing.T) {
	assert := assert.New(t)
	mockServ := &mockGetService{}

	mockOn := domain.UsersList{
		&domain.User{
			Id:        "378927492",
			FirstName: "first_name",
			LastName:  "last_name",
			Email:     "email@email.com",
			Password:  "jdisjdijs",
			IsActive:  true,
		},
	}

	mockServ.On("GetAll").Return(mockOn, nil)

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockServ,
	}

	want, wantErr := userService.GetAll()

	mockServ.AssertExpectations(t)

	expectedResult := []*in.UserRespBody{
		{
			Id:        "378927492",
			FirstName: "first_name",
			LastName:  "last_name",
			Email:     "email@email.com",
		},
	}

	assert.Equal(want, expectedResult)
	assert.Nil(wantErr)

}

func testGetAllError(t *testing.T) {
	assert := assert.New(t)
	mockServ := &mockGetService{}

	var mockOn domain.UsersList

	mockServ.On("GetAll").Return(mockOn, errors.New("error"))

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockServ,
	}

	want, wantErr := userService.GetAll()

	mockServ.AssertExpectations(t)

	assert.Nil(want)
	assert.Error(wantErr)

}

func testGetByIdOk(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mockGetService{}

	mockOn := &domain.User{
		Id:        "378927492",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email@email.com",
		Password:  "jdisjdijs",
		IsActive:  true,
	}

	mockRepo.On("GetById").Return(mockOn, nil)

	userService := &userService{
		storeService:  nil,
		updateService: nil,
		deleteService: nil,
		getService:    mockRepo,
	}

	want, wantErr := userService.GetById("id")

	expectedResult := &in.UserRespBody{
		Id:        "378927492",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email@email.com",
	}

	assert.Equal(want, expectedResult)
	assert.Nil(wantErr)

}

func testGetByIdError(t *testing.T) {
	assert := assert.New(t)
	mockServ := &mockGetService{}

	var mockOn *domain.User

	mockServ.On("GetById").Return(mockOn, errors.New("Any Error"))

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

func testStore(t *testing.T) {
	assert := assert.New(t)

	type getEmailResults struct {
		exists bool
		resp   *domain.User
		err    error
	}

	type testCase struct {
		description     string
		expectedResult  *in.UserRespBody
		expectedError   error
		mockOn          *domain.User
		getEmailResults getEmailResults
	}

	mockStoreOk := &domain.User{
		Id:        "378927492",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email@email.com",
		Password:  "jdisjdijs",
		IsActive:  true,
	}

	for _, scenario := range []testCase{
		{
			description: "store ok",
			expectedResult: &in.UserRespBody{
				Id:        "378927492",
				FirstName: "first_name",
				LastName:  "last_name",
				Email:     "email@email.com",
			},
			mockOn:          mockStoreOk,
			expectedError:   nil,
			getEmailResults: getEmailResults{false, mockStoreOk, nil},
		},

		{
			description:     "store error",
			expectedResult:  nil,
			mockOn:          nil,
			expectedError:   errors.New("any Error"),
			getEmailResults: getEmailResults{true, &domain.User{}, errors.New("any Error")},
		},
	} {
		t.Run(scenario.description, func(*testing.T) {

			mockGetSrv := &mockGetService{}
			mockStoreSrv := &mockStoreService{}

			mockGetSrv.On("GetByEmail").Return(
				scenario.getEmailResults.exists,
				scenario.getEmailResults.resp,
				scenario.getEmailResults.err,
			)

			mockStoreSrv.On("Store").Return(
				scenario.mockOn,
				scenario.expectedError,
			)

			userService := &userService{
				storeService:  mockStoreSrv,
				updateService: nil,
				deleteService: nil,
				getService:    mockGetSrv,
			}

			req := &in.UserReqBody{
				FirstName: "first_name",
				LastName:  "last_name",
				Email:     "email@email.com",
				Password:  "pass123",
			}

			want, wantError := userService.Store(req)

			if !scenario.getEmailResults.exists {
				assert.Equal(want, scenario.expectedResult)
				assert.Nil(wantError)
				return
			}

			assert.Nil(want)
			assert.Error(wantError)

		})
	}
}

func testUpdate(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		description    string
		getByIdWant    *domain.User
		getByIdError   error
		updateWant     *domain.User
		updateError    error
		expectedResult *in.UserRespBody
	}

	for _, scenario := range []testCase{
		{
			description: "update ok",
			getByIdWant: &domain.User{
				Id:        "378927492",
				FirstName: "old_first_name",
				LastName:  "old_last_name",
				Email:     "old_email@email.com",
				Password:  "jdisjdijs",
				IsActive:  true,
			},
			getByIdError: nil,
			updateWant: &domain.User{
				Id:        "378927492",
				FirstName: "new_first_name",
				LastName:  "new_last_name",
				Email:     "new_email@email.com",
				Password:  "jdisjdijs",
				IsActive:  true,
			},
			updateError: nil,
			expectedResult: &in.UserRespBody{
				Id:        "378927492",
				FirstName: "new_first_name",
				LastName:  "new_last_name",
				Email:     "new_email@email.com",
			},
		},
		{
			description:    "update error",
			getByIdWant:    nil,
			getByIdError:   errors.New("any error"),
			updateWant:     nil,
			updateError:    errors.New("any error"),
			expectedResult: nil,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			mockGetSrv := &mockGetService{}
			mockUpdateSrv := &mockUpdateService{}

			mockGetSrv.On("GetById").Return(scenario.getByIdWant, scenario.getByIdError)
			mockUpdateSrv.On("Update").Return(scenario.updateWant, scenario.updateError)

			userSrv := &userService{
				storeService:  nil,
				updateService: mockUpdateSrv,
				deleteService: nil,
				getService:    mockGetSrv,
			}

			req := &in.UserUpdateReq{
				Id:        "378927492",
				FirstName: "new_first_name",
				LastName:  "new_last_name",
				Email:     "email@email.com",
			}

			want, wantErr := userSrv.Update(req)

			if wantErr != nil {
				assert.Nil(want)
				assert.Error(wantErr)
				return
			}

			assert.Equal(want, scenario.expectedResult)
			assert.Nil(wantErr)

		})
	}

}

func TestUserService(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"GetAllOk":     testGetAllOk,
		"GetAllError":  testGetAllError,
		"GetByIdOk":    testGetByIdOk,
		"GetByIdError": testGetByIdError,
		"Store":        testStore,
		"Update":       testUpdate,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
