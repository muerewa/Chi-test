package main

import (
	"Chi-test/internal/rest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Route("/todo/{id}", func(r chi.Router) {
		r.Get("/", handlers.GetTask)
		r.Post("/", handlers.CreateTask)
		r.Delete("/", handlers.DeleteTask)
		r.Put("/", handlers.PutTask)
	})
	http.ListenAndServe(":8080", r)
}
