package repository

import (
	"github.com/google/uuid"
)

func (r *repository) DeleteClientsById(id uuid.UUID) error {
	_, err := r.deleteClientsById.Exec(id)
	return err
}
