package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
}

// SystemRoot handller
func SystemRoot(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Book reader service"})
}

// Router defines all routers
func Router() *chi.Mux {
	router.Get("/", SystemRoot)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/users", userRoutes())
	})

	return router
}

func userRoutes() http.Handler {
	route := chi.NewRouter()
	route.Group(func(r chi.Router) {
		r.Post("/", NewReaderAPI().CreateUser)
		r.Get("/{user_id:[0-9]+}", NewReaderAPI().GetUser)
	})
	return route
}
