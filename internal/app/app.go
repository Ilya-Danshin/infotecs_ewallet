package app

import "EWallet/internal/transport"

type App struct {
	s *transport.Server
}

func New() (*App, error) {
	s, err := transport.New()
	if err != nil {
		return nil, err
	}

	return &App{
		s: s,
	}, nil
}

func (a *App) Run() error {
	err := a.s.Run()
	if err != nil {
		return err
	}

	return nil
}
