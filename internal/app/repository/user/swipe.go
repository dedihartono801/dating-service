package repository

import (
	"context"
	"database/sql"
	"dating-service/internal/entity"
	"dating-service/pkg/helper"
)

func (r *userRepository) CreateSwipe(ctx context.Context, tx *sql.Tx, request entity.SwipeRequest) error {
	valueCtx := helper.GetValueContext(ctx)
	query := `INSERT INTO "swipe" (
	swipper_user_id,
    target_user_id, 
    swipe_type) 
    VALUES ($1, $2, $3)`

	_, err := tx.ExecContext(ctx, query,
		valueCtx.UserId,
		request.TargetUserId,
		request.SwipeType,
	)

	if err != nil {
		return helper.HandleError(err)
	}

	return nil
}

func (r *userRepository) CreateMatch(ctx context.Context, tx *sql.Tx, request entity.SwipeRequest) error {
	valueCtx := helper.GetValueContext(ctx)
	query := `INSERT INTO "match" (
	user_id,
    matched_user_id) 
    VALUES ($1, $2)`

	_, err := tx.ExecContext(ctx, query,
		valueCtx.UserId,
		request.TargetUserId,
	)

	if err != nil {
		return helper.HandleError(err)
	}

	return nil
}
