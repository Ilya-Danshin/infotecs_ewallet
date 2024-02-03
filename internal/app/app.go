package app

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"EWallet/internal/config"
	"EWallet/internal/database/walletRepository"
	"EWallet/internal/transport"
)

type App struct {
	cfg *config.Config
	s   *transport.Server
}

func New() (*App, error) {
	a := &App{}
	var err error

	a.cfg, err = config.New()
	if err != nil {
		return nil, err
	}

	conn, err := a.connectToDB()
	if err != nil {
		return nil, err
	}

	repo, err := walletRepository.New(conn)
	if err != nil {
		return nil, err
	}
	// Test case
	uid, err := uuid.DefaultGenerator.NewV4()

	err = repo.InsertWallet(context.Background(), uid, 100.0)
	if err != nil {
		return nil, err
	}

	s, err := transport.New(a.cfg.Server)
	if err != nil {
		return nil, err
	}

	a.s = s

	return a, nil
}

func (a *App) Run() error {
	err := a.s.Run()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) connectToDB() (*pgxpool.Pool, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		a.cfg.Database.Host, a.cfg.Database.User, a.cfg.Database.Password, a.cfg.Database.DatabaseName, a.cfg.Database.Port)

	conn, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
