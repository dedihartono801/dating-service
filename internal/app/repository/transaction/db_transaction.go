package repository

import (
	"context"
	"database/sql"
	"dating-service/pkg/helper"
)

func (r *transactionRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {

	tx, err := r.Database.BeginTx(ctx, nil)
	if err != nil {
		return nil, helper.HandleError(err)
	}
	return tx, nil
}

func (r *transactionRepository) RollbackTx(ctx context.Context, tx *sql.Tx) error {

	if err := tx.Rollback(); err != nil {
		return helper.HandleError(err)
	}
	return nil
}

func (r *transactionRepository) CommitTx(ctx context.Context, tx *sql.Tx) error {

	if err := tx.Commit(); err != nil {
		return helper.HandleError(err)
	}
	return nil
}
