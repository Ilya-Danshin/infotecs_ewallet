package database

import (
	"context"
	"time"

	"github.com/gofrs/uuid"

	"EWallet/internal/models"
)

type WalletRepository interface {
	InsertWallet(ctx context.Context, walletId uuid.UUID, balance float32) error
	SelectWallet(ctx context.Context, walletId uuid.UUID) (*models.Wallet, error)
	UpdateWalletBalance(ctx context.Context, from, to uuid.UUID, amount float32) error
	DeleteWallet(ctx context.Context, walletId uuid.UUID) error

	InsertTransaction(ctx context.Context, time time.Time, from uuid.UUID, to uuid.UUID, money float32) error
	SelectTransactionsByWallet(ctx context.Context, walletId uuid.UUID) ([]*models.Transaction, error)
}
