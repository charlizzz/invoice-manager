package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all functions to execute queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx execute a function within a database transaction (unexported)
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := New(tx)
	queryErr := fn(query)
	if queryErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			//combine the 2 errors in one
			return fmt.Errorf("transaction error: %v, rollback error: %v", queryErr, rollbackErr)
		}

		return queryErr
	}

	return tx.Commit()
}

type TransactionTxParams struct {
	InvoiceID  int32
	NewBalance int64
	UserID     int32
}

// TransactionTx performs an app transaction
func (store *Store) TransactionTx(ctx context.Context, arg TransactionTxParams) error {
	err := store.execTx(ctx, func(query *Queries) error {
		var err error

		_, err = query.UpdateUserBalance(ctx, UpdateUserBalanceParams{
			ID:      arg.UserID,
			Balance: arg.NewBalance,
		})
		if err != nil {
			return err
		}

		_, err = query.UpdateInvoiceStatus(ctx, arg.InvoiceID)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
