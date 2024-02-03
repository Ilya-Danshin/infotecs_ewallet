package walletService

import (
	"EWallet/internal/config"
	"context"

	"github.com/gofrs/uuid"

	"EWallet/internal/database"
	"EWallet/internal/models"
)

type Service struct {
	cfg  config.Service
	repo database.WalletRepository
}

func New(repo database.WalletRepository, cfg config.Service) (*Service, error) {
	return &Service{
		cfg:  cfg,
		repo: repo}, nil
}

func (s *Service) CreateWallet(ctx context.Context) (*models.Wallet, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	err = s.repo.InsertWallet(ctx, uid, s.cfg.DefaultBalance)
	if err != nil {
		return nil, err
	}

	return &models.Wallet{
		Id:      uid,
		Balance: s.cfg.DefaultBalance,
	}, nil
}

func (s *Service) GetBalance(ctx context.Context, id uuid.UUID) (*models.Wallet, error) {
	w, err := s.repo.SelectWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (s *Service) IsWalletExist(ctx context.Context, id uuid.UUID) bool {
	w, err := s.repo.SelectWallet(ctx, id)
	if err != nil {
		return false
	}
	if w == nil {
		return false
	}

	return true
}

func (s *Service) CreateTransaction(ctx context.Context, from, to uuid.UUID, amount float32) error {
	return nil
}

func (s *Service) GetHistory(ctx context.Context, id uuid.UUID) ([]*models.Transaction, error) {
	return nil, nil
}
