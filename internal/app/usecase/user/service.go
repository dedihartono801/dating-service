package user

import (
	"context"
	repository "dating-service/internal/app/repository/user"
	"dating-service/internal/entity"
	"dating-service/pkg/identifier"
	"dating-service/pkg/validator"
)

type Service interface {
	Signup(ctx context.Context, request *entity.SignupRequest) error
	Login(ctx context.Context, request *entity.LoginRequest) (*entity.LoginResponse, error)
	UserList(ctx context.Context) ([]*entity.GetUserListResponse, error)
	Swipe(ctx context.Context, request *entity.SwipeRequest) error
}

type service struct {
	repository repository.UserRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewUserService(
	repository repository.UserRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}
