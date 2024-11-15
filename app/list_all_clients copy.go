package app

import (
	"log"

	"github.com/google/uuid"
)

func (a AppService) ListClientById(id uuid.UUID) (Client, error) {
	client, err := a.clientRepo.QueryClientById(id)
	if err != nil {
		log.Println("Error al extraer GetLineup", err)
		return Client{}, err
	}
	return client, nil
}
