package app

import (
	"errors"

	"github.com/google/uuid"
)

func (a AppService) DeleteClientsById(id uuid.UUID) error {
	exists, err := a.clientRepo.ClientExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("el cliente no existe")
	}

	return a.clientRepo.DeleteClientsById(id)
}
