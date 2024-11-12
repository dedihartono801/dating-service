package transaction

import (
	"context"
	"dating-service/internal/entity"
	"dating-service/pkg/constant"
	"dating-service/pkg/helper"
	"net/http"
)

func (s *service) Purchase(ctx context.Context, request *entity.TransactionRequest) error {

	if err := s.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return err
	}

	tx, err := s.repository.BeginTx(ctx)
	if err != nil {
		return err
	}

	exist, err := s.repository.CheckPaymentMethod(ctx, request.PaymentMethodId)
	if err != nil {
		return err
	}

	if !exist {
		err = helper.Error(http.StatusConflict, constant.ErrorPaymentMethodNotFound, nil)
		return err
	}

	dtPackage, err := s.repository.GetPackageById(ctx, request.PackageTypeId)
	if err != nil {
		return err
	}

	if dtPackage.Price > request.Amount {
		err = helper.Error(http.StatusConflict, constant.ErrorAmountNotEnough, nil)
		return err
	}

	if request.Amount > dtPackage.Price {
		err = helper.Error(http.StatusConflict, constant.ErrorAmountTooMuch, nil)
		return err
	}

	id, err := s.repository.CreateTransaction(ctx, tx, *request)
	if err != nil {
		return err
	}

	if err := s.repository.CreateSubscription(ctx, tx, *request, id); err != nil {
		s.repository.RollbackTx(ctx, tx)
		return err
	}

	if dtPackage.Name == constant.Premium {
		err = s.repository.UpdateUserIsPremium(ctx, tx)
		if err != nil {
			s.repository.RollbackTx(ctx, tx)
			return err
		}
	} else {
		err = s.repository.UpdateUserIsVerified(ctx, tx)
		if err != nil {
			s.repository.RollbackTx(ctx, tx)
			return err
		}
	}

	if err := s.repository.CommitTx(ctx, tx); err != nil {
		return err
	}

	return nil
}

func (s *service) PackageList(ctx context.Context) ([]*entity.PackageResponse, error) {
	user, err := s.repository.GetPackages(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) PaymentMethodList(ctx context.Context) ([]*entity.PaymentMethodResponse, error) {
	user, err := s.repository.GetPaymentMethods(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
