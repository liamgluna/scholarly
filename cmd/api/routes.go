package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.MethodNotAllowed(app.methodNotAllowedResponse)
	router.NotFound(app.notFoundResponse)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Cruel World!"))
	})

	return router
}
