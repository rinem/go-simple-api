package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"

	"github.com/rinem/go-simple-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Auth middleware
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
