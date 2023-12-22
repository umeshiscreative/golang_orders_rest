package application

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRouters(),
	}
	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: app.router,
	}

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("Failed to start the Server %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
