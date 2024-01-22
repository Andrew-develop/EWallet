package repository

import (
	"EWallet/internal/adapter/storage/postgres"
	"EWallet/internal/core/domain"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type WalletRepository struct {
	db *postgres.DB
}

func NewWalletRepository(db *postgres.DB) *WalletRepository {
	return &WalletRepository{
		db,
	}
}

func (wr *WalletRepository) CreateWallet(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error) {
	q := "INSERT INTO wallets (id, balance) VALUES ($1, $2) RETURNING *"

	err := wr.db.QueryRow(ctx, q, wallet.Id, wallet.Balance).Scan(
		&wallet.Id,
		&wallet.Balance,
	)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (wr *WalletRepository) GetWalletByID(ctx context.Context, id uuid.UUID) (*domain.Wallet, error) {
	q := "SELECT * FROM wallets WHERE id = $1 limit 1"

	var wallet domain.Wallet
	err := wr.db.QueryRow(ctx, q, id).Scan(
		&wallet.Id,
		&wallet.Balance,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &wallet, nil
}
