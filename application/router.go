package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRouters() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/orders", loadOrderRouter)

	return router
}

func loadOrderRouter(route chi.Router) {
	route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
