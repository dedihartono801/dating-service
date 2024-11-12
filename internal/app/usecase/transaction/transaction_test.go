package transaction

import (
	"context"
	"database/sql"
	mock_repository "dating-service/internal/app/repository/transaction/mocks"
	"dating-service/internal/entity"
	"dating-service/pkg/constant"
	"dating-service/pkg/identifier"
	"dating-service/pkg/validator"
	"errors"
	"testing"

	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPackageList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.HeaderContext, entity.ValueContext{
		UserId:     "1",
		Gender:     "male",
		Email:      "contoh@gmail.com",
		IsVerified: true,
		IsPremium:  true,
	})

	mockRepository := mock_repository.NewMockTransactionRepository(ctl)
	feature := NewTransactionService(mockRepository, nil, nil)

	t.Run("Get List packages Data", func(t *testing.T) {
		mockPackage := &entity.PackageResponse{
			Id:    1,
			Name:  "Premium",
			Price: 100000,
		}
		mockResults := []*entity.PackageResponse{mockPackage}
		mockRepository.EXPECT().GetPackages(ctx).Return(mockResults, nil)

		expectResponse := []*entity.PackageResponse{mockPackage}

		resp, err := feature.PackageList(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expectResponse, resp)
	})

	t.Run("Get Error List packages", func(t *testing.T) {
		mockRepository.EXPECT().GetPackages(ctx).Return(nil, errors.New("error"))

		_, err := feature.PackageList(ctx)
		assert.NotNil(t, err)
	})
}

func TestPaymentMethodList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.HeaderContext, entity.ValueContext{
		UserId:     "1",
		Gender:     "male",
		Email:      "contoh@gmail.com",
		IsVerified: true,
		IsPremium:  true,
	})

	mockRepository := mock_repository.NewMockTransactionRepository(ctl)
	feature := NewTransactionService(mockRepository, nil, nil)

	t.Run("Get List Payment Method Data", func(t *testing.T) {
		mockPackage := &entity.PaymentMethodResponse{
			Id:   1,
			Name: "OVO",
		}
		mockResults := []*entity.PaymentMethodResponse{mockPackage}
		mockRepository.EXPECT().GetPaymentMethods(ctx).Return(mockResults, nil)

		expectResponse := []*entity.PaymentMethodResponse{mockPackage}

		resp, err := feature.PaymentMethodList(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expectResponse, resp)
	})

	t.Run("Get Error List Payment", func(t *testing.T) {
		mockRepository.EXPECT().GetPaymentMethods(ctx).Return(nil, errors.New("error"))

		_, err := feature.PaymentMethodList(ctx)
		assert.NotNil(t, err)
	})
}

func TestPurchase(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.HeaderContext, entity.ValueContext{
		UserId:     "1",
		Gender:     "male",
		Email:      "contoh@gmail.com",
		IsVerified: true,
		IsPremium:  true,
	})

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())

	mockRepository := mock_repository.NewMockTransactionRepository(ctl)
	feature := NewTransactionService(mockRepository, validator, identifier) // Adjust as needed

	t.Run("successful purchase for Premium package", func(t *testing.T) {
		request := &entity.TransactionRequest{
			PaymentMethodId: 1,
			Currency:        "IDR",
			Amount:          10000,
			PackageTypeId:   1,
		}

		tx := &sql.Tx{}

		// Set up expected calls and return values
		mockRepository.EXPECT().BeginTx(gomock.Any()).Return(tx, nil)
		mockRepository.EXPECT().CheckPaymentMethod(gomock.Any(), request.PaymentMethodId).Return(true, nil)
		mockRepository.EXPECT().GetPackageById(gomock.Any(), request.PackageTypeId).Return(entity.PackageType{
			Price: 10000,
			Name:  "premium",
		}, nil)
		mockRepository.EXPECT().CreateTransaction(gomock.Any(), gomock.Any(), *request).Return(1, nil)
		mockRepository.EXPECT().CreateSubscription(gomock.Any(), gomock.Any(), *request, 1).Return(nil)
		mockRepository.EXPECT().UpdateUserIsPremium(gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CommitTx(gomock.Any(), gomock.Any()).Return(nil)

		// Execute the test
		err := feature.Purchase(ctx, request)
		assert.Nil(t, err)
	})

	t.Run("failed purchase for not enough amount", func(t *testing.T) {
		request := &entity.TransactionRequest{
			PaymentMethodId: 1,
			Currency:        "IDR",
			Amount:          5000,
			PackageTypeId:   1,
		}

		tx := &sql.Tx{}

		// Set up expected calls and return values
		mockRepository.EXPECT().BeginTx(gomock.Any()).Return(tx, nil)
		mockRepository.EXPECT().CheckPaymentMethod(gomock.Any(), request.PaymentMethodId).Return(true, nil)
		mockRepository.EXPECT().GetPackageById(gomock.Any(), request.PackageTypeId).Return(entity.PackageType{
			Price: 10000,
			Name:  "premium",
		}, nil)

		// Execute the test
		err := feature.Purchase(ctx, request)
		assert.NotNil(t, err)
	})

	t.Run("failed purchase for Amount is too much", func(t *testing.T) {
		request := &entity.TransactionRequest{
			PaymentMethodId: 1,
			Currency:        "IDR",
			Amount:          20000,
			PackageTypeId:   1,
		}

		tx := &sql.Tx{}

		// Set up expected calls and return values
		mockRepository.EXPECT().BeginTx(gomock.Any()).Return(tx, nil)
		mockRepository.EXPECT().CheckPaymentMethod(gomock.Any(), request.PaymentMethodId).Return(true, nil)
		mockRepository.EXPECT().GetPackageById(gomock.Any(), request.PackageTypeId).Return(entity.PackageType{
			Price: 10000,
			Name:  "premium",
		}, nil)

		// Execute the test
		err := feature.Purchase(ctx, request)
		assert.NotNil(t, err)
	})

	t.Run("successful purchase for Verified package", func(t *testing.T) {
		request := &entity.TransactionRequest{
			PaymentMethodId: 1,
			Currency:        "IDR",
			Amount:          10000,
			PackageTypeId:   1,
		}

		tx := &sql.Tx{}

		// Set up expected calls and return values
		mockRepository.EXPECT().BeginTx(gomock.Any()).Return(tx, nil)
		mockRepository.EXPECT().CheckPaymentMethod(gomock.Any(), request.PaymentMethodId).Return(true, nil)
		mockRepository.EXPECT().GetPackageById(gomock.Any(), request.PackageTypeId).Return(entity.PackageType{
			Price: 10000,
			Name:  "verified",
		}, nil)
		mockRepository.EXPECT().CreateTransaction(gomock.Any(), gomock.Any(), *request).Return(1, nil)
		mockRepository.EXPECT().CreateSubscription(gomock.Any(), gomock.Any(), *request, 1).Return(nil)
		mockRepository.EXPECT().UpdateUserIsVerified(gomock.Any(), gomock.Any()).Return(nil)
		mockRepository.EXPECT().CommitTx(gomock.Any(), gomock.Any()).Return(nil)

		// Execute the test
		err := feature.Purchase(ctx, request)
		assert.Nil(t, err)
	})

	t.Run("failed purchase for required package id", func(t *testing.T) {
		request := &entity.TransactionRequest{
			PaymentMethodId: 1,
			Currency:        "IDR",
			Amount:          10000,
		}

		// Execute the test
		err := feature.Purchase(ctx, request)
		assert.NotNil(t, err)
	})
}
