package app

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateClientRequest struct {
	Name     *string
	LastName *string
	Email    *string
	Age      *int
	Birthday *string
}

func (a AppService) UpdateClientById(ctx *gin.Context, id uuid.UUID, req UpdateClientRequest) error {
	if req.Name == nil && req.LastName == nil && req.Email == nil && req.Age == nil && req.Birthday == nil {
		return errors.New("no hay datos para actualizar")
	}

	err := a.clientRepo.UpdateClientById(ctx, id, req)
	if err != nil {
		return err
	}

	return nil
}
