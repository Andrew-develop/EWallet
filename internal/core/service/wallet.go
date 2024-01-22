package service

import (
	"EWallet/internal/core/domain"
	"EWallet/internal/core/port"
	"context"
	"github.com/google/uuid"
)

type WalletService struct {
	repo port.WalletRepository
}

func NewWalletService(repo port.WalletRepository) *WalletService {
	return &WalletService{
		repo,
	}
}

func (ws *WalletService) Create(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error) {
	_, err := ws.repo.CreateWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (ws *WalletService) GetWallet(ctx context.Context, id uuid.UUID) (*domain.Wallet, error) {
	var wallet *domain.Wallet

	wallet, err := ws.repo.GetWalletByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
