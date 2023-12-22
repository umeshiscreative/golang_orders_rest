package application

import (
	"fmt"
	"net/http"
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

func (app *App) Start() error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: app.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		e := fmt.Errorf("Failed to start the Server !!!")
		return e
	}
	return nil
}
