package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

func (r *repository) ClientExistsById(id uuid.UUID) (bool, error) {
	var exists bool
	err := r.clientExistsById.QueryRow(id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}
