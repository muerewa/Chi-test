package main

import (
	"Chi-test/internal/config"
	"Chi-test/internal/domain/service"
	"Chi-test/internal/repository/postgresql"
	"Chi-test/internal/rest/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()
	repo, err := postgresql.New(cfg.StoragePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	taskService := service.NewTaskService(repo, log)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Post("/todo", taskHandler.CreateTask)
	r.Route("/todo/{id}", func(r chi.Router) {
		r.Get("/", taskHandler.GetTask)
		r.Delete("/", taskHandler.DeleteTask)
		r.Put("/", taskHandler.PutTask)
	})
	http.ListenAndServe(":8080", r)
}
