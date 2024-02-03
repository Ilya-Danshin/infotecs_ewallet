package services

import (
	"EWallet/internal/models"
	"context"

	"github.com/gofrs/uuid"
)

type WalletService interface {
	CreateWallet(ctx context.Context) (*models.Wallet, error)
	GetBalance(ctx context.Context, id uuid.UUID) (*models.Wallet, error)
	IsWalletExist(ctx context.Context, id uuid.UUID) bool

	CreateTransaction(ctx context.Context, from, to uuid.UUID, amount float32) error
	GetHistory(ctx context.Context, id uuid.UUID) ([]*models.Transaction, error)
}
