package app

import (
	"EWallet/internal/config"
	"EWallet/internal/transport"
)

type App struct {
	cfg *config.Config
	s   *transport.Server
}

func New() (*App, error) {
	a := &App{}

	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	a.cfg = cfg

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
