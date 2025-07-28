package handler

import (
	"apiRestFulSimpes/internal/model"
	"apiRestFulSimpes/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

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
	go sucessCreateTask(newID)
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

func sucessCreateTask(newID int) {
	log.Printf("INFO: Tarefa %d criada com sucesso em segundo plano.", newID)
}

func checkDB(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func checkES(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func CheckHealthHandler(w http.ResponseWriter, r *http.Request) {
	channelServices := make(chan bool)
	go checkDB(channelServices)
	go checkES(channelServices)

	for i := 0; i < 2; i++ {
		select {
		case <-channelServices:
			continue
		case <-time.After(3 * time.Second):
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func somaCounter(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	*counter++
	mu.Unlock()
	wg.Done()
}

func RaceHandler(w http.ResponseWriter, r *http.Request) {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var counter int
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go somaCounter(&counter, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(counter)
}
