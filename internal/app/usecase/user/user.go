package user

import (
	"context"
	"dating-service/cmd/middleware"
	"dating-service/internal/entity"
	"dating-service/pkg/constant"
	"dating-service/pkg/helper"
	"net/http"
	"time"
)

func (s *service) Swipe(ctx context.Context, request *entity.SwipeRequest) error {

	if err := s.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return err
	}

	premium, err := s.repository.CheckPremium(ctx)
	if err != nil {
		return err
	}

	if !premium.IsPremium {
		swipped, err := s.repository.CountUserSwipped(ctx)
		if err != nil {
			return err
		}
		if swipped > 10 {
			return helper.Error(http.StatusConflict, constant.ErrorAlreadySwiped, nil)
		}
	}

	exist, err := s.repository.CheckUserSwipped(ctx, *request)
	if err != nil {
		return err
	}

	if exist {
		err = helper.Error(http.StatusConflict, constant.ErrorAlreadySwipedPerson, nil)
		return err
	}

	tx, err := s.repository.BeginTx(ctx)
	if err != nil {
		return err
	}

	if err := s.repository.CreateSwipe(ctx, tx, *request); err != nil {
		return err
	}

	if request.SwipeType == constant.Like {
		if err := s.repository.CreateMatch(ctx, tx, *request); err != nil {
			s.repository.RollbackTx(ctx, tx)
			return err
		}
	}

	if err := s.repository.CommitTx(ctx, tx); err != nil {
		return err
	}

	return nil
}

func (s *service) Signup(ctx context.Context, request *entity.SignupRequest) error {

	if err := s.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return err
	}

	if request.Password != "" {
		request.Password = helper.EncryptPassword(request.Password)
	}

	check, err := s.repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		statusCode, _, _ := helper.TrimMesssage(err)
		if statusCode != http.StatusNotFound {
			return err
		}
	}

	if check != nil {
		err = helper.Error(http.StatusConflict, constant.ErrorEmailAlreadyExists, nil)
		return err
	}

	if err := s.repository.Signup(ctx, *request); err != nil {
		return err
	}

	return nil
}

func (s *service) Login(ctx context.Context, request *entity.LoginRequest) (*entity.LoginResponse, error) {

	if err := s.validator.Validate(request); err != nil {
		err = helper.Error(http.StatusBadRequest, err.Error(), err)
		return nil, err
	}

	user, err := s.repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		statusCode, _, _ := helper.TrimMesssage(err)
		if statusCode == http.StatusNotFound {
			err = helper.Error(http.StatusUnauthorized, constant.ErrorEmailNotFound, err)
			return nil, err
		}
		return nil, err
	}

	if user.Password != helper.EncryptPassword(request.Password) {
		err = helper.Error(http.StatusUnauthorized, constant.ErrorPasswordWrong, nil)
		return nil, err
	}

	expirationTime := time.Now().Add(time.Hour * time.Duration(24))
	token, err := middleware.GenerateToken(user.Id, user.Email, user.Gender, user.IsVerified, user.IsPremium)
	if err != nil {
		return nil, err
	}

	resp := entity.LoginResponse{
		Token:     token,
		ExpiredAt: expirationTime.Format(time.RFC3339),
	}

	return &resp, nil
}

func (s *service) UserList(ctx context.Context) ([]*entity.GetUserListResponse, error) {
	premium, err := s.repository.CheckPremium(ctx)
	if err != nil {
		return nil, err
	}

	if !premium.IsPremium {
		swipped, err := s.repository.CountUserSwipped(ctx)
		if err != nil {
			return nil, err
		}

		if swipped > 10 {
			return nil, helper.Error(http.StatusConflict, constant.ErrorAlreadySwiped, nil)
		}

	}
	user, err := s.repository.GetUserList(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
