package repository

import (
	"context"
	"database/sql"
	"dating-service/database"
	"dating-service/internal/entity"
	"dating-service/pkg/helper"
)

type UserRepository interface {
	Signup(ctx context.Context, request entity.SignupRequest) error
	GetUserByEmail(ctx context.Context, email string) (*entity.GetUserDetailResponse, error)
	GetUserList(ctx context.Context) ([]*entity.GetUserListResponse, error)
	CreateSwipe(ctx context.Context, tx *sql.Tx, request entity.SwipeRequest) error
	CreateMatch(ctx context.Context, tx *sql.Tx, request entity.SwipeRequest) error
	CheckUserSwipped(ctx context.Context, request entity.SwipeRequest) (resp bool, err error)
	CountUserSwipped(ctx context.Context) (int, error)
	CheckPremium(ctx context.Context) (*entity.GetUserDetailResponse, error)
	BeginTx(ctx context.Context) (*sql.Tx, error)
	RollbackTx(ctx context.Context, tx *sql.Tx) error
	CommitTx(ctx context.Context, tx *sql.Tx) error
}

type userRepository struct {
	Database *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		Database: db,
	}
}

func (r *userRepository) Signup(ctx context.Context, request entity.SignupRequest) error {
	query := `INSERT INTO "user" (
    email, 
    password, 
    first_name, 
    last_name, 
    gender, 
    age, 
    date_of_birth, 
    bio, 
    location) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.Database.DB.ExecContext(ctx, query,
		request.Email,
		request.Password,
		request.FirstName,
		request.LastName,
		request.Gender,
		request.Age,
		request.DateOfBirth,
		request.Bio,
		request.Location,
	)

	if err != nil {
		return helper.HandleError(err)
	}

	return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.GetUserDetailResponse, error) {
	query := `select id, email, password, gender, is_verified, is_premium from "user" where email = $1`

	var response entity.GetUserDetailResponse
	if err := r.Database.DB.GetContext(ctx, &response, query, email); err != nil {
		return nil, helper.HandleError(err)
	}

	return &response, nil
}

func (r *userRepository) CheckUserSwipped(ctx context.Context, request entity.SwipeRequest) (resp bool, err error) {
	valueCtx := helper.GetValueContext(ctx)
	query := `SELECT EXISTS (SELECT 1 FROM "swipe" WHERE created_at::date = CURRENT_DATE AND swipper_user_id = $1 AND target_user_id = $2)`

	if err = r.Database.GetContext(ctx, &resp, query, valueCtx.UserId, request.TargetUserId); err != nil {
		err = helper.HandleError(err)
	}

	return
}

func (r *userRepository) GetUserList(ctx context.Context) ([]*entity.GetUserListResponse, error) {
	valueCtx := helper.GetValueContext(ctx)
	query := `SELECT 
				u.id,
				u.first_name,
				u.last_name,
				u.profile_picture,
				u.bio,
				u.location,
				u.age,
				u.Gender,
				u.is_verified,
				u.is_premium
			FROM "user" u
			LEFT JOIN swipe s ON u.id = s.target_user_id AND s.swipper_user_id = $1 AND s.created_at::date = CURRENT_DATE
			WHERE s.swipper_user_id IS NULL AND u.gender != $2
			ORDER BY RANDOM()
			LIMIT 10`

	var response []*entity.GetUserListResponse
	if err := r.Database.DB.SelectContext(ctx, &response, query, valueCtx.UserId, valueCtx.Gender); err != nil {
		return nil, helper.HandleError(err)
	}

	return response, nil
}

func (r *userRepository) CountUserSwipped(ctx context.Context) (int, error) {
	valueCtx := helper.GetValueContext(ctx)
	query := `SELECT COUNT(*) FROM "swipe" WHERE created_at::date = CURRENT_DATE AND swipper_user_id = $1`

	var count int
	err := r.Database.QueryRowContext(ctx, query, valueCtx.UserId).Scan(&count)
	if err != nil {
		return 0, helper.HandleError(err)
	}

	return count, nil
}

func (r *userRepository) CheckPremium(ctx context.Context) (*entity.GetUserDetailResponse, error) {
	valueCtx := helper.GetValueContext(ctx)
	query := `select is_premium from "user" where id = $1`

	var response entity.GetUserDetailResponse
	if err := r.Database.DB.GetContext(ctx, &response, query, valueCtx.UserId); err != nil {
		return nil, helper.HandleError(err)
	}

	return &response, nil
}
