package main

import (
	"apiRestFulSimpes/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", handler.HelloHandler)
	r.Get("/tasks", handler.TasksListHandler)
	r.Get("/tasks/{ID}", handler.TaskHandler)
	http.ListenAndServe(":8080", r)
}
