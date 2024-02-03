package app

import (
	"EWallet/internal/transport"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"EWallet/internal/config"
	"EWallet/internal/database/walletRepository"
	"EWallet/internal/services/walletService"
	"EWallet/internal/transport/echoServer"
)

type App struct {
	cfg *config.Config
	s   transport.WalletTransport
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

	service, err := walletService.New(repo, a.cfg.Service)
	if err != nil {
		return nil, err
	}

	a.s, err = echoServer.New(service, a.cfg.Server)
	if err != nil {
		return nil, err
	}

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
