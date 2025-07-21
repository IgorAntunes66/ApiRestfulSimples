package main

import (
	"apiRestFulSimpes/internal/handler"
	"apiRestFulSimpes/internal/pkg"
	"apiRestFulSimpes/internal/repository"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	db, err := pkg.ConectaDB()
	if err != nil {
		log.Fatalf("Erro ao iniciar o banco de dados: %v", err)
	}

	repo := repository.NewRepository(db)
	apiServer := handler.NewApiServer(*repo)

	r := chi.NewRouter()
	r.Get("/", handler.HelloHandler)
	r.Get("/tasks", apiServer.TasksListHandler)
	r.Get("/tasks/{ID}", apiServer.TaskHandler)
	http.ListenAndServe(":8080", r)
}
