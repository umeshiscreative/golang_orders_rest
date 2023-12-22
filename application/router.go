package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/umeshiscreative/reset-service/handler"
)

func loadRouters() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(route chi.Router) {
	orderHandler := &handler.Order{}

	route.Get("/", orderHandler.GetOrder)
}
