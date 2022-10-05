package app

import v1 "github.com/ValikoDorodnov/go_sample/internal/delivery/http/v1"

type App struct {
	Handler *v1.Handler
}

func NewApp(handler *v1.Handler) *App {
	return &App{
		Handler: handler,
	}
}
