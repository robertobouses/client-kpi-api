package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClientRepository interface {
	QueryAllClients() ([]Client, error)
	QueryClientById(uuid.UUID) (Client, error)
	InsertClients(req Client) error
	UpdateClientById(ctx *gin.Context, id uuid.UUID, req UpdateClientRequest) error
	DeleteClientsById(id uuid.UUID) error
	ClientExistsById(id uuid.UUID) (bool, error)
}

func NewApp(clientRepository ClientRepository) AppService {
	return AppService{
		clientRepo: clientRepository,
	}
}

type AppService struct {
	clientRepo ClientRepository
}
