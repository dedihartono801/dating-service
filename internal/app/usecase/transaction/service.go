package transaction

import (
	"context"
	repository "dating-service/internal/app/repository/transaction"
	"dating-service/internal/entity"
	"dating-service/pkg/identifier"
	"dating-service/pkg/validator"
)

type Service interface {
	Purchase(ctx context.Context, request *entity.TransactionRequest) error
	PaymentMethodList(ctx context.Context) ([]*entity.PaymentMethodResponse, error)
	PackageList(ctx context.Context) ([]*entity.PackageResponse, error)
}

type service struct {
	repository repository.TransactionRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewTransactionService(
	repository repository.TransactionRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}
