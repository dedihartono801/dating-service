package user

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	mock_repository "dating-service/internal/app/repository/user/mocks"
	"dating-service/internal/entity"
	"dating-service/pkg/constant"
	"dating-service/pkg/identifier"
	"dating-service/pkg/validator"

	"dating-service/pkg/helper"

	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())

	mockRepository := mock_repository.NewMockUserRepository(ctrl)
	feature := NewUserService(mockRepository, validator, identifier) // Adjust as needed

	ctx := context.Background()

	t.Run("Successful - Premium User", func(t *testing.T) {
		// Mock repository calls for a premium user
		mockRepository.EXPECT().CheckPremium(ctx).Return(&entity.GetUserDetailResponse{IsPremium: true}, nil)
		mockRepository.EXPECT().GetUserList(ctx).Return([]*entity.GetUserListResponse{
			{
				Id:         "1",
				FirstName:  "User1",
				LastName:   "User1",
				Bio:        "User1",
				Location:   "User1",
				Age:        1,
				Gender:     "User1",
				IsVerified: true,
				IsPremium:  true,
			},
		}, nil)

		users, err := feature.UserList(ctx)

		assert.Nil(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, "User1", users[0].FirstName)
	})

	t.Run("Successful - Non-premium User with <10 Swipes", func(t *testing.T) {
		// Mock repository calls for a non-premium user with less than 10 swipes
		mockRepository.EXPECT().CheckPremium(ctx).Return(&entity.GetUserDetailResponse{IsPremium: false}, nil)
		mockRepository.EXPECT().CountUserSwipped(ctx).Return(5, nil)
		mockRepository.EXPECT().GetUserList(ctx).Return([]*entity.GetUserListResponse{
			{
				Id:         "1",
				FirstName:  "User1",
				LastName:   "User1",
				Bio:        "User1",
				Location:   "User1",
				Age:        1,
				Gender:     "User1",
				IsVerified: true,
				IsPremium:  false,
			},
		}, nil)

		users, err := feature.UserList(ctx)

		assert.Nil(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, "User1", users[0].FirstName)
	})

	t.Run("Conflict Error - Non-premium User with >10 Swipes", func(t *testing.T) {
		// Mock repository calls for a non-premium user with more than 10 swipes
		mockRepository.EXPECT().CheckPremium(ctx).Return(&entity.GetUserDetailResponse{IsPremium: false}, nil)
		mockRepository.EXPECT().CountUserSwipped(ctx).Return(11, nil)

		users, err := feature.UserList(ctx)

		assert.Nil(t, users)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), constant.ErrorAlreadySwiped)
	})

	t.Run("Error - CheckPremium Failure", func(t *testing.T) {
		// Mock repository call that returns an error on CheckPremium
		mockRepository.EXPECT().CheckPremium(ctx).Return(nil, errors.New("database error"))

		users, err := feature.UserList(ctx)

		assert.Nil(t, users)
		assert.NotNil(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("Error - GetUserList Failure", func(t *testing.T) {
		// Mock repository calls for a premium user but return an error on GetUserList
		mockRepository.EXPECT().CheckPremium(ctx).Return(&entity.GetUserDetailResponse{IsPremium: true}, nil)
		mockRepository.EXPECT().GetUserList(ctx).Return(nil, errors.New("database error"))

		users, err := feature.UserList(ctx)

		assert.Nil(t, users)
		assert.NotNil(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockUserRepository(ctrl)
	validator := validator.NewValidator(validatorv10.New())
	identifier := identifier.NewIdentifier()
	svc := NewUserService(mockRepository, validator, identifier)

	ctx := context.Background()

	t.Run("Validation Error", func(t *testing.T) {
		request := &entity.LoginRequest{} // Missing required fields

		resp, err := svc.Login(ctx, request)

		assert.Nil(t, resp)
		assert.NotNil(t, err)
	})

	t.Run("Email Not Found", func(t *testing.T) {
		request := &entity.LoginRequest{
			Email:    "notfound@example.com",
			Password: "password123",
		}

		// Mock repository call to simulate user not found
		mockRepository.EXPECT().GetUserByEmail(ctx, request.Email).Return(nil, helper.Error(http.StatusNotFound, "User not found", nil))

		resp, err := svc.Login(ctx, request)

		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), constant.ErrorEmailNotFound)
	})

	t.Run("Incorrect Password", func(t *testing.T) {
		request := &entity.LoginRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}

		user := &entity.GetUserDetailResponse{
			Id:       "1",
			Email:    request.Email,
			Password: helper.EncryptPassword("correctpassword"),
		}

		// Mock repository call to return the user
		mockRepository.EXPECT().GetUserByEmail(ctx, request.Email).Return(user, nil)

		resp, err := svc.Login(ctx, request)

		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), constant.ErrorPasswordWrong)
	})

	t.Run("Successful Login", func(t *testing.T) {
		request := &entity.LoginRequest{
			Email:    "test@example.com",
			Password: "correctpassword",
		}

		user := &entity.GetUserDetailResponse{
			Id:         "1",
			Email:      request.Email,
			Password:   helper.EncryptPassword(request.Password),
			Gender:     "male",
			IsVerified: true,
			IsPremium:  true,
		}

		mockRepository.EXPECT().GetUserByEmail(ctx, request.Email).Return(user, nil)

		resp, err := svc.Login(ctx, request)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.ExpiredAt)
	})
}

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.HeaderContext, entity.ValueContext{
		UserId:     "1",
		Gender:     "male",
		Email:      "contoh@gmail.com",
		IsVerified: true,
		IsPremium:  true,
	})

	// Create mock dependencies
	mockRepository := mock_repository.NewMockUserRepository(ctrl)
	validator := validator.NewValidator(validatorv10.New())

	// Instantiate the service with dependencies
	svc := NewUserService(mockRepository, validator, nil)

	t.Run("Validation Error", func(t *testing.T) {
		request := &entity.SignupRequest{} // Empty request to trigger validation error

		err := svc.Signup(ctx, request)

		assert.NotNil(t, err)
	})

	t.Run("Email Already Exists", func(t *testing.T) {
		request := &entity.SignupRequest{
			Email:       "existing@example.com",
			Password:    "password123",
			FirstName:   "John",
			LastName:    "Doe",
			Gender:      "male",
			Age:         30,
			DateOfBirth: "1990-01-01",
			Bio:         "I'm a software developer",
			Location:    "New York",
		}

		// Mock repository call to return a user for the email check
		mockRepository.EXPECT().GetUserByEmail(ctx, request.Email).Return(&entity.GetUserDetailResponse{
			Id:         "1",
			Email:      request.Email,
			Password:   helper.EncryptPassword("password123"),
			Gender:     "male",
			IsVerified: true,
			IsPremium:  true,
		}, nil)

		err := svc.Signup(ctx, request)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), constant.ErrorEmailAlreadyExists)
	})

	t.Run("Repository Error on GetUserByEmail", func(t *testing.T) {
		request := &entity.SignupRequest{
			Email:       "existing@example.com",
			Password:    "password123",
			FirstName:   "John",
			LastName:    "Doe",
			Gender:      "male",
			Age:         30,
			DateOfBirth: "1990-01-01",
			Bio:         "I'm a software developer",
			Location:    "New York",
		}

		// Mock repository call to return an unexpected error
		mockRepository.EXPECT().GetUserByEmail(ctx, request.Email).Return(nil, helper.Error(http.StatusInternalServerError, "database error", nil))

		err := svc.Signup(ctx, request)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "database error")
	})

	t.Run("Successful Signup", func(t *testing.T) {
		request := &entity.SignupRequest{
			Email:       "newuser@example.com",
			Password:    "password123",
			FirstName:   "John",
			LastName:    "Doe",
			Gender:      "male",
			Age:         30,
			DateOfBirth: "1990-01-01",
			Bio:         "I'm a software developer",
			Location:    "New York",
		}

		// Encrypt the password in the request to match the expected format
		encryptedPassword := helper.EncryptPassword(request.Password)
		expectedRequest := *request
		expectedRequest.Password = encryptedPassword

		// Mock GetUserByEmail to return a not found error to simulate new user
		mockRepository.EXPECT().GetUserByEmail(gomock.Any(), request.Email).Return(nil, helper.Error(http.StatusNotFound, "User not found", nil))

		// Use gomock.Eq to match the exact fields of SignupRequest
		mockRepository.EXPECT().Signup(gomock.Any(), gomock.Eq(expectedRequest)).Return(nil)

		err := svc.Signup(ctx, request)

		assert.Nil(t, err)
	})
}

func TestSwipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.HeaderContext, entity.ValueContext{
		UserId:     "1",
		Gender:     "male",
		Email:      "contoh@gmail.com",
		IsVerified: true,
		IsPremium:  true,
	})

	mockRepository := mock_repository.NewMockUserRepository(ctrl)
	validator := validator.NewValidator(validatorv10.New())
	svc := NewUserService(mockRepository, validator, nil)

	t.Run("Validation error", func(t *testing.T) {
		// Create an invalid request (e.g., missing required fields)
		request := &entity.SwipeRequest{}

		err := svc.Swipe(ctx, request)

		assert.NotNil(t, err)
	})

	t.Run("User already swiped", func(t *testing.T) {
		request := &entity.SwipeRequest{
			TargetUserId: 2,
			SwipeType:    constant.Like,
		}

		// Mock premium check, swipe count, and existing swipe check
		mockRepository.EXPECT().CheckPremium(ctx).Return(&entity.GetUserDetailResponse{IsPremium: true}, nil)
		mockRepository.EXPECT().CheckUserSwipped(ctx, *request).Return(true, nil)

		err := svc.Swipe(ctx, request)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), constant.ErrorAlreadySwipedPerson)
	})

	t.Run("Successful swipe with match", func(t *testing.T) {
		request := &entity.SwipeRequest{
			TargetUserId: 2,
			SwipeType:    constant.Like,
		}

		tx := &sql.Tx{} // Mock transaction object

		// Mock the complete flow for a successful swipe and match creation
		mockRepository.EXPECT().CheckPremium(ctx).Return(&entity.GetUserDetailResponse{IsPremium: true}, nil)
		mockRepository.EXPECT().CheckUserSwipped(ctx, *request).Return(false, nil)
		mockRepository.EXPECT().BeginTx(ctx).Return(tx, nil)
		mockRepository.EXPECT().CreateSwipe(ctx, tx, *request).Return(nil)

		// Only expect the CreateMatch to be called once when SwipeType is "Like"
		mockRepository.EXPECT().CreateMatch(ctx, tx, *request).Return(nil).Times(1) // Adjusted Times here

		mockRepository.EXPECT().CommitTx(ctx, tx).Return(nil)

		err := svc.Swipe(ctx, request)

		assert.Nil(t, err)
	})

	t.Run("Successful swipe without match", func(t *testing.T) {
		request := &entity.SwipeRequest{
			TargetUserId: 2,
			SwipeType:    constant.Pass,
		}

		tx := &sql.Tx{} // Mock transaction object

		// Mock the complete flow for a successful swipe without a match
		mockRepository.EXPECT().CheckPremium(ctx).Return(&entity.GetUserDetailResponse{IsPremium: true}, nil)
		mockRepository.EXPECT().CheckUserSwipped(ctx, *request).Return(false, nil)
		mockRepository.EXPECT().BeginTx(ctx).Return(tx, nil)
		mockRepository.EXPECT().CreateSwipe(ctx, tx, *request).Return(nil)
		mockRepository.EXPECT().CommitTx(ctx, tx).Return(nil)

		err := svc.Swipe(ctx, request)

		assert.Nil(t, err)
	})

}
