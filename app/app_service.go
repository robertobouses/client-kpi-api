package app

import "github.com/google/uuid"

type ClientRepository interface {
	QueryAllClients() ([]Client, error)
	QueryClientById(uuid.UUID) (Client, error)
}

func NewApp(clientRepository ClientRepository) AppService {
	return AppService{
		clientRepo: clientRepository,
	}
}

type AppService struct {
	clientRepo ClientRepository
}
