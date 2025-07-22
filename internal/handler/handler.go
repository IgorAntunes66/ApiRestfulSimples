package handler

import (
	"apiRestFulSimpes/internal/model"
	"apiRestFulSimpes/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, Gopher!")
}

func (s *ApiServer) TasksListHandler(w http.ResponseWriter, r *http.Request) {
	tarefas, err := s.repo.FindAllTask()
	if err != nil {
		http.Error(w, "Erro ao receber os dados do banco de dados", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tarefas)
	if err != nil {
		http.Error(w, "Erro ao converter a slice para json", http.StatusInternalServerError)
	}
}

func (s *ApiServer) TaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Erro ao converter o ID da URL para inteiro? %v", err)
	}

	tarefa, err := s.repo.FindTaskByID(idInt)
	if err != nil {
		http.Error(w, "Erro ao receber os dados do banco de dados", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tarefa)
	if err != nil {
		http.Error(w, "Erro ao transformar os dados em json", http.StatusInternalServerError)
	}
}

func (s *ApiServer) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tarefa model.Task
	err := json.NewDecoder(r.Body).Decode(&tarefa)
	if err != nil {
		http.Error(w, "Erro ao decodificar a requisição", http.StatusBadRequest)
	}

	newID, err := s.repo.InsertTask(tarefa)
	tarefa.ID = newID
	if err != nil {
		http.Error(w, "Erro ao inserir tarefa no banco de dados", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tarefa)
}

func (s *ApiServer) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tarefa model.Task

	id := chi.URLParam(r, "ID")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter o ID para inteiro", http.StatusInternalServerError)
	}

	err = json.NewDecoder(r.Body).Decode(&tarefa)
	if err != nil {
		http.Error(w, "Erro ao decodificar a requisição", http.StatusBadRequest)
	}

	err = s.repo.UpdateTask(idInt, tarefa)
	if err == repository.ErrRegistroNaoEncontrado {
		http.Error(w, repository.ErrRegistroNaoEncontrado.Error(), http.StatusNotFound)
	} else if err != nil {
		http.Error(w, "Erro ao atualizar a tarefa no banco de dados", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *ApiServer) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro ao converter o ID para inteiro", http.StatusInternalServerError)
	}

	err = s.repo.DeleteTask(idInt)
	if err == repository.ErrRegistroNaoEncontrado {
		http.Error(w, repository.ErrRegistroNaoEncontrado.Error(), http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
