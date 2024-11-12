package repository

import (
	"context"
	"database/sql"
	"dating-service/database"
	"dating-service/internal/entity"
	"dating-service/pkg/helper"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, tx *sql.Tx, request entity.TransactionRequest) (int, error)
	CreateSubscription(ctx context.Context, tx *sql.Tx, request entity.TransactionRequest, trxId int) error
	GetPackageById(ctx context.Context, id int) (resp entity.PackageType, err error)
	CheckPaymentMethod(ctx context.Context, id int) (resp bool, err error)
	UpdateUserIsVerified(ctx context.Context, tx *sql.Tx) error
	UpdateUserIsPremium(ctx context.Context, tx *sql.Tx) error
	CountUserSwipped(ctx context.Context) (int, error)
	GetPackages(ctx context.Context) ([]*entity.PackageResponse, error)
	GetPaymentMethods(ctx context.Context) ([]*entity.PaymentMethodResponse, error)
	BeginTx(ctx context.Context) (*sql.Tx, error)
	RollbackTx(ctx context.Context, tx *sql.Tx) error
	CommitTx(ctx context.Context, tx *sql.Tx) error
}

type transactionRepository struct {
	Database *database.Database
}

func NewTransactionRepository(db *database.Database) TransactionRepository {
	return &transactionRepository{
		Database: db,
	}
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, tx *sql.Tx, request entity.TransactionRequest) (int, error) {
	valueCtx := helper.GetValueContext(ctx)
	query := `INSERT INTO "transaction" (
	user_id,
	payment_method_id,
    currency, 
    amount,
	status) 
    VALUES ($1, $2, $3, $4, 'paid')
	RETURNING id`

	var trxId int
	err := tx.QueryRowContext(ctx, query,
		valueCtx.UserId,
		request.PaymentMethodId,
		request.Currency,
		request.Amount,
	).Scan(&trxId)

	if err != nil {
		return 0, helper.HandleError(err)
	}

	return int(trxId), nil
}

func (r *transactionRepository) CreateSubscription(ctx context.Context, tx *sql.Tx, request entity.TransactionRequest, trxId int) error {
	valueCtx := helper.GetValueContext(ctx)
	query := `INSERT INTO "subscription" (
	user_id,
	transaction_id,
    package_type_id,
	is_active,
	expires_at) 
    VALUES ($1, $2, $3, TRUE, NOW() + INTERVAL '1 MONTH')`

	_, err := tx.ExecContext(ctx, query,
		valueCtx.UserId,
		trxId,
		request.PackageTypeId,
	)

	if err != nil {
		return helper.HandleError(err)
	}

	return nil
}

func (r *transactionRepository) CheckPaymentMethod(ctx context.Context, id int) (resp bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM "payment_method" WHERE id = $1)`

	if err = r.Database.GetContext(ctx, &resp, query, id); err != nil {
		err = helper.HandleError(err)
	}

	return
}

func (r *transactionRepository) GetPackageById(ctx context.Context, id int) (resp entity.PackageType, err error) {
	query := `SELECT name, price FROM "package_type" WHERE id = $1`

	if err = r.Database.GetContext(ctx, &resp, query, id); err != nil {
		err = helper.HandleError(err)
	}

	return
}

func (r *transactionRepository) UpdateUserIsVerified(ctx context.Context, tx *sql.Tx) error {
	valueCtx := helper.GetValueContext(ctx)
	query := `UPDATE "user" 
	          SET is_verified = TRUE
	          WHERE id = $1`

	_, err := tx.ExecContext(ctx, query,
		valueCtx.UserId,
	)

	if err != nil {
		return helper.HandleError(err)
	}

	return nil
}

func (r *transactionRepository) UpdateUserIsPremium(ctx context.Context, tx *sql.Tx) error {
	valueCtx := helper.GetValueContext(ctx)
	query := `UPDATE "user" 
	          SET is_premium = TRUE
	          WHERE id = $1`

	_, err := tx.ExecContext(ctx, query,
		valueCtx.UserId,
	)

	if err != nil {
		return helper.HandleError(err)
	}

	return nil
}

func (r *transactionRepository) CountUserSwipped(ctx context.Context) (int, error) {
	valueCtx := helper.GetValueContext(ctx)
	query := `SELECT COUNT(*) FROM "swipe" WHERE created_at::date = CURRENT_DATE AND swipper_user_id = $1`

	var count int
	err := r.Database.QueryRowContext(ctx, query, valueCtx.UserId).Scan(&count)
	if err != nil {
		return 0, helper.HandleError(err)
	}

	return count, nil
}

func (r *transactionRepository) GetPackages(ctx context.Context) ([]*entity.PackageResponse, error) {
	query := `SELECT 
				id,
				name,
				price
			FROM "package_type"`

	var response []*entity.PackageResponse
	if err := r.Database.DB.SelectContext(ctx, &response, query); err != nil {
		return nil, helper.HandleError(err)
	}

	return response, nil
}

func (r *transactionRepository) GetPaymentMethods(ctx context.Context) ([]*entity.PaymentMethodResponse, error) {
	query := `SELECT 
				id,
				name
			FROM "payment_method"`

	var response []*entity.PaymentMethodResponse
	if err := r.Database.DB.SelectContext(ctx, &response, query); err != nil {
		return nil, helper.HandleError(err)
	}

	return response, nil
}
