package app

import (
	"github.com/google/uuid"
)

func (a AppService) DeleteClientsById(id uuid.UUID) error {
	exists, err := a.clientRepo.ClientExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrClientNotFound
	}

	return a.clientRepo.DeleteClientsById(id)
}
