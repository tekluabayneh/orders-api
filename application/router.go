package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tekluabayneh/orders-api/handler"
	"net/http"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Get("/", orderHandler.List)
	router.Post("/", orderHandler.Create)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateById)
	router.Delete("/{id}", orderHandler.DelteById)
}
