package repository

import (
	"EWallet/internal/adapter/storage/postgres"
	"EWallet/internal/core/domain"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	db *postgres.DB
}

func NewTransactionRepository(db *postgres.DB) *TransactionRepository {
	return &TransactionRepository{
		db,
	}
}

func (tr *TransactionRepository) CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	tx, err := tr.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	q := "UPDATE wallets SET balance = balance - $1 WHERE id = $2"
	_, err = tx.Exec(ctx, q, transaction.Amount, transaction.From)
	if err != nil {
		return nil, err
	}

	q = "UPDATE wallets SET balance = balance + $1 WHERE id = $2"
	_, err = tx.Exec(ctx, q, transaction.Amount, transaction.To)
	if err != nil {
		return nil, err
	}

	q = "INSERT INTO transactions (id, from_wallet, to_wallet, amount, date_time) VALUES ($1, $2, $3, $4, $5) RETURNING *"
	err = tx.QueryRow(ctx, q, transaction.Id, transaction.From, transaction.To, transaction.Amount, transaction.DateTime).Scan(
		&transaction.Id,
		&transaction.From,
		&transaction.To,
		&transaction.Amount,
		&transaction.DateTime)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (tr *TransactionRepository) ListTransactions(ctx context.Context, id uuid.UUID) ([]domain.Transaction, error) {
	q := "SELECT date_time, from_wallet, to_wallet, amount FROM transactions WHERE from_wallet = $1 OR to_wallet = $1"

	var transaction domain.Transaction
	var transactions []domain.Transaction

	rows, err := tr.db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&transaction.DateTime,
			&transaction.From,
			&transaction.To,
			&transaction.Amount,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
