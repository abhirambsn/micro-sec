package app

import (
	"fmt"
	"context"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context, addr string) error {
	server := &http.Server{
		Addr: addr,
		Handler: a.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("Failed to Start Server: %w", err)
	}

	return nil
}

