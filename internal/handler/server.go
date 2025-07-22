package handler

import (
	"apiRestFulSimpes/internal/repository"
)

type ApiServer struct {
	repo *repository.Repository
}

func NewApiServer(repo *repository.Repository) *ApiServer {
	return &ApiServer{
		repo: repo,
	}
}
