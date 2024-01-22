package service

import (
	"EWallet/internal/core/domain"
	"EWallet/internal/core/port"
	"context"
	"github.com/google/uuid"
)

type TransactionService struct {
	repo port.TransactionRepository
}

func NewTransactionService(repo port.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo,
	}
}

func (ts *TransactionService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	_, err := ts.repo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (ts *TransactionService) ListTransactions(ctx context.Context, id uuid.UUID) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	transactions, err := ts.repo.ListTransactions(ctx, id)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
