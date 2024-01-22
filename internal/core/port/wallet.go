package port

import (
	"context"
	"github.com/google/uuid"

	"EWallet/internal/core/domain"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error)
	GetWalletByID(ctx context.Context, id uuid.UUID) (*domain.Wallet, error)
}

type WalletService interface {
	Create(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error)
	GetWallet(ctx context.Context, id uuid.UUID) (*domain.Wallet, error)
}
