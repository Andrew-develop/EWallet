package port

import (
	"context"
	"github.com/google/uuid"

	"EWallet/internal/core/domain"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error)
	ListTransactions(ctx context.Context, id uuid.UUID) ([]domain.Transaction, error)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error)
	ListTransactions(ctx context.Context, id uuid.UUID) ([]domain.Transaction, error)
}
