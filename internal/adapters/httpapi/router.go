package httpapi

import (
	"net/http"

	"github.com/edmiltonVinicius/go-api-catalog/internal/adapters/httpapi/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handler *handlers.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Route("/products", func(r chi.Router) {
			r.Post("/", handler.CreateProduct)
			r.Get("/{id}", handler.GetProduct)
		})
	})

	return r
}
