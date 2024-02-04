package walletService

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"

	"EWallet/internal/config"
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
	if amount < 0.0 {
		return errors.New("amount less than 0")
	}

	if !s.IsWalletExist(ctx, to) {
		return errors.New("\"to\" doesn't exist")
	}

	fromWallet, err := s.GetBalance(ctx, from)
	if err != nil {
		return err
	}

	if fromWallet.Balance < amount {
		return errors.New("not enough money")
	}

	err = s.repo.UpdateWalletBalance(ctx, from, to, amount)
	if err != nil {
		return err
	}

	err = s.logTransaction(ctx, from, to, amount)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetHistory(ctx context.Context, id uuid.UUID) ([]*models.Transaction, error) {
	t, err := s.repo.SelectTransactionsByWallet(ctx, id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *Service) logTransaction(ctx context.Context, from, to uuid.UUID, amount float32) error {
	err := s.repo.InsertTransaction(ctx, time.Now(), from, to, amount)
	if err != nil {
		return err
	}

	return nil
}
